# Savanna Informatics Backend Interview

From the [Interview Question](https://firebasestorage.googleapis.com/v0/b/creadable-22c39.appspot.com/o/Savannah%20Informatics%20Back%20End%20Dev-%20Challenge.pdf?alt=media&token=051b7a05-90d3-4dce-9403-c4ab493a9602).

## Tech & Tools Used

- Postgres (Database)
- Auth0 (Authorization/Authentication)
- Golang (Web API golang-fiber)
- RabbitMQ (Message Queues)
- Node JS (Backend for sending messages)
- SQLC (Generate SQLC type safe code)
- Golang Migrate (Golang SQL migrate tool)

## Getting Started

- Setup your Auth0 at [Auth0](https://auth0.com/docs/libraries#backend)
- Run `make postgresdb` :setup database
- Run `make createdb` :create db
- Run `make migrateup` :migrate database migrations to get latest db schema.Check number of migrtations.
- Use your own .env file(Check Makefile for used values) Ensure you have the following environmental variables
  - APP_PORT
  - POSTGRES_USERNAME
  - POSTGRES_PASSWORD
  - DB_PORT
  - AUTH0_URL
  - AUTH0_CLIENTID
  - AUTH0_CLIENT_SECRET
  - RABBITMQ_SERVER
- Run `go run cmd/main.go`

## Swagger Documentation

Find swagger documentation at **/api/v1/swagger**
![swagger docs](https://firebasestorage.googleapis.com/v0/b/creadable-22c39.appspot.com/o/Screenshot%20from%202024-06-10%2019-34-21.png?alt=media&token=ad34f93d-0e6e-4d46-8e3d-53c9417f6764)

## Architecture

Hexagonal Architecture (Ports and Adapters)

## Whats Remaining

### Deployment

- k8s using Helm/kops
- CICD using jenkins
- code analysis using sonarqube in CICD

## Other

- **Docker Image** is at kevinkimutai/orderapp1.0
- SMS repo at [sms service](https://github.com/kevinkimutai/savanna-interview-sms)
