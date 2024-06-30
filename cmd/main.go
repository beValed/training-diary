package main

import (
	"log"
	"os"
	"training-diary/configs/entities"
	"training-diary/internal/handlers"

	"github.com/joho/godotenv"
	"github.com/yanzay/tbot/v2"
)

var (
	app   entities.Application
	bot   *tbot.Server
	token string
)

func main() {
	bot = tbot.New(token)
	app.Client = bot.Client()

	bot.HandleMessage("/start", handlers.StartBot(&app))
	bot.HandleMessage("/inform", handlers.GetOptions(&app))
	bot.HandleCallback(handlers.CallbackHandler(&app))

	log.Fatal(bot.Start())
}

func init() {
	e := godotenv.Load()
	if e != nil {
		log.Println(e)
	}
	token = os.Getenv("TELEGRAM_TOKEN")
}
