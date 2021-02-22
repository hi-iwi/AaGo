package ae

func NewHttpError(err error) *Error {
	if err == nil {
		return nil
	}
	pos := Caller(1)

	return NewError(500, pos+" http: "+err.Error())

}
