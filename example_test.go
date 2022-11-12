package errl

import (
	"errors"
	"fmt"
	"io"
)

var (
	ErrUnknow = errors.New("unknow error")
	ErrRead   = errors.New("ReadCloseErrorGenerator.Read: generated error")
	ErrClose  = errors.New("ReadCloseErrorGenerator.Close: generated error")
)

type ReadCloseErrorGenerator struct{}

func (r ReadCloseErrorGenerator) Read(p []byte) (n int, err error) {
	return 0, ErrRead
}

func (r ReadCloseErrorGenerator) Close() error {
	return ErrClose
}

func ReadAllAndClose(readerCloser io.ReadCloser, consume Consumer[[]byte]) (err error) {
	errList := New()
	defer func() {
		errList.TryAdd(readerCloser.Close())
		err = errList.ParseErrors()
	}()

	res, err := io.ReadAll(readerCloser)
	Process(errList, res, err, consume)
	return
}

func ExampleReadCloseErrorGenerator() {
	var readCloser io.ReadCloser = ReadCloseErrorGenerator{}

	dataConsumer := func(data []byte) {
		fmt.Printf("Bytes readed as string: %s", string(data))
	}

	errHandler := func(err error) {
		fmt.Println("(LAST) ERROR:", err)

		errs, ok := err.(Errors)
		if !ok {
			fmt.Println("This error not support Errors interface")
			return
		}

		fmt.Println("FIRST ERROR:", errs.First())

		fmt.Println("ERROR LIST (First to Last)")
		it := errs.Iterator()
		for it.HasNext() {
			fmt.Println(it.Next())
		}

		if errs.Contains(ErrRead) {
			fmt.Println("ErrRead FOUND")
		}
		if !errs.Contains(ErrUnknow) {
			fmt.Println("ErrUnknow NOT FOUND")
		}
	}

	if ProcessError(ReadAllAndClose(readCloser, dataConsumer), errHandler) {
		fmt.Println("Errors handled successfully")
	} else {
		fmt.Println("Application run without errors")
	}
}
