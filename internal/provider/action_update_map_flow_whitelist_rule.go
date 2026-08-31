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
var _ action.Action = (*UpdateMapFlowWhitelistRuleAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*UpdateMapFlowWhitelistRuleAction)(nil)

// UpdateMapFlowWhitelistRuleAction is the generated Terraform action implementation.
type UpdateMapFlowWhitelistRuleAction struct {
	client *client.Client
}

// UpdateMapFlowWhitelistRuleActionModel describes the action configuration shape.
type UpdateMapFlowWhitelistRuleActionModel struct {
	Alias      types.String `tfsdk:"alias"`
	BodyRuleId types.Int64  `tfsdk:"body_rule_id" json:"ruleId"`
	Flow5G     types.Object `tfsdk:"flow5_g" json:"flow5g"`
	Gtp        types.Object `tfsdk:"gtp"`
	RuleId     types.String `tfsdk:"rule_id"`
	Sip        types.Object `tfsdk:"sip"`
}

// NewUpdateMapFlowWhitelistRuleAction returns a new instance of the generated action.
func NewUpdateMapFlowWhitelistRuleAction() action.Action {
	return &UpdateMapFlowWhitelistRuleAction{}
}

// Metadata returns the action type name.
func (r *UpdateMapFlowWhitelistRuleAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_update_map_flow_whitelist_rule"
}

// Schema returns the action schema.
func (r *UpdateMapFlowWhitelistRuleAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "update flowRule of a 'secondLevel/flowWhitelist' map", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "alias of the target map", Required: true}, "body_rule_id": schema.Int64Attribute{Required: true}, "flow5_g": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Whitelist 5g Rule GTP match Definition. Private class", Optional: true, Attributes: map[string]schema.Attribute{"dnn": schema.StringAttribute{MarkdownDescription: "Domain Network Name pattern.  Alphanumeric, '.', '-', and '*' allowed.", Optional: true}, "type": schema.StringAttribute{MarkdownDescription: "Set 5G WL-DB lookup type", Optional: true}, "whitelist_databases": schema.ListAttribute{MarkdownDescription: "Attach whitelist databases to the map", Optional: true, ElementType: types.StringType}}}, "gtp": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Whitelist Rule GTP match Definition. Private class", Optional: true, Attributes: map[string]schema.Attribute{"apn": schema.StringAttribute{MarkdownDescription: "Access Point Name pattern.  Alphanumeric, '.', '-', and '*' allowed.", Optional: true}, "interface": schema.StringAttribute{MarkdownDescription: "interface type. Mutually exclusive with version. required till H 5.6", Optional: true}, "type": schema.StringAttribute{MarkdownDescription: "Set GTP WL-DB lookup type", Optional: true}, "version": schema.StringAttribute{MarkdownDescription: "mutually exclusive with interface", Optional: true}, "whitelist_databases": schema.ListAttribute{MarkdownDescription: "Attach whitelist databases to the map", Optional: true, ElementType: types.StringType}}}, "rule_id": schema.StringAttribute{MarkdownDescription: "id of the rule to update", Required: true}, "sip": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Whitelist Rule Sip match definition", Optional: true, Attributes: map[string]schema.Attribute{"type": schema.StringAttribute{MarkdownDescription: "all:Whitelist based on caller/callee/source/destination IP address, bothAddr: Whitelist source/destination IP address, bothId:Whitelist Caller/Callee Id's, calleeId: Whitelist Callee ID, callerId: Whitelist Caller ID, destIp: Whitelist based on destination IP address, srcIp: Whitelist based on Source IP address", Optional: true}}}}}
}

// Invoke executes the action against the remote API.
func (r *UpdateMapFlowWhitelistRuleAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UpdateMapFlowWhitelistRuleActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *UpdateMapFlowWhitelistRuleAction) invokeRemote(ctx context.Context, config *UpdateMapFlowWhitelistRuleActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/maps/{alias}/flowWhitelistRules/pass/{ruleId}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	reqPath = strings.ReplaceAll(reqPath, "{ruleId}", url.PathEscape(config.RuleId.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_map_flow_whitelist_rule", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_map_flow_whitelist_rule", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_map_flow_whitelist_rule", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_map_flow_whitelist_rule", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_map_flow_whitelist_rule", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_map_flow_whitelist_rule", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_map_flow_whitelist_rule", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_map_flow_whitelist_rule", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_map_flow_whitelist_rule", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_map_flow_whitelist_rule", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_map_flow_whitelist_rule", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_update_map_flow_whitelist_rule", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_map_flow_whitelist_rule", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *UpdateMapFlowWhitelistRuleAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
