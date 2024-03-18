package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// [1]
// We know we want our Countdown function to write data somewhere
// and io.Writer is the de-facto way of capturing that as an interface in Go.

// - In main we will send to os.Stdout so our users see the countdown printed to the terminal.
// - In test we will send to bytes.Buffer so our tests can capture what data is being generated.

// [01]
// func Countdown(out *bytes.Buffer) {
// 	fmt.Fprint(out, "3")
// }

// [02]

// We know that while *bytes.Buffer works, it would be better to use a general purpose interface instead.
// func Countdown(out io.Writer) {
// 	for i := 3; i > 0; i-- {
// 		// use fmt.Fprintln to print to out with our number followed by a newline character.
// 		fmt.Fprintln(out, i)
// 	}
// 	// Finally use fmt.Fprint to send "Go!" afterward.
// 	fmt.Fprint(out, "Go!")
// }

// [03]
// refactor magic values into named constants
const finalWord = "Go!"
const countdownStart = 3

// func Countdown(out io.Writer, sleeper Sleeper) {
// 	for i := countdownStart; i > 0; i-- {
// 		fmt.Fprintln(out, i)
// 		// add 1 second pauses
// 		// time.Sleep(1 * time.Second)
// 		// [07c] use the sleeper instead (the tests will now pass without waiting a second each time)
// 		sleeper.Sleep()
// 	}
// 	fmt.Fprint(out, finalWord)
// }

// Now our tests take 3 seconds to run:
// - Every forward-thinking post about software development emphasises the importance of quick feedback loops.
// - Slow tests ruin developer productivity.
// - Imagine if the requirements get more sophisticated warranting more tests. Are we happy with 3s added to the test run for every new test of Countdown?
// We have not tested an important property of our function:
// - We have a dependency on Sleeping which we need to extract so we can then control it in our tests.
// - If we can mock time.Sleep we can use dependency injection to use it instead of a "real" time.Sleep and then we can spy on the calls to make assertions on them.

// [04]
// Let's define our dependency as an interface.
// This lets us then use a real Sleeper in main and a spy sleeper in our tests.
// By using an interface our Countdown function is oblivious to this and adds some flexibility for the caller.

type Sleeper interface {
	Sleep()
}

// [05]
// I made a design decision that our Countdown function would not be responsible for how long the sleep is.
// This simplifies our code a little for now at least and means a user of our function can configure that sleepiness however they like.
// Now we need to make a mock of it for our tests to use.

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

// Spies are a kind of mock which can record how a dependency is used.
// They can record the arguments sent in, how many times it has been called, etc.
// In our case, we're keeping track of how many times Sleep() is called so we can check it in our test.

// [07]
// Let's create a real sleeper which implements the interface we need
type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// [08]
// Our latest change only asserts that it has slept 3 times, but those sleeps could occur out of sequence.
// If you run your tests they should still be passing even though the implementation is wrong.
// func Countdown(out io.Writer, sleeper Sleeper) {
// 	for i := countdownStart; i > 0; i-- {
// 		sleeper.Sleep()
// 	}

// 	for i := countdownStart; i > 0; i-- {
// 		fmt.Fprintln(out, i)
// 	}

// 	fmt.Fprint(out, finalWord)
// }

// Let's use spying again with a new test to check the order of operations is correct.
// We have two different dependencies and we want to record all of their operations into one list.
// So we'll create one spy for them both.

// Our SpyCountdownOperations implements both io.Writer and Sleeper, recording every call into one slice.
// In this test we're only concerned about the order of operations,
// so just recording them as list of named operations is sufficient.

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

// [09]
// The test should now fail. Revert Countdown back to how it was to fix the test.
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

// [11]
// A nice feature would be for the Sleeper to be configurable.
// This means that we can adjust the sleep time in our main program.
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// [12]
// We are using duration to configure the time slept and sleep as a way to pass in a sleep function.
// The signature of sleep is the same as for time.Sleep
// allowing us to use time.Sleep in our real implementation and the following spy in our tests:

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

// [13]
// All we need to do now is implement the Sleep function for ConfigurableSleeper.
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func main() {
	// [06]
	// not enough arguments in call to Countdown, have(*os.File), want(io.Writer, Sleeper)
	// Countdown(os.Stdout)

	// [07b]
	// We can then use it in our real application like so
	// sleeper := &DefaultSleeper{}

	// [14]
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}

	Countdown(os.Stdout, sleeper)
}
