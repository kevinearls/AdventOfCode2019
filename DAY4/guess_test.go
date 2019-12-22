package DAY4

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestGuessPassword(t *testing.T) {
	min := 145852
	max := 616942
	count := 0

	for i := min; i <= max; i++ {
		s := strconv.Itoa(i)
		hasMatch := hasDigitMatch(s)
		increasing := digitsAreIncreasing(s)

		if hasMatch && increasing {
			count++
		}
	}

	fmt.Printf("COUNT: %d\n", count)
	assert.Equal(t, 1767, count)
}

func hasDigitMatch(target string) bool {
	hasMatch := false
	for j:= 0; j < len(target) -1; j++ {
		if target[j] == target[j+1] {
			hasMatch = true
			break
		}
	}

	return hasMatch
}

func digitsAreIncreasing(target string) bool {
	increasing := true
	for j:= 0; j < len(target) -1; j++ {
		if target[j] > target[j+1] {
			increasing = false
			break
		}
	}

	return increasing
}
