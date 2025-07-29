# orbitallog

<img width="256" alt="Image" src="https://github.com/user-attachments/assets/6b67e78e-1839-47bf-ad73-4494f604c133" />

Uma biblioteca Go **simples e segura** para escrita de logs em arquivo, com suporte a:

- Criação automática de arquivo por dia (`prefixo_DD-MM-YYYY.log`)
- Limpeza automática de arquivos antigos com base na idade máxima configurada
- Operação **thread-safe** com `sync.Mutex`
- Sem goroutines em background — ideal para scripts e aplicações de execução única
- Uso simples, com API pequena e direta

## Características

- ✅ **Arquivo diário**: Um log por dia, com nome padronizado
- ✅ **Limpeza automática**: Remove logs antigos no início da execução
- ✅ **Thread-safe**: Seguro para uso com múltiplas goroutines
- ✅ **API mínima**: Apenas o essencial para escrever e fechar
- ✅ **Sem dependências externas**
- ✅ **Ideal para scripts e CLIs**

## Instalação

```bash
go get github.com/v1ctorw1ll1an/orbitallog@v.0.0.1
```

## Uso Básico

```go
package main

import (
	"time"

	"github.com/v1ctorw1ll1an/orbitallog"
)

func main() {
	// Cria o logger (mantém logs por 7 dias)
	logger, err := orbitallog.New("logs", "app", 7*24*time.Hour)
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	logger.Printf("Servidor iniciado em %v", time.Now())
	logger.Printf("Evento: %s", "Conexão recebida")
}
```

## API

```go
// Cria novo logger
func New(logDir, prefix string, maxAge time.Duration) (*Logger, error)

// Escreve no log (thread-safe)
func (l *Logger) Printf(format string, v ...any)

// Fecha o arquivo de log
func (l *Logger) Close() error
```

## Estrutura dos Arquivos

Arquivos seguem o padrão:

```
{logDir}/{prefix}_{DD-MM-YYYY}.log
```

Exemplo:

```
logs/
├── app_29-07-2025.log
├── app_28-07-2025.log
└── app_27-07-2025.log
```

## Formato das Mensagens

```
2025/07/29 15:42:01 Servidor iniciado em 2025-07-29 15:42:01
2025/07/29 15:42:05 Evento: Conexão recebida
```

## Testes

O projeto já inclui testes unitários e exemplos:

```bash
go test ./... -v
```

- **logger_test.go** → Testa criação, escrita, limpeza e fechamento
- **example_test.go** → Mostra uso básico, gerando documentação no `pkg.go.dev`

## Boas Práticas

1. Crie o logger uma vez por execução e use `defer Close()`
2. Configure `maxAge` para evitar acúmulo excessivo de logs
3. Sempre use nomes de prefixo claros para diferenciar módulos/serviços
4. Se rodar em produção, aponte `logDir` para um local seguro e com permissão de escrita

## Licença

MIT License — veja o arquivo LICENSE para mais detalhes.
