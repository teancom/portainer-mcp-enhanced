package mcp

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// metaAction maps an action name to its handler and access metadata.
type metaAction struct {
	name     string
	handler  func(s *PortainerMCPServer) server.ToolHandlerFunc
	readOnly bool // true = always available; false = hidden in read-only mode
}

// metaToolDef describes a single grouped meta-tool.
type metaToolDef struct {
	name        string
	description string
	actions     []metaAction
	annotation  mcp.ToolAnnotation
}

// boolPtr is a convenience helper for creating *bool values.
func boolPtr(v bool) *bool { return &v }

// metaToolDefinitions returns the complete list of meta-tool groups.
// Each group aggregates several existing granular tools behind a single
// "action" enum parameter. Read-only mode filters write actions at
// registration time, so the enum only exposes permitted actions.
func metaToolDefinitions() []metaToolDef {
	return []metaToolDef{
		{
			name:        "manage_environments",
			description: "Manage Portainer environments, environment groups, and tags. Actions: listEnvironments, getEnvironment, deleteEnvironment, snapshotEnvironment, snapshotAllEnvironments, updateEnvironmentTags, updateEnvironmentUserAccesses, updateEnvironmentTeamAccesses, listEnvironmentGroups, createEnvironmentGroup, updateEnvironmentGroupName, updateEnvironmentGroupEnvironments, updateEnvironmentGroupTags, listEnvironmentTags, createEnvironmentTag, deleteEnvironmentTag. Set 'action' parameter to choose.",
			actions: []metaAction{
				{name: "listEnvironments", handler: (*PortainerMCPServer).HandleGetEnvironments, readOnly: true},
				{name: "getEnvironment", handler: (*PortainerMCPServer).HandleGetEnvironment, readOnly: true},
				{name: "deleteEnvironment", handler: (*PortainerMCPServer).HandleDeleteEnvironment, readOnly: false},
				{name: "snapshotEnvironment", handler: (*PortainerMCPServer).HandleSnapshotEnvironment, readOnly: false},
				{name: "snapshotAllEnvironments", handler: (*PortainerMCPServer).HandleSnapshotAllEnvironments, readOnly: false},
				{name: "updateEnvironmentTags", handler: (*PortainerMCPServer).HandleUpdateEnvironmentTags, readOnly: false},
				{name: "updateEnvironmentUserAccesses", handler: (*PortainerMCPServer).HandleUpdateEnvironmentUserAccesses, readOnly: false},
				{name: "updateEnvironmentTeamAccesses", handler: (*PortainerMCPServer).HandleUpdateEnvironmentTeamAccesses, readOnly: false},
				{name: "listEnvironmentGroups", handler: (*PortainerMCPServer).HandleGetEnvironmentGroups, readOnly: true},
				{name: "createEnvironmentGroup", handler: (*PortainerMCPServer).HandleCreateEnvironmentGroup, readOnly: false},
				{name: "updateEnvironmentGroupName", handler: (*PortainerMCPServer).HandleUpdateEnvironmentGroupName, readOnly: false},
				{name: "updateEnvironmentGroupEnvironments", handler: (*PortainerMCPServer).HandleUpdateEnvironmentGroupEnvironments, readOnly: false},
				{name: "updateEnvironmentGroupTags", handler: (*PortainerMCPServer).HandleUpdateEnvironmentGroupTags, readOnly: false},
				{name: "listEnvironmentTags", handler: (*PortainerMCPServer).HandleGetEnvironmentTags, readOnly: true},
				{name: "createEnvironmentTag", handler: (*PortainerMCPServer).HandleCreateEnvironmentTag, readOnly: false},
				{name: "deleteEnvironmentTag", handler: (*PortainerMCPServer).HandleDeleteEnvironmentTag, readOnly: false},
			},
			annotation: mcp.ToolAnnotation{
				Title:           "Manage Environments",
				ReadOnlyHint:    boolPtr(false),
				DestructiveHint: boolPtr(true),
				IdempotentHint:  boolPtr(false),
				OpenWorldHint:   boolPtr(false),
			},
		},
		{
			name:        "manage_stacks",
			description: "Manage Docker stacks (Compose and Edge deployments). Actions: listStacks, listRegularStacks, getStack, getStackFile, inspectStackFile, createStack, updateStack, deleteStack, updateStackGit, redeployStackGit, startStack, stopStack, migrateStack. Set 'action' parameter to choose.",
			actions: []metaAction{
				{name: "listStacks", handler: (*PortainerMCPServer).HandleGetStacks, readOnly: true},
				{name: "listRegularStacks", handler: (*PortainerMCPServer).HandleListRegularStacks, readOnly: true},
				{name: "getStack", handler: (*PortainerMCPServer).HandleInspectStack, readOnly: true},
				{name: "getStackFile", handler: (*PortainerMCPServer).HandleGetStackFile, readOnly: true},
				{name: "inspectStackFile", handler: (*PortainerMCPServer).HandleInspectStackFile, readOnly: true},
				{name: "createStack", handler: (*PortainerMCPServer).HandleCreateStack, readOnly: false},
				{name: "updateStack", handler: (*PortainerMCPServer).HandleUpdateStack, readOnly: false},
				{name: "deleteStack", handler: (*PortainerMCPServer).HandleDeleteStack, readOnly: false},
				{name: "updateStackGit", handler: (*PortainerMCPServer).HandleUpdateStackGit, readOnly: false},
				{name: "redeployStackGit", handler: (*PortainerMCPServer).HandleRedeployStackGit, readOnly: false},
				{name: "startStack", handler: (*PortainerMCPServer).HandleStartStack, readOnly: false},
				{name: "stopStack", handler: (*PortainerMCPServer).HandleStopStack, readOnly: false},
				{name: "migrateStack", handler: (*PortainerMCPServer).HandleMigrateStack, readOnly: false},
			},
			annotation: mcp.ToolAnnotation{
				Title:           "Manage Stacks",
				ReadOnlyHint:    boolPtr(false),
				DestructiveHint: boolPtr(true),
				IdempotentHint:  boolPtr(false),
				OpenWorldHint:   boolPtr(false),
			},
		},
		{
			name:        "manage_access_groups",
			description: "Manage access groups for environment-level permissions. Actions: listAccessGroups, createAccessGroup, updateAccessGroupName, updateAccessGroupUserAccesses, updateAccessGroupTeamAccesses, addEnvironmentToAccessGroup, removeEnvironmentFromAccessGroup. Set 'action' parameter to choose.",
			actions: []metaAction{
				{name: "listAccessGroups", handler: (*PortainerMCPServer).HandleGetAccessGroups, readOnly: true},
				{name: "createAccessGroup", handler: (*PortainerMCPServer).HandleCreateAccessGroup, readOnly: false},
				{name: "updateAccessGroupName", handler: (*PortainerMCPServer).HandleUpdateAccessGroupName, readOnly: false},
				{name: "updateAccessGroupUserAccesses", handler: (*PortainerMCPServer).HandleUpdateAccessGroupUserAccesses, readOnly: false},
				{name: "updateAccessGroupTeamAccesses", handler: (*PortainerMCPServer).HandleUpdateAccessGroupTeamAccesses, readOnly: false},
				{name: "addEnvironmentToAccessGroup", handler: (*PortainerMCPServer).HandleAddEnvironmentToAccessGroup, readOnly: false},
				{name: "removeEnvironmentFromAccessGroup", handler: (*PortainerMCPServer).HandleRemoveEnvironmentFromAccessGroup, readOnly: false},
			},
			annotation: mcp.ToolAnnotation{
				Title:           "Manage Access Groups",
				ReadOnlyHint:    boolPtr(false),
				DestructiveHint: boolPtr(true),
				IdempotentHint:  boolPtr(false),
				OpenWorldHint:   boolPtr(false),
			},
		},
		{
			name:        "manage_users",
			description: "Manage Portainer user accounts and roles. Actions: listUsers, getUser, createUser, deleteUser, updateUserRole. Set 'action' parameter to choose.",
			actions: []metaAction{
				{name: "listUsers", handler: (*PortainerMCPServer).HandleGetUsers, readOnly: true},
				{name: "getUser", handler: (*PortainerMCPServer).HandleGetUser, readOnly: true},
				{name: "createUser", handler: (*PortainerMCPServer).HandleCreateUser, readOnly: false},
				{name: "deleteUser", handler: (*PortainerMCPServer).HandleDeleteUser, readOnly: false},
				{name: "updateUserRole", handler: (*PortainerMCPServer).HandleUpdateUserRole, readOnly: false},
			},
			annotation: mcp.ToolAnnotation{
				Title:           "Manage Users",
				ReadOnlyHint:    boolPtr(false),
				DestructiveHint: boolPtr(true),
				IdempotentHint:  boolPtr(false),
				OpenWorldHint:   boolPtr(false),
			},
		},
		{
			name:        "manage_teams",
			description: "Manage Portainer teams and membership. Actions: listTeams, getTeam, createTeam, deleteTeam, updateTeamName, updateTeamMembers. Set 'action' parameter to choose.",
			actions: []metaAction{
				{name: "listTeams", handler: (*PortainerMCPServer).HandleGetTeams, readOnly: true},
				{name: "getTeam", handler: (*PortainerMCPServer).HandleGetTeam, readOnly: true},
				{name: "createTeam", handler: (*PortainerMCPServer).HandleCreateTeam, readOnly: false},
				{name: "deleteTeam", handler: (*PortainerMCPServer).HandleDeleteTeam, readOnly: false},
				{name: "updateTeamName", handler: (*PortainerMCPServer).HandleUpdateTeamName, readOnly: false},
				{name: "updateTeamMembers", handler: (*PortainerMCPServer).HandleUpdateTeamMembers, readOnly: false},
			},
			annotation: mcp.ToolAnnotation{
				Title:           "Manage Teams",
				ReadOnlyHint:    boolPtr(false),
				DestructiveHint: boolPtr(true),
				IdempotentHint:  boolPtr(false),
				OpenWorldHint:   boolPtr(false),
			},
		},
		{
			name:        "manage_docker",
			description: "Interact with Docker environments via dashboards and proxy API calls. Actions: getDockerDashboard, dockerProxy. Set 'action' parameter to choose.",
			actions: []metaAction{
				{name: "getDockerDashboard", handler: (*PortainerMCPServer).HandleGetDockerDashboard, readOnly: true},
				{name: "dockerProxy", handler: (*PortainerMCPServer).HandleDockerProxy, readOnly: false},
			},
			annotation: mcp.ToolAnnotation{
				Title:           "Manage Docker",
				ReadOnlyHint:    boolPtr(false),
				DestructiveHint: boolPtr(true),
				IdempotentHint:  boolPtr(false),
				OpenWorldHint:   boolPtr(true),
			},
		},
		{
			name:        "manage_kubernetes",
			description: "Interact with Kubernetes environments via dashboards, namespaces, kubeconfig, and proxy API calls. Actions: getKubernetesResourceStripped, getKubernetesDashboard, listKubernetesNamespaces, getKubernetesConfig, kubernetesProxy. Set 'action' parameter to choose.",
			actions: []metaAction{
				{name: "getKubernetesResourceStripped", handler: (*PortainerMCPServer).HandleKubernetesProxyStripped, readOnly: true},
				{name: "getKubernetesDashboard", handler: (*PortainerMCPServer).HandleGetKubernetesDashboard, readOnly: true},
				{name: "listKubernetesNamespaces", handler: (*PortainerMCPServer).HandleListKubernetesNamespaces, readOnly: true},
				{name: "getKubernetesConfig", handler: (*PortainerMCPServer).HandleGetKubernetesConfig, readOnly: true},
				{name: "kubernetesProxy", handler: (*PortainerMCPServer).HandleKubernetesProxy, readOnly: false},
			},
			annotation: mcp.ToolAnnotation{
				Title:           "Manage Kubernetes",
				ReadOnlyHint:    boolPtr(false),
				DestructiveHint: boolPtr(true),
				IdempotentHint:  boolPtr(false),
				OpenWorldHint:   boolPtr(true),
			},
		},
		{
			name:        "manage_helm",
			description: "Manage Helm repositories, charts, and releases on Kubernetes environments. Actions: listHelmRepositories, searchHelmCharts, listHelmReleases, getHelmReleaseHistory, addHelmRepository, removeHelmRepository, installHelmChart, deleteHelmRelease. Set 'action' parameter to choose.",
			actions: []metaAction{
				{name: "listHelmRepositories", handler: (*PortainerMCPServer).HandleListHelmRepositories, readOnly: true},
				{name: "searchHelmCharts", handler: (*PortainerMCPServer).HandleSearchHelmCharts, readOnly: true},
				{name: "listHelmReleases", handler: (*PortainerMCPServer).HandleListHelmReleases, readOnly: true},
				{name: "getHelmReleaseHistory", handler: (*PortainerMCPServer).HandleGetHelmReleaseHistory, readOnly: true},
				{name: "addHelmRepository", handler: (*PortainerMCPServer).HandleAddHelmRepository, readOnly: false},
				{name: "removeHelmRepository", handler: (*PortainerMCPServer).HandleRemoveHelmRepository, readOnly: false},
				{name: "installHelmChart", handler: (*PortainerMCPServer).HandleInstallHelmChart, readOnly: false},
				{name: "deleteHelmRelease", handler: (*PortainerMCPServer).HandleDeleteHelmRelease, readOnly: false},
			},
			annotation: mcp.ToolAnnotation{
				Title:           "Manage Helm",
				ReadOnlyHint:    boolPtr(false),
				DestructiveHint: boolPtr(true),
				IdempotentHint:  boolPtr(false),
				OpenWorldHint:   boolPtr(false),
			},
		},
		{
			name:        "manage_registries",
			description: "Manage container registries (Quay, Azure, DockerHub, GitLab, ECR, custom). Actions: listRegistries, getRegistry, createRegistry, updateRegistry, deleteRegistry. Set 'action' parameter to choose.",
			actions: []metaAction{
				{name: "listRegistries", handler: (*PortainerMCPServer).HandleListRegistries, readOnly: true},
				{name: "getRegistry", handler: (*PortainerMCPServer).HandleGetRegistry, readOnly: true},
				{name: "createRegistry", handler: (*PortainerMCPServer).HandleCreateRegistry, readOnly: false},
				{name: "updateRegistry", handler: (*PortainerMCPServer).HandleUpdateRegistry, readOnly: false},
				{name: "deleteRegistry", handler: (*PortainerMCPServer).HandleDeleteRegistry, readOnly: false},
			},
			annotation: mcp.ToolAnnotation{
				Title:           "Manage Registries",
				ReadOnlyHint:    boolPtr(false),
				DestructiveHint: boolPtr(true),
				IdempotentHint:  boolPtr(false),
				OpenWorldHint:   boolPtr(false),
			},
		},
		{
			name:        "manage_templates",
			description: "Manage custom and application templates for stack deployment. Actions: listCustomTemplates, getCustomTemplate, getCustomTemplateFile, createCustomTemplate, deleteCustomTemplate, listAppTemplates, getAppTemplateFile. Set 'action' parameter to choose.",
			actions: []metaAction{
				{name: "listCustomTemplates", handler: (*PortainerMCPServer).HandleListCustomTemplates, readOnly: true},
				{name: "getCustomTemplate", handler: (*PortainerMCPServer).HandleGetCustomTemplate, readOnly: true},
				{name: "getCustomTemplateFile", handler: (*PortainerMCPServer).HandleGetCustomTemplateFile, readOnly: true},
				{name: "createCustomTemplate", handler: (*PortainerMCPServer).HandleCreateCustomTemplate, readOnly: false},
				{name: "deleteCustomTemplate", handler: (*PortainerMCPServer).HandleDeleteCustomTemplate, readOnly: false},
				{name: "listAppTemplates", handler: (*PortainerMCPServer).HandleListAppTemplates, readOnly: true},
				{name: "getAppTemplateFile", handler: (*PortainerMCPServer).HandleGetAppTemplateFile, readOnly: true},
			},
			annotation: mcp.ToolAnnotation{
				Title:           "Manage Templates",
				ReadOnlyHint:    boolPtr(false),
				DestructiveHint: boolPtr(true),
				IdempotentHint:  boolPtr(false),
				OpenWorldHint:   boolPtr(false),
			},
		},
		{
			name:        "manage_backups",
			description: "Manage Portainer server backups and restore (local and S3). Actions: getBackupStatus, getBackupS3Settings, createBackup, backupToS3, restoreFromS3. Set 'action' parameter to choose.",
			actions: []metaAction{
				{name: "getBackupStatus", handler: (*PortainerMCPServer).HandleGetBackupStatus, readOnly: true},
				{name: "getBackupS3Settings", handler: (*PortainerMCPServer).HandleGetBackupS3Settings, readOnly: true},
				{name: "createBackup", handler: (*PortainerMCPServer).HandleCreateBackup, readOnly: false},
				{name: "backupToS3", handler: (*PortainerMCPServer).HandleBackupToS3, readOnly: false},
				{name: "restoreFromS3", handler: (*PortainerMCPServer).HandleRestoreFromS3, readOnly: false},
			},
			annotation: mcp.ToolAnnotation{
				Title:           "Manage Backups",
				ReadOnlyHint:    boolPtr(false),
				DestructiveHint: boolPtr(true),
				IdempotentHint:  boolPtr(false),
				OpenWorldHint:   boolPtr(false),
			},
		},
		{
			name:        "manage_webhooks",
			description: "Manage webhooks for container services and automated deployments. Actions: listWebhooks, createWebhook, deleteWebhook. Set 'action' parameter to choose.",
			actions: []metaAction{
				{name: "listWebhooks", handler: (*PortainerMCPServer).HandleListWebhooks, readOnly: true},
				{name: "createWebhook", handler: (*PortainerMCPServer).HandleCreateWebhook, readOnly: false},
				{name: "deleteWebhook", handler: (*PortainerMCPServer).HandleDeleteWebhook, readOnly: false},
			},
			annotation: mcp.ToolAnnotation{
				Title:           "Manage Webhooks",
				ReadOnlyHint:    boolPtr(false),
				DestructiveHint: boolPtr(true),
				IdempotentHint:  boolPtr(false),
				OpenWorldHint:   boolPtr(false),
			},
		},
		{
			name:        "manage_edge",
			description: "Manage Edge compute jobs and update schedules for remote environments. Actions: listEdgeJobs, getEdgeJob, getEdgeJobFile, createEdgeJob, deleteEdgeJob, listEdgeUpdateSchedules. Set 'action' parameter to choose.",
			actions: []metaAction{
				{name: "listEdgeJobs", handler: (*PortainerMCPServer).HandleListEdgeJobs, readOnly: true},
				{name: "getEdgeJob", handler: (*PortainerMCPServer).HandleGetEdgeJob, readOnly: true},
				{name: "getEdgeJobFile", handler: (*PortainerMCPServer).HandleGetEdgeJobFile, readOnly: true},
				{name: "createEdgeJob", handler: (*PortainerMCPServer).HandleCreateEdgeJob, readOnly: false},
				{name: "deleteEdgeJob", handler: (*PortainerMCPServer).HandleDeleteEdgeJob, readOnly: false},
				{name: "listEdgeUpdateSchedules", handler: (*PortainerMCPServer).HandleListEdgeUpdateSchedules, readOnly: true},
			},
			annotation: mcp.ToolAnnotation{
				Title:           "Manage Edge",
				ReadOnlyHint:    boolPtr(false),
				DestructiveHint: boolPtr(true),
				IdempotentHint:  boolPtr(false),
				OpenWorldHint:   boolPtr(false),
			},
		},
		{
			name:        "manage_settings",
			description: "Manage Portainer server settings, public settings, and SSL configuration. Actions: getSettings, getPublicSettings, updateSettings, getSSLSettings, updateSSLSettings. Set 'action' parameter to choose.",
			actions: []metaAction{
				{name: "getSettings", handler: (*PortainerMCPServer).HandleGetSettings, readOnly: true},
				{name: "getPublicSettings", handler: (*PortainerMCPServer).HandleGetPublicSettings, readOnly: true},
				{name: "updateSettings", handler: (*PortainerMCPServer).HandleUpdateSettings, readOnly: false},
				{name: "getSSLSettings", handler: (*PortainerMCPServer).HandleGetSSLSettings, readOnly: true},
				{name: "updateSSLSettings", handler: (*PortainerMCPServer).HandleUpdateSSLSettings, readOnly: false},
			},
			annotation: mcp.ToolAnnotation{
				Title:           "Manage Settings",
				ReadOnlyHint:    boolPtr(false),
				DestructiveHint: boolPtr(false),
				IdempotentHint:  boolPtr(true),
				OpenWorldHint:   boolPtr(false),
			},
		},
		{
			name:        "manage_system",
			description: "Portainer system info, roles, MOTD, and authentication. Actions: getSystemStatus, listRoles, getMOTD, authenticate, logout. Set 'action' parameter to choose.",
			actions: []metaAction{
				{name: "getSystemStatus", handler: (*PortainerMCPServer).HandleGetSystemStatus, readOnly: true},
				{name: "listRoles", handler: (*PortainerMCPServer).HandleListRoles, readOnly: true},
				{name: "getMOTD", handler: (*PortainerMCPServer).HandleGetMOTD, readOnly: true},
				{name: "authenticate", handler: (*PortainerMCPServer).HandleAuthenticateUser, readOnly: true},
				{name: "logout", handler: (*PortainerMCPServer).HandleLogout, readOnly: false},
			},
			annotation: mcp.ToolAnnotation{
				Title:           "Manage System",
				ReadOnlyHint:    boolPtr(false),
				DestructiveHint: boolPtr(false),
				IdempotentHint:  boolPtr(false),
				OpenWorldHint:   boolPtr(false),
			},
		},
	}
}
