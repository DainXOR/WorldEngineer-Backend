package types

type Optional[T any] struct {
	value   T
	present bool
}

func OptionalOf[T any](value T, condition ...bool) Optional[T] {

	if len(condition) > 0 {
		if condition[0] {
			return Optional[T]{value: value, present: true}
		}
		return OptionalEmpty[T]()

	} else {
		return Optional[T]{value: value, present: true}
	}
}

func OptionalEmpty[T any]() Optional[T] {
	var zeroValue T
	return Optional[T]{value: zeroValue, present: false}
}

func (o Optional[T]) IsPresent() bool {
	return o.present
}

func (o Optional[T]) Get() T {
	return o.value
}

func (o Optional[T]) GetOrElse(defaultValue T) T {
	if o.present {
		return o.value
	}
	return defaultValue
}
