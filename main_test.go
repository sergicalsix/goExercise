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

	if pred != ans {
		t.Errorf("expected %d, got %d", ans, pred)
	}

	// devideNum(a, c)

}

func TestAssert(t *testing.T) {
	var (
		a, b int = 5, 2
		ans  int = 2
	)

	assert.Equal(t, devideNum(a, b), ans)
	// devideNum(a, c)

}
