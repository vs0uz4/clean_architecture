# Description: Makefile for clean architecture

TOOLS = migrate protoc go

DB_USER=root
DB_PASSWORD=root
DB_HOST=localhost
DB_PORT=3306
DB_NAME=orders
DB_URL=mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)

MIGRATE_CMD=migrate -path=sql/migrations -database "$(DB_URL)"

PROTOFILE_DIR=internal/infra/grpc/protofiles
GRAPHQL_PKG=github.com/99designs/gqlgen

.DEFAULT_GOAL := help

.PHONY: check_tools migrate-up migrate-down migrate-drop gen-proto gen-graphql server test

help:  ## Exibe este menu de ajuda
	@echo "Opções disponíveis no Makefile:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

check_tools: ## Verifica se as ferramentas necessárias estão instaladas
	@for tool in $(TOOLS); do \
		if ! command -v $$tool &> /dev/null; then \
			echo "$$tool not found! Please install it."; \
			exit 1; \
		fi; \
	done

migration-up: check_tools ## Aplica as migrações na base de dados
	@echo "Running migration up"
	@$(MIGRATE_CMD) up

migration-down: check_tools ## Reverte as migrações aplicadas na base de dados
	@echo "Running migration down"
	@$(MIGRATE_CMD) down

migration-drop: check_tools ## Elimina as migrações aplicadas na base de dados
	@echo "Running migration drop"
	@$(MIGRATE_CMD) drop

gen-proto: check_tools ## Efetua a geração dos arquivos protobuffer
	@echo "Generating proto"
	@protoc --go_out=. --go-grpc_out=. $(PROTOFILE_DIR)/order.proto

gen-graphql: check_tools ## Efetua a geração dos arquivos graphql
	@echo "Generating graphql"
	@go run $(GRAPHQL_PKG) generate 

server: check_tools ## Inicializa o servidor da aplicação
	@echo "Running server"
	@cd cmd/ordersystem && go run main.go wire_gen.go

test: check_tools ## Executa a suite de testes
	@echo "Running test"
	@go test -v ./... -coverprofile=coverage.out
	@go tool cover -func=coverage.out
