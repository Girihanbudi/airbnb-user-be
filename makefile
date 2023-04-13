.PHONY: injectapp
injectapp:
	cd ./internal/app && wire

.PHONY: runapp
runapp:
	go run ./cmd/app/main.go

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
	go run db/migration/main.go -migration=up

.PHONY: migratedown
migratedown:
	go run db/migration/main.go -migration=down