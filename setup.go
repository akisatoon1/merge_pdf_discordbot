package main

import (
	"fmt"
	"merge_pdf/infra/merger"
	"merge_pdf/infra/sender"
	"merge_pdf/infra/server"
	"merge_pdf/usecase"

	"github.com/bwmarrin/discordgo"
)

func setupServer(token, channelID string) (usecase.Server, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, fmt.Errorf("Error creating Discord session: %w", err)
	}

	mgr := merger.NewMerger()
	sdr := sender.NewSender(dg, channelID)

	wtr := usecase.NewProcessor(mgr, sdr, channelID)

	return server.NewServer(dg, wtr), nil
}
