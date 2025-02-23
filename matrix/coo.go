package matrix

import (
	"fmt"
	"sort"
)

// COOMatrix represents a sparse matrix in Coordinate (COO) format
type COOMatrix struct {
	Values     []float64 // Non-zero values
	RowIndices []int     // Row indices
	ColIndices []int     // Column indices
	Rows       int       // Number of rows
	Cols       int       // Number of columns
}

// NewCOOMatrix creates a new COO matrix from the given values and indices
func NewCOOMatrix(values []float64, rowIndices, colIndices []int, rows, cols int) (*COOMatrix, error) {
	if len(values) != len(rowIndices) || len(values) != len(colIndices) {
		return nil, fmt.Errorf("values, row indices, and column indices must have same length")
	}

	// Check indices are within bounds
	for i, row := range rowIndices {
		if row < 0 || row >= rows {
			return nil, fmt.Errorf("row index out of bounds at position %d", i)
		}
		if colIndices[i] < 0 || colIndices[i] >= cols {
			return nil, fmt.Errorf("column index out of bounds at position %d", i)
		}
	}

	return &COOMatrix{
		Values:     values,
		RowIndices: rowIndices,
		ColIndices: colIndices,
		Rows:       rows,
		Cols:       cols,
	}, nil
}

// ToCSR converts the COO matrix to CSR format
func (m *COOMatrix) ToCSR() (*CSRMatrix, error) {
	if len(m.Values) == 0 {
		// Handle empty matrix
		rowPtr := make([]int, m.Rows+1)
		return NewCSRMatrix([]float64{}, rowPtr, []int{}, m.Rows, m.Cols)
	}

	// Create index array for sorting
	indices := make([]int, len(m.Values))
	for i := range indices {
		indices[i] = i
	}

	// Sort by row, then by column
	sort.Slice(indices, func(i, j int) bool {
		if m.RowIndices[indices[i]] != m.RowIndices[indices[j]] {
			return m.RowIndices[indices[i]] < m.RowIndices[indices[j]]
		}
		return m.ColIndices[indices[i]] < m.ColIndices[indices[j]]
	})

	// Create sorted arrays
	values := make([]float64, len(m.Values))
	colIndices := make([]int, len(m.Values))
	rowPtr := make([]int, m.Rows+1)

	// Fill the arrays
	currentRow := -1
	for i, idx := range indices {
		row := m.RowIndices[idx]
		
		// Fill empty rows
		for r := currentRow + 1; r <= row; r++ {
			rowPtr[r] = i
		}
		currentRow = row
		
		values[i] = m.Values[idx]
		colIndices[i] = m.ColIndices[idx]
	}

	// Fill remaining empty rows
	for r := currentRow + 1; r <= m.Rows; r++ {
		rowPtr[r] = len(values)
	}

	return NewCSRMatrix(values, rowPtr, colIndices, m.Rows, m.Cols)
}

// FromCSR converts a CSR matrix to COO format
func FromCSR(csr *CSRMatrix) (*COOMatrix, error) {
	nnz := len(csr.Values)
	rowIndices := make([]int, nnz)
	
	// Generate row indices from row pointers
	for i := 0; i < csr.Rows; i++ {
		for j := csr.RowPtr[i]; j < csr.RowPtr[i+1]; j++ {
			rowIndices[j] = i
		}
	}

	return NewCOOMatrix(
		append([]float64{}, csr.Values...),
		rowIndices,
		append([]int{}, csr.ColIndices...),
		csr.Rows,
		csr.Cols,
	)
}
