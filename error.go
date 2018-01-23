package errslice

type Error []error

func (e Error) Error() (res string) {
	for i, e1 := range e {
		if i > 0 {
			res += ","
		}
		res += e1.Error()
	}
	return
}

