package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGrade(t *testing.T) {
	// Testing decimals
	out19 := getGrade("6/10")
	assert.Equal(t, "Invalid Input", out19)
	out20 := getGrade("Other words that are not good")
	assert.Equal(t, "Invalid Input", out20)
	out10 := getGrade("Words are no good")
	assert.Equal(t, "Invalid Input", out10)
	out := getGrade("86.8577")
	assert.Equal(t, "You got a B.", out)
	out2 := getGrade("98.212")
	assert.Equal(t, "You got an A!", out2)
	out3 := getGrade("56.5")
	assert.Equal(t, "You got an E.", out3)
	out4 := getGrade("64.76452")
	assert.Equal(t, "You got a D.", out4)
	out5 := getGrade("76.4234")
	assert.Equal(t, "You got a C.", out5)
	out6 := getGrade("0.5")
	assert.Equal(t, "You got an E.", out6)
	out11 := getGrade("103")
	assert.Equal(t, "Invalid Input", out11)
	out12 := getGrade("-56")
	assert.Equal(t, "Invalid Input", out12)
	out13 := getGrade("92")
	assert.Equal(t, "You got an A!", out13)
	out14 := getGrade("89")
	assert.Equal(t, "You got a B.", out14)
	out15 := getGrade("71")
	assert.Equal(t, "You got a C.", out15)
	out16 := getGrade("64")
	assert.Equal(t, "You got a D.", out16)
	out17 := getGrade("55")
	assert.Equal(t, "You got an E.", out17)
	out18 := getGrade("23")
	assert.Equal(t, "You got an E.", out18)

}
