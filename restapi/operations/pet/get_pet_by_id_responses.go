// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/senomas/xgoapi/models"
)

// GetPetByIDOKCode is the HTTP code returned for type GetPetByIDOK
const GetPetByIDOKCode int = 200

/*GetPetByIDOK successful operation

swagger:response getPetByIdOK
*/
type GetPetByIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Pet `json:"body,omitempty"`
}

// NewGetPetByIDOK creates GetPetByIDOK with default headers values
func NewGetPetByIDOK() *GetPetByIDOK {

	return &GetPetByIDOK{}
}

// WithPayload adds the payload to the get pet by Id o k response
func (o *GetPetByIDOK) WithPayload(payload *models.Pet) *GetPetByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get pet by Id o k response
func (o *GetPetByIDOK) SetPayload(payload *models.Pet) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPetByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPetByIDBadRequestCode is the HTTP code returned for type GetPetByIDBadRequest
const GetPetByIDBadRequestCode int = 400

/*GetPetByIDBadRequest Invalid ID supplied

swagger:response getPetByIdBadRequest
*/
type GetPetByIDBadRequest struct {
}

// NewGetPetByIDBadRequest creates GetPetByIDBadRequest with default headers values
func NewGetPetByIDBadRequest() *GetPetByIDBadRequest {

	return &GetPetByIDBadRequest{}
}

// WriteResponse to the client
func (o *GetPetByIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetPetByIDNotFoundCode is the HTTP code returned for type GetPetByIDNotFound
const GetPetByIDNotFoundCode int = 404

/*GetPetByIDNotFound Pet not found

swagger:response getPetByIdNotFound
*/
type GetPetByIDNotFound struct {
}

// NewGetPetByIDNotFound creates GetPetByIDNotFound with default headers values
func NewGetPetByIDNotFound() *GetPetByIDNotFound {

	return &GetPetByIDNotFound{}
}

// WriteResponse to the client
func (o *GetPetByIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
