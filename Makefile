build:
		go build -v ./cmd/main.go

run:
		./main -LOGIN=postgres, -PASSWORD=123, -HOST=localhost, -PORT=5432, -DB_NAME=postgres

docker:
		docker-compose up -d

db:
		docker exec -it tic-tac-toe_postgress psql -U postgres

migrate:
		migrate -path db/migrations -database "postgresql://postgres:123@localhost:5432/postgres?sslmode=disable" -verbose $(mode)

.PHONY: build run docker db migrate


.DEFAULT_GOAL := build

