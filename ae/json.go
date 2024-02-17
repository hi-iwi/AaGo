package ae

func NewJsonError(err error) *Error {
	if err == nil {
		return nil
	}
	msg, pos := CallerMsg(err.Error(), 1)
	return NewError(500, pos+" json marshal/unmarshal error: "+msg)
}
