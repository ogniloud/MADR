package storage

import (
	"github.com/ogniloud/madr/pkg/flashcards/models"
)

//go:generate go run github.com/vektra/mockery/v2@v2.36.0 --name Storage
type Storage interface {
	// GetDecksByUserId возвращает все колоды, имеющиеся у пользователя.
	GetDecksByUserId(id models.UserId) (models.Decks, error)

	// GetFlashcardsByDeckId возращает карточки в колоде.
	GetFlashcardsByDeckId(id models.DeckId) ([]models.Flashcard, error)

	GetFlashcardById(id models.FlashcardId) (models.Flashcard, error)

	// GetCardsByUserCDBox возвращает id все карточек с истёкшим CoolDown,
	// а также удовлетворяющих пределам по количеству по Box'ам.
	//
	// Например, если limits = {1:5, 2:7}, будет возвращено не более 5 карт с Box == 1
	// и не более 7 карт с Box == 2.
	GetCardsByUserCDBox(id models.UserId, cd models.CoolDown, limits map[models.Box]int) ([]models.FlashcardId, error)

	// GetCardsByUserCDBoxDeck возвращает те же id карт, что и GetCardsByUserCDBox, но внутри колоды.
	GetCardsByUserCDBoxDeck(id models.UserId, cd models.CoolDown, limits map[models.Box]int, deckId models.DeckId) ([]models.FlashcardId, error)

	GetLeitnerByUserIdCardId(id models.UserId, flashcardId models.FlashcardId) (models.UserLeitner, error)

	GetUserInfo(uid models.UserId) (models.UserInfo, error)

	PutAllFlashcards(id models.DeckId, cards []models.Flashcard) error

	PutNewDeck(config models.DeckConfig) error

	PutAllUserLeitner(uls []models.UserLeitner) error

	DeleteFlashcardFromDeck(id models.DeckId, cardId models.FlashcardId) error

	// DeleteUserDeck удаляет запись из models.DeckConfig о том, что пользователь userId
	// имеет колоду id
	DeleteUserDeck(userId models.UserId, id models.DeckId) error

	// UpdateLeitner обновляет запись в базе данным при совпадении LeitnerId
	UpdateLeitner(ul models.UserLeitner) error
}
