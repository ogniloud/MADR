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

	// GetNextRandomDeck returns a random card from the user's deck whose CoolDown has expired.
	GetNextRandomDeck(ctx context.Context, uid models.UserId, id models.DeckId, down models.CoolDown) (models.FlashcardId, error)

	// Rate moves the card into the box relative to the mark.
	Rate(ctx context.Context, uid models.UserId, id models.FlashcardId, mark models.Mark) error

	// MakeMatching returns a matching comprised from words from the entire set of user cards with expired CoolDown.
	MakeMatching(ctx context.Context, uid models.UserId, down models.CoolDown, size int) (models.Matching, error)

	// MakeText returns a text generated by a ***AI_SERVICE_NAME*** from the entire set of user cards with expired CoolDown.
	MakeText(ctx context.Context, generator IPalmService, uid models.UserId, down models.CoolDown, size int) (models.Text, error)
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
	maxBox := 4 //todo:will be removed in 2 semester
	p := make([]float32, maxBox)
	p[0] = 0

	for i := range p[1:] {
		p[i+1] = p[i] + float32(math.Pow(.5, float64(i+1)))
	}

	return &Service{dsrv: s, p: p}
}

func (s *Service) GetNextRandomN(ctx context.Context, uid models.UserId, down models.CoolDown, n int) []models.FlashcardId {
	decks, err := s.dsrv.LoadDecks(ctx, uid)
	if err != nil {
		return nil
	}

	keys := decks.Keys()
	rand.Shuffle(len(keys), func(i, j int) {
		keys[i], keys[j] = keys[j], keys[i]
	})

	var cards []models.FlashcardId = nil
	maxIter := n
	for i := 0; i < len(keys) && maxIter > 0; i++ {
		card, err := s.GetNextRandomDeck(ctx, uid, keys[i], down)
		if err != nil {
			continue
		}
		maxIter--
		cards = append(cards, card)
	}

	return cards
}

// GetNextRandom returns a random card from the entire set of user cards with expired CoolDown.
func (s *Service) GetNextRandom(ctx context.Context, uid models.UserId, down models.CoolDown) (models.FlashcardId, error) {
	cards := s.GetNextRandomN(ctx, uid, down, 1)
	cardAmt := len(cards)
	if cardAmt == 0 {
		return 0, ErrNoCards
	} else if cardAmt > 1 {
		return 0, ErrCardsAmt
	}

	return cards[0], nil
}

// GetNextRandomDeckN returns a random card from the user's deck whose CoolDown has expired.
func (s *Service) GetNextRandomDeckN(ctx context.Context, uid models.UserId, id models.DeckId, down models.CoolDown, n int) ([]models.FlashcardId, error) {
	ids, err := s.dsrv.GetFlashcardsIdByDeckId(ctx, id)
	if err != nil {
		return nil, err
	}

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
	maxIter := n
	var cards []models.FlashcardId = nil

	for i := 0; i < len(ltns) && maxIter > 0; i++ {
		curLtn := ltns[i]
		if curLtn.Box != b {
			continue
		}
		maxIter--
		cards = append(cards, curLtn.FlashcardId)

	}

	maxIter = n - len(cards)

	for i := 0; i < len(ltns) && i < maxIter; i++ {
		curLtn := ltns[i]
		if curLtn.Box == b {
			continue
		}
		maxIter--
		cards = append(cards, curLtn.FlashcardId)
	}

	return cards, nil
}

// GetNextRandomDeck returns a random card from the user's deck whose CoolDown has expired.
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
		return err
	}

	box, err := s.dsrv.UserMaxBox(ctx, uid)
	if err != nil {
		return err
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

	return s.dsrv.UpdateLeitner(ctx, l)
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

func (s *Service) MakeMatching(ctx context.Context, uid models.UserId, down models.CoolDown, size int) (models.Matching, error) {
	cardIds := s.GetNextRandomN(ctx, uid, down, size)
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
		matching[words[i]] = matching[answers[i]]
	}

	return models.Matching{M: matching, Cards: cards}, nil
}

func (s *Service) MakeText(ctx context.Context, _ IPalmService, uid models.UserId, down models.CoolDown, size int) (models.Text, error) {
	cardIds := s.GetNextRandomN(ctx, uid, down, size)
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
