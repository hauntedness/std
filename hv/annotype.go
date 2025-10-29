package hv

// HvOption is used to distinguish Option[T] with other types.
// you may need this in rare corner cases.
// e.g. when you want to test a reflect.Value is hv.Option.
// It is hard to instantiate all possible hv.Option[T].
type HvOption interface {
	isOption()
}

var _ HvOption = Option[struct{}]{}

// Is implements IsOption.
func (o Option[T]) isOption() {}
