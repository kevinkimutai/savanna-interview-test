CREATE TABLE customers (
"customer_id" varchar PRIMARY KEY,
"name" varchar NOT NULL,
"created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE orders (
"order_id" varchar PRIMARY KEY,
"customer_id" varchar,
"total_amount" numeric NOT NULL,
"created_at" timestamptz NOT NULL DEFAULT (now()),
FOREIGN KEY (customer_id) REFERENCES customers(customer_id)
);

CREATE TABLE products (
"product_id" bigserial PRIMARY KEY,
"name" varchar NOT NULL,
"price" numeric NOT NULL,
"image_url" varchar NOT NULL,
"created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE order_items (
    "order_item_id" bigserial PRIMARY KEY,
    "order_id" varchar NOT NULL,
    "product_id" bigint NOT NULL,
    "quantity" integer NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    FOREIGN KEY (order_id) REFERENCES orders(order_id),
    FOREIGN KEY (product_id) REFERENCES products(product_id)
);

