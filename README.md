# Tech Challenge FIAP

Aplicação responsável pela gestão de pedidos da hamburgueria Zé do Burguer via totem de auto atendimento.

## Documentação

[DDD](https://miro.com/app/board/uXjVMjkFsPU=/?share_link_id=958233804889)

[Arquitetura](#arquitetura)

[Stack](#stack-utilizada)

[Instalação](#instalação)

[APIs](#documentação-da-api)

---

## Arquitetura

### Hexagonal

![Hexagonal](./doc/arquitetura/hexagonal.svg)

### Estrutura do projeto

- doc 
- infra
- src
    - **Adapter**: Módulo responsável por realizar a recepção e armazenamento de dados, e a integração com sistemas ou serviços de terceiros
        - **Inbound**
            - **Controller**: Camada responsável por processar a validação dos dados e direcionar a requisição para o serviço;
            - **Handler**: Camada responsável por definir o meio de recepção das requisições; ex: REST API, GraphQL, Mensageria
        - **Outbound**: Camada onde realizamos a implementação das ports **repository** e **external**
            - **Repository**: Camada responsável por realizar a integração com o banco de dados e serviços externos; Ex: MySQL, PostgreSQL, DynamoDB, Integração com Mercado Pago, Integração com Mensageria
    - **Core**: Módulo responsável pelo coração do negócio
        - **Domain**: Camada responsável pelas entidades do negócio; 
        - **Port**: Camada responsável por definir as interfaces de **Service**, **Repository** e **External**;
        - **Service**: Camada responsável pela implementação da regra de negócio;

---

## Stack utilizada

**Linguagem:** Go lang (v1.21)

**Banco de dados:** PostgreSQL

**Ambiente:** Docker v24.0.5 | Docker-compose v2.20.2


---

## Instalação

Clone o projeto

```bash
  git clone https://github.com/ViniAlvesMartins/tech-challenge-fiap.git
```

Entre no diretório do projeto

```bash
  cd tech-challenge-fiap
```

Crie o arquivo `.env` apartir do `.env.example`

```bash
  cp .env.example .env
```

Inicie a aplicação

```bash
  docker-compose up
```


## Documentação da API

[Collection_Insomnia](./doc/apis/insomnia.json)
#### Base URL: http://localhost:8080

### Cliente

#### Cadastro de Cliente

```http
  POST /client
```

| Parâmetro   |   Tipo    | Local     | Descrição                           |
| :---------- | :-------  | :---------| :---------------------------------- |
| cpf         |   string  | body      | **Obrigatório**. O CPF do cliente   |
| name        |   string  | body      | **Obrigatório**. O nome do cliente  |
| email       |   string  | body      | **Obrigatório**. O email do cliente |

#### Response

```http
StatusCode: 200
Response: {
	"MessageError": "",
	"Data": {
		"id": 48,
		"cpf": 19034721801,
		"name": "cliente1",
		"email": "cli1ent1a22@teste.com"
	}
}
```
```http
StatusCode: 409
Response: {
	"MessageError": "Client already exists",
	"Data": null
}

```

#### Consulta de Cliente por CPF

```http
  GET /client
```

| Parâmetro   |   Tipo    | Local     | Descrição                           |
| :---------- | :-------  | :---------| :---------------------------------- |
| cpf         |   string  | Query     | **Obrigatório**. O CPF do cliente   |

#### Response

```http
StatusCode: 200
Response: {
	"MessageError": "",
	"Data": {
		"id": 48,
		"cpf": 19034721801,
		"name": "cliente1",
		"email": "cli1ent1a22@teste.com"
	}
}
```
```http
StatusCode: 404
Response: {
	"MessageError": "Not found",
	"Data": null
}
```

### Produto

##### Enum de categoria
| Id | Nome           |
| :- | :------------- |
|  1 | Lanche         |
|  2 | Bebida         |
|  3 | Acompanhamento |
|  4 | Sobremesa      |

#### Cadastro de Produto

```http
  POST /product
```

| Parâmetro   |   Tipo    | Local     | Descrição                           |
| :---------- | :-------  | :---------| :---------------------------------- |
| name_product|   string  | Body      | **Obrigatório**. O nome do produto  |
| description |   string  | Body      | **Obrigatório**. A descrição do produto |
| price       |   decimal | Body      | **Obrigatório**. O preço do produto |
| category_id |   number  | Body      | **Obrigatório**. A categoria do produto   |

#### Response

```http
StatusCode: 201
Response: {
	"MessageError": "",
	"Data": {
		"id": 23,
		"name_product": "teste",
		"description": "descricao teste",
		"price": 10,
		"category_id": 2,
		"active": true
	}
}
```

```http
StatusCode: 400
Response: {
	"Errors": [
		{
			"Field": "Price",
			"Message": "required"
		}
	]
}
```

```http
StatusCode: 404
Response: {
	"MessageError": "Category not found",
	"Data": null
}
```

#### Consulta de Produto por categoria

```http
  GET /category/{category_id}/product
```

| Parâmetro   |   Tipo    | Local     | Descrição                           |
| :---------- | :-------  | :---------| :---------------------------------- |
| category_id |   number  | pathParam | **Obrigatório**. O códgio da categoria  |

#### Response

```http
StatusCode: 200
Response: {
	"MessageError": "",
	"Data": [
		{
			"id": 3,
			"name_product": "batata frita",
			"description": "quente",
			"price": 5,
			"category_id": 3,
			"active": true
		},
		{
			"id": 6,
			"name_product": "nuggets",
			"description": "quente",
			"price": 7,
			"category_id": 3,
			"active": true
		}
	]
}
```

```http
StatusCode: 404
Response: {
	"MessageError": "Category Not found",
	"Data": null
}
```

```http
StatusCode: 404
Response: {
	"MessageError": "Product Not found",
	"Data": null
}
```

#### Edição de Produto

```http
  PUT /product/{product_id}
```

| Parâmetro   |   Tipo    | Local     | Descrição                           |
| :---------- | :-------  | :---------| :---------------------------------- |
| product_id  |   number  | pathParam | **Obrigatório**. O códgio do produto|

#### Response

```http
StatusCode: 200
Response: {
	"MessageError": "",
	"Data": {
		"id": 6,
		"name_product": "teste",
		"description": "descricao teste",
		"price": 10,
		"category_id": 1,
		"active": true
	}
}
```

```http
StatusCode: 404
Response: {
	"MessageError": "Not found",
	"Data": null
}
```

```http
StatusCode: 400
Response: {
	"Errors": [
		{
			"Field": "NameProduct",
			"Message": "required"
		},
		{
			"Field": "Description",
			"Message": "required"
		},
		{
			"Field": "Price",
			"Message": "required"
		}
	]
}
```

#### Delete de Produto

```http
  DELETE /product/{product_id}
```

| Parâmetro   |   Tipo    | Local     | Descrição                           |
| :---------- | :-------  | :---------| :---------------------------------- |
| product_id  |   number  | pathParam | **Obrigatório**. O códgio do produto|

#### Response

```http
StatusCode: 200
Response: {}
```

```http
StatusCode: 404
Response: {
	"MessageError": "Not found",
	"Data": null
}
```

### Pedido

#### Criação de pedido

```http
  POST /order
```

| Parâmetro   |   Tipo       | Local     | Descrição                           |
| :---------- | :-------     | :---------| :---------------------------------- |
| client_id   |   number     | body      | O código do cliente                 |
| products    |list(products)| body      | **Obrigatório**. Lista de produtos  |

#### Exemplo da lista de produtos

```http
"products": [
        {
            "id": 2
        },
        {
            "id": 5
        }
    ]
```

#### Response

```http
StatusCode: 200
Response: {
	"MessageError": "",
	"Data": {
		"id": 37,
		"client_id": null,
		"status_order": "AWAITING_PAYMENT",
		"amount": 6,
		"created_at": "2023-10-31T04:09:21.493587439Z",
		"products": [
			{
				"id": 2,
				"name_product": "",
				"description": "",
				"price": 0,
				"category_id": 0,
				"active": false
			},
			{
				"id": 5,
				"name_product": "",
				"description": "",
				"price": 0,
				"category_id": 0,
				"active": false
			}
		]
	}
}
```
```http
StatusCode: 404
Response: {
	"MessageError": "Product not found 563",
	"Data": null
}

```

```http
StatusCode: 400
Response: {
	"MessageError": "Product is required",
	"Data": null
}
```
