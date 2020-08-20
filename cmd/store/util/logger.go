package util

import (
	"os"

	"github.com/bamboovir/dockerfile/cmd/types"
	log "github.com/sirupsen/logrus"
)

// SetupLogger defination
func SetupLogger(logger *log.Logger, prop *types.RootArgs) {
	level, err := log.ParseLevel(prop.LogLevel)

	if err != nil {
		logger.SetLevel(log.TraceLevel)
	} else {
		logger.SetLevel(level)
		if level == log.TraceLevel {
			logger.SetReportCaller(true)
		}
	}

	if prop.Verbose {
		logger.SetLevel(log.TraceLevel)
		logger.SetReportCaller(true)
	}

	if !prop.Verbose && prop.LogLevel == "" {
		logger.SetLevel(log.FatalLevel)
	}

	logger.SetFormatter(&log.TextFormatter{
		PadLevelText:              true,
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		DisableTimestamp:          true,
	})

	logger.SetOutput(os.Stderr)

	logger.Infof("logger init success to => %s\n", logger.Level.String())
}
