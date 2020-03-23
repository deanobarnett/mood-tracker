.PHONY: help build build.test build.dependencies lint logs run shell start stop test test-with-artifacts test.shell docker.test docker.test-with-artifacts

default: help

GIT_SHA = $(shell git rev-parse HEAD)

help: ## Show this help
	@echo "Mood Tracker"
	@echo "========"
	@echo
	@fgrep -h " ## " $(MAKEFILE_LIST) | fgrep -v fgrep | sed -Ee 's/([a-z.]*):[^#]*##(.*)/\1##\2/' | column -t -s "##"

build: ## Build the application
	@docker-compose exec service go install -v ./cmd/...

build.test: ## Build the test container
	@docker-compose build test

lint: ## Lint the application
	@docker-compose run --rm test golangci-lint run

logs: ## Show the application logs
	@docker-compose logs --follow service

run: ## Run the application locally in interactive mode
	@docker-compose up --build service

shell: running ## Create a shell in the application container
	@docker-compose exec service /bin/bash

start: ## Run the application locally in the background
	@docker-compose up --build -d service

stop: ## Stop the running application
	@docker-compose down

test: ## Test the application
	@docker-compose run --rm test

test.shell: ## Open a shell in the test container
	@docker-compose run --rm test /bin/bash

docker.test: ## Test the package from within a container
	@go test -cover $(if $(path), $(path), ./...)
