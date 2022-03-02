package erratum

func Use(opener ResourceOpener, input string) (err error) {
	resource, err := opener()

	if _, ok := err.(TransientError); ok {
		return Use(opener, input)
	} else if err != nil {
		return err
	}

	defer resource.Close()

	defer func() {
		if r := recover(); r != nil {
			if frobErr, ok := r.(FrobError); ok {
				resource.Defrob(frobErr.defrobTag)
			}
			err = r.(error)
		}
	}()

	resource.Frob(input)
	return err
}
