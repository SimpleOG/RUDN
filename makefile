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
mgu1:
	migrate -path  db/edumigrations -database "postgresql://root:1234@localhost:5440/education?sslmode=disable" -verbose up
mgd1:
	migrate -path  db/edumigrations -database "postgresql://root:1234@localhost:5440/education?sslmode=disable" -verbose down
createdbtest:
	docker exec -it postgres createdb --username=root --owner=root test_education
dropdbtest:
	docker exec -it postgres dropdb test_education
mgu2:
	migrate -path  db/edumigrations -database "postgresql://root:1234@localhost:5440/test_education?sslmode=disable" -verbose up
mgd2:
	migrate -path  db/edumigrations -database "postgresql://root:1234@localhost:5440/test_education?sslmode=disable" -verbose down
sqlc:
	sqlc generate
start:
	docker start postgres
test:
	go test -v -cover ./db/tests
.PHONY: createmigration  createdb createdbtest  postgres createdb dropdb mgu1 mgd1 dbtest dropdbtest mgu2 mgd2 sqlc image test start