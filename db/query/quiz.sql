-- name: CreateQuiz :one
INSERT INTO quizzes (
  event_id,
  content,
  quiz_genre
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetQuizzesByEventId :one
SELECT * FROM quizzes
WHERE event_id = $1 LIMIT 1;