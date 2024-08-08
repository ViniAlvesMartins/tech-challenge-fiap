CREATE SCHEMA IF NOT EXISTS ze_burguer;

CREATE TYPE ze_burguer.order_status AS ENUM ('AWAITING_PAYMENT', 'PAID', 'RECEIVED', 'PREPARING', 'READY', 'FINISHED', 'CANCELED');

-- Create clients table
CREATE TABLE IF NOT EXISTS ze_burguer.clients (
    "id" BIGSERIAL NOT NULL,
    "cpf" BIGINT NOT NULL UNIQUE,
    "name" VARCHAR(55) NOT NULL,
    "email" VARCHAR(55) NOT NULL UNIQUE,
    CONSTRAINT "PK_Clients" PRIMARY KEY ("id")
);

-- Sample data clients
INSERT INTO ze_burguer.clients( "cpf", "name", "email")
VALUES (58482697005, 'Ruperto Edmenson', 'redmenson0@go.com');

INSERT INTO ze_burguer.clients ( "cpf", "name", "email")
VALUES (76891019095, 'Andrej Hazeman', 'ahazeman1@rediff.com');

INSERT INTO ze_burguer.clients ( "cpf", "name", "email")
VALUES (56077689025, 'Bax Kleimt', 'bkleimt2@cocolog-nifty.com');

INSERT INTO ze_burguer.clients ( "cpf", "name", "email")
VALUES (58039865000, 'Brett Moralis', 'bmoralis3@posterous.com');

INSERT INTO ze_burguer.clients ( "cpf", "name", "email")
VALUES (79155626068, 'Dorey Sapsed', 'dsapsed4@dailymail.co.uk');

INSERT INTO ze_burguer.clients ( "cpf", "name", "email")
VALUES (75075366023, 'Doralin Spincke', 'dspincke5@illinois.edu');

-- Create orders table
CREATE TABLE IF NOT EXISTS ze_burguer.orders (
    "id" BIGSERIAL NOT NULL,
    "client_id" INT NULL,
    "order_status" ze_burguer.order_status NOT NULL,
    "amount" FLOAT NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    CONSTRAINT "PK_order" PRIMARY KEY ("id"),
    CONSTRAINT "FK_client" FOREIGN KEY ("client_id") REFERENCES ze_burguer.clients(id)
);

-- Sample data orders
INSERT INTO ze_burguer.orders( "client_id", "order_status", "created_at", "amount")
VALUES (1, 'AWAITING_PAYMENT', '2023-10-13 11:30:30', 17.51);
INSERT INTO ze_burguer.orders( "client_id", "order_status", "created_at", "amount")
VALUES (2, 'PREPARING', '2023-10-13 11:31:30', 20.50);
INSERT INTO ze_burguer.orders( "client_id", "order_status", "created_at", "amount")
VALUES (3, 'RECEIVED', '2023-10-13 11:32:30', 15);
INSERT INTO ze_burguer.orders( "client_id", "order_status", "created_at", "amount")
VALUES (4, 'FINISHED', '2023-10-13 11:33:30', 17);

-- Create categories table
CREATE TABLE IF NOT EXISTS ze_burguer.categories (
    "id" INT NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    CONSTRAINT "PK_category" PRIMARY KEY ("id")
);

-- Sample data category
INSERT INTO ze_burguer.categories( "id", "name")
VALUES (1, 'Lanche');
INSERT INTO ze_burguer.categories( "id", "name")
VALUES (2, 'Bebida');
INSERT INTO ze_burguer.categories( "id", "name")
VALUES (3, 'Acompanhamento');
INSERT INTO ze_burguer.categories( "id", "name")
VALUES (4, 'Sobremesa');

-- Create category product table
CREATE TABLE IF NOT EXISTS ze_burguer.products (
    "id" BIGSERIAL NOT NULL,
    "category_id" INT NOT NULL,
    "product_name" VARCHAR(255) NOT NULL,
    "price" FLOAT NOT NULL,
    "description" VARCHAR(300) NOT NULL,
    "active" BOOLEAN NOT NULL,
    CONSTRAINT "PK_products" PRIMARY KEY ("id"),
    CONSTRAINT "FK_category" FOREIGN KEY ("category_id") REFERENCES ze_burguer.categories(id)
);

-- Sample data products
INSERT INTO ze_burguer.products( "category_id", "product_name", "price", "description", "active")
VALUES (1, 'X-salada', 10, 'Delicioso hamburger com salada, queijo, tomate e alface', true);
INSERT INTO ze_burguer.products( "category_id", "product_name", "price", "description", "active")
VALUES (2, 'Suco de laranja', 2.50, 'Suco natural, extraído de laranjas selecionadas', true);
INSERT INTO ze_burguer.products( "category_id", "product_name", "price", "description", "active")
VALUES (3, 'Batata frita', 5, 'Batatas fritas sequinhas e crocantes', true);
INSERT INTO ze_burguer.products( "category_id", "product_name", "price", "description", "active")
VALUES (1, 'X-bacon', 10, 'Delicioso hamburger com salada, queijo, tomate, alface e bacon', true);
INSERT INTO ze_burguer.products( "category_id", "product_name", "price", "description", "active")
VALUES (2, 'Suco de limão', 3.50, 'Suco natural, extraído de limões selecionados', true);
INSERT INTO ze_burguer.products( "category_id", "product_name", "price", "description", "active")
VALUES (3, 'Nuggets', 7, 'Uma caixinha com 4 unidades dos mais deliciosos Nuggets ', true);

-- Create orders products.
CREATE TABLE IF NOT EXISTS ze_burguer.orders_products (
    "id" BIGSERIAL NOT NULL,
    "order_id" INT NOT NULL,
    "product_id" INT NOT NULL,
    CONSTRAINT "FK_id_order" FOREIGN KEY ("order_id") REFERENCES  ze_burguer.orders(id),
    CONSTRAINT "FK_id_product" FOREIGN KEY ("product_id") REFERENCES ze_burguer.products(id)
); 