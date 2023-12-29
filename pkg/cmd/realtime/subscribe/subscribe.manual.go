package subscribe

import (
	"time"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/reubenmiller/go-c8y-cli/v2/pkg/c8ysubscribe"
	"github.com/reubenmiller/go-c8y-cli/v2/pkg/cmd/subcommand"
	"github.com/reubenmiller/go-c8y-cli/v2/pkg/cmdutil"
	"github.com/reubenmiller/go-c8y-cli/v2/pkg/completion"
	"github.com/reubenmiller/go-c8y-cli/v2/pkg/flags"
	"github.com/reubenmiller/go-c8y/pkg/c8y"
	"github.com/spf13/cobra"
)

type CmdSubscribe struct {
	flagChannel string
	flagCount   int64
	actionTypes []string

	*subcommand.SubCommand

	factory *cmdutil.Factory
}

func NewCmdSubscribe(f *cmdutil.Factory) *CmdSubscribe {
	ccmd := &CmdSubscribe{
		factory: f,
	}

	cmd := &cobra.Command{
		Use:   "subscribe",
		Short: "Subscribe to realtime notifications",
		Long:  `Subscribe to realtime notifications`,
		Example: heredoc.Doc(`
$ c8y realtime subscribe --channel "/measurements/*" --duration 90s

Subscribe to all measurements for 90 seconds
		`),
		RunE: ccmd.RunE,
	}

	// Flags
	cmd.Flags().StringVar(&ccmd.flagChannel, "channel", "", "Channel name i.e. \"/measurements/12345\" or \"/measurements/*\"")
	cmd.Flags().String("duration", "30s", "Duration to subscribe for. i.e. 30s, 1m")
	cmd.Flags().Int64Var(&ccmd.flagCount, "count", 0, "Max number of realtime notifications to wait for")
	cmd.Flags().StringSliceVar(&ccmd.actionTypes, "actionTypes", nil, "Filter by realtime action types, i.e. CREATE,UPDATE,DELETE")

	completion.WithOptions(
		cmd,
		completion.WithValidateSet("actionTypes", "CREATE", "UPDATE", "DELETE"),
		completion.WithValidateSet(
			"channel",
			c8y.RealtimeAlarms("*"),
			c8y.RealtimeAlarmsWithChildren("*"),
			c8y.RealtimeEvents("*"),
			c8y.RealtimeManagedObjects("*"),
			c8y.RealtimeMeasurements("*"),
			c8y.RealtimeOperations("*"),
		),
	)

	ccmd.SubCommand = subcommand.NewSubCommand(cmd)

	return ccmd
}

func (n *CmdSubscribe) RunE(cmd *cobra.Command, args []string) error {
	client, err := n.factory.Client()
	if err != nil {
		return err
	}
	cfg, err := n.factory.Config()
	if err != nil {
		return err
	}
	log, err := n.factory.Logger()
	if err != nil {
		return err
	}
	duration, err := flags.GetDurationFlag(cmd, "duration", true, time.Second)
	if err != nil {
		return err
	}
	opts := c8ysubscribe.Options{
		Timeout:     duration,
		MaxMessages: n.flagCount,
		ActionTypes: n.actionTypes,
		OnMessage: func(msg string) error {
			return n.factory.WriteJSONToConsole(cfg, cmd, "", []byte(msg))
		},
	}
	return c8ysubscribe.Subscribe(client, log, n.flagChannel, opts)
}
