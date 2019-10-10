package prettifiers_test

import (
	"strings"
	"time"

	"github.com/globocom/prettylog/config"
	. "github.com/globocom/prettylog/prettifiers"

	"fmt"

	"github.com/fatih/color"
	"github.com/globocom/prettylog/parsers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/viper"
)

const (
	TIME         = "2019-10-16T07:51:08-03:00"
	LOGGER       = "root"
	CALLER       = "main.go:10"
	MESSAGE      = "foobar"
	FIELD1_NAME  = "field1"
	FIELD1_VALUE = "foo"
	FIELD2_NAME  = "field2"
	FIELD2_VALUE = "42"
	FIELD3_NAME  = "field3"
	FIELD3_VALUE = "true"
	FIELD4_NAME  = "field4"
	FIELD4_VALUE = "1 2 3 4 5"
)

var _ = Describe("Default prettifier", func() {
	var sut Prettifier

	BeforeEach(func() {
		setDefaultConfig()
		config.Load(false)
		sut = &DefaultPrettifier{}
	})

	Context("level is DEBUG", func() {
		It("should return a formatted string containing all fields", func() {
			// Arrange
			line := getParsedLine(TIME, "debug")

			// Act
			result := sut.Prettify(line)

			// Assert
			Expect(result).To(BeIdenticalTo(getFormattedLine(TIME, "debug", color.FgMagenta)))
		})
	})

	Context("level is INFO", func() {
		It("should return a formatted string containing all fields", func() {
			// Arrange
			line := getParsedLine(TIME, "info")

			// Act
			result := sut.Prettify(line)

			// Assert
			Expect(result).To(BeIdenticalTo(getFormattedLine(TIME, "info", color.FgBlue)))
		})
	})

	Context("level is WARN", func() {
		It("should return a formatted string containing all fields", func() {
			// Arrange
			line := getParsedLine(TIME, "warn")

			// Act
			result := sut.Prettify(line)

			// Assert
			Expect(result).To(BeIdenticalTo(getFormattedLine(TIME, "warn", color.FgYellow)))
		})
	})

	Context("level is ERROR", func() {
		It("should return a formatted string containing all fields", func() {
			// Arrange
			line := getParsedLine(TIME, "error")

			// Act
			result := sut.Prettify(line)

			// Assert
			Expect(result).To(BeIdenticalTo(getFormattedLine(TIME, "error", color.FgRed)))
		})
	})

	Context("timestamp format setting is defined", func() {
		It("should return timestamp formatted", func() {

			// Set temporary config
			config.GetSettings().Timestamp.Format = "02/01/2006"

			// Arrange
			tsFormat := config.GetSettings().Timestamp.Format
			parsedTime, err := time.Parse(time.RFC3339, TIME)
			if err != nil {
				fmt.Println(err)
			}
			ts := parsedTime.Format(tsFormat)
			line := getParsedLine(ts, "error")

			// Act
			result := sut.Prettify(line)

			// Unset config
			config.GetSettings().Timestamp.Format = "02/01/2006"

			// Assert
			Expect(result).To(BeIdenticalTo(getFormattedLine(ts, "error", color.FgRed)))
		})
	})
})

func setDefaultConfig() {
	viper.Set("timestamp.format", time.RFC3339)
	viper.Set("timestamp.key", "time")
	viper.Set("logger.key", "logger")
	viper.Set("logger.padding", 0)
	viper.Set("caller.key", "caller")
	viper.Set("caller.visible", true)
	viper.Set("caller.padding", 0)
	viper.Set("level.key", "level")
	viper.Set("level.padding", 0)
	viper.Set("message.key", "msg")
	viper.Set("message.padding", 0)
}

func getParsedLine(timestamp string, level string) *parsers.ParsedLine {
	return &parsers.ParsedLine{
		Timestamp: timestamp,
		Logger:    LOGGER,
		Level:     level,
		Caller:    CALLER,
		Message:   MESSAGE,
		Fields: [][]string{
			{FIELD1_NAME, FIELD1_VALUE},
			{FIELD2_NAME, FIELD2_VALUE},
			{FIELD3_NAME, FIELD3_VALUE},
			{FIELD4_NAME, FIELD4_VALUE},
		},
	}
}

func getFormattedLine(timestamp string, level string, levelColor color.Attribute) string {
	return fmt.Sprintf("%s %s %s %s %s %s=%s %s=%s %s=%s %s=\"%s\" ",
		color.New(color.FgYellow).Add(color.Faint).Sprint(timestamp),
		color.New(color.FgWhite).Add(color.Faint).Sprint(LOGGER),
		color.New(color.FgWhite).Add(color.Faint).Sprint(CALLER),
		color.New(levelColor).Sprint(strings.ToUpper(level)),
		MESSAGE,
		color.New(levelColor).Sprint(FIELD1_NAME), FIELD1_VALUE,
		color.New(levelColor).Sprint(FIELD2_NAME), FIELD2_VALUE,
		color.New(levelColor).Sprint(FIELD3_NAME), FIELD3_VALUE,
		color.New(levelColor).Sprint(FIELD4_NAME), FIELD4_VALUE)
}
