package spinner

import (
	"fmt"
	"time"
)

func Party() string {
	return "🎉"
}

func Pageantry(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
