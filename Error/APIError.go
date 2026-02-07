// You can edit this code!
// Click here and start typing.
package main

import (
	"errors"
	"fmt"
)

type APIError struct {
	code    int
	message string
	err     error
}

func (apierr *APIError) Error() string {
	return fmt.Sprintf("Code : %d Message :%s Error:%v", apierr.code, apierr.message, apierr.err)
}

func ThrowAPIError() error {
	return &APIError{
		code:    404,
		message: "Not Found",
		err:     errors.New("Wrapped Error"),
	}
}
func main() {
	var apiErr *APIError
	err := ThrowAPIError()
	if err != nil {
		if errors.As(err, &apiErr) {
			fmt.Println(err)
		}
	}
}
