package main

import (
	"fmt"
	"merge_pdf/internal/infra/bot"
	"merge_pdf/internal/infra/merger"
	"merge_pdf/internal/infra/sender"
	"merge_pdf/internal/usecase"

	"github.com/bwmarrin/discordgo"
)

func setupBot(token, channelID string) (usecase.Bot, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, fmt.Errorf("Error creating Discord session: %w", err)
	}

	mgr := merger.NewMerger()
	sdr := sender.NewSender(dg, channelID)

	wtr := usecase.NewProcessor(mgr, sdr, channelID)

	return bot.NewBot(dg, wtr), nil
}
