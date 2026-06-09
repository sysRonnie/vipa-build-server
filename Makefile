APP_ENV ?= dev

assets:
	bash bobby.sh
	templ generate
	npx tailwindcss -i ./public/input.css -o ./public/styles.css

run: assets
	APP_ENV=$(APP_ENV) go run cmd/main.go

run-dev:
	$(MAKE) run APP_ENV=dev

run-prod:
	$(MAKE) run APP_ENV=prod

build: assets
	go build -o bin/server cmd/main.go

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	APP_ENV=$(APP_ENV) go run cmd/migrate/main.go up

migrate-up-prod:
	$(MAKE) migrate-up APP_ENV=prod

migrate-down:
	APP_ENV=$(APP_ENV) go run cmd/migrate/main.go down

migrate-down-prod:
	$(MAKE) migrate-down APP_ENV=prod

migrate-force:
ifndef version
	@echo "Usage: make migrate-force version=<version>"
	@exit 1
endif
	APP_ENV=$(APP_ENV) go run cmd/migrate/main.go force $(version)