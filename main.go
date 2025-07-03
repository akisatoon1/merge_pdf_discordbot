package main

import (
	"fmt"
	"merge_pdf/internal/config"
)

func main() {
	c := config.NewConfig()
	if err := c.Init(); err != nil {
		panic(fmt.Sprintf("Error initializing config: %v", err))
	}

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
