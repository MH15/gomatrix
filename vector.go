package main

import "sync"

// Vector is the functions for vector operations.
// Vector must be implemented for each numberical type.
type Vector []float64

// Vec makes a new Vector
func Vec(s ...float64) Vector {
	v := make([]float64, len(s))
	for i := range s {
		v[i] = s[i]
	}
	return v
}

func ZeroVec(l int) Vector {
	return make(Vector, l)
}

// Len returns the length
func (a Vector) Len() int {
	return len(a)
}

// At returns the value at position p in vector
func (a Vector) At(p int) float64 {
	return a[p]
}

// Equal checks if vectors are the same size and are equal element-wise.
func Equal(a, b Vector) bool {
	panicLength(a, b)
	for i := 0; i < a.Len(); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// EqualSize checks dimensions of all vectos
func EqualSize(a ...Vector) bool {
	allSame := true
	len := a[0].Len()
	for _, v := range a {
		if v.Len() != len {
			allSame = false
		}
	}
	return allSame

}

// Dot returns the dot product
func Dot(a, b Vector) float64 {
	panicLength(a, b)
	var dot float64
	for i := 0; i < a.Len(); i++ {
		dot += a[i] * b[i]
	}
	return dot
}

/// OPERATIONS
func (a Vector) add(b Vector) {
	panicLength(a, b)
	for i := 0; i < a.Len(); i++ {
		a[i] += b[i]
	}
}
func (a Vector) fastAdd(b Vector) {
	panicLength(a, b)

	var wg sync.WaitGroup
	wg.Add(a.Len())

	for i := 0; i < a.Len(); i++ {
		go func(i int) {
			defer wg.Done()
			a[i] += b[i]
		}(i)
		a[i] += b[i]
	}

	wg.Wait()

}

func (a Vector) subtract(b Vector) {
	panicLength(a, b)
	for i := 0; i < a.Len(); i++ {
		a[i] -= b[i]
	}
}
func (a Vector) scalarMult(s float64) {
	for i := 0; i < a.Len(); i++ {
		a[i] *= s
	}
}

/// PANICS
func panicLength(a, b Vector) {
	if a.Len() != b.Len() {
		panic("matrix.ErrShape")
	}
}
