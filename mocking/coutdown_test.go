package mocking

import (
	"slices"
	"testing"
)

const sleep = "sleep"
const write = "write"

type SpySleeper struct {
	Calls int
}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountdown(t *testing.T) {
	spySleepPrinter := &SpyCountdownOperations{}

	Countdown(spySleepPrinter, spySleepPrinter)

	want := []string{
		write,
		sleep,
		write,
		sleep,
		write,
		sleep,
		write,
	}

	if !slices.Equal(want, spySleepPrinter.Calls) {
		t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
	}
}
