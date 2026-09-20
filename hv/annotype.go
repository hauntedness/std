package hv

// MapAny map T to any, this is useful when you need to distinguish Option[T] with other types.
// although you may only need this in rare corner cases.
// e.g. when you want to test a reflect.Value is hv.Option.
// It is hard to instantiate all possible hv.Option[T].
// you can check whether the value can implements value.(interface { MapAny() Option[any] }).
func (o Option[T]) MapAny() Option[any] {
	return From(any(o.value), o.isPresent)
}
