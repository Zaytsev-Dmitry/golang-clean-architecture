# Golang Clean Architecture

![Go Version](https://img.shields.io/badge/Go-1.24-blue)
![Visitors](https://api.visitorbadge.io/api/visitors?path=https://github.com/Zaytsev-Dmitry/golang-clean-architecture&label=Repository%20Visits&countColor=%230c7ebe&style=flat&labelStyle=none)

🚀 **Golang Clean Architecture** is a highly scalable, maintainable, and testable project structure following SOLID principles and best practices. This repository is designed to help developers build robust applications with a clean separation of concerns.

## ✨ Features
- **Modular Structure** – Clean and organized codebase.
- **SOLID Principles** – Ensures maintainability and flexibility.
- **Separation of Concerns** – Divides the project into well-defined layers.
- **Testing Friendly** – Unit and integration tests supported.

### **Upcoming Features ✨**

![Apache Kafka](https://img.shields.io/badge/Apache%20Kafka-000?style=for-the-badge&logo=apachekafka)
![RabbitMQ](https://img.shields.io/badge/Rabbitmq-FF6600?style=for-the-badge&logo=rabbitmq&logoColor=white)
![MongoDB](https://img.shields.io/badge/MongoDB-%234ea94b.svg?style=for-the-badge&logo=mongodb&logoColor=white)
![Redis](https://img.shields.io/badge/redis-%23DD0031.svg?&style=for-the-badge&logo=redis&logoColor=white)
![GithubActions](https://img.shields.io/badge/GitHub_Actions-2088FF?style=for-the-badge&logo=github-actions&logoColor=white)

## ⚡ Tech Stack

- **[Golang 1.24+](https://golang.org/)**
- **[Cleanenv](https://github.com/ilyakaznacheev/cleanenv/)**
- **[Gin](https://github.com/gin-gonic/gin/)**
- **[Sqlx](https://github.com/jmoiron/sqlx/)**
- **[Oapi-codegen](https://github.com/oapi-codegen/runtime/)**
- **[Postgres](https://www.postgresql.org/)**

## 📂 Project Structure
```plaintext
📦project-root                     # Root directory of the project
├── api                            # API interfaces of the project
│   └── http                       # HTTP API (Swagger, OpenAPI, and other specifications)
│       ├── note-api-web.gen.go    # Automatically generated API code
│       ├── note-api-web.yml       # API description in YAML format
│       └── swagger.go             # Swagger configuration file
├── cmd                            # Entry points of the application
├── configs                        # Application configuration files
├── data                           # Directory containing data needed for local deployment
│   └── postgres                   # Data related to the PostgreSQL database
├── docs                           # Project documentation
├── internal                       # Internal business logic (not accessible to external modules)
│   ├── app                        # Application logic
│   │   ├── domain/                # Core domain entities (models)
│   │   ├── services/              # Services and business logic
│   │   ├── usecases/              # Business use cases implementation (Application Layer)
│   │   ├── ports                  # Interfaces for interaction between layers
│   │   │   ├── in                 # Input ports
│   │   │   │   ├── delegate/      # Delegation of calls to business logic
│   │   │   └── out                # Output ports (database, external services)
│   │   │       └── dao            # Data Access Layer (database operations)
│   │   │           ├── dto/       # Data Transfer Objects (DTO)
│   │   │           ├── queries/   # SQL queries
│   │   │           └── repository # Repository pattern implementation
│   ├── infrastructure             # Infrastructure services
│   │   └── transport              # Transport layer (REST, gRPC)
│   │       ├── http               # HTTP API
│   │       │   ├── controllers/   # Controllers for handling HTTP requests
│   │       │   ├── handlers/      # Request handling logic
│   │       │   ├── middleware/    # Middleware (logging, authorization, etc.)
│   │       │   └── presenters/    # Data transformation for API response format
│   │       └── grpc/              # gRPC API (if used)
├── pkg                            # Packages reusable across multiple projects
│   ├── errors/                    # Global error handling
│   ├── utils/                     # Utility functions
│   ├── config_loader/             # Configuration loading and parsing
│   └── marshalling/               # Data serialization/deserialization
└── third_party                    # External dependencies and integrations
    └── swagger_ui/                # Swagger UI files for API documentation
```

## 🚀 Getting Started
### Prerequisites
- [Golang 1.24+](https://golang.org/)
- Docker (optional, for containerization)

### Installation
```sh
    git clone https://github.com/your-repo/golang-clean-architecture.git
    cd golang-clean-architecture
    go mod tidy
    go run cmd/main.go
```

### Running Tests
```sh
    go test ./... (in progess now)
```

## 🤝 Contributing

We ❤️ contributions! If you want to improve this project, feel free to:
1. Fork the repository
2. Create a new feature branch
3. Make your changes
4. Submit a pull request 🚀

---

## 📬 Contact & Collaboration

Have questions, ideas, or just want to connect? Let’s chat! 📩

[![Gmail](https://img.shields.io/badge/Gmail-D14836?style=for-the-badge&logo=gmail&logoColor=white)](mailto:zaytsev.dmitry9228@gmail.com)  
[![Telegram](https://img.shields.io/badge/Telegram-2CA5E0?style=for-the-badge&logo=telegram&logoColor=white)](https://t.me/zaytsev_dv)

---

### ⭐ Don't forget to give this repository a star if you find it useful!

