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
func TestMassPart1(t *testing.T) {
	mass := calculateMass(12)
	assert.Equal(t, 2, mass)

	mass = calculateMass(14)
	assert.Equal(t, 2, mass)

	mass = calculateMass(1969)
	assert.Equal(t, 654, mass)

	mass = calculateMass(100756)
	assert.Equal(t, 33583, mass)
}


/*
A module of mass 14 requires 2 fuel. This fuel requires no further fuel (2 divided by 3 and rounded down is 0, which would call for a negative fuel), so the total fuel required is still just 2.
At first, a module of mass 1969 requires 654 fuel. Then, this fuel requires 216 more fuel (654 / 3 - 2). 216 then requires 70 more fuel, which requires 21 fuel, which requires 5 fuel, which requires no further fuel. So, the total fuel required for a module of mass 1969 is 654 + 216 + 70 + 21 + 5 = 966.
The fuel required by a module of mass 100756 and its fuel is: 33583 + 11192 + 3728 + 1240 + 411 + 135 + 43 + 12 + 2 = 50346.
 */
func TestMassPart2(t *testing.T) {
	mass := calculateMassWithFuel(12)
	assert.Equal(t, 2, mass)

	mass = calculateMassWithFuel(14)
	assert.Equal(t, 2, mass)

	mass = calculateMassWithFuel(1969)
	assert.Equal(t, 966, mass)

	mass = calculateMassWithFuel(100756)
	assert.Equal(t, 50346, mass)
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

func TestGeneratePart2Sum(t *testing.T) {
	data := readFile("data.txt")
	assert.Equal(t, 100, len(data))

	sum := 0
	for _, i := range data {
		sum += calculateMassWithFuel(i)
	}
	fmt.Printf("SUM: %d\n", sum)
}