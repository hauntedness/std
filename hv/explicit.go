package hv

// boolean prevent other packages from implementing IsOption and IsRequired.
type boolean struct{}

// IsOption is used to distinguish Option[T] with other types.
// you may need this in rare corner cases.
// e.g. when you want to test a reflect.Value is hv.Option.
// It is hard to instantiate all possible hv.Option[T].
type IsOption interface {
	isOption() boolean
}

var _ IsOption = Option[struct{}]{}

// Is implements IsOption.
func (o Option[T]) isOption() boolean {
	// see interface IsOption.
	return boolean{}
}
