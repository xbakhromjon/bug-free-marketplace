CREATE TABLE IF NOT EXISTS "orders" (
    "id" SERIAL PRIMARY KEY NOT NULL,
    "number" VARCHAR(255) NOT NULL,
    "basket_id" INTEGER NOT NULL,
    "total_price" INTEGER NOT NULL,
    "status" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE "orders" ADD FOREIGN KEY ("basket_id") REFERENCES "basket" ("id");