package merger

import (
	"merge_pdf/internal/usecase"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

var testdir string

func TestMain(m *testing.M) {
	setTestdir()
	os.Exit(m.Run())
}

func setTestdir() {
	_, thisFile, _, _ := runtime.Caller(0)
	testdir = filepath.Dir(thisFile)
}

func testdataPath(filename string) string {
	return filepath.Join(testdir, "testdata", filename)
}

func inputTestdata(t *testing.T) []*usecase.PDF {
	a, err := os.ReadFile(testdataPath("a.pdf"))
	if err != nil {
		t.Fatalf("failed to read a.pdf: %v", err)
	}
	b, err := os.ReadFile(testdataPath("b.pdf"))
	if err != nil {
		t.Fatalf("failed to read b.pdf: %v", err)
	}
	pdfA := usecase.NewPDF(a)
	pdfB := usecase.NewPDF(b)
	return []*usecase.PDF{pdfA, pdfB}
}

func saveFile(t *testing.T, pdf *usecase.PDF) {
	path := filepath.Join(testdir, "test_merged.pdf")
	content := pdf.Content()
	if err := os.WriteFile(path, content, 0644); err != nil {
		t.Fatalf("failed to save PDF file: %v", err)
	}
}

func TestMerger_Merge(t *testing.T) {
	testdata := inputTestdata(t)
	merger := NewMerger()
	pdf, err := merger.Merge(testdata)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if pdf == nil {
		t.Fatal("expected a merged PDF, got nil")
	}
	saveFile(t, pdf)
}
