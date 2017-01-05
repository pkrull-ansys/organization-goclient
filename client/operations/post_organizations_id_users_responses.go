package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/3dsim/organization-goclient/models"
)

// PostOrganizationsIDUsersReader is a Reader for the PostOrganizationsIDUsers structure.
type PostOrganizationsIDUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOrganizationsIDUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostOrganizationsIDUsersCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostOrganizationsIDUsersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostOrganizationsIDUsersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostOrganizationsIDUsersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostOrganizationsIDUsersCreated creates a PostOrganizationsIDUsersCreated with default headers values
func NewPostOrganizationsIDUsersCreated() *PostOrganizationsIDUsersCreated {
	return &PostOrganizationsIDUsersCreated{}
}

/*PostOrganizationsIDUsersCreated handles this case with default header values.

Successfully created the user
*/
type PostOrganizationsIDUsersCreated struct {
	Payload *models.User
}

func (o *PostOrganizationsIDUsersCreated) Error() string {
	return fmt.Sprintf("[POST /organizations/{id}/users][%d] postOrganizationsIdUsersCreated  %+v", 201, o.Payload)
}

func (o *PostOrganizationsIDUsersCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrganizationsIDUsersUnauthorized creates a PostOrganizationsIDUsersUnauthorized with default headers values
func NewPostOrganizationsIDUsersUnauthorized() *PostOrganizationsIDUsersUnauthorized {
	return &PostOrganizationsIDUsersUnauthorized{}
}

/*PostOrganizationsIDUsersUnauthorized handles this case with default header values.

Not authorized
*/
type PostOrganizationsIDUsersUnauthorized struct {
	Payload *models.Error
}

func (o *PostOrganizationsIDUsersUnauthorized) Error() string {
	return fmt.Sprintf("[POST /organizations/{id}/users][%d] postOrganizationsIdUsersUnauthorized  %+v", 401, o.Payload)
}

func (o *PostOrganizationsIDUsersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrganizationsIDUsersForbidden creates a PostOrganizationsIDUsersForbidden with default headers values
func NewPostOrganizationsIDUsersForbidden() *PostOrganizationsIDUsersForbidden {
	return &PostOrganizationsIDUsersForbidden{}
}

/*PostOrganizationsIDUsersForbidden handles this case with default header values.

Forbidden
*/
type PostOrganizationsIDUsersForbidden struct {
	Payload *models.Error
}

func (o *PostOrganizationsIDUsersForbidden) Error() string {
	return fmt.Sprintf("[POST /organizations/{id}/users][%d] postOrganizationsIdUsersForbidden  %+v", 403, o.Payload)
}

func (o *PostOrganizationsIDUsersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrganizationsIDUsersDefault creates a PostOrganizationsIDUsersDefault with default headers values
func NewPostOrganizationsIDUsersDefault(code int) *PostOrganizationsIDUsersDefault {
	return &PostOrganizationsIDUsersDefault{
		_statusCode: code,
	}
}

/*PostOrganizationsIDUsersDefault handles this case with default header values.

error
*/
type PostOrganizationsIDUsersDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post organizations ID users default response
func (o *PostOrganizationsIDUsersDefault) Code() int {
	return o._statusCode
}

func (o *PostOrganizationsIDUsersDefault) Error() string {
	return fmt.Sprintf("[POST /organizations/{id}/users][%d] PostOrganizationsIDUsers default  %+v", o._statusCode, o.Payload)
}

func (o *PostOrganizationsIDUsersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
