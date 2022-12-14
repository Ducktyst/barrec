// Code generated by go-swagger; DO NOT EDIT.

package specops

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ducktyst/bar_recomend/internal/app/apihandler/generated/specmodels"
)

// PostRecommendationsAnalyzeOKCode is the HTTP code returned for type PostRecommendationsAnalyzeOK
const PostRecommendationsAnalyzeOKCode int = 200

/*
PostRecommendationsAnalyzeOK перечисление альтернатив

swagger:response postRecommendationsAnalyzeOK
*/
type PostRecommendationsAnalyzeOK struct {

	/*
	  In: Body
	*/
	Payload []*specmodels.Recommendation `json:"body,omitempty"`
}

// NewPostRecommendationsAnalyzeOK creates PostRecommendationsAnalyzeOK with default headers values
func NewPostRecommendationsAnalyzeOK() *PostRecommendationsAnalyzeOK {

	return &PostRecommendationsAnalyzeOK{}
}

// WithPayload adds the payload to the post recommendations analyze o k response
func (o *PostRecommendationsAnalyzeOK) WithPayload(payload []*specmodels.Recommendation) *PostRecommendationsAnalyzeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post recommendations analyze o k response
func (o *PostRecommendationsAnalyzeOK) SetPayload(payload []*specmodels.Recommendation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostRecommendationsAnalyzeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*specmodels.Recommendation, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PostRecommendationsAnalyzeBadRequestCode is the HTTP code returned for type PostRecommendationsAnalyzeBadRequest
const PostRecommendationsAnalyzeBadRequestCode int = 400

/*
PostRecommendationsAnalyzeBadRequest ошибка

swagger:response postRecommendationsAnalyzeBadRequest
*/
type PostRecommendationsAnalyzeBadRequest struct {

	/*
	  In: Body
	*/
	Payload *specmodels.GenericError `json:"body,omitempty"`
}

// NewPostRecommendationsAnalyzeBadRequest creates PostRecommendationsAnalyzeBadRequest with default headers values
func NewPostRecommendationsAnalyzeBadRequest() *PostRecommendationsAnalyzeBadRequest {

	return &PostRecommendationsAnalyzeBadRequest{}
}

// WithPayload adds the payload to the post recommendations analyze bad request response
func (o *PostRecommendationsAnalyzeBadRequest) WithPayload(payload *specmodels.GenericError) *PostRecommendationsAnalyzeBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post recommendations analyze bad request response
func (o *PostRecommendationsAnalyzeBadRequest) SetPayload(payload *specmodels.GenericError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostRecommendationsAnalyzeBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
