package ae

func NewSerializationError(err error) *Error {
	if err == nil {
		return nil
	}
	pos := Caller(1)
	return NewError(500, pos+" serialization/unserialization error: "+err.Error())
}
