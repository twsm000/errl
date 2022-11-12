package errl

import (
	"bytes"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestProcessTestSuite(t *testing.T) {
	suite.Run(t, new(ProcessTestSuite))
}

type ProcessTestSuite struct {
	suite.Suite
	errList *ErrorList
}

func (suite *ProcessTestSuite) SetupTest() {
	suite.errList = New()
}

func (suite *ProcessTestSuite) TestProcessShouldReturnFalseWhenErrorIsReturned() {
	f, err := os.Open("gibrish.unknow")
	got := Process(suite.errList, f, err, nil)
	assert.False(suite.T(), got)
	assert.False(suite.T(), suite.errList.IsEmpty())
}

func (suite *ProcessTestSuite) TestProcessShouldReturnTrueWhenNoErrorIsReturned() {
	const expected string = "TestProcessShouldReturnTrueWhenNoErrorIsReturned"

	buf := bytes.Buffer{}

	size, err := buf.WriteString(expected)
	consumer := func(bytesReaded int) {
		assert.EqualValues(suite.T(), len(expected), bytesReaded)
	}

	got := Process(suite.errList, size, err, consumer)
	assert.True(suite.T(), got)
	assert.True(suite.T(), suite.errList.IsEmpty())
}

func TestProcessErrorTestSuite(t *testing.T) {
	suite.Run(t, new(ProcessErrorTestSuite))
}

type ProcessErrorTestSuite struct {
	suite.Suite
	ErrTest error
}

func (suite *ProcessErrorTestSuite) SetupTest() {
	suite.ErrTest = errors.New("this is an error")
}

func (suite *ProcessErrorTestSuite) TestProcessErrorShouldReturnFalseWhenNoErrorIsPassed() {
	assert.False(suite.T(), ProcessError(nil, nil))
}

func (suite *ProcessErrorTestSuite) TestProcessErrorShouldReturnTrueWhenErrorIsPassed() {
	errHandler := func(err error) {
		assert.True(suite.T(), errors.Is(err, suite.ErrTest))
	}

	assert.True(suite.T(), ProcessError(suite.ErrTest, errHandler))
}
