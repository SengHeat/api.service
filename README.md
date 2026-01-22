## **2. Large Go Project (Microservices/Enterprise)**
```
myapp/
├── cmd/
│   ├── api/                    # Main API service
│   │   └── main.go
│   ├── worker/                 # Background worker
│   │   └── main.go
│   └── migrate/                # Database migration tool
│       └── main.go
├── internal/
│   ├── api/
│   │   ├── handlers/
│   │   │   ├── user/
│   │   │   │   ├── create.go
│   │   │   │   ├── get.go
│   │   │   │   ├── update.go
│   │   │   │   └── delete.go
│   │   │   └── product/
│   │   │       ├── create.go
│   │   │       └── get.go
│   │   ├── middleware/
│   │   │   ├── auth.go
│   │   │   ├── cors.go
│   │   │   └── ratelimit.go
│   │   └── routes/
│   │       └── routes.go
│   ├── domain/                 # Domain layer (business entities)
│   │   ├── user/
│   │   │   ├── entity.go      # User entity
│   │   │   ├── repository.go  # Repository interface
│   │   │   └── service.go     # Service interface
│   │   └── product/
│   │       ├── entity.go
│   │       ├── repository.go
│   │       └── service.go
│   ├── infrastructure/         # External dependencies
│   │   ├── database/
│   │   │   ├── postgres.go
│   │   │   └── migrations/
│   │   ├── cache/
│   │   │   └── redis.go
│   │   └── queue/
│   │       └── rabbitmq.go
│   ├── repository/             # Repository implementations
│   │   ├── user_postgres.go
│   │   └── product_postgres.go
│   └── service/                # Service implementations
│       ├── user_service.go
│       └── product_service.go
├── pkg/                        # Public libraries
│   ├── logger/
│   │   └── logger.go
│   ├── errors/
│   │   └── errors.go
│   ├── validator/
│   │   └── validator.go
│   └── jwt/
│       └── jwt.go
├── config/
│   ├── config.go
│   └── config.yaml
├── docker/
│   ├── Dockerfile
│   └── docker-compose.yml
├── scripts/
│   ├── build.sh
│   └── deploy.sh
├── tests/
│   ├── integration/
│   │   └── user_test.go
│   └── unit/
│       └── service_test.go
├── docs/
│   ├── api.md
│   └── swagger.yaml
├── .env
├── .gitignore
├── go.mod
├── go.sum
├── Makefile
└── README.md
```# api.service
