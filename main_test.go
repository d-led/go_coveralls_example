package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnswers(t *testing.T) {
	assert.Equal(t, 42, toCover(true))
}

func TestQuestions(t *testing.T) {
	assert.PanicsWithValue(t, "oh no!", func() {
		toCover(false)
	})
}
