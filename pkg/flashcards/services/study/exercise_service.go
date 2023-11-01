package study

import "github.com/ogniloud/madr/pkg/flashcards/services/deck"

type Matching struct{}

type Sentence struct{}

type IExerciseService interface {
	MakeMatching() Matching
	MakeSentence() Sentence
}

type ExerciseService struct {
	s *deck.Service
}

func NewExerciseService(s *deck.Service) ExerciseService {
	return ExerciseService{s: s}
}

func (e ExerciseService) MakeMatching() Matching {
	//TODO implement me
	panic("implement me")
}

func (e ExerciseService) MakeSentence() Sentence {
	//TODO implement me
	panic("implement me")
}
