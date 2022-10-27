// Code generated by go-swagger; DO NOT EDIT.

package specmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Recommendation recommendation
//
// swagger:model Recommendation
type Recommendation struct {

	// articul
	Articul string `json:"articul,omitempty"`

	// barcode
	Barcode string `json:"barcode,omitempty"`

	// price
	Price float64 `json:"price,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this recommendation
func (m *Recommendation) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this recommendation based on context it is used
func (m *Recommendation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Recommendation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Recommendation) UnmarshalBinary(b []byte) error {
	var res Recommendation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}