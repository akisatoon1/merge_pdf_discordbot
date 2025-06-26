package usecase

type Merger interface {
	Merge() (*PDF, error)
}
