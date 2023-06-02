package main

import (
	"bytes"
	"strconv"
)

var (
	Version    = "devel"
	Commit     = ""
	CommitDate = ""
	TreeState  = ""
)

const versionCommandOneLine = "Show version information."

const versionCommandHelp = `
Prints all available version data in a YAML-like format.
`

var versionCommand = &Command{
	Names:      []Name{Name{"version"}},
	OneLine:    versionCommandOneLine,
	Help:       versionCommandHelp,
	MinArgs:    0,
	MaxArgs:    0,
	HasMaxArgs: true,
	Run: func(params Params) int {
		versionMain(params.Out)
		return 0
	},
}

func versionMain(out *bytes.Buffer) {
	printValue := func(name string, value string, force bool) {
		if force || value != "" {
			out.WriteString(name)
			out.WriteString(": ")
			out.WriteString(strconv.Quote(value))
			out.WriteString("\n")
		}
	}

	printValue("version", Version, true)
	printValue("commit", Commit, false)
	printValue("commitDate", CommitDate, false)
	printValue("treeState", TreeState, false)
}
