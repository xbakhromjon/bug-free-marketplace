CREATE TABLE "cart_items" (
                              "id" SERIAL PRIMARY KEY NOT NULL,
                              "product_id" int NOT NULL,
                              "cart_id" int NOT NULL,
                              "unit_price" integer NOT NULL,
                              "quantity" integer NOT NULL,
);

ALTER TABLE "cart_items" ADD FOREIGN KEY ("product_id") REFERENCES "product" ("id");
ALTER TABLE "cart_items" ADD FOREIGN KEY ("cart_id") REFERENCES "cart" ("id");
