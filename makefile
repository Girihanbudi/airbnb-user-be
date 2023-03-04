.PHONY: wireapp
wireapp:
	cd ./internal/app && wire

.PHONY: documentation
docs:
	swag init -g ./cmd/app/main.go -o ./docs