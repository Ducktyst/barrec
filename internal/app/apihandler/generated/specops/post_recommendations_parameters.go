// Code generated by go-swagger; DO NOT EDIT.

package specops

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"mime/multipart"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

// PostRecommendationsMaxParseMemory sets the maximum size in bytes for
// the multipart form parser for this operation.
//
// The default value is 32 MB.
// The multipart parser stores up to this + 10MB.
var PostRecommendationsMaxParseMemory int64 = 32 << 20

// NewPostRecommendationsParams creates a new PostRecommendationsParams object
//
// There are no default values defined in the spec.
func NewPostRecommendationsParams() PostRecommendationsParams {

	return PostRecommendationsParams{}
}

// PostRecommendationsParams contains all the bound params for the post recommendations operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostRecommendations
type PostRecommendationsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Изображение со штрикодом
	  Required: true
	  In: formData
	*/
	Content io.ReadCloser
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostRecommendationsParams() beforehand.
func (o *PostRecommendationsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := r.ParseMultipartForm(PostRecommendationsMaxParseMemory); err != nil {
		if err != http.ErrNotMultipart {
			return errors.New(400, "%v", err)
		} else if err := r.ParseForm(); err != nil {
			return errors.New(400, "%v", err)
		}
	}

	content, contentHeader, err := r.FormFile("content")
	if err != nil {
		res = append(res, errors.New(400, "reading file %q failed: %v", "content", err))
	} else if err := o.bindContent(content, contentHeader); err != nil {
		// Required: true
		res = append(res, err)
	} else {
		o.Content = &runtime.File{Data: content, Header: contentHeader}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindContent binds file parameter Content.
//
// The only supported validations on files are MinLength and MaxLength
func (o *PostRecommendationsParams) bindContent(file multipart.File, header *multipart.FileHeader) error {
	return nil
}
