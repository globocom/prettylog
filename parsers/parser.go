package parsers

import "errors"

type LineParser interface {
	Parse(string) (*ParsedLine, error)
}

type ParsedLine struct {
	Timestamp string
	Logger    string
	Caller    string
	Level     string
	Message   string
	Fields    [][]string
}

var ErrNonParseableLine = errors.New("line could not be parsed")

func containsKey(key string, configKeys []string) bool {
	for _, i := range configKeys {
		if i == key {
			return true
		}
	}
	return false
}
