package solvers

import (
	"fmt"
	"math"

	"github.com/gosparse/matrix"
)

// CGSolver represents a Conjugate Gradient solver
type CGSolver struct {
	MaxIter   int
	Tolerance float64
}

// NewCGSolver creates a new Conjugate Gradient solver
func NewCGSolver(maxIter int, tolerance float64) *CGSolver {
	return &CGSolver{
		MaxIter:   maxIter,
		Tolerance: tolerance,
	}
}

// Solve solves the system Ax = b using the Conjugate Gradient method
func (cg *CGSolver) Solve(A *matrix.CSRMatrix, b []float64) ([]float64, error) {
	if A.Rows != len(b) {
		return nil, fmt.Errorf("matrix and vector dimensions mismatch")
	}

	n := len(b)
	x := make([]float64, n) // Initial guess x = 0

	// r = b - Ax
	Ax, err := A.MatVec(x)
	if err != nil {
		return nil, err
	}

	r := make([]float64, n)
	for i := range b {
		r[i] = b[i] - Ax[i]
	}

	p := make([]float64, n)
	copy(p, r)

	rsold := dot(r, r)

	if math.Sqrt(rsold) < cg.Tolerance {
		return x, nil
	}

	for iter := 0; iter < cg.MaxIter; iter++ {
		Ap, err := A.MatVec(p)
		if err != nil {
			return nil, err
		}

		alpha := rsold / dot(p, Ap)

		// x = x + alpha*p
		for i := range x {
			x[i] += alpha * p[i]
		}

		// r = r - alpha*Ap
		for i := range r {
			r[i] -= alpha * Ap[i]
		}

		rsnew := dot(r, r)
		if math.Sqrt(rsnew) < cg.Tolerance {
			return x, nil
		}

		beta := rsnew / rsold

		// p = r + beta*p
		for i := range p {
			p[i] = r[i] + beta*p[i]
		}

		rsold = rsnew
	}

	return nil, fmt.Errorf("maximum iterations reached without convergence")
}

// dot computes the dot product of two vectors
func dot(a, b []float64) float64 {
	sum := 0.0
	for i := range a {
		sum += a[i] * b[i]
	}
	return sum
}
