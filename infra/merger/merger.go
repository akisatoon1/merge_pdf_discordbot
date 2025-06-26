package merger

import (
	"bytes"
	"errors"
	"io"
	"merge_pdf/usecase"

	pdfapi "github.com/pdfcpu/pdfcpu/pkg/api"
)

type merger struct {
	pdfs []usecase.PDF
}

func NewMerger(pdfs []usecase.PDF) usecase.Merger {
	return &merger{pdfs: pdfs}
}

func (m *merger) Merge() (*usecase.PDF, error) {
	if len(m.pdfs) < 2 {
		return nil, errors.New("at least two PDFs are required for merging")
	}

	var pdfReaders []io.ReadSeeker
	for _, pdf := range m.pdfs {
		pdfReaders = append(pdfReaders, bytes.NewReader(pdf.Content()))
	}

	var buf bytes.Buffer
	err := pdfapi.MergeRaw(pdfReaders, &buf, false, nil)
	if err != nil {
		return nil, err
	}

	return usecase.NewPDF(buf.Bytes()), nil
}
