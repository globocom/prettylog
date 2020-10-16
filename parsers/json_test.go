package parsers_test

import (
	"github.com/globocom/prettylog/config"
	. "github.com/globocom/prettylog/parsers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/viper"
)

var _ = Describe("JSON line parser", func() {
	BeforeEach(func() {
		setDefaultConfig()
		config.Load(false)
	})

	Context("Line isn't parseable", func() {
		It("should return an error", func() {
			// Arrange
			line := "NOT JSON"
			sut := &JsonLineParser{}

			// Act
			_, err := sut.Parse(line)

			// Assert
			Expect(err).To(MatchError(ErrNonParseableLine))
		})
	})

	Context("Line is parseable", func() {
		It("should return a parsed line", func() {
			// Arrange
			line := `{
				"t":"2017-01-01",
				"lg": "foo",
				"ln": "main.go:10",
				"lvl": "warn",
				"field1": "bar",
				"field2": 42,
				"field3": true,
				"msg": "expected"
			}`

			sut := &JsonLineParser{}

			// Act
			parsed, err := sut.Parse(line)

			// Assert
			Expect(err).To(Succeed())
			Expect(parsed).To(Equal(&ParsedLine{
				Timestamp: "2017-01-01",
				Logger:    "foo",
				Caller:    "main.go:10",
				Level:     "warn",
				Message: "expected",
				Fields: [][]string{
					{"field1", "bar"},
					{"field2", "42"},
					{"field3", "true"},
				},
			}))
		})
	})
})

func setDefaultConfig() {
	viper.Set("timestamp.key", "t")
	viper.Set("logger.key", "lg")
	viper.Set("caller.key", "ln")
	viper.Set("level.key", "lvl")
	viper.Set("message.key", "msg")
}
