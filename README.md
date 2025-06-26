# 📘 Tutorial: RESTful API Server for a Bookmark Manager in Go

## 🧠 Overview

This project implements a RESTful API server for managing bookmarks using Go. It provides endpoints to create, retrieve, update, and delete bookmarks, along with structured logging and OpenAPI documentation.


## 🔧 Technologies Used

* Go (latest stable version)
* `gin-gonic/gin`: HTTP web framework
* `uber-go/zap`: Structured logging
* `swaggo/swag`: OpenAPI (Swagger) generation
* `net/http`: Underlying HTTP protocol support

📂 Project Structure
```
bookmark-api/
├── cmd/
│   └── server/               # Entrypoint for the service
├── internal/
│   └── bookmarks/           # Business logic and handlers
├── pkg/
│   └── middleware/          # Logging and error middleware
├── docs/                    # Swagger documentation
├── go.mod
├── main.go
└── README.md
```

## 🚀 Getting Started
### Prerequisites
* Go installed (version 1.18 or later)
* Go modules enabled
* A code editor (e.g., VSCode, GoLand)
* Postman or cURL for testing API endpoints
* Optional: Docker for containerization


### Installation
1. Clone the repository:
    ```bash
       git clone git@github.com/tazzledazzle/bookmark-manager.git
        cd bookmark-manager
    ```

2. Install dependencies:
    ```bash
       go mod tidy
    ```

3. Generate Swagger documentation:
    ```bash
       swag init --parseDependency --parseInternal
    ```

4. Run the server
    ```bash
       go run cmd/server/main.go
    ```

5. Access the API documentation at `http://localhost:8080/swagger/index.html`
6. Use Postman or cURL to test the API endpoints
