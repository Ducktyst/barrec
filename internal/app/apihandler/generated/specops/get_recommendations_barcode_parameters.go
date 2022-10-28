// Code generated by go-swagger; DO NOT EDIT.

package specops

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetRecommendationsBarcodeParams creates a new GetRecommendationsBarcodeParams object
//
// There are no default values defined in the spec.
func NewGetRecommendationsBarcodeParams() GetRecommendationsBarcodeParams {

	return GetRecommendationsBarcodeParams{}
}

// GetRecommendationsBarcodeParams contains all the bound params for the get recommendations barcode operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetRecommendationsBarcode
type GetRecommendationsBarcodeParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*штрихкод товара
	  Required: true
	  In: path
	*/
	Barcode string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetRecommendationsBarcodeParams() beforehand.
func (o *GetRecommendationsBarcodeParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rBarcode, rhkBarcode, _ := route.Params.GetOK("barcode")
	if err := o.bindBarcode(rBarcode, rhkBarcode, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindBarcode binds and validates parameter Barcode from path.
func (o *GetRecommendationsBarcodeParams) bindBarcode(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Barcode = raw

	return nil
}