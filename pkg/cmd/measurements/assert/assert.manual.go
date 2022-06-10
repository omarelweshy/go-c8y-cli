package assert

import (
	cmdCount "github.com/reubenmiller/go-c8y-cli/v2/pkg/cmd/measurements/assert/count"
	"github.com/reubenmiller/go-c8y-cli/v2/pkg/cmd/subcommand"
	"github.com/reubenmiller/go-c8y-cli/v2/pkg/cmdutil"
	"github.com/spf13/cobra"
)

type SubCmdAssert struct {
	*subcommand.SubCommand
}

func NewSubCommand(f *cmdutil.Factory) *SubCmdAssert {
	ccmd := &SubCmdAssert{}

	cmd := &cobra.Command{
		Use:   "assert",
		Short: "Measurement assertions",
		Long:  `Assertions related to Cumulocity measurements`,
	}

	// Subcommands
	cmd.AddCommand(cmdCount.NewCmdCount(f).GetCommand())

	ccmd.SubCommand = subcommand.NewSubCommand(cmd)

	return ccmd
}
