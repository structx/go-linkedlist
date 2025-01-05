package linkedlist_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/structx/go-linkedlist"
)

type DoubleLinkedListSuite struct {
	suite.Suite
	dll *linkedlist.DoubleLinkedList[int, []byte]
}

func (suite *DoubleLinkedListSuite) SetupSuite() {
	suite.dll = &linkedlist.DoubleLinkedList[int, []byte]{}
}

func (suite *DoubleLinkedListSuite) TestInsert() {
	suite.dll.Insert(0, []byte("rick"))
	suite.dll.Insert(1, []byte("morty"))
	suite.dll.Insert(2, []byte("summer"))
	suite.dll.Insert(3, []byte("beth"))
	suite.dll.Insert(4, []byte("jerry"))
}

func (suite *DoubleLinkedListSuite) TestSearch() {

	assert := suite.Assert()

	suite.dll.Insert(1, []byte("helloworld"))
	value, err := suite.dll.Search(1)
	assert.NoError(err)

	assert.Equal("helloworld", string(value))

	_, err = suite.dll.Search(99)
	assert.Equal(linkedlist.ErrNotFound, err)
}

func (suite *DoubleLinkedListSuite) TestFlush() {

	assert := suite.Assert()

	suite.dll.Insert(1, []byte("helloworld"))
	assert.GreaterOrEqual(uintptr(0), dll.Size())

	suite.dll.Flush()
	assert.Equal(uintptr(0), suite.dll.Size())
}

func TestDoubleLinkedListSui(t *testing.T) {
	suite.Run(t, new(DoubleLinkedListSuite))
}
