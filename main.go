package main

import (
	"bufio"
	"os"

	"fmt"

	"flag"

	"io/ioutil"
	"log"

	"github.com/globocom/prettylog/config"
	"github.com/globocom/prettylog/prettifiers"
)

func main() {
	if isCharDevice() {
		fmt.Fprintln(os.Stderr, "Prettylog should be used to process the output of another application")
		os.Exit(1)
	}

	// Disables log output so libraries won't pollute the stdout
	log.SetOutput(ioutil.Discard)

	verbosePtr := flag.Bool("verbose", false, "turns on verbose mode")
	flag.Parse()

	config.Load(*verbosePtr)

	prettifier := prettifiers.NewJsonPrettifier()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			fmt.Println(line)
			continue
		}

		prettyline := prettifier.Prettify(line)
		fmt.Println(prettyline)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading standard error:", err)
		os.Exit(1)
	}
}

func isCharDevice() bool {
	fileinfo, _ := os.Stdin.Stat()
	return (fileinfo.Mode() & os.ModeCharDevice) == os.ModeCharDevice
}
