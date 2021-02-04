// Code generated from specification version 1.0.0: DO NOT EDIT
package cmd

import (
	"io"
	"net/http"
	"net/url"

	"github.com/reubenmiller/go-c8y-cli/pkg/flags"
	"github.com/reubenmiller/go-c8y-cli/pkg/mapbuilder"
	"github.com/reubenmiller/go-c8y/pkg/c8y"
	"github.com/spf13/cobra"
)

type NewAuditCmd struct {
	*baseCmd
}

func NewNewAuditCmd() *NewAuditCmd {
	ccmd := &NewAuditCmd{}
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new audit record",
		Long:  `Create a new audit record for a given action`,
		Example: `
$ c8y auditRecords create --type "ManagedObject" --time "0s" --text "Managed Object updated: my_Prop: value" --source $Device.id --activity "Managed Object updated" --severity "information"
Create an audit record for a custom managed object update
        `,
		PreRunE: validateCreateMode,
		RunE:    ccmd.RunE,
	}

	cmd.SilenceUsage = true

	cmd.Flags().String("type", "", "Identifies the type of this audit record. (required)")
	cmd.Flags().String("time", "0s", "Time of the audit record. Defaults to current timestamp.")
	cmd.Flags().String("text", "", "Text description of the audit record. (required)")
	cmd.Flags().String("source", "", "An optional ManagedObject that the audit record originated from (required)")
	cmd.Flags().String("activity", "", "The activity that was carried out. (required)")
	cmd.Flags().String("severity", "", "The severity of action: critical, major, minor, warning or information. (required)")
	cmd.Flags().String("user", "", "The user responsible for the audited action.")
	cmd.Flags().String("application", "", "The application used to carry out the audited action.")
	addDataFlag(cmd)
	addProcessingModeFlag(cmd)

	flags.WithOptions(
		cmd,
		flags.WithPipelineSupport(""),
	)

	// Required flags
	cmd.MarkFlagRequired("type")
	cmd.MarkFlagRequired("text")
	cmd.MarkFlagRequired("source")
	cmd.MarkFlagRequired("activity")
	cmd.MarkFlagRequired("severity")

	ccmd.baseCmd = newBaseCmd(cmd)

	return ccmd
}

func (n *NewAuditCmd) RunE(cmd *cobra.Command, args []string) error {
	var err error
	// query parameters
	queryValue := url.QueryEscape("")
	query := url.Values{}

	err = flags.WithQueryParameters(
		cmd,
		query,
	)
	if err != nil {
		return newUserError(err)
	}
	err = flags.WithQueryOptions(
		cmd,
		query,
	)
	if err != nil {
		return newUserError(err)
	}

	queryValue, err = url.QueryUnescape(query.Encode())

	if err != nil {
		return newSystemError("Invalid query parameter")
	}

	// headers
	headers := http.Header{}
	if cmd.Flags().Changed("processingMode") {
		if v, err := cmd.Flags().GetString("processingMode"); err == nil && v != "" {
			headers.Add("X-Cumulocity-Processing-Mode", v)
		}
	}

	err = flags.WithHeaders(
		cmd,
		headers,
	)
	if err != nil {
		return newUserError(err)
	}

	// form data
	formData := make(map[string]io.Reader)

	// body
	body := mapbuilder.NewInitializedMapBuilder()
	err = flags.WithBody(
		cmd,
		body,
		flags.WithDataValue(FlagDataName),
		flags.WithStringValue("type", "type"),
		flags.WithRelativeTimestamp("time", "time", ""),
		flags.WithStringValue("text", "text"),
		flags.WithStringValue("source", "source.id"),
		flags.WithStringValue("activity", "activity"),
		flags.WithStringValue("severity", "severity"),
		flags.WithStringValue("user", "user"),
		flags.WithStringValue("application", "application"),
	)
	if err != nil {
		return newUserError(err)
	}

	if err := setLazyDataTemplateFromFlags(cmd, body); err != nil {
		return newUserError("Template error. ", err)
	}
	if err := body.Validate(); err != nil {
		return newUserError("Body validation error. ", err)
	}

	// path parameters
	pathParameters := make(map[string]string)
	err = flags.WithPathParameters(
		cmd,
		pathParameters,
	)

	path := replacePathParameters("/audit/auditRecords", pathParameters)

	req := c8y.RequestOptions{
		Method:       "POST",
		Path:         path,
		Query:        queryValue,
		Body:         body,
		FormData:     formData,
		Header:       headers,
		IgnoreAccept: false,
		DryRun:       globalFlagDryRun,
	}

	return processRequestAndResponseWithWorkers(cmd, &req, PipeOption{"", false})
}
