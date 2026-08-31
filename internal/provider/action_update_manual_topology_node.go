package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*UpdateManualTopologyNodeAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*UpdateManualTopologyNodeAction)(nil)

// UpdateManualTopologyNodeAction is the generated Terraform action implementation.
type UpdateManualTopologyNodeAction struct {
	client *client.Client
}

// UpdateManualTopologyNodeActionModel describes the action configuration shape.
type UpdateManualTopologyNodeActionModel struct {
	BodyTopoNodeId types.String `tfsdk:"body_topo_node_id" json:"topoNodeId"`
	Comment        types.String `tfsdk:"comment"`
	Model          types.String `tfsdk:"model"`
	NodeAlias      types.String `tfsdk:"node_alias" json:"nodeAlias"`
	TopoNodeId     types.String `tfsdk:"topo_node_id"`
	Type           types.String `tfsdk:"type"`
	Vendor         types.String `tfsdk:"vendor"`
}

// NewUpdateManualTopologyNodeAction returns a new instance of the generated action.
func NewUpdateManualTopologyNodeAction() action.Action {
	return &UpdateManualTopologyNodeAction{}
}

// Metadata returns the action type name.
func (r *UpdateManualTopologyNodeAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_update_manual_topology_node"
}

// Schema returns the action schema.
func (r *UpdateManualTopologyNodeAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Update Manual Topology Node", Attributes: map[string]schema.Attribute{"body_topo_node_id": schema.StringAttribute{MarkdownDescription: "Node's unique identifier. Auto-assigned. This is required while updating existing node", Optional: true}, "comment": schema.StringAttribute{Optional: true}, "model": schema.StringAttribute{MarkdownDescription: "Node model, if available", Optional: true}, "node_alias": schema.StringAttribute{MarkdownDescription: "Node alias. User-assigned", Required: true}, "topo_node_id": schema.StringAttribute{MarkdownDescription: "Topology Node Id", Required: true}, "type": schema.StringAttribute{MarkdownDescription: "Annotation: type of a node. E.g. switch, router, etc...", Optional: true}, "vendor": schema.StringAttribute{MarkdownDescription: "Node vendor, if available", Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *UpdateManualTopologyNodeAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UpdateManualTopologyNodeActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *UpdateManualTopologyNodeAction) invokeRemote(ctx context.Context, config *UpdateManualTopologyNodeActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/topology/manual/nodes/{topoNodeId}"
	reqPath = strings.ReplaceAll(reqPath, "{topoNodeId}", url.PathEscape(config.TopoNodeId.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_manual_topology_node", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_manual_topology_node", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_manual_topology_node", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_manual_topology_node", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_manual_topology_node", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_manual_topology_node", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_manual_topology_node", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_manual_topology_node", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_manual_topology_node", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_manual_topology_node", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_manual_topology_node", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_update_manual_topology_node", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_manual_topology_node", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *UpdateManualTopologyNodeAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Action Configure Type", fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData))
		return
	}
	r.client = c
}
