CREATE TYPE "customrole" AS ENUM (
  'merchant',
  'user',
  'admin'
);

CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "name" varchar(255) NOT NULL,
  "phone_number" varchar(12) NOT NULL,
  "password" varchar(128) NOT NULL,
  "role" customrole NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (current_timestamp),
  "updated_at" timestamp NOT NULL DEFAULT (current_timestamp),
  "deleted_at" timestamp
);

CREATE UNIQUE INDEX "unique_phone_number" ON "users" ("phone_number");
