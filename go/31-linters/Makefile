.PHONY: help format vet lint lint-revive lint-staticcheck lint-golangci lint-fast lint-fix test build clean

help: ## Mostra esta mensagem de ajuda
	@echo "Comandos disponíveis:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

format: ## Formata código e organiza imports
	@echo "📝 Formatando código..."
	@if command -v goimports >/dev/null 2>&1; then \
		goimports -w .; \
		echo "✅ Código formatado!"; \
	else \
		echo "❌ goimports não encontrado. Instale com: go install golang.org/x/tools/cmd/goimports@latest"; \
		exit 1; \
	fi

vet: ## Executa go vet
	@echo "🔎 Executando go vet..."
	@go vet ./...
	@echo "✅ go vet passou!"

lint-revive: ## Executa revive
	@echo "✨ Executando revive..."
	@if command -v revive >/dev/null 2>&1; then \
		revive ./...; \
		echo "✅ revive passou!"; \
	else \
		echo "❌ revive não encontrado. Instale com: go install github.com/mgechev/revive@latest"; \
		exit 1; \
	fi

lint-staticcheck: ## Executa staticcheck
	@echo "🔍 Executando staticcheck..."
	@if command -v staticcheck >/dev/null 2>&1; then \
		staticcheck ./...; \
		echo "✅ staticcheck passou!"; \
	else \
		echo "❌ staticcheck não encontrado. Instale com: go install honnef.co/go/tools/cmd/staticcheck@latest"; \
		exit 1; \
	fi

lint-golangci: ## Executa golangci-lint (completo)
	@echo "🚀 Executando golangci-lint..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
		echo "✅ golangci-lint passou!"; \
	else \
		echo "❌ golangci-lint não encontrado. Instale com: brew install golangci-lint"; \
		exit 1; \
	fi

lint-fast: ## Executa golangci-lint (modo rápido)
	@echo "⚡ Executando golangci-lint (modo rápido)..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run --fast; \
		echo "✅ golangci-lint (fast) passou!"; \
	else \
		echo "❌ golangci-lint não encontrado. Instale com: brew install golangci-lint"; \
		exit 1; \
	fi

lint-fix: ## Executa golangci-lint e tenta corrigir automaticamente
	@echo "🔧 Executando golangci-lint com correções automáticas..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run --fix; \
		echo "✅ Correções aplicadas!"; \
	else \
		echo "❌ golangci-lint não encontrado. Instale com: brew install golangci-lint"; \
		exit 1; \
	fi

lint: lint-revive lint-staticcheck ## Executa linters individuais (revive + staticcheck)

test: ## Executa testes
	@echo "🧪 Executando testes..."
	@go test -v ./...
	@echo "✅ Testes concluídos!"

build: ## Compila o projeto
	@echo "🔨 Compilando..."
	@go build -o app .
	@echo "✅ Compilação concluída!"

clean: ## Limpa arquivos gerados
	@echo "🧹 Limpando..."
	@rm -f app
	@go clean
	@echo "✅ Limpeza concluída!"

all: format vet lint-fast test ## Executa tudo (formato, vet, lint rápido, testes)

ci: format vet lint-golangci test ## Executa tudo para CI/CD (formato, vet, lint completo, testes)



