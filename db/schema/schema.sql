CREATE TABLE "urls" (
  "id" varchar(5) PRIMARY KEY,
  "long_url" varchar NOT NULL,
  "short_url" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);