package main

import (
	"log"
	"strings"

	"github.com/fatih/color"
	"github.com/urfave/cli"

	"io/ioutil"

	"os"

	"github.com/globocom/prettylog/config"
	"github.com/globocom/prettylog/parsers"
	"github.com/globocom/prettylog/prettifiers"
)

func init() {
	// Disables log output so libraries won't pollute the stdout
	log.SetOutput(ioutil.Discard)
}

func main() {
	app := cli.NewApp()
	app.Name = "Prettylog"
	app.Usage = "Logs for human beings"
	app.UsageText = "some-app | prettylog"
	app.Description = "Prettylog processes JSON logs and prints them in a human-friendly format"
	app.Version = "1.1.0"
	app.HideHelp = true
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "color",
			Usage: "Colorize the output. Valid values: auto, always, never",
			Value: "auto",
		},
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "Enable verbose mode",
		},
	}

	app.Action = defaultAction
	app.Run(os.Args)
}

func defaultAction(ctx *cli.Context) error {
	if isCharDevice() {
		cli.ShowAppHelp(ctx)
		return nil
	}

	config.Load(ctx.Bool("verbose"))
	enableColorizedOutput(ctx.String("color"))

	input := &InputReader{
		Parser:     &parsers.JsonLineParser{},
		Prettifier: &prettifiers.DefaultPrettifier{},
	}
	err := input.Read(os.Stdin, os.Stdout)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	return nil
}

func isCharDevice() bool {
	fileinfo, _ := os.Stdin.Stat()
	return (fileinfo.Mode() & os.ModeCharDevice) == os.ModeCharDevice
}

func enableColorizedOutput(value string) {
	switch strings.ToLower(value) {
	case "always":
		color.NoColor = false
	case "never":
		color.NoColor = true
	}
}
