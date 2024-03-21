package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type SpyCountdownOperations struct {
	Calls []string
}

const sleep = "sleep"

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

const write = "write"

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("test countdown without order", func(t *testing.T) {
		b := &bytes.Buffer{}
		s := &SpySleeper{}

		Countdown(b, s)

		got := b.String()
		want := `3
2
1
Go!
`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		if s.Calls != 3 {
			t.Errorf("not enough calls to sleeper, want 3 got %d", s.Calls)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		s := &SpyCountdownOperations{}
		Countdown(s, s)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, s.Calls) {
			t.Errorf("wanted calls %v got %v", want, s.Calls)
		}
	})

}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}

}
