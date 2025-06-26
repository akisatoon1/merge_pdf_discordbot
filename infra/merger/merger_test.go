package merger

import (
	"merge_pdf/usecase"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

// テストファイルのディレクトリ
var testdir string

func init() {
	_, thisFile, _, _ := runtime.Caller(0)
	testdir = filepath.Dir(thisFile)
}

func TestMerge(t *testing.T) {
	testdata := inputTestdata()

	// マージをする
	merger := NewMerger(testdata)
	pdf, err := merger.Merge()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if pdf == nil {
		t.Fatal("expected a merged PDF, got nil")
	}

	// マージができているか確認するため
	if err := saveFile(pdf); err != nil {
		t.Fatalf("failed to save PDF file: %v", err)
	}
}

func testdataPath(filename string) string {
	return filepath.Join(testdir, "testdata", filename)
}

func inputTestdata() []usecase.PDF {
	a, _ := os.ReadFile(testdataPath("a.pdf"))
	b, _ := os.ReadFile(testdataPath("b.pdf"))

	pdfA := usecase.NewPDF(a)
	pdfB := usecase.NewPDF(b)

	return []usecase.PDF{*pdfA, *pdfB}
}

func saveFile(pdf *usecase.PDF) error {
	path := filepath.Join(testdir, "test_merged.pdf")
	content := pdf.Content()

	if err := os.WriteFile(path, content, 0644); err != nil {
		return err
	}

	return nil
}
