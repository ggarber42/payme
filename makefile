.PHONY: run-app
run-app:
	@echo "Running app.."
	go run cmd/api/api.go

.PHONY: unit-test
unit-test:
	@echo "Running unit test"
	@go test ./tests/unit/... --tags=unit -v
