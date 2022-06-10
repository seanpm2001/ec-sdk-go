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

// Actor actor
//
// swagger:model Actor
type Actor struct {

	// identity id
	IdentityID string `json:"identity_id,omitempty"`

	// identity type
	// Enum: [User Client]
	IdentityType string `json:"identity_type,omitempty"`

	// portal type id
	// Enum: [Customer Partner Wholesaler Uber OpenCdn]
	PortalTypeID string `json:"portal_type_id,omitempty"`

	// user id
	UserID int32 `json:"user_id,omitempty"`
}

// Validate validates this actor
func (m *Actor) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIdentityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePortalTypeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var actorTypeIdentityTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["User","Client"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		actorTypeIdentityTypePropEnum = append(actorTypeIdentityTypePropEnum, v)
	}
}

const (

	// ActorIdentityTypeUser captures enum value "User"
	ActorIdentityTypeUser string = "User"

	// ActorIdentityTypeClient captures enum value "Client"
	ActorIdentityTypeClient string = "Client"
)

// prop value enum
func (m *Actor) validateIdentityTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, actorTypeIdentityTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Actor) validateIdentityType(formats strfmt.Registry) error {
	if swag.IsZero(m.IdentityType) { // not required
		return nil
	}

	// value enum
	if err := m.validateIdentityTypeEnum("identity_type", "body", m.IdentityType); err != nil {
		return err
	}

	return nil
}

var actorTypePortalTypeIDPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Customer","Partner","Wholesaler","Uber","OpenCdn"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		actorTypePortalTypeIDPropEnum = append(actorTypePortalTypeIDPropEnum, v)
	}
}

const (

	// ActorPortalTypeIDCustomer captures enum value "Customer"
	ActorPortalTypeIDCustomer string = "Customer"

	// ActorPortalTypeIDPartner captures enum value "Partner"
	ActorPortalTypeIDPartner string = "Partner"

	// ActorPortalTypeIDWholesaler captures enum value "Wholesaler"
	ActorPortalTypeIDWholesaler string = "Wholesaler"

	// ActorPortalTypeIDUber captures enum value "Uber"
	ActorPortalTypeIDUber string = "Uber"

	// ActorPortalTypeIDOpenCdn captures enum value "OpenCdn"
	ActorPortalTypeIDOpenCdn string = "OpenCdn"
)

// prop value enum
func (m *Actor) validatePortalTypeIDEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, actorTypePortalTypeIDPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Actor) validatePortalTypeID(formats strfmt.Registry) error {
	if swag.IsZero(m.PortalTypeID) { // not required
		return nil
	}

	// value enum
	if err := m.validatePortalTypeIDEnum("portal_type_id", "body", m.PortalTypeID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this actor based on context it is used
func (m *Actor) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Actor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Actor) UnmarshalBinary(b []byte) error {
	var res Actor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}