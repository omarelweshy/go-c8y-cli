package c8yfetcher

import (
	"context"
	"fmt"
	"regexp"

	"github.com/pkg/errors"
	"github.com/reubenmiller/go-c8y-cli/v2/pkg/cmdutil"
	"github.com/reubenmiller/go-c8y/pkg/c8y"
)

type UIPluginFetcher struct {
	*CumulocityFetcher

	// Look for shared plugins if no local ones are found
	EnableSharedPlugins bool
}

func NewUIPluginFetcher(factory *cmdutil.Factory) *UIPluginFetcher {
	return &UIPluginFetcher{
		CumulocityFetcher: &CumulocityFetcher{
			factory: factory,
		},
	}
}

func (f *UIPluginFetcher) getByID(id string) ([]fetcherResultSet, error) {
	app, resp, err := f.Client().Application.GetApplication(
		c8y.WithDisabledDryRunContext(context.Background()),
		id,
	)

	if err != nil {
		return nil, errors.Wrap(err, "Could not fetch by id")
	}

	results := make([]fetcherResultSet, 1)
	results[0] = fetcherResultSet{
		ID:    app.ID,
		Name:  app.Name,
		Value: resp.JSON(),
	}
	return results, nil
}

// getByName returns applications matching a given using regular expression
func (f *UIPluginFetcher) getByName(name string) ([]fetcherResultSet, error) {
	serverOptions := &c8y.ApplicationOptions{
		PaginationOptions: *c8y.NewPaginationOptions(2000),
		Type:              c8y.ApplicationTypeHosted,
	}
	serverOptions.WithHasVersions(true)
	if f.EnableSharedPlugins && f.Client().TenantName != "" {
		// Ignore microservices which don't match the owner
		// so that microservices of sub tenants don't get returned.
		serverOptions.Owner = f.Client().TenantName
	}

	col, _, err := f.Client().Application.GetApplications(
		c8y.WithDisabledDryRunContext(context.Background()),
		serverOptions,
	)

	if err != nil {
		return nil, errors.Wrap(err, "could not fetch ui plugin")
	}

	pattern, err := regexp.Compile("^" + regexp.QuoteMeta(name) + "$")

	if err != nil {
		return nil, errors.Wrap(err, "invalid regex")
	}

	results := make([]fetcherResultSet, 0)

	// Note: Match against both name and contextPath
	// as the contextPath is used by UI plugins as a reference
	for i, app := range col.Applications {
		if pattern.MatchString(app.Name) || pattern.MatchString(app.ContextPath) {
			results = append(results, fetcherResultSet{
				ID:    app.ID,
				Name:  app.Name,
				Value: col.Items[i],
			})
		}
	}

	// If not results are found, then also include any matches (not just in the current tenant)
	if len(results) == 0 && f.EnableSharedPlugins {

		// Run request against, but without the tenant filter
		serverOptions.Availability = c8y.ApplicationAvailabilityShared
		serverOptions.Owner = ""
		col, _, err := f.Client().Application.GetApplications(
			c8y.WithDisabledDryRunContext(context.Background()),
			serverOptions,
		)
		if err != nil {
			return nil, errors.Wrap(err, "could not fetch applications")
		}

		for i, app := range col.Applications {
			if pattern.MatchString(app.Name) || pattern.MatchString(app.ContextPath) {
				results = append(results, fetcherResultSet{
					ID:    app.ID,
					Name:  app.Name,
					Value: col.Items[i],
				})
			}
		}
	}

	return results, nil
}

// Find UI Plugins returns plugins given either an id or search text
// @values: An array of ids, or names (with wildcards)
// @lookupID: Lookup the data if an id is given. If a non-id text is given, the result will always be looked up.
func FindUIPlugins(factory *cmdutil.Factory, values []string, lookupID bool, format string, resolveSharedPlugins bool) ([]entityReference, error) {
	f := NewUIPluginFetcher(factory)
	f.EnableSharedPlugins = resolveSharedPlugins

	formattedValues, err := lookupEntity(f, values, lookupID, format)

	if err != nil {
		return nil, err
	}

	results := []entityReference{}

	invalidLookups := []string{}
	for _, item := range formattedValues {
		if item.ID != "" {
			results = append(results, item)
		} else {
			if item.Name != "" {
				invalidLookups = append(invalidLookups, item.Name)
			}
		}
	}

	var errors error

	if len(invalidLookups) > 0 {
		errors = fmt.Errorf("no results %v", invalidLookups)
	}

	return results, errors
}
