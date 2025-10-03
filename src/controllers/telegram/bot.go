package telegram

import (
	"log"
	"suggestio/src/config"

	"github.com/gin-gonic/gin"
)

func StartServer(cfg *config.Config) error {
	router := gin.Default()

	// Webhook endpoint
	router.POST("/webhook", func(c *gin.Context) {
		handleUpdate(c, cfg.BotToken)
	})

	log.Printf("Bot server in ascolto su :%s", cfg.Port)
	return router.Run(":" + cfg.Port)
}
