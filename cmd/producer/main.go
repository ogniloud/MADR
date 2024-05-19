package main

import (
	"context"
	"github.com/ogniloud/madr/internal/wordmaster"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

func producer() error {
	w := &kafka.Writer{
		Addr:  kafka.TCP("localhost:9092", "localhost:9093"),
		Topic: "wiktionary",
	}
	defer w.Close()

	pos := "verb"
	msg, _ := proto.Marshal(&wordmaster.WiktionaryRequest{
		Word: &wordmaster.RequestId{
			WiktionaryLanguage: wordmaster.SupportedLanguage_RU,
			WordLanguage:       wordmaster.SupportedLanguage_EN,
			WordId: &wordmaster.WordId{
				Word:         "ejaculate",
				PartOfSpeech: &pos,
			},
		},
		Contents: &wordmaster.RequestedContents{
			Definition: true,
			Examples:   true,
			Etymology:  true,
			Ipa:        true,
		},
		Source: &wordmaster.SourceId{
			DeckId: 14,
			CardId: 88,
		},
	})
	err := w.WriteMessages(context.Background(), kafka.Message{
		Value: msg,
	})
	return err
}

func main() {
	producer()
}
