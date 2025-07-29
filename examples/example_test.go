package orbitallog

import (
	"time"

	"github.com/v1ctorw1ll1an/orbitallog"
)

func ExampleLogger() {
	// Criar logger na pasta "logs", prefixo "app", guardando logs por 7 dias
	logger, err := orbitallog.New("logs", "app", 7*24*time.Hour)
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	logger.Printf("Servidor iniciado em %v", time.Now())
	logger.Printf("Evento: %s", "Conexão recebida")

	// Output:
}
