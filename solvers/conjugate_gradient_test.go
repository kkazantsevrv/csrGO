package solvers

import (
	"math"
	"testing"

	"github.com/gosparse/matrix"
)

func TestCGSolver(t *testing.T) {
	// Create a simple symmetric positive-definite matrix
	A, _ := matrix.NewCSRMatrix(
		[]float64{2.0, -1.0, -1.0, 2.0},
		[]int{0, 2, 4},
		[]int{0, 1, 0, 1},
		2, 2,
	)

	b := []float64{1.0, 1.0}

	solver := NewCGSolver(1000, 1e-10)
	x, err := solver.Solve(A, b)

	if err != nil {
		t.Errorf("Solver failed: %v", err)
	}

	// Check if Ax = b
	Ax, _ := A.MatVec(x)
	for i := range b {
		if math.Abs(Ax[i]-b[i]) > 1e-8 {
			t.Errorf("Solution is not accurate: Ax[%d] = %f, b[%d] = %f", i, Ax[i], i, b[i])
		}
	}
}

func TestCGSolverConvergence(t *testing.T) {
	// Create a larger matrix to test convergence
	n := 10
	values := make([]float64, n)
	rowPtr := make([]int, n+1)
	colIndices := make([]int, n)

	for i := 0; i < n; i++ {
		values[i] = 2.0
		rowPtr[i] = i
		colIndices[i] = i
	}
	rowPtr[n] = n

	A, _ := matrix.NewCSRMatrix(values, rowPtr, colIndices, n, n)

	b := make([]float64, n)
	for i := range b {
		b[i] = 1.0
	}

	solver := NewCGSolver(1000, 1e-10)
	x, err := solver.Solve(A, b)

	if err != nil {
		t.Errorf("Solver failed to converge: %v", err)
	}

	// Check solution
	for i := range x {
		if math.Abs(x[i]-0.5) > 1e-8 {
			t.Errorf("Incorrect solution value at x[%d] = %f, expected 0.5", i, x[i])
		}
	}
}
