// Code generated by go-swagger; DO NOT EDIT.

package keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "mlss-controlcenter-go/pkg/models"
)

// AddKeyOKCode is the HTTP code returned for type AddKeyOK
const AddKeyOKCode int = 200

/*AddKeyOK keyPair interceptor.

swagger:response addKeyOK
*/
type AddKeyOK struct {

	/*
	  In: Body
	*/
	Payload *models.Result `json:"body,omitempty"`
}

// NewAddKeyOK creates AddKeyOK with default headers values
func NewAddKeyOK() *AddKeyOK {

	return &AddKeyOK{}
}

// WithPayload adds the payload to the add key o k response
func (o *AddKeyOK) WithPayload(payload *models.Result) *AddKeyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add key o k response
func (o *AddKeyOK) SetPayload(payload *models.Result) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddKeyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddKeyUnauthorizedCode is the HTTP code returned for type AddKeyUnauthorized
const AddKeyUnauthorizedCode int = 401

/*AddKeyUnauthorized Unauthorized

swagger:response addKeyUnauthorized
*/
type AddKeyUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddKeyUnauthorized creates AddKeyUnauthorized with default headers values
func NewAddKeyUnauthorized() *AddKeyUnauthorized {

	return &AddKeyUnauthorized{}
}

// WithPayload adds the payload to the add key unauthorized response
func (o *AddKeyUnauthorized) WithPayload(payload *models.Error) *AddKeyUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add key unauthorized response
func (o *AddKeyUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddKeyUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddKeyNotFoundCode is the HTTP code returned for type AddKeyNotFound
const AddKeyNotFoundCode int = 404

/*AddKeyNotFound url to add keyPair not found.

swagger:response addKeyNotFound
*/
type AddKeyNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddKeyNotFound creates AddKeyNotFound with default headers values
func NewAddKeyNotFound() *AddKeyNotFound {

	return &AddKeyNotFound{}
}

// WithPayload adds the payload to the add key not found response
func (o *AddKeyNotFound) WithPayload(payload *models.Error) *AddKeyNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add key not found response
func (o *AddKeyNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddKeyNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
