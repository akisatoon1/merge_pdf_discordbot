package config

import (
	"fmt"
	"os"
)

type config struct {
	DiscordBotToken  string
	DiscordChannelID string
}

func NewConfig() *config {
	c := &config{}
	c.load()
	if err := c.validate(); err != nil {
		fmt.Println("Error loading configuration:", err)
		os.Exit(1)
	}
	return c
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
