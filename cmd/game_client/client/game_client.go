package client

import (
	"VOU-Server/proto/gen"
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

// GameClient is a client to call Game service RPCs
type GameClient struct {
	service gen.GameServiceClient
}

// NewGameClient returns a new Game client
func NewGameClient(cc *grpc.ClientConn) *GameClient {
	service := gen.NewGameServiceClient(cc)
	return &GameClient{service}
}

// SendQuestionAnswer calls AnswerQuestion RPC
func (gameClient *GameClient) SendQuestionAnswer(
	answer1 string, // cuong_admin
	answer2 string, // lehuynhcuong
	question_num int,
) error {
	answers := []string{answer1, answer2}
	username := []string{"cuong_admin", "lehuynhcuong"}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := gameClient.service.AnswerQuestion(ctx)
	if err != nil {
		return fmt.Errorf("cannot answer question: %v", err)
	}

	waitResponse := make(chan error)
	// go routine to receive responses
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				// log.Print("no more responses")
				waitResponse <- nil
				return
			}
			if err != nil {
				waitResponse <- fmt.Errorf("cannot receive stream response: %v", err)
				return
			}

			log.Print("received response: ", res)
		}
	}()

	// send requests
	for i := 0; i < 2; i++ {
		req := &gen.AnswerQuestionRequest{
			EventId:     5, // code cứng, lấy trong table event
			Username:    username[i],
			Answer:      answers[i],
			QuestionNum: int32(question_num),
		}

		err := stream.Send(req)
		if err != nil {
			return fmt.Errorf("cannot send stream request: %v - %v", err, stream.RecvMsg(nil))
		}

		log.Print("sent request: ", req)
	}

	err = stream.CloseSend()
	if err != nil {
		return fmt.Errorf("cannot close send: %v", err)
	}

	err = <-waitResponse
	return err
}

func (gameClient *GameClient) GetQuestion(question_num int) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := gameClient.service.GetQuestion(ctx)
	if err != nil {
		return fmt.Errorf("cannot get question: %v", err)
	}

	waitResponse := make(chan error)
	// go routine to receive responses
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				// log.Print("no more responses")
				waitResponse <- nil
				return
			}
			if err != nil {
				waitResponse <- fmt.Errorf("cannot receive stream response: %v", err)
				return
			}

			log.Print("received response: ", res)
		}
	}()

	// send requests
	for i := 0; i < 2; i++ {
		req := &gen.GetQuestionRequest{
			EventId:     5, // code cứng, lấy trong table event
			QuestionNum: int32(question_num),
		}

		err := stream.Send(req)
		if err != nil {
			return fmt.Errorf("cannot send stream request: %v - %v", err, stream.RecvMsg(nil))
		}

		log.Print("sent request: ", req)
	}

	err = stream.CloseSend()
	if err != nil {
		return fmt.Errorf("cannot close send: %v", err)
	}

	err = <-waitResponse
	return err
}
