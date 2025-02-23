package matrix

import (
	"testing"
	"math"
)

func TestNewCOOMatrix(t *testing.T) {
	tests := []struct {
		name       string
		values     []float64
		rowIndices []int
		colIndices []int
		rows       int
		cols       int
		wantErr    bool
	}{
		{
			name:       "valid matrix",
			values:     []float64{1.0, 2.0, 3.0},
			rowIndices: []int{0, 0, 1},
			colIndices: []int{0, 1, 1},
			rows:       2,
			cols:       2,
			wantErr:    false,
		},
		{
			name:       "mismatched lengths",
			values:     []float64{1.0, 2.0},
			rowIndices: []int{0, 0, 1},
			colIndices: []int{0, 1},
			rows:       2,
			cols:       2,
			wantErr:    true,
		},
		{
			name:       "out of bounds row",
			values:     []float64{1.0},
			rowIndices: []int{2},
			colIndices: []int{0},
			rows:       2,
			cols:       2,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewCOOMatrix(tt.values, tt.rowIndices, tt.colIndices, tt.rows, tt.cols)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCOOMatrix() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCOOToCSRConversion(t *testing.T) {
	// Create a COO matrix
	coo, _ := NewCOOMatrix(
		[]float64{1.0, 3.0, 2.0},  // Values out of order
		[]int{0, 1, 0},            // Row indices
		[]int{0, 1, 1},            // Column indices
		2, 2,                      // Dimensions
	)

	// Convert to CSR
	csr, err := coo.ToCSR()
	if err != nil {
		t.Fatalf("Failed to convert to CSR: %v", err)
	}

	// Check dimensions
	if csr.Rows != 2 || csr.Cols != 2 {
		t.Errorf("Wrong dimensions after conversion")
	}

	// Check if values are correctly sorted
	expectedValues := []float64{1.0, 2.0, 3.0}
	expectedRowPtr := []int{0, 2, 3}
	expectedColIndices := []int{0, 1, 1}

	// Check values
	if len(csr.Values) != len(expectedValues) {
		t.Errorf("Wrong number of values")
	}
	for i, v := range expectedValues {
		if math.Abs(csr.Values[i]-v) > 1e-15 {
			t.Errorf("Wrong value at position %d", i)
		}
	}

	// Check row pointers
	if len(csr.RowPtr) != len(expectedRowPtr) {
		t.Errorf("Wrong row pointer array length")
	}
	for i, v := range expectedRowPtr {
		if csr.RowPtr[i] != v {
			t.Errorf("Wrong row pointer at position %d", i)
		}
	}

	// Check column indices
	if len(csr.ColIndices) != len(expectedColIndices) {
		t.Errorf("Wrong column indices array length")
	}
	for i, v := range expectedColIndices {
		if csr.ColIndices[i] != v {
			t.Errorf("Wrong column index at position %d", i)
		}
	}
}

func TestCSRToCOOConversion(t *testing.T) {
	// Create a CSR matrix
	csr, _ := NewCSRMatrix(
		[]float64{1.0, 2.0, 3.0},
		[]int{0, 2, 3},
		[]int{0, 1, 1},
		2, 2,
	)

	// Convert to COO
	coo, err := FromCSR(csr)
	if err != nil {
		t.Fatalf("Failed to convert to COO: %v", err)
	}

	// Check dimensions
	if coo.Rows != 2 || coo.Cols != 2 {
		t.Errorf("Wrong dimensions after conversion")
	}

	// Expected values
	expectedRowIndices := []int{0, 0, 1}

	// Check values (should be preserved)
	for i, v := range csr.Values {
		if math.Abs(coo.Values[i]-v) > 1e-15 {
			t.Errorf("Wrong value at position %d", i)
		}
	}

	// Check row indices
	for i, v := range expectedRowIndices {
		if coo.RowIndices[i] != v {
			t.Errorf("Wrong row index at position %d", i)
		}
	}

	// Check column indices (should be preserved)
	for i, v := range csr.ColIndices {
		if coo.ColIndices[i] != v {
			t.Errorf("Wrong column index at position %d", i)
		}
	}
}

func TestEmptyMatrixConversion(t *testing.T) {
	// Create empty COO matrix
	coo, _ := NewCOOMatrix(
		[]float64{},
		[]int{},
		[]int{},
		2, 2,
	)

	// Convert to CSR
	csr, err := coo.ToCSR()
	if err != nil {
		t.Fatalf("Failed to convert empty matrix to CSR: %v", err)
	}

	// Check dimensions
	if csr.Rows != 2 || csr.Cols != 2 {
		t.Errorf("Wrong dimensions for empty matrix")
	}

	// Check that arrays are empty
	if len(csr.Values) != 0 || len(csr.ColIndices) != 0 {
		t.Errorf("Non-empty arrays in empty matrix")
	}

	// Check row pointers are all zero
	expectedRowPtr := make([]int, 3) // rows + 1
	for i, v := range csr.RowPtr {
		if v != expectedRowPtr[i] {
			t.Errorf("Wrong row pointer at position %d", i)
		}
	}
}
