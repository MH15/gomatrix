// Package matrix code to create common matrix structures
package matrix

import "github.com/MH15/gomatrix/pkg/vector"

// Zeros creates a zero-filled matrix
func Zeros(m, n int) Matrix {
	if notPositive(m, n) {
		panic("matrix.dimensionsMustbePositive")
	}
	return Matrix{el: make(vector.Vector, m*n), stride: n, class: FloatingPoint}
}

// Ones creates a one-filled matrix
func Ones(m, n int) Matrix {
	mat := Zeros(m, n)
	iterateRows(&mat, func(i, j int) {
		mat.set(i, j, 1)
	})
	return mat
}

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

// Diagonal creates a n-by-n matrix with a Vector of n elements spread along the
// diagonal entries of the matrix
func Diagonal(v vector.Vector, super ...int) Matrix {
	mat := Zeros(len(v), len(v))
	iterateRows(&mat, func(i, j int) {
		if i == j {
			mat.set(i, j, v.At(i))
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
