package ae

func NewFileError(tag string, err error) *Error {
	if err == nil {
		return nil
	}

	return NewError(500, tag+" file error: "+err.Error())

}
