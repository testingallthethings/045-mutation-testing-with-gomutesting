package mutation_test

import (
	"github.com/stretchr/testify/suite"
	"mutation"
	"testing"
)

type MyIntSuite struct {
	suite.Suite
}

func TestMyIntSuite(t *testing.T) {
	suite.Run(t, new(MyIntSuite))
}

func (s *MyIntSuite) TestNotGreaterThan() {
	var nine = mutation.NewMyInt(9)

	s.False(nine.IsGreaterThan(10))
}

func (s *MyIntSuite) TestEqualNotGreaterThan() {
	var nine = mutation.NewMyInt(9)

	s.False(nine.IsGreaterThan(9))
}

func (s *MyIntSuite) TestIsGreaterThan() {
	var nine = mutation.NewMyInt(9)

	s.True(nine.IsGreaterThan(8))
}

func (s *MyIntSuite) TestNotLessThan() {
	var nine = mutation.NewMyInt(9)

	s.False(nine.IsLessThan(8))
}

func (s *MyIntSuite) TestEqualNotLessThan() {
	var nine = mutation.NewMyInt(9)

	s.False(nine.IsLessThan(9))
}

func (s *MyIntSuite) TestIsLessThan() {
	var nine = mutation.NewMyInt(9)

	s.True(nine.IsLessThan(10))
}
