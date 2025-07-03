package main

import (
	"fmt"
	"merge_pdf/internal/config"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	c := config.NewConfig()

	bot, err := setupBot(c.DiscordBotToken, c.DiscordChannelID)
	if err != nil {
		panic(fmt.Sprintf("Error setting up server: %v", err))
	}

	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	if err := bot.Start(); err != nil {
		panic(err)
	}
	fmt.Println("Bot is shutting down.")
}

func init() {
	if len(os.Args) == 2 && os.Args[1] == "-e" {
		fmt.Println("Read .env")
		err := godotenv.Load(".env")
		if err != nil {
			panic("Error loading .env file")
		}
	}
}
