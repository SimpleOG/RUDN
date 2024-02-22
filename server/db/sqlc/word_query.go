package db

type WordQuery interface {
	ReadItAll() error
	FillTeacherHours(name string) error
	FillWord(name string, fields []string) (string, string, error)
	FillTables() error
	TakeInfo(fields []string, name string) error
}

var _ WordQuery = (*Queries)(nil)
