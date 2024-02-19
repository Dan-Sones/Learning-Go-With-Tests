package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const countdownStart = 3
const finalWord = "Go!"

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type DefaultSleeper struct {
}

func (D DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, sleeper Sleeper) {

	for i := countdownStart; i > 0; i-- {
		// Takes in an io.writer  and passes a string to it
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}
func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
