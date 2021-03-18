// TODO

package cmd

import (
	"github.com/MakeNowJust/heredoc/v2"
	"github.com/reubenmiller/go-c8y-cli/pkg/c8yfetcher"
	"github.com/reubenmiller/go-c8y-cli/pkg/cmd/subcommand"
	"github.com/reubenmiller/go-c8y-cli/pkg/flags"
	"github.com/reubenmiller/go-c8y/pkg/c8y"
	"github.com/spf13/cobra"
)

type subscribeMeasurementCmd struct {
	*subcommand.SubCommand

	flagDurationSec int64
	flagCount       int64
}

func NewSubscribeMeasurementCmd() *subscribeMeasurementCmd {
	ccmd := &subscribeMeasurementCmd{}

	cmd := &cobra.Command{
		Use:   "subscribe",
		Short: "Subscribe to realtime measurements",
		Long:  `Subscribe to realtime measurements`,
		Example: heredoc.Doc(`
$ c8y measurements subscribe --device 12345
Subscribe to measurements (in realtime) for device 12345

$ c8y measurements subscribe --device 12345 --duration 30
Subscribe to measurements (in realtime) for device 12345 for 30 seconds

$ c8y measurements subscribe --count 10
Subscribe to measurements (in realtime) for all devices, and stop after receiving 10 measurements
		`),
		RunE: ccmd.subscribeMeasurement,
	}

	cmd.SilenceUsage = true

	cmd.Flags().StringSlice("device", []string{""}, "Device ID")
	cmd.Flags().Int64Var(&ccmd.flagDurationSec, "duration", 30, "Timeout in seconds")
	cmd.Flags().Int64Var(&ccmd.flagCount, "count", 0, "Max number of realtime notifications to wait for")

	// Required flags

	ccmd.SubCommand = subcommand.NewSubCommand(cmd)

	return ccmd
}

func (n *subscribeMeasurementCmd) subscribeMeasurement(cmd *cobra.Command, args []string) error {

	inputIterators, err := flags.NewRequestInputIterators(cmd)
	if err != nil {
		return err
	}

	// path parameters
	path := flags.NewStringTemplate("{device}")
	err = flags.WithPathParameters(
		cmd,
		path,
		inputIterators,
		flags.WithStringDefaultValue("*", "device", "device"),
		c8yfetcher.WithDeviceByNameFirstMatch(client, args, "device", "device"),
	)
	if err != nil {
		return err
	}

	device, _, err := path.Execute(false)
	if err != nil {
		return err
	}

	return subscribe(c8y.RealtimeMeasurements(device), n.flagDurationSec, n.flagCount, cmd)
}
