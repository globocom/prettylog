package prettifiers

import (
	"bytes"

	"strings"

	"fmt"

	"github.com/fatih/color"
	"github.com/tidwall/gjson"
	"github.com/globocom/pretty-log/config"
)

const (
	DEBUG_LEVEL = "debug"
	INFO_LEVEL  = "info"
	WARN_LEVEL  = "warn"
	ERROR_LEVEL = "error"

	SEPARATOR       = " "
	FIELD_SEPARATOR = "="
)

var (
	noColorFunc = func(a ...interface{}) string { return fmt.Sprint(a...) }

	timeColorFunc       = color.New(color.FgYellow).Add(color.Faint).SprintFunc()
	loggerColorFunc     = color.New(color.FgWhite).Add(color.Faint).SprintFunc()
	callerColorFunc     = color.New(color.FgWhite).Add(color.Faint).SprintFunc()
	messageColorFunc    = noColorFunc
	fieldValueColorFunc = noColorFunc
	levelColorMap       = map[string]func(...interface{}) string{
		DEBUG_LEVEL: color.New(color.FgMagenta).SprintFunc(),
		INFO_LEVEL:  color.New(color.FgBlue).SprintFunc(),
		WARN_LEVEL:  color.New(color.FgYellow).SprintFunc(),
		ERROR_LEVEL: color.New(color.FgRed).SprintFunc(),
	}
)

type parsedLine struct {
	Timestamp string
	Logger    string
	Caller    string
	Level     string
	Message   string
	Fields    [][]string
}

type jsonPrettifier struct{}

func (p *jsonPrettifier) Prettify(line string) string {
	if !gjson.Valid(line) {
		return line
	}

	settings := config.GetSettings()
	parsed := parseLine(settings, line)
	return generateFormattedLine(settings, parsed)
}

func parseLine(settings *config.Settings, line string) *parsedLine {
	parsed := &parsedLine{}

	gjson.Parse(line).ForEach(func(key, value gjson.Result) bool {
		switch key.String() {
		case settings.Timestamp.Key:
			parsed.Timestamp = value.String()
		case settings.Logger.Key:
			parsed.Logger = value.String()
		case settings.Caller.Key:
			parsed.Caller = value.String()
		case settings.Level.Key:
			parsed.Level = value.String()
		case settings.Message.Key:
			parsed.Message = value.String()
		default:
			parsed.Fields = append(parsed.Fields, []string{key.String(), value.String()})
		}
		return true
	})

	return parsed
}

func generateFormattedLine(settings *config.Settings, parsed *parsedLine) string {
	levelColorFunc := getLevelColorFunc(parsed.Level)
	buffer := &bytes.Buffer{}

	if settings.Timestamp.Visible {
		buffer.WriteString(timeColorFunc(parsed.Timestamp))
		buffer.WriteString(SEPARATOR)
	}

	if parsed.Logger != "" {
		buffer.WriteString(loggerColorFunc(parsed.Logger))
		buffer.WriteString(SEPARATOR)
	}

	if settings.Caller.Visible {
		buffer.WriteString(callerColorFunc(parsed.Caller))
		buffer.WriteString(SEPARATOR)
	}

	buffer.WriteString(levelColorFunc(strings.ToUpper(parsed.Level)))
	buffer.WriteString(SEPARATOR)

	buffer.WriteString(messageColorFunc(parsed.Message))
	buffer.WriteString(SEPARATOR)

	for _, field := range parsed.Fields {
		buffer.WriteString(levelColorFunc(field[0]))
		buffer.WriteString(FIELD_SEPARATOR)
		buffer.WriteString(fieldValueColorFunc(field[1]))
		buffer.WriteString(SEPARATOR)
	}

	return buffer.String()
}

func getLevelColorFunc(level string) func(...interface{}) string {
	if value, exists := levelColorMap[strings.ToLower(level)]; exists {
		return value
	} else {
		return color.New(color.FgWhite).SprintFunc()
	}
}

func NewJsonPrettifier() Prettifier {
	return &jsonPrettifier{}
}
