# Tech Challenge FIAP

Aplicação responsável pela gestão de pedidos da hamburgueria Zé do Burguer via totem de auto atendimento.

## Oficial Image

[Docker Image](https://hub.docker.com/repository/docker/marcosilva/ze_burguer/general)

## Documentação

[DDD](https://miro.com/app/board/uXjVMjkFsPU=/?share_link_id=958233804889)

[Arquitetura](#arquitetura)

[Stack](#stack-utilizada)

[Instalação Docker](#instalação)

[Instalação Kubernetes](#instalação-k8s)

[APIs](#documentação-da-api)

---

## Arquitetura

### Clean Architecture

![Clean_Architecture](./doc/arquitetura/clean_arch.svg)

### Estrutura do projeto

- doc
- infra: Módulo responsável pela gestão da infra e configurações externas utilizadas na aplicação. Ex: Migrations, docker, config.go
- kustomize: Módulo responsável pela gestão dos arquivos kubernetes
- src
	- **External**: Módulo responsável por realizar a recepção dos dados e realizar a chamada para a controller
		- **Handler**: Camada responsável por definir o meio de recepção das requisições; ex: REST API, GraphQL, Mensageria
        - **Database**: Camada onde realizamos a configuração do banco de dados;
		- **Repository**: Camada responsável por realizar a integração com o banco de dados ; Ex: MySQL, PostgreSQL, DynamoDB, etc;
		- **Service**: Camada responsável por realizar a integração com serviços externos; Ex: Integração com mercado pago, integração com serviços AWS, etc;
    - **Controller**: Módulo responsável por realizar a validação dos dados e realizar a chamada para as UseCases necessárias
		- **Controller**: Camada responsável por realizar a validação dos dados e realizar a chamada para as UseCases necessárias;
		- **Serializer**: Camada responsável por definir o **Input** e **Output** da API;
    - **Application**: Módulo responsável pelo coração do negócio
        - **Modules**: Camada responsável pelos responses do **Service**; 
        - **Contract**: Camada responsável por definir as interfaces de **Service** e **Repository**;
        - **UseCases**: Camada responsável pela implementação da regra de negócio;
	- **Entities**: Módulo responsável pela definição das entidades da aplicação
        - **Entity**: Camada responsável pela definição das entidades que o sistema vai utilizar; Ex: Pedido, cliente
        - **Enum**: Camada responsável por definir os enums utilizados pelo negócio. Ex: Status do pedido, status do pagamento

---

## Stack utilizada

**Linguagem:** Go lang (v1.21)

**Banco de dados:** PostgreSQL

**Ambiente:** Docker v24.0.5 | Docker-compose v2.20.2

**Kubernetes:** V1.27.2

---

## Instalação Docker

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
## Instalação Kubernetes

Requisitos Cluster Kubernetes:

- Docker Kubernetes
- Metrics Server (Instalado)


Para criar os recursos 

```bash
  kubectl apply -k kustomize/
```

Com a execução acima será criado a seguinte infraestrutura:

Services
 - ze-burguer: NodePort 30443
 - postgres: ClusterIP 5432

Deployments
 - ze-burguer: HPA (2-5 Pods) - CPU Average Usage metrics
 - postgres: 1 Pod

![K8S](./doc/infra/kubernetes.png)

## Documentação da API

[Collection_Insomnia](./doc/apis/insomnia.json)
#### Base URL: http://localhost:8080
#### Doc Swagger: http://localhost:8080/docs/index.html