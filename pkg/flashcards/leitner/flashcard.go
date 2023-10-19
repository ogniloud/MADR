package leitner

import (
	"github.com/ogniloud/madr/pkg/flashcards/types"
	"time"
)

// CoolDownTime is a cool down relatively real time
type CoolDownTime time.Time

// Passed if timestamp of Now is later than cool down timestamp
func (cd CoolDownTime) Passed() bool {
	return time.Now().After(time.Time(cd))
}

// returns new cool down for each level
func cooldown(l types.Level) CoolDownTime {
	return CoolDownTime(time.Now().Add(time.Duration((l*l+1)*24) * time.Hour))
}

type Translation string

func (t Translation) Answer() any {
	return t
}
