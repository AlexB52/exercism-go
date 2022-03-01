package erratum

func Use(opener ResourceOpener, input string) error {
	resource, err := opener()

	if err != nil {
		switch err.(type) {
		case TransientError:
			return Use(opener, input)
		default:
			return err
		}
	}

	resource.Frob(input)
	err = resource.Close()
	if err != nil {
		return err
	}
	return err
}
