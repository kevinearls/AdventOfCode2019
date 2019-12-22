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
		hasMatch := hasDigitMatch2(s)
		increasing := digitsAreIncreasing(s)

		if hasMatch && increasing {
			count++
		}
	}

	fmt.Printf("COUNT: %d\n", count)
	//assert.Equal(t, 1767, count)  // For Part 1
	assert.Equal(t, 1192, count)  // For Part 2

}

// This is for part 1
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

// This is for part 2
func hasDigitMatch2(target string) bool {
	hasMatch := false
	for j:= 0; j < len(target) -1; j++ {
		if target[j] == target[j+1] {

			// Peek one character behind or ahead to make sure we only have 2 of the same in a row
			if (j == 0 || target[j-1] != target[j]) && (j >= len(target)-2 || target[j+1] != target[j+2]) {
				hasMatch = true
				break
			}
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

func TestHasDigitMatch2(t *testing.T) {
	assert.True(t, hasDigitMatch2("112233"), "First")
	assert.False(t, hasDigitMatch2("123444"), "second")
	assert.True(t, hasDigitMatch2("111122"), "Third")
}


