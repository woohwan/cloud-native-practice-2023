package ch04

import (
	"context"
	"fmt"
	"testing"
	"time"
)

/*
callsCountFunction returns a function that increments the int that
callCounter points to whenever it is run
*/
func callsCountFunction(callCounter *int) Effector {
	return func(ctx context.Context) (string, error) {
		*callCounter++
		return fmt.Sprintf("call %d", *callCounter), nil
	}
}

/*
TestThrottleMax1 tests whether a max of 1 call per duration is respected
*/
func TestThrottleMax1(t *testing.T) {
	const max uint = 1

	callsCounter := 0
	effector := callsCountFunction(&callsCounter)

	ctx := context.Background()
	throttle := Throttle(effector, max, max, time.Second)

	for i := 0; i < 100; i++ {
		throttle(ctx)
	}

	if callsCounter == 0 {
		t.Error("test is broken; got", callsCounter)
	}

	if callsCounter > int(max) {
		t.Error("max is broken; got", callsCounter)
	}
}

/*
TestThrottleContextTimeout tests whether a Throttle will return an error
when its context is canceled.
*/
