include .env
export

LOCAL ?= true

ifeq ($(LOCAL), true)
	DB_CONN_HOST = 127.0.0.1
else
	DB_CONN_HOST = ${DB_HOST}
endif

migrate-up:
	@echo -e "\033[44m \033[97m Running database migrations up... \033[0m"
	migrate -path internal/infra/database/migrations -database postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_CONN_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable up
	@echo -e "\033[42m Migrations completed. \033[0m"

migrate-single:
	@echo -e "\033[44m \033[97m Running single database migration... \033[0m"
	migrate -path internal/infra/database/migrations -database postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_CONN_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable up 1
	@echo -e "\033[42m Single migration completed. \033[0m"

migrate-down:
	@echo -e "\033[44m \033[97m Running database migrations down... \033[0m"
	migrate -path internal/infra/database/migrations -database postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_CONN_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable down
	@echo -e "\033[42m Migrations rolled back. \033[0m"

migrate-prev:
	@echo -e "\033[44m \033[97m Running database migrations to previous version... \033[0m"
	migrate -path internal/infra/database/migrations -database postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_CONN_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable down 1
	@echo -e "\033[42m Migration rolled back. \033[0m"

migrate-create:
	@echo -e "\033[44m \033[97m Creating new migration file... \033[0m"
	migrate create -ext sql -dir internal/infra/database/migrations -seq $(name)
	@echo -e "\033[42m Migration file created.\033[0m"

migrate-fix:
	@read -p "Enter version to force: " version; \
	read -p "Are you sure you want to force the migration to version $$version? This can be dangerous. [y/N]: " confirm; \
	if [ "$$confirm" = "y" ] || [ "$$confirm" = "Y" ]; then \
		echo -e "\033[44m \033[97m Forcing migration to version $$version... \033[0m"; \
		migrate -path internal/infra/database/migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_CONN_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" force $$version; \
		echo -e "\033[42m Migration version forced.\033[0m"; \
	else \
		echo -e "\033[41m \033[97m Migration force aborted. \033[0m"; \
	fi