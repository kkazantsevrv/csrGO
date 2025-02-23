package utils

// MatrixError represents an error that occurred during matrix operations
type MatrixError struct {
    message string
}

func (e *MatrixError) Error() string {
    return e.message
}

// NewMatrixError creates a new matrix error
func NewMatrixError(message string) *MatrixError {
    return &MatrixError{message: message}
}
