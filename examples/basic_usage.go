package main

import (
	"fmt"

	"github.com/gosparse/matrix"
	"github.com/gosparse/solvers"
)

func main() {
	// Create a sample sparse matrix in CSR format
	values := []float64{4.0, 1.0, 1.0, 4.0}
	rowPtr := []int{0, 0, 1, 1}
	colIndices := []int{0, 1, 0, 1}
	coo, _ := matrix.NewCOOMatrix(values, rowPtr, colIndices, 2, 2)
	A, _ := coo.ToCSR()
	// A, err := matrix.NewCSRMatrix(values, rowPtr, colIndices, 2, 2)
	// if err != nil {
	// 	fmt.Printf("Error creating matrix: %v\n", err)
	// 	return
	// }
	// Create right-hand side vector
	b := []float64{1.0, 1.0}

	// Create solver
	solver := solvers.NewCGSolver(1000, 1e-10)

	// Solve system
	x, err := solver.Solve(A, b)
	if err != nil {
		fmt.Printf("Error solving system: %v\n", err)
		return
	}

	fmt.Println("Solution:", x)
}
