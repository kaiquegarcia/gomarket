package util

import (
	"fmt"
	"strings"
)

type AskBoolOption func(c *askBoolConfig)
type askBoolConfig struct {
	shouldAppendOptions bool
	trueSymbol          string
	falseSymbol         string
}

func AskBoolCLI(text string, options ...AskBoolOption) bool {
	config := &askBoolConfig{
		shouldAppendOptions: true,
		trueSymbol:          "Y",
		falseSymbol:         "N",
	}

	for _, opt := range options {
		opt(config)
	}

	if strings.ToUpper(config.trueSymbol) == strings.ToUpper(config.falseSymbol) {
		panic("development error: trueSymbol cannot be equal to falseSymbol")
	}

	if config.shouldAppendOptions {
		text += fmt.Sprintf(" (%s/%s)", config.trueSymbol, config.falseSymbol)
	}
	return strings.ToUpper(AskCLI(text)) == strings.ToUpper(config.trueSymbol)
}

func ShouldAppendOptionsToText(shouldAppendOptions bool) AskBoolOption {
	return func(c *askBoolConfig) {
		c.shouldAppendOptions = shouldAppendOptions
	}
}

func WithTrueSymbol(trueSymbol string) AskBoolOption {
	return func(c *askBoolConfig) {
		c.trueSymbol = trueSymbol
	}
}

func WithFalseSymbol(falseSymbol string) AskBoolOption {
	return func(c *askBoolConfig) {
		c.falseSymbol = falseSymbol
	}
}
