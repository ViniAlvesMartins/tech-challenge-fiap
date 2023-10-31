# Tech Challenge FIAP

Aplicação responsável pela gestão de pedidos da hamburgueria Zé do Burguer via totem de auto atendimento.

### Building image
`$ make build`

### Running dev environment
`$ make run-dev`

### Running prod environment
`$ make run-prod`

### Generate mocks
`$ make mocks`

### Running tests
`$ make test`


## Documentação

[DDD](https://miro.com/app/board/uXjVMjkFsPU=/?share_link_id=958233804889)

[Arquitetura](#arquitetura)

[Stack](#stack-utilizada)

[Instalação](#instalação)

[APIs]()

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
            - **Repository**: Camada responsável por realizar a integração com o banco de dados; Ex: MySQL, PostgreSQL, DynamoDB
            - **External**: Camada responsável por realizar a integração com sistema ou serviços de terceiros; Ex: Integração com Mercado Pago, Integração com Mensageria
    - **Core**: Módulo responsável pelo coração do negócio
        - **Domain**: Camada responsável pelas entidades do negócio; 
        - **Port**: Camada responsável por definir as interfaces de **Service**, **Repository** e **External**;
        - **Service**: Camada responsável pela implementação da regra de negócio;

---

## Stack utilizada

**Linguagem:** Go lang (v1.21)

**Banco de dados:** PostgreSQL

**Ambiente:** Docker

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

Inicie a aplicação com o ambiente de DEV

```bash
  docker-compose up
```

ou inicie a aplicação com o ambiente de PROD (sem seeder, apenas com a seed de categoria)

```bash
  docker-compose -f docker-compose.prod.yaml up
```

---
