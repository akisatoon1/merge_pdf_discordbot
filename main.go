package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	TOKEN     string
	ChannelID string
)

func main() {
	svr, err := setupServer(TOKEN, ChannelID)
	if err != nil {
		panic(fmt.Sprintf("Error setting up server: %v", err))
	}

	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	if err := svr.Serve(); err != nil {
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

	TOKEN = os.Getenv("DISCORD_BOT_TOKEN")
	if TOKEN == "" {
		panic("DISCORD_BOT_TOKEN is not set in .env file")
	}

	ChannelID = os.Getenv("DISCORD_CHANNEL_ID")
	if ChannelID == "" {
		panic("DISCORD_CHANNEL_ID is not set in .env file")
	}
}
