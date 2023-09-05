package errors

import "fmt"

type AppError struct {
    Code    int
    Message string
}

func (e *AppError) Error() string {
    return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}
