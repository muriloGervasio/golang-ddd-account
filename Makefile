migration_create:
	@echo "Creating migration"
	migrate create -ext sql -dir internal/infrastructure/database/migrations/ -seq $(name)