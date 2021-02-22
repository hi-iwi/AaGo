package ae

func NewFileError(err error) *Error {
 	if err == nil {
		return nil
	}
	pos := Caller(1)
	return NewError(500, pos+" file error: "+err.Error())

}
