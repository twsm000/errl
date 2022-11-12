package errl

import "errors"

func New() *ErrorList {
	return new(ErrorList)
}

var emptyIt = &emptyIterator{}

type ErrorList struct {
	errs []error
}

func (e *ErrorList) First() error {
	if e.IsEmpty() {
		return nil
	}
	return e.errs[0]
}

func (e *ErrorList) Last() error {
	if e.IsEmpty() {
		return nil
	}
	return e.errs[len(e.errs)-1]
}

func (e *ErrorList) TryAdd(err error) bool {
	if err != nil {
		e.errs = append(e.errs, err)
		return true
	}
	return false
}

func (e *ErrorList) Iterator() Iterator[error] {
	if e.IsEmpty() {
		return emptyIt
	}
	return &errorListIterator{list: e}
}

func (e *ErrorList) Contains(target error) bool {
	it := e.Iterator()
	for it.HasNext() {
		if errors.Is(it.Next(), target) {
			return true
		}
	}
	return false
}

func (e *ErrorList) IsEmpty() bool {
	return e.errs == nil
}

func (e *ErrorList) ParseErrors() Errors {
	if e.IsEmpty() {
		return nil
	}
	return newErrorCollection(e)
}
