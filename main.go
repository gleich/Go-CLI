package main

import (
	"github.com/fatih/color"
	"github.com/urfave/cli"
)

var (
	red = color.New(color.FgRed)
)

func main() {
	cliApp := cli.NewApp()
	cliApp.Name = "go-cli"
	cliApp.Usage = "Simple math quizzes"
	cliApp.Version = "0.1.0"
}
