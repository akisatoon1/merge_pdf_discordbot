package sender

import (
	"bytes"
	"errors"
	"merge_pdf/usecase"

	"github.com/bwmarrin/discordgo"
)

type sender struct {
	session   *discordgo.Session
	channelID string
}

func NewSender(s *discordgo.Session, channelID string) usecase.Sender {
	return &sender{
		session:   s,
		channelID: channelID,
	}
}

func (s *sender) Send(pdf *usecase.PDF) error {
	if s.session == nil {
		return errors.New("discord session is nil")
	}

	if pdf == nil {
		return nil // No PDF to send
	}

	cont := pdf.Content()
	file := &discordgo.File{
		Name:        "merged.pdf",
		ContentType: "application/pdf",
		Reader:      bytes.NewReader(cont),
	}

	_, err := s.session.ChannelMessageSendComplex(s.channelID, &discordgo.MessageSend{
		Files: []*discordgo.File{file},
	})

	return err
}
