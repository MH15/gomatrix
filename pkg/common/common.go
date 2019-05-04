// create common matrices and vectors
package common

import (
	"strconv"
)

/// HELPERS
func notPositive(x ...int) bool {
	result := false
	for _, i := range x {
		if i < 0 {
			result := true
			return result
		}
	}
	return result
}

// FloatToString converts a float to a string
func FloatToString(inputNum float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(inputNum, 'f', 6, 64)
}
