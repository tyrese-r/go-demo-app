# Go Demo App

A simple REST API built with Go and Gin framework demonstrating user authentication, SQLite database integration, and
JWT token-based authentication.

## Features

- User registration and login
- JWT token authentication
- SQLite database with GORM ORM
- Structured logging
- Docker containerization
- GitHub Actions CI/CD pipeline

## Tech Stack

- [Go 1.23](https://golang.org/)
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/) with SQLite
- [JWT](https://github.com/golang-jwt/jwt)
- [Docker](https://www.docker.com/) & Docker Compose

## Getting Started

### Prerequisites

- Go 1.23 or higher
- Docker and Docker Compose (optional)
- Bruno API Client (optional, for testing API endpoints)

### Local Development

1. Clone the repository:
   ```bash
   git clone https://github.com/tyrese-r/go-demo-app.git
   cd go-demo-app
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

4. The server will start at http://localhost:8080

### Using Docker

1. Build and run with Docker Compose:
   ```bash
   docker-compose up -d
   ```

2. The server will be available at http://localhost:8080

## API Endpoints

### Authentication

- **Register a new user**
    - `POST /api/register`
    - Request body:
      ```json
      {
        "username": "johndoe",
        "password": "securepassword123"
      }
      ```

- **Login**
    - `POST /api/login`
    - Request body:
      ```json
      {
        "username": "johndoe",
        "password": "securepassword123"
      }
      ```
    - Response:
      ```json
      {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
      }
      ```

### User Management

- **Get User by Username**
    - `GET /api/users/:username`
    - Response:
      ```json
      {
        "id": 1,
        "username": "johndoe",
        "roles": ["user"]
      }
      ```

- **Create User**
    - `POST /api/users`
    - Request body:
      ```json
      {
        "username": "johndoe",
        "password": "securepassword123",
        "roles": ["user"]
      }
      ```

### System Information

- **Get System Stats**
    - `GET /api/stats`
    - Response:
      ```json
      {
        "uptime": "3h2m41s",
        "system": {
          "hostname": "server-name",
          "goroutines": 12,
          "memory": {
            "allocated": 12,
            "totalAlloc": 35,
            "sys": 72,
            "numGC": 7
          },
          "cpu": {
            "numCPU": 8,
            "goOS": "linux",
            "goArch": "amd64"
          }
        },
        "build": {
          "goVersion": "go1.23",
          "buildDate": "2025-02-27T10:15:30Z",
          "version": "0.1.0"
        }
      }
      ```

## Testing API Endpoints

### Using Bruno

This project includes a Bruno API collection for testing the endpoints. Bruno is an open-source API client that helps
you test your APIs locally.

1. Install [Bruno](https://www.usebruno.com/)
2. Open the `.bruno` folder in Bruno
3. The collection includes all the API endpoints configured for testing

### Using cURL

You can also test the API with cURL. Here's a typical workflow:

```bash
# 1. Register a new user
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{"username": "testuser", "password": "password123"}'

# 2. Login to get JWT token
TOKEN=$(curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"username": "testuser", "password": "password123"}' \
  | grep -o '"token":"[^"]*' | sed 's/"token":"//')

echo "Your token: $TOKEN"

# 3. Get user information
curl -X GET http://localhost:8080/api/users/testuser \
  -H "Authorization: Bearer $TOKEN"

# 4. Check system stats
curl -X GET http://localhost:8080/api/stats
```

## Project Structure

```
├── internal/
│   ├── db/               # Database configuration and connections
│   │   ├── schema/       # Database models
│   │   └── sqlite.go     # SQLite configuration
│   ├── handlers/         # HTTP handlers
│   ├── repositories/     # Data access layer
│   ├── services/         # Business logic
│   └── utils/            # Utility packages
│       ├── hello/        # Sample package with tests
│       ├── logger/       # Logging configuration
│       └── secrets/      # Environment variables handling
├── .bruno/               # Bruno API collection
├── .github/workflows/    # GitHub Actions CI/CD configuration
├── Dockerfile            # Docker configuration
├── docker-compose.yml    # Docker Compose services
├── go.mod                # Go module dependencies
├── go.sum                # Go module checksums
└── main.go               # Application entry point
```

## Development Commands

- Build: `go build -v ./...`
- Run: `go run main.go`
- Run all tests: `go test -v ./...`
- Run a specific test: `go test -v ./internal/utils/hello/hello_test.go -run TestHelloWorld`
- Lint: `golangci-lint run ./...`

## Deployment

### Automated Deployment with GitHub Actions

The app is automatically deployed to an OVH VPS using GitHub Actions when changes are pushed to the master branch. The
workflow:

1. Builds and tests the Go application
2. Creates a Docker image and pushes it to Docker Hub
3. Deploys the image to the production server at https://go-sandbox.tbertie.dev

### Setting Up the Deployment

To set up the deployment pipeline, you need to:

1. Configure GitHub Secrets:
    - `DEPLOY_SERVER_HOST`: Your server IP address
    - `DEPLOY_SERVER_USERNAME`: SSH username for your server
    - `DEPLOY_SERVER_KEY`: Private SSH key for authentication
    - `DOCKERHUB_TOKEN`: Docker Hub API token

2. Configure GitHub Variables:
    - `DOCKERHUB_USERNAME`: Your Docker Hub username

3. Server Configuration:
    - Install Docker and Docker Compose on your VPS
    - Configure a reverse proxy (like Caddy) separately on your server
    - Create the deployment directory: `/home/$USERNAME/go-demo-app`
    - Set up proper DNS records pointing to your VPS IP

The GitHub Actions workflow will:

1. Copy the docker-compose.yml to your VPS
2. Pull the latest Docker image from Docker Hub
3. Create a logs directory and .env file with necessary variables
4. Restart the application with docker-compose