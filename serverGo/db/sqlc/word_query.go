package db

type WordQuery interface {
	ReadItAll() error
	FillTeacherHours(name string) error
	TakeInfo(fields *[]string, name, semType string) ([]map[string]string, error)
}

var _ WordQuery = (*Queries)(nil)
