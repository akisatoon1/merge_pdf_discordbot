package usecase

import "github.com/bwmarrin/discordgo"

type Bot interface {
	Start() error
	AddHandler(handler func(s *discordgo.Session, m *discordgo.MessageCreate))
}
