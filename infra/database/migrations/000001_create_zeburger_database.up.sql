CREATE SCHEMA IF NOT EXISTS ze_burguer;

-- Create clients table
CREATE TABLE IF NOT EXISTS ze_burguer.clients (
    "id" BIGSERIAL NOT NULL,
    "cpf" INT NOT NULL UNIQUE,
    "name" VARCHAR(55) NOT NULL,
    "email" VARCHAR(55) NOT NULL UNIQUE,
    CONSTRAINT "PK_Clients" PRIMARY KEY ("id")
);

-- Create enum status
CREATE TYPE status_order AS ENUM ('WAITING', 'RECEIVED', 'PREPARING', 'READY', 'FINISHED');
-- Create orders table
CREATE TABLE IF NOT EXISTS ze_burguer.orders (
    "id" BIGSERIAL NOT NULL,
    "client_id" INT NOT NULL,
    "status_order" status_order NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    CONSTRAINT "PK_order" PRIMARY KEY ("id"),
    CONSTRAINT "FK_client" FOREIGN KEY ("client_id") REFERENCES ze_burguer.clients(id)
);

-- Create category table
CREATE TABLE IF NOT EXISTS ze_burguer.category (
    "id" INT NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    CONSTRAINT "PK_category" PRIMARY KEY ("id")
); 

-- Create category product table
CREATE TABLE IF NOT EXISTS ze_burguer.products (
    "id" BIGSERIAL NOT NULL,
    "category_id" INT NOT NULL,
    "name_product" VARCHAR(255) NOT NULL,
    "price" FLOAT NOT NULL,
    "description" VARCHAR(300) NOT NULL,
    "image_product" VARCHAR(10) NULL, 
    CONSTRAINT "PK_products" PRIMARY KEY ("id"),
    CONSTRAINT "FK_category" FOREIGN KEY ("category_id") REFERENCES ze_burguer.category(id)
);

-- Create orders products.
CREATE TABLE IF NOT EXISTS ze_burguer.orders_products (
    "id" BIGSERIAL NOT NULL,
    "order_id" INT NOT NULL,
    "product_id" INT NOT NULL,
    CONSTRAINT "FK_id_order" FOREIGN KEY ("order_id") REFERENCES  ze_burguer.orders(id),
    CONSTRAINT "FK_id_product" FOREIGN KEY ("product_id") REFERENCES ze_burguer.products(id)
); 