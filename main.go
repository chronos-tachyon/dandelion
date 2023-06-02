package main

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/adrg/xdg"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const appName = "dandelion"

var reNodeName = regexp.MustCompile(`^[A-Za-z][0-9A-Za-z]*(?:[_-][0-9A-Za-z]+)*$`)

var (
	programName string
	commandList []*Command
)

func main() {
	stdout := bytes.NewBuffer(make([]byte, 0, 4096))
	stderr := bytes.NewBuffer(make([]byte, 0, 256))
	var exitCode int
	defer func() {
		_, _ = io.Copy(os.Stderr, stderr)
		_, _ = io.Copy(os.Stdout, stdout)
		os.Exit(exitCode)
	}()

	programName = "dandelion"
	commandList = []*Command{
		versionCommand,
		helpAllCommand,
		helpCommand,
		serveCommand,
	}

	var rawArgs []string
	var rawArgsLen uint
	if len(os.Args) > 0 {
		programName = filepath.Base(os.Args[0])
		rawArgs = os.Args[1:]
		rawArgsLen = uint(len(rawArgs))
	}

	var wantHelp bool
	var wantVersion bool
	var nodeName string
	var hasNodeName bool

	args := make([]string, 0, rawArgsLen)
	i := uint(0)
	for i < rawArgsLen {
		rawArg := rawArgs[i]
		i++

		switch {
		case rawArg == "--":
			args = append(args, rawArgs[i:]...)
			rawArgs = nil
			rawArgsLen = 0

		case rawArg == "--help":
			wantHelp = true

		case rawArg == "--version":
			wantVersion = true

		case rawArg == "--name":
			nodeName = rawArgs[i]
			hasNodeName = true
			i++

		case strings.HasPrefix(rawArg, "--name="):
			nodeName = strings.TrimPrefix(rawArg, "--name=")
			hasNodeName = true

		case strings.HasPrefix(rawArg, "--"):
			rawArg = rawArg[2:]
			log.Error().Msgf("unknown long flag --%s", rawArg)
			exitCode = 1
			return

		case strings.HasPrefix(rawArg, "-"):
			rawArg = rawArg[1:]
			runes := []rune(rawArg)
			runesLen := uint(len(runes))
			j := uint(0)
			for j < runesLen {
				ch := runes[j]
				j++
				switch ch {
				case 'h':
					wantHelp = true
				case 'V':
					wantVersion = true
				case 'n':
					if j < runesLen {
						nodeName = string(runes[j:])
						hasNodeName = true
						runes = nil
						runesLen = 0
					} else {
						nodeName = rawArgs[i]
						hasNodeName = true
						i++
					}
				default:
					log.Error().Msgf("unknown short flag -%c", ch)
					exitCode = 1
					return
				}
			}

		default:
			args = append(args, rawArg)
		}
	}

	if wantVersion {
		versionMain(stdout)
		return
	}

	if wantHelp {
		help(stdout, args)
		return
	}

	argsLen := uint(len(args))
	if argsLen <= 0 {
		helpAll(stdout)
		return
	}

	if !hasNodeName {
		nodeName = "main"
		if value, ok := os.LookupEnv("DANDELION_NODE"); ok {
			nodeName = value
		}
	}

	if !reNodeName.MatchString(nodeName) {
		log.Error().Msgf("%q is not a valid name for your Dandelion node; it must match %v", nodeName, reNodeName)
		exitCode = 1
		return
	}

	var params Params
	params.Out = stdout
	params.Err = stderr

	params.ConfigHome = filepath.Join(xdg.ConfigHome, appName)
	params.DataHome = filepath.Join(xdg.DataHome, appName)
	params.StateHome = filepath.Join(xdg.StateHome, appName)
	params.CacheHome = filepath.Join(xdg.CacheHome, appName)
	params.RuntimeDir = xdg.RuntimeDir
	params.FontDirs = xdg.FontDirs
	params.AppDirs = xdg.ApplicationDirs
	params.UserDirs = xdg.UserDirs
	params.NodeName = nodeName

	params.ConfigDirs = make([]string, len(xdg.ConfigDirs))
	for i, base := range xdg.ConfigDirs {
		params.ConfigDirs[i] = filepath.Join(base, appName)
	}

	params.DataDirs = make([]string, len(xdg.DataDirs))
	for i, base := range xdg.DataDirs {
		params.DataDirs[i] = filepath.Join(base, appName)
	}

	if x, ok := ExactCommand(args); ok {
		log.Logger.UpdateContext(func(c zerolog.Context) zerolog.Context {
			return c.Strs("cmd", x.Name)
		})

		argsLen := uint(len(x.Args))
		if argsLen < x.Command.MinArgs {
			log.Fatal().
				Strs("args", x.Args).
				Msgf("need at least %d positional args; got %d", x.Command.MinArgs, argsLen)
			panic(nil)
		}
		if x.Command.HasMaxArgs && argsLen > x.Command.MaxArgs {
			log.Fatal().
				Strs("args", x.Args).
				Msgf("need at most %d positional args; got %d", x.Command.MaxArgs, argsLen)
			panic(nil)
		}
		params.Args = x.Args
		exitCode = x.Command.Run(params)
		return
	}

	log.Error().Strs("cmd", args).Msg("unknown command")
	helpFuzzy(stdout, args, FuzzyCommands(args))
	exitCode = 1
}

func stringsEqual(a []string, b []string) bool {
	aLen := uint(len(a))
	bLen := uint(len(b))
	if aLen != bLen {
		return false
	}
	for i := uint(0); i < aLen; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
