package lily_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/structx/lily"
)

type LinkedListSuite struct {
	suite.Suite
	ll *lily.LinkedList[int]
}

func (suite *LinkedListSuite) SetupSuite() {
	suite.ll = lily.NewLinkedListInt()
}

func (suite *LinkedListSuite) TestInsert() {
	ll.Insert(0, []byte("rick"))
	ll.Insert(1, []byte("morty"))
	ll.Insert(2, []byte("summer"))
	ll.Insert(3, []byte("beth"))
	ll.Insert(4, []byte("jerry"))
}

func (suite *LinkedListSuite) TestSearch() {

	assert := suite.Assert()

	suite.ll.Insert(1, []byte("helloworld"))

	value, err := suite.ll.Search(1)
	assert.NoError(err)

	assert.Equal("helloworld", string(value))

	_, err = suite.ll.Search(2)
	assert.Equal(lily.ErrNotFound, err)
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
