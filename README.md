# MicroKit

This Microservices Starter Template provides a foundation for building scalable and maintainable microservices-based applications using Go. It includes basic implementations of essential microservices patterns and practices.

## Table of Contents
- [MicroKit](#microkit)
  - [Table of Contents](#table-of-contents)
  - [Introduction](#introduction)
  - [Project Structure](#project-structure)
  - [Prerequisites](#prerequisites)
  - [Getting Started](#getting-started)
  - [Services](#services)
    - [API Gateway](#api-gateway)
    - [Service A](#service-a)
    - [Service B](#service-b)
  - [Shared Components](#shared-components)
  - [Configuration](#configuration)
  - [Database](#database)
  - [Message Queue](#message-queue)
  - [Service Discovery](#service-discovery)
  - [Logging](#logging)
  - [Testing](#testing)
  - [CI/CD](#cicd)
  - [Deployment](#deployment)
  - [Monitoring and Observability](#monitoring-and-observability)
  - [Scaling](#scaling)
  - [Troubleshooting](#troubleshooting)
  - [Contributing](#contributing)
  - [License](#license)

## Introduction
This Microservices Starter Template provides a foundation for building scalable and maintainable microservices-based applications using Go. It includes basic implementations of essential microservices patterns and practices.

## Project Structure
```
microservice-starter-template/
├── .github/
│   └── workflows/
│       └── ci.yml
├── api-gateway/
│   ├── Dockerfile
│   └── main.go
├── service-a/
│   ├── Dockerfile
│   ├── main.go
│   ├── handlers/
│   ├── models/
│   └── tests/
├── service-b/
│   ├── Dockerfile
│   ├── main.go
│   ├── handlers/
│   ├── models/
│   └── tests/
├── shared/
│   ├── config/
│   ├── database/
│   ├── logger/
│   └── middleware/
├── proto/
│   └── service.proto
├── scripts/
│   ├── setup.sh
│   └── run-tests.sh
├── prometheus/
│   └── prometheus.yml
├── docker-compose.yml
├── .env.example
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

## Prerequisites
- Go 1.16 or later
- Docker and Docker Compose
- Git

## Getting Started
1. Clone the repository:
   ```
   git clone https://github.com/yourusername/microservice-starter-template.git
   cd microservice-starter-template
   ```

2. Run the setup script:
   ```
   ./scripts/setup.sh
   ```

3. Start the services:
   ```
   docker-compose up
   ```

## Services

### API Gateway
The API Gateway serves as the entry point for client requests. It routes requests to appropriate microservices and handles cross-cutting concerns like authentication and rate limiting.

Key files:
- `api-gateway/main.go`: Main entry point
- `api-gateway/Dockerfile`: Docker configuration

### Service A
Service A is an example microservice that demonstrates basic CRUD operations.

Key files:
- `service-a/main.go`: Main entry point
- `service-a/handlers/`: Request handlers
- `service-a/models/`: Data models
- `service-a/tests/`: Unit and integration tests
- `service-a/Dockerfile`: Docker configuration

### Service B
Service B is another example microservice, demonstrating inter-service communication.

Key files:
- `service-b/main.go`: Main entry point
- `service-b/handlers/`: Request handlers
- `service-b/models/`: Data models
- `service-b/tests/`: Unit and integration tests
- `service-b/Dockerfile`: Docker configuration

## Shared Components
The `shared/` directory contains components used across multiple services:

- `config/`: Configuration management using Viper
- `database/`: Database connection and ORM setup
- `logger/`: Centralized logging using Zap
- `middleware/`: Shared middleware functions

## Configuration
- Environment variables are used for configuration. See `.env.example` for available options.
- Viper is used for configuration management. See `shared/config/` for implementation.

## Database
- PostgreSQL is used as the primary database.
- GORM is used as the ORM. See `shared/database/` for connection setup.

## Message Queue
- RabbitMQ is used for asynchronous communication between services.
- See `shared/rabbitmq/` for implementation details.

## Service Discovery
- Consul is used for service discovery and registration.
- See `shared/discovery/` for implementation details.

## Logging
- Zap is used for structured logging.
- See `shared/logger/` for implementation details.

## Testing
- Unit tests and integration tests are included for each service.
- Run tests using:
  ```
  ./scripts/run-tests.sh
  ```

## CI/CD
- GitHub Actions is used for CI/CD. See `.github/workflows/ci.yml` for the workflow definition.
- The workflow includes steps for testing, building Docker images, and a placeholder for deployment.

## Deployment
- Docker images are built for each service.
- A basic deployment step is included in the CI/CD workflow, which can be customized based on your hosting environment.

## Monitoring and Observability
- Prometheus is used for metrics collection. See `prometheus/prometheus.yml` for configuration.
- Grafana can be used for visualization (not included in the template).

## Scaling
- Services can be scaled horizontally by adjusting the number of containers in `docker-compose.yml` or your production orchestration tool (e.g., Kubernetes).

## Troubleshooting
- Check logs using `docker-compose logs [service_name]`
- Ensure all required environment variables are set
- Verify network connectivity between services

## Contributing
Contributions are welcome! Please read our contributing guidelines and code of conduct before submitting pull requests.

## License
This project is licensed under the MIT License - see the LICENSE file for details.