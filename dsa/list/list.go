package list

type List[T any] interface {
	Add(pos int, v T) error
	Delete(pos int) (T, error)
	Value(pos int) (T, error)
	Len() int
}

type errorConst string

func (e errorConst) Error() string {
	return string(e)
}

const (
	ErrInvalidPosition = errorConst("invalid position")
)
