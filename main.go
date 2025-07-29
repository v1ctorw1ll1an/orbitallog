package main

import (
	"time"

	"github.com/v1ctorw1ll1an/orbitallog/orbitallog"
)

func main() {
	oneWeek := time.Second * 20 //7 * 24 * time.Hour
	log, err := orbitallog.New("logs", "app", oneWeek)
	if err != nil {
		panic(err)
	}
	defer log.Close()

	log.Printf("Servidor iniciado em 20 segundos depois %v", time.Now())
	log.Printf("Evento: %s", "Conexão recebida")
}
