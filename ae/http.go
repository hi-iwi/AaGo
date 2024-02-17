package ae

func NewHttpError(err error) *Error {
	if err == nil {
		return nil
	}
	msg, pos := CallerMsg(err.Error(), 1)
	return NewError(500, pos+" http: "+msg)

}
