package http_server

import (
	"testing"
)

func TestCLI(t *testing.T) {
	playerStore := &StubPlayerStore{}
	cli := &CLI{playerStore: playerStore}
	cli.PlayPoker()

	if len(playerStore.winCalls) != 1 {
		t.Fatal("expected a win call but didn't get any")
	}
}
