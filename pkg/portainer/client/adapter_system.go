package client

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/portainer/client-api-go/v2/pkg/client/kubernetes"
	"github.com/portainer/client-api-go/v2/pkg/client/roles"
	apimodels "github.com/portainer/client-api-go/v2/pkg/models"
)

// ListRoles lists all roles.
func (a *portainerAPIAdapter) ListRoles() ([]*apimodels.PortainereeRole, error) {
	params := roles.NewRoleListParams()
	resp, err := a.swagger.Roles.RoleList(params, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list roles: %w", err)
	}
	return resp.Payload, nil
}

// GetMOTD retrieves the message of the day.
func (a *portainerAPIAdapter) GetMOTD() (map[string]any, error) {
	// Use raw HTTP to avoid SDK Hash type mismatch
	// (SDK defines Hash as []int64, but newer API versions return a string).
	op := &runtime.ClientOperation{
		ID:                 "MOTD",
		Method:             "GET",
		PathPattern:        "/motd",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{a.scheme},
		Params: runtime.ClientRequestWriterFunc(func(req runtime.ClientRequest, reg strfmt.Registry) error {
			return nil
		}),
		AuthInfo: a.httpTransport.DefaultAuthentication,
		Reader: runtime.ClientResponseReaderFunc(func(resp runtime.ClientResponse, consumer runtime.Consumer) (any, error) {
			var result map[string]any
			if err := consumer.Consume(resp.Body(), &result); err != nil {
				return nil, err
			}
			return result, nil
		}),
	}
	res, err := a.httpTransport.Submit(op)
	if err != nil {
		return nil, fmt.Errorf("failed to get MOTD: %w", err)
	}
	return res.(map[string]any), nil
}

// GetDockerDashboard retrieves the Docker dashboard data for a specific environment.
// Uses raw HTTP GET because the SDK sends POST but newer Portainer versions require GET.
func (a *portainerAPIAdapter) GetDockerDashboard(environmentId int64) (*apimodels.DockerDashboardResponse, error) {
	op := &runtime.ClientOperation{
		ID:                 "DockerDashboard",
		Method:             "GET",
		PathPattern:        fmt.Sprintf("/docker/%d/dashboard", environmentId),
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{a.scheme},
		Params: runtime.ClientRequestWriterFunc(func(req runtime.ClientRequest, reg strfmt.Registry) error {
			return nil
		}),
		AuthInfo: a.httpTransport.DefaultAuthentication,
		Reader: runtime.ClientResponseReaderFunc(func(resp runtime.ClientResponse, consumer runtime.Consumer) (any, error) {
			var result apimodels.DockerDashboardResponse
			if err := consumer.Consume(resp.Body(), &result); err != nil {
				return nil, err
			}
			return &result, nil
		}),
	}
	res, err := a.httpTransport.Submit(op)
	if err != nil {
		return nil, fmt.Errorf("failed to get docker dashboard: %w", err)
	}
	return res.(*apimodels.DockerDashboardResponse), nil
}

// GetKubernetesDashboard retrieves the Kubernetes dashboard data for a specific environment.
// Uses raw HTTP GET because the SDK expects an array but the API returns a single object.
func (a *portainerAPIAdapter) GetKubernetesDashboard(environmentId int64) (*apimodels.KubernetesK8sDashboard, error) {
	op := &runtime.ClientOperation{
		ID:                 "KubernetesDashboard",
		Method:             "GET",
		PathPattern:        fmt.Sprintf("/kubernetes/%d/dashboard", environmentId),
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{a.scheme},
		Params: runtime.ClientRequestWriterFunc(func(req runtime.ClientRequest, reg strfmt.Registry) error {
			return nil
		}),
		AuthInfo: a.httpTransport.DefaultAuthentication,
		Reader: runtime.ClientResponseReaderFunc(func(resp runtime.ClientResponse, consumer runtime.Consumer) (any, error) {
			var result apimodels.KubernetesK8sDashboard
			if err := consumer.Consume(resp.Body(), &result); err != nil {
				return nil, err
			}
			return &result, nil
		}),
	}
	res, err := a.httpTransport.Submit(op)
	if err != nil {
		return nil, fmt.Errorf("failed to get kubernetes dashboard: %w", err)
	}
	return res.(*apimodels.KubernetesK8sDashboard), nil
}

// GetKubernetesNamespaces retrieves the Kubernetes namespaces for a specific environment.
func (a *portainerAPIAdapter) GetKubernetesNamespaces(environmentId int64) ([]*apimodels.PortainerK8sNamespaceInfo, error) {
	params := kubernetes.NewGetKubernetesNamespacesParams().WithID(environmentId)
	resp, err := a.swagger.Kubernetes.GetKubernetesNamespaces(params, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get kubernetes namespaces: %w", err)
	}
	return resp.Payload, nil
}

// GetKubernetesConfig retrieves the Kubernetes config for a specific environment.
func (a *portainerAPIAdapter) GetKubernetesConfig(environmentId int64) (interface{}, error) {
	params := kubernetes.NewGetKubernetesConfigParams().WithIds([]int64{environmentId})
	resp, err := a.swagger.Kubernetes.GetKubernetesConfig(params, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get kubernetes config: %w", err)
	}
	return resp.Payload, nil
}
