package matrix

import (
    "testing"
    "math"
)

func TestNewCSRMatrix(t *testing.T) {
    values := []float64{1.0, 2.0, 3.0}
    rowPtr := []int{0, 2, 3}
    colIndices := []int{0, 1, 1}
    
    m, err := NewCSRMatrix(values, rowPtr, colIndices, 2, 2)
    if err != nil {
        t.Errorf("Failed to create CSR matrix: %v", err)
    }
    
    if m.Rows != 2 || m.Cols != 2 {
        t.Errorf("Wrong matrix dimensions")
    }
}

func TestGet(t *testing.T) {
    values := []float64{1.0, 2.0, 3.0}
    rowPtr := []int{0, 2, 3}
    colIndices := []int{0, 1, 1}
    
    m, _ := NewCSRMatrix(values, rowPtr, colIndices, 2, 2)
    
    tests := []struct {
        i, j     int
        expected float64
    }{
        {0, 0, 1.0},
        {0, 1, 2.0},
        {1, 1, 3.0},
        {1, 0, 0.0}, // Zero element
    }
    
    for _, test := range tests {
        got := m.Get(test.i, test.j)
        if math.Abs(got-test.expected) > 1e-15 {
            t.Errorf("Get(%d,%d) = %f; want %f", test.i, test.j, got, test.expected)
        }
    }
}

func TestFromDense(t *testing.T) {
    dense := [][]float64{
        {1.0, 0.0, 2.0},
        {0.0, 3.0, 0.0},
    }
    
    m, err := FromDense(dense)
    if err != nil {
        t.Errorf("Failed to create CSR from dense: %v", err)
    }
    
    if len(m.Values) != 3 {
        t.Errorf("Wrong number of non-zero elements")
    }
}
