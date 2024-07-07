package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// We define a Sleeper interface that has a Sleep method.
type Sleeper interface {
	Sleep()
}

// ConfigurableSleeper is a struct that has a duration and a sleep function.
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// Sleep method for ConfigurableSleeper
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// SpyCountdownOperations is a struct that records the calls made to it.
type SpyCountdownOperations struct {
	Calls []string
}

// Sleep method for SpyCountdownOperations
func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

// Write method for SpyCountdownOperations
func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const finalWord = "Go!"
const countdownStart = 3
const write = "write"
const sleep = "sleep"

// Countdown function that takes an io.Writer and a Sleeper.
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	// Create a ConfigurableSleeper with a duration of 1 second and a sleep function of time.Sleep.
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
