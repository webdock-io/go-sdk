package client

import (
	"encoding/json"
	"fmt"
)

// Nullable is an optional request field that can also be sent as an explicit
// JSON null. It mirrors the Node SDK's `field?: T | null` options:
//
//   - the zero value is omitted from the request body (like undefined),
//   - NewNullable(v) sends v,
//   - Null[T]() sends null (for example to clear a callback URL).
//
// Fields of this type use the `omitzero` JSON tag option.
type Nullable[T any] struct {
	value T
	set   bool
	null  bool
}

// NewNullable returns a Nullable that sends v.
func NewNullable[T any](v T) Nullable[T] {
	return Nullable[T]{value: v, set: true}
}

// Null returns a Nullable that sends an explicit JSON null.
func Null[T any]() Nullable[T] {
	return Nullable[T]{set: true, null: true}
}

// IsZero reports whether the field is unset and should be omitted.
func (n Nullable[T]) IsZero() bool { return !n.set }

// IsNull reports whether the field is an explicit null.
func (n Nullable[T]) IsNull() bool { return n.set && n.null }

// Get returns the value and whether a non-null value is set.
func (n Nullable[T]) Get() (T, bool) { return n.value, n.set && !n.null }

func (n Nullable[T]) MarshalJSON() ([]byte, error) {
	if !n.set || n.null {
		return []byte("null"), nil
	}
	return json.Marshal(n.value)
}

func (n *Nullable[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*n = Null[T]()
		return nil
	}
	var v T
	if err := json.Unmarshal(data, &v); err != nil {
		return fmt.Errorf("decoding nullable value: %w", err)
	}
	*n = NewNullable(v)
	return nil
}
