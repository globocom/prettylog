package input

import (
	"bufio"
	"errors"
	"fmt"
	"io"

	"github.com/globocom/prettylog/parsers"
	"github.com/globocom/prettylog/prettifiers"
)

type FilterFunc func(*parsers.ParsedLine) bool

type Reader struct {
	Parser     parsers.LineParser
	Prettifier prettifiers.Prettifier
	Filter     FilterFunc
}

func (r *Reader) Start(input io.Reader, output io.Writer) error {
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

		if r.Filter(parsed) {
			fmt.Fprintln(output, r.Prettifier.Prettify(parsed))
		}
	}

	if err := scanner.Err(); err != nil {
		return errors.New("error: failed to read input: " + err.Error())
	}

	return nil
}
