postgres:
	docker run --name urlshortener -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	docker exec -it urlshortener createdb --username=root --owner=root urlshortener

server:
	go run main.go

mock:
	mockgen -source=repository/urlshortener/postgresql/postgresql.go \
  			-destination=repository/urlshortener/postgresql/mock/postgresql.go \
  			-package=mockurlshortener \
  			-self_package=github.com/trevinwisaksana/trevin-urlshortener

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: sqlc