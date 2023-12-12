CREATE TABLE "cart" (
                        "id" SERIAL PRIMARY KEY NOT NULL,
                        "user_id" int NOT NULL
);

ALTER TABLE "cart" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
