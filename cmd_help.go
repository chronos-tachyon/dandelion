package main

import (
	"bytes"
	"fmt"
)

const helpFlags = `
global flags:
 -h, --help         show usage information
 -V, --version      show version information
 -n, --name=NAME    set instance name [default: $DANDELION_NODE]
`

const helpEnv = `
environment variables:
  DANDELION_NODE=<name>
	specify default value of -n / --name flag
	default: "main"

  LOG_OUTPUT=<stdin|stdout|path>
	specify logging output
	default: "stderr"

  LOG_FORMAT=<console|json>
	specify logging format
	default: "console" if $LOG_OUTPUT is a TTY, "json" otherwise

  LOG_COLOR=<yes|no|auto>
	specify logging colorization for when $LOG_FORMAT is "console"
	default: "auto", which is the same as "yes" if $LOG_OUTPUT is a TTY, "no" otherwise

  LOG_LEVEL=<trace|debug|info|warn|error|fatal>
	specify logging level
	default: "info"

  LOG_TIMEFORMAT=<zerolog time format>
	specify a custom time format string
	default: "" if $LOG_FORMAT is "console", "UNIXMS" otherwise
`

const helpCommandOneLine = "Show detailed usage information for the specified command."

const helpCommandHelp = `
If no arguments are given, then this command prints general usage
information to stdout, followed by the same output tha would be
produced by "help all".

If one or more arguments are given, they are taken as the name of a
command or group of commands to match.

If an exact match is found, then the detailed help for the matching
command is written to stdout.

If one or more fuzzy matches are found, then the names of the matching
commands are printed to stdout along with their one-line descriptions.
`

var helpCommand = &Command{
	Names:      []Name{Name{"help"}},
	OneLine:    helpCommandOneLine,
	Help:       helpCommandHelp,
	MinArgs:    0,
	MaxArgs:    0,
	HasMaxArgs: false,
	Run: func(params Params) int {
		help(params.Out, params.Args)
		return 0
	},
}

const helpAllCommandOneLine = "Show list of all available commands."

const helpAllCommandHelp = `
Lists all available commands, including aliases.
`

var helpAllCommand = &Command{
	Names:      []Name{Name{"help", "all"}},
	OneLine:    helpAllCommandOneLine,
	Help:       helpAllCommandHelp,
	MinArgs:    0,
	MaxArgs:    0,
	HasMaxArgs: true,
	Run: func(params Params) int {
		helpAll(params.Out)
		return 0
	},
}

func help(out *bytes.Buffer, args []string) {
	argsLen := uint(len(args))
	if argsLen <= 0 {
		helpAll(out)
		return
	}

	if x, ok := ExactCommand(args); ok {
		helpExact(out, x)
		return
	}

	helpFuzzy(out, args, FuzzyCommands(args))
}

func helpAll(out *bytes.Buffer) {
	out.WriteString("Available commands:\n\n")
	for _, cmd := range commandList {
		UsageLine(out, programName, cmd.Names[0], cmd)
		for _, alias := range cmd.Names[1:] {
			fmt.Fprintf(out, "  [alias] %v\n", alias)
		}
	}
	out.WriteString(helpFlags)
	out.WriteString(helpEnv)
}

func helpExact(out *bytes.Buffer, exact Match) {
	Usage(out, programName, exact.Name, exact.Command)
	out.WriteString(helpFlags)
	out.WriteString(helpEnv)
}

func helpFuzzy(out *bytes.Buffer, args []string, fuzzy []Match) {
	if len(fuzzy) <= 0 {
		fmt.Fprintf(out, "No matches found for %q.\n", args)
	} else {
		fmt.Fprintf(out, "No exact match found for %q.  Fuzzy matches:\n\n", args)
		for _, x := range fuzzy {
			UsageLine(out, programName, x.Name, x.Command)
		}
	}
	out.WriteString(helpFlags)
	out.WriteString(helpEnv)
}
