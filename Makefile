DB_URL=postgresql://root:root@localhost:5432/aahar?sslmode=disable

postgres:
	docker run --name aahar-pg -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:15-alpine

createdb:
	docker exec -it aahar-pg createdb --username=root --owner=root aahar

dropdb:
	docker exec -it aahar-pg dropdb aahar

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate

sqlc-win:
	docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate

destroy:
	docker stop aahar-pg
	docker rm aahar-pg

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc destroy sqlc-win test