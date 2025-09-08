# Event Driven Architecture API Design With Kafka, Django, Gin, PostgreSQL, MongoDB, and Debezium

This repository contains a sample implementation of an **Event Driven Architecture (EDA)** using **Apache Kafka**. The project demonstrates how to design and implement APIs that leverage Kafka for event-driven communication between microservices built with **Django** and **Gin (Go)**, with persistence in **PostgreSQL** and **MongoDB**, and real-time change data capture powered by **Debezium**.

## Table of Contents

* [Introduction](#introduction)
* [Architecture Overview](#architecture-overview)
* [Prerequisites](#prerequisites)
* [Installation](#installation)
* [Configuration](#configuration)
* [Usage](#usage)
* [API Endpoints](#api-endpoints)
* [Sample Requests and Responses](#sample-requests-and-responses)
* [Monitoring](#monitoring)
* [Contributing](#contributing)
* [License](#license)

## Introduction

**Event Driven Architecture (EDA)** is a design paradigm where services communicate through events instead of direct API calls.

* **Apache Kafka** acts as the central event bus
* **Django service**: Handles administrative API operations and produces events to Kafka
* **Gin (Go) service**: Consumes Kafka events and exposes query endpoints
* **PostgreSQL**: Used as the relational data store
* **MongoDB**: Used as a NoSQL data store for flexible queries
* **Debezium**: Captures database changes from PostgreSQL and streams them into Kafka topics

This project showcases how to integrate all these components into a cohesive EDA-based API system.

## Architecture Overview

```
+----------------+       +-----------+       +----------------+
| Django (Python)| --->  |   Kafka   | --->  | Gin (Go) API   |
|  (Producer)    |       | (Event Bus)|      | (Consumer)     |
+----------------+       +-----------+       +----------------+
        |                                          |
        |                                          |
  PostgreSQL <---- Debezium ----> Kafka <----> MongoDB
```

### Data Flow

1. **Django** produces events to Kafka whenever a write operation occurs (create, update, delete)
2. **Debezium** streams database changes (CDC) from PostgreSQL into Kafka topics
3. **Gin (Go)** consumes events from Kafka and updates MongoDB for efficient retrieval
4. Clients interact with both APIs depending on whether they need to **write** (Django) or **read** (Gin)

## Prerequisites

Before running the project, ensure you have:

* [Docker](https://www.docker.com/get-started) and [Docker Compose](https://docs.docker.com/compose/install/) installed
* Python **3.8+** installed
* Go **1.18+** installed
* Basic knowledge of Kafka, REST APIs, and databases

## Installation

1. Clone the repository:

   ```bash
   git clone <repository-url>
   cd todo-eda
   ```

2. Build and start the Docker containers:

   ```bash
   docker-compose up -d
   ```

   This sets up:
   * Kafka + Zookeeper
   * Django API service
   * Gin API service
   * PostgreSQL
   * MongoDB
   * Debezium

3. Wait for all services to be healthy. You can check the status with:

   ```bash
   docker-compose ps
   ```

4. Services will be available at:
   * Django API → `http://localhost:8000`
   * Gin API → `http://localhost:8080`
   * PostgreSQL → `localhost:5432`
   * MongoDB → `localhost:27017`
   * Kafka → `localhost:9092`

## Configuration

Configuration files are located in the following locations:

* **Kafka settings** → `docker-compose.yml`
* **Django settings** → `todo_cud/settings.py`
* **Gin settings** → `config/config.go`
* **Debezium connector config** → `debezium/connector.json`

### Environment Variables

Key environment variables you may need to adjust:

```bash
# Kafka
KAFKA_BROKER_URL=kafka:9092

# PostgreSQL
POSTGRES_DB=todos
POSTGRES_USER=user
POSTGRES_PASSWORD=password

# MongoDB
MONGO_URL=mongodb://mongodb:27017
MONGO_DB=todos
```

Make sure to adjust **Kafka broker addresses** and **database credentials** if needed.

## Usage

### Basic Workflow

1. **Produce events** → Send `POST`, `PUT`, `DELETE` requests to the Django API
2. **Consume events** → Gin API automatically updates its read models from Kafka
3. **Query data** → Use the Gin API to fetch todos from MongoDB
4. **Change Data Capture** → Debezium continuously captures PostgreSQL changes and pushes them to Kafka

### Testing the Flow

1. Create a todo via Django API
2. Verify the event appears in Kafka
3. Check that the Gin API can query the updated data
4. Monitor Debezium for CDC events

You can also monitor Kafka topics using [Kafka Tool](http://www.kafkatool.com/) or:

```bash
docker exec -it kafka kafka-console-consumer --topic todos --from-beginning --bootstrap-server kafka:9092
```

## API Endpoints

### Django API (Producer/Write Service)

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/api/todos/` | Create a new todo (produces Kafka event) |
| `PUT` | `/api/todos/{id}/` | Update a todo (produces Kafka event) |
| `DELETE` | `/api/todos/{id}/` | Delete a todo (produces Kafka event) |
| `GET` | `/api/todos/` | List todos (direct from PostgreSQL) |
| `GET` | `/api/todos/{id}/` | Get todo by ID (direct from PostgreSQL) |

### Gin API (Consumer/Read Service)

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/api/todos/` | Retrieve all todos (from MongoDB) |
| `GET` | `/api/todos/{id}/` | Retrieve a todo by ID (from MongoDB) |
| `GET` | `/health` | Health check endpoint |

## Sample Requests and Responses

### Creating a Todo (Django)

**Request:**
```bash
curl -X POST http://localhost:8000/todos/ \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Learn Kafka",
    "description": "Study event-driven architecture",
    "completed": false
  }'
```

**Response:**
```json
{
  "status": "success",
  "status_code": 202,
  "data": {},
  "message": "Todo create successfully queued"
}
```

**Kafka Event Produced:**
```json
{
  "operation": "create", 
  "data": {
    "title": "Check Order Ops New Test", 
    "description": "Check Order Ops New 001"
  },
  "timestamp": "2025-09-08T08:07:36.506264"
}
```

### Querying Todos (Gin)

**Request:**
```bash
curl -X GET http://localhost:8080/todos/
```

**Response:**
```json
[
    {
        "id": 16,
        "title": "Link the golang debezium consumer New",
        "description": "Link the golang debezium consumer Ok",
        "created_at": "2025-09-06T12:56:56.855869Z",
        "updated_at": "2025-09-06T12:56:56.855905Z",
        "completed": false
    },
    {
        "id": 18,
        "title": "Check Order Ops New",
        "description": "Check Order Ops New 001",
        "created_at": "2025-09-08T08:01:47.627717Z",
        "updated_at": "2025-09-08T08:01:47.628237Z",
        "completed": false
    },
    {
        "id": 19,
        "title": "Check Order Ops New Test",
        "description": "Check Order Ops New 001",
        "created_at": "2025-09-08T08:07:36.544031Z",
        "updated_at": "2025-09-08T08:07:36.544061Z",
        "completed": false
    }
]
```

### Updating a Todo (Django)

**Request:**
```bash
curl -X PUT http://localhost:8000/todos/1/ \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Learn Kafka",
    "description": "Master event-driven architecture",
    "completed": true
  }'
```

**Kafka Event Produced:**
```json
{
  "operation": "update", 
  "data": {
    "title": "Check Order Ops New Test", 
    "description": "Check Order Ops New 001"
  },
  "timestamp": "2025-09-08T08:07:36.506264"
}
```

### Debezium CDC Event

When PostgreSQL data changes, Debezium produces events like:

```json
{
  "schema": {...},
  "payload": {
    "before": null,
    "after": {
      "id": 1,
      "title": "Learn Kafka",
      "description": "Master event-driven architecture",
      "completed": true
    },
    "source": {
      "version": "1.9.5.Final",
      "connector": "postgresql",
      "name": "postgres-connector",
      "ts_ms": 1693648200000,
      "snapshot": "false",
      "db": "todoapp",
      "table": "todos"
    },
    "op": "u",
    "ts_ms": 1693648200456
  }
}
```

## Monitoring

### Kafka Topics

Monitor the following Kafka topics:

* `todos` - Application events from Django
* `postgres.public.todos` - CDC events from Debezium

### Health Checks

* Django: `http://localhost:8000/health/`
* Gin: `http://localhost:8080/health`
* Kafka UI (if enabled): `http://localhost:8081`

### Debugging

Check service logs:

```bash
# View all logs
docker-compose logs

# View specific service logs
docker-compose logs django
docker-compose logs gin-service
docker-compose logs kafka
```

## Development

### Local Development Setup

For local development without Docker:

1. Start Kafka and databases:
   ```bash
   docker-compose up kafka zookeeper postgres mongodb
   ```

2. Run Django service:
   ```bash
   cd django-service
   pip install -r requirements.txt
   python manage.py migrate
   python manage.py runserver
   ```

3. Run Gin service:
   ```bash
   cd gin-service
   go mod download
   go run main.go
   ```

### Running Tests

```bash
# Django tests
cd django-service
python manage.py test

# Go tests
cd gin-service
go test ./...
```

## Contributing

To contribute:

1. Fork the repository
2. Create a feature branch:
   ```bash
   git checkout -b feature/your-feature-name
   ```
3. Make your changes and add tests
4. Commit your changes:
   ```bash
   git commit -m "Add your feature description"
   ```
5. Push to your branch:
   ```bash
   git push origin feature/your-feature-name
   ```
6. Open a Pull Request

### Guidelines

* Follow the existing code style
* Add tests for new features
* Update documentation as needed
* Ensure all services work together properly

## Troubleshooting

### Common Issues

1. **Kafka connection errors**: Ensure Kafka is fully started before other services
2. **Database connection issues**: Check PostgreSQL and MongoDB are accessible
3. **Debezium not capturing changes**: Verify PostgreSQL has WAL logging enabled
4. **Event consumption delays**: Check Kafka consumer group status

### Useful Commands

```bash
# Reset Kafka consumer groups
docker exec -it kafka kafka-consumer-groups --bootstrap-server kafka:9092 --group gin-consumer --reset-offsets --to-earliest --topic todos --execute

# Check Kafka topics
docker exec -it kafka kafka-topics --list --bootstrap-server kafka:9092

# View PostgreSQL logs
docker-compose logs postgres
```

## License

This project is licensed under the **MIT License**. See the `LICENSE` file for details.

---

## Next Steps

* Add authentication and authorization
* Implement event sourcing patterns
* Add more complex business logic
* Set up monitoring and alerting
* Deploy to production environment