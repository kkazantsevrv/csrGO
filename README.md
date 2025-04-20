# GoSparse

GoSparse is a Go library for sparse matrix operations and linear system equation solving, similar to AMGCL. It provides efficient implementations for working with sparse matrices in Compressed Sparse Row (CSR) format and includes solvers for linear systems.

## Features

- Multiple sparse matrix formats:
  - Compressed Sparse Row (CSR) format
  - Coordinate (COO) format
  - Format conversions (CSR ↔ COO)
- Basic matrix operations (addition, matrix-vector multiplication)
- Conjugate Gradient solver for sparse linear systems
- Efficient memory usage for sparse matrices

## Installation

```bash
go get github.com/gosparse
```

## Usage Example

Here's a simple example of solving a linear system Ax = b using the Conjugate Gradient solver:

```go
package main

import (
    "fmt"
    "github.com/gosparse/matrix"
    "github.com/gosparse/solvers"
)

func main() {
    // Create a sample sparse matrix in CSR format
    values := []float64{4.0, 1.0, 1.0, 4.0}
    rowPtr := []int{0, 2, 4}
    colIndices := []int{0, 1, 0, 1}

    A, _ := matrix.NewCSRMatrix(values, rowPtr, colIndices, 2, 2)

    // Create right-hand side vector
    b := []float64{1.0, 1.0}

    // Create solver with maximum iterations and tolerance
    solver := solvers.NewCGSolver(1000, 1e-10)

    // Solve the system
    x, err := solver.Solve(A, b)
    if err != nil {
        fmt.Printf("Error solving system: %v\n", err)
        return
    }

    fmt.Println("Solution:", x)
}
```

## Matrix Format Examples

### CSR Format

The Compressed Sparse Row format is efficient for matrix operations:

```go
// Create a CSR matrix
csrMatrix, _ := matrix.NewCSRMatrix(
    []float64{1.0, 2.0, 3.0},  // Values
    []int{0, 2, 3},            // Row pointers
    []int{0, 1, 1},            // Column indices
    2, 2,                      // Dimensions
)
```

### COO Format

The Coordinate format is convenient for building matrices incrementally:

```go
// Create a COO matrix
cooMatrix, _ := matrix.NewCOOMatrix(
    []float64{1.0, 2.0, 3.0},  // Values
    []int{0, 0, 1},            // Row indices
    []int{0, 1, 1},            // Column indices
    2, 2,                      // Dimensions
)

// Convert to CSR format for efficient operations
csrMatrix, _ := cooMatrix.ToCSR()
```

## API Documentation

### Matrix Package

- `NewCSRMatrix`: Creates a new CSR matrix from values, row pointers, and column indices
- `NewCOOMatrix`: Creates a new COO matrix from values, row indices, and column indices
- `ToCSR`: Converts a COO matrix to CSR format
- `FromCSR`: Converts a CSR matrix to COO format
- `Get`: Retrieves a value from the matrix at given indices
- `Set`: Sets a value in the matrix (note: inefficient for CSR format)
- `MatVec`: Performs matrix-vector multiplication
- `Add`: Adds two CSR matrices

### Solvers Package

- `NewCGSolver`: Creates a new Conjugate Gradient solver
- `Solve`: Solves a linear system using the Conjugate Gradient method

## Testing

Run the tests using:

```bash
go test ./...
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License
TODO
