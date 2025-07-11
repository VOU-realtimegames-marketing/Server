package task

import "google.golang.org/genai"

type PayloadGenQuiz struct {
	QuizGenre string `json:"quiz_genre"`
	EventId   int64  `json:"event_id"`
	QuizNum   int32  `json:"quiz_num"`
}

type PayloadQuizCreated struct {
	EventId int64 `json:"event_id"`
}

// sudo rabbitmqctl list_queues name messages_ready messages_unacknowledged
type Quiz struct {
	Question string   `json:"question"`
	Answer   string   `json:"answer"`
	Options  []string `json:"options"`
}

var ConfigGenQuiz = &genai.GenerateContentConfig{
	ResponseMIMEType: "application/json",
	ResponseSchema: &genai.Schema{
		Type: genai.TypeArray,
		Items: &genai.Schema{
			Type: genai.TypeObject,
			Properties: map[string]*genai.Schema{
				"question": {Type: genai.TypeString},
				"answer":   {Type: genai.TypeString},
				"options": {
					Type:  genai.TypeArray,
					Items: &genai.Schema{Type: genai.TypeString},
				},
			},
			Required: []string{"question", "answer", "options"},
		},
	},
}

const QuizGeneratePrompt = `
The following is a Quizzes Generator Agent that is designed to generate quizzes based on: number of quizzes generated and genre.
Example: Generate 2 quizzes about football. (Number of quizzes: 2, Genre: football)
The output must be a JSON string in the exact following format:
[
	{
		question: "Which country won the 2022 FIFA World Cup?",
		answer: "Argentina",
		options: ["Brazil", "France", "Argentina", "Germany"]
	},
	{
		question: "What is the name of the most prestigious club competition in European football?",
		answer: "UEFA Champions League",
		options: ["UEFA Europa League", "UEFA Champions League", "FIFA Club World Cup", "FA Cup"]
	}
]

Generate %d quizzes about %s.
Output:
`
