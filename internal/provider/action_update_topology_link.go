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
var _ action.Action = (*UpdateTopologyLinkAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*UpdateTopologyLinkAction)(nil)

// UpdateTopologyLinkAction is the generated Terraform action implementation.
type UpdateTopologyLinkAction struct {
	client *client.Client
}

// UpdateTopologyLinkActionModel describes the action configuration shape.
type UpdateTopologyLinkActionModel struct {
	BodyTopoLinkId types.String `tfsdk:"body_topo_link_id" json:"topoLinkId"`
	Comment        types.String `tfsdk:"comment"`
	Connections    types.List   `tfsdk:"connections"`
	Endpoint1      types.Object `tfsdk:"endpoint1"`
	Endpoint2      types.Object `tfsdk:"endpoint2"`
	TopoLinkId     types.String `tfsdk:"topo_link_id"`
}

// NewUpdateTopologyLinkAction returns a new instance of the generated action.
func NewUpdateTopologyLinkAction() action.Action {
	return &UpdateTopologyLinkAction{}
}

// Metadata returns the action type name.
func (r *UpdateTopologyLinkAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_update_topology_link"
}

// Schema returns the action schema.
func (r *UpdateTopologyLinkAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Update Manual Topology Link", Attributes: map[string]schema.Attribute{"body_topo_link_id": schema.StringAttribute{MarkdownDescription: "Link's unique identifier. Auto-assigned. This is required while updating existing links", Optional: true}, "comment": schema.StringAttribute{Optional: true}, "connections": schema.ListNestedAttribute{MarkdownDescription: "Actual port ids which are part of the physical connection", Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"port1": schema.StringAttribute{Optional: true}, "port2": schema.StringAttribute{Optional: true}}}}, "endpoint1": schema.SingleNestedAttribute{MarkdownDescription: "Node link endpoint Create/Update Spec", Required: true, Attributes: map[string]schema.Attribute{"comment": schema.StringAttribute{Optional: true}, "component_alias": schema.StringAttribute{MarkdownDescription: "For 'gigamon': this is either portId (e.g. '1/1/x1') or Gigastream alias (e.g. 'gs1').  For 'manual': this is user provided or may be omitted", Optional: true}, "component_type": schema.StringAttribute{MarkdownDescription: "Type of the endpoint component. This is either 'Port' or 'GigaStream'", Required: true}, "port": schema.StringAttribute{MarkdownDescription: "For 'gigamon': this is the portId (e.g. '1/1/x1'). For 'manual': this is user provided or may be omitted", Optional: true}, "topo_node_id": schema.StringAttribute{MarkdownDescription: "Topology graph node Id. References one of the Topology Nodes in the graph", Required: true}, "topo_node_type": schema.StringAttribute{MarkdownDescription: "topology node type", Optional: true}}}, "endpoint2": schema.SingleNestedAttribute{MarkdownDescription: "Node link endpoint Create/Update Spec", Required: true, Attributes: map[string]schema.Attribute{"comment": schema.StringAttribute{Optional: true}, "component_alias": schema.StringAttribute{MarkdownDescription: "For 'gigamon': this is either portId (e.g. '1/1/x1') or Gigastream alias (e.g. 'gs1').  For 'manual': this is user provided or may be omitted", Optional: true}, "component_type": schema.StringAttribute{MarkdownDescription: "Type of the endpoint component. This is either 'Port' or 'GigaStream'", Required: true}, "port": schema.StringAttribute{MarkdownDescription: "For 'gigamon': this is the portId (e.g. '1/1/x1'). For 'manual': this is user provided or may be omitted", Optional: true}, "topo_node_id": schema.StringAttribute{MarkdownDescription: "Topology graph node Id. References one of the Topology Nodes in the graph", Required: true}, "topo_node_type": schema.StringAttribute{MarkdownDescription: "topology node type", Optional: true}}}, "topo_link_id": schema.StringAttribute{MarkdownDescription: "Topology Link Id", Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *UpdateTopologyLinkAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UpdateTopologyLinkActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *UpdateTopologyLinkAction) invokeRemote(ctx context.Context, config *UpdateTopologyLinkActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/topology/manual/links/{topoLinkId}"
	reqPath = strings.ReplaceAll(reqPath, "{topoLinkId}", url.PathEscape(config.TopoLinkId.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_topology_link", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_topology_link", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPatch, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_topology_link", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_topology_link", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_topology_link", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_topology_link", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_topology_link", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_topology_link", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_topology_link", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_topology_link", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_topology_link", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_update_topology_link", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_topology_link", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *UpdateTopologyLinkAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
