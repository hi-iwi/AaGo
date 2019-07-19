package ae

func NewRedisError(err error) *Error {
	if err == nil {
		return nil
	}

	if err.Error() == "redigo: nil returned" {
		return NewError(404, "redis not found")
	}

	return NewError(500, "redis error: "+err.Error())

}
