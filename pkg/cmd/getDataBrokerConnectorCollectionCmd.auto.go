// Code generated from specification version 1.0.0: DO NOT EDIT
package cmd

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/reubenmiller/go-c8y-cli/pkg/mapbuilder"
	"github.com/reubenmiller/go-c8y/pkg/c8y"
	"github.com/spf13/cobra"
)

type getDataBrokerConnectorCollectionCmd struct {
	*baseCmd
}

func newGetDataBrokerConnectorCollectionCmd() *getDataBrokerConnectorCollectionCmd {
	ccmd := &getDataBrokerConnectorCollectionCmd{}

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Get collection of data broker connectors",
		Long:  ``,
		Example: `
$ c8y databroker list
Get a list of data broker connectors
        `,
		PreRunE: nil,
		RunE:    ccmd.getDataBrokerConnectorCollection,
	}

	cmd.SilenceUsage = true

	// Required flags

	ccmd.baseCmd = newBaseCmd(cmd)

	return ccmd
}

func (n *getDataBrokerConnectorCollectionCmd) getDataBrokerConnectorCollection(cmd *cobra.Command, args []string) error {

	commonOptions, err := getCommonOptions(cmd)
	if err != nil {
		return newUserError(fmt.Sprintf("Failed to get common options. err=%s", err))
	}

	// query parameters
	queryValue := url.QueryEscape("")
	query := url.Values{}
	commonOptions.AddQueryParameters(&query)
	queryValue, err = url.QueryUnescape(query.Encode())

	if err != nil {
		return newSystemError("Invalid query parameter")
	}

	// headers
	headers := http.Header{}

	// form data
	formData := make(map[string]io.Reader)

	// body
	body := mapbuilder.NewMapBuilder()

	// path parameters
	pathParameters := make(map[string]string)

	path := replacePathParameters("/databroker/connectors", pathParameters)

	req := c8y.RequestOptions{
		Method:       "GET",
		Path:         path,
		Query:        queryValue,
		Body:         body.GetMap(),
		FormData:     formData,
		Header:       headers,
		IgnoreAccept: false,
		DryRun:       globalFlagDryRun,
	}

	return processRequestAndResponse([]c8y.RequestOptions{req}, commonOptions)
}
