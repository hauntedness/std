package hv

// required prevent other packages from implementing IsOption.
type required struct{}

// IsOption is used to distinguish Option[T] with other types.
// you may need this in rare corner cases.
// e.g. when you want to test a reflect.Value is hv.Option.
// It is hard to instantiate all possible hv.Option[T].
type IsOption interface {
	isOption(required)
}

var _ IsOption = Option[struct{}]{}

// Is implements IsOption.
func (o Option[T]) isOption(required) {
	// see interface IsOption.
}
