CREATE TABLE "basket"(
                         "id" SERIAL PRIMARY KEY NOT NULL,
                         "user_id" int NOT NULL,
                         "purchased" boolean
);

ALTER TABLE "basket" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
