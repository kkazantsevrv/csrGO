package matrix

import (
    "fmt"
)

// MatVec multiplies CSR matrix with a vector
func (m *CSRMatrix) MatVec(vec []float64) ([]float64, error) {
    if len(vec) != m.Cols {
        return nil, fmt.Errorf("vector length mismatch: expected %d, got %d", m.Cols, len(vec))
    }
    
    result := make([]float64, m.Rows)
    for i := 0; i < m.Rows; i++ {
        sum := 0.0
        for j := m.RowPtr[i]; j < m.RowPtr[i+1]; j++ {
            sum += m.Values[j] * vec[m.ColIndices[j]]
        }
        result[i] = sum
    }
    
    return result, nil
}

// Add adds two CSR matrices
func (m *CSRMatrix) Add(other *CSRMatrix) (*CSRMatrix, error) {
    if m.Rows != other.Rows || m.Cols != other.Cols {
        return nil, fmt.Errorf("matrix dimensions mismatch")
    }
    
    // Pre-calculate size of result
    nnz := 0
    rowPtr := make([]int, m.Rows+1)
    
    for i := 0; i < m.Rows; i++ {
        rowPtr[i] = nnz
        
        // Merge row entries from both matrices
        p1 := m.RowPtr[i]
        p2 := other.RowPtr[i]
        
        for p1 < m.RowPtr[i+1] || p2 < other.RowPtr[i+1] {
            if p1 == m.RowPtr[i+1] {
                nnz++
                p2++
            } else if p2 == other.RowPtr[i+1] {
                nnz++
                p1++
            } else if m.ColIndices[p1] < other.ColIndices[p2] {
                nnz++
                p1++
            } else if m.ColIndices[p1] > other.ColIndices[p2] {
                nnz++
                p2++
            } else {
                nnz++
                p1++
                p2++
            }
        }
    }
    rowPtr[m.Rows] = nnz
    
    // Allocate result arrays
    values := make([]float64, nnz)
    colIndices := make([]int, nnz)
    
    // Fill result arrays
    pos := 0
    for i := 0; i < m.Rows; i++ {
        p1 := m.RowPtr[i]
        p2 := other.RowPtr[i]
        
        for p1 < m.RowPtr[i+1] || p2 < other.RowPtr[i+1] {
            if p1 == m.RowPtr[i+1] {
                values[pos] = other.Values[p2]
                colIndices[pos] = other.ColIndices[p2]
                pos++
                p2++
            } else if p2 == other.RowPtr[i+1] {
                values[pos] = m.Values[p1]
                colIndices[pos] = m.ColIndices[p1]
                pos++
                p1++
            } else if m.ColIndices[p1] < other.ColIndices[p2] {
                values[pos] = m.Values[p1]
                colIndices[pos] = m.ColIndices[p1]
                pos++
                p1++
            } else if m.ColIndices[p1] > other.ColIndices[p2] {
                values[pos] = other.Values[p2]
                colIndices[pos] = other.ColIndices[p2]
                pos++
                p2++
            } else {
                values[pos] = m.Values[p1] + other.Values[p2]
                colIndices[pos] = m.ColIndices[p1]
                pos++
                p1++
                p2++
            }
        }
    }
    
    return NewCSRMatrix(values, rowPtr, colIndices, m.Rows, m.Cols)
}
