build:
	@go build -o bin/auto_school cmd/main.go

run: build
	@./bin/auto_school

migration:
	@migrate create -ext sql -dir cmd/migrations $(filter-out $@, $(MAKECMDGOALS))

migrate-up:
	@go run cmd/migations/main.go up

migrate-down:
	@go run cmd/migations/main.go down