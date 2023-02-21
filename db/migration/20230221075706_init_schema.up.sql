CREATE TABLE "restaurant" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "name" varchar,
  "address" varchar,
  "phone_number" varchar,
  "email" varchar
);

CREATE TABLE "menu" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "restaurant_id" int,
  "name" varchar,
  "description" varchar,
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "item" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "menu_id" int,
  "name" varchar,
  "description" varchar,
  "image" varchar,
  "price" decimal,
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "order" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "restaurant_id" int,
  "customer_name" varchar,
  "customer_address" varchar,
  "customer_phone_number" varchar,
  "order_time" timestamp DEFAULT (now())
);

CREATE TABLE "payment" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "order_id" int,
  "amount" decimal
);

ALTER TABLE "menu" ADD FOREIGN KEY ("restaurant_id") REFERENCES "restaurant" ("id");

ALTER TABLE "item" ADD FOREIGN KEY ("menu_id") REFERENCES "menu" ("id");

ALTER TABLE "order" ADD FOREIGN KEY ("restaurant_id") REFERENCES "restaurant" ("id");

ALTER TABLE "payment" ADD FOREIGN KEY ("order_id") REFERENCES "order" ("id");
