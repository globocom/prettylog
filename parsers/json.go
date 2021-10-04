package parsers

import (
	"github.com/globocom/prettylog/config"
	"github.com/tidwall/gjson"
)

type JsonLineParser struct{}

func (*JsonLineParser) Parse(line string) (*ParsedLine, error) {
	if !gjson.Valid(line) {
		return nil, ErrNonParseableLine
	}

	settings := config.GetSettings()
	parsed := &ParsedLine{}

	gjson.Parse(line).ForEach(func(key, value gjson.Result) bool {
		switch {
		case containsKey(key.String(), settings.Timestamp.Key):
			parsed.Timestamp = value.String()
		case containsKey(key.String(), settings.Logger.Key):
			parsed.Logger = value.String()
		case containsKey(key.String(), settings.Caller.Key):
			parsed.Caller = value.String()
		case containsKey(key.String(), settings.Level.Key):
			parsed.Level = value.String()
		case containsKey(key.String(), settings.Message.Key):
			parsed.Message = value.String()
		default:
			parsed.Fields = append(parsed.Fields, []string{key.String(), value.String()})
		}
		return true
	})

	return parsed, nil
}
