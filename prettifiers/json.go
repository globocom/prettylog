package prettifiers

import (
	"encoding/json"
	"fmt"

	"bytes"

	"strings"

	"github.com/fatih/color"
	"github.com/spf13/viper"
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
	timeColorFunc   = color.New(color.FgYellow).Add(color.Faint).SprintFunc()
	loggerColorFunc = color.New(color.FgWhite).Add(color.Faint).SprintFunc()
	callerColorFunc = color.New(color.FgWhite).Add(color.Faint).SprintFunc()
	levelColorMap   = map[string]func(...interface{}) string{
		DEBUG_LEVEL: color.New(color.FgMagenta).SprintFunc(),
		INFO_LEVEL:  color.New(color.FgBlue).SprintFunc(),
		WARN_LEVEL:  color.New(color.FgYellow).SprintFunc(),
		ERROR_LEVEL: color.New(color.FgRed).SprintFunc(),
	}
)

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
	var data map[string]interface{}
	err := json.Unmarshal([]byte(line), &data)
	if err != nil {
		return line
	}

	time := getAndRemoveField(data, p.TimestampField)
	level := getAndRemoveField(data, p.LevelField)
	logger := getAndRemoveField(data, p.LoggerField)
	caller := getAndRemoveField(data, p.CallerField)
	msg := getAndRemoveField(data, p.MessageField)

	levelColorFunc := getLevelColorFunc(level)

	buffer := &bytes.Buffer{}

	if p.ShowTimestamp {
		buffer.WriteString(timeColorFunc(time))
		buffer.WriteString(SEPARATOR)
	}

	if logger != "" {
		buffer.WriteString(loggerColorFunc(logger))
		buffer.WriteString(SEPARATOR)
	}

	if p.ShowCaller {
		buffer.WriteString(callerColorFunc(caller))
		buffer.WriteString(SEPARATOR)
	}

	buffer.WriteString(levelColorFunc(strings.ToUpper(level)))
	buffer.WriteString(SEPARATOR)

	buffer.WriteString(msg)
	buffer.WriteString(SEPARATOR)

	for field, value := range data {
		buffer.WriteString(levelColorFunc(field))
		buffer.WriteString(FIELD_SEPARATOR)
		buffer.WriteString(fmt.Sprintf("%v", value))
		buffer.WriteString(SEPARATOR)
	}

	return buffer.String()
}

func getAndRemoveField(data map[string]interface{}, field string) string {
	if value, exists := data[field]; exists {
		delete(data, field)
		return fmt.Sprintf("%v", value)
	} else {
		return ""
	}
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
