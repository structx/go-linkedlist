package lili_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/structx/lili"
)

type LinkedListSuite struct {
	suite.Suite
	ll *lili.LinkedList[int]
}

func (suite *LinkedListSuite) SetupSuite() {
	suite.ll = lili.NewLinkedListInt()
}

func (suite *LinkedListSuite) TestInsert() {
	suite.ll.Insert(0, []byte("rick"))
	suite.ll.Insert(1, []byte("morty"))
	suite.ll.Insert(2, []byte("summer"))
	suite.ll.Insert(3, []byte("beth"))
	suite.ll.Insert(4, []byte("jerry"))
}

func (suite *LinkedListSuite) TestSearch() {

	assert := suite.Assert()

	suite.ll.Insert(99, []byte("helloworld"))

	value, err := suite.ll.Search(99)
	assert.NoError(err)

	assert.Equal("helloworld", string(value.([]byte)))

	_, err = suite.ll.Search(101)
	assert.Equal(lili.ErrNotFound, err)
}

func (suite *LinkedListSuite) TestFlush() {

	assert := suite.Assert()

	suite.ll.Insert(1, []byte("helloworld"))

	suite.ll.Flush()
	assert.Equal(uintptr(0), suite.ll.Size())
}

func TestLinkedListSuite(t *testing.T) {
	suite.Run(t, new(LinkedListSuite))
}
