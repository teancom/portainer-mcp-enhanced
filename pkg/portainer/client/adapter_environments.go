package client

import (
	"fmt"

	"github.com/portainer/client-api-go/v2/pkg/client/endpoints"
	"github.com/portainer/client-api-go/v2/pkg/client/tags"
	"github.com/portainer/client-api-go/v2/pkg/client/teams"
	"github.com/portainer/client-api-go/v2/pkg/client/users"
)

// DeleteTag deletes a tag by ID using the low-level Swagger client.
func (a *portainerAPIAdapter) DeleteTag(id int64) error {
	params := tags.NewTagDeleteParams().WithID(id)
	_, err := a.swagger.Tags.TagDelete(params, nil)
	if err != nil {
		return fmt.Errorf("failed to delete tag: %w", err)
	}
	return nil
}

// DeleteTeam deletes a team by ID using the low-level Swagger client.
func (a *portainerAPIAdapter) DeleteTeam(id int64) error {
	params := teams.NewTeamDeleteParams().WithID(id)
	_, err := a.swagger.Teams.TeamDelete(params, nil)
	if err != nil {
		return fmt.Errorf("failed to delete team: %w", err)
	}
	return nil
}

// DeleteUser deletes a user by ID using the low-level Swagger client.
func (a *portainerAPIAdapter) DeleteUser(id int64) error {
	params := users.NewUserDeleteParams().WithID(id)
	_, err := a.swagger.Users.UserDelete(params, nil)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

// DeleteEndpoint deletes an endpoint by ID using the low-level Swagger client.
func (a *portainerAPIAdapter) DeleteEndpoint(id int64) error {
	params := endpoints.NewEndpointDeleteParams().WithID(id)
	_, err := a.swagger.Endpoints.EndpointDelete(params, nil)
	if err != nil {
		return fmt.Errorf("failed to delete endpoint: %w", err)
	}
	return nil
}

// SnapshotEndpoint triggers a snapshot for a single endpoint.
func (a *portainerAPIAdapter) SnapshotEndpoint(id int64) error {
	params := endpoints.NewEndpointSnapshotParams().WithID(id)
	_, err := a.swagger.Endpoints.EndpointSnapshot(params, nil)
	if err != nil {
		return fmt.Errorf("failed to snapshot endpoint: %w", err)
	}
	return nil
}

// SnapshotAllEndpoints triggers a snapshot for all endpoints.
func (a *portainerAPIAdapter) SnapshotAllEndpoints() error {
	params := endpoints.NewEndpointSnapshotsParams()
	_, err := a.swagger.Endpoints.EndpointSnapshots(params, nil)
	if err != nil {
		return fmt.Errorf("failed to snapshot all endpoints: %w", err)
	}
	return nil
}
