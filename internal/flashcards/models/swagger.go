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
type SwaggerAddFlashcardToDeckCreatedResponse struct{}

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
type SwaggerDeleteFlashcardFromDeckNoContentResponse struct{}

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
type SwaggerNewDeckWithFlashcardsNoContentResponse struct{}

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
type SwaggerDeleteDeckNoContentResponse struct{}

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

// swagger:response randomNCardsDeckOkResponse
type SwaggerRandomNCardsDeckOkResponse struct {
	// in: body
	Body RandomNCardsDeckResponse
}

// swagger:response randomNCardsDeckBadRequestError
type SwaggerRandomNCardsDeckBadRequest struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response randomNCardsDeckInternalServerError
type SwaggerRandomNCardsDeckInternalServerError struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response randomCardOkResponse
type SwaggerRandomCardOkResponse struct {
	// in: body
	Body RandomCardResponse
}

// swagger:response randomCardBadRequestError
type SwaggerRandomCardBadRequest struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response randomCardInternalServerError
type SwaggerRandomCardInternalServerError struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response randomNCardsOkResponse
type SwaggerRandomNCardsOkResponse struct {
	// in: body
	Body RandomNCardsResponse
}

// swagger:response randomNCardsBadRequestError
type SwaggerRandomNCardsBadRequest struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response randomNCardsInternalServerError
type SwaggerRandomNCardsInternalServerError struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response randomCardDeckOkResponse
type SwaggerRandomCardDeckOkResponse struct {
	// in: body
	Body RandomCardDeckResponse
}

// swagger:response randomCardDeckBadRequestError
type SwaggerRandomCardDeckBadRequest struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response randomCardDeckInternalServerError
type SwaggerRandomCardDeckInternalServerError struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response rateNoContentResponse
type SwaggerRateNoContentResponse struct{}

// swagger:response rateBadRequestError
type SwaggerRateBadRequest struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response rateInternalServerError
type SwaggerRateInternalServerError struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response randomMatchingOkResponse
type SwaggerRandomMatchingOkResponse struct {
	// in: body
	Body RandomMatchingResponse
}

// swagger:response randomMatchingBadRequestError
type SwaggerRandomMatchingBadRequest struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response randomMatchingInternalServerError
type SwaggerRandomMatchingInternalServerError struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response randomMatchingDeckOkResponse
type SwaggerRandomMatchingDeckOkResponse struct {
	// in: body
	Body RandomMatchingDeckResponse
}

// swagger:response randomMatchingDeckBadRequestError
type SwaggerRandomMatchingDeckBadRequest struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response randomMatchingDeckInternalServerError
type SwaggerRandomMatchingDeckInternalServerError struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response randomTextOkResponse
type SwaggerRandomTextOkResponse struct {
	// in: body
	Body RandomTextResponse
}

// swagger:response randomTextBadRequestError
type SwaggerRandomTextBadRequest struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response randomTextInternalServerError
type SwaggerRandomTextInternalServerError struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response randomTextDeckOkResponse
type SwaggerRandomTextDeckOkResponse struct {
	// in: body
	Body RandomTextDeckResponse
}

// swagger:response randomTextDeckBadRequestError
type SwaggerRandomTextDeckBadRequest struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response randomTextDeckInternalServerError
type SwaggerRandomTextDeckInternalServerError struct {
	// in: body
	Body ioutil.GenericError
}
