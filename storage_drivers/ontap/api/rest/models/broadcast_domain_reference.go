// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BroadcastDomainReference Broadcast domain UUID along with a readable name. Either the UUID or both names may be provided on input.
//
// swagger:model broadcast_domain_reference
type BroadcastDomainReference struct {

	// links
	Links *BroadcastDomainReferenceLinks `json:"_links,omitempty"`

	// ipspace
	Ipspace *BroadcastDomainReferenceIpspace `json:"ipspace,omitempty"`

	// Name of the broadcast domain, scoped to its IPspace
	// Example: bd1
	Name string `json:"name,omitempty"`

	// Broadcast domain UUID
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this broadcast domain reference
func (m *BroadcastDomainReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIpspace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BroadcastDomainReference) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *BroadcastDomainReference) validateIpspace(formats strfmt.Registry) error {
	if swag.IsZero(m.Ipspace) { // not required
		return nil
	}

	if m.Ipspace != nil {
		if err := m.Ipspace.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipspace")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this broadcast domain reference based on the context it is used
func (m *BroadcastDomainReference) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIpspace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BroadcastDomainReference) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *BroadcastDomainReference) contextValidateIpspace(ctx context.Context, formats strfmt.Registry) error {

	if m.Ipspace != nil {
		if err := m.Ipspace.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipspace")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BroadcastDomainReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BroadcastDomainReference) UnmarshalBinary(b []byte) error {
	var res BroadcastDomainReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BroadcastDomainReferenceIpspace broadcast domain reference ipspace
//
// swagger:model BroadcastDomainReferenceIpspace
type BroadcastDomainReferenceIpspace struct {

	// Name of the broadcast domain's IPspace
	// Example: ipspace1
	Name string `json:"name,omitempty"`
}

// Validate validates this broadcast domain reference ipspace
func (m *BroadcastDomainReferenceIpspace) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this broadcast domain reference ipspace based on context it is used
func (m *BroadcastDomainReferenceIpspace) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BroadcastDomainReferenceIpspace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BroadcastDomainReferenceIpspace) UnmarshalBinary(b []byte) error {
	var res BroadcastDomainReferenceIpspace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BroadcastDomainReferenceLinks broadcast domain reference links
//
// swagger:model BroadcastDomainReferenceLinks
type BroadcastDomainReferenceLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this broadcast domain reference links
func (m *BroadcastDomainReferenceLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BroadcastDomainReferenceLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this broadcast domain reference links based on the context it is used
func (m *BroadcastDomainReferenceLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BroadcastDomainReferenceLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BroadcastDomainReferenceLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BroadcastDomainReferenceLinks) UnmarshalBinary(b []byte) error {
	var res BroadcastDomainReferenceLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HELLO RIPPY