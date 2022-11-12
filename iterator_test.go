package errl

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestErrorListIteratorTestSuite(t *testing.T) {
	suite.Run(t, new(ErrorListIteratorTestSuite))
}

type ErrorListIteratorTestSuite struct {
	suite.Suite
	errList *ErrorList
}

func (suite *ErrorListIteratorTestSuite) SetupTest() {
	suite.errList = New()
}

func (suite *ErrorListIteratorTestSuite) TestHasNextShouldReturnFalseWhenEmptyListIsPassed() {
	it := errorListIterator{list: suite.errList}
	assert.False(suite.T(), it.HasNext())
	for it.HasNext() {
		assert.FailNow(suite.T(), "expect empty iterator")
	}
}

func (suite *ErrorListIteratorTestSuite) TestNextShouldReturnErrosInOrderTheyWereAdded() {
	AddErrosToTheList(suite.errList)
	it := errorListIterator{list: suite.errList}
	var index int
	for it.HasNext() {
		assert.True(suite.T(), errors.Is(it.Next(), ErrTests[index]))
		index++
	}
	assert.EqualValues(suite.T(), 3, index)
}

func (suite *ErrorListIteratorTestSuite) TestResetShouldReturnToFirstError() {
	AddErrosToTheList(suite.errList)
	it := errorListIterator{list: suite.errList}
	var index int
	var firstErr error
	for it.HasNext() {
		if index == 0 {
			firstErr = it.Next()
		} else {
			assert.NotErrorIs(suite.T(), it.Next(), firstErr)
		}
		index++
	}
	it.Reset()
	assert.True(suite.T(), it.HasNext())
	assert.ErrorIs(suite.T(), it.Next(), firstErr)
}

func TestEmptyIteratorTestSuit(t *testing.T) {
	suite.Run(t, new(EmptyIteratorTestSuit))
}

type EmptyIteratorTestSuit struct {
	suite.Suite
	it emptyIterator
}

func (suite *EmptyIteratorTestSuit) SetupTest() {
	suite.it = emptyIterator{}
}

func (suite *EmptyIteratorTestSuit) TestHasNextShouldAlwaysReturnFalse() {
	assert.False(suite.T(), suite.it.HasNext())
}

func (suite *EmptyIteratorTestSuit) TestNextShouldAlwaysReturnNil() {
	assert.Nil(suite.T(), suite.it.Next())
}

func (suite *EmptyIteratorTestSuit) TestRestShouldDoNothingAndNotPanic() {
	assert.NotPanics(suite.T(), func() {
		suite.it.Reset()
	})
}
