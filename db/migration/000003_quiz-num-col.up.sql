ALTER TABLE "quizzes"
ADD "quiz_num" integer NOT NULL DEFAULT 5;

CREATE TABLE "user_answers" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "event_id" bigint NOT NULL,
  "num_correct" integer NOT NULL DEFAULT 1
);

ALTER TABLE "user_answers" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "user_answers" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id");