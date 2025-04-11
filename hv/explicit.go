package hv

// np prevent other packages from implementing IsOption and IsRequired.
type np struct{}

// IsOption is used to distinguish Option[T] with other types.
// you may need this in rare corner cases.
// e.g. when you want to test a reflect.Value is hv.Option.
// It is hard to instantiate all possible hv.Option[T].
type IsOption interface {
	isOption(np)
}

var _ IsOption = Option[struct{}]{}

// Is implements IsOption.
func (o Option[T]) isOption(np) {
	// see interface IsOption.
}

// IsRequired is used to distinguish Required[T] with other types.
// you may need this in rare corner cases.
// e.g. when you want to test a reflect.Value is hv.Required.
// It is hard to instantiate all possible hv.Required[T].
type IsRequired interface {
	isRequired(np)
}

var _ IsRequired = Required[struct{}]{}

// Is implements IsRequired.
func (o Required[T]) isRequired(np) {
	// see interface IsRequired.
}
