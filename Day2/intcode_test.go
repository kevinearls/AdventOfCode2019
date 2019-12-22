package Day2

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimple(t *testing.T) {
	wtf := []int{1,9,10,3,2,3,11,0,99,30,40,50}
	result := process(wtf)
	fmt.Printf("%v\n", result)

	expectedResult := []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}
	assert.Equal(t, wtf, expectedResult)
}

func TestExample1(t *testing.T) {
	example := []int{1,0,0,0,99}
	expectedResult := []int{2,0,0,0,99}
	result := process(example)
	assert.Equal(t, expectedResult, result)
}

func TestExample2(t *testing.T) {
	example := []int{2,3,0,3,99}
	expectedResult := []int{2,3,0,6,99}
	result := process(example)
	assert.Equal(t, expectedResult, result)
}

func TestExample3(t *testing.T) {
	example := []int{2,4,4,5,99,0}
	expectedResult := []int{2,4,4,5,99,9801}
	result := process(example)
	assert.Equal(t, expectedResult, result)
}

func TestExample4(t *testing.T) {
	example := []int{1,1,1,4,99,5,6,0,99}
	expectedResult := []int{30,1,1,4,2,5,6,0,99}
	result := process(example)
	assert.Equal(t, expectedResult, result)
}

func TestPart1(t *testing.T) {
	example := []int{1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,6,1,19,1,5,19,23,2,6,23,27,1,27,5,31,2,9,31,35,1,5,35,39,2,6,39,
		43,2,6,43,47,1,5,47,51,2,9,51,55,1,5,55,59,1,10,59,63,1,63,6,67,1,9,67,71,1,71,6,75,1,75,13,79,2,79,13,83,2,9,
		83,87,1,87,5,91,1,9,91,95,2,10,95,99,1,5,99,103,1,103,9,107,1,13,107,111,2,111,10,115,1,115,5,119,2,13,119,123,
		1,9,123,127,1,5,127,131,2,131,6,135,1,135,5,139,1,139,6,143,1,143,6,147,1,2,147,151,1,151,5,0,99,2,14,0,0}

	//To do this, before running the program, replace position 1 with the value 12 and replace position 2 with the value 2.
	example[1] = 12
	example[2] = 2

	// 337024 is too low
	fmt.Printf("Length %d\n", len(example))
	result := process(example)
	fmt.Printf("RESULT[0]%d\n", result[0])
}