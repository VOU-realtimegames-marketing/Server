package gapi

import (
	db "VOU-Server/db/sqlc"
	"VOU-Server/proto/gen"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetQuestion(stream grpc.BidiStreamingServer[gen.GetQuestionRequest, gen.GetQuestionResponse]) error {
	for {
		streamCtx := stream.Context()
		err := contextError(streamCtx)
		if err != nil {
			return err
		}

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("no more data")
			break
		}
		if err != nil {
			return logError(status.Errorf(codes.Unknown, "cannot receive stream request: %v", err))
		}

		questionNum := req.GetQuestionNum()
		eventId := req.GetEventId()

		event, err := server.store.GetEventById(streamCtx, eventId)
		if err != nil {
			if errors.Is(err, db.ErrRecordNotFound) {
				return status.Errorf(codes.NotFound, "event not found")
			}
			return status.Errorf(codes.Internal, "failed to get event: %s", err)
		}

		previousQuestionDuration := TOTAL_QUESTION_TIME * (questionNum - 1) // since start of event
		fmt.Println("previousQuestionDuration-get", previousQuestionDuration)

		// CHECK IF QUESTION NUM IS VALID
		time1 := event.StartTime.Add(
			time.Duration(previousQuestionDuration) * time.Second)
		time2 := time.Now()
		if time1.After(time2) {
			fmt.Println("2.time1", time1)
			fmt.Println("2.time2", time2)

			return status.Errorf(codes.InvalidArgument, "Your question number exceeds current one")
		}

		// GET QUIZ RECORD FROM DB
		quizzesDB, err := server.store.GetQuizzesByEventId(streamCtx, eventId)
		if err != nil {
			if errors.Is(err, db.ErrRecordNotFound) {
				return status.Errorf(codes.NotFound, "no quiz found")
			}
			return status.Errorf(codes.Internal, "failed to get quiz: %s", err)
		}

		if questionNum > quizzesDB.QuizNum {
			return status.Errorf(codes.InvalidArgument, "question number exceeds total questions")
		}

		var quizzes []struct {
			Question string   `json:"question"`
			Answer   string   `json:"answer"`
			Options  []string `json:"options"`
		}
		err = json.Unmarshal(quizzesDB.Content, &quizzes)
		if err != nil {
			return status.Errorf(codes.Internal, "failed to unmarshal quiz content: %s", err)
		}

		res := &gen.GetQuestionResponse{
			Question: quizzes[questionNum-1].Question,
			Options:  quizzes[questionNum-1].Options,
		}

		err = stream.Send(res)
		if err != nil {
			return logError(status.Errorf(codes.Unknown, "cannot send stream response: %v", err))
		}
	}
	return nil
}
