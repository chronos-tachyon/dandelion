package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/mattn/go-isatty"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	logOutputSTDOUT  = "stdout"
	logOutputSTDERR  = "stderr"
	logFormatJSON    = "json"
	logFormatConsole = "console"
	logColorAuto     = "auto"
	logColorYes      = "yes"
	logColorNo       = "no"
	logColorOn       = "on"
	logColorOff      = "off"
	logColorOne      = "1"
	logColorZero     = "0"
)

var logFormatArray = [...]string{
	logFormatJSON,
	logFormatConsole,
}

var logColorArray = [...]string{
	logColorAuto,
	logColorNo,
}

var logTimeFormatMap = map[string]string{
	"kitchen":    "3:04PM",
	"kitchen.s":  "3:04:05PM",
	"kitchen.ms": "3:04:05.999PM",
	"kitchen.us": "3:04:05.999999PM",
	"kitchen.ns": "3:04:05.999999999PM",
	"rfc822":     "02 Jan 2006 15:04 -0700",
	"rfc822.s":   "02 Jan 2006 15:04:05 -0700",
	"rfc822.ms":  "02 Jan 2006 15:04:05.999 -0700",
	"rfc822.us":  "02 Jan 2006 15:04:05.999999 -0700",
	"rfc822.ns":  "02 Jan 2006 15:04:05.999999999 -0700",
	"rfc1123":    "Mon, 02 Jan 2006 15:04 -0700",
	"rfc1123.s":  "Mon, 02 Jan 2006 15:04:05 -0700",
	"rfc1123.ms": "Mon, 02 Jan 2006 15:04:05.999 -0700",
	"rfc1123.us": "Mon, 02 Jan 2006 15:04:05.999999 -0700",
	"rfc1123.ns": "Mon, 02 Jan 2006 15:04:05.999999999 -0700",
	"rfc3339":    "2006-01-02T15:04Z07:00",
	"rfc3339.s":  "2006-01-02T15:04:05Z07:00",
	"rfc3339.ms": "2006-01-02T15:04:05.999Z07:00",
	"rfc3339.us": "2006-01-02T15:04:05.999999Z07:00",
	"rfc3339.ns": "2006-01-02T15:04:05.999999999Z07:00",
	"iso8601":    "2006-01-02T15:04Z07:00",
	"iso8601.s":  "2006-01-02T15:04:05Z07:00",
	"iso8601.ms": "2006-01-02T15:04:05.999Z07:00",
	"iso8601.us": "2006-01-02T15:04:05.999999Z07:00",
	"iso8601.ns": "2006-01-02T15:04:05.999999999Z07:00",
}

func getenv(name string, defaultValue string) string {
	if value, found := os.LookupEnv(name); found {
		return value
	}
	return defaultValue
}

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	zerolog.DurationFieldUnit = time.Second
	zerolog.DurationFieldInteger = false

	if str, found := os.LookupEnv("LOG_LEVEL"); found {
		level, err := zerolog.ParseLevel(str)
		if err != nil {
			panic(err)
		}
		zerolog.SetGlobalLevel(level)
	}

	color := triStateAuto
	if str, found := os.LookupEnv("LOG_COLOR"); found {
		if err := color.Parse(str); err != nil {
			panic(err)
		}
	}

	var logOutputFile *os.File
	switch logOutput := getenv("LOG_OUTPUT", logOutputSTDERR); logOutput {
	case logOutputSTDOUT:
		logOutputFile = os.Stdout
	case logOutputSTDERR:
		logOutputFile = os.Stderr
	default:
		var err error
		logOutputFile, err = os.OpenFile(logOutput, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o666)
		if err != nil {
			panic(fmt.Errorf("failed to open log file %q for appending: %w", logOutput, err))
		}
	}

	defaultLogFormat := logFormatJSON
	switch {
	case isatty.IsTerminal(logOutputFile.Fd()):
		defaultLogFormat = logFormatConsole
	case isatty.IsCygwinTerminal(logOutputFile.Fd()):
		defaultLogFormat = logFormatConsole
	default:
		if color == triStateAuto {
			color = triStateNo
		}
	}

	var logWriter io.Writer
	var c *zerolog.ConsoleWriter
	switch logFormat := getenv("LOG_FORMAT", defaultLogFormat); logFormat {
	case logFormatJSON:
		logWriter = logOutputFile
	case logFormatConsole:
		c = &zerolog.ConsoleWriter{Out: logOutputFile, NoColor: color == triStateNo}
		logWriter = c
	default:
		panic(fmt.Errorf("unknown log format %q; expected one of %q", logFormat, logFormatArray[:]))
	}

	if logTimeFormat, found := os.LookupEnv("LOG_TIMEFORMAT"); found {
		key := strings.ToLower(strings.ReplaceAll(logTimeFormat, "µ", "u"))
		if value, found := logTimeFormatMap[key]; found {
			logTimeFormat = value
		}

		if c == nil {
			zerolog.TimeFieldFormat = logTimeFormat
		} else {
			c.TimeFormat = logTimeFormat
		}
	}

	log.Logger = zerolog.New(logWriter).With().Timestamp().Logger()
	zerolog.DefaultContextLogger = &log.Logger
}

type triState byte

const (
	triStateAuto triState = iota
	triStateYes
	triStateNo
)

var triStateGoNames = [...]string{"triStateAuto", "triStateYes", "triStateNo"}
var triStateNames = [...]string{"auto", "yes", "no"}
var triStateMap = map[string]triState{
	"":     triStateAuto,
	"auto": triStateAuto,

	"1":    triStateYes,
	"y":    triStateYes,
	"yes":  triStateYes,
	"t":    triStateYes,
	"true": triStateYes,
	"on":   triStateYes,

	"0":     triStateNo,
	"n":     triStateNo,
	"no":    triStateNo,
	"f":     triStateNo,
	"false": triStateNo,
	"off":   triStateNo,
}

func (enum triState) GoString() string {
	if enum < triState(len(triStateGoNames)) {
		return triStateGoNames[enum]
	}
	return fmt.Sprintf("triState(%d)", uint(enum))
}

func (enum triState) String() string {
	if enum < triState(len(triStateNames)) {
		return triStateNames[enum]
	}
	return triStateNames[0]
}

func (enum triState) MarshalText() ([]byte, error) {
	return []byte(enum.String()), nil
}

func (enum *triState) Parse(input string) error {
	*enum = 0

	if value, found := triStateMap[input]; found {
		*enum = value
		return nil
	}

	lc := strings.ToLower(input)
	if value, found := triStateMap[lc]; found {
		*enum = value
		return nil
	}

	return fmt.Errorf("unknown tri-state value %q", input)
}
