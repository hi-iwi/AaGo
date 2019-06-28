package ae

func NewRedisError(err error) *Error {
	if err == nil {
		return nil
	}

	switch err {

	}

	return NewError(500, "redis error: "+err.Error())

}
