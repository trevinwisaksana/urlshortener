CREATE TABLE "urls" (
  "id" bigserial PRIMARY KEY,
  "long_url" varchar NOT NULL,
  "short_url" varchar NOT NULL UNIQUE,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "owner" varchar NOT NULL
);

CREATE INDEX ON "urls" ("owner");

CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "urls" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");
