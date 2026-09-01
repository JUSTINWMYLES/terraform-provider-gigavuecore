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
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*UpdateMapPriorityAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*UpdateMapPriorityAction)(nil)

// UpdateMapPriorityAction is the generated Terraform action implementation.
type UpdateMapPriorityAction struct {
	client *client.Client
}

// UpdateMapPriorityActionModel describes the action configuration shape.
type UpdateMapPriorityActionModel struct {
	ClusterId    types.String `tfsdk:"cluster_id"`
	Id           types.String `tfsdk:"id"`
	MapAlias     types.String `tfsdk:"map_alias" json:"mapAlias"`
	MapChainId   types.String `tfsdk:"map_chain_id" json:"mapChainId"`
	PriorityType types.String `tfsdk:"priority_type" json:"priorityType"`
	RefMapAlias  types.String `tfsdk:"ref_map_alias" json:"refMapAlias"`
	SrcPortsAsId types.String `tfsdk:"src_ports_as_id" json:"srcPortsAsId"`
}

// NewUpdateMapPriorityAction returns a new instance of the generated action.
func NewUpdateMapPriorityAction() action.Action {
	return &UpdateMapPriorityAction{}
}

// Metadata returns the action type name.
func (r *UpdateMapPriorityAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_update_map_priority"
}

// Schema returns the action schema.
func (r *UpdateMapPriorityAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Update Map priority within its Chain", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "id": schema.StringAttribute{MarkdownDescription: "srcPortsAsId or mapChainId for which the map order is to be redefined. Note that mapChainID works only for classic maps.", Required: true}, "map_alias": schema.StringAttribute{MarkdownDescription: "alias of the map(either cluster map or fabric map) to update priority for", Required: true}, "map_chain_id": schema.StringAttribute{MarkdownDescription: "mapChain ID - replaces srcPortsAsId - should be used anywhere that requires srcPortsAsId", Optional: true}, "priority_type": schema.StringAttribute{MarkdownDescription: "map priority update type within its map chain", Required: true}, "ref_map_alias": schema.StringAttribute{MarkdownDescription: "when 'priorityType' is 'before' or 'after', this field specifies the reference map (either cluster map or fabric map) for the update", Optional: true}, "src_ports_as_id": schema.StringAttribute{MarkdownDescription: "(Deprecated - use mapChainId instead) mapChain ID", Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *UpdateMapPriorityAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UpdateMapPriorityActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *UpdateMapPriorityAction) invokeRemote(ctx context.Context, config *UpdateMapPriorityActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/mapChains/{id}/mapPriority"
	reqPath = strings.ReplaceAll(reqPath, "{id}", url.PathEscape(config.Id.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_map_priority", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_map_priority", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_map_priority", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_map_priority", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_map_priority", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_map_priority", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_map_priority", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_map_priority", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_map_priority", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_map_priority", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_map_priority", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_update_map_priority", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_map_priority", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *UpdateMapPriorityAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
