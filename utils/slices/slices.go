package slices

func IsEmpty[T any](sv []T) bool {
	return len(sv) == 0
}

func Pop[T any](sptr *[]T) T {
	sv := *sptr
	rIdx := len(sv) - 1
	tmp := sv[rIdx]
	*sptr = sv[0:rIdx]
	return tmp
}

func PopLeft[T any](sptr *[]T) T {
	sv := *sptr
	tmp := sv[0]
	*sptr = sv[1:]
	return tmp
}
