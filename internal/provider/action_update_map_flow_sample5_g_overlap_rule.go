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
var _ action.Action = (*UpdateMapFlowSample5GOverlapRuleAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*UpdateMapFlowSample5GOverlapRuleAction)(nil)

// UpdateMapFlowSample5GOverlapRuleAction is the generated Terraform action implementation.
type UpdateMapFlowSample5GOverlapRuleAction struct {
	client *client.Client
}

// UpdateMapFlowSample5GOverlapRuleActionModel describes the action configuration shape.
type UpdateMapFlowSample5GOverlapRuleActionModel struct {
	Alias      types.String `tfsdk:"alias"`
	BodyRuleId types.Int64  `tfsdk:"body_rule_id" json:"ruleId"`
	Comment    types.String `tfsdk:"comment"`
	Flow5G     types.Object `tfsdk:"flow5_g" json:"flow5g"`
	Percentage types.Int64  `tfsdk:"percentage"`
	RuleId     types.String `tfsdk:"rule_id"`
}

// NewUpdateMapFlowSample5GOverlapRuleAction returns a new instance of the generated action.
func NewUpdateMapFlowSample5GOverlapRuleAction() action.Action {
	return &UpdateMapFlowSample5GOverlapRuleAction{}
}

// Metadata returns the action type name.
func (r *UpdateMapFlowSample5GOverlapRuleAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_update_map_flow_sample5_g_overlap_rule"
}

// Schema returns the action schema.
func (r *UpdateMapFlowSample5GOverlapRuleAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "update flowSample5gOverlapRule of a 'secondLevel/flowSample5gOverlap' map", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "alias of the target map", Required: true}, "body_rule_id": schema.Int64Attribute{Required: true}, "comment": schema.StringAttribute{Optional: true}, "flow5_g": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Sample Overlap Rule 5g match Definition. Private class", Required: true, Attributes: map[string]schema.Attribute{"dnn": schema.StringAttribute{MarkdownDescription: "Access Point Name pattern.  Alphanumeric, '.', '-', and '*' allowed.", Optional: true}, "gpsi": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true}, "nas_5_qi": schema.StringAttribute{MarkdownDescription: "5G QoS Indicator. Valid 5qi value <1 - 255>", Optional: true}, "nci": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true}, "nsiid": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the SD value, it is treated as prefix", Optional: true}, "pei": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true}, "plmn_id": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the MNC value, it is treated as prefix", Optional: true}, "supi": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true}, "tac": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true}}}, "percentage": schema.Int64Attribute{Required: true}, "rule_id": schema.StringAttribute{MarkdownDescription: "id of the rule to update", Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *UpdateMapFlowSample5GOverlapRuleAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UpdateMapFlowSample5GOverlapRuleActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *UpdateMapFlowSample5GOverlapRuleAction) invokeRemote(ctx context.Context, config *UpdateMapFlowSample5GOverlapRuleActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/maps/{alias}/flowSample5gOverlapRules/pass/{ruleId}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	reqPath = strings.ReplaceAll(reqPath, "{ruleId}", url.PathEscape(config.RuleId.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_map_flow_sample5_g_overlap_rule", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_map_flow_sample5_g_overlap_rule", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_map_flow_sample5_g_overlap_rule", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_map_flow_sample5_g_overlap_rule", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_map_flow_sample5_g_overlap_rule", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_map_flow_sample5_g_overlap_rule", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_map_flow_sample5_g_overlap_rule", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_map_flow_sample5_g_overlap_rule", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_map_flow_sample5_g_overlap_rule", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_map_flow_sample5_g_overlap_rule", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_map_flow_sample5_g_overlap_rule", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_update_map_flow_sample5_g_overlap_rule", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_map_flow_sample5_g_overlap_rule", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *UpdateMapFlowSample5GOverlapRuleAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
