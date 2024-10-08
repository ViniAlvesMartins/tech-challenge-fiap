# Tech Challenge FIAP

Aplicação responsável pela gestão de pedidos da hamburgueria Zé do Burguer via totem de auto atendimento.

## Oficial Image

[Docker Image](https://hub.docker.com/repository/docker/marcosilva/ze_burguer/general)

## Documentação

[DDD](https://miro.com/app/board/uXjVMjkFsPU=/?share_link_id=958233804889)

[Arquitetura](#arquitetura)

[Arquitetura Cloud](#arquitetura-cloud)

[Stack](#stack-utilizada)

[SAGA](#saga-coreografia)

[Instalação Docker](#instalação-docker)

[Instalação Kubernetes](#instalação-kubernetes)

[APIs](#documentação-da-api)

[SONAR](#sonar)

[OWASP ZAP](#owasp-zap)

[Relatorio RIPD](#Relatorio-RIPD)

---

## Arquitetura

### Clean Architecture

![Clean_Architecture](./doc/arquitetura/clean_arch.svg)

### Estrutura do projeto


- cmd
- doc
- infra: Módulo responsável pela gestão da infra e configurações externas utilizadas na aplicação. Ex: Migrations, docker, config.go, kustomize
- internal
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

## Database

![postgresql](./doc/arquitetura/database.png)

Alguns pontos foram considerados para o uso de banco relacional na nossa solução, dentre eles podemos destacar: a natureza relacional entre as entidades e sua estrutura pouco flexível; o conhecimento difundido do SQL, dentre os integrante do grupo, como linguagem para interface com o banco de dados; conformidade com ACID(atomicidade, consistência, isolamento e durabilidade).

Dada a natureza da aplicação desenvolvida, levamos em consideração também a vantagem dos bancos de dados relacionais, em relação á flexibilidade para executar consultas mais complexas,
estratégia de indexação a mais tempo sendo colocadas á prova e a consistência dos dados sobre os bancos não relacionais.

Entre as opções de RDBMS, escolhemos o PostgreSQL por conta de extensa adoção no mercado, aumentando a gama de material de suporte disponível; seu controle de concorrência, que por padrão evita que transações não salvas afetem outras transações;
compatibilidade com as principais soluções de banco de dados como serviço do mercado(Amazon RDS, Microsoft Azure e Google GCP). O PostgreSQL também oferece diversas funções que facilitam o desenvolvimento, como:
uma vasta quantidade de tipos de dados; suporte nativo a UUID(identificador único universal);e controle de acesso granular, permitindo acesso á apenas o que for necessário ao usuário da aplicação.

---

## Arquitetura cloud

![Arquitetura_cloud_v2](./doc/arquitetura/cloud_arch_v2.png)

![Arquitetura_cloud_v2_detail](./doc/arquitetura/cloud_arch_v2_detail.png)

---

## Saga coreografia

O padrão SAGA coreografado é vantajoso para aplicações pequenas por sua simplicidade, desacoplamento de serviços, escalabilidade, resiliência e flexibilidade. Ele evita a complexidade de um orquestrador central, facilitando a manutenção e evolução independente dos serviços. A comunicação entre serviços por eventos promove baixo acoplamento e resiliência, permitindo que cada serviço escale e evolua separadamente. Em uma aplicação de hamburguer, por exemplo, os serviços de pedidos, inventário e pagamentos podem coordenar suas ações por meio de eventos sem um controlador central.

---

## Stack utilizada

**Linguagem:** Go (v1.22)

**Banco de dados:** PostgreSQL

**Ambiente:** Docker v24.0.5 | Docker-compose v2.20.2

**Kubernetes:** V1.27.2

---

## Instalação docker

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

Inicie a infra

```bash
  docker-compose -f docker-compose-infra.yaml up -d
  ou
  make start-infra
```

Inicie a aplicação

```bash
  docker-compose up dev-app-orders -d
  ou
  make run-app
```
## Instalação kubernetes

Requisitos Cluster Kubernetes:

- Docker Kubernetes
- Metrics Server (Instalado)
	- [Guia de instalação Cluster](https://github.com/kubernetes-sigs/metrics-server?tab=readme-ov-file#installation)

Requisitos Criação dos Recursos:

- Kubectl 
	- [Guia de instalação local](https://kubernetes.io/docs/tasks/tools/)

Para criar os recursos 

```bash
  kubectl apply -k kustomize/
```

Com a execução acima será criado a seguinte infraestrutura:

Services
 - ze-burguer: NodePort 30443
 - postgres: NodePort 30432

Deployments
 - ze-burguer: HPA (2-5 Pods) - CPU Average Usage metrics
 - postgres: 1 Pod

![K8S](./doc/infra/kubernetes.png)

## Documentação da API

[Collection_Postman](./doc/apis/Ze_burguer.postman_collection.json) (necessário importar as ENVS que costam na mesma pasta)

#### Base URL: http://localhost:8080

#### Doc Swagger: http://localhost:8080/docs/index.html

#### Passo a passo para execução das APIs

  - Categorias dos produtos:
    - | Id | Categoria  |
      | -- | ---------- |
      | 1  | Lanche     |
      | 2  | Bebida     |
      | 3  | Acompanhamento   |
      | 4  | Sobremesa  |

  - Status do pedido:
    - | Status | 
      | ------ |
      | RECEIVED  |
      | AWAITING_PAYMENT  |
    - | PAID   |
    - | PREPARING  | 
      | READY  | 
      | FINISHED |

 - Passo 1: Cadastrar os Produtos desejados (`http://localhost:port/products`)
 - Passo 2: Cadastrar o cliente (etapa opcional) (`http://localhost:port/clients`)
 - Passo 3: Criar um pedido com os produtos cadastrados (pode ou não informar o id do cliente cadastrado) (`http://localhost:port/orders`)

## Sonar

![Sonar](./doc/sonar.png)

## OWASP ZAP

Relatório se encontra na pasta ./doc/relatorio/orders-report.html
![OWASP ZAP](./doc/relatorio/owasp.png)

## Relatorio RIPD

![Relatorio RIPD](./doc/RIPD.docx)