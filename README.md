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
```

## Features

- [X] Domein-Driven Desing
- [X] Dockerize
- [X] Rest API
- [X] Fiber
- [X] Swagger
- [ ] Kafka
- [ ] Kubernetes Deployment File
- [ ] Unit Test
- [ ] Integration testing
- [ ] E2E testing

## Project Layout

```Bash
$ tree 
.
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
│   │   ├── handler
│   │   │   ├── comment_handler.go
│   │   │   └── task_handler.go
│   │   └── router
│   │       └── router.go
│   ├── domain
│   │   ├── model
│   │   │   ├── comment.go
│   │   │   └── task.go
│   │   ├── repository
│   │   │   ├── CommentRepository.go
│   │   │   └── TaskRepository.go
│   │   └── service
│   │       ├── CommentService.go
│   │       ├── CommentServiceRepository.go
│   │       ├── TaskService.go
│   │       └── TaskServiceRepository.go
│   └── infrastructure
│       ├── broker
│       │   ├── consumer
│       │   ├── interfaces
│       │   └── producer
│       ├── config
│       │   └── config.go
│       └── persistence
│           ├── connect.go
│           └── repository
│               ├── CommentRepository.go
│               └── TaskRepository.go
├── docker-compose.yml
├── Dockerfile
├── Makefile
├──  start.sh
├── LICENSE
├── README.md
├── go.mod
├── go.sum
└──  main.go
```
