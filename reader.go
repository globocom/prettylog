package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"

	"github.com/globocom/prettylog/parsers"
	"github.com/globocom/prettylog/prettifiers"
)

type InputReader struct {
	Parser     parsers.LineParser
	Prettifier prettifiers.Prettifier
}

func (r *InputReader) Read(input io.Reader, output io.Writer) error {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			fmt.Fprintln(output, line)
			continue
		}

		parsed, err := r.Parser.Parse(line)
		if err != nil {
			fmt.Fprintln(output, line)
			continue
		}

		fmt.Fprintln(output, r.Prettifier.Prettify(parsed))
	}

	if err := scanner.Err(); err != nil {
		return errors.New("error: failed to read input: " + err.Error())
	}

	return nil
}
