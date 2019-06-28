package ae

func NewHttpError(err error) *Error {
	if err == nil {
		return nil
	}

	switch err {

	}

	return NewError(500, "http: "+err.Error())

}
