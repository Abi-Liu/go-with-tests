package main

import (
	"bytes"
	"testing"
)

type spySleeper struct {
	count int
}

func (s *spySleeper) Sleep() {
	s.count++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	s := &spySleeper{}
	countdown(buffer, s)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if s.count != 3 {
		t.Errorf("got %q want %q", s.count, 3)
	}
}
