package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGrade(t *testing.T) {
	// Testing decimals
	out := getGrade(86.8577)
	assert.Equal(t, "You got a B.", out)
	out2 := getGrade(98.212)
	assert.Equal(t, "You got an A!", out2)
	out3 := getGrade(56.5)
	assert.Equal(t, "You got an E.", out3)
	out4 := getGrade(64.76452)
	assert.Equal(t, "You got a D.", out4)
	out5 := getGrade(76.4234)
	assert.Equal(t, "You got a C.", out5)
	out6 := getGrade(0.5)
	assert.Equal(t, "You got an E.", out6)

	// testing whole numbers
	for i := 90; i <= 100; i++ {
		res := getGrade(float64(i))
		assert.Equal(t, "You got an A!", res)
	}
	for i := 80; i <= 89; i++ {
		res := getGrade(float64(i))
		assert.Equal(t, "You got a B.", res)
	}
	for i := 70; i <= 79; i++ {
		res := getGrade(float64(i))
		assert.Equal(t, "You got a C.", res)
	}
	for i := 60; i <= 69; i++ {
		res := getGrade(float64(i))
		assert.Equal(t, "You got a D.", res)
	}
	for i := 0; i <= 59; i++ {
		res := getGrade(float64(i))
		assert.Equal(t, "You got an E.", res)
	}

}
