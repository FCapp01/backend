CREATE TABLE "restaurant" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY NOT NULL,
  "name" varchar NOT NULL,
  "address" varchar NOT NULL,
  "phone_number" varchar NOT NULL,
  "email" varchar NOT NULL
);

CREATE TABLE "menu" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY NOT NULL,
  "restaurant_id" int NOT NULL,
  "name" varchar NOT NULL,
  "description" varchar,
  "updated_at" timestamp DEFAULT (now()) NOT NULL
);

CREATE TABLE "item" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY NOT NULL,
  "menu_id" int NOT NULL,
  "name" varchar NOT NULL,
  "description" varchar,
  "image" varchar,
  "price" decimal NOT NULL,
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "order" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "restaurant_id" int NOT NULL,
  "customer_name" varchar NOT NULL,
  "customer_address" varchar NOT NULL,
  "customer_phone_number" varchar NOT NULL,
  "order_time" timestamp DEFAULT (now()) NOT NULL
);

CREATE TABLE "payment" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY NOT NULL,
  "order_id" int NOT NULL,
  "amount" decimal NOT NULL
);

ALTER TABLE "menu" ADD FOREIGN KEY ("restaurant_id") REFERENCES "restaurant" ("id");

ALTER TABLE "item" ADD FOREIGN KEY ("menu_id") REFERENCES "menu" ("id");

ALTER TABLE "order" ADD FOREIGN KEY ("restaurant_id") REFERENCES "restaurant" ("id");

ALTER TABLE "payment" ADD FOREIGN KEY ("order_id") REFERENCES "order" ("id");
