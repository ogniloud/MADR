package study

type IPalmService interface {
	GenerateTextWithWords(words []string) (string, error)
}
