package prettifiers

import (
	"bytes"

	"strings"

	"github.com/fatih/color"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
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
	timeColorFunc       = color.New(color.FgYellow).Add(color.Faint).SprintFunc()
	loggerColorFunc     = color.New(color.FgWhite).Add(color.Faint).SprintFunc()
	callerColorFunc     = color.New(color.FgWhite).Add(color.Faint).SprintFunc()
	messageColorFunc    = color.New(color.FgWhite).SprintFunc()
	fieldValueColorFunc = color.New(color.FgWhite).SprintFunc()
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

type jsonPrettifier struct {
	TimestampField string
	LevelField     string
	LoggerField    string
	CallerField    string
	MessageField   string
	ShowTimestamp  bool
	ShowCaller     bool
}

func (p *jsonPrettifier) Prettify(line string) string {
	if !gjson.Valid(line) {
		return line
	}

	parsed := p.parseLine(line)
	return p.generateFormattedLine(parsed)
}

func (p *jsonPrettifier) parseLine(line string) *parsedLine {
	parsed := &parsedLine{}

	gjson.Parse(line).ForEach(func(key, value gjson.Result) bool {
		switch key.String() {
		case p.TimestampField:
			parsed.Timestamp = value.String()
		case p.LoggerField:
			parsed.Logger = value.String()
		case p.CallerField:
			parsed.Caller = value.String()
		case p.LevelField:
			parsed.Level = value.String()
		case p.MessageField:
			parsed.Message = value.String()
		default:
			parsed.Fields = append(parsed.Fields, []string{key.String(), value.String()})
		}
		return true
	})

	return parsed
}

func (p *jsonPrettifier) generateFormattedLine(parsed *parsedLine) string {
	levelColorFunc := getLevelColorFunc(parsed.Level)
	buffer := &bytes.Buffer{}

	if p.ShowTimestamp {
		buffer.WriteString(timeColorFunc(parsed.Timestamp))
		buffer.WriteString(SEPARATOR)
	}

	if parsed.Logger != "" {
		buffer.WriteString(loggerColorFunc(parsed.Logger))
		buffer.WriteString(SEPARATOR)
	}

	if p.ShowCaller {
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
	return &jsonPrettifier{
		TimestampField: viper.GetString("fields.timestamp"),
		LoggerField:    viper.GetString("fields.logger"),
		LevelField:     viper.GetString("fields.level"),
		CallerField:    viper.GetString("fields.caller"),
		MessageField:   viper.GetString("fields.message"),
		ShowTimestamp:  viper.GetBool("show.timestamp"),
		ShowCaller:     viper.GetBool("show.caller"),
	}
}
