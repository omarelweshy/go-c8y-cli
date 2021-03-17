package cmd

import (
	"context"
	"fmt"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/reubenmiller/go-c8y-cli/pkg/cmd/subcommand"
	"github.com/reubenmiller/go-c8y-cli/pkg/cmderrors"
	"github.com/reubenmiller/go-c8y/pkg/c8y"
	"github.com/spf13/cobra"
)

type newServiceUserCmd struct {
	*subcommand.SubCommand

	name    string
	tenants []string
	roles   []string
}

func NewNewServiceUserCmd() *newServiceUserCmd {
	ccmd := &newServiceUserCmd{}

	cmd := &cobra.Command{
		Use:   "createServiceUser",
		Short: "New application service user",
		Long:  ``,
		Example: heredoc.Doc(`
$ c8y microservices createServiceUser --name my-user
Create new application service user
		`),
		PreRunE: validateCreateMode,
		RunE:    ccmd.doProcedure,
	}

	cmd.SilenceUsage = true

	addDataFlagWithoutTemplates(cmd)
	cmd.Flags().StringVar(&ccmd.name, "name", "", "Name of application")
	cmd.Flags().StringSliceVar(&ccmd.tenants, "tenants", []string{}, "Tenant to subscribe to. If left blank than the application will not generate the service user")
	cmd.Flags().StringSliceVar(&ccmd.roles, "roles", []string{}, "Roles which should be assigned to the service user")

	// Required flags
	_ = cmd.MarkFlagRequired("name")

	ccmd.SubCommand = subcommand.NewSubCommand(cmd)

	return ccmd
}

func (n *newServiceUserCmd) getApplicationDetails() *c8y.Application {

	app := c8y.Application{}

	// Set application properties
	app.Name = n.name
	app.Key = app.Name + "-key"
	app.Type = "MICROSERVICE"
	app.ContextPath = app.Name

	if len(n.roles) > 0 {
		app.RequiredRoles = n.roles
	}

	return &app
}

func (n *newServiceUserCmd) doProcedure(cmd *cobra.Command, args []string) error {
	var application *c8y.Application
	var response *c8y.Response
	var err error

	applicationDetails := n.getApplicationDetails()

	if applicationDetails.Name == "" {
		return cmderrors.NewUserError("Could not detect application name for the given input")
	}

	Logger.Info("Creating new application")
	application, response, err = client.Application.Create(context.Background(), applicationDetails)

	if err != nil {
		return fmt.Errorf("failed to create microservice. %s", err)
	}

	// App subscription
	if len(n.tenants) > 0 {
		for _, tenant := range n.tenants {
			Logger.Infof("Subscribing to application in tenant %s", tenant)
			_, resp, err := client.Tenant.AddApplicationReference(context.Background(), tenant, application.Self)

			if err != nil {
				if resp != nil && resp.StatusCode == 409 {
					Logger.Infof("microservice is already enabled")
				} else {
					return fmt.Errorf("Failed to subscribe to application. %s", err)
				}
			}
		}
	}

	commonOptions, err := cliConfig.GetOutputCommonOptions(cmd)
	if err != nil {
		return err
	}

	_, err = processResponse(response, err, commonOptions)
	return err
}
