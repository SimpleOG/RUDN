createmigration:
	migrate create -ext sql -dir ./serverGo/db/Migrations -seq new_migration
image:
	docker pull postgres:12-alpine
postgres:
	docker run --name postgres  --network network -p 5441:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=1234 -d postgres:12-alpine
createdb:
	docker exec -it postgres createdb --username=root --owner=root education
dropdb:
	docker exec -it postgres dropdb education
mgu:
	migrate -path  ./serverGo/db/Migrations -database "postgresql://root:1234@localhost:5441/education?sslmode=disable" -verbose up
mgd:
	migrate  -path  ./serverGo/db/Migrations -database "postgresql://root:1234@localhost:5441/education?sslmode=disable" -verbose down
sqlc:
	cd serverGo &&	sqlc generate
server:
	cd serverGo && go run main.go
client:
	cd client && npm start
runserv:
	docker run --name server --network network -p 8080:8080  server
dockbuild:
	cd serverGo && docker build -t server:latest .
test:
	go test -v -cover ./db/tests
network:
	docker network create network
compose:
	cd docker && docker compose up --build
proto:
	cd serverPy  &&python -m  grpc_tools.protoc -I proto --python_out=python/pb --grpc_python_out=python/pb proto/generator.proto
#после генерации при попытке запуска будет возникать ошибка, исправить заменив импорт в файле from . import generator_pb2 as generator__pb2
goproto:
	cd serverGo &&  protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
        --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
        --grpc-gateway_out=pb	--grpc-gateway_opt=paths=source_relative		\
        proto/*.proto
.PHONY:  start_serv compose dockbuild runserv server client createmigration  createdb   postgres createdb dropdb mgd dbtest  sqlc image test start_post