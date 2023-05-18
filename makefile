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

.PHONY: setrpc
setrpc:
	protoc --go-grpc_out=require_unimplemented_servers=false:internal/app/user/api --go_out=internal/app/user/api internal/app/user/api/rpc/user.proto

.PHONY: migrateup
migrateup:
	go run cmd/migration/main.go -migration=up

.PHONY: migratedown
migratedown:
	go run cmd/migration/main.go -migration=down

.PHONY: serverprivatekey
serverprivatekey:
	openssl genrsa -out server.key 2048
	openssl ecparam -genkey -name secp384r1 -out server.key

.PHONY: serverpublickey
serverpublickey:
	openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650