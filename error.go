package gochimp

import (
	"fmt"
)

// ErrorResponse ...
type ErrorResponse struct {
	Type   string `json:"type"`
	Title  string `json:"title"`
	Status int    `json:"status"`
	Detail string `json:"detail"`
}

// Error ...
func (e ErrorResponse) Error() string {
	return fmt.Sprintf("Error %d %s (%s)", e.Status, e.Title, e.Detail)
}
