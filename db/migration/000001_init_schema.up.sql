CREATE TYPE "events_status" AS ENUM (
  'generating',
  'not_accepted',
  'ready'
);

CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "role" varchar NOT NULL DEFAULT 'user',
  "photo" varchar NOT NULL DEFAULT 'default-user.jpg',
  "active" boolean NOT NULL DEFAULT true,
  "is_email_verified" boolean NOT NULL DEFAULT false,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "verify_emails" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "email" varchar NOT NULL,
  "secret_code" varchar NOT NULL,
  "is_used" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "expired_at" timestamptz NOT NULL DEFAULT (now() + interval '15 minutes')
);

CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "username" varchar NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfer_voucher" (
  "id" bigserial PRIMARY KEY,
  "from" varchar NOT NULL,
  "to" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "voucher_id" bigint NOT NULL
);

CREATE TABLE "voucher_owner" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "voucher_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "stores" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "name" varchar NOT NULL,
  "business_type" varchar NOT NULL
);

CREATE TABLE "branchs" (
  "id" bigserial PRIMARY KEY,
  "store_id" bigint NOT NULL,
  "name" varchar NOT NULL,
  "position" varchar UNIQUE NOT NULL,
  "city_name" varchar NOT NULL,
  "country" varchar NOT NULL,
  "address" varchar NOT NULL,
  "emoji" varchar NOT NULL
);

CREATE TABLE "games" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "photo" varchar NOT NULL DEFAULT 'default-game.jpg',
  "type" varchar NOT NULL DEFAULT 'phone-shake',
  "play_guide" varchar,
  "gift_allowed" boolean NOT NULL DEFAULT true
);

CREATE TABLE "events" (
  "id" bigserial PRIMARY KEY,
  "game_id" bigint NOT NULL,
  "store_id" bigint NOT NULL,
  "owner" varchar NOT NULL,
  "name" varchar NOT NULL,
  "photo" varchar NOT NULL DEFAULT 'default-game.jpg',
  "voucher_quantity" int NOT NULL DEFAULT 5,
  "status" events_status NOT NULL DEFAULT 'generating',
  "start_time" timestamptz NOT NULL DEFAULT (now() + interval '60 minutes'),
  "end_time" timestamptz NOT NULL DEFAULT (now() + interval '75 minutes')
);

CREATE TABLE "quizzes" (
  "id" bigserial PRIMARY KEY,
  "event_id" bigint NOT NULL,
  "content" jsonb NOT NULL,
  "quiz_genre" varchar NOT NULL DEFAULT 'miscellaneous',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "vouchers" (
  "id" bigserial PRIMARY KEY,
  "event_id" bigint NOT NULL,
  "qr_code" varchar,
  "type" varchar NOT NULL DEFAULT 'phone',
  "status" varchar NOT NULL DEFAULT 'valid',
  "expires_at" timestamptz NOT NULL
);

CREATE INDEX ON "transfer_voucher" ("from");

CREATE INDEX ON "transfer_voucher" ("to");

CREATE INDEX ON "transfer_voucher" ("from", "to");

CREATE INDEX ON "voucher_owner" ("username");

CREATE INDEX ON "stores" ("owner");

CREATE UNIQUE INDEX ON "branchs" ("store_id", "name");

CREATE INDEX ON "events" ("store_id");

CREATE INDEX ON "events" ("game_id");

ALTER TABLE "verify_emails" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "sessions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "transfer_voucher" ADD FOREIGN KEY ("from") REFERENCES "users" ("username");

ALTER TABLE "transfer_voucher" ADD FOREIGN KEY ("to") REFERENCES "users" ("username");

ALTER TABLE "transfer_voucher" ADD FOREIGN KEY ("voucher_id") REFERENCES "vouchers" ("id");

ALTER TABLE "voucher_owner" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "voucher_owner" ADD FOREIGN KEY ("voucher_id") REFERENCES "vouchers" ("id");

ALTER TABLE "stores" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

ALTER TABLE "branchs" ADD FOREIGN KEY ("store_id") REFERENCES "stores" ("id");

ALTER TABLE "events" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

ALTER TABLE "events" ADD FOREIGN KEY ("game_id") REFERENCES "games" ("id");

ALTER TABLE "events" ADD FOREIGN KEY ("store_id") REFERENCES "stores" ("id");

ALTER TABLE "quizzes" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id");

ALTER TABLE "vouchers" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id");

ALTER TABLE "vouchers"
ADD CONSTRAINT "status_check"
CHECK ("status" IN ('valid', 'expired'));

ALTER TABLE "games"
ADD CONSTRAINT "game_type"
CHECK ("type" IN ('phone-shake', 'quiz'));