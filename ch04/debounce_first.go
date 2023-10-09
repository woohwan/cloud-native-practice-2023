package ch04

import (
	"context"
	"sync"
	"time"
)

func DebounceFirst(circuit Circuit, d time.Duration) Circuit {
	var threshold time.Time
	var result string
	var err error
	var m sync.Mutex

	return func(ctx context.Context) (string, error) {
		m.Lock()

		defer func() {
			threshold = time.Now().Add(d)
			m.Unlock()
		}()

		// threshold가 nil이므로 처음은 무조건 에러 발생
		if time.Now().Before(threshold) {
			return result, err
		}
		result, err := circuit(ctx)

		return result, err
	}

}
