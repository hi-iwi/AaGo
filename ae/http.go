package ae

import "fmt"

func NewHttpError(err error) *Error {
	if err == nil {
		return nil
	}

	switch err {

	}

	return NewError(500, fmt.Sprintf("http: %s", err))

}
