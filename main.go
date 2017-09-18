package main

import (
	"bufio"
	"os"

	"fmt"

	"github.com/globocom/prettylog/config"
	"github.com/globocom/prettylog/prettifiers"
)

func main() {
	err := config.Load()
	if err != nil {
		panic(err)
	}

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
