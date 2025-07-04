package usecase

import (
	"io"
	"sort"

	"github.com/bwmarrin/discordgo"
)

type Handler interface {
	MergeAndSend(s *discordgo.Session, m *discordgo.MessageCreate)
	Response(s *discordgo.Session, m *discordgo.MessageCreate)
}

type handler struct {
	merger    Merger
	sender    Sender
	channelID string
}

func NewHandler(merger Merger, sender Sender, channelID string) Handler {
	return &handler{
		merger:    merger,
		sender:    sender,
		channelID: channelID,
	}
}

func (w *handler) Response(s *discordgo.Session, m *discordgo.MessageCreate) {
	if s == nil || m == nil {
		return
	}

	if m.Author.Bot {
		return
	}

	if m.ChannelID != w.channelID {
		return
	}

	s.ChannelMessageSend(m.ChannelID, "メッセージを受信しました。")
}

func (w *handler) MergeAndSend(s *discordgo.Session, m *discordgo.MessageCreate) {
	if s == nil || m == nil {
		return
	}

	if m.Author.Bot {
		return
	}

	if m.ChannelID != w.channelID {
		return
	}

	sortByFilename(m.Attachments)

	pdfs, err := download(s, m.Attachments)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "PDFの取得に失敗しました")
		return
	}

	merged, err := w.merger.Merge(pdfs)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "PDFの結合に失敗しました")
		return
	}

	fn := getFilename(m)
	err = w.sender.Send(merged, fn)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "PDF送信に失敗しました")
	}
}

func getFilename(m *discordgo.MessageCreate) string {
	c := m.Content
	if c == "" {
		return "merged.pdf"
	}
	return c + ".pdf"
}

func download(s *discordgo.Session, atts []*discordgo.MessageAttachment) ([]*PDF, error) {
	urls := getURLs(atts)

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

func getURLs(atts []*discordgo.MessageAttachment) []string {
	var urls []string
	for _, att := range atts {
		if att == nil {
			continue
		}

		if att.ContentType == "application/pdf" || (len(att.Filename) > 4 && att.Filename[len(att.Filename)-4:] == ".pdf") {
			urls = append(urls, att.URL)
		}
	}
	return urls
}

// 添付ファイルをファイル名で昇順ソートして取得
func sortByFilename(atts []*discordgo.MessageAttachment) {
	sort.Slice(atts, func(i, j int) bool {
		if atts[i] == nil {
			return false
		}
		if atts[j] == nil {
			return true
		}
		return atts[i].Filename < atts[j].Filename
	})
}
