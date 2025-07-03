package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	DiscordBotToken  string
	DiscordChannelID string
}

func NewConfig() *config {
	return &config{}
}

func (c *config) Init() error {
	err := eMode()
	if err != nil {
		return fmt.Errorf("Error initializing config: %w", err)
	}

	c.load()
	if err := c.validate(); err != nil {
		return fmt.Errorf("config validation failed: %w", err)
	}
	return nil
}

func (c *config) load() {
	c.DiscordBotToken = os.Getenv("DISCORD_BOT_TOKEN")
	c.DiscordChannelID = os.Getenv("DISCORD_CHANNEL_ID")
}

func (c *config) validate() error {
	if c.DiscordBotToken == "" {
		return fmt.Errorf("Discord bot token is required")
	}
	if c.DiscordChannelID == "" {
		return fmt.Errorf("Discord channel ID is required")
	}
	return nil
}

func eMode() error {
	if len(os.Args) == 2 && os.Args[1] == "-e" {
		fmt.Println("Read .env")
		err := godotenv.Load(".env")
		if err != nil {
			return fmt.Errorf("Error loading .env file: %w", err)
		}
	}
	return nil
}
