DB_URL=postgresql://root:123456@localhost:5432/simple_bank?sslmode=disable

tidy:
	go mod tidy

server:
	go run main.go

templ:
	templ generate

db:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123456 -d postgres

createdb:
	docker exec -it postgres createdb --username=root --owner=root ikarus_travel

dropdb:
	docker exec -it postgres dropdb ikarus_travel

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

.PHONY: tidy server templ db createdb dropdb migrateup migratedown new_migration migrateup1 migratedown1
