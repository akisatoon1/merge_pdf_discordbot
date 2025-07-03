package main

import (
	"fmt"
	"merge_pdf/infra/merger"
	"merge_pdf/infra/sender"
	"merge_pdf/infra/server"
	"merge_pdf/usecase"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	TOKEN     string
	ChannelID string
)

func main() {
	dg, err := discordgo.New("Bot " + TOKEN)
	if err != nil {
		panic("Error creating Discord session: " + err.Error())
	}

	mgr := merger.NewMerger()
	sdr := sender.NewSender(dg, ChannelID)

	wtr := usecase.NewProcessor(mgr, sdr, ChannelID)

	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	svr := server.NewServer(dg, wtr)
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
