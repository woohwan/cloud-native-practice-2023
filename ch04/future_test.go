package ch04

import (
	"context"
	"strings"
	"testing"
	"time"
)

/*
TestFuture just runs slow functions, and makes sure that it return the
expected result after the expected amount of time
*/
func TestFuture(t *testing.T) {
	start := time.Now()

	ctx := context.Background()
	future := SlowFunction(ctx)

	res, err := future.Result()
	if err != nil {
		t.Error(err)
		return
	}

	if !strings.HasPrefix(res, "I slept for") {
		t.Error("unexpected output: ", res)
	}
	elapsedCheck(t, start, 0)
}

func elapsedCheck(t *testing.T, start time.Time, seconds int) {
	elapsed := int(time.Since(start).Seconds())

	if seconds != elapsed {
		t.Errorf("expected %d seconds; got %d\n", seconds, elapsed)
	}
}
