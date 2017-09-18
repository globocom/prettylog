package prettifiers

import (
	"bytes"

	"strings"

	"github.com/fatih/color"
	"github.com/tidwall/gjson"
	"github.com/globocom/prettylog/config"
)

const (
	SEPARATOR       = " "
	FIELD_SEPARATOR = "="
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
	buffer := &bytes.Buffer{}

	if settings.Timestamp.Visible {
		writeTo(buffer, parsed.Timestamp, 0, settings.Timestamp.Color)
	}

	if settings.Logger.Visible && parsed.Logger != "" {
		writeTo(buffer, parsed.Logger, settings.Logger.Padding, settings.Logger.Color)
	}

	if settings.Caller.Visible {
		writeTo(buffer, parsed.Caller, settings.Caller.Padding, settings.Caller.Color)
	}

	if settings.Level.Visible {
		writeTo(buffer, strings.ToUpper(parsed.Level), settings.Level.Padding, settings.Level.GetColorAttr(parsed.Level))
	}

	writeTo(buffer, parsed.Message, settings.Message.Padding, settings.Message.Color)
	writeFieldsTo(buffer, parsed.Fields, settings.Level.GetColorAttr(parsed.Level))

	return buffer.String()
}

func writeTo(buffer *bytes.Buffer, value string, padding int, colorAttrs []color.Attribute) {
	color := parseColor(colorAttrs)
	value = padRight(value, padding)

	if color != nil {
		value = color.Sprint(value)
	}

	buffer.WriteString(value)
	buffer.WriteString(SEPARATOR)
}

func writeFieldsTo(buffer *bytes.Buffer, fields [][]string, colorsAttrs []color.Attribute) {
	color := parseColor(colorsAttrs)

	for _, field := range fields {
		if color != nil {
			buffer.WriteString(color.Sprint(field[0]))
		} else {
			buffer.WriteString(field[0])
		}
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

func parseColor(attributes []color.Attribute) *color.Color {
	var c *color.Color

	if len(attributes) > 0 {
		c = color.New(attributes[0])
	}

	if len(attributes) > 1 {
		c.Add(attributes[1])
	}

	if len(attributes) > 2 {
		c.Add(attributes[2])
	}

	return c
}

func NewJsonPrettifier() Prettifier {
	return &jsonPrettifier{}
}
