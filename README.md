# MicroKit

This Microservices Starter Template provides a foundation for building scalable and maintainable microservices-based applications using Go. It includes basic implementations of essential microservices patterns and practices.

## Table of Contents
- [MicroKit](#microkit)
  - [Table of Contents](#table-of-contents)
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

## Prerequisites
- Go 1.16 or later
- Docker and Docker Compose
- Git

## Getting Started
1. Clone the repository:
   ```
   git clone MuxSphere/microkit.git
   cd microkit
   ```

2. Run the setup script:
   ```
   ./scripts/setup.sh
   ```
This script initializes the project, installs dependencies, and sets up initial configurations.

3. Start the services:
   ```
   docker-compose up
   ```
This command builds and starts all the services defined in docker-compose.yml.

## Services

### API Gateway
The API Gateway serves as the entry point for client requests. It routes requests to appropriate microservices and handles cross-cutting concerns like authentication and rate limiting.

Key files:
- `api-gateway/main.go`: Main entry point
- `api-gateway/Dockerfile`: Docker configuration for building the API Gateway image

The API Gateway uses a reverse proxy to route requests to the appropriate services based on the URL path.

### Service A
Service A is an example microservice that demonstrates basic CRUD operations.

Key files:
- `service-a/main.go`: Main entry point
- `service-a/handlers/`: Request handlers
- `service-a/main_test.go`: Unit and integration tests
- `service-a/Dockerfile`: Docker configuration

### Service B
Service B is another example microservice, demonstrating inter-service communication.

Key files:
- `service-a/main.go`: Main entry point
- `service-a/handlers/`: Request handlers
- `service-a/main_test.go`: Unit and integration tests
- `service-a/Dockerfile`: Docker configuration

## Shared Components
The `shared/` directory contains components used across multiple services:

- `config/`: Configuration management using Viper
- `database/`: Database connection and ORM setup
- `logger/`: Centralized logging using Zap structured and efficient logging

## Configuration
- Environment variables are used for configuration. See `.env.example` for available options.
- Viper is used for configuration management. See `shared/config/` for implementation.

To configure your services, copy  `.env.example`  to `.env` and modify the values as needed. The application will automatically load these environment variables.

## Database
- PostgreSQL is used as the primary database.
- GORM is used as the ORM. See `shared/database/` for connection setup.

## Message Queue
- RabbitMQ is used for asynchronous communication between services.
- See `shared/rabbitmq/` for implementation details.

RabbitMQ allows for decoupled, scalable communication between services. It's particularly useful for handling background tasks and ensuring reliable message delivery.

## Service Discovery
- Consul is used for service discovery and registration.
- See `shared/discovery/` for implementation details.

When a service starts, it registers itself with Consul. Other services can then discover and communicate with it dynamically.

To use Consul:

 1. Ensure Consul is running (it's included in the docker-compose.yml)
 2. Services automatically register themselves on startup
 3. Use the Consul API or DNS interface to discover other services

## Logging
- Zap is used for structured logging.
- See `shared/logger/` for implementation details.

Zap provides fast, structured logging. Use the logger like this:
```
logger.Info("This is an info log", zap.String("key", "value"))
logger.Error("This is an error log", zap.Error(err))
```

## Testing
- Unit tests and integration tests are included for each service.
- Run tests using:
  ```
  ./scripts/run-tests.sh
  ```

## CI/CD
- GitHub Actions is used for CI/CD. See `.github/workflows/ci.yml` for the workflow definition.
- The workflow includes steps for testing, building Docker images, and a placeholder for deployment.

To use this workflow:

 1. Replace `your-dockerhub-username` with your actual DockerHub username.
 2. Add `DOCKERHUB_USERNAME` and `DOCKERHUB_TOKEN` secrets in your GitHub repository settings.
 3. Customize the deployment step according to your preferred deployment method.

## Deployment
- Docker images are built for each service.
- A basic deployment step is included in the CI/CD workflow, which can be customized based on your hosting environment.

## Monitoring and Observability
- Prometheus is used for metrics collection. See `prometheus/prometheus.yml` for configuration.
- Grafana can be used for visualization (not included in the template).

When you have Prometheus set up and services are exposing metrics, configure some basic dashboards in Grafana. Do this manually through the Grafana UI:

 1. Access Grafana at http://localhost:3000 (default credentials: admin/admin)
 2. Add Prometheus as a data source (http://prometheus:9090)
 3. Create a new dashboard with panels for:
   - Total HTTP requests
   - HTTP request rate
   - HTTP response codes
   - Service uptime

## Scaling
- Services can be scaled horizontally by adjusting the number of containers in `docker-compose.yml` or your production orchestration tool (e.g., Kubernetes).

To scale a service using Docker Compose:
```
  docker-compose up --scale service-a=3 --scale service-b=2
```
This command would start 3 instances of Service A and 2 instances of Service B.

## Troubleshooting
- Check logs using `docker-compose logs [service_name]`
- Ensure all required environment variables are set
- Verify network connectivity between services

Common issues:

 1. Services failing to start: Check the logs for error messages
 2. Database connection issues: Ensure the database is running and the connection string is correct
 3. Inter-service communication failures: Check if services are registered correctly with Consul

## Contributing
Contributions are welcome! 

## License
This project is licensed under the MIT License 