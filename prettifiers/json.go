package prettifiers

import (
	"bytes"

	"strings"

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
	noColor        = (*color.Color)(nil)
	timestampColor = color.New(color.FgYellow).Add(color.Faint)
	loggerColor    = color.New(color.FgWhite).Add(color.Faint)
	callerColor    = color.New(color.FgWhite).Add(color.Faint)
	messageColor   = noColor
	levelColor     = map[string]*color.Color{
		DEBUG_LEVEL: color.New(color.FgMagenta),
		INFO_LEVEL:  color.New(color.FgBlue),
		WARN_LEVEL:  color.New(color.FgYellow),
		ERROR_LEVEL: color.New(color.FgRed),
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
	levelColor := getLevelColor(parsed.Level)
	buffer := &bytes.Buffer{}

	if settings.Timestamp.Visible {
		writeTo(buffer, parsed.Timestamp, 0, timestampColor)
	}

	if settings.Logger.Visible && parsed.Logger != "" {
		writeTo(buffer, parsed.Logger, settings.Logger.Padding, loggerColor)
	}

	if settings.Caller.Visible {
		writeTo(buffer, parsed.Caller, settings.Caller.Padding, callerColor)
	}

	if settings.Level.Visible {
		writeTo(buffer, strings.ToUpper(parsed.Level), settings.Level.Padding, levelColor)
	}

	writeTo(buffer, parsed.Message, settings.Message.Padding, messageColor)
	writeFieldsTo(buffer, parsed.Fields, levelColor)

	return buffer.String()
}

func writeTo(buffer *bytes.Buffer, value string, padding int, color *color.Color) {
	value = padRight(value, padding)

	if color != nil {
		value = color.Sprint(value)
	}

	buffer.WriteString(value)
	buffer.WriteString(SEPARATOR)
}

func writeFieldsTo(buffer *bytes.Buffer, fields [][]string, color *color.Color) {
	for _, field := range fields {
		buffer.WriteString(color.Sprint(field[0]))
		buffer.WriteString(FIELD_SEPARATOR)
		buffer.WriteString(field[1])
		buffer.WriteString(SEPARATOR)
	}
}

func padRight(str string, size int) string {
	size = size - len(str)
	if size < 0 {
		size = 0
	}
	return str + strings.Repeat(" ", size)
}

func getLevelColor(level string) *color.Color {
	if value, exists := levelColor[strings.ToLower(level)]; exists {
		return value
	} else {
		return nil
	}
}

func NewJsonPrettifier() Prettifier {
	return &jsonPrettifier{}
}
