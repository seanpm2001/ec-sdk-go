// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CustomerCommit customer commit
//
// swagger:model CustomerCommit
type CustomerCommit struct {

	// at id
	AtID string `json:"@id,omitempty"`

	// at type
	AtType string `json:"@type,omitempty"`

	// amount
	Amount float64 `json:"amount,omitempty"`
}

// Validate validates this customer commit
func (m *CustomerCommit) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this customer commit based on context it is used
func (m *CustomerCommit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustomerCommit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomerCommit) UnmarshalBinary(b []byte) error {
	var res CustomerCommit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
