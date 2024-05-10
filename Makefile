include .env
export

LOCAL_BIN:=$(CURDIR)/bin
PATH:=$(LOCAL_BIN):$(PATH)

# HELP =================================================================================================================
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

run:
	sh build.sh
	sh output/bootstrap.sh
.PHONY: run


migrate-create:  ### create new migration
	migrate create -ext sql -dir migrations create_chat_app_table
.PHONY: migrate-create

migrate-up: ### migration up
	migrate -database  '$(PG_MIGRATE_URL)?sslmode=disable' -path migrations up
.PHONY: migrate-up


migrate-down: ### migration down
	migrate -path migrations -database '$(PG_MIGRATE_URL)?sslmode=disable' down
.PHONY: migrate-down

cert:
	openssl genrsa -out cert/id_rsa 4096
	openssl rsa -in cert/id_rsa -pubout -out cert/id_rsa.pub
.PHONY: cert

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	proto/*.proto
.PHONY: proto