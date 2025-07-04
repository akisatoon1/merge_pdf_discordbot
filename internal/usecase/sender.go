package usecase

type Sender interface {
	Send(pdf *PDF, filename string) error
}
