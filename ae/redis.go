package ae

func NewRedisError(tag string, err error) *Error {
	if err == nil {
		return nil
	}

	if err.Error() == "redigo: nil returned" {
		return NewError(404, tag+" redis not found")
	}

	return NewError(500, tag+" redis: "+err.Error())

}
