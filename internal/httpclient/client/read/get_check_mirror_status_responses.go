// Code generated by go-swagger; DO NOT EDIT.

package read

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ory/keto/internal/httpclient/models"
)

// GetCheckMirrorStatusReader is a Reader for the GetCheckMirrorStatus structure.
type GetCheckMirrorStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCheckMirrorStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCheckMirrorStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCheckMirrorStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCheckMirrorStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCheckMirrorStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCheckMirrorStatusOK creates a GetCheckMirrorStatusOK with default headers values
func NewGetCheckMirrorStatusOK() *GetCheckMirrorStatusOK {
	return &GetCheckMirrorStatusOK{}
}

/* GetCheckMirrorStatusOK describes a response with status code 200, with default header values.

getCheckResponse
*/
type GetCheckMirrorStatusOK struct {
	Payload *models.GetCheckResponse
}

func (o *GetCheckMirrorStatusOK) Error() string {
	return fmt.Sprintf("[GET /relation-tuples/check][%d] getCheckMirrorStatusOK  %+v", 200, o.Payload)
}
func (o *GetCheckMirrorStatusOK) GetPayload() *models.GetCheckResponse {
	return o.Payload
}

func (o *GetCheckMirrorStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCheckResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCheckMirrorStatusBadRequest creates a GetCheckMirrorStatusBadRequest with default headers values
func NewGetCheckMirrorStatusBadRequest() *GetCheckMirrorStatusBadRequest {
	return &GetCheckMirrorStatusBadRequest{}
}

/* GetCheckMirrorStatusBadRequest describes a response with status code 400, with default header values.

genericError
*/
type GetCheckMirrorStatusBadRequest struct {
	Payload *models.GenericError
}

func (o *GetCheckMirrorStatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /relation-tuples/check][%d] getCheckMirrorStatusBadRequest  %+v", 400, o.Payload)
}
func (o *GetCheckMirrorStatusBadRequest) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetCheckMirrorStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCheckMirrorStatusForbidden creates a GetCheckMirrorStatusForbidden with default headers values
func NewGetCheckMirrorStatusForbidden() *GetCheckMirrorStatusForbidden {
	return &GetCheckMirrorStatusForbidden{}
}

/* GetCheckMirrorStatusForbidden describes a response with status code 403, with default header values.

getCheckResponse
*/
type GetCheckMirrorStatusForbidden struct {
	Payload *models.GetCheckResponse
}

func (o *GetCheckMirrorStatusForbidden) Error() string {
	return fmt.Sprintf("[GET /relation-tuples/check][%d] getCheckMirrorStatusForbidden  %+v", 403, o.Payload)
}
func (o *GetCheckMirrorStatusForbidden) GetPayload() *models.GetCheckResponse {
	return o.Payload
}

func (o *GetCheckMirrorStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCheckResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCheckMirrorStatusInternalServerError creates a GetCheckMirrorStatusInternalServerError with default headers values
func NewGetCheckMirrorStatusInternalServerError() *GetCheckMirrorStatusInternalServerError {
	return &GetCheckMirrorStatusInternalServerError{}
}

/* GetCheckMirrorStatusInternalServerError describes a response with status code 500, with default header values.

genericError
*/
type GetCheckMirrorStatusInternalServerError struct {
	Payload *models.GenericError
}

func (o *GetCheckMirrorStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /relation-tuples/check][%d] getCheckMirrorStatusInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCheckMirrorStatusInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetCheckMirrorStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}