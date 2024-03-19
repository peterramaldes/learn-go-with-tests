package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	ds := &DefaultSleeper{}
	Countdown(os.Stdout, ds)
}

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const finalWord = "Go!"
const countdownStart = 3

func Countdown(w io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
		s.Sleep()
	}
	fmt.Fprintln(w, finalWord)
}
