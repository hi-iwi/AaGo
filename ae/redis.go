package ae

func NewRedisError(err error) *Error {
	if err == nil {
		return nil
	}
	pos := Caller(1)
	if err.Error() == "redigo: nil returned" {
		return NewError(404, "Cache Not Found")
	}

	return NewError(500, pos+" redis: "+err.Error())
}
