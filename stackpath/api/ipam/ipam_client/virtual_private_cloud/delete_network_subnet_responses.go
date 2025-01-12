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

// DeleteNetworkSubnetReader is a Reader for the DeleteNetworkSubnet structure.
type DeleteNetworkSubnetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworkSubnetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteNetworkSubnetNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteNetworkSubnetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteNetworkSubnetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteNetworkSubnetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteNetworkSubnetNoContent creates a DeleteNetworkSubnetNoContent with default headers values
func NewDeleteNetworkSubnetNoContent() *DeleteNetworkSubnetNoContent {
	return &DeleteNetworkSubnetNoContent{}
}

/* DeleteNetworkSubnetNoContent describes a response with status code 204, with default header values.

No content
*/
type DeleteNetworkSubnetNoContent struct {
}

func (o *DeleteNetworkSubnetNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/v1alpha/stacks/{stack_id}/networks/{network_id}/subnets/{subnet_id}][%d] deleteNetworkSubnetNoContent ", 204)
}

func (o *DeleteNetworkSubnetNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNetworkSubnetUnauthorized creates a DeleteNetworkSubnetUnauthorized with default headers values
func NewDeleteNetworkSubnetUnauthorized() *DeleteNetworkSubnetUnauthorized {
	return &DeleteNetworkSubnetUnauthorized{}
}

/* DeleteNetworkSubnetUnauthorized describes a response with status code 401, with default header values.

Returned when an unauthorized request is attempted.
*/
type DeleteNetworkSubnetUnauthorized struct {
	Payload *ipam_models.APIStatus
}

func (o *DeleteNetworkSubnetUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /ipam/v1alpha/stacks/{stack_id}/networks/{network_id}/subnets/{subnet_id}][%d] deleteNetworkSubnetUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteNetworkSubnetUnauthorized) GetPayload() *ipam_models.APIStatus {
	return o.Payload
}

func (o *DeleteNetworkSubnetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNetworkSubnetInternalServerError creates a DeleteNetworkSubnetInternalServerError with default headers values
func NewDeleteNetworkSubnetInternalServerError() *DeleteNetworkSubnetInternalServerError {
	return &DeleteNetworkSubnetInternalServerError{}
}

/* DeleteNetworkSubnetInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type DeleteNetworkSubnetInternalServerError struct {
	Payload *ipam_models.APIStatus
}

func (o *DeleteNetworkSubnetInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /ipam/v1alpha/stacks/{stack_id}/networks/{network_id}/subnets/{subnet_id}][%d] deleteNetworkSubnetInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteNetworkSubnetInternalServerError) GetPayload() *ipam_models.APIStatus {
	return o.Payload
}

func (o *DeleteNetworkSubnetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNetworkSubnetDefault creates a DeleteNetworkSubnetDefault with default headers values
func NewDeleteNetworkSubnetDefault(code int) *DeleteNetworkSubnetDefault {
	return &DeleteNetworkSubnetDefault{
		_statusCode: code,
	}
}

/* DeleteNetworkSubnetDefault describes a response with status code -1, with default header values.

Default error structure.
*/
type DeleteNetworkSubnetDefault struct {
	_statusCode int

	Payload *ipam_models.APIStatus
}

// Code gets the status code for the delete network subnet default response
func (o *DeleteNetworkSubnetDefault) Code() int {
	return o._statusCode
}

func (o *DeleteNetworkSubnetDefault) Error() string {
	return fmt.Sprintf("[DELETE /ipam/v1alpha/stacks/{stack_id}/networks/{network_id}/subnets/{subnet_id}][%d] DeleteNetworkSubnet default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteNetworkSubnetDefault) GetPayload() *ipam_models.APIStatus {
	return o.Payload
}

func (o *DeleteNetworkSubnetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
