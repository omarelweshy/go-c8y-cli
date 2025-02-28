// Code generated from specification version 1.0.0: DO NOT EDIT
package createbinary

import (
	"io"
	"net/http"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/reubenmiller/go-c8y-cli/v2/pkg/c8ybinary"
	"github.com/reubenmiller/go-c8y-cli/v2/pkg/c8yfetcher"
	"github.com/reubenmiller/go-c8y-cli/v2/pkg/cmd/subcommand"
	"github.com/reubenmiller/go-c8y-cli/v2/pkg/cmderrors"
	"github.com/reubenmiller/go-c8y-cli/v2/pkg/cmdutil"
	"github.com/reubenmiller/go-c8y-cli/v2/pkg/completion"
	"github.com/reubenmiller/go-c8y-cli/v2/pkg/flags"
	"github.com/reubenmiller/go-c8y-cli/v2/pkg/mapbuilder"
	"github.com/reubenmiller/go-c8y/pkg/c8y"
	"github.com/spf13/cobra"
)

// CreateBinaryCmd command
type CreateBinaryCmd struct {
	*subcommand.SubCommand

	factory *cmdutil.Factory
}

// NewCreateBinaryCmd creates a command to Create event binary
func NewCreateBinaryCmd(f *cmdutil.Factory) *CreateBinaryCmd {
	ccmd := &CreateBinaryCmd{
		factory: f,
	}
	cmd := &cobra.Command{
		Use:   "createBinary",
		Short: "Create event binary",
		Long:  `Upload a new binary file to an event`,
		Example: heredoc.Doc(`
$ c8y events createBinary --id 12345 --file ./myfile.log
Add a binary to an event

$ c8y events createBinary --id 12345 --file ./myfile.log --name "myfile-2022-03-31.log"
Add a binary to an event using a custom name
        `),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return f.CreateModeEnabled()
		},
		RunE: ccmd.RunE,
	}

	cmd.SilenceUsage = true

	cmd.Flags().StringSlice("id", []string{""}, "Event id (required) (accepts pipeline)")
	cmd.Flags().String("file", "", "File to be uploaded as a binary (required)")
	cmd.Flags().String("name", "", "Set the name of the binary file. This will be the name of the file when it is downloaded in the UI")

	completion.WithOptions(
		cmd,
	)

	flags.WithOptions(
		cmd,
		flags.WithProcessingMode(),

		flags.WithExtendedPipelineSupport("id", "id", true),
	)

	// Required flags
	_ = cmd.MarkFlagRequired("file")

	ccmd.SubCommand = subcommand.NewSubCommand(cmd)

	return ccmd
}

// RunE executes the command
func (n *CreateBinaryCmd) RunE(cmd *cobra.Command, args []string) error {
	cfg, err := n.factory.Config()
	if err != nil {
		return err
	}
	// Runtime flag options
	flags.WithOptions(
		cmd,
		flags.WithRuntimePipelineProperty(),
	)
	client, err := n.factory.Client()
	if err != nil {
		return err
	}
	inputIterators, err := cmdutil.NewRequestInputIterators(cmd, cfg)
	if err != nil {
		return err
	}

	// query parameters
	query := flags.NewQueryTemplate()
	err = flags.WithQueryParameters(
		cmd,
		query,
		inputIterators,
		flags.WithCustomStringSlice(func() ([]string, error) { return cfg.GetQueryParameters(), nil }, "custom"),
	)
	if err != nil {
		return cmderrors.NewUserError(err)
	}

	queryValue, err := query.GetQueryUnescape(true)

	if err != nil {
		return cmderrors.NewSystemError("Invalid query parameter")
	}

	// headers
	headers := http.Header{}
	err = flags.WithHeaders(
		cmd,
		headers,
		inputIterators,
		flags.WithCustomStringSlice(func() ([]string, error) { return cfg.GetHeader(), nil }, "header"),
		flags.WithProcessingModeValue(),
	)
	if err != nil {
		return cmderrors.NewUserError(err)
	}

	// form data
	formData := make(map[string]io.Reader)
	err = flags.WithFormDataOptions(
		cmd,
		formData,
		inputIterators, flags.WithOptionBuilder().
			Append(flags.WithFormDataFile("file", "data")...).
			Build()...,
	)
	if err != nil {
		return cmderrors.NewUserError(err)
	}

	// body
	body := mapbuilder.NewInitializedMapBuilder(true)
	err = flags.WithBody(
		cmd,
		body,
		inputIterators,
		flags.WithDataFlagValue(),
		flags.WithStringValue("name", "name"),
		cmdutil.WithTemplateValue(n.factory),
		flags.WithTemplateVariablesValue(),
	)
	if err != nil {
		return cmderrors.NewUserError(err)
	}

	// path parameters
	path := flags.NewStringTemplate("event/events/{id}/binaries")
	err = flags.WithPathParameters(
		cmd,
		path,
		inputIterators,
		c8yfetcher.WithIDSlice(args, "id", "id"),
	)
	if err != nil {
		return err
	}

	req := c8y.RequestOptions{
		Method:         "POST",
		Path:           path.GetTemplate(),
		Query:          queryValue,
		Body:           body,
		FormData:       formData,
		Header:         headers,
		IgnoreAccept:   cfg.IgnoreAcceptHeader(),
		DryRun:         cfg.ShouldUseDryRun(cmd.CommandPath()),
		PrepareRequest: c8ybinary.AddProgress(cmd, "file", cfg.GetProgressBar(n.factory.IOStreams.ErrOut, n.factory.IOStreams.IsStderrTTY())),
	}

	return n.factory.RunWithWorkers(client, cmd, &req, inputIterators)
}
