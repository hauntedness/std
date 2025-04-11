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

type Required[T any] struct {
	value T
}

func Require[T any](v T) Required[T] {
	return Required[T]{value: v}
}

func MustOk[T any](v T, ok bool) Required[T] {
	if !ok {
		panic("Required: must ok")
	}
	return Required[T]{value: v}
}

func Must[T any](v T, err error) Required[T] {
	if err != nil {
		panic("Required: must no error")
	}
	return Required[T]{value: v}
}

func (r Required[T]) Get() T {
	return r.value
}

// MarshalJSON encodes Option into json.
func (r Required[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.value)
}

// UnmarshalJSON decodes Option from json.
func (r *Required[T]) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte("null")) {
		return errors.New("Required[T].UnmarshalJSON: require not null")
	}
	err := json.Unmarshal(b, &r.value)
	if err != nil {
		return err
	}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface.
func (r Required[T]) MarshalText() ([]byte, error) {
	return json.Marshal(r)
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (r *Required[T]) UnmarshalText(data []byte) error {
	return json.Unmarshal(data, r)
}

// MarshalBinary is the interface implemented by an object that can marshal itself into a binary form.
func (r Required[T]) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(r.value); err != nil {
		return []byte{}, err
	}

	return append([]byte{1}, buf.Bytes()...), nil
}

// UnmarshalBinary is the interface implemented by an object that can unmarshal a binary representation of itself.
func (r *Required[T]) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return errors.New("Required[T].UnmarshalBinary: no data")
	}

	if data[0] == 0 {
		return errors.New("Required[T].UnmarshalBinary: invalid data")
	}

	buf := bytes.NewBuffer(data[1:])
	dec := gob.NewDecoder(buf)
	err := dec.Decode(&r.value)
	if err != nil {
		return err
	}

	return nil
}

// GobEncode implements the gob.GobEncoder interface.
func (r Required[T]) GobEncode() ([]byte, error) {
	return r.MarshalBinary()
}

// GobDecode implements the gob.GobDecoder interface.
func (r *Required[T]) GobDecode(data []byte) error {
	return r.UnmarshalBinary(data)
}

// Scan implements the SQL sql.Scanner interface.
func (r *Required[T]) Scan(src any) error {
	if src == nil {
		return errors.New("Required[T]: require not nil")
	}

	// is is only possible to assert interfaces, so convert first
	var t T
	if tScanner, ok := any(&t).(sql.Scanner); ok {
		if err := tScanner.Scan(src); err != nil {
			return fmt.Errorf("failed to scan: %w", err)
		}

		r.value = t
		return nil
	}

	if av, err := driver.DefaultParameterConverter.ConvertValue(src); err == nil {
		if v, ok := av.(T); ok {
			r.value = v
			return nil
		}
	}

	return r.scanConvertValue(src)
}

// Value implements the driver Valuer interface.
func (r Required[T]) Value() (driver.Value, error) {
	return driver.DefaultParameterConverter.ConvertValue(r.value)
}

func (r *Required[T]) scanConvertValue(src any) error {
	// we try to convertAssign values that we can't directly assign because ConvertValue
	// will return immediately for v that is already a Value, even if it is a different
	// Value type than the one we expect here.
	var st sql.Null[T]
	err := st.Scan(src)
	if err != nil {
		return err
	}
	r.value = st.V
	return nil
}
