package ae

func NewFileError(err error) *Error {
	if err == nil {
		return nil
	}
	m, pos := CallerMsg(err.Error(), 1)
	return NewError(500, pos+" file error: "+m)

}
