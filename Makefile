DB_URL=postgresql://postgres:postgre2022@localhost:5432/simple_bank?sslmode=disable

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate

