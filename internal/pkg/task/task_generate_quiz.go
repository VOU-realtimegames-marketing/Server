package task

type PayloadGenQuiz struct {
	QuizGenre string `json:"quiz_genre"`
	EventId   int64  `json:"event_id"`
}

type PayloadQuizCreated struct {
	Status bool `json:"status"`
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
