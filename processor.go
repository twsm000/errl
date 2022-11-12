package errl

type Consumer[T any] func(data T)

func Process[T any](erl *ErrorList, t T, err error, consume Consumer[T]) bool {
	if erl.TryAdd(err) {
		return false
	}
	consume(t)
	return true
}

func ProcessError(err error, consume Consumer[error]) bool {
	if err == nil {
		return false
	}
	consume(err)
	return true
}
