// Code generated by go-swagger; DO NOT EDIT.

package access_codes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/infobloxopen/b1ddi-go-client/models"
)

// AccessCodesUpdateAccessCodeReader is a Reader for the AccessCodesUpdateAccessCode structure.
type AccessCodesUpdateAccessCodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccessCodesUpdateAccessCodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAccessCodesUpdateAccessCodeCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAccessCodesUpdateAccessCodeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAccessCodesUpdateAccessCodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewAccessCodesUpdateAccessCodeConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAccessCodesUpdateAccessCodeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /access_codes/{payload.access_key}] access_codesUpdateAccessCode", response, response.Code())
	}
}

// NewAccessCodesUpdateAccessCodeCreated creates a AccessCodesUpdateAccessCodeCreated with default headers values
func NewAccessCodesUpdateAccessCodeCreated() *AccessCodesUpdateAccessCodeCreated {
	return &AccessCodesUpdateAccessCodeCreated{}
}

/*
AccessCodesUpdateAccessCodeCreated describes a response with status code 201, with default header values.

PUT operation response
*/
type AccessCodesUpdateAccessCodeCreated struct {
	Payload *models.AtcfwAccessCodeUpdateResponse
}

// IsSuccess returns true when this access codes update access code created response has a 2xx status code
func (o *AccessCodesUpdateAccessCodeCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this access codes update access code created response has a 3xx status code
func (o *AccessCodesUpdateAccessCodeCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this access codes update access code created response has a 4xx status code
func (o *AccessCodesUpdateAccessCodeCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this access codes update access code created response has a 5xx status code
func (o *AccessCodesUpdateAccessCodeCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this access codes update access code created response a status code equal to that given
func (o *AccessCodesUpdateAccessCodeCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the access codes update access code created response
func (o *AccessCodesUpdateAccessCodeCreated) Code() int {
	return 201
}

func (o *AccessCodesUpdateAccessCodeCreated) Error() string {
	return fmt.Sprintf("[PUT /access_codes/{payload.access_key}][%d] accessCodesUpdateAccessCodeCreated  %+v", 201, o.Payload)
}

func (o *AccessCodesUpdateAccessCodeCreated) String() string {
	return fmt.Sprintf("[PUT /access_codes/{payload.access_key}][%d] accessCodesUpdateAccessCodeCreated  %+v", 201, o.Payload)
}

func (o *AccessCodesUpdateAccessCodeCreated) GetPayload() *models.AtcfwAccessCodeUpdateResponse {
	return o.Payload
}

func (o *AccessCodesUpdateAccessCodeCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AtcfwAccessCodeUpdateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessCodesUpdateAccessCodeBadRequest creates a AccessCodesUpdateAccessCodeBadRequest with default headers values
func NewAccessCodesUpdateAccessCodeBadRequest() *AccessCodesUpdateAccessCodeBadRequest {
	return &AccessCodesUpdateAccessCodeBadRequest{}
}

/*
	AccessCodesUpdateAccessCodeBadRequest describes a response with status code 400, with default header values.

- 'name' length cannot exceed 256 characters limit
- 'description' length cannot exceed 256 characters limit
- Expiration date must be after activation date
- Cannot enter expired Bypass Code
*/
type AccessCodesUpdateAccessCodeBadRequest struct {
	Payload *AccessCodesUpdateAccessCodeBadRequestBody
}

// IsSuccess returns true when this access codes update access code bad request response has a 2xx status code
func (o *AccessCodesUpdateAccessCodeBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this access codes update access code bad request response has a 3xx status code
func (o *AccessCodesUpdateAccessCodeBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this access codes update access code bad request response has a 4xx status code
func (o *AccessCodesUpdateAccessCodeBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this access codes update access code bad request response has a 5xx status code
func (o *AccessCodesUpdateAccessCodeBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this access codes update access code bad request response a status code equal to that given
func (o *AccessCodesUpdateAccessCodeBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the access codes update access code bad request response
func (o *AccessCodesUpdateAccessCodeBadRequest) Code() int {
	return 400
}

func (o *AccessCodesUpdateAccessCodeBadRequest) Error() string {
	return fmt.Sprintf("[PUT /access_codes/{payload.access_key}][%d] accessCodesUpdateAccessCodeBadRequest  %+v", 400, o.Payload)
}

func (o *AccessCodesUpdateAccessCodeBadRequest) String() string {
	return fmt.Sprintf("[PUT /access_codes/{payload.access_key}][%d] accessCodesUpdateAccessCodeBadRequest  %+v", 400, o.Payload)
}

func (o *AccessCodesUpdateAccessCodeBadRequest) GetPayload() *AccessCodesUpdateAccessCodeBadRequestBody {
	return o.Payload
}

func (o *AccessCodesUpdateAccessCodeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AccessCodesUpdateAccessCodeBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessCodesUpdateAccessCodeNotFound creates a AccessCodesUpdateAccessCodeNotFound with default headers values
func NewAccessCodesUpdateAccessCodeNotFound() *AccessCodesUpdateAccessCodeNotFound {
	return &AccessCodesUpdateAccessCodeNotFound{}
}

/*
	AccessCodesUpdateAccessCodeNotFound describes a response with status code 404, with default header values.

- Threat Feed and TI rules must contain existing threat feeds and TI lists
- Custom Redirect rules must contain existing Custom Redirect
- Custom List rules must contain existing Custom List
*/
type AccessCodesUpdateAccessCodeNotFound struct {
	Payload *AccessCodesUpdateAccessCodeNotFoundBody
}

// IsSuccess returns true when this access codes update access code not found response has a 2xx status code
func (o *AccessCodesUpdateAccessCodeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this access codes update access code not found response has a 3xx status code
func (o *AccessCodesUpdateAccessCodeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this access codes update access code not found response has a 4xx status code
func (o *AccessCodesUpdateAccessCodeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this access codes update access code not found response has a 5xx status code
func (o *AccessCodesUpdateAccessCodeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this access codes update access code not found response a status code equal to that given
func (o *AccessCodesUpdateAccessCodeNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the access codes update access code not found response
func (o *AccessCodesUpdateAccessCodeNotFound) Code() int {
	return 404
}

func (o *AccessCodesUpdateAccessCodeNotFound) Error() string {
	return fmt.Sprintf("[PUT /access_codes/{payload.access_key}][%d] accessCodesUpdateAccessCodeNotFound  %+v", 404, o.Payload)
}

func (o *AccessCodesUpdateAccessCodeNotFound) String() string {
	return fmt.Sprintf("[PUT /access_codes/{payload.access_key}][%d] accessCodesUpdateAccessCodeNotFound  %+v", 404, o.Payload)
}

func (o *AccessCodesUpdateAccessCodeNotFound) GetPayload() *AccessCodesUpdateAccessCodeNotFoundBody {
	return o.Payload
}

func (o *AccessCodesUpdateAccessCodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AccessCodesUpdateAccessCodeNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessCodesUpdateAccessCodeConflict creates a AccessCodesUpdateAccessCodeConflict with default headers values
func NewAccessCodesUpdateAccessCodeConflict() *AccessCodesUpdateAccessCodeConflict {
	return &AccessCodesUpdateAccessCodeConflict{}
}

/*
	AccessCodesUpdateAccessCodeConflict describes a response with status code 409, with default header values.

- 'name' value must be unique among bypass codes belonging to the same account
*/
type AccessCodesUpdateAccessCodeConflict struct {
	Payload *AccessCodesUpdateAccessCodeConflictBody
}

// IsSuccess returns true when this access codes update access code conflict response has a 2xx status code
func (o *AccessCodesUpdateAccessCodeConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this access codes update access code conflict response has a 3xx status code
func (o *AccessCodesUpdateAccessCodeConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this access codes update access code conflict response has a 4xx status code
func (o *AccessCodesUpdateAccessCodeConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this access codes update access code conflict response has a 5xx status code
func (o *AccessCodesUpdateAccessCodeConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this access codes update access code conflict response a status code equal to that given
func (o *AccessCodesUpdateAccessCodeConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the access codes update access code conflict response
func (o *AccessCodesUpdateAccessCodeConflict) Code() int {
	return 409
}

func (o *AccessCodesUpdateAccessCodeConflict) Error() string {
	return fmt.Sprintf("[PUT /access_codes/{payload.access_key}][%d] accessCodesUpdateAccessCodeConflict  %+v", 409, o.Payload)
}

func (o *AccessCodesUpdateAccessCodeConflict) String() string {
	return fmt.Sprintf("[PUT /access_codes/{payload.access_key}][%d] accessCodesUpdateAccessCodeConflict  %+v", 409, o.Payload)
}

func (o *AccessCodesUpdateAccessCodeConflict) GetPayload() *AccessCodesUpdateAccessCodeConflictBody {
	return o.Payload
}

func (o *AccessCodesUpdateAccessCodeConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AccessCodesUpdateAccessCodeConflictBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessCodesUpdateAccessCodeInternalServerError creates a AccessCodesUpdateAccessCodeInternalServerError with default headers values
func NewAccessCodesUpdateAccessCodeInternalServerError() *AccessCodesUpdateAccessCodeInternalServerError {
	return &AccessCodesUpdateAccessCodeInternalServerError{}
}

/*
	AccessCodesUpdateAccessCodeInternalServerError describes a response with status code 500, with default header values.

- Internal server error occurred
*/
type AccessCodesUpdateAccessCodeInternalServerError struct {
	Payload *AccessCodesUpdateAccessCodeInternalServerErrorBody
}

// IsSuccess returns true when this access codes update access code internal server error response has a 2xx status code
func (o *AccessCodesUpdateAccessCodeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this access codes update access code internal server error response has a 3xx status code
func (o *AccessCodesUpdateAccessCodeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this access codes update access code internal server error response has a 4xx status code
func (o *AccessCodesUpdateAccessCodeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this access codes update access code internal server error response has a 5xx status code
func (o *AccessCodesUpdateAccessCodeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this access codes update access code internal server error response a status code equal to that given
func (o *AccessCodesUpdateAccessCodeInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the access codes update access code internal server error response
func (o *AccessCodesUpdateAccessCodeInternalServerError) Code() int {
	return 500
}

func (o *AccessCodesUpdateAccessCodeInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /access_codes/{payload.access_key}][%d] accessCodesUpdateAccessCodeInternalServerError  %+v", 500, o.Payload)
}

func (o *AccessCodesUpdateAccessCodeInternalServerError) String() string {
	return fmt.Sprintf("[PUT /access_codes/{payload.access_key}][%d] accessCodesUpdateAccessCodeInternalServerError  %+v", 500, o.Payload)
}

func (o *AccessCodesUpdateAccessCodeInternalServerError) GetPayload() *AccessCodesUpdateAccessCodeInternalServerErrorBody {
	return o.Payload
}

func (o *AccessCodesUpdateAccessCodeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AccessCodesUpdateAccessCodeInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
AccessCodesUpdateAccessCodeBadRequestBody access codes update access code bad request body
swagger:model AccessCodesUpdateAccessCodeBadRequestBody
*/
type AccessCodesUpdateAccessCodeBadRequestBody struct {

	// error
	Error *AccessCodesUpdateAccessCodeBadRequestBodyError `json:"error,omitempty"`
}

// Validate validates this access codes update access code bad request body
func (o *AccessCodesUpdateAccessCodeBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AccessCodesUpdateAccessCodeBadRequestBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessCodesUpdateAccessCodeBadRequest" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accessCodesUpdateAccessCodeBadRequest" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this access codes update access code bad request body based on the context it is used
func (o *AccessCodesUpdateAccessCodeBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AccessCodesUpdateAccessCodeBadRequestBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {

		if swag.IsZero(o.Error) { // not required
			return nil
		}

		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessCodesUpdateAccessCodeBadRequest" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accessCodesUpdateAccessCodeBadRequest" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AccessCodesUpdateAccessCodeBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AccessCodesUpdateAccessCodeBadRequestBody) UnmarshalBinary(b []byte) error {
	var res AccessCodesUpdateAccessCodeBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AccessCodesUpdateAccessCodeBadRequestBodyError access codes update access code bad request body error
swagger:model AccessCodesUpdateAccessCodeBadRequestBodyError
*/
type AccessCodesUpdateAccessCodeBadRequestBodyError struct {

	// code
	// Example: INVALID_ARGUMENT
	Code string `json:"code,omitempty"`

	// message
	// Example: rules' must not be empty
	Message string `json:"message,omitempty"`

	// status
	// Example: 400
	Status string `json:"status,omitempty"`
}

// Validate validates this access codes update access code bad request body error
func (o *AccessCodesUpdateAccessCodeBadRequestBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this access codes update access code bad request body error based on context it is used
func (o *AccessCodesUpdateAccessCodeBadRequestBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AccessCodesUpdateAccessCodeBadRequestBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AccessCodesUpdateAccessCodeBadRequestBodyError) UnmarshalBinary(b []byte) error {
	var res AccessCodesUpdateAccessCodeBadRequestBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AccessCodesUpdateAccessCodeConflictBody access codes update access code conflict body
swagger:model AccessCodesUpdateAccessCodeConflictBody
*/
type AccessCodesUpdateAccessCodeConflictBody struct {

	// error
	Error *AccessCodesUpdateAccessCodeConflictBodyError `json:"error,omitempty"`
}

// Validate validates this access codes update access code conflict body
func (o *AccessCodesUpdateAccessCodeConflictBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AccessCodesUpdateAccessCodeConflictBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessCodesUpdateAccessCodeConflict" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accessCodesUpdateAccessCodeConflict" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this access codes update access code conflict body based on the context it is used
func (o *AccessCodesUpdateAccessCodeConflictBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AccessCodesUpdateAccessCodeConflictBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {

		if swag.IsZero(o.Error) { // not required
			return nil
		}

		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessCodesUpdateAccessCodeConflict" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accessCodesUpdateAccessCodeConflict" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AccessCodesUpdateAccessCodeConflictBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AccessCodesUpdateAccessCodeConflictBody) UnmarshalBinary(b []byte) error {
	var res AccessCodesUpdateAccessCodeConflictBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AccessCodesUpdateAccessCodeConflictBodyError access codes update access code conflict body error
swagger:model AccessCodesUpdateAccessCodeConflictBodyError
*/
type AccessCodesUpdateAccessCodeConflictBodyError struct {

	// code
	// Example: ALREADY_EXISTS
	Code string `json:"code,omitempty"`

	// message
	// Example: Bypass Code with name 'acc_code' already exists
	Message string `json:"message,omitempty"`

	// status
	// Example: 409
	Status string `json:"status,omitempty"`
}

// Validate validates this access codes update access code conflict body error
func (o *AccessCodesUpdateAccessCodeConflictBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this access codes update access code conflict body error based on context it is used
func (o *AccessCodesUpdateAccessCodeConflictBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AccessCodesUpdateAccessCodeConflictBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AccessCodesUpdateAccessCodeConflictBodyError) UnmarshalBinary(b []byte) error {
	var res AccessCodesUpdateAccessCodeConflictBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AccessCodesUpdateAccessCodeInternalServerErrorBody access codes update access code internal server error body
swagger:model AccessCodesUpdateAccessCodeInternalServerErrorBody
*/
type AccessCodesUpdateAccessCodeInternalServerErrorBody struct {

	// error
	Error *AccessCodesUpdateAccessCodeInternalServerErrorBodyError `json:"error,omitempty"`
}

// Validate validates this access codes update access code internal server error body
func (o *AccessCodesUpdateAccessCodeInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AccessCodesUpdateAccessCodeInternalServerErrorBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessCodesUpdateAccessCodeInternalServerError" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accessCodesUpdateAccessCodeInternalServerError" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this access codes update access code internal server error body based on the context it is used
func (o *AccessCodesUpdateAccessCodeInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AccessCodesUpdateAccessCodeInternalServerErrorBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {

		if swag.IsZero(o.Error) { // not required
			return nil
		}

		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessCodesUpdateAccessCodeInternalServerError" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accessCodesUpdateAccessCodeInternalServerError" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AccessCodesUpdateAccessCodeInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AccessCodesUpdateAccessCodeInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res AccessCodesUpdateAccessCodeInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AccessCodesUpdateAccessCodeInternalServerErrorBodyError access codes update access code internal server error body error
swagger:model AccessCodesUpdateAccessCodeInternalServerErrorBodyError
*/
type AccessCodesUpdateAccessCodeInternalServerErrorBodyError struct {

	// code
	// Example: INTERNAL
	Code string `json:"code,omitempty"`

	// message
	// Example: Internal Server Error
	Message string `json:"message,omitempty"`

	// status
	// Example: 500
	Status string `json:"status,omitempty"`
}

// Validate validates this access codes update access code internal server error body error
func (o *AccessCodesUpdateAccessCodeInternalServerErrorBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this access codes update access code internal server error body error based on context it is used
func (o *AccessCodesUpdateAccessCodeInternalServerErrorBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AccessCodesUpdateAccessCodeInternalServerErrorBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AccessCodesUpdateAccessCodeInternalServerErrorBodyError) UnmarshalBinary(b []byte) error {
	var res AccessCodesUpdateAccessCodeInternalServerErrorBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AccessCodesUpdateAccessCodeNotFoundBody access codes update access code not found body
swagger:model AccessCodesUpdateAccessCodeNotFoundBody
*/
type AccessCodesUpdateAccessCodeNotFoundBody struct {

	// error
	Error *AccessCodesUpdateAccessCodeNotFoundBodyError `json:"error,omitempty"`
}

// Validate validates this access codes update access code not found body
func (o *AccessCodesUpdateAccessCodeNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AccessCodesUpdateAccessCodeNotFoundBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessCodesUpdateAccessCodeNotFound" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accessCodesUpdateAccessCodeNotFound" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this access codes update access code not found body based on the context it is used
func (o *AccessCodesUpdateAccessCodeNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AccessCodesUpdateAccessCodeNotFoundBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {

		if swag.IsZero(o.Error) { // not required
			return nil
		}

		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessCodesUpdateAccessCodeNotFound" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accessCodesUpdateAccessCodeNotFound" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AccessCodesUpdateAccessCodeNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AccessCodesUpdateAccessCodeNotFoundBody) UnmarshalBinary(b []byte) error {
	var res AccessCodesUpdateAccessCodeNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AccessCodesUpdateAccessCodeNotFoundBodyError access codes update access code not found body error
swagger:model AccessCodesUpdateAccessCodeNotFoundBodyError
*/
type AccessCodesUpdateAccessCodeNotFoundBodyError struct {

	// code
	// Example: NOT_FOUND
	Code string `json:"code,omitempty"`

	// message
	// Example: Invalid Rule: List 'custom-list-a.com' not found
	Message string `json:"message,omitempty"`

	// status
	// Example: 404
	Status string `json:"status,omitempty"`
}

// Validate validates this access codes update access code not found body error
func (o *AccessCodesUpdateAccessCodeNotFoundBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this access codes update access code not found body error based on context it is used
func (o *AccessCodesUpdateAccessCodeNotFoundBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AccessCodesUpdateAccessCodeNotFoundBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AccessCodesUpdateAccessCodeNotFoundBodyError) UnmarshalBinary(b []byte) error {
	var res AccessCodesUpdateAccessCodeNotFoundBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
