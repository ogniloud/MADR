package flashcards

import (
	"github.com/ogniloud/madr/pkg/flashcards/types"
	"time"
)

type CoolDownTime time.Time

func (cd CoolDownTime) Passed() bool {
	return time.Now().After(time.Time(cd))
}

func cooldown(l types.Level) CoolDownTime {
	return CoolDownTime(time.Now().Add(time.Duration((l*l+1)*24) * time.Hour))
}

type Translation string

func (t Translation) Answer() any {
	return t
}
