package models

import (
	"encoding/json"
	"fmt"
	"io"
)

// ToJSON serializes the given interface into a string-based JSON format.
func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)

	err := e.Encode(i)
	if err != nil {
		return fmt.Errorf("unable to encode JSON: %w", err)
	}

	return nil
}

// FromJSON deserializes the object from JSON string
// in an io.Reader to the given interface.
func FromJSON(i interface{}, r io.Reader) error {
	d := json.NewDecoder(r)

	err := d.Decode(i)
	if err != nil {
		return fmt.Errorf("unable to decode JSON: %w", err)
	}

	return nil
}
