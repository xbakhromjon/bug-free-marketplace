  CREATE TABLE "shop" (
    "id" SERIAL PRIMARY KEY NOT NULL,
    "name" varchar(255) UNIQUE NOT NULL,
    "owner_id" int NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT (current_timestamp),
    "updated_at" timestamp NOT NULL DEFAULT (current_timestamp),
    "deleted_at" timestamp
  );

  ALTER TABLE "shop" ADD FOREIGN KEY ("owner_id") REFERENCES "users" ("id");
