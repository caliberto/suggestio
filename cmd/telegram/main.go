package main

import (
	"log"
	"suggestio/src/config"
	"suggestio/src/controllers/telegram"
)

func main() {
	// Carica config (token, porta, ecc.)
	cfg := config.Load()

	// Avvia server Telegram bot
	if err := telegram.StartServer(cfg); err != nil {
		log.Fatalf("errore avvio bot: %v", err)
	}
}
