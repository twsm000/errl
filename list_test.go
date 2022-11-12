package errl

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var (
	Err1     = errors.New("error 1")
	Err2     = errors.New("error 2")
	Err3     = errors.New("error 3")
	ErrTests = []error{Err1, Err2, Err3}
)

func AddErrosToTheList(errList *ErrorList) {
	errList.TryAdd(Err1)
	errList.TryAdd(Err2)
	errList.TryAdd(Err3)
}

func TestErrorListTestSuite(t *testing.T) {
	suite.Run(t, new(ErrorListTestSuite))
}

type ErrorListTestSuite struct {
	suite.Suite
	errList *ErrorList
}

func (suite *ErrorListTestSuite) SetupTest() {
	suite.errList = New()
}

func (suite *ErrorListTestSuite) TestNewShouldReturnEmptyList() {
	assert.NotNil(suite.T(), suite.errList)
	assert.True(suite.T(), suite.errList.IsEmpty())
}

func (suite *ErrorListTestSuite) TestFirstShouldReturnNilWhenNoErrorAddedToTheList() {
	assert.Nil(suite.T(), suite.errList.First())
}

func (suite *ErrorListTestSuite) TestFirstShouldReturnTheFirstErrorRaised() {
	AddErrosToTheList(suite.errList)
	assert.ErrorIs(suite.T(), suite.errList.First(), Err1)
}

func (suite *ErrorListTestSuite) TestLastShouldReturnNilWhenNoErrorAddedToTheList() {
	assert.Nil(suite.T(), suite.errList.Last())
}

func (suite *ErrorListTestSuite) TestLastShouldReturnTheLastErrorRaised() {
	AddErrosToTheList(suite.errList)
	assert.ErrorIs(suite.T(), suite.errList.Last(), Err3)
}

func (suite *ErrorListTestSuite) TestIteratorShouldNeverBeNil() {
	assert.NotNil(suite.T(), suite.errList.Iterator())
	assert.EqualValues(suite.T(), New().Iterator(), suite.errList.Iterator())
}

func (suite *ErrorListTestSuite) TestIteratorShouldReturnErrorsWhenIsNotEmpty() {
	AddErrosToTheList(suite.errList)
	it := suite.errList.Iterator()
	assert.True(suite.T(), it.HasNext())
	assert.ErrorIs(suite.T(), it.Next(), Err1)
	assert.ErrorIs(suite.T(), it.Next(), Err2)
	assert.ErrorIs(suite.T(), it.Next(), Err3)
}

func (suite *ErrorListTestSuite) TestContainsShouldReturnFalseWhenErrorNotExistsInTheList() {
	AddErrosToTheList(suite.errList)
	err := errors.New("a new error")
	assert.False(suite.T(), suite.errList.Contains(err))
}

func (suite *ErrorListTestSuite) TestContainsShouldReturnTrueWhenErrorExistsInTheList() {
	AddErrosToTheList(suite.errList)
	assert.True(suite.T(), suite.errList.Contains(Err2))
}

func (suite *ErrorListTestSuite) TestParseErrorShouldReturnNilWhenListIsEmpty() {
	assert.Nil(suite.T(), suite.errList.ParseErrors())
}

func (suite *ErrorListTestSuite) TestParseErrorShouldReturnErrorThatImplementsErrors() {
	AddErrosToTheList(suite.errList)
	err := suite.errList.ParseErrors()
	assert.NotNil(suite.T(), err)
	assert.NotPanics(suite.T(), func() {
		_, ok := err.(Errors)
		assert.True(suite.T(), ok)
	})
}
