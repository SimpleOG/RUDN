package db

type WordQuery interface {
	ReadItAll() error
	FillTeacherHours(name string) error
	FillWord(name string) (string, string, error)
}

var _ WordQuery = (*Queries)(nil)
