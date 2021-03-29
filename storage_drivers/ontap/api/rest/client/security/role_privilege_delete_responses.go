// Code generated by go-swagger; DO NOT EDIT.

package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// RolePrivilegeDeleteReader is a Reader for the RolePrivilegeDelete structure.
type RolePrivilegeDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RolePrivilegeDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRolePrivilegeDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRolePrivilegeDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRolePrivilegeDeleteOK creates a RolePrivilegeDeleteOK with default headers values
func NewRolePrivilegeDeleteOK() *RolePrivilegeDeleteOK {
	return &RolePrivilegeDeleteOK{}
}

/* RolePrivilegeDeleteOK describes a response with status code 200, with default header values.

OK
*/
type RolePrivilegeDeleteOK struct {
}

func (o *RolePrivilegeDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /security/roles/{owner.uuid}/{name}/privileges/{path}][%d] rolePrivilegeDeleteOK ", 200)
}

func (o *RolePrivilegeDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRolePrivilegeDeleteDefault creates a RolePrivilegeDeleteDefault with default headers values
func NewRolePrivilegeDeleteDefault(code int) *RolePrivilegeDeleteDefault {
	return &RolePrivilegeDeleteDefault{
		_statusCode: code,
	}
}

/* RolePrivilegeDeleteDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 5636172 | User accounts detected with this role assigned. Update or delete those accounts before deleting this role. |
| 5636173 | This feature requires an effective cluster version of 9.6 or later. |
| 13434890 | Vserver-ID failed for Vserver roles. |
| 13434893 | The SVM does not exist. |

*/
type RolePrivilegeDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the role privilege delete default response
func (o *RolePrivilegeDeleteDefault) Code() int {
	return o._statusCode
}

func (o *RolePrivilegeDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /security/roles/{owner.uuid}/{name}/privileges/{path}][%d] role_privilege_delete default  %+v", o._statusCode, o.Payload)
}
func (o *RolePrivilegeDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *RolePrivilegeDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}