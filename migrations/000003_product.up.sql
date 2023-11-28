CREATE TABLE "product" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "name" varchar(255) NOT NULL,
  "shop_id" int NOT NULL,
  "price" integer NOT NULL,
  "count" integer NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (current_timestamp),
  "updated_at" timestamp NOT NULL DEFAULT (current_timestamp),
  "deleted_at" timestamp
);

ALTER TABLE "product" ADD FOREIGN KEY ("shop_id") REFERENCES "shop" ("id");
