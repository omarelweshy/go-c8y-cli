// Code generated from specification version 1.0.0: DO NOT EDIT
package cmd

import (
	"io"
	"net/http"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/reubenmiller/go-c8y-cli/pkg/cmd/subcommand"
	"github.com/reubenmiller/go-c8y-cli/pkg/cmderrors"
	"github.com/reubenmiller/go-c8y-cli/pkg/completion"
	"github.com/reubenmiller/go-c8y-cli/pkg/flags"
	"github.com/reubenmiller/go-c8y-cli/pkg/mapbuilder"
	"github.com/reubenmiller/go-c8y/pkg/c8y"
	"github.com/spf13/cobra"
)

// UpdateRetentionRuleCmd command
type UpdateRetentionRuleCmd struct {
	*subcommand.SubCommand
}

// NewUpdateRetentionRuleCmd creates a command to Update retention rule
func NewUpdateRetentionRuleCmd() *UpdateRetentionRuleCmd {
	ccmd := &UpdateRetentionRuleCmd{}
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update retention rule",
		Long: `Update an existing retention rule, i.e. change maximum number of days or the data type.
`,
		Example: heredoc.Doc(`
$ c8y retentionRules get --id 12345
Update a retention rule
        `),
		PreRunE: validateUpdateMode,
		RunE:    ccmd.RunE,
	}

	cmd.SilenceUsage = true

	cmd.Flags().String("id", "", "Retention rule id (required) (accepts pipeline)")
	cmd.Flags().String("dataType", "", "RetentionRule will be applied to this type of documents, possible values [ALARM, AUDIT, EVENT, MEASUREMENT, OPERATION, *]. (required)")
	cmd.Flags().String("fragmentType", "", "RetentionRule will be applied to documents with fragmentType.")
	cmd.Flags().String("type", "", "RetentionRule will be applied to documents with type.")
	cmd.Flags().String("source", "", "RetentionRule will be applied to documents with source.")
	cmd.Flags().Int("maximumAge", 0, "Maximum age of document in days.")
	cmd.Flags().Bool("editable", false, "Whether the rule is editable. Can be updated only by management tenant.")
	addDataFlag(cmd)
	addProcessingModeFlag(cmd)

	completion.WithOptions(
		cmd,
		completion.WithValidateSet("dataType", "ALARM", "AUDIT", "EVENT", "MEASUREMENT", "OPERATION", "*"),
	)

	flags.WithOptions(
		cmd,
		flags.WithExtendedPipelineSupport("id", "id", true),
	)

	// Required flags
	_ = cmd.MarkFlagRequired("dataType")

	ccmd.SubCommand = subcommand.NewSubCommand(cmd)

	return ccmd
}

// RunE executes the command
func (n *UpdateRetentionRuleCmd) RunE(cmd *cobra.Command, args []string) error {
	var err error
	inputIterators, err := flags.NewRequestInputIterators(cmd)
	if err != nil {
		return err
	}

	// query parameters
	query := flags.NewQueryTemplate()
	err = flags.WithQueryParameters(
		cmd,
		query,
		inputIterators,
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
		inputIterators,
	)
	if err != nil {
		return cmderrors.NewUserError(err)
	}

	// body
	body := mapbuilder.NewInitializedMapBuilder()
	err = flags.WithBody(
		cmd,
		body,
		inputIterators,
		WithDataValue(),
		flags.WithStringValue("dataType", "dataType"),
		flags.WithStringValue("fragmentType", "fragmentType"),
		flags.WithStringValue("type", "type"),
		flags.WithStringValue("source", "source"),
		flags.WithIntValue("maximumAge", "maximumAge"),
		flags.WithBoolValue("editable", "editable", ""),
		WithTemplateValue(),
		WithTemplateVariablesValue(),
	)
	if err != nil {
		return cmderrors.NewUserError(err)
	}

	// path parameters
	path := flags.NewStringTemplate("/retention/retentions/{id}")
	err = flags.WithPathParameters(
		cmd,
		path,
		inputIterators,
		flags.WithStringValue("id", "id"),
	)
	if err != nil {
		return err
	}

	req := c8y.RequestOptions{
		Method:       "PUT",
		Path:         path.GetTemplate(),
		Query:        queryValue,
		Body:         body,
		FormData:     formData,
		Header:       headers,
		IgnoreAccept: cliConfig.IgnoreAcceptHeader(),
		DryRun:       cliConfig.DryRun(),
	}

	return processRequestAndResponseWithWorkers(cmd, &req, inputIterators)
}
