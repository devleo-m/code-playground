.PHONY: format vet test quality clean help

# Variáveis
GO := go
GOIMPORTS := goimports
PACKAGES := ./...

# Ajuda
help:
	@echo "Comandos disponíveis:"
	@echo "  make format    - Formatar código com goimports"
	@echo "  make vet       - Verificar qualidade com go vet"
	@echo "  make test      - Executar testes"
	@echo "  make quality   - Executar todas as verificações"
	@echo "  make clean     - Limpar arquivos gerados"
	@echo "  make help      - Mostrar esta ajuda"

# Formatar código
format:
	@echo "📝 Formatando código com goimports..."
	@$(GOIMPORTS) -w $(PACKAGES)
	@echo "✅ Formatação concluída!"

# Verificar qualidade
vet:
	@echo "🔍 Verificando qualidade com go vet..."
	@$(GO) vet $(PACKAGES)
	@echo "✅ Verificação concluída!"

# Executar testes
test:
	@echo "🧪 Executando testes..."
	@$(GO) test -v $(PACKAGES)
	@echo "✅ Testes concluídos!"

# Verificar formatação (sem modificar)
check-format:
	@echo "🔍 Verificando formatação..."
	@if [ "$$($(GOIMPORTS) -l $(PACKAGES) | wc -l)" -gt 0 ]; then \
		echo "❌ Código não está formatado corretamente:"; \
		$(GOIMPORTS) -l $(PACKAGES); \
		exit 1; \
	fi
	@echo "✅ Formatação OK!"

# Tudo junto: qualidade completa
quality: format vet test
	@echo ""
	@echo "✅ Qualidade verificada com sucesso!"
	@echo "📊 Resumo:"
	@echo "   - Código formatado ✓"
	@echo "   - go vet passou ✓"
	@echo "   - Testes passaram ✓"

# Limpar
clean:
	@echo "🧹 Limpando arquivos gerados..."
	@$(GO) clean
	@echo "✅ Limpeza concluída!"

# Instalar goimports se não estiver instalado
install-tools:
	@echo "📦 Instalando ferramentas..."
	@$(GO) install golang.org/x/tools/cmd/goimports@latest
	@echo "✅ Ferramentas instaladas!"

