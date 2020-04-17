// Code generated by go-swagger; DO NOT EDIT.

package stage_resource

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	"strconv"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"

	"github.com/keptn/keptn/configuration-service/models"
)

// PostProjectProjectNameStageStageNameResourceHandlerFunc turns a function with the right signature into a post project project name stage stage name resource handler
type PostProjectProjectNameStageStageNameResourceHandlerFunc func(PostProjectProjectNameStageStageNameResourceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostProjectProjectNameStageStageNameResourceHandlerFunc) Handle(params PostProjectProjectNameStageStageNameResourceParams) middleware.Responder {
	return fn(params)
}

// PostProjectProjectNameStageStageNameResourceHandler interface for that can handle valid post project project name stage stage name resource params
type PostProjectProjectNameStageStageNameResourceHandler interface {
	Handle(PostProjectProjectNameStageStageNameResourceParams) middleware.Responder
}

// NewPostProjectProjectNameStageStageNameResource creates a new http.Handler for the post project project name stage stage name resource operation
func NewPostProjectProjectNameStageStageNameResource(ctx *middleware.Context, handler PostProjectProjectNameStageStageNameResourceHandler) *PostProjectProjectNameStageStageNameResource {
	return &PostProjectProjectNameStageStageNameResource{Context: ctx, Handler: handler}
}

/*PostProjectProjectNameStageStageNameResource swagger:route POST /project/{projectName}/stage/{stageName}/resource Stage Resource postProjectProjectNameStageStageNameResource

Create list of new resources for the stage

*/
type PostProjectProjectNameStageStageNameResource struct {
	Context *middleware.Context
	Handler PostProjectProjectNameStageStageNameResourceHandler
}

func (o *PostProjectProjectNameStageStageNameResource) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostProjectProjectNameStageStageNameResourceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostProjectProjectNameStageStageNameResourceBody post project project name stage stage name resource body
// swagger:model PostProjectProjectNameStageStageNameResourceBody
type PostProjectProjectNameStageStageNameResourceBody struct {

	// resources
	Resources []*models.Resource `json:"resources"`
}

// Validate validates this post project project name stage stage name resource body
func (o *PostProjectProjectNameStageStageNameResourceBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostProjectProjectNameStageStageNameResourceBody) validateResources(formats strfmt.Registry) error {

	if swag.IsZero(o.Resources) { // not required
		return nil
	}

	for i := 0; i < len(o.Resources); i++ {
		if swag.IsZero(o.Resources[i]) { // not required
			continue
		}

		if o.Resources[i] != nil {
			if err := o.Resources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resources" + "." + "resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostProjectProjectNameStageStageNameResourceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostProjectProjectNameStageStageNameResourceBody) UnmarshalBinary(b []byte) error {
	var res PostProjectProjectNameStageStageNameResourceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
