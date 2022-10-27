// Code generated by go-swagger; DO NOT EDIT.

package specmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// KeyValue key value
//
// swagger:model KeyValue
type KeyValue struct {

	// Ключ
	// Read Only: true
	Key string `json:"key"`

	// Значение
	// Read Only: true
	Value string `json:"value"`
}

// Validate validates this key value
func (m *KeyValue) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this key value based on the context it is used
func (m *KeyValue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KeyValue) contextValidateKey(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "key", "body", string(m.Key)); err != nil {
		return err
	}

	return nil
}

func (m *KeyValue) contextValidateValue(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "value", "body", string(m.Value)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KeyValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KeyValue) UnmarshalBinary(b []byte) error {
	var res KeyValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}