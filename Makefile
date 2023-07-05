DOCKER_COMPOSE=docker-compose
APP_NAME?=postgres
DATABASE_TESTS_URL=postgres://postgres:postgres@db:5432/dataloader?sslmode=disable


.DEFAULT_GOAL := build

build:
	@echo "BUILDING THE APP"
	-$(DOCKER_COMPOSE) build $(APP_NAME)

build:
	-$(DOCKER_COMPOSE) up --build -d

run:
	-$(DOCKER_COMPOSE) up go-app

stop:
	@echo "STOPPING CONTAINERS"
	-$(DOCKER_COMPOSE) stop

down:
	@echo "REMOVING CONTAINERS"
	-$(DOCKER_COMPOSE) down

remove:
	@echo "REMOVING CONTAINERS AND VOLUMES"
	-$(DOCKER_COMPOSE) down -v

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

setup: build up-db

execute:
	go run cmd/main.go

test:
	-$(DOCKER_COMPOSE) up test

.PHONY:test