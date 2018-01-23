package errslice

// Append takes two errors and
func Append(left, right error) (err error) {
	if right == nil {
		return left
	}
	if left == nil {
		return right
	}
	var err1 Error
	if l1, ok := left.(Error); ok {
		err1 = append(err1, l1...)
	} else {
		err1 = append(err1, left)
	}
	if l1, ok := right.(Error); ok {
		err1 = append(err1, l1...)
	} else {
		err1 = append(err1, right)
	}
	return err1
}


