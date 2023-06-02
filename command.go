package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/adrg/xdg"
)

type Command struct {
	Names      []Name
	OneLine    string
	Help       string
	MinArgs    uint
	MaxArgs    uint
	HasMaxArgs bool
	Run        func(Params) int
}

type Params struct {
	Out *bytes.Buffer
	Err *bytes.Buffer

	Args       []string
	ConfigDirs []string
	DataDirs   []string
	FontDirs   []string
	AppDirs    []string
	UserDirs   xdg.UserDirectories
	ConfigHome string
	DataHome   string
	StateHome  string
	CacheHome  string
	RuntimeDir string
	NodeName   string
}

type Name []string

func (name Name) AppendTo(out []byte, verbose bool) []byte {
	if verbose {
		out = append(out, '[')
		for i, str := range name {
			if i > 0 {
				out = append(out, ',', ' ')
			}
			out = strconv.AppendQuote(out, str)
		}
		out = append(out, ']')
		return out
	}

	for i, str := range name {
		if i > 0 {
			out = append(out, ' ')
		}
		out = append(out, str...)
	}
	return out
}

func (name Name) GoString() string {
	var scratch [64]byte
	return string(name.AppendTo(scratch[:0], true))
}

func (name Name) String() string {
	var scratch [64]byte
	return string(name.AppendTo(scratch[:0], false))
}

func (name Name) Equals(other Name) bool {
	nameLen := uint(len(name))
	otherLen := uint(len(other))
	if nameLen != otherLen {
		return false
	}
	for i := uint(0); i < nameLen; i++ {
		a := name[i]
		b := other[i]
		if a != b {
			return false
		}
	}
	return true
}

func (name Name) CompareTo(other Name) int {
	nameLen := uint(len(name))
	otherLen := uint(len(other))

	minLen := nameLen
	if minLen > otherLen {
		minLen = otherLen
	}

	for i := uint(0); i < minLen; i++ {
		a := name[i]
		b := other[i]
		if cmp := strings.Compare(a, b); cmp != 0 {
			return cmp
		}
	}
	if nameLen < otherLen {
		return -1
	}
	if nameLen > otherLen {
		return 1
	}
	return 0
}

type Match struct {
	Command *Command
	Name    Name
	Args    []string
}

func ExactCommand(args []string) (Match, bool) {
	for _, cmd := range commandList {
		for _, name := range cmd.Names {
			_, nameLen, n := commonPrefixLength(args, name)
			if n < nameLen {
				continue
			}

			rest := args[nameLen:]
			restLen := uint(len(rest))
			if restLen < cmd.MinArgs {
				continue
			}
			if cmd.HasMaxArgs && restLen > cmd.MaxArgs {
				continue
			}
			return Match{Command: cmd, Name: name, Args: rest}, true
		}
	}
	return Match{}, false
}

func FuzzyCommands(args []string) []Match {
	out := make([]Match, 0, len(commandList))
	for _, cmd := range commandList {
		for _, name := range cmd.Names {
			argsLen, nameLen, n := commonPrefixLength(args, name)
			if n < argsLen && n < nameLen {
				continue
			}

			var rest []string
			var restLen uint
			if n >= nameLen {
				rest = args[nameLen:]
				restLen = uint(len(rest))
				if restLen >= cmd.MinArgs {
					continue
				}
			}
			out = append(out, Match{Command: cmd, Name: name, Args: rest})
		}
	}
	return out
}

func Usage(out *bytes.Buffer, programName string, name Name, cmd *Command) {
	if cmd.OneLine != "" {
		out.WriteString(cmd.OneLine)
		out.WriteByte('\n')
	}
	out.WriteString("Usage: ")
	UsageLine(out, programName, name, cmd)
	out.WriteString(cmd.Help)
}

func UsageLine(out *bytes.Buffer, programName string, name Name, cmd *Command) {
	out.WriteString(programName)
	for _, piece := range name {
		out.WriteByte(' ')
		out.WriteString(piece)
	}
	var argIndex uint
	for argIndex < cmd.MinArgs {
		fmt.Fprintf(out, " <arg%d>", argIndex)
		argIndex++
	}
	if cmd.HasMaxArgs {
		for argIndex < cmd.MaxArgs {
			fmt.Fprintf(out, " [<arg%d>]", argIndex)
			argIndex++
		}
		out.WriteByte('\n')
		return
	}
	out.WriteString(" [...]\n")
}

func commonPrefixLength[A ~[]string, B ~[]string](a A, b B) (aLen uint, bLen uint, prefixLen uint) {
	aLen = uint(len(a))
	bLen = uint(len(b))
	prefixLen = uint(0)
	for prefixLen < aLen && prefixLen < bLen && a[prefixLen] == b[prefixLen] {
		prefixLen++
	}
	return
}
