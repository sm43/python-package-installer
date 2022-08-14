
.PHONY: vendor
vendor:
	@echo Generating vendor directory
	@go mod tidy -compat=1.17 && go mod vendor

.PHONY: run
run:
	@echo Starting the service ...
	@go run ./main.go