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
