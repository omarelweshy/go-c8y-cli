package cmd

import (
	"github.com/MakeNowJust/heredoc/v2"
	"github.com/reubenmiller/go-c8y-cli/pkg/cmd/subcommand"
	"github.com/reubenmiller/go-c8y-cli/pkg/cmderrors"
	"github.com/reubenmiller/go-c8y/pkg/c8y"
	"github.com/spf13/cobra"
)

type subscribeManagedObjectCmd struct {
	*subcommand.SubCommand

	flagDurationSec int64
	flagCount       int64
}

func NewSubscribeManagedObjectCmd() *subscribeManagedObjectCmd {
	ccmd := &subscribeManagedObjectCmd{}

	cmd := &cobra.Command{
		Use:   "subscribe",
		Short: "Subscribe to realtime managedObjects",
		Long:  `Subscribe to realtime managedObjects`,
		Example: heredoc.Doc(`
$ c8y inventory subscribe --device 12345
Subscribe to managedObjects (in realtime) for device 12345

$ c8y inventory subscribe --device 12345 --duration 30
Subscribe to managedObjects (in realtime) for device 12345 for 30 seconds

$ c8y inventory subscribe --count 10
Subscribe to managedObjects (in realtime) for all devices, and stop after receiving 10 managedObjects
		`),
		RunE: ccmd.subscribeManagedObject,
	}

	cmd.SilenceUsage = true

	cmd.Flags().StringSlice("device", []string{""}, "Device ID")
	cmd.Flags().Int64Var(&ccmd.flagDurationSec, "duration", 30, "Timeout in seconds")
	cmd.Flags().Int64Var(&ccmd.flagCount, "count", 0, "Max number of realtime notifications to wait for")

	// Required flags

	ccmd.SubCommand = subcommand.NewSubCommand(cmd)

	return ccmd
}

func (n *subscribeManagedObjectCmd) subscribeManagedObject(cmd *cobra.Command, args []string) error {

	// options
	device := "*"

	if cmd.Flags().Changed("device") {
		deviceInputValues, deviceValue, err := getFormattedDeviceSlice(cmd, args, "device")

		if err != nil {
			return cmderrors.NewUserError("no matching devices found", deviceInputValues, err)
		}

		if len(deviceValue) == 0 {
			return cmderrors.NewUserError("no matching devices found", deviceInputValues)
		}

		for _, item := range deviceValue {
			if item != "" {
				device = newIDValue(item).GetID()
			}
		}
	}

	return subscribe(c8y.RealtimeManagedObjects(device), n.flagDurationSec, n.flagCount, cmd)
}
