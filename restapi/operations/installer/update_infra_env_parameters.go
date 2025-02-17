// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"

	"github.com/openshift/assisted-service/models"
)

// NewUpdateInfraEnvParams creates a new UpdateInfraEnvParams object
//
// There are no default values defined in the spec.
func NewUpdateInfraEnvParams() UpdateInfraEnvParams {

	return UpdateInfraEnvParams{}
}

// UpdateInfraEnvParams contains all the bound params for the update infra env operation
// typically these are obtained from a http.Request
//
// swagger:parameters UpdateInfraEnv
type UpdateInfraEnvParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The properties to update.
	  Required: true
	  In: body
	*/
	InfraEnvUpdateParams *models.InfraEnvUpdateParams
	/*The infra-env to be updated.
	  Required: true
	  In: path
	*/
	InfraEnvID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateInfraEnvParams() beforehand.
func (o *UpdateInfraEnvParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.InfraEnvUpdateParams
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("infraEnvUpdateParams", "body", ""))
			} else {
				res = append(res, errors.NewParseError("infraEnvUpdateParams", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(context.Background())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.InfraEnvUpdateParams = &body
			}
		}
	} else {
		res = append(res, errors.Required("infraEnvUpdateParams", "body", ""))
	}

	rInfraEnvID, rhkInfraEnvID, _ := route.Params.GetOK("infra_env_id")
	if err := o.bindInfraEnvID(rInfraEnvID, rhkInfraEnvID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindInfraEnvID binds and validates parameter InfraEnvID from path.
func (o *UpdateInfraEnvParams) bindInfraEnvID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("infra_env_id", "path", "strfmt.UUID", raw)
	}
	o.InfraEnvID = *(value.(*strfmt.UUID))

	if err := o.validateInfraEnvID(formats); err != nil {
		return err
	}

	return nil
}

// validateInfraEnvID carries on validations for parameter InfraEnvID
func (o *UpdateInfraEnvParams) validateInfraEnvID(formats strfmt.Registry) error {

	if err := validate.FormatOf("infra_env_id", "path", "uuid", o.InfraEnvID.String(), formats); err != nil {
		return err
	}
	return nil
}
