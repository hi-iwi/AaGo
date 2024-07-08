package ae

func Check(es ...*Error) *Error {
	for _, e := range es {
		if e != nil {
			return e
		}
	}
	return nil
}

func CheckError(es ...error) error {
	for _, err := range es {
		if err != nil {
			return err
		}
	}
	return nil
}
