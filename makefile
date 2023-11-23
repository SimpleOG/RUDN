createmigration:
	migrate create -ext sql -dir ./server/db/Migrations -seq init_education
image:
	docker pull postgres:12-alpine
postgres:
	docker run --name postgres -p 5440:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=1234 -d postgres:12-alpine
createdb:
	docker exec -it postgres createdb --username=root --owner=root education
dropdb:
	docker exec -it postgres dropdb education
mgu1:
	migrate -path  ./server/db/edumigrations -database "postgresql://root:1234@localhost:5440/education?sslmode=disable" -verbose up
mgd1:
	migrate -path  ./server/db/edumigrations -database "postgresql://root:1234@localhost:5440/education?sslmode=disable" -verbose down
createdbtest:
	docker exec -it postgres createdb --username=root --owner=root test_education
dropdbtest:
	docker exec -it postgres dropdb test_education
mgu2:
	migrate -path  ./server/db/Migrations -database "postgresql://root:1234@localhost:5440/test_education?sslmode=disable" -verbose up
mgd2:
	migrate -path  ./server/db/Migrations -database "postgresql://root:1234@localhost:5440/test_education?sslmode=disable" -verbose down
sqlc:
	cd server &&	sqlc generate
start_post:
	docker start postgres
server:
	cd server && go run main.go
client:
	cd client && npm start

test:
	go test -v -cover ./db/tests
.PHONY: server client createmigration  createdb createdbtest  postgres createdb dropdb mgu1 mgd1 dbtest dropdbtest mgu2 mgd2 sqlc image test start_post