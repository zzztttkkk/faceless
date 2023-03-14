package result

type Result[T any] struct {
	val T
	err error
}

func WithValue[T any](val T) Result[T] {
	return Result[T]{
		val: val,
		err: nil,
	}
}

func WithError[T any](err error) Result[T] {
	return Result[T]{
		err: err,
	}
}

func (r *Result[T]) Must() T {
	if r.err != nil {
		panic(r.err)
	}
	return r.val
}

func (r *Result[T]) Error() error {
	return r.err
}

func (r *Result[T]) IsOK() bool {
	return r.err == nil
}

func (r *Result[T]) IsErr() bool {
	return r.err != nil
}

func (r *Result[T]) Or(fn func() T) T {
	if r.err != nil {
		return fn()
	}
	return r.val
}

func (r *Result[T]) OrValue(v T) T {
	if r.err != nil {
		return v
	}
	return r.val
}

func Call[S any, A any](f func(A) (S, error), a A) Result[S] {
	v, e := f(a)
	if e != nil {
		return WithError[S](e)
	}
	return WithValue(v)
}

func Call2[S any, A1 any, A2 any](f func(A1, A2) (S, error), a1 A1, a2 A2) Result[S] {
	v, e := f(a1, a2)
	if e != nil {
		return WithError[S](e)
	}
	return WithValue(v)
}

func Call3[S any, A1 any, A2 any, A3 any](f func(A1, A2, A3) (S, error), a1 A1, a2 A2, a3 A3) Result[S] {
	v, e := f(a1, a2, a3)
	if e != nil {
		return WithError[S](e)
	}
	return WithValue(v)
}

func Call4[S any, A1 any, A2 any, A3 any, A4 any](f func(A1, A2, A3, A4) (S, error), a1 A1, a2 A2, a3 A3, a4 A4) Result[S] {
	v, e := f(a1, a2, a3, a4)
	if e != nil {
		return WithError[S](e)
	}
	return WithValue(v)
}
