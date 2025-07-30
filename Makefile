# Nome do binário (se quiser buildar o projeto)
BINARY_NAME=app

# Diretórios do código
SRC=./...

# Comandos principais
.PHONY: all install-tools errcheck lint build run clean

# Executa todos os checks e builda
all: install-tools errcheck build

# Instala as ferramentas necessárias
install-tools:
	@echo "==> Instalando ferramentas..."
	@go install github.com/kisielk/errcheck@latest

# Checa se todos os erros estão sendo tratados
errcheck:
	@echo "==> Rodando errcheck..."
	@errcheck $(SRC)

# Lint (se quiser integrar no futuro com golangci-lint)
lint:
	@echo "==> Rodando lint..."
	@golangci-lint run $(SRC) || true

# Builda o projeto
build:
	@echo "==> Compilando..."
	@go build -o $(BINARY_NAME) .

# Roda o projeto
run: build
	@echo "==> Executando..."
	@./$(BINARY_NAME)

# Limpa binários gerados
clean:
	@echo "==> Limpando..."
	@rm -f $(BINARY_NAME)
