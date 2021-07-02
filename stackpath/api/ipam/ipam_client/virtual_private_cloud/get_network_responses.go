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

// GetNetworkReader is a Reader for the GetNetwork structure.
type GetNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetNetworkUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetNetworkInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetNetworkDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetworkOK creates a GetNetworkOK with default headers values
func NewGetNetworkOK() *GetNetworkOK {
	return &GetNetworkOK{}
}

/* GetNetworkOK describes a response with status code 200, with default header values.

GetNetworkOK get network o k
*/
type GetNetworkOK struct {
	Payload *ipam_models.NetworkGetNetworkResponse
}

func (o *GetNetworkOK) Error() string {
	return fmt.Sprintf("[GET /ipam/v1/stacks/{stack_id}/networks/{network_id}][%d] getNetworkOK  %+v", 200, o.Payload)
}
func (o *GetNetworkOK) GetPayload() *ipam_models.NetworkGetNetworkResponse {
	return o.Payload
}

func (o *GetNetworkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.NetworkGetNetworkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkUnauthorized creates a GetNetworkUnauthorized with default headers values
func NewGetNetworkUnauthorized() *GetNetworkUnauthorized {
	return &GetNetworkUnauthorized{}
}

/* GetNetworkUnauthorized describes a response with status code 401, with default header values.

Returned when an unauthorized request is attempted.
*/
type GetNetworkUnauthorized struct {
	Payload *ipam_models.APIStatus
}

func (o *GetNetworkUnauthorized) Error() string {
	return fmt.Sprintf("[GET /ipam/v1/stacks/{stack_id}/networks/{network_id}][%d] getNetworkUnauthorized  %+v", 401, o.Payload)
}
func (o *GetNetworkUnauthorized) GetPayload() *ipam_models.APIStatus {
	return o.Payload
}

func (o *GetNetworkUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkInternalServerError creates a GetNetworkInternalServerError with default headers values
func NewGetNetworkInternalServerError() *GetNetworkInternalServerError {
	return &GetNetworkInternalServerError{}
}

/* GetNetworkInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type GetNetworkInternalServerError struct {
	Payload *ipam_models.APIStatus
}

func (o *GetNetworkInternalServerError) Error() string {
	return fmt.Sprintf("[GET /ipam/v1/stacks/{stack_id}/networks/{network_id}][%d] getNetworkInternalServerError  %+v", 500, o.Payload)
}
func (o *GetNetworkInternalServerError) GetPayload() *ipam_models.APIStatus {
	return o.Payload
}

func (o *GetNetworkInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkDefault creates a GetNetworkDefault with default headers values
func NewGetNetworkDefault(code int) *GetNetworkDefault {
	return &GetNetworkDefault{
		_statusCode: code,
	}
}

/* GetNetworkDefault describes a response with status code -1, with default header values.

Default error structure.
*/
type GetNetworkDefault struct {
	_statusCode int

	Payload *ipam_models.APIStatus
}

// Code gets the status code for the get network default response
func (o *GetNetworkDefault) Code() int {
	return o._statusCode
}

func (o *GetNetworkDefault) Error() string {
	return fmt.Sprintf("[GET /ipam/v1/stacks/{stack_id}/networks/{network_id}][%d] GetNetwork default  %+v", o._statusCode, o.Payload)
}
func (o *GetNetworkDefault) GetPayload() *ipam_models.APIStatus {
	return o.Payload
}

func (o *GetNetworkDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
