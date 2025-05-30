# gocourse-challenge-03

Welcome to the `gocourse-challenge-03` repository!

This project is part of the **GoExpert** specialization course and is based on the code from [github.com/devfullcycle/20-CleanArch](https://github.com/devfullcycle/20-CleanArch).

The challenge focuses on creating a use case for listing orders using different technologies.

## Challenge Overview

In this challenge, you will implement a use case to list orders across three different interfaces:

- REST API (GET /order)
- **gRPC** Service (ListOrders)
- **GraphQL** Query (ListOrders)

Additionally, you will set up the necessary database migrations and API specifications.

## What Needed to Be Done

- Create the use case for listing orders.
- Implement the listing functionality for each technology:
  - REST endpoint (`GET /order`)
  - **gRPC** service (`ListOrders`)
  - **GraphQL** query (`ListOrders`)
- Set up database migrations and the `api.http` file with requests to create and list orders.
- Use Docker to manage the database setup, with `Dockerfile` and `docker-compose.yaml`.

## Steps to Verify the Usecases

### 1. Clone the Repository
```bash
git clone https://github.com/dchfarah/gocourse-challenge-03.git
```

### 2. Navigate to the Project Folder
```bash
cd gocourse-challenge-03
```

### 3. Install Dependencies
```bash
go mod tidy
```

### 4. Start the Services and App with Docker
```bash
sudo docker compose up -d
```

## API Endpoints and Ports

**REST API**:

Available at `http://localhost:8000`

**gRPC Service**:

Accessible at `tcp://localhost:50051`

**GraphQL**:

Endpoint at `http://localhost:8080`

## Additional Resources

### REST API

* API request specifications are located in the `api/` folder: `create_order.http` and `list_orders.http`

### gRPC Service

* To test the **gRPC** service, you can use [**Evans**](https://github.com/ktr0731/evans):

1. Run Evans:
   ```bash
   evans -r repl
   ```
2. Create some orders:
   ```bash
   call CreateOrder
   ```
   and provide the values ​​that are requested to create an order.
3. List the orders:
   ```bash
   call ListOrders
   ```

### GraphQL Query

1. Access the address `http://localhost:8080` through your browser
2. Below are examples of calls for creating and consulting orders
    ```graphql
    mutation createOrder {
        createOrder(input: {
            id: "change-id",
            Price: 10.0,
            Tax: 0.5
        }) {
            id
            Price
            Tax
            FinalPrice
        }
    }
    ```
    ```graphql
    query listOrders {
        listOrders {
            id
            Price
            Tax
            FinalPrice
        }
    }
    ```

## Notes

To view events in **RabbitMQ**, create a queue called `orders` with the `amq.direct` exchange.
