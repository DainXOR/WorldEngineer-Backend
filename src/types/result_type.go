package types

type Result[T, E any] struct {
	value Optional[T]
	err   Optional[E]
}

func ResultErr[T, E any](err E) Result[T, E] {
	return Result[T, E]{value: OptionalEmpty[T](), err: OptionalOf(err)}
}
func ResultOk[T, E any](value T) Result[T, E] {
	return Result[T, E]{value: OptionalOf(value), err: OptionalEmpty[E]()}
}
func ResultOf[T, E any](value T, err E, condition bool) Result[T, E] {
	if condition {
		return ResultOk[T, E](value)
	}
	return ResultErr[T](err)
}

func (r Result[T, E]) IsOk() bool {
	return !r.err.IsPresent()
}

func (r Result[T, E]) IsErr() bool {
	return r.err.IsPresent()
}

func (r *Result[T, E]) Value() T {
	return r.value.Get()
}
func (r *Result[T, E]) ValueOr(value T) T {
	if r.IsOk() {
		return r.Value()
	}
	return value
}

func (r *Result[T, E]) Error() E {
	return r.err.Get()
}
func (r *Result[T, E]) ErrorOr(err E) E {
	if r.IsErr() {
		return r.Error()
	}
	return err
}

func (r *Result[T, E]) Get() (T, E) {
	return r.Value(), r.Error()
}
func (r *Result[T, E]) GetRaw() (Optional[T], Optional[E]) {
	return r.value, r.err
}
