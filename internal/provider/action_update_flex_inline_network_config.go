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
var _ action.Action = (*UpdateFlexInlineNetworkConfigAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*UpdateFlexInlineNetworkConfigAction)(nil)

// UpdateFlexInlineNetworkConfigAction is the generated Terraform action implementation.
type UpdateFlexInlineNetworkConfigAction struct {
	client *client.Client
}

// UpdateFlexInlineNetworkConfigActionModel describes the action configuration shape.
type UpdateFlexInlineNetworkConfigActionModel struct {
	Alias               types.String  `tfsdk:"alias"`
	ClusterId           types.String  `tfsdk:"cluster_id" json:"clusterId"`
	ConfigData          types.Dynamic `tfsdk:"config_data" json:"configData"`
	ConfigStatus        types.String  `tfsdk:"config_status" json:"configStatus"`
	ConfigStatusReasons types.List    `tfsdk:"config_status_reasons" json:"configStatusReasons"`
	ConfigType          types.String  `tfsdk:"config_type" json:"configType"`
	HealthState         types.String  `tfsdk:"health_state" json:"healthState"`
	HealthStateReasons  types.Dynamic `tfsdk:"health_state_reasons" json:"healthStateReasons"`
	SolutionAlias       types.String  `tfsdk:"solution_alias" json:"solutionAlias"`
}

// NewUpdateFlexInlineNetworkConfigAction returns a new instance of the generated action.
func NewUpdateFlexInlineNetworkConfigAction() action.Action {
	return &UpdateFlexInlineNetworkConfigAction{}
}

// Metadata returns the action type name.
func (r *UpdateFlexInlineNetworkConfigAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_update_flex_inline_network_config"
}

// Schema returns the action schema.
func (r *UpdateFlexInlineNetworkConfigAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Updates inline network details used in the given solution", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Solution alias for which inline network is to be updated", Required: true}, "cluster_id": schema.StringAttribute{Optional: true}, "config_data": schema.DynamicAttribute{MarkdownDescription: "Data holding configuration details", Optional: true}, "config_status": schema.StringAttribute{MarkdownDescription: "Status of the created config object", Optional: true}, "config_status_reasons": schema.ListAttribute{Optional: true, ElementType: types.StringType}, "config_type": schema.StringAttribute{MarkdownDescription: "Type of the configuration object", Optional: true}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Optional: true}, "health_state_reasons": schema.DynamicAttribute{Optional: true}, "solution_alias": schema.StringAttribute{MarkdownDescription: "Alias of the solution", Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *UpdateFlexInlineNetworkConfigAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UpdateFlexInlineNetworkConfigActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *UpdateFlexInlineNetworkConfigAction) invokeRemote(ctx context.Context, config *UpdateFlexInlineNetworkConfigActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/flexInline/{alias}/inline/networks"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_flex_inline_network_config", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_flex_inline_network_config", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPatch, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_flex_inline_network_config", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_flex_inline_network_config", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_flex_inline_network_config", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_flex_inline_network_config", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_flex_inline_network_config", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_flex_inline_network_config", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_flex_inline_network_config", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_flex_inline_network_config", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_flex_inline_network_config", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_update_flex_inline_network_config", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_flex_inline_network_config", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *UpdateFlexInlineNetworkConfigAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
