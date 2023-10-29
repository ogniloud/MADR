package studying

type Matching struct{}

type Sentence struct{}

type IExerciseService interface {
	MakeMatching() Matching
	MakeSentence() Sentence
}
