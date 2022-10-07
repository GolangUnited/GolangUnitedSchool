default:
	$(info *********** COMMAND NOT SELECTED ***********)

install-dep:
	@echo "Install dependencies..."
	go install github.com/swaggo/swag/cmd/swag@v1.7.8
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.49.0
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.1
	go install github.com/vektra/mockery/v2@latest
	@echo "Done"

lint:
	@echo "Linting..."
	go fmt ./...
	golangci-lint run --skip-files ".*_gen.go" ./...
	@echo "Done"

migrate:
	@echo "Migrate..."
	bash scripts/migrate.sh
	@echo "Done"