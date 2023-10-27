package ae

func NewJsonError(err error) *Error {
	if err == nil {
		return nil
	}
	pos := Caller(1)
	return NewError(500, pos+" json marshal/unmarshal error: "+err.Error())
}
