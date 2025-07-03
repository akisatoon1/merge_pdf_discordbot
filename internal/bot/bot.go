package bot

import (
	"fmt"
	"merge_pdf/internal/infra/bot"
	"merge_pdf/internal/infra/merger"
	"merge_pdf/internal/infra/sender"
	"merge_pdf/internal/usecase"

	"github.com/bwmarrin/discordgo"
)

func NewBot(token, channelID string) (usecase.Bot, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, fmt.Errorf("Error creating Discord session: %w", err)
	}

	mgr := merger.NewMerger()
	sdr := sender.NewSender(dg, channelID)

	proc := usecase.NewProcessor(mgr, sdr, channelID)

	b := bot.NewBot(dg)
	b.AddHandler(proc.MergeAndSend)
	b.AddHandler(proc.Response)
	return b, nil
}
