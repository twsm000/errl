package errl

type errorListIterator struct {
	index int
	list  *ErrorList
}

func (ei *errorListIterator) HasNext() bool {
	return ei.index < len(ei.list.errs)
}

func (ei *errorListIterator) Next() error {
	err := ei.list.errs[ei.index]
	ei.index++
	return err
}

func (ei *errorListIterator) Reset() {
	ei.index = 0
}

type emptyIterator struct{}

func (ei *emptyIterator) HasNext() bool {
	return false
}

func (ei *emptyIterator) Next() error {
	return nil
}

func (ei *emptyIterator) Reset() {
}
