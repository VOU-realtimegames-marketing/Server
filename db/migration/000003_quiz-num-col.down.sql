ALTER TABLE "quizzes"
DROP COLUMN "quiz_num";

ALTER TABLE IF EXISTS "user_answers" DROP CONSTRAINT IF EXISTS "user_answers_username_fkey";

ALTER TABLE IF EXISTS "user_answers" DROP CONSTRAINT IF EXISTS "user_answers_event_id_fkey";

ALTER TABLE IF EXISTS "user_answers" DROP CONSTRAINT IF EXISTS "user_answers_quizzes_id_fkey";

DROP TABLE IF EXISTS "user_answers";
