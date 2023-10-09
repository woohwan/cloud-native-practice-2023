package ch04

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"
)

func failAfter(threshold int) Circuit {
	count := 0

	// Service function. Fails after 5 tries.
	return func(ctx context.Context) (string, error) {
		count++
		if count > threshold {
			return "", errors.New("INTENTIONAL FAIL!")
		}
		return "Success", nil
	}
}

// TestFailAfter5 tests that the failAfter5 functions acts as expected
func TestCircuitBreakerFailAfter5(t *testing.T) {
	circuit := failAfter(5)
	ctx := context.Background()

	for count := 1; count <= 5; count++ {
		_, err := circuit(ctx)
		t.Logf("attempt %d: %v", count, err)

		switch {
		case count <= 5 && err != nil:
			t.Error("expected no error; got", err)
		case count > 5 && err == nil:
			t.Error("expected err; got none")
		}
	}
}

// TestBreaker tests that the Breaker function automatically closes and reopens.
func TestCircuitBreaker(t *testing.T) {
	// Serivce function. Fail after 5 tries
	circuit := failAfter(5)

	// A circuit breaker the opens after one failed attempt
	breaker := Breaker(circuit, 1)

	ctx := context.Background()

	circuitOpen := false
	doesCircuitOpen := false
	doesCircuitReclose := false
	count := 0

	for range time.NewTicker(time.Second).C {
		_, err := breaker(ctx)
		if err != nil {
			// Does the circuit open?
			if strings.HasPrefix(err.Error(), "service unreachable") {
				if !circuitOpen {
					circuitOpen = true
					doesCircuitOpen = true

					t.Log("circuit has opened")
				}
			} else {
				// Does it close again
				if circuitOpen {
					circuitOpen = false
					doesCircuitReclose = true

					t.Log("circuit has automatically closed")
				}
			}
		} else {
			t.Log("circuit closed and operational")
		}
		count++
		if count >= 10 {
			break
		}
	}

	if !doesCircuitOpen {
		t.Error("circuit didn't appear to open")
	}

	if !doesCircuitReclose {
		t.Error("circuit didn't appear to close  after time")
	}
}
