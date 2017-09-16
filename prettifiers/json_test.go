package prettifiers_test

import (
	"strings"

	"github.com/globocom/pretty-log/config"
	. "github.com/globocom/pretty-log/prettifiers"

	"fmt"

	"github.com/fatih/color"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/viper"
)

const (
	TIME         = "2017-01-01T12:10:11"
	LOGGER       = "root"
	CALLER       = "main.go:10"
	MESSAGE      = "foobar"
	FIELD1_NAME  = "field1"
	FIELD1_VALUE = "foo"
	FIELD2_NAME  = "field2"
	FIELD2_VALUE = "42"
	FIELD3_NAME  = "field3"
	FIELD3_VALUE = "true"
)

var _ = Describe("JSON prettifier", func() {
	var sut Prettifier

	BeforeEach(func() {
		config.Load()
		setDefaultConfig()
		sut = NewJsonPrettifier()
	})

	Context("Line isn't a JSON string", func() {
		It("should return the line itself", func() {
			Expect(sut.Prettify("foobar")).To(Equal("foobar"))
		})
	})

	Context("Line is a JSON string", func() {
		Context("level is DEBUG", func() {
			It("should return a formatted string containing all fields", func() {
				// Arrange
				line := getSampleJson("debug")

				// Act
				result := sut.Prettify(line)

				// Assert
				Expect(result).To(BeIdenticalTo(getFormattedLine("debug", color.FgMagenta)))
			})
		})

		Context("level is INFO", func() {
			It("should return a formatted string containing all fields", func() {
				// Arrange
				line := getSampleJson("info")

				// Act
				result := sut.Prettify(line)

				// Assert
				Expect(result).To(BeIdenticalTo(getFormattedLine("info", color.FgBlue)))
			})
		})

		Context("level is WARN", func() {
			It("should return a formatted string containing all fields", func() {
				// Arrange
				line := getSampleJson("warn")

				// Act
				result := sut.Prettify(line)

				// Assert
				Expect(result).To(BeIdenticalTo(getFormattedLine("warn", color.FgYellow)))
			})
		})

		Context("level is ERROR", func() {
			It("should return a formatted string containing all fields", func() {
				// Arrange
				line := getSampleJson("error")

				// Act
				result := sut.Prettify(line)

				// Assert
				Expect(result).To(BeIdenticalTo(getFormattedLine("error", color.FgRed)))
			})
		})
	})
})

func setDefaultConfig() {
	viper.Set("timestamp.key", "time")
	viper.Set("logger.key", "logger")
	viper.Set("caller.key", "caller")
	viper.Set("caller.visible", true)
	viper.Set("level.key", "level")
	viper.Set("message.key", "msg")
}

func getSampleJson(level string) string {
	return fmt.Sprintf(`{
		"time": "%s",
		"logger": "%s",
		"level": "%s",
		"caller": "%s",
		"msg": "%s",
		"%s": "%s",
		"%s": %s,
		"%s": %s
		}`, TIME, LOGGER, level, CALLER, MESSAGE, FIELD1_NAME, FIELD1_VALUE, FIELD2_NAME, FIELD2_VALUE, FIELD3_NAME, FIELD3_VALUE)
}

func getFormattedLine(level string, levelColor color.Attribute) string {
	return fmt.Sprintf("%s %s %s %s %s %s=%s %s=%s %s=%s ",
		color.New(color.FgYellow).Add(color.Faint).Sprint(TIME),
		color.New(color.FgWhite).Add(color.Faint).Sprint(LOGGER),
		color.New(color.FgWhite).Add(color.Faint).Sprint(CALLER),
		color.New(levelColor).Sprint(strings.ToUpper(level)),
		MESSAGE,
		color.New(levelColor).Sprint(FIELD1_NAME), FIELD1_VALUE,
		color.New(levelColor).Sprint(FIELD2_NAME), FIELD2_VALUE,
		color.New(levelColor).Sprint(FIELD3_NAME), FIELD3_VALUE)
}
