createmigration1:
	migrate create -ext sql -dir db/migrations -seq init_education
createmigration2:
	migrate create -ext sql -dir db/migrations -seq init_users
image:
	docker pull postgres:12-alpine
postgres:
	docker run --name postgres -p 5440:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=1234 -d postgres:12-alpine
createdb1:
	docker exec -it postgres createdb --username=root --owner=root education
dropdb1:
	docker exec -it postgres dropdb education
migrateup1:
	migrate -path  db/edumigrations -database "postgresql://root:1234@localhost:5440/education?sslmode=disable" -verbose up
migratedown1:
	migrate -path  db/edumigrations -database "postgresql://root:1234@localhost:5440/education?sslmode=disable" -verbose down
createdb2:
	docker exec -it postgres createdb --username=root --owner=root users
dropdb2:
	docker exec -it postgres dropdb users
migrateup2:
	migrate -path  db/usersMigrations -database "postgresql://root:1234@localhost:5440/users?sslmode=disable" -verbose up
migratedown2:
	migrate -path  db/usersMigrations -database "postgresql://root:1234@localhost:5440/users?sslmode=disable" -verbose down
sqlc:
	sqlc generate
start:
	docker start postgres
test:
	go test -v -cover ./...
.PHONY: createmigration1 createmigration2  postgres createdb1 dropdb1 migrateup1 migratedown1 createdb2 dropdb2 migrateup2 migratedown2 sqlc image test start