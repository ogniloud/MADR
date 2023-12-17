package study

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"slices"

	"github.com/ogniloud/madr/internal/flashcards/models"
	"github.com/ogniloud/madr/internal/flashcards/services/deck"
)

var ErrNoCards = fmt.Errorf("no cards")
var ErrCardsAmt = fmt.Errorf("wrong amount of cards returned")

const (
	Bad = models.Mark(iota)
	Satisfactory
	Excellent
)

type IStudyService interface {
	// GetNextRandom returns a random card from the entire set of user cards with expired CoolDown.
	GetNextRandom(ctx context.Context, uid models.UserId, down models.CoolDown) (models.FlashcardId, error)

	// GetNextRandomN returns at most n random card from the entire set of user cards with expired CoolDown.
	GetNextRandomN(ctx context.Context, uid models.UserId, down models.CoolDown, n int) ([]models.FlashcardId, error)

	// GetNextRandomDeck returns a random card from the user's deck with expired CoolDown.
	GetNextRandomDeck(ctx context.Context, uid models.UserId, id models.DeckId, down models.CoolDown) (models.FlashcardId, error)

	// GetNextRandomDeckN returns at most n random cards from the user's deck with expired CoolDown.
	GetNextRandomDeckN(ctx context.Context, uid models.UserId, id models.DeckId, down models.CoolDown, n int) ([]models.FlashcardId, error)

	// Rate moves the card into the box relative to the mark.
	Rate(ctx context.Context, uid models.UserId, id models.FlashcardId, mark models.Mark) error

	// MakeMatching returns a matching comprised from words from the entire set of user cards with expired CoolDown.
	MakeMatching(ctx context.Context, uid models.UserId, down models.CoolDown, size int) (models.Matching, error)

	// MakeMatchingDeck returns a matching comprised from words from the given deck with cards with expired CoolDown.
	MakeMatchingDeck(ctx context.Context, uid models.UserId, id models.DeckId, down models.CoolDown, size int) (models.Matching, error)

	// MakeText returns a text generated by a ***AI_SERVICE_NAME*** from the entire set of user cards with expired CoolDown.
	MakeText(ctx context.Context, uid models.UserId, down models.CoolDown, size int) (models.Text, error)

	// MakeTextDeck returns a text generated by a ***AI_SERVICE_NAME*** from the given deck with cards with expired CoolDown.
	MakeTextDeck(ctx context.Context, uid models.UserId, id models.DeckId, down models.CoolDown, size int) (models.Text, error)
}

type Service struct {
	// dsrv stands for deck.Service
	dsrv deck.IService

	//plmsrv stands for PalmService
	plmsrv IPalmService

	// p is a probability distribution for card selection functions,
	// the list must be equal to the maximum number of boxes the user has.
	p []float32 // for each i: 0 <= p[i] < 1
	// Example: Let p = [0, 0,5, 0.9] - three boxes: 0, 1, 2.
	// |____________<_______<__)
	// 0           0.5     0.9 1
	//
	// r = 0.6
	// Since p[1] <= r < p[2], we choose the flashcard from the box 1.
}

func NewService(s deck.IService) IStudyService {
	maxBox := 5 //todo:will be removed in 2 semester
	p := make([]float32, maxBox)
	p[0] = 0

	for i := range p[1:] {
		p[i+1] = p[i] + float32(math.Pow(.5, float64(i+1)))
	}

	return &Service{dsrv: s, p: p}
}

func (s *Service) ConvertIdsToCards(ctx context.Context, uid models.UserId, down models.CoolDown, n int, ids []models.FlashcardId) ([]models.FlashcardId, error) {
	ltns := make([]models.UserLeitner, 0, len(ids))

	for _, id := range ids {
		ltn, err := s.dsrv.GetLeitnerByUserIdCardId(ctx, uid, id)
		if err != nil {
			return nil, err
		}

		ltns = append(ltns, ltn)
	}

	ltns = slices.DeleteFunc(ltns, func(leitner models.UserLeitner) bool {
		return !leitner.CoolDown.IsPassed(down)
	})

	if len(ltns) == 0 {
		return nil, nil
	}

	// shuffle and choose next card
	rand.Shuffle(len(ltns), func(i, j int) {
		ltns[i], ltns[j] = ltns[j], ltns[i]
	})

	b := s.box()
	taken := 0
	var cards []models.FlashcardId = nil

	for i := 0; i < len(ltns) && taken < n; i++ {
		curLtn := ltns[i]
		if curLtn.Box != b {
			continue
		}
		taken++
		cards = append(cards, curLtn.FlashcardId)
	}

	for i := 0; i < len(ltns) && taken < n; i++ {
		curLtn := ltns[i]
		if curLtn.Box == b {
			continue
		}
		taken++
		cards = append(cards, curLtn.FlashcardId)
	}

	return cards, nil
}

// GetNextRandomN returns at most n random card from the entire set of user cards with expired CoolDown.
func (s *Service) GetNextRandomN(ctx context.Context, uid models.UserId, down models.CoolDown, n int) ([]models.FlashcardId, error) {
	decks, err := s.dsrv.LoadDecks(ctx, uid)
	if err != nil {
		return nil, err
	}

	keys := decks.Keys()

	var ids []models.FlashcardId
	for _, i := range keys {
		currentCards, err := s.dsrv.GetFlashcardsIdByDeckId(ctx, i)
		if err != nil {
			continue
		}
		ids = append(ids, currentCards...)
	}

	return s.ConvertIdsToCards(ctx, uid, down, n, ids)
}

// GetNextRandom returns a random card from the entire set of user cards with expired CoolDown.
func (s *Service) GetNextRandom(ctx context.Context, uid models.UserId, down models.CoolDown) (models.FlashcardId, error) {
	cards, err := s.GetNextRandomN(ctx, uid, down, 1)
	if err != nil {
		return 0, err
	}
	cardAmt := len(cards)
	if cardAmt == 0 {
		return 0, ErrNoCards
	} else if cardAmt > 1 {
		return 0, ErrCardsAmt
	}

	return cards[0], nil
}

// GetNextRandomDeckN returns at most n random cards from the user's deck with expired CoolDown.
func (s *Service) GetNextRandomDeckN(ctx context.Context, uid models.UserId, id models.DeckId, down models.CoolDown, n int) ([]models.FlashcardId, error) {
	ids, err := s.dsrv.GetFlashcardsIdByDeckId(ctx, id)
	if err != nil {
		return nil, err
	}

	return s.ConvertIdsToCards(ctx, uid, down, n, ids)
}

// GetNextRandomN returns at most n random card from the entire set of user cards with expired CoolDown.
func (s *Service) GetNextRandomDeck(ctx context.Context, uid models.UserId, id models.DeckId, down models.CoolDown) (models.FlashcardId, error) {
	cards, err := s.GetNextRandomDeckN(ctx, uid, id, down, 1)

	if err != nil {
		return 0, err
	}
	cardsAmt := len(cards)
	if cardsAmt == 0 {
		return 0, ErrNoCards
	}
	if cardsAmt > 1 {
		return 0, ErrCardsAmt
	}

	return cards[0], nil
}

// Rate moves the card into the box relative to the mark.
func (s *Service) Rate(ctx context.Context, uid models.UserId, id models.FlashcardId, mark models.Mark) error {
	l, err := s.dsrv.GetLeitnerByUserIdCardId(ctx, uid, id)
	if err != nil {
		return fmt.Errorf("leitner not found: %w", err)
	}

	box, err := s.dsrv.UserMaxBox(ctx, uid)
	if err != nil {
		return fmt.Errorf("user max box: %w", err)
	}

	switch mark {
	case Bad:
		if l.Box != 0 {
			l.Box--
		}
	case Satisfactory:
	case Excellent:
		if l.Box != box {
			l.Box++
		}
	}

	if err := s.dsrv.UpdateLeitner(ctx, l); err != nil {
		return fmt.Errorf("update leitner failed: %w", err)
	}

	return nil
}

// let p = [.5 .7 .8 .9 .95] and r is a random float
// returned Box is the minimum index i such that r >= p[i]
// so, for each i => 0 <= p[i] < 1.
func (s *Service) box() models.Box {
	r := rand.Float32()
	for i, v := range s.p {
		if r >= v && (i == len(s.p)-1 || s.p[i+1] > r) {
			return i
		}
	}

	return len(s.p)
}

// Creates a matching from the given cardIds
func (s *Service) MakeMatchingBase(ctx context.Context, cardIds []models.FlashcardId) (models.Matching, error) {
	cards := make(map[string]models.Flashcard)
	var words []string = nil
	var answers []string = nil

	for _, v := range cardIds {
		card, err := s.dsrv.GetFlashcardById(ctx, v)
		if err != nil {
			continue
		}
		_, ok := cards[card.W]
		if ok {
			continue
		}
		cards[card.W] = card
		words = append(words, card.W)
		answers = append(answers, card.A)
	}

	if len(cards) < 2 {
		return models.Matching{}, ErrCardsAmt
	}

	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	rand.Shuffle(len(answers), func(i, j int) {
		answers[i], answers[j] = answers[j], answers[i]
	})
	matching := make(map[string]string)
	for i := range words {
		matching[words[i]] = answers[i]
	}

	return models.Matching{M: matching, Cards: cards}, nil
}

// MakeMatching returns a matching comprised from words from the entire set of user cards with expired CoolDown.
func (s *Service) MakeMatching(ctx context.Context, uid models.UserId, down models.CoolDown, size int) (models.Matching, error) {
	cardIds, err := s.GetNextRandomN(ctx, uid, down, size)
	if err != nil {
		return models.Matching{}, err
	}
	return s.MakeMatchingBase(ctx, cardIds)
}

// MakeMatchingDeck returns a matching comprised from words from the given deck with cards with expired CoolDown.
func (s *Service) MakeMatchingDeck(ctx context.Context, uid models.UserId, id models.DeckId, down models.CoolDown, size int) (models.Matching, error) {
	cardIds, err := s.GetNextRandomDeckN(ctx, uid, id, down, size)
	if err != nil {
		return models.Matching{}, err
	}
	return s.MakeMatchingBase(ctx, cardIds)
}

// Creates the text from the given cardIds
func (s *Service) MakeTextBase(ctx context.Context, cardIds []models.FlashcardId) (models.Text, error) {
	cards := make(map[string]models.Flashcard)
	var words []string = nil

	for _, v := range cardIds {
		card, err := s.dsrv.GetFlashcardById(ctx, v)
		if err != nil {
			continue
		}
		_, ok := cards[card.W]
		if ok {
			continue
		}
		cards[card.W] = card
		words = append(words, card.W)
	}

	if len(cards) < 2 {
		return models.Text{}, ErrCardsAmt
	}

	text, err := s.plmsrv.GenerateTextWithWords(words)
	if err != nil {
		return models.Text{}, err
	}

	return models.Text{T: hideWords(words, text), Opts: words, Cards: cards}, nil
}

// MakeText returns a text generated by a ***AI_SERVICE_NAME*** from the entire set of user cards with expired CoolDown.
func (s *Service) MakeText(ctx context.Context, uid models.UserId, down models.CoolDown, size int) (models.Text, error) {
	cardIds, err := s.GetNextRandomN(ctx, uid, down, size)
	if err != nil {
		return models.Text{}, err
	}
	return s.MakeTextBase(ctx, cardIds)
}

// MakeTextDeck returns a text generated by a ***AI_SERVICE_NAME*** from the given deck with cards with expired CoolDown.
func (s *Service) MakeTextDeck(ctx context.Context, uid models.UserId, id models.DeckId, down models.CoolDown, size int) (models.Text, error) {
	cardIds, err := s.GetNextRandomDeckN(ctx, uid, id, down, size)
	if err != nil {
		return models.Text{}, err
	}
	return s.MakeTextBase(ctx, cardIds)
}

// Escapes words that are to be hidden from the user in the string
func hideWords(words []string, text string) string {
	result := text
	for _, word := range words {
		result = regexp.
			MustCompile("(?i)\\b"+word+"\\b").
			ReplaceAllStringFunc(result, func(match string) string {
				return "<div style=\"display:none;\">" + match + "</div>"
			})
	}
	return result
}
