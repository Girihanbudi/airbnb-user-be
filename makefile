.PHONY: wireapp
wireapp:
	cd ./internal/app && wire

.PHONY: documentation
docs:
	swag init -g ./cmd/app/main.go -o ./docs

.PHONY: gqlinit
gqlinit:
	go get github.com/99designs/gqlgen@v0.17.25
	go run github.com/99designs/gqlgen init

.PHONY: gqlgenerate
gqlgenerate:
	go get github.com/99designs/gqlgen@v0.17.25
	go run github.com/99designs/gqlgen generate

.PHONY: gqlrun
gqlrun:
	go run cmd/gql/main.go

.PHONY: migrateup
migrateup:
	go run db/main.go -migration=up

.PHONY: migratedown
migratedown:
	go run db/main.go -migration=down