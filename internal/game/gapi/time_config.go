package gapi

import "time"

// 0------ 6 -----12 ----- 18 ------ 24 -------30
const (
	// Time to answer a question in seconds
	QUESTION_DISPLAY_TIME int32 = 5

	// Time for the server to handle the answer in seconds
	SERVER_HANDLE_TIME int32 = 2

	// Time to show the correct answer in seconds
	ANSWER_DISPLAY_TIME int32 = 1

	// Time to show the leaderboard in seconds
	LEADERBOARD_DISPLAY_TIME int32 = 3

	// Total time to answer a question in seconds
	TOTAL_QUESTION_TIME = QUESTION_DISPLAY_TIME + SERVER_HANDLE_TIME + ANSWER_DISPLAY_TIME

	TIME_TOLERANCE = 200 * time.Millisecond
)
