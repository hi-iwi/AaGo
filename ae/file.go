package ae

func NewFileError(err error) *Error {
	if err == nil {
		return nil
	}

	return NewError(500, "file error: "+err.Error())

}
