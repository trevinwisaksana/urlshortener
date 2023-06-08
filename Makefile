postgres:
	docker run --name urlshortener -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	docker exec -it urlshortener createdb --username=root --owner=root urlshortener

server:
	go run main.go

redirect:
	go run cmd/redirect/redirect.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/trevinwisaksana/trevin-urlshortener/db/sqlc Store

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: sqlc