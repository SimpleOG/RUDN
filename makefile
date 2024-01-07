createmigration:
	migrate create -ext sql -dir ./server/db/Migrations -seq init_education
image:
	docker pull postgres:12-alpine
postgres:
	docker run --name postgres  --network network -p 5441:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=1234 -d postgres:12-alpine
createdb:
	docker exec -it postgres createdb --username=root --owner=root education
dropdb:
	docker exec -it postgres dropdb education
mgu:
	migrate -path  ./server/db/Migrations -database "postgresql://root:1234@localhost:5441/education?sslmode=disable" -verbose up
mgd:
	migrate -path  ./server/db/Migrations -database "postgresql://root:1234@localhost:5441/education?sslmode=disable" -verbose down
sqlc:
	cd server &&	sqlc generate
server:
	cd server && go run main.go
client:
	cd client && npm start
runserv:
	cd server && docker run --name server --network network -p 8080:8080 -e DB_SOURCE="postgresql://root:1234@postgres:5455/education?sslmode=disable" server
dockbuild:
	cd server && docker build -t server:latest .
test:
	go test -v -cover ./db/tests
network:
	docker network create network
compose:
	cd server && docker compose up

.PHONY: start_serv compose dockbuild runserv server client createmigration  createdb   postgres createdb dropdb mgd dbtest  sqlc image test start_post