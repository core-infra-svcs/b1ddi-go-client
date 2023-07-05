// Code generated by go-swagger; DO NOT EDIT.

package security_policy_rules

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

// SecurityPolicyRulesListSecurityPolicyRulesReader is a Reader for the SecurityPolicyRulesListSecurityPolicyRules structure.
type SecurityPolicyRulesListSecurityPolicyRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityPolicyRulesListSecurityPolicyRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityPolicyRulesListSecurityPolicyRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewSecurityPolicyRulesListSecurityPolicyRulesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /security_policy_rules] security_policy_rulesListSecurityPolicyRules", response, response.Code())
	}
}

// NewSecurityPolicyRulesListSecurityPolicyRulesOK creates a SecurityPolicyRulesListSecurityPolicyRulesOK with default headers values
func NewSecurityPolicyRulesListSecurityPolicyRulesOK() *SecurityPolicyRulesListSecurityPolicyRulesOK {
	return &SecurityPolicyRulesListSecurityPolicyRulesOK{}
}

/*
SecurityPolicyRulesListSecurityPolicyRulesOK describes a response with status code 200, with default header values.

GET operation response
*/
type SecurityPolicyRulesListSecurityPolicyRulesOK struct {
	Payload *models.AtcfwSecurityPolicyRuleMultiResponse
}

// IsSuccess returns true when this security policy rules list security policy rules o k response has a 2xx status code
func (o *SecurityPolicyRulesListSecurityPolicyRulesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security policy rules list security policy rules o k response has a 3xx status code
func (o *SecurityPolicyRulesListSecurityPolicyRulesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security policy rules list security policy rules o k response has a 4xx status code
func (o *SecurityPolicyRulesListSecurityPolicyRulesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this security policy rules list security policy rules o k response has a 5xx status code
func (o *SecurityPolicyRulesListSecurityPolicyRulesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this security policy rules list security policy rules o k response a status code equal to that given
func (o *SecurityPolicyRulesListSecurityPolicyRulesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the security policy rules list security policy rules o k response
func (o *SecurityPolicyRulesListSecurityPolicyRulesOK) Code() int {
	return 200
}

func (o *SecurityPolicyRulesListSecurityPolicyRulesOK) Error() string {
	return fmt.Sprintf("[GET /security_policy_rules][%d] securityPolicyRulesListSecurityPolicyRulesOK  %+v", 200, o.Payload)
}

func (o *SecurityPolicyRulesListSecurityPolicyRulesOK) String() string {
	return fmt.Sprintf("[GET /security_policy_rules][%d] securityPolicyRulesListSecurityPolicyRulesOK  %+v", 200, o.Payload)
}

func (o *SecurityPolicyRulesListSecurityPolicyRulesOK) GetPayload() *models.AtcfwSecurityPolicyRuleMultiResponse {
	return o.Payload
}

func (o *SecurityPolicyRulesListSecurityPolicyRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AtcfwSecurityPolicyRuleMultiResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityPolicyRulesListSecurityPolicyRulesInternalServerError creates a SecurityPolicyRulesListSecurityPolicyRulesInternalServerError with default headers values
func NewSecurityPolicyRulesListSecurityPolicyRulesInternalServerError() *SecurityPolicyRulesListSecurityPolicyRulesInternalServerError {
	return &SecurityPolicyRulesListSecurityPolicyRulesInternalServerError{}
}

/*
	SecurityPolicyRulesListSecurityPolicyRulesInternalServerError describes a response with status code 500, with default header values.

- Internal server error occurred
*/
type SecurityPolicyRulesListSecurityPolicyRulesInternalServerError struct {
	Payload *SecurityPolicyRulesListSecurityPolicyRulesInternalServerErrorBody
}

// IsSuccess returns true when this security policy rules list security policy rules internal server error response has a 2xx status code
func (o *SecurityPolicyRulesListSecurityPolicyRulesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this security policy rules list security policy rules internal server error response has a 3xx status code
func (o *SecurityPolicyRulesListSecurityPolicyRulesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security policy rules list security policy rules internal server error response has a 4xx status code
func (o *SecurityPolicyRulesListSecurityPolicyRulesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this security policy rules list security policy rules internal server error response has a 5xx status code
func (o *SecurityPolicyRulesListSecurityPolicyRulesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this security policy rules list security policy rules internal server error response a status code equal to that given
func (o *SecurityPolicyRulesListSecurityPolicyRulesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the security policy rules list security policy rules internal server error response
func (o *SecurityPolicyRulesListSecurityPolicyRulesInternalServerError) Code() int {
	return 500
}

func (o *SecurityPolicyRulesListSecurityPolicyRulesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /security_policy_rules][%d] securityPolicyRulesListSecurityPolicyRulesInternalServerError  %+v", 500, o.Payload)
}

func (o *SecurityPolicyRulesListSecurityPolicyRulesInternalServerError) String() string {
	return fmt.Sprintf("[GET /security_policy_rules][%d] securityPolicyRulesListSecurityPolicyRulesInternalServerError  %+v", 500, o.Payload)
}

func (o *SecurityPolicyRulesListSecurityPolicyRulesInternalServerError) GetPayload() *SecurityPolicyRulesListSecurityPolicyRulesInternalServerErrorBody {
	return o.Payload
}

func (o *SecurityPolicyRulesListSecurityPolicyRulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SecurityPolicyRulesListSecurityPolicyRulesInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
SecurityPolicyRulesListSecurityPolicyRulesInternalServerErrorBody security policy rules list security policy rules internal server error body
swagger:model SecurityPolicyRulesListSecurityPolicyRulesInternalServerErrorBody
*/
type SecurityPolicyRulesListSecurityPolicyRulesInternalServerErrorBody struct {

	// error
	Error *SecurityPolicyRulesListSecurityPolicyRulesInternalServerErrorBodyError `json:"error,omitempty"`
}

// Validate validates this security policy rules list security policy rules internal server error body
func (o *SecurityPolicyRulesListSecurityPolicyRulesInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SecurityPolicyRulesListSecurityPolicyRulesInternalServerErrorBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("securityPolicyRulesListSecurityPolicyRulesInternalServerError" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("securityPolicyRulesListSecurityPolicyRulesInternalServerError" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this security policy rules list security policy rules internal server error body based on the context it is used
func (o *SecurityPolicyRulesListSecurityPolicyRulesInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SecurityPolicyRulesListSecurityPolicyRulesInternalServerErrorBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {

		if swag.IsZero(o.Error) { // not required
			return nil
		}

		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("securityPolicyRulesListSecurityPolicyRulesInternalServerError" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("securityPolicyRulesListSecurityPolicyRulesInternalServerError" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SecurityPolicyRulesListSecurityPolicyRulesInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SecurityPolicyRulesListSecurityPolicyRulesInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res SecurityPolicyRulesListSecurityPolicyRulesInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SecurityPolicyRulesListSecurityPolicyRulesInternalServerErrorBodyError security policy rules list security policy rules internal server error body error
swagger:model SecurityPolicyRulesListSecurityPolicyRulesInternalServerErrorBodyError
*/
type SecurityPolicyRulesListSecurityPolicyRulesInternalServerErrorBodyError struct {

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

// Validate validates this security policy rules list security policy rules internal server error body error
func (o *SecurityPolicyRulesListSecurityPolicyRulesInternalServerErrorBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this security policy rules list security policy rules internal server error body error based on context it is used
func (o *SecurityPolicyRulesListSecurityPolicyRulesInternalServerErrorBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SecurityPolicyRulesListSecurityPolicyRulesInternalServerErrorBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SecurityPolicyRulesListSecurityPolicyRulesInternalServerErrorBodyError) UnmarshalBinary(b []byte) error {
	var res SecurityPolicyRulesListSecurityPolicyRulesInternalServerErrorBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
