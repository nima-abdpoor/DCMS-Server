postgresql:
	docker run --name postgre14.5 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14.5-alpine

createDB:
	docker exec -it postgre14.5 createdb --username=root --owner=root DCMS-Test

createDBProduction:
	docker exec -it postgre14.5 createdb --username=root --owner=root DCMS

dropDB:
	docker exec -it postgre14.5 dropdb DCMS-Test

dropDBProduction:
	docker exec -it postgre14.5 dropdb DCMS

migrateUp:
	./migrate --path /home/nima/GolandProjects/DCMS-Server/db/postgresql/migration --database "postgresql://root:secret@localhost:5432/DCMS-Test?sslmode=disable" --verbose up

migrateDown:
	./migrate --path /home/nima/GolandProjects/DCMS-Server/db/postgresql/migration --database "postgresql://root:secret@localhost:5432/DCMS-Test?sslmode=disable" --verbose down

migrateUpW:
	migrate -database "postgresql://root:secret@localhost:5432/DCMS-Test?sslmode=disable" -path ./db/postgresql/migration up

migrateDownW:
	migrate -database "postgresql://root:secret@localhost:5432/DCMS-Test?sslmode=disable" -path ./db/postgresql/migration down

migrateUpProduction:
	./migrate --path /home/nima/GolandProjects/DCMS-Server/db/postgresql/migration --database "postgresql://root:secret@localhost:5432/DCMS?sslmode=disable" --verbose up

migrateDownProduction:
	./migrate --path /home/nima/GolandProjects/DCMS-Server/db/postgresql/migration --database "postgresql://root:secret@localhost:5432/DCMS?sslmode=disable" --verbose down

migrateUpWProduction:
	migrate -database "postgresql://root:secret@localhost:5432/DCMS?sslmode=disable" -path ./db/postgresql/migration up

migrateDownWProduction:
	migrate -database "postgresql://root:secret@localhost:5432/DCMS?sslmode=disable" -path ./db/postgresql/migration down

sqlc:
	sqlc generate

sqlcW:
	docker run --rm -v "C:\Users\Nima\GolandProjects\DCMS:/src" -w /src kjconroy/sqlc generate

test:
	go test -v ./...

server:
	go run Main.go

resetServer:
	make dropDB
	make dropDBProduction
	make createDB
	make createDBProduction
	make migrateUp
	make migrateUpW
	make migrateUpProduction
	make migrateUpWProduction

startServerW:
	make postgresql
	make createDB
	make createDBProduction
	make migrateUpW
	make migrateUpWProduction

.PHONY: postgresql createDB dropDB migrateUp migrateDownW migrateDownW migrateUp sqlc test server resetServer startServer