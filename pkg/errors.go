package pkg

import "fmt"

type InvalidParameterError struct {
	Reason string
	Code   string
}

func (e InvalidParameterError) Error() string {
	return fmt.Sprintf("%s:%s", e.Code, e.Reason)
}
