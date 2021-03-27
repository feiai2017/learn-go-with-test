package http_server

type CLI struct {
	playerStore PlayerStore
}

func (cli *CLI) PlayPoker() {
	cli.playerStore.RecordWin("Cleo")
}
