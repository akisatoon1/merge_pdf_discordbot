package sender

import (
	"merge_pdf/usecase"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	testdir    string
	TOKEN      string
	CHANNEL_ID string
)

func TestMain(m *testing.M) {
	setTestdir()
	loadEnv()
	os.Exit(m.Run())
}

func setTestdir() {
	_, thisFile, _, _ := runtime.Caller(0)
	testdir = filepath.Dir(thisFile)
}

func loadEnv() {
	envPath := filepath.Join(testdir, ".env")
	_ = godotenv.Load(envPath)

	TOKEN = os.Getenv("DISCORD_TOKEN")
	CHANNEL_ID = os.Getenv("DISCORD_CHANNEL_ID")
}

func testdataPath(filename string) string {
	return filepath.Join(testdir, "testdata", filename)
}

func newSession(token string) (*discordgo.Session, error) {
	return discordgo.New("Bot " + token)
}

func inputTestdata() *usecase.PDF {
	a, err := os.ReadFile(testdataPath("a.pdf"))
	if err != nil {
		return nil
	}
	return usecase.NewPDF(a)
}

func TestSender_Send(t *testing.T) {
	if TOKEN == "" || CHANNEL_ID == "" {
		t.Skip("DISCORD_TOKEN or DISCORD_CHANNEL_ID is not set")
	}

	s, err := newSession(TOKEN)
	if err != nil {
		t.Fatalf("Failed to create session: %v", err)
	}

	sdr := NewSender(s, CHANNEL_ID)
	pdf := inputTestdata()
	if pdf == nil {
		t.Fatal("Failed to load test PDF data")
	}

	err = sdr.Send(pdf)
	if err != nil {
		t.Errorf("Failed to send PDF: %v", err)
	}
}
