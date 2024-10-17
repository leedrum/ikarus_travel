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

.PHONY: tidy server templ db createdb dropdb
