package goption

import (
	"encoding/json"
	"errors"
)

var ErrEmpty = errors.New("GOption is empty")

type GOption[T any] struct {
	value   T
	defined bool
}

// Some creates a new GOption with a value
func Some[T any](value T) GOption[T] {
	return GOption[T]{
		value:   value,
		defined: true,
	}
}

// None creates a new GOption with no value
func None[T any]() GOption[T] {
	return GOption[T]{
		defined: false,
	}
}

// IsDefined returns true if the GOption has a defined value
func (o GOption[T]) IsDefined() bool {
	return o.defined
}

// IsEmpty returns true if the GOption has no value
func (o GOption[T]) IsEmpty() bool {
	return !o.defined
}

// GetOrElse returns the value of the GOption if it exists, otherwise it returns the default value
func (o GOption[T]) GetOrElse(defaultValue T) T {
	if o.IsEmpty() {
		return defaultValue
	}
	return o.value
}

// Get returns the value of the GOption and boolean if it is defined
func (o GOption[T]) Get() (T, bool) {
	return o.value, o.IsDefined()
}

// MarshalJSON returns the JSON representation of the GOption
func (o GOption[T]) MarshalJSON() ([]byte, error) {
	if o.IsEmpty() {
		return []byte("null"), nil
	}
	return json.Marshal(o.value)
}

// UnmarshalJSON parses the JSON representation of the GOption
func (o *GOption[T]) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}
	var val T
	err := json.Unmarshal(data, &val)
	if err != nil {
		return err
	}
	o.value = val
	o.defined = true
	return nil
}
