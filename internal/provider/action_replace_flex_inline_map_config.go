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
var _ action.Action = (*ReplaceFlexInlineMapConfigAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*ReplaceFlexInlineMapConfigAction)(nil)

// ReplaceFlexInlineMapConfigAction is the generated Terraform action implementation.
type ReplaceFlexInlineMapConfigAction struct {
	client *client.Client
}

// ReplaceFlexInlineMapConfigActionModel describes the action configuration shape.
type ReplaceFlexInlineMapConfigActionModel struct {
	Alias               types.String  `tfsdk:"alias"`
	BypassStrategy      types.String  `tfsdk:"bypass_strategy"`
	ClusterId           types.String  `tfsdk:"cluster_id" json:"clusterId"`
	ConfigData          types.Dynamic `tfsdk:"config_data" json:"configData"`
	ConfigStatus        types.String  `tfsdk:"config_status" json:"configStatus"`
	ConfigStatusReasons types.List    `tfsdk:"config_status_reasons" json:"configStatusReasons"`
	ConfigType          types.String  `tfsdk:"config_type" json:"configType"`
	HealthState         types.String  `tfsdk:"health_state" json:"healthState"`
	HealthStateReasons  types.List    `tfsdk:"health_state_reasons" json:"healthStateReasons"`
	SolutionAlias       types.String  `tfsdk:"solution_alias" json:"solutionAlias"`
}

// NewReplaceFlexInlineMapConfigAction returns a new instance of the generated action.
func NewReplaceFlexInlineMapConfigAction() action.Action {
	return &ReplaceFlexInlineMapConfigAction{}
}

// Metadata returns the action type name.
func (r *ReplaceFlexInlineMapConfigAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_replace_flex_inline_map_config"
}

// Schema returns the action schema.
func (r *ReplaceFlexInlineMapConfigAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Replace the existing map config with the provided configs", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Updates maps for the given solution", Required: true}, "bypass_strategy": schema.StringAttribute{MarkdownDescription: "if bypass, Inline network's traffic path would be bypass during the solution deployment and would return to the original state on successful deployment, If skipBypass, Inline network's traffic path would not be changed during the deployment", Optional: true}, "cluster_id": schema.StringAttribute{Optional: true}, "config_data": schema.DynamicAttribute{MarkdownDescription: "Data holding configuration details", Optional: true}, "config_status": schema.StringAttribute{MarkdownDescription: "Status of the created config object", Optional: true}, "config_status_reasons": schema.ListAttribute{Optional: true, ElementType: types.StringType}, "config_type": schema.StringAttribute{MarkdownDescription: "Type of the configuration object", Optional: true}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Optional: true}, "health_state_reasons": schema.ListNestedAttribute{Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Optional: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Optional: true}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Optional: true}}}}, "solution_alias": schema.StringAttribute{MarkdownDescription: "Alias of the solution", Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *ReplaceFlexInlineMapConfigAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config ReplaceFlexInlineMapConfigActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *ReplaceFlexInlineMapConfigAction) invokeRemote(ctx context.Context, config *ReplaceFlexInlineMapConfigActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/flexInline/{alias}/maps"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_replace_flex_inline_map_config", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_replace_flex_inline_map_config", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_replace_flex_inline_map_config", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.BypassStrategy.IsNull() {
		query.Set("bypassStrategy", config.BypassStrategy.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_replace_flex_inline_map_config", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_replace_flex_inline_map_config", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_replace_flex_inline_map_config", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_replace_flex_inline_map_config", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_replace_flex_inline_map_config", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_replace_flex_inline_map_config", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_replace_flex_inline_map_config", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_replace_flex_inline_map_config", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_replace_flex_inline_map_config", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_replace_flex_inline_map_config", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *ReplaceFlexInlineMapConfigAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
