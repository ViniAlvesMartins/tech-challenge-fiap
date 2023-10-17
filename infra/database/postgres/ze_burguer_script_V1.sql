CREATE TABLE ze_burguer.clients (
	id serial PRIMARY KEY,
	cpf INTEGER NOT NULL UNIQUE,
	name VARCHAR (255),
	email VARCHAR (255) UNIQUE
);