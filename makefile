createmigration:
	migrate create -ext sql -dir db/migrations -seq init_schema
image:
	docker pull postgres:12-alpine
postgres:
	docker run --name postgres -p 5440:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=1234 -d postgres:12-alpine
createdb:
	docker exec -it postgres createdb --username=root --owner=root education
dropdb:
	docker exec -it postgres dropdb education
migrateup:
	migrate -path  db/migrations -database "postgresql://root:1234@localhost:5440/education?sslmode=disable" -verbose up
migratedown:
	migrate -path  db/migrations -database "postgresql://root:1234@localhost:5440/education?sslmode=disable" -verbose down
sqlc:
	sqlc generate
start:
	docker start postgres
test:
	go test -v -cover ./...
server:
	go run ./cmd/main.go
.PHONY: createmigration  postgres createdb dropdb migrateup migratedown sqlc image test start server