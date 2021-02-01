// Code generated from specification version 1.0.0: DO NOT EDIT
package cmd

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/reubenmiller/go-c8y-cli/pkg/flags"
	"github.com/reubenmiller/go-c8y-cli/pkg/mapbuilder"
	"github.com/reubenmiller/go-c8y/pkg/c8y"
	"github.com/spf13/cobra"
)

type NewRetentionRuleCmd struct {
	*baseCmd
}

func NewNewRetentionRuleCmd() *NewRetentionRuleCmd {
	ccmd := &NewRetentionRuleCmd{}
	cmd := &cobra.Command{
		Use:   "create",
		Short: "New retention rule",
		Long: `Create a new retention rule to managed when data is deleted in the tenant
`,
		Example: `
$ c8y retentionRules create --dataType ALARM --maximumAge 180
Create a retention rule
        `,
		PreRunE: validateCreateMode,
		RunE:    ccmd.RunE,
	}

	cmd.SilenceUsage = true

	cmd.Flags().String("dataType", "", "RetentionRule will be applied to this type of documents, possible values [ALARM, AUDIT, EVENT, MEASUREMENT, OPERATION, *]. (required)")
	cmd.Flags().String("fragmentType", "", "RetentionRule will be applied to documents with fragmentType.")
	cmd.Flags().String("type", "", "RetentionRule will be applied to documents with type.")
	cmd.Flags().String("source", "", "RetentionRule will be applied to documents with source.")
	cmd.Flags().Int("maximumAge", 0, "Maximum age of document in days. (required)")
	cmd.Flags().Bool("editable", false, "Whether the rule is editable. Can be updated only by management tenant.")
	addProcessingModeFlag(cmd)

	flags.WithOptions(
		cmd,
		flags.WithPipelineSupport(""),
	)

	// Required flags
	cmd.MarkFlagRequired("dataType")
	cmd.MarkFlagRequired("maximumAge")

	ccmd.baseCmd = newBaseCmd(cmd)

	return ccmd
}

func (n *NewRetentionRuleCmd) RunE(cmd *cobra.Command, args []string) error {
	// query parameters
	queryValue := url.QueryEscape("")
	query := url.Values{}

	err := flags.WithQueryOptions(
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

	// form data
	formData := make(map[string]io.Reader)

	// body
	body := mapbuilder.NewInitializedMapBuilder()
	body.SetMap(getDataFlag(cmd))
	if v, err := cmd.Flags().GetString("dataType"); err == nil {
		if v != "" {
			body.Set("dataType", v)
		}
	} else {
		return newUserError(fmt.Sprintf("Flag [%s] does not exist. %s", "dataType", err))
	}
	if v, err := cmd.Flags().GetString("fragmentType"); err == nil {
		if v != "" {
			body.Set("fragmentType", v)
		}
	} else {
		return newUserError(fmt.Sprintf("Flag [%s] does not exist. %s", "fragmentType", err))
	}
	if v, err := cmd.Flags().GetString("type"); err == nil {
		if v != "" {
			body.Set("type", v)
		}
	} else {
		return newUserError(fmt.Sprintf("Flag [%s] does not exist. %s", "type", err))
	}
	if v, err := cmd.Flags().GetString("source"); err == nil {
		if v != "" {
			body.Set("source", v)
		}
	} else {
		return newUserError(fmt.Sprintf("Flag [%s] does not exist. %s", "source", err))
	}
	if v, err := cmd.Flags().GetInt("maximumAge"); err == nil {
		body.Set("maximumAge", v)
	} else {
		return newUserError(fmt.Sprintf("Flag [%s] does not exist. %s", "maximumAge", err))
	}
	if cmd.Flags().Changed("editable") {
		if v, err := cmd.Flags().GetBool("editable"); err == nil {
			body.Set("editable", v)
		} else {
			return newUserError("Flag does not exist")
		}
	}
	if err := setDataTemplateFromFlags(cmd, body); err != nil {
		return newUserError("Template error. ", err)
	}
	if err := body.Validate(); err != nil {
		return newUserError("Body validation error. ", err)
	}

	// path parameters
	pathParameters := make(map[string]string)

	path := replacePathParameters("/retention/retentions", pathParameters)

	req := c8y.RequestOptions{
		Method:       "POST",
		Path:         path,
		Query:        queryValue,
		Body:         body.GetMap(),
		FormData:     formData,
		Header:       headers,
		IgnoreAccept: false,
		DryRun:       globalFlagDryRun,
	}

	return processRequestAndResponseWithWorkers(cmd, &req, "")
}
