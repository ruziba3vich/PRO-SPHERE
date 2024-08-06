package errors

type (
	ProError struct {
		Message string `json:"message"`
		Err     string `json:"error"`
	}
)
