package Day1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
For a mass of 12, divide by 3 and round down to get 4, then subtract 2 to get 2.
For a mass of 14, dividing by 3 and rounding down still yields 4, so the fuel required is also 2.
For a mass of 1969, the fuel required is 654.
For a mass of 100756, the fuel required is 33583.
 */
func TestMass(t *testing.T) {
	mass := calculateMass(12)
	assert.Equal(t, 2, mass)

	mass = calculateMass(14)
	assert.Equal(t, 2, mass)

	mass = calculateMass(1969)
	assert.Equal(t, 654, mass)

	mass = calculateMass(100756)
	assert.Equal(t, 33583, mass)
}

func TestLoadData(t *testing.T) {
	data := readFile("data.txt")
	assert.Equal(t, 100, len(data))
}

func TestGeneratePart1Sum(t *testing.T) {
	data := readFile("data.txt")
	assert.Equal(t, 100, len(data))

	sum := 0
	for _, i := range data {
		sum += calculateMass(i)
	}
	fmt.Printf("SUM: %d\n", sum)
}

