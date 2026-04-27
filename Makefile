run:
	bash bobby.sh
	templ generate
	npx tailwindcss -i ./public/input.css -o ./public/styles.css 
	go run cmd/main.go

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))


migrate-up:
	@go run cmd/migrate/main.go up

migrate-down: 
	@go run cmd/migrate/main.go down

migrate-force:
ifndef version
	@echo "Usage: make migrate-force version=<version>"
	@exit 1
endif
	@go run cmd/migrate/main.go force $(version)