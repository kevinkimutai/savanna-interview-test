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
- Jenkins (CICD)
- Next JS (frontend)
- [Testify](github.com/stretchr/testify) (Testing package for assertions and mocks)

## Project Overview

This project demonstrates the development of a backend Golang REST service using Jenkins as CICD pipeline ensuring automated testing and deployment to Kubernetes using Helm as well as RabbitMQ in ensuring asynchronous communication between services.
After successful order.message is published to rabbitmq and sms service subscribes to queue (sms service sends message content using Africas Talking Api)
Use sandbox to see message.

## Features Implemented

- Backend Service: Developed a robust GoLang backend service that handles CRUD operations for customer,product and order data via REST APIs.
- Database: Designed and deployed a PostgreSQL database schema optimized for storing customer,product,order and their relationships.
- Frontend Integration: Integrated a responsive Next.js frontend that communicates with the backend APIs to display and manipulate data.
- CI/CD Pipeline: Implemented a Jenkins pipeline configured to automate building, testing, and deploying the application to Kubernetes clusters.
- RabbitMQ for asynchronous communication between services.

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

## Database

migration files are located at ./migrations
Schema:

![db](https://firebasestorage.googleapis.com/v0/b/creadable-22c39.appspot.com/o/Screenshot%20from%202024-06-27%2013-28-58.png?alt=media&token=0044a866-0d31-4613-83c1-a25b83d00965)

## Jenkins CICD Pipeline Config

![pipeline](https://miro.medium.com/v2/resize:fit:680/1*bCXV9mUXqOww-xubDRJf4g.png)

The pipeline is defined in a Jenkinsfile located in the root of this repository.
./Jenkinsfile

- Checkout: Fetches the code from the GitHub repository on every merge to main branch.
- Build: Compiles the GoLang application and generates the executable.
- Test: Executes unit tests for the GoLang application.
- SonarQube Analysis: Runs static code analysis using SonarQube to ensure code quality and security standards.
- Building Docker Image: Builds a Docker image of the application using a specified Dockerfile.
- Deploy Image: Pushes the Docker image to Docker Hub with version tags and deploys it to Kubernetes using Helm charts.
- Remove Unused Docker Image: Cleans up Docker images after deployment to optimize storage usage.
- Kubernetes Deploy: Deploys the application to Kubernetes cluster using Helm.
- Cleanup: Cleans up the workspace by deleting temporary files and artifacts.

## Frontend

Frontend implemented using Next js,shadcn for styling and Auth0 for authentication and authorization

- Github link: [Frontend](https://github.com/kevinkimutai/savanna_frontend)
- Live link: [Live](https://savanna-pi.vercel.app)

![frontend-snippet](https://firebasestorage.googleapis.com/v0/b/creadable-22c39.appspot.com/o/Screenshot%20from%202024-06-27%2013-59-13.png?alt=media&token=fc5745a7-b032-4b24-a479-0e90f76f9631)

## Swagger Documentation

Swagger middleware for documentation

```
	cfg := swagger.Config{
		BasePath: "/api/v1/",
		FilePath: "./docs/swagger.json",
		Path:     "swagger",
		Title:    "Swagger Order API Docs",
	}
	app.Use(swagger.New(cfg))
```

Find swagger documentation at **/api/v1/swagger**
![swagger docs](https://firebasestorage.googleapis.com/v0/b/creadable-22c39.appspot.com/o/Screenshot%20from%202024-06-10%2019-34-21.png?alt=media&token=ad34f93d-0e6e-4d46-8e3d-53c9417f6764)

## SMS service

Service responsible for sending messages using the AfricasTalking Api

- Github link: [sms](https://github.com/kevinkimutai/savanna-interview-sms)

## Web Security

All security middleware are located at the server adapter `./internal/adapters/server/server.go`

- cors:Only requested URLS,methods and headers allowed.
- CSRF (Cross-Site Request Forgery) Protection:ensures requests originate from trusted sources
- Helmet: sets HTTP headers to secure the application.
  -Rate Limiting: Requests are rate-limited to prevent abuse and ensure fair usage of resources.

## Architecture

Hexagonal Architecture (Ports and Adapters)
