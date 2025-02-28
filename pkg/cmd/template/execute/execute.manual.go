package execute

import (
	"bytes"
	"strconv"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/reubenmiller/go-c8y-cli/v2/pkg/cmd/subcommand"
	"github.com/reubenmiller/go-c8y-cli/v2/pkg/cmderrors"
	"github.com/reubenmiller/go-c8y-cli/v2/pkg/cmdutil"
	"github.com/reubenmiller/go-c8y-cli/v2/pkg/flags"
	"github.com/reubenmiller/go-c8y-cli/v2/pkg/iterator"
	"github.com/reubenmiller/go-c8y-cli/v2/pkg/mapbuilder"
	"github.com/reubenmiller/go-c8y-cli/v2/pkg/worker"
	"github.com/spf13/cobra"
)

type CmdExecute struct {
	*subcommand.SubCommand

	factory *cmdutil.Factory
}

func NewCmdExecute(f *cmdutil.Factory) *CmdExecute {
	ccmd := &CmdExecute{
		factory: f,
	}

	cmd := &cobra.Command{
		Use:   "execute",
		Short: "Execute a jsonnet template",
		Long:  `Execute a jsonnet template and return the output. Useful when creating new templates`,
		Example: heredoc.Doc(`
Example 1:
$ c8y template execute --template ./mytemplate.jsonnet

Verify a jsonnet template and show the output after it is evaluated

Example 2:
$ c8y template execute --template ./mytemplate.jsonnet --templateVars "name=testme" --data "name=myDeviceName"

Verify a jsonnet template and specify input data to be used as the input when evaluating the template

Example 3:
$ echo '{"name": "external_source"}' | c8y template execute --template "{name: input.value.name, nestedValues: input.value}"
$ => {"name":"external_source","nestedValues":{"name":"external_source"}}

Pass external json data into the template, and reference it via the "input.value" variable
		`),
		RunE: ccmd.newTemplate,
	}

	cmdutil.DisableEncryptionCheck(cmd)
	cmd.SilenceUsage = true

	cmd.Flags().String("input", "", "input (accepts pipeline)")

	flags.WithOptions(
		cmd,
		flags.WithData(),
		f.WithTemplateFlag(cmd),
		flags.WithExtendedPipelineSupport("input", "", false),
	)

	cmdutil.DisableAuthCheck(cmd)
	ccmd.SubCommand = subcommand.NewSubCommand(cmd).SetRequiredFlags(flags.FlagDataTemplateName)

	return ccmd
}

func formatOutput(response []byte) []byte {
	// Strip trailing newline (if present)
	if bytes.HasSuffix(response, []byte("\n")) {
		response = response[:len(response)-1]
	}

	// Unquote string to convert any escape sequences to their representation
	if bytes.HasPrefix(response, []byte("\"")) && bytes.HasSuffix(response, []byte("\"")) {
		if out, err := strconv.Unquote(string(response)); err == nil {
			response = []byte(out)
		}
	}
	return response
}

func (n *CmdExecute) newTemplate(cmd *cobra.Command, args []string) error {
	cfg, err := n.factory.Config()
	if err != nil {
		return err
	}
	inputIterators, err := cmdutil.NewRequestInputIterators(cmd, cfg)
	if err != nil {
		return err
	}

	if !cmd.Flags().Changed(flags.FlagDataTemplateName) {
		return &flags.ParameterError{
			Name: flags.FlagDataTemplateName,
			Err:  flags.ErrParameterMissing,
		}
	}

	// body
	body := mapbuilder.NewInitializedMapBuilder(true)
	err = flags.WithBody(
		cmd,
		body,
		inputIterators,
		flags.WithOverrideValue("input", "input"),
		flags.WithDataFlagValue(),
		cmdutil.WithTemplateValue(n.factory),
		flags.WithTemplateVariablesValue(),
		flags.WithStringValue("input", "input", ""),
	)

	if err != nil {
		return cmderrors.NewUserError(err)
	}

	var iter iterator.Iterator
	if inputIterators.Total > 0 {
		iter = mapbuilder.NewMapBuilderIterator(body)
	} else {
		iter = iterator.NewBoundIterator(mapbuilder.NewMapBuilderIterator(body), 1)
	}

	commonOptions, err := cfg.GetOutputCommonOptions(cmd)
	if err != nil {
		return err
	}
	commonOptions.DisableResultPropertyDetection()

	return n.factory.RunWithGenericWorkers(cmd, inputIterators, iter, func(j worker.Job) (any, error) {
		output := formatOutput(j.Value.([]byte))
		err := n.factory.WriteOutput(output, cmdutil.OutputContext{
			Input: j.Input,
		}, &commonOptions)
		return nil, err
	})
}
