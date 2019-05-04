// create common matrices and vectors
package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Identity creates an identity matrix
func Identity(m, n int) Matrix {
	// This is a naive implementaiton for initial design clarity
	// using the el field instead of rows method would be faster
	mat := Zeros(m, n)
	iterateRows(&mat, func(i, j int) {
		if i == j {
			mat.set(i, j, 1)
		}
	})
	return mat
}

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

func arrayToString(a []float64, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
	//return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
	//return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}

// FloatToString converts a float to a string
func FloatToString(inputNum float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(inputNum, 'f', 6, 64)
}
