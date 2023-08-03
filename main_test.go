package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	var (
		a, b, c       int    = 5, 2, 0
		ans           int    = 2
		error_message string = "error: division by zero"
	)

	// test case1: 5 / 2
	result, err := divideNum(a, b)
	if err != nil {
		t.Errorf("error should not be nil when division by zero")
	}
	if result != ans {
		t.Errorf("expected %d, got %d", ans, result)
	}

	// test case2: 5 / 0
	_, err2 := divideNum(a, c)
	if err2 == nil {
		t.Errorf("error should be nil when division by zero")
	}
	if err2.Error() != error_message {
		t.Errorf("expected: %s, got %s", error_message, err2.Error())
	}

}

func TestAssert(t *testing.T) {
	assert := assert.New(t)

	var (
		a, b, c       int    = 5, 3, 0
		ans           int    = 2
		error_message string = "error: division by zero"
	)

	// test case1: 5 / 2
	result, err := divideNum(a, b)
	assert.Nil(err, "error should be nil when division by zero")
	assert.Equal(ans, result, "expected result does not match the actual result")

	// test case2: 5 / 0
	_, err2 := divideNum(a, c)
	assert.NotNil(err2, "error should not be nil when division by zero")
	assert.Equal(error_message, err2.Error(), "expected error message does not match the actual error message")
}
