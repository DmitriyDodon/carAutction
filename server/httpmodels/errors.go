package httpmodels

import "fmt"

type CustomHttpError struct {
	Message string `json:"message"`
}

func (c CustomHttpError) Error() string {
	return fmt.Sprintf(" Message: %s", c.Message)
}

var UnprocessableEntity = CustomHttpError{
	Message: "Invalid entity.",
}

var ServerError = CustomHttpError{
	Message: "Iternal server error.",
}

var NotFoundError = CustomHttpError{
	Message: "Not found.",
}
