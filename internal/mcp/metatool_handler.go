package mcp

import (
	"context"
	"fmt"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// RegisterMetaTools builds and registers all meta-tools on the MCP server.
// In read-only mode, write actions are excluded from the action enum and
// their handlers are not registered. If a meta-tool has no available
// actions after filtering (e.g. all are write-only and read-only is on),
// it is silently skipped.
func (s *PortainerMCPServer) RegisterMetaTools() {
	defs := metaToolDefinitions()
	for _, def := range defs {
		s.registerOneMetaTool(def)
	}
}

// registerOneMetaTool builds a single meta-tool from its definition,
// filtering actions by read-only mode, and registers it. It merges parameter
// schemas from the granular tools (if available) into the meta-tool's schema
// so that LLM callers can discover all available parameters.
func (s *PortainerMCPServer) registerOneMetaTool(def metaToolDef) {
	// Filter actions based on read-only mode
	available := make([]metaAction, 0, len(def.actions))
	for _, a := range def.actions {
		if s.readOnly && !a.readOnly {
			continue
		}
		available = append(available, a)
	}

	if len(available) == 0 {
		return
	}

	// Build action enum values and handler dispatch map
	actionNames := make([]string, len(available))
	handlers := make(map[string]server.ToolHandlerFunc, len(available))
	for i, a := range available {
		actionNames[i] = a.name
		handlers[a.name] = a.handler(s)
	}

	// Compute annotation: if ALL remaining actions are read-only, mark the
	// meta-tool as read-only. Otherwise use the definition's annotation.
	annotation := def.annotation
	allReadOnly := true
	for _, a := range available {
		if !a.readOnly {
			allReadOnly = false
			break
		}
	}
	if allReadOnly {
		annotation.ReadOnlyHint = boolPtr(true)
		annotation.DestructiveHint = boolPtr(false)
	}

	// Build the MCP tool programmatically
	tool := mcp.NewTool(def.name,
		mcp.WithDescription(def.description),
		mcp.WithToolAnnotation(annotation),
		mcp.WithString("action",
			mcp.Required(),
			mcp.Description(fmt.Sprintf("The operation to perform. Available actions: %s", strings.Join(actionNames, ", "))),
			mcp.Enum(actionNames...),
		),
	)

	// Merge parameters from granular tools into the meta-tool schema.
	// This allows LLM callers to discover all available parameters for all actions
	// in a single meta-tool, making the schema more complete.
	if s.tools != nil && len(s.tools) > 0 {
		mergeParametersFromGranularTools(&tool, available, s.tools)
	}

	// Register the meta-tool with a routing handler
	s.srv.AddTool(tool, makeMetaHandler(def.name, handlers))
}

// mergeParametersFromGranularTools extracts parameter schemas from granular
// tools and merges them into the meta-tool's InputSchema. Parameters are
// merged from all available actions, so LLM callers see all possible
// parameters they might need to pass based on the chosen action.
//
// If the same parameter name appears in multiple actions, it is only added
// once (first occurrence wins). The tools map is expected to have keys
// matching the action names in camelCase.
func mergeParametersFromGranularTools(tool *mcp.Tool, available []metaAction, tools map[string]mcp.Tool) {
	if tool.InputSchema.Properties == nil {
		tool.InputSchema.Properties = make(map[string]any)
	}

	// Iterate through all available actions and merge their parameters
	for _, action := range available {
		granularTool, ok := tools[action.name]
		if !ok {
			// Skip if the tool is not found in the tools map. This can happen
			// in test setups where tools are not loaded.
			continue
		}

		// Extract properties from the granular tool's input schema
		if granularTool.InputSchema.Properties == nil {
			continue
		}

		// Merge each property from the granular tool, skipping duplicates
		for propName, propSchema := range granularTool.InputSchema.Properties {
			if _, exists := tool.InputSchema.Properties[propName]; !exists {
				tool.InputSchema.Properties[propName] = propSchema
			}
		}
	}
}

// makeMetaHandler creates a ToolHandlerFunc that routes to the correct
// sub-handler based on the "action" parameter.
func makeMetaHandler(metaToolName string, handlers map[string]server.ToolHandlerFunc) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		actionRaw, ok := request.GetArguments()["action"]
		if !ok {
			return mcp.NewToolResultError("missing required parameter: action"), nil
		}

		action, ok := actionRaw.(string)
		if !ok || action == "" {
			return mcp.NewToolResultError("parameter 'action' must be a non-empty string"), nil
		}

		handler, ok := handlers[action]
		if !ok {
			available := make([]string, 0, len(handlers))
			for k := range handlers {
				available = append(available, k)
			}
			return mcp.NewToolResultError(fmt.Sprintf(
				"unknown action '%s' for tool '%s'. Available actions: %s",
				action, metaToolName, strings.Join(available, ", "),
			)), nil
		}

		return handler(ctx, request)
	}
}
