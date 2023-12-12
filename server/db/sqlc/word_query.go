package db

type WordQuery interface {
	ReadItAll() error
	FillTeacherHours(name string) error
	TeacherHours(name string) (Teacher_InfoRow, error)
}

var _ WordQuery = (*Queries)(nil)
