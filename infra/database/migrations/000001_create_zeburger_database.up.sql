CREATE SCHEMA IF NOT EXISTS ze_burguer;

-- Create clients table
CREATE TABLE IF NOT EXISTS ze_burguer.clients (
    "id" BIGSERIAL NOT NULL,
    "cpf" INT NOT NULL UNIQUE,
    "name" VARCHAR(55) NOT NULL,
    "email" VARCHAR(55) NOT NULL UNIQUE,
    CONSTRAINT "PK_Clients" PRIMARY KEY ("id")
); 
-- Sample data clients
INSERT INTO ze_burguer.clients( "cpf", "name", "email")
VALUES (142358967, 'dbmussarelo', 'dbmussarelo@emailo.com');

INSERT INTO ze_burguer.clients ( "cpf", "name", "email")
VALUES (142358987, 'dbcalabresso', 'dbcalabresso@emailo.com');

INSERT INTO ze_burguer.clients ( "cpf", "name", "email")
VALUES (185358987, 'dbtroncudo', 'dbtroncudo@emailo.com');

INSERT INTO ze_burguer.clients ( "cpf", "name", "email")
VALUES (185898987, 'dbcasosbahio', 'dbcasosbahio@emailo.com');

INSERT INTO ze_burguer.clients ( "cpf", "name", "email")
VALUES (185898222, 'dbludmilo', 'dbludmilo@emailo.com');

INSERT INTO ze_burguer.clients ( "cpf", "name", "email")
VALUES (185898224, 'dbdelicio', 'dbdelicio@emailo.com');

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
-- Sample data orders
INSERT INTO ze_burguer.orders( "client_id", "status_order", "created_at")
VALUES (1, 'WAITING', '2023-10-13 11:30:30');
INSERT INTO ze_burguer.orders( "client_id", "status_order", "created_at")
VALUES (2, 'PREPARING', '2023-10-13 11:31:30');
INSERT INTO ze_burguer.orders( "client_id", "status_order", "created_at")
VALUES (3, 'RECEIVED', '2023-10-13 11:32:30');
INSERT INTO ze_burguer.orders( "client_id", "status_order", "created_at")
VALUES (4, 'FINISHED', '2023-10-13 11:33:30');

-- Create category table
CREATE TABLE IF NOT EXISTS ze_burguer.category (
    "id" INT NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    CONSTRAINT "PK_category" PRIMARY KEY ("id")
); 
-- Sample data category
INSERT INTO ze_burguer.category( "id", "name")
VALUES (1, 'lanche');
INSERT INTO ze_burguer.category( "id", "name")
VALUES (2, 'bebida');
INSERT INTO ze_burguer.category( "id", "name")
VALUES (3, 'acompanhamento');

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
-- Sample data products
INSERT INTO ze_burguer.products( "category_id", "name_product", "price", "description")
VALUES (1, 'x-salada', 10, 'hamburgão com salada');
INSERT INTO ze_burguer.products( "category_id", "name_product", "price", "description")
VALUES (2, 'suco de laranja', 2.50, 'gelado');
INSERT INTO ze_burguer.products( "category_id", "name_product", "price", "description")
VALUES (3, 'batata frita', 5, 'quente');
INSERT INTO ze_burguer.products( "category_id", "name_product", "price", "description")
VALUES (1, 'x-bacon', 10, 'hamburgão com bacon');
INSERT INTO ze_burguer.products( "category_id", "name_product", "price", "description")
VALUES (2, 'suco de limão', 3.50, 'gelado');
INSERT INTO ze_burguer.products( "category_id", "name_product", "price", "description")
VALUES (3, 'nuggets', 7, 'quente');

-- Create orders products.
CREATE TABLE IF NOT EXISTS ze_burguer.orders_products (
    "id_orders" INT NOT NULL,
    "id_products" INT NOT NULL,
    CONSTRAINT "FK_id_orders" FOREIGN KEY ("id_orders") REFERENCES  ze_burguer.orders(id),
    CONSTRAINT "FK_id_products" FOREIGN KEY ("id_products") REFERENCES ze_burguer.products(id)
); 