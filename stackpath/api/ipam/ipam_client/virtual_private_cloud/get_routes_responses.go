// Code generated by go-swagger; DO NOT EDIT.

package virtual_private_cloud

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/stackpath/terraform-provider-stackpath/stackpath/api/ipam/ipam_models"
)

// GetRoutesReader is a Reader for the GetRoutes structure.
type GetRoutesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoutesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoutesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetRoutesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRoutesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetRoutesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRoutesOK creates a GetRoutesOK with default headers values
func NewGetRoutesOK() *GetRoutesOK {
	return &GetRoutesOK{}
}

/* GetRoutesOK describes a response with status code 200, with default header values.

GetRoutesOK get routes o k
*/
type GetRoutesOK struct {
	Payload *ipam_models.NetworkGetRoutesResponse
}

func (o *GetRoutesOK) Error() string {
	return fmt.Sprintf("[GET /ipam/v1alpha/stacks/{stack_id}/networks/{network_id}/routes][%d] getRoutesOK  %+v", 200, o.Payload)
}
func (o *GetRoutesOK) GetPayload() *ipam_models.NetworkGetRoutesResponse {
	return o.Payload
}

func (o *GetRoutesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.NetworkGetRoutesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutesUnauthorized creates a GetRoutesUnauthorized with default headers values
func NewGetRoutesUnauthorized() *GetRoutesUnauthorized {
	return &GetRoutesUnauthorized{}
}

/* GetRoutesUnauthorized describes a response with status code 401, with default header values.

Returned when an unauthorized request is attempted.
*/
type GetRoutesUnauthorized struct {
	Payload *ipam_models.APIStatus
}

func (o *GetRoutesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /ipam/v1alpha/stacks/{stack_id}/networks/{network_id}/routes][%d] getRoutesUnauthorized  %+v", 401, o.Payload)
}
func (o *GetRoutesUnauthorized) GetPayload() *ipam_models.APIStatus {
	return o.Payload
}

func (o *GetRoutesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutesInternalServerError creates a GetRoutesInternalServerError with default headers values
func NewGetRoutesInternalServerError() *GetRoutesInternalServerError {
	return &GetRoutesInternalServerError{}
}

/* GetRoutesInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type GetRoutesInternalServerError struct {
	Payload *ipam_models.APIStatus
}

func (o *GetRoutesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /ipam/v1alpha/stacks/{stack_id}/networks/{network_id}/routes][%d] getRoutesInternalServerError  %+v", 500, o.Payload)
}
func (o *GetRoutesInternalServerError) GetPayload() *ipam_models.APIStatus {
	return o.Payload
}

func (o *GetRoutesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutesDefault creates a GetRoutesDefault with default headers values
func NewGetRoutesDefault(code int) *GetRoutesDefault {
	return &GetRoutesDefault{
		_statusCode: code,
	}
}

/* GetRoutesDefault describes a response with status code -1, with default header values.

Default error structure.
*/
type GetRoutesDefault struct {
	_statusCode int

	Payload *ipam_models.APIStatus
}

// Code gets the status code for the get routes default response
func (o *GetRoutesDefault) Code() int {
	return o._statusCode
}

func (o *GetRoutesDefault) Error() string {
	return fmt.Sprintf("[GET /ipam/v1alpha/stacks/{stack_id}/networks/{network_id}/routes][%d] GetRoutes default  %+v", o._statusCode, o.Payload)
}
func (o *GetRoutesDefault) GetPayload() *ipam_models.APIStatus {
	return o.Payload
}

func (o *GetRoutesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
