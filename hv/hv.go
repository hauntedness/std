// github.com/hauntedness/std/hv is a fork of github.com/samber/mo/option.
package hv

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
)

// Some builds an Option when value is present.
func Some[T any](value T) Option[T] {
	return Option[T]{
		isPresent: true,
		value:     value,
	}
}

// None builds an Option when value is absent.
func None[T any]() Option[T] {
	return Option[T]{
		isPresent: false,
	}
}

// FromTuple builds a Some Option when second argument is true, or None.
func FromTuple[T any](value T, ok bool) Option[T] {
	if ok {
		return Some(value)
	}
	return None[T]()
}

// FromPointr builds a Some Option when value is not nil, or None.
func FromPointr[T any](value *T) Option[T] {
	if value != nil {
		return Some(*value)
	}
	return None[T]()
}

// Option is a container for an optional value of type T. If value exists, Option is
// of type Some. If the value is absent, Option is of type None.
type Option[T any] struct {
	isPresent bool
	value     T
}

// IsPresent returns false when value is absent.
func (o Option[T]) IsPresent() bool {
	return o.isPresent
}

// IsAbsent returns false when value is present.
func (o Option[T]) IsAbsent() bool {
	return !o.isPresent
}

// Get returns value and presence.
func (o Option[T]) Get() (T, bool) {
	if !o.isPresent {
		return *new(T), false
	}

	return o.value, true
}

// MustGet panic if value is not present.
func (o Option[T]) MustGet() T {
	if !o.isPresent {
		panic("value is not present")
	}
	return o.value
}

// OrElse returns value if present or default value.
func (o Option[T]) OrElse(fallback T) T {
	if !o.isPresent {
		return fallback
	}

	return o.value
}

// OrEmpty returns value if present or empty value.
func (o Option[T]) OrEmpty() T {
	return o.value
}

// ToPointr returns value if present or a nil pointer.
func (o Option[T]) ToPointr() *T {
	if !o.isPresent {
		return nil
	}

	return &o.value
}

// MarshalJSON encodes Option into json.
func (o Option[T]) MarshalJSON() ([]byte, error) {
	if o.isPresent {
		return json.Marshal(o.value)
	}

	// if anybody find a way to support `omitempty` param, please contribute!
	return json.Marshal(nil)
}

// UnmarshalJSON decodes Option from json.
func (o *Option[T]) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte("null")) {
		o.isPresent = false
		return nil
	}

	err := json.Unmarshal(b, &o.value)
	if err != nil {
		return err
	}

	o.isPresent = true
	return nil
}

// IsZero support json/v2.
func (o Option[T]) IsZero() bool {
	return o.IsAbsent()
}

// MarshalText implements the encoding.TextMarshaler interface.
func (o Option[T]) MarshalText() ([]byte, error) {
	return json.Marshal(o)
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (o *Option[T]) UnmarshalText(data []byte) error {
	return json.Unmarshal(data, o)
}

// MarshalBinary is the interface implemented by an object that can marshal itself into a binary form.
func (o Option[T]) MarshalBinary() ([]byte, error) {
	if !o.isPresent {
		return []byte{0}, nil
	}

	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(o.value); err != nil {
		return []byte{}, err
	}

	return append([]byte{1}, buf.Bytes()...), nil
}

// UnmarshalBinary is the interface implemented by an object that can unmarshal a binary representation of itself.
func (o *Option[T]) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return errors.New("Option[T].UnmarshalBinary: no data")
	}

	if data[0] == 0 {
		o.isPresent = false
		o.value = *new(T)
		return nil
	}

	buf := bytes.NewBuffer(data[1:])
	dec := gob.NewDecoder(buf)
	err := dec.Decode(&o.value)
	if err != nil {
		return err
	}

	o.isPresent = true
	return nil
}

// GobEncode implements the gob.GobEncoder interface.
func (o Option[T]) GobEncode() ([]byte, error) {
	return o.MarshalBinary()
}

// GobDecode implements the gob.GobDecoder interface.
func (o *Option[T]) GobDecode(data []byte) error {
	return o.UnmarshalBinary(data)
}

// Scan implements the SQL sql.Scanner interface.
func (o *Option[T]) Scan(src any) error {
	if src == nil {
		o.isPresent = false
		o.value = *new(T)
		return nil
	}

	// is is only possible to assert interfaces, so convert first
	var t T
	if tScanner, ok := any(&t).(sql.Scanner); ok {
		if err := tScanner.Scan(src); err != nil {
			return fmt.Errorf("failed to scan: %w", err)
		}

		o.isPresent = true
		o.value = t
		return nil
	}

	if av, err := driver.DefaultParameterConverter.ConvertValue(src); err == nil {
		if v, ok := av.(T); ok {
			o.isPresent = true
			o.value = v
			return nil
		}
	}

	return o.scanConvertValue(src)
}

// Value implements the driver Valuer interface.
func (o Option[T]) Value() (driver.Value, error) {
	if !o.isPresent {
		return driver.Value(nil), nil
	}

	return driver.DefaultParameterConverter.ConvertValue(o.value)
}

func (o *Option[T]) scanConvertValue(src any) error {
	// we try to convertAssign values that we can't directly assign because ConvertValue
	// will return immediately for v that is already a Value, even if it is a different
	// Value type than the one we expect here.
	var st sql.Null[T]
	err := st.Scan(src)
	if err != nil {
		return err
	}
	o.isPresent = true
	o.value = st.V
	return nil
}
