package ctenv

import (
	"os"
	"strconv"
	"sync"
)

var (
	slotCh   chan struct{}
	slotOnce sync.Once
)

func initConcurrency(limitFromOption int) {
	slotOnce.Do(func() {
		limit := limitFromOption
		if limit <= 0 {
			if s := os.Getenv("DBTEST_MAX_CONCURRENCY"); s != "" {
				if n, err := strconv.Atoi(s); err == nil && n > 0 {
					limit = n
				}
			}
		}
	})
}
