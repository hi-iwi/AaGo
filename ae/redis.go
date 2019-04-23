package ae

import "fmt"

func NewRedisError(err error) *Error {
	if err == nil {
		return nil
	}

	switch err {

	}

	return NewError(500, fmt.Sprintf("redis: %s", err))

}
