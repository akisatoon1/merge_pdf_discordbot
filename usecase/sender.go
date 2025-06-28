package usecase

type Sender interface {
	Send(*PDF) error
}
