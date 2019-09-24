// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/models"
)

// CreateRoleReader is a Reader for the CreateRole structure.
type CreateRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateRoleCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewCreateRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateRoleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCreateRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateRoleCreated creates a CreateRoleCreated with default headers values
func NewCreateRoleCreated() *CreateRoleCreated {
	return &CreateRoleCreated{}
}

/*CreateRoleCreated handles this case with default header values.

Role
*/
type CreateRoleCreated struct {
	Payload *models.Role
}

func (o *CreateRoleCreated) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles][%d] createRoleCreated  %+v", 201, o.Payload)
}

func (o *CreateRoleCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Role)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRoleUnauthorized creates a CreateRoleUnauthorized with default headers values
func NewCreateRoleUnauthorized() *CreateRoleUnauthorized {
	return &CreateRoleUnauthorized{}
}

/*CreateRoleUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type CreateRoleUnauthorized struct {
}

func (o *CreateRoleUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles][%d] createRoleUnauthorized ", 401)
}

func (o *CreateRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRoleForbidden creates a CreateRoleForbidden with default headers values
func NewCreateRoleForbidden() *CreateRoleForbidden {
	return &CreateRoleForbidden{}
}

/*CreateRoleForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type CreateRoleForbidden struct {
}

func (o *CreateRoleForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles][%d] createRoleForbidden ", 403)
}

func (o *CreateRoleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRoleDefault creates a CreateRoleDefault with default headers values
func NewCreateRoleDefault(code int) *CreateRoleDefault {
	return &CreateRoleDefault{
		_statusCode: code,
	}
}

/*CreateRoleDefault handles this case with default header values.

errorResponse
*/
type CreateRoleDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the create role default response
func (o *CreateRoleDefault) Code() int {
	return o._statusCode
}

func (o *CreateRoleDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles][%d] createRole default  %+v", o._statusCode, o.Payload)
}

func (o *CreateRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
