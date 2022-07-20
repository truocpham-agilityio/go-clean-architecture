CURRENT_DIR = $(shell pwd)
MIGRATE=docker run -v $(CURRENT_DIR)/db/migrations/mysql:/migrations --network host migrate/migrate -path=/migrations -database "mysql://test:root@tcp(localhost:3307)/go-clean-architecture"

migrate-up:
	$(MIGRATE) -verbose up

migrate-down:
	$(MIGRATE) -verbose down -all

.PHONY: migrate-up migrate-down
