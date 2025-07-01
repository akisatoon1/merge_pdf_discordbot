package usecase

import (
	"io"

	"github.com/bwmarrin/discordgo"
)

type Processor interface {
	MergeAndSend(s *discordgo.Session, m *discordgo.MessageCreate)
}

type processor struct {
	merger    Merger
	sender    Sender
	channelID string
}

func NewProcessor(merger Merger, sender Sender, channelID string) Processor {
	return &processor{
		merger:    merger,
		sender:    sender,
		channelID: channelID,
	}
}

func (w *processor) MergeAndSend(s *discordgo.Session, m *discordgo.MessageCreate) {
	if s == nil || m == nil {
		return
	}

	if m.Author.Bot {
		return
	}

	if m.ChannelID != w.channelID {
		return
	}

	pdfs, err := download(s, m)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "PDFの取得に失敗しました")
		return
	}

	merged, err := w.merger.Merge(pdfs)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "PDFの結合に失敗しました")
		return
	}

	err = w.sender.Send(merged)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "PDF送信に失敗しました")
	}
}

func download(s *discordgo.Session, m *discordgo.MessageCreate) ([]*PDF, error) {
	urls := getURLs(m)

	var pdfs []*PDF
	for _, url := range urls {
		resp, err := s.Client.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		pdfs = append(pdfs, NewPDF(data))
	}

	return pdfs, nil
}

func getURLs(m *discordgo.MessageCreate) []string {
	var urls []string
	for _, att := range m.Attachments {
		if att.ContentType == "application/pdf" || (len(att.Filename) > 4 && att.Filename[len(att.Filename)-4:] == ".pdf") {
			urls = append(urls, att.URL)
		}
	}
	return urls
}
