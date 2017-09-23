package prettifiers

import "github.com/globocom/prettylog/parsers"

type Prettifier interface {
	Prettify(line *parsers.ParsedLine) string
}
