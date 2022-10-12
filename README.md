# To-Do App

## Introduction

This application is create to gain experience with CRUD operations and REST API structures.

## Quick Start

```Bash
# make file
$ make -f Makefile
$ make run
# script
$ ./start.sh
# docker
$ docker-compose up -d
# consumer
$ go run ./consumer/main.go
```

## Features

- [X] Domein-Driven Design
- [X] Dockerize
- [X] PostgreSQL
- [X] Rest API
- [X] Fiber
- [X] Swagger
- [X] Kafka
- [X] Unit Test
- [ ] Kubernetes Deployment File

## Project Layout

```Bash
$ tree 
.
├── consumer
│   ├── internal
│   │   ├── application
│   │   │   └── consume.go
│   │   ├── domain
│   │   │   ├── black_list
│   │   │   │   └── black_list.go
│   │   │   └── comment
│   │   │       ├── comment.go
│   │   │       └── comment_repository.go
│   │   ├── infrastructure
│   │   │   ├── consumer
│   │   │   │   └── reader.go
│   │   │   └── db
│   │   │       ├── connect_db.go
│   │   │       └── repository
│   │   │           └── comment_repository.go
│   │   └── service
│   │       └── consumer_service.go
│   └── main.go
├── db
│   └── migrations
│       ├── create_comments_table.sql
│       ├── create_tasks_table.sql
│       ├── delete_comments.table.sql
│       └── delete_tasks_table.sql
├── doc
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── internal
│   ├── application
│   │   └── handler
│   │       ├── comment_handler.go
│   │       ├── task_handler.go
│   │       └── task_handler_test.go
│   ├── domain
│   │   ├── comment
│   │   │   ├── comment.go
│   │   │   └── comment_repository.go
│   │   └── task
│   │       ├── task.go
│   │       └── task_repository.go
│   ├── infrastructure
│   │   ├── broker
│   │   │   └── producer
│   │   │       └── producer.go
│   │   ├── config
│   │   │   └── config.go
│   │   └── persistence
│   │       ├── connect.go
│   │       └── repository
│   │           ├── comment_repository.go
│   │           └── task_repository.go
│   └── service
│       ├── CommentService.go
│       └── TaskService.go
├── docker-compose.yml
├── Dockerfile
├── Dockerfile.test1
├── go.mod
├── go.sum
├── LICENSE
├── main.go
├── Makefile
├── README.md
└── start.sh
```
