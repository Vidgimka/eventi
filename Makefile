include .env
export
PROJECT_ROOT := $(subst \,/,$(CURDIR))
export PROJECT_ROOT

env-up:
	@docker compose up -d eventi_postgres
env-down: 
	@docker compose down eventi_postgres

env-cleanup:
	@echo "Clear all environment volume files? [y/N]:"
	@cmd /v:on /c "set /p ans=& if /i \"!ans!\"==\"y\" ( \
		docker compose down eventi_postgres && \
		rmdir /s /q out\pgdata && \
		echo Environment files deleted) \
		else (echo Deletion cancelled)"

env-forwarder-up:
	@docker compose up -d port_forwarder

env-forwarder-down:
	@docker compose down port_forwarder

migrate-create:
	@if "$(seq)"=="" ( \
		echo Missing required parameter seq && \
		exit 1 \
	)
	@docker compose run --rm eventi_migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action:
	@if "$(action)"=="" ( \
	echo Missing required parameter action && \
	exit 1 \
	)
	@docker compose run --rm eventi_migrate \
	-path /migrations \
	-database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@eventi_postgres:5432/${POSTGRES_DB}?sslmode=disable" \
	"${action}"

app-run:
	@go run cmd/eventi/main.go