package task

type PayloadGenQuiz struct {
	QuizGenre string `json:"quiz_genre"`
	EventId   int64  `json:"event_id"`
}

type PayloadQuizCreated struct {
	Status bool `json:"status"`
}
