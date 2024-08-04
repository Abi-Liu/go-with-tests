package poker

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	store   PlayerStore
	scanner *bufio.Scanner
}

const dbFileName = "game.db.json"

func NewCli(store PlayerStore, in io.Reader) *CLI {
	scanner := bufio.NewScanner(in)

	return &CLI{
		store:   store,
		scanner: scanner,
	}
}

func (c *CLI) PlayPoker() {
	c.scanner.Scan()
	c.store.RecordWin(extractWinner(c.scanner.Text()))
}

func extractWinner(line string) string {
	return strings.Split(line, " ")[0]
}
