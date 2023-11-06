createmigration:
	migrate create -ext sql -dir db/migrations -seq init_education
image:
	docker pull postgres:12-alpine
postgres:
	docker run --name postgres -p 5440:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=1234 -d postgres:12-alpine
createdb:
	docker exec -it postgres createdb --username=root --owner=root education
dropdb:
	docker exec -it postgres dropdb education
migrateup1:
	migrate -path  db/edumigrations -database "postgresql://root:1234@localhost:5440/education?sslmode=disable" -verbose up
migratedown1:
	migrate -path  db/edumigrations -database "postgresql://root:1234@localhost:5440/education?sslmode=disable" -verbose down
createdbtest:
	docker exec -it postgres createdb --username=root --owner=root test_education
dropdbtest:
	docker exec -it postgres dropdb test_education
migrateup2:
	migrate -path  db/edumigrations -database "postgresql://root:1234@localhost:5440/test_education?sslmode=disable" -verbose up
migratedown2:
	migrate -path  db/edumigrations -database "postgresql://root:1234@localhost:5440/test_education?sslmode=disable" -verbose down
sqlc:
	sqlc generate
start:
	docker start postgres
test:
	go test -v -cover ./db/tests
.PHONY: createmigration1 createdbtest  postgres createdb dropdb migrateup1 migratedown1 dbtest dropdbtest migrateup2 migratedown2 sqlc image test start