package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	var (
		a, b int = 5, 2
		ans  int = 2
	)

	var pred int = devideNum(a, b)

	// test case1: 5 / 2
	if pred != ans {
		t.Errorf("expected %d, got %d", ans, pred)
	}

	// test case1: 5 / 2
	// devideNum(a, c)

}

func TestAssert(t *testing.T) {
	var (
		a, b int = 5, 2
		ans  int = 2
	)
	// test case1: 5 / 2
	assert.Equal(t, devideNum(a, b), ans)
	// test case2: 5 / 2
	// devideNum(a, c)

}
