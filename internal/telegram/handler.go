package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Update struct {
	UpdateID int `json:"update_id"`
	Message  struct {
		Text string `json:"text"`
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

// Handler principale
func handleUpdate(c *gin.Context, token string) {
	var update Update
	if err := c.ShouldBindJSON(&update); err != nil {
		log.Println("errore parsing update:", err)
		c.Status(http.StatusBadRequest)
		return
	}

	text := update.Message.Text
	chatID := update.Message.Chat.ID

	switch text {
	case "/start":
		sendMessage(token, chatID, "👋 Ciao! Benvenuto su *Suggestio*. Scrivi /recommend per ricevere un consiglio!")
	case "/help":
		sendMessage(token, chatID, "📖 Comandi disponibili:\n/start - Avvia il bot\n/recommend - Ricevi una raccomandazione\n/help - Mostra questo messaggio")
	default:
		sendMessage(token, chatID, "❓ Comando non riconosciuto. Usa /help per la lista completa.")
	}

	c.Status(http.StatusOK)
}

// Funzione di utilità per inviare messaggi
func sendMessage(token string, chatID int64, text string) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)

	body, _ := json.Marshal(map[string]interface{}{
		"chat_id":    chatID,
		"text":       text,
		"parse_mode": "Markdown",
	})

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Println("errore invio messaggio:", err)
		return
	}
	defer resp.Body.Close()
}
