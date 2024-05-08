DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable

network:
	docker network create employee-network

postgres:
	docker run --name postgres --network employee-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root employee_data

dropdb:
	docker exec -it postgres dropdb employee_data

migrateup:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/employee_data?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/employee_data?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/pzanwar/employee/db/sqlc Store

dockerbuild:
	docker build -t employee:latest .

dockerrun:
	docker run --name employee --network employee-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:secret@postgres:5432/employee_data?sslmode=disable" employee:latest 


.PHONY: network postgres createdb dropdb migrateup migratedown  sqlc test server mock 