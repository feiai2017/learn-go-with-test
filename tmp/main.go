package main

import (
	"fmt"
	"github.com/Learn-go-with-tests/mocking"
	"os"
	"time"
)

func main() {
	sleeper := &mocking.ConfigurableSleeper{Duration: 1 * time.Millisecond, Sleepa: time.Sleep}
	mocking.Countdown(os.Stdout, sleeper)
	fmt.Println(sleeper)
}
