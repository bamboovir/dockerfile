package types

import (
	log "github.com/sirupsen/logrus"
)

// State defination
type State struct {
	// RootArgs
	RootArgs *RootArgs
	// Logger
	Logger *log.Logger
}

// RootArgs defination
type RootArgs struct {
	// Verbose
	Verbose bool
	// LogLevel
	LogLevel string
}

// InspectArgs defination
type InspectArgs struct {
	// Path
	Path     string
	Formater string
}
