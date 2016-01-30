// This file was generated by typhen-api

package battle

import (
	"app/typhenapi/core"
	"errors"
)

var _ = errors.New

// AccelerationRequestObject is a kind of TyphenAPI type.
type AccelerationRequestObject struct {
}

// Coerce the fields.
func (t *AccelerationRequestObject) Coerce() error {
	return nil
}

// Bytes creates the byte array.
func (t *AccelerationRequestObject) Bytes(serializer typhenapi.Serializer) ([]byte, error) {
	if err := t.Coerce(); err != nil {
		return nil, err
	}

	data, err := serializer.Serialize(t)
	if err != nil {
		return nil, err
	}

	return data, nil
}
