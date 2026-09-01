package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*UpdateTopologyVizConfigsAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*UpdateTopologyVizConfigsAction)(nil)

// UpdateTopologyVizConfigsAction is the generated Terraform action implementation.
type UpdateTopologyVizConfigsAction struct {
	client *client.Client
}

// UpdateTopologyVizConfigsActionModel describes the action configuration shape.
type UpdateTopologyVizConfigsActionModel struct {
	AliasConfig               types.List   `tfsdk:"alias_config" json:"aliasConfig"`
	HierarchicalTags          types.List   `tfsdk:"hierarchical_tags" json:"hierarchicalTags"`
	LinkRepresentationConfigs types.List   `tfsdk:"link_representation_configs" json:"linkRepresentationConfigs"`
	NetworkDevicesEnabled     types.Bool   `tfsdk:"network_devices_enabled" json:"networkDevicesEnabled"`
	PlacementTags             types.List   `tfsdk:"placement_tags" json:"placementTags"`
	SubGroupKeys              types.List   `tfsdk:"sub_group_keys" json:"subGroupKeys"`
	Switch                    types.Bool   `tfsdk:"switch"`
	ToolsViewEnabled          types.Bool   `tfsdk:"tools_view_enabled" json:"toolsViewEnabled"`
	TopologyType              types.String `tfsdk:"topology_type" json:"topologyType"`
}

// NewUpdateTopologyVizConfigsAction returns a new instance of the generated action.
func NewUpdateTopologyVizConfigsAction() action.Action {
	return &UpdateTopologyVizConfigsAction{}
}

// Metadata returns the action type name.
func (r *UpdateTopologyVizConfigsAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_update_topology_viz_configs"
}

// Schema returns the action schema.
func (r *UpdateTopologyVizConfigsAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Configuring topology hierarchical and placement tags", Attributes: map[string]schema.Attribute{"alias_config": schema.ListNestedAttribute{MarkdownDescription: "Node group titles can be set using this property. This is applicable only for smart sankey.", Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"key": schema.StringAttribute{Optional: true}, "value": schema.StringAttribute{Optional: true}}}}, "hierarchical_tags": schema.ListAttribute{MarkdownDescription: "List of topology hierarchical tag keys(Ids). This will allow the user to drill down from a top level hierarchy view to a lower level detailed view", Required: true, ElementType: types.StringType}, "link_representation_configs": schema.ListNestedAttribute{Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"color_code": schema.StringAttribute{MarkdownDescription: "The color code of the link", Optional: true}, "link_format": schema.StringAttribute{MarkdownDescription: "The format of the link", Optional: true}, "link_type": schema.StringAttribute{MarkdownDescription: "Represents the type of the link", Optional: true}}}}, "network_devices_enabled": schema.BoolAttribute{MarkdownDescription: "set this to 'true' in-order to view network devices information in topology view, false by default", Optional: true}, "placement_tags": schema.ListNestedAttribute{Required: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"name": schema.StringAttribute{MarkdownDescription: "Unique alias for the placement config. This will help the user to  give a meaning full name for placement tag config", Required: true}, "tag_key": schema.StringAttribute{MarkdownDescription: "Unique placement tag key(Id)", Required: true}, "tag_values": schema.ListAttribute{MarkdownDescription: "List of placement tag values associated with the specified tag key", Required: true, ElementType: types.StringType}}}}, "sub_group_keys": schema.ListAttribute{Optional: true, ElementType: types.StringType}, "switch": schema.BoolAttribute{MarkdownDescription: "true/false, used for defining the update if for changing the topology type. true means topology type update otherwise some other properties update", Optional: true}, "tools_view_enabled": schema.BoolAttribute{MarkdownDescription: "set this to 'true' in-order to view tools information in topology view, false by default", Optional: true}, "topology_type": schema.StringAttribute{MarkdownDescription: "Topology type. Decided based on the number of managed nodes and tags configuration. 'TAG_BASED_SANKEY' if the topology tags are configured,'SMART_SANKEY' if the tags are not configured and number of managed nodes is lesser than or equal to 100, 'NOT_DEFINED' if the tags are not configured and number of managed nodes is greater than 100", Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *UpdateTopologyVizConfigsAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UpdateTopologyVizConfigsActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *UpdateTopologyVizConfigsAction) invokeRemote(ctx context.Context, config *UpdateTopologyVizConfigsActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/topoviz/config"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_topology_viz_configs", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_topology_viz_configs", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_topology_viz_configs", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.Switch.IsNull() {
		query.Set("switch", strconv.FormatBool(config.Switch.ValueBool()))
	}
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_topology_viz_configs", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 204) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_topology_viz_configs", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_topology_viz_configs", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_topology_viz_configs", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_topology_viz_configs", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_topology_viz_configs", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_topology_viz_configs", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_topology_viz_configs", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_update_topology_viz_configs", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_topology_viz_configs", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *UpdateTopologyVizConfigsAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
