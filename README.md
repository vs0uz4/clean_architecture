# Go Clean Architecture

> [!IMPORTANT]  
> Para poder executar o projeto contido neste repositório é necessário que se tenha o Go instalado no computador. Para maiores informações siga o site <https://go.dev/>

## Desafio GoLang Pós GoExpert - Clean Architecture

Este projeto é parte de um desafio da Pós GoExpert, nele são cobertos os conhecimentos em Go Rotines, channels, contextos, tratamentos de erros, packages, Clean Architecture, gRPC, GraphQL, APIRest, Eventos, DI.

O Desafio consiste em entregar um `endpoint` para listagem das Ordens de Compras realizadas.

### Requisitos a serem seguidos

A listagem de `orders` precisa ser realizada da seguinte maneira:

* O `endpoint` REST atenderá na rota (GET /order);
* Um `service` ListOrders com gRPC;
* Uma `query` ListOrders em GraphQL.

> Devemos criar as `migrations` necessárias e também o arquivo `api.http` com as requests para criar e listar as `orders`.

Para a criação do banco de dados, devemos utilizar o Docker (Dockerfile|docker-compose.yaml), com isso ao rodar o comando `docker compose up` tudo deverá subir, já deixando todo o ambiente de INFRA preparado, levantando o RabbitMQ e o MySQL.

Criar documentação `README.md` descrevendo passo a passo como executar a aplicação e em quais portas cada serviço deverá responder.

### Executando os Sistemas

Para executar o sistema, devemos primeiramente provisionar a INFRA necessária, para isto precisamos rodar o seguinte comando abaixo:

```shell
❯ docker-compose up
```

> O comando acima deve ser executado a partir da pasta raiz do projeto, onde encontramos o arquivo `docker-compose.yaml`

Após toda a INFRA necessária estar de pé, bastar rodarmos o projeto. Para isto, estando na pasta raiz do projeto, execute os seguintes comandos:

```shell
❯ cd cmd/ordersystem
❯ go run main.go wire_gen.go
```

Na janela do terminal, você deverá ver uma mensagem parecida com o exemplo abaixo:

```shell
❯ go run main.go wire_gen.go
Starting web server on port :8000
Starting gRPC server on port 50051
Starting GraphQL server on port 8080
```

### Informações dos Serviços

**APIRestful - [Porta 8000]**
GET /order  - Listagem de todas as `orders`
POST /order - Criação de uma `order`

**GraphQL - [Porta 8080]**
Query
    - orders: [Order!]!
Mutation
    - createOrder(input: OrderInput): Order

WIP