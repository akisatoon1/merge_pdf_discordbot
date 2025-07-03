package usecase

type Merger interface {
	Merge([]*PDF) (*PDF, error)
}
