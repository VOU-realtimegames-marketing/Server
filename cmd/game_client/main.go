package main

import (
	"VOU-Server/cmd/game_client/client"
	"fmt"
	"log"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func testAnswerQuestion(gameClient *client.GameClient) {

	n := 5
	i := 0

	answer1 := []string{"Paris", "Mount Everest", "H2O", "Leonardo da Vinci", "Jupiter"}
	answer2 := []string{"Paris", "Mount Everest", "O2", "Leonardo da Vinci", "Mars"}
	for {
		if i >= n {
			break
		}

		fmt.Print("answer question (y/n)? ")
		var answer string
		fmt.Scan(&answer)

		if strings.ToLower(answer) != "y" {
			break
		}

		err := gameClient.SendQuestionAnswer(answer1[i], answer2[i], i+1)
		if err != nil {
			log.Fatal(err)
		}
		i++
	}
}

func main() {

	cc, err := grpc.NewClient("0.0.0.0:50054", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	gameClient := client.NewGameClient(cc)
	testAnswerQuestion(gameClient)
}
