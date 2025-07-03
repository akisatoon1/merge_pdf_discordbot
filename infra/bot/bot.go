package bot

import (
	"fmt"
	"merge_pdf/usecase"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

type bot struct {
	session   *discordgo.Session
	processor usecase.Processor
}

func NewBot(s *discordgo.Session, proc usecase.Processor) usecase.Bot {
	return &bot{
		session:   s,
		processor: proc,
	}
}

func (s *bot) Start() error {
	if s.session == nil {
		return fmt.Errorf("discord session is nil")
	}

	s.session.AddHandler(s.processor.MergeAndSend)

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
