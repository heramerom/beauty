package main

import "strings"

type logLevel int8

const (
	levelDebug logLevel = iota
	levelInfo
	levelWarn
	levelError
	levelUnknown
)

func ParseLevel(level string) logLevel {
	switch strings.ToLower(level) {
	case "e", "error":
		return levelError
	case "i", "info":
		return levelInfo
	case "w", "warn", "warning":
		return levelWarn
	case "d", "debug":
		return levelDebug
	}
	return levelUnknown
}

func (l logLevel) Color() []byte {
	switch l {
	case levelDebug:
		return colorLightCyan
	case levelInfo:
		return colorLightGreen
	case levelWarn:
		return colorLightYellow
	case levelError:
		return colorLightRed
	}
	return colorDefault
}
