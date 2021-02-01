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

type NewApplicationCmd struct {
	*baseCmd
}

func NewNewApplicationCmd() *NewApplicationCmd {
	ccmd := &NewApplicationCmd{}
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new application",
		Long:  `Create a new application using explicit settings`,
		Example: `
$ c8y applications create --name myapp --type HOSTED --key "myapp-key" --contextPath "myapp"
Create a new hosted application
        `,
		PreRunE: validateCreateMode,
		RunE:    ccmd.RunE,
	}

	cmd.SilenceUsage = true

	addDataFlag(cmd)
	cmd.Flags().String("name", "", "Name of application (required)")
	cmd.Flags().String("key", "", "Shared secret of application (required)")
	cmd.Flags().String("type", "", "Type of application. Possible values are EXTERNAL, HOSTED, MICROSERVICE (required)")
	cmd.Flags().String("availability", "", "Access level for other tenants. Possible values are : MARKET, PRIVATE (default)")
	cmd.Flags().String("contextPath", "", "contextPath of the hosted application. Required when application type is HOSTED")
	cmd.Flags().String("resourcesUrl", "", "URL to application base directory hosted on an external server. Required when application type is HOSTED")
	cmd.Flags().String("resourcesUsername", "", "authorization username to access resourcesUrl")
	cmd.Flags().String("resourcesPassword", "", "authorization password to access resourcesUrl")
	cmd.Flags().String("externalUrl", "", "URL to the external application. Required when application type is EXTERNAL")
	addProcessingModeFlag(cmd)

	flags.WithOptions(
		cmd,
		flags.WithPipelineSupport(""),
	)

	// Required flags
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("key")
	cmd.MarkFlagRequired("type")

	ccmd.baseCmd = newBaseCmd(cmd)

	return ccmd
}

func (n *NewApplicationCmd) RunE(cmd *cobra.Command, args []string) error {
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
	if v, err := cmd.Flags().GetString("name"); err == nil {
		if v != "" {
			body.Set("name", v)
		}
	} else {
		return newUserError(fmt.Sprintf("Flag [%s] does not exist. %s", "name", err))
	}
	if v, err := cmd.Flags().GetString("key"); err == nil {
		if v != "" {
			body.Set("key", v)
		}
	} else {
		return newUserError(fmt.Sprintf("Flag [%s] does not exist. %s", "key", err))
	}
	if v, err := cmd.Flags().GetString("type"); err == nil {
		if v != "" {
			body.Set("type", v)
		}
	} else {
		return newUserError(fmt.Sprintf("Flag [%s] does not exist. %s", "type", err))
	}
	if v, err := cmd.Flags().GetString("availability"); err == nil {
		if v != "" {
			body.Set("availability", v)
		}
	} else {
		return newUserError(fmt.Sprintf("Flag [%s] does not exist. %s", "availability", err))
	}
	if v, err := cmd.Flags().GetString("contextPath"); err == nil {
		if v != "" {
			body.Set("contextPath", v)
		}
	} else {
		return newUserError(fmt.Sprintf("Flag [%s] does not exist. %s", "contextPath", err))
	}
	if v, err := cmd.Flags().GetString("resourcesUrl"); err == nil {
		if v != "" {
			body.Set("resourcesUrl", v)
		}
	} else {
		return newUserError(fmt.Sprintf("Flag [%s] does not exist. %s", "resourcesUrl", err))
	}
	if v, err := cmd.Flags().GetString("resourcesUsername"); err == nil {
		if v != "" {
			body.Set("resourcesUsername", v)
		}
	} else {
		return newUserError(fmt.Sprintf("Flag [%s] does not exist. %s", "resourcesUsername", err))
	}
	if v, err := cmd.Flags().GetString("resourcesPassword"); err == nil {
		if v != "" {
			body.Set("resourcesPassword", v)
		}
	} else {
		return newUserError(fmt.Sprintf("Flag [%s] does not exist. %s", "resourcesPassword", err))
	}
	if v, err := cmd.Flags().GetString("externalUrl"); err == nil {
		if v != "" {
			body.Set("externalUrl", v)
		}
	} else {
		return newUserError(fmt.Sprintf("Flag [%s] does not exist. %s", "externalUrl", err))
	}
	if err := setDataTemplateFromFlags(cmd, body); err != nil {
		return newUserError("Template error. ", err)
	}
	if err := body.Validate(); err != nil {
		return newUserError("Body validation error. ", err)
	}

	// path parameters
	pathParameters := make(map[string]string)

	path := replacePathParameters("/application/applications", pathParameters)

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
