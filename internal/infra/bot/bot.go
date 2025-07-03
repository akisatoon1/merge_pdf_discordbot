package bot

import (
	"fmt"
	"merge_pdf/internal/usecase"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

type bot struct {
	session *discordgo.Session
}

func NewBot(s *discordgo.Session) usecase.Bot {
	return &bot{
		session: s,
	}
}

func (s *bot) Start() error {
	if s.session == nil {
		return fmt.Errorf("discord session is nil")
	}

	err := s.session.Open()
	if err != nil {
		return err
	}
	defer s.session.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	return nil
}

func (s *bot) AddHandler(handler func(s *discordgo.Session, m *discordgo.MessageCreate)) {
	if s.session == nil {
		fmt.Println("Discord session is nil, cannot add handler")
		return
	}
	s.session.AddHandler(handler)
}
