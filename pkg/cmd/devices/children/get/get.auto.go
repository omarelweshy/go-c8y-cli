// Code generated from specification version 1.0.0: DO NOT EDIT
package get

import (
	"fmt"
	"io"
	"net/http"

	"github.com/MakeNowJust/heredoc/v2"
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

// GetCmd command
type GetCmd struct {
	*subcommand.SubCommand

	factory *cmdutil.Factory
}

// NewGetCmd creates a command to Get child
func NewGetCmd(f *cmdutil.Factory) *GetCmd {
	ccmd := &GetCmd{
		factory: f,
	}
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get child",
		Long:  `Get a child of a device`,
		Example: heredoc.Doc(`
$ c8y devices children get --id 12345 --child 12345 --childType addition
Get an existing child addition reference
        `),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		RunE: ccmd.RunE,
	}

	cmd.SilenceUsage = true

	cmd.Flags().StringSlice("id", []string{""}, "Managed object id (required) (accepts pipeline)")
	cmd.Flags().String("childType", "", "Child relationship type (required)")
	cmd.Flags().StringSlice("child", []string{""}, "Child managed object id (required)")

	completion.WithOptions(
		cmd,
		completion.WithDevice("id", func() (*c8y.Client, error) { return ccmd.factory.Client() }),
		completion.WithValidateSet("childType", "addition", "asset", "device"),
		completion.WithDevice("child", func() (*c8y.Client, error) { return ccmd.factory.Client() }),
	)

	flags.WithOptions(
		cmd,

		flags.WithExtendedPipelineSupport("id", "id", true, "deviceId", "source.id", "managedObject.id", "id"),
		flags.WithPipelineAliases("id", "deviceId", "source.id", "managedObject.id", "id"),
		flags.WithPipelineAliases("child", "deviceId", "source.id", "managedObject.id", "id"),

		flags.WithCollectionProperty("managedObject"),
	)

	// Required flags
	_ = cmd.MarkFlagRequired("childType")
	_ = cmd.MarkFlagRequired("child")

	ccmd.SubCommand = subcommand.NewSubCommand(cmd)

	return ccmd
}

// RunE executes the command
func (n *GetCmd) RunE(cmd *cobra.Command, args []string) error {
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
	commonOptions, err := cfg.GetOutputCommonOptions(cmd)
	if err != nil {
		return cmderrors.NewUserError(fmt.Sprintf("Failed to get common options. err=%s", err))
	}
	commonOptions.AddQueryParameters(query)

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
	)
	if err != nil {
		return cmderrors.NewUserError(err)
	}

	// form data
	formData := make(map[string]io.Reader)
	err = flags.WithFormDataOptions(
		cmd,
		formData,
		inputIterators,
	)
	if err != nil {
		return cmderrors.NewUserError(err)
	}

	// body
	body := mapbuilder.NewInitializedMapBuilder(false)
	err = flags.WithBody(
		cmd,
		body,
		inputIterators,
	)
	if err != nil {
		return cmderrors.NewUserError(err)
	}

	// path parameters
	path := flags.NewStringTemplate("inventory/managedObjects/{id}/{childType}/{child}")
	err = flags.WithPathParameters(
		cmd,
		path,
		inputIterators,
		c8yfetcher.WithDeviceByNameFirstMatch(n.factory, args, "id", "id"),
		flags.WithInventoryChildType("childType", "childType"),
		c8yfetcher.WithDeviceByNameFirstMatch(n.factory, args, "child", "child"),
	)
	if err != nil {
		return err
	}

	req := c8y.RequestOptions{
		Method:       "GET",
		Path:         path.GetTemplate(),
		Query:        queryValue,
		Body:         body,
		FormData:     formData,
		Header:       headers,
		IgnoreAccept: cfg.IgnoreAcceptHeader(),
		DryRun:       cfg.ShouldUseDryRun(cmd.CommandPath()),
	}

	return n.factory.RunWithWorkers(client, cmd, &req, inputIterators)
}
