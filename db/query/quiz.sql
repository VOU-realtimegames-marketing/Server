-- name: CreateQuiz :one
INSERT INTO quizzes (
  event_id,
  content,
  quiz_genre
) VALUES (
  $1, $2, $3
) RETURNING *;