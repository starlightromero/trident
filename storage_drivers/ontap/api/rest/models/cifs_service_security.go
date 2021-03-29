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

// CifsServiceSecurity cifs service security
//
// swagger:model cifs_service_security
type CifsServiceSecurity struct {

	// Specifies whether AES-128 and AES-256 encryption is enabled for all Kerberos-based communication with the Active Directory KDC.
	// To take advantage of the strongest security with Kerberos-based communication, AES-256 and AES-128 encryption can be enabled on the CIFS server.
	// Kerberos-related communication for CIFS is used during CIFS server creation on the SVM, as well
	// as during the SMB session setup phase.
	// The CIFS server supports the following encryption types for Kerberos communication:
	//     * RC4-HMAC
	//     * DES
	//     * AES
	// When the CIFS server is created, the domain controller creates a computer machine account in
	// Active Directory. After a newly created machine account authenticates, the KDC and the CIFS server
	// negotiates encryption types. At this time, the KDC becomes aware of the encryption capabilities of
	// the particular machine account and uses those capabilities in subsequent communication with the
	// CIFS server.
	// In addition to negotiating encryption types during CIFS server creation, the encryption types are
	// renegotiated when a machine account password is reset.
	//
	KdcEncryption *bool `json:"kdc_encryption,omitempty"`

	// Specifies what level of access an anonymous user is granted. An anonymous user (also known as a "null user") can list or enumerate certain types of system information from Windows hosts on the network, including user names and details, account policies, and share names. Access for the anonymous user can be controlled by specifying one of three access restriction settings.
	//  The available values are:
	//  * no_restriction   - No access restriction for an anonymous user.
	//  * no_enumeration   - Enumeration is restricted for an anonymous user.
	//  * no_access        - All access is restricted for an anonymous user.
	//
	// Enum: [no_restriction no_enumeration no_access]
	RestrictAnonymous *string `json:"restrict_anonymous,omitempty"`

	// Specifies whether encryption is required for incoming CIFS traffic.
	SmbEncryption *bool `json:"smb_encryption,omitempty"`

	// Specifies whether signing is required for incoming CIFS traffic. SMB signing helps to ensure that network traffic between the CIFS server and the client is not compromised.
	//
	SmbSigning *bool `json:"smb_signing,omitempty"`
}

// Validate validates this cifs service security
func (m *CifsServiceSecurity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRestrictAnonymous(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var cifsServiceSecurityTypeRestrictAnonymousPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["no_restriction","no_enumeration","no_access"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cifsServiceSecurityTypeRestrictAnonymousPropEnum = append(cifsServiceSecurityTypeRestrictAnonymousPropEnum, v)
	}
}

const (

	// BEGIN RIPPY DEBUGGING
	// cifs_service_security
	// CifsServiceSecurity
	// restrict_anonymous
	// RestrictAnonymous
	// no_restriction
	// END RIPPY DEBUGGING
	// CifsServiceSecurityRestrictAnonymousNoRestriction captures enum value "no_restriction"
	CifsServiceSecurityRestrictAnonymousNoRestriction string = "no_restriction"

	// BEGIN RIPPY DEBUGGING
	// cifs_service_security
	// CifsServiceSecurity
	// restrict_anonymous
	// RestrictAnonymous
	// no_enumeration
	// END RIPPY DEBUGGING
	// CifsServiceSecurityRestrictAnonymousNoEnumeration captures enum value "no_enumeration"
	CifsServiceSecurityRestrictAnonymousNoEnumeration string = "no_enumeration"

	// BEGIN RIPPY DEBUGGING
	// cifs_service_security
	// CifsServiceSecurity
	// restrict_anonymous
	// RestrictAnonymous
	// no_access
	// END RIPPY DEBUGGING
	// CifsServiceSecurityRestrictAnonymousNoAccess captures enum value "no_access"
	CifsServiceSecurityRestrictAnonymousNoAccess string = "no_access"
)

// prop value enum
func (m *CifsServiceSecurity) validateRestrictAnonymousEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cifsServiceSecurityTypeRestrictAnonymousPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CifsServiceSecurity) validateRestrictAnonymous(formats strfmt.Registry) error {
	if swag.IsZero(m.RestrictAnonymous) { // not required
		return nil
	}

	// value enum
	if err := m.validateRestrictAnonymousEnum("restrict_anonymous", "body", *m.RestrictAnonymous); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cifs service security based on context it is used
func (m *CifsServiceSecurity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CifsServiceSecurity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CifsServiceSecurity) UnmarshalBinary(b []byte) error {
	var res CifsServiceSecurity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HELLO RIPPY