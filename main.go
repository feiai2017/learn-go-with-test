package main

import (
	"fmt"
	"github.com/Learn-go-with-tests/mocking"
	"os"
)

func main() {
	sleeper := &mocking.CountdownOperationsSpy{}
	mocking.Countdown(os.Stdout, sleeper)
	fmt.Println(sleeper)
}
