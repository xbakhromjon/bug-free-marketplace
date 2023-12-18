CREATE TABLE "cart_items"(
    "id" SERIAL PRIMARY KEY NOT NULL,
    "cart_id" int NOT NULL,
    "product_id" int NOT NULL,
    "quantity" int
);

ALTER TABLE "cart_itemS" ADD FOREIGN KEY ("cart_id") REFERENCES "cart" ("id");

ALTER TABLE "cart_items" ADD FOREIGN KEY ("product_id") REFERENCES "product" ("id");