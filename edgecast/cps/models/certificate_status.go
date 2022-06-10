// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CertificateStatus certificate status
//
// swagger:model CertificateStatus
type CertificateStatus struct {

	// at id
	AtID string `json:"@id,omitempty"`

	// at type
	AtType string `json:"@type,omitempty"`

	// error message
	ErrorMessage string `json:"error_message,omitempty"`

	// order validation
	OrderValidation *OrderValidation `json:"order_validation,omitempty"`

	// requires attention
	RequiresAttention bool `json:"requires_attention,omitempty"`

	// status
	// Enum: [Processing DomainControlValidation OtherValidation Deployment Active Deleted]
	Status string `json:"status,omitempty"`

	// step
	// Read Only: true
	Step int32 `json:"step,omitempty"`
}

// Validate validates this certificate status
func (m *CertificateStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrderValidation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificateStatus) validateOrderValidation(formats strfmt.Registry) error {
	if swag.IsZero(m.OrderValidation) { // not required
		return nil
	}

	if m.OrderValidation != nil {
		if err := m.OrderValidation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("order_validation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("order_validation")
			}
			return err
		}
	}

	return nil
}

var certificateStatusTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Processing","DomainControlValidation","OtherValidation","Deployment","Active","Deleted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		certificateStatusTypeStatusPropEnum = append(certificateStatusTypeStatusPropEnum, v)
	}
}

const (

	// CertificateStatusStatusProcessing captures enum value "Processing"
	CertificateStatusStatusProcessing string = "Processing"

	// CertificateStatusStatusDomainControlValidation captures enum value "DomainControlValidation"
	CertificateStatusStatusDomainControlValidation string = "DomainControlValidation"

	// CertificateStatusStatusOtherValidation captures enum value "OtherValidation"
	CertificateStatusStatusOtherValidation string = "OtherValidation"

	// CertificateStatusStatusDeployment captures enum value "Deployment"
	CertificateStatusStatusDeployment string = "Deployment"

	// CertificateStatusStatusActive captures enum value "Active"
	CertificateStatusStatusActive string = "Active"

	// CertificateStatusStatusDeleted captures enum value "Deleted"
	CertificateStatusStatusDeleted string = "Deleted"
)

// prop value enum
func (m *CertificateStatus) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, certificateStatusTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CertificateStatus) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this certificate status based on the context it is used
func (m *CertificateStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOrderValidation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStep(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificateStatus) contextValidateOrderValidation(ctx context.Context, formats strfmt.Registry) error {

	if m.OrderValidation != nil {
		if err := m.OrderValidation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("order_validation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("order_validation")
			}
			return err
		}
	}

	return nil
}

func (m *CertificateStatus) contextValidateStep(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "step", "body", int32(m.Step)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CertificateStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificateStatus) UnmarshalBinary(b []byte) error {
	var res CertificateStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
