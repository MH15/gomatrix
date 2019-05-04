package main

var (
	ErrIndexOutOfRange     = "matrix: index out of range"
	ErrRowAccess           = "matrix: row index out of range"
	ErrColAccess           = "matrix: column index out of range"
	ErrVectorAccess        = "matrix: vector index out of range"
	ErrZeroLength          = "matrix: zero length in matrix definition"
	ErrRowLength           = "matrix: row length mismatch"
	ErrColLength           = "matrix: col length mismatch"
	ErrSquare              = "matrix: expect square matrix"
	ErrNormOrder           = "matrix: invalid norm order for matrix"
	ErrSingular            = "matrix: matrix is singular"
	ErrShape               = "matrix: dimension mismatch"
	ErrIllegalStride       = "matrix: illegal stride"
	ErrPivot               = "matrix: malformed pivot list"
	ErrTriangle            = "matrix: triangular storage mismatch"
	ErrTriangleSet         = "matrix: triangular set out of bounds"
	ErrSliceLengthMismatch = "matrix: input slice length mismatch"
	ErrNotPSD              = "matrix: input not positive symmetric definite"
	ErrFailedEigen         = "matrix: eigendecomposition not successful"
)
