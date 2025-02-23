package matrix

import (
    "testing"
    "math"
)

func TestMatVec(t *testing.T) {
    values := []float64{1.0, 2.0, 3.0}
    rowPtr := []int{0, 2, 3}
    colIndices := []int{0, 1, 1}
    
    m, _ := NewCSRMatrix(values, rowPtr, colIndices, 2, 2)
    vec := []float64{1.0, 1.0}
    
    result, err := m.MatVec(vec)
    if err != nil {
        t.Errorf("MatVec failed: %v", err)
    }
    
    expected := []float64{3.0, 3.0}
    for i := range result {
        if math.Abs(result[i]-expected[i]) > 1e-15 {
            t.Errorf("MatVec result[%d] = %f; want %f", i, result[i], expected[i])
        }
    }
}

func TestAdd(t *testing.T) {
    m1, _ := NewCSRMatrix(
        []float64{1.0, 2.0},
        []int{0, 1, 2},
        []int{0, 1},
        2, 2,
    )
    
    m2, _ := NewCSRMatrix(
        []float64{1.0, 3.0},
        []int{0, 1, 2},
        []int{0, 1},
        2, 2,
    )
    
    result, err := m1.Add(m2)
    if err != nil {
        t.Errorf("Add failed: %v", err)
    }
    
    expected := []float64{2.0, 5.0}
    for i := range expected {
        if math.Abs(result.Values[i]-expected[i]) > 1e-15 {
            t.Errorf("Add result.Values[%d] = %f; want %f", i, result.Values[i], expected[i])
        }
    }
}
