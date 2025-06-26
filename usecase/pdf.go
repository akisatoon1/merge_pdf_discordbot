package usecase

type PDF struct {
	content []byte
}

func NewPDF(content []byte) *PDF {
	return &PDF{content: content}
}

func (p *PDF) Content() []byte {
	return p.content
}
