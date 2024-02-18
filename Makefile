migration_create:
	@echo "Creating migration"
	migrate create -ext sql -dir internal/infrastructure/database/migrations/ -seq $(name)

migration_up:
	@echo "Running migration"
	migrate -path internal/infrastructure/database/migrations/ -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up