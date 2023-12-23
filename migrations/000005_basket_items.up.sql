CREATE TABLE "basket_items"(
                               "id" SERIAL PRIMARY KEY NOT NULL,
                               "basket_id" int NOT NULL,
                               "product_id" int NOT NULL,
                               "quantity" int
);

ALTER TABLE "basket_items" ADD FOREIGN KEY ("basket_id") REFERENCES "basket" ("id");

ALTER TABLE "basket_items" ADD FOREIGN KEY ("product_id") REFERENCES "product" ("id");
