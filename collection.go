package errl

func newErrorCollection(collection ErrorCollection) *errorCollection {
	return &errorCollection{
		error:           collection.Last(),
		ErrorCollection: collection,
	}
}

type errorCollection struct {
	error
	ErrorCollection
}
