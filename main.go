package main

import (
	"fmt"
	"os"

	"github.com/joakimen/goji/cmd"
)

func main() {
	run(os.Args)
}

func run(args []string) {
	app := cmd.NewApp()
	err := app.Run(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
