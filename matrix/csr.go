package matrix

import (
    "fmt"
    "math"
)

// CSRMatrix represents a sparse matrix in Compressed Sparse Row format
type CSRMatrix struct {
    Values    []float64 // Non-zero values
    RowPtr    []int     // Row pointers
    ColIndices []int    // Column indices
    Rows      int       // Number of rows
    Cols      int       // Number of columns
}

// NewCSRMatrix creates a new CSR matrix from the given values
func NewCSRMatrix(values []float64, rowPtr []int, colIndices []int, rows, cols int) (*CSRMatrix, error) {
    if len(rowPtr) != rows+1 {
        return nil, fmt.Errorf("invalid row pointer array length: expected %d, got %d", rows+1, len(rowPtr))
    }
    
    if len(values) != len(colIndices) {
        return nil, fmt.Errorf("values and column indices must have same length")
    }
    
    return &CSRMatrix{
        Values:     values,
        RowPtr:     rowPtr,
        ColIndices: colIndices,
        Rows:       rows,
        Cols:       cols,
    }, nil
}

// Get returns the value at position (i,j)
func (m *CSRMatrix) Get(i, j int) float64 {
    if i < 0 || i >= m.Rows || j < 0 || j >= m.Cols {
        return 0.0
    }
    
    // Search in the row
    for k := m.RowPtr[i]; k < m.RowPtr[i+1]; k++ {
        if m.ColIndices[k] == j {
            return m.Values[k]
        }
    }
    return 0.0
}

// Set sets the value at position (i,j)
// Note: This is not efficient for CSR format and should be used sparingly
func (m *CSRMatrix) Set(i, j int, value float64) error {
    if i < 0 || i >= m.Rows || j < 0 || j >= m.Cols {
        return fmt.Errorf("index out of bounds")
    }
    
    // Find if element exists
    for k := m.RowPtr[i]; k < m.RowPtr[i+1]; k++ {
        if m.ColIndices[k] == j {
            m.Values[k] = value
            return nil
        }
    }
    
    // Element doesn't exist, would need matrix reconstruction
    return fmt.Errorf("cannot set new non-zero element in CSR format directly")
}

// FromDense creates a CSR matrix from a dense matrix representation
func FromDense(dense [][]float64) (*CSRMatrix, error) {
    if len(dense) == 0 {
        return nil, fmt.Errorf("empty matrix")
    }
    
    rows := len(dense)
    cols := len(dense[0])
    
    var values []float64
    var colIndices []int
    rowPtr := make([]int, rows+1)
    
    count := 0
    for i := 0; i < rows; i++ {
        rowPtr[i] = count
        for j := 0; j < cols; j++ {
            if math.Abs(dense[i][j]) > 1e-15 {
                values = append(values, dense[i][j])
                colIndices = append(colIndices, j)
                count++
            }
        }
    }
    rowPtr[rows] = count
    
    return NewCSRMatrix(values, rowPtr, colIndices, rows, cols)
}
