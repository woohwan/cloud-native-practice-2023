package ch04

import (
	"context"
	"errors"
	"sync"
	"time"
)

type Circuit func(context.Context) (string, error)

func Breaker(circuit Circuit, failureThreshold uint) Circuit {

	var consecutiveFailure int = 0
	// Time of the last interaction with downstream servcie
	var lastAttempt = time.Now()

	var m sync.RWMutex

	// Construct and return the Circuit Closure
	return func(ctx context.Context) (string, error) {
		m.RLock()

		d := consecutiveFailure - int(failureThreshold)
		// 연속 실패 수가 기준 실패 수를 넘기고, 그 차의 승수만큼 시간이 흐를 경우 에러 번환
		/*
			예를들어, 차가 1일 경우, 마지막 성공 시간 후 2초 (2의 1승) 흐르면, 에러 발생
		*/
		if d >= 0 {
			shouldRetryAt := lastAttempt.Add(time.Second * 2 << d)
			if !time.Now().After(shouldRetryAt) {
				m.RUnlock()
				return "", errors.New("service unreachable")
			}
		}
		m.RUnlock()

		response, err := circuit(ctx)

		m.Lock() // Lock around shared resources
		defer m.Unlock()

		lastAttempt = time.Now() // Record time of attempt

		if err != nil { //Circuit returned error
			consecutiveFailure++
			return response, err
		}

		consecutiveFailure = 0 // Reset failures counter
		return response, nil
	}

}
