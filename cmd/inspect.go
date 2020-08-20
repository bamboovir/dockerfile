package cmd

import (
	"os"

	"github.com/bamboovir/dockerfile/cmd/types"
	"github.com/bamboovir/dockerfile/lib/cli"
	"github.com/bamboovir/dockerfile/lib/inspect"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// InspectProp define
type InspectProp struct {
	cli.Meta
	*types.InspectArgs
}

// InspectArgsCollector defination
func InspectArgsCollector(cmd *cobra.Command) *types.InspectArgs {
	props := &types.InspectArgs{}
	cmd.Flags().StringVar(&props.Path, "path", "", "dockerfile path / dir")
	cmd.Flags().StringVarP(&props.Formater, "formater", "f", "{{.From}}", "formater")
	return props
}

// InspectPropCollector defination
func InspectPropCollector(props *types.InspectArgs) (*InspectProp, error) {
	prop := &InspectProp{}
	prop.Meta = cli.DefaultMeta()
	prop.InspectArgs = props
	return prop, nil
}

// NewInspectCMD defination
func NewInspectCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "inspect",
		Short: "inspect",
	}

	args := InspectArgsCollector(cmd)

	cmd.PreRun = func(cmd *cobra.Command, args []string) {
		cmd.Parent().PreRun(cmd.Parent(), args)
	}

	cmd.Run = func(cmd *cobra.Command, _ []string) {
		props, err := InspectPropCollector(args)
		if err != nil {
			os.Exit(1)
		}

		err = Inspect(props)
		if err != nil {
			os.Exit(1)
		}
		os.Exit(0)
	}

	return cmd
}

// Inspect defination
func Inspect(props *InspectProp) error {
	errmsg := "inspect failed"

	rst, err := inspect.Inspect(props.FS, props.Path, props.Formater)
	if err != nil {
		log.Errorf("%v, %s\n", err, errmsg)
		return errors.Wrap(err, errmsg)
	}

	for _, s := range rst {
		props.OutWriter.Write([]byte(s))
		props.OutWriter.Write([]byte("\n"))
	}
	
	return nil
}
