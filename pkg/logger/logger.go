package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

type Logger struct {
	l *zerolog.Logger
}

var L *Logger

func Load(level string) {
	var lvl zerolog.Level

	switch strings.ToLower(level) {
	case "error":
		lvl = zerolog.ErrorLevel
	case "warn":
		lvl = zerolog.WarnLevel
	case "info":
		lvl = zerolog.InfoLevel
	case "debug":
		lvl = zerolog.DebugLevel
	}

	zerolog.SetGlobalLevel(lvl)

	consoleOut := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "15:04:05",
		NoColor:    false,
	}

	l := zerolog.New(consoleOut).With().Timestamp().Logger()

	L = &Logger{
		l: &l,
	}
}

func (lg *Logger) Debug(message interface{}, args ...interface{}) {
	lg.logAtLevel(zerolog.DebugLevel, message, args...)
}

func (lg *Logger) Info(message string, args ...interface{}) {
	lg.logAtLevel(zerolog.InfoLevel, message, args...)
}

func (lg *Logger) Warn(message string, args ...interface{}) {
	lg.logAtLevel(zerolog.WarnLevel, message, args...)
}

func (lg *Logger) Error(message interface{}, args ...interface{}) {
	lg.logAtLevel(zerolog.ErrorLevel, message, args...)
}

func (lg *Logger) ErrorStackTrace(err error) {
	lg.l.Error().Stack().Err(err).Msg("")
}

func (lg *Logger) Fatal(message interface{}, args ...interface{}) {
	lg.logAtLevel(zerolog.FatalLevel, message, args...)
	os.Exit(1)
}

func (lg *Logger) logAtLevel(level zerolog.Level, message interface{}, args ...interface{}) {
	msgStr := lg.formatMessage(message)
	event := lg.l.WithLevel(level)
	if len(args) == 0 {
		event.Msg(msgStr)
	} else {
		event.Msgf(msgStr, args...)
	}
}

func (lg *Logger) formatMessage(message interface{}) string {
	switch msg := message.(type) {
	case error:
		return msg.Error()
	case string:
		return msg
	default:
		return fmt.Sprintf("%v", message)
	}
}
