package errl

type Errors interface {
	error
	ErrorCollection
}

type ErrorCollection interface {
	First() error
	Last() error
	TryAdd(err error) bool
	Iterator() Iterator[error]
	Contains(err error) bool
}

type Iterator[T any] interface {
	HasNext() bool
	Next() T
	Reset()
}
