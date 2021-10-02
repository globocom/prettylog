package main

import (
	"log"
	"strings"

	"github.com/fatih/color"
	"github.com/urfave/cli"

	"io/ioutil"

	"os"

	"github.com/globocom/prettylog/config"
	"github.com/globocom/prettylog/input"
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
	app.UsageText = "some-app | prettylog [options...]"
	app.Description = "Prettylog processes JSON logs and prints them in a human-friendly format"
	app.Version = "1.1.0"
	app.HideHelp = true
	app.Flags = []cli.Flag{
		cli.StringSliceFlag{
			Name:  "logger",
			Usage: "Output lines containing the provided loggers only",
		},
		cli.StringSliceFlag{
			Name:  "level",
			Usage: "Output lines containing the provided levels only",
		},
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
	app.Run(os.Args) // #nosec
}

func defaultAction(ctx *cli.Context) error {
	if isCharDevice() {
		cli.ShowAppHelp(ctx) // #nosec
		return nil
	}

	config.Load(ctx.Bool("verbose"))
	enableColorizedOutput(ctx.String("color"))

	reader := &input.Reader{
		Parser:     &parsers.JsonLineParser{},
		Prettifier: &prettifiers.DefaultPrettifier{},
		Filter: func(line *parsers.ParsedLine) bool {
			filtered := true
			if loggers := ctx.StringSlice("logger"); len(loggers) > 0 {
				filtered = filtered && contains(loggers, line.Logger)
			}
			if levels := ctx.StringSlice("level"); len(levels) > 0 {
				filtered = filtered && contains(levels, line.Level)
			}
			return filtered
		},
	}
	err := reader.Start(os.Stdin, os.Stdout)
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

func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
