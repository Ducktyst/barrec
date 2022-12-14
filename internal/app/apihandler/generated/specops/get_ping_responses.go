// Code generated by go-swagger; DO NOT EDIT.

package specops

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ducktyst/bar_recomend/internal/app/apihandler/generated/specmodels"
)

// GetPingOKCode is the HTTP code returned for type GetPingOK
const GetPingOKCode int = 200

/*
GetPingOK успешный ответ

swagger:response getPingOK
*/
type GetPingOK struct {

	/*
	  In: Body
	*/
	Payload *specmodels.Pong `json:"body,omitempty"`
}

// NewGetPingOK creates GetPingOK with default headers values
func NewGetPingOK() *GetPingOK {

	return &GetPingOK{}
}

// WithPayload adds the payload to the get ping o k response
func (o *GetPingOK) WithPayload(payload *specmodels.Pong) *GetPingOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get ping o k response
func (o *GetPingOK) SetPayload(payload *specmodels.Pong) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPingOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
