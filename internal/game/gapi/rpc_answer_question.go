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

func (server *Server) AnswerQuestion(stream grpc.BidiStreamingServer[gen.AnswerQuestionRequest, gen.AnswerQuestionResponse]) error {
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

		eventId := req.GetEventId()
		questionNum := req.GetQuestionNum()
		answer := req.GetAnswer()
		username := req.GetUsername()

		event, err := server.store.GetEventById(streamCtx, eventId)
		if err != nil {
			if errors.Is(err, db.ErrRecordNotFound) {
				return status.Errorf(codes.NotFound, "event not found")
			}
			return status.Errorf(codes.Internal, "failed to get event: %s", err)
		}

		previousQuestionDuration := TOTAL_QUESTION_TIME * (questionNum - 1) // since start of event
		fmt.Println("previousQuestionDuration-answer", previousQuestionDuration)

		// CHECK IF QUESTION NUM IS VALID
		time1 := event.StartTime.Add(
			time.Duration(previousQuestionDuration+QUESTION_DISPLAY_TIME) * time.Second)
		time2 := time.Now().Add(-1 * TIME_TOLERANCE)
		if time1.Before(time2) && answer != "" {
			fmt.Println("1.time1", time1)
			fmt.Println("1.time2", time2)

			return status.Errorf(codes.InvalidArgument, "question time is over")
		}

		time1 = event.StartTime.Add(
			time.Duration(previousQuestionDuration) * time.Second)
		time2 = time.Now()
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

		var quizzes []struct {
			Question string   `json:"question"`
			Answer   string   `json:"answer"`
			Options  []string `json:"options"`
		}
		err = json.Unmarshal(quizzesDB.Content, &quizzes)
		if err != nil {
			return status.Errorf(codes.Internal, "failed to unmarshal quiz content: %s", err)
		}

		res := &gen.AnswerQuestionResponse{
			CorrectAnswer: quizzes[questionNum-1].Answer,
			IsWinner:      false,
			IsCorrect:     false,
		}

		var userAnswer db.UserAnswer
		// CORRECT ANSWER
		if answer == quizzes[questionNum-1].Answer {
			userAnswer, err = server.store.UpdateUserAnswer(streamCtx, db.UpdateUserAnswerParams{
				EventID:  eventId,
				Username: username,
			})

			if err != nil {
				if errors.Is(err, db.ErrRecordNotFound) {
					_, err = server.store.CreateUserAnswer(streamCtx, db.CreateUserAnswerParams{
						EventID:  eventId,
						Username: username,
					})
					if err != nil {
						return status.Errorf(codes.Internal, "failed to create user answer: %s", err)
					}
				} else {
					return status.Errorf(codes.Internal, "failed to update user answer: %s", err)
				}
			}

			res.IsCorrect = true
		}

		if questionNum == quizzesDB.QuizNum && userAnswer.NumCorrect == quizzesDB.QuizNum {
			// Last question, check if user win the game
			res.IsWinner = true
			voucher, err := server.store.CreateVoucher(streamCtx, db.CreateVoucherParams{
				EventID:   event.ID,
				ExpiresAt: event.EndTime,
			})
			if err != nil {
				return status.Errorf(codes.Internal, "failed to create voucher: %s", err)
			}

			_, err = server.store.CreateVoucherOwner(streamCtx, db.CreateVoucherOwnerParams{
				VoucherID: voucher.ID,
				Username:  username,
			})
			if err != nil {
				return status.Errorf(codes.Internal, "failed to create voucher owner: %s", err)
			}
		}

		err = stream.Send(res)
		if err != nil {
			return logError(status.Errorf(codes.Unknown, "cannot send stream response: %v", err))
		}
	}
	return nil
}
