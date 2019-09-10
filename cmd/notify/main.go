package main

import (
	"os"

	"github.com/KeisukeYamashita/notifyer/cli"
)

func main() {
	c := cli.NewCLI(os.Stdout, os.Stderr, os.Stdin)
	os.Exit(c.Execute(os.Args[1:]))
}
