package merger

import (
	"bytes"
	"errors"
	"io"
	"merge_pdf/internal/usecase"

	pdfapi "github.com/pdfcpu/pdfcpu/pkg/api"
)

type merger struct{}

func NewMerger() usecase.Merger {
	return &merger{}
}

func (m *merger) Merge(pdfs []*usecase.PDF) (*usecase.PDF, error) {
	if len(pdfs) < 2 {
		return nil, errors.New("at least two PDFs are required for merging")
	}

	var pdfReaders []io.ReadSeeker
	for _, pdf := range pdfs {
		if pdf == nil {
			continue
		}
		pdfReaders = append(pdfReaders, bytes.NewReader(pdf.Content()))
	}

	var buf bytes.Buffer
	err := pdfapi.MergeRaw(pdfReaders, &buf, false, nil)
	if err != nil {
		return nil, err
	}

	return usecase.NewPDF(buf.Bytes()), nil
}
