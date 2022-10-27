// Code generated by go-swagger; DO NOT EDIT.

package specops

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostRecommendationsHandlerFunc turns a function with the right signature into a post recommendations handler
type PostRecommendationsHandlerFunc func(PostRecommendationsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PostRecommendationsHandlerFunc) Handle(params PostRecommendationsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PostRecommendationsHandler interface for that can handle valid post recommendations params
type PostRecommendationsHandler interface {
	Handle(PostRecommendationsParams, interface{}) middleware.Responder
}

// NewPostRecommendations creates a new http.Handler for the post recommendations operation
func NewPostRecommendations(ctx *middleware.Context, handler PostRecommendationsHandler) *PostRecommendations {
	return &PostRecommendations{Context: ctx, Handler: handler}
}

/*
	PostRecommendations swagger:route POST /recommendations/ postRecommendations

Получить рекомендации по штрихкоду
*/
type PostRecommendations struct {
	Context *middleware.Context
	Handler PostRecommendationsHandler
}

func (o *PostRecommendations) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostRecommendationsParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
