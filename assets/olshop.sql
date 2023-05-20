CREATE TABLE "products" (
  "id" integer PRIMARY KEY,
  "product_name" varchar,
  "product_price" float,
  "product_qty" integer,
  "product_category" varchar,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "users" (
  "id" integer PRIMARY KEY,
  "name" varchar,
  "email" varchar,
  "password" varchar,
  "role" enum,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "carts" (
  "user_id" integer,
  "product_id" integer,
  "amount" integer,
  "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("user_id", "product_id")
);

CREATE TABLE "payments" (
  "id" integer PRIMARY KEY,
  "user_id" integer,
  "product_id" integer,
  "amount" integer,
  "total_price" float,
  "created_at" timestamp,
  "updated_at" timestamp
);

ALTER TABLE "carts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "carts" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

CREATE TABLE "users_payments" (
  "users_id" integer,
  "payments_user_id" integer,
  PRIMARY KEY ("users_id", "payments_user_id")
);

ALTER TABLE "users_payments" ADD FOREIGN KEY ("users_id") REFERENCES "users" ("id");

ALTER TABLE "users_payments" ADD FOREIGN KEY ("payments_user_id") REFERENCES "payments" ("user_id");


CREATE TABLE "payments_products" (
  "payments_product_id" integer,
  "products_id" integer,
  PRIMARY KEY ("payments_product_id", "products_id")
);

ALTER TABLE "payments_products" ADD FOREIGN KEY ("payments_product_id") REFERENCES "payments" ("product_id");

ALTER TABLE "payments_products" ADD FOREIGN KEY ("products_id") REFERENCES "products" ("id");

