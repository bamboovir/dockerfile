package main

import (
	"fmt"
	"os"

	"github.com/bamboovir/dockerfile/cmd"
	"github.com/bamboovir/dockerfile/cmd/store"
	"github.com/bamboovir/dockerfile/cmd/types"
	log "github.com/sirupsen/logrus"
)

func main() {
	state := &types.State{
		Logger: log.StandardLogger(),
	}

	rootCMD := cmd.NewRootCMD(os.Args[1:], state, store.NewBasicStore())

	if err := rootCMD.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
