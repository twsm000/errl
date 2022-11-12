package errl

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestErrorCollectionTestSuite(t *testing.T) {
	suite.Run(t, new(ErrorCollectionTestSuite))
}

type ErrorCollectionTestSuite struct {
	suite.Suite
	err *errorCollection
}

func (suite *ErrorCollectionTestSuite) SetupTest() {
	errList := New()
	AddErrosToTheList(errList)
	suite.err = newErrorCollection(errList)
}

func (suite *ErrorCollectionTestSuite) TestErrorShouldReturnTheLastErrorOfCollection() {
	assert.ErrorIs(suite.T(), suite.err.error, Err3)
	assert.ErrorIs(suite.T(), suite.err.error, suite.err.Last())
}

func (suite *ErrorCollectionTestSuite) TestErrorShouldImplementsErrorsInterface() {
	assert.NotPanics(suite.T(), func() {
		var errs Errors = suite.err
		assert.True(suite.T(), errs.Contains(Err3))
		assert.True(suite.T(), errs.Contains(Err1))
		assert.True(suite.T(), errs.Contains(Err2))
	})
}
