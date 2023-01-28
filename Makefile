#!make
include .env
export $(shell sed 's/=.*//' .env)

.SILENT:


build:
	go mod download && go build -o ./.bin/app ./cmd/main.go

run: build
	./.bin/app

migrate:
	migrate -path ./schema -database postgres://pguser:${DB_PASSWORD}@localhost:5431/devdb?sslmode=disable up

migrate-down:
	migrate -path ./schema -database postgres://pguser:${DB_PASSWORD}@localhost:5431/devdb?sslmode=disable down
