package main

import (
	"github.com/rs/zerolog/log"

	"github.com/chronos-tachyon/dandelion/config"
)

const serveCommandOneLine = "Run a Dandelion server."

const serveCommandHelp = ``

var serveCommand = &Command{
	Names:      []Name{Name{"serve"}},
	OneLine:    serveCommandOneLine,
	Help:       serveCommandHelp,
	MinArgs:    0,
	MaxArgs:    0,
	HasMaxArgs: true,
	Run: func(params Params) int {
		log.Info().Interface("params", params).Msg("hello from ServeMain!")
		var file config.File
		if err := file.Load(params.NodeName, params.ConfigHome, params.ConfigDirs); err != nil {
			log.Error().Err(err).Msg("failed to load configuration")
			return 1
		}
		log.Info().Interface("file", file).Msg("configuration loaded")
		return 0
	},
}
