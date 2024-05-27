#!make
SHELL := /bin/bash
export NOW = $(shell date +"%F %T")

run:
	@echo "$(NOW) running..."
	@go run main.go

wire:
	@echo "$(NOW) generating wire..."
	@cd app && wire
	@echo "done"

unit-test:
	@go test ./... -race -cover -count=1

docker-up:
	@docker compose up --build -d