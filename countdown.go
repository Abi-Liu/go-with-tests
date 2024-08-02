package main

import (
	"fmt"
	"io"
	"time"
)

type SleeperInterface interface {
	Sleep()
}

type Sleeper struct {
	duration time.Duration
}

func (s *Sleeper) Sleep() {
	time.Sleep(s.duration)
}

func countdown(w io.Writer, s SleeperInterface) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(w, i)
		s.Sleep()
	}
	fmt.Fprint(w, "Go!")
}
