// Package models describes models for application.
//
// Documentation for MADR API.
//
// Schemes: http
// BasePath: /
// Version: 0.0.1
// Contact: Artyom Blaginin<pelageech@mail.ru>
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package models

import "github.com/ogniloud/madr/internal/ioutil"

// swagger:response loadDecksOkResponse
type SwaggerLoadDecksOkResponse struct {
	// in: body
	Body LoadDecksResponse
}

// swagger:response loadDecksBadRequestError
type SwaggerLoadDecksBadRequestError struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response signUpInternalServerError
type SwaggerLoadDecksInternalServerError struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response getFlashcardsByDeckIdOkResponse
type SwaggerGetFlashcardsByDeckIdOkResponse struct {
	// in: body
	Body GetFlashcardsByDeckIdResponse
}

// swagger:response getFlashcardsByDeckIdBadRequestError
type SwaggerGetFlashcardsByDeckIdBadRequestError struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response getFlashcardsByDeckIdInternalServerError
type SwaggerGetFlashcardsByDeckIdInternalServerError struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response addFlashcardToDeckCreatedResponse
type SwaggerAddFlashcardToDeckCreatedResponse struct {
	// in: body
	Body AddFlashcardToDeckResponse
}

// swagger:response addFlashcardToDeckBadRequestError
type SwaggerAddFlashcardToDeckBadRequestError struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response addFlashcardToDeckInternalServerError
type SwaggerAddFlashcardToDeckInternalServerError struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response deleteFlashcardFromDeckNoContentResponse
type SwaggerDeleteFlashcardFromDeckNoContentResponse struct {
	// in: body
	Body DeleteFlashcardFromDeckResponse
}

// swagger:response deleteFlashcardFromDeckBadRequestError
type SwaggerDeleteFlashcardFromDeckBadRequestError struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response deleteFlashcardFromDeckInternalServerError
type SwaggerDeleteFlashcardFromDeckInternalServerError struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response newDeckWithFlashcardsCreatedResponse
type SwaggerNewDeckWithFlashcardsNoContentResponse struct {
	// in: body
	Body NewDeckWithFlashcardsResponse
}

// swagger:response newDeckWithFlashcardsBadRequestError
type SwaggerNewDeckWithFlashcardsBadRequestError struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response newDeckWithFlashcardsInternalServerError
type SwaggerNewDeckWithFlashcardsInternalServerError struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response deleteDeckNoContentResponse
type SwaggerDeleteDeckNoContentResponse struct {
	// in: body
	Body DeleteDeckResponse
}

// swagger:response deleteDeckBadRequestError
type SwaggerDeleteDeckBadRequestError struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response deleteDeckInternalServerError
type SwaggerDeleteDeckInternalServerError struct {
	// in: body
	Body ioutil.GenericError
}
