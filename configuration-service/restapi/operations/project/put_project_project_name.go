// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PutProjectProjectNameHandlerFunc turns a function with the right signature into a put project project name handler
type PutProjectProjectNameHandlerFunc func(PutProjectProjectNameParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutProjectProjectNameHandlerFunc) Handle(params PutProjectProjectNameParams) middleware.Responder {
	return fn(params)
}

// PutProjectProjectNameHandler interface for that can handle valid put project project name params
type PutProjectProjectNameHandler interface {
	Handle(PutProjectProjectNameParams) middleware.Responder
}

// NewPutProjectProjectName creates a new http.Handler for the put project project name operation
func NewPutProjectProjectName(ctx *middleware.Context, handler PutProjectProjectNameHandler) *PutProjectProjectName {
	return &PutProjectProjectName{Context: ctx, Handler: handler}
}

/*PutProjectProjectName swagger:route PUT /project/{projectName} Project putProjectProjectName

INTERNAL Endpoint: Update the specified project

*/
type PutProjectProjectName struct {
	Context *middleware.Context
	Handler PutProjectProjectNameHandler
}

func (o *PutProjectProjectName) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPutProjectProjectNameParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
