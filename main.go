package main

import (
	"log"

	"github.com/urfave/cli"

	"io/ioutil"

	"os"

	"bufio"
	"fmt"

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

	parser := &parsers.JsonLineParser{}
	prettifier := &prettifiers.DefaultPrettifier{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			fmt.Fprintln(os.Stdout, line)
			continue
		}

		parsed, err := parser.Parse(line)
		if err != nil {
			fmt.Fprintln(os.Stdout, line)
			continue
		}

		fmt.Fprintln(os.Stdout, prettifier.Prettify(parsed))
	}

	if err := scanner.Err(); err != nil {
		return cli.NewExitError("error: failed to read input: "+err.Error(), 1)
	}
	return nil
}

func isCharDevice() bool {
	fileinfo, _ := os.Stdin.Stat()
	return (fileinfo.Mode() & os.ModeCharDevice) == os.ModeCharDevice
}
