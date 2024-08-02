package main

import (
	"os"
	"time"
)

func main() {
	s := &Sleeper{duration: 1 * time.Second}
	countdown(os.Stdout, s)
}
