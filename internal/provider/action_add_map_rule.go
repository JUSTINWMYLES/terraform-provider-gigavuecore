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
var _ action.Action = (*AddMapRuleAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*AddMapRuleAction)(nil)

// AddMapRuleAction is the generated Terraform action implementation.
type AddMapRuleAction struct {
	client *client.Client
}

// AddMapRuleActionModel describes the action configuration shape.
type AddMapRuleActionModel struct {
	Alias     types.String  `tfsdk:"alias"`
	Bidi      types.Bool    `tfsdk:"bidi"`
	ClusterId types.String  `tfsdk:"cluster_id"`
	Comment   types.String  `tfsdk:"comment"`
	IpRewrite types.Object  `tfsdk:"ip_rewrite" json:"ipRewrite"`
	Matches   types.Dynamic `tfsdk:"matches"`
	Rewrite   types.Object  `tfsdk:"rewrite"`
	RuleId    types.Int64   `tfsdk:"rule_id" json:"ruleId"`
	RuleType  types.String  `tfsdk:"rule_type"`
	VlanTag   types.Object  `tfsdk:"vlan_tag" json:"vlanTag"`
}

// NewAddMapRuleAction returns a new instance of the generated action.
func NewAddMapRuleAction() action.Action {
	return &AddMapRuleAction{}
}

// Metadata returns the action type name.
func (r *AddMapRuleAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_add_map_rule"
}

// Schema returns the action schema.
func (r *AddMapRuleAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Add new rule to a 'regular/byRule', 'inline/byRule' or 'firstLevel/byRule' map", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "alias of the target map", Required: true}, "bidi": schema.BoolAttribute{Optional: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "comment": schema.StringAttribute{Optional: true}, "ip_rewrite": schema.SingleNestedAttribute{MarkdownDescription: "IpRewrite options on the packets", Optional: true, Attributes: map[string]schema.Attribute{"dst_ip": schema.StringAttribute{Optional: true}, "src_ip": schema.StringAttribute{Optional: true}}}, "matches": schema.DynamicAttribute{MarkdownDescription: "Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched", Required: true}, "rewrite": schema.SingleNestedAttribute{MarkdownDescription: "Rewrite options on the packets", Optional: true, Attributes: map[string]schema.Attribute{"dst_mac": schema.StringAttribute{Optional: true}, "src_mac": schema.StringAttribute{Optional: true}}}, "rule_id": schema.Int64Attribute{Required: true}, "rule_type": schema.StringAttribute{MarkdownDescription: "map rule type", Required: true}, "vlan_tag": schema.SingleNestedAttribute{MarkdownDescription: "This field is only  applicable for pass rules", Optional: true, Attributes: map[string]schema.Attribute{"tag_protocol_id": schema.StringAttribute{Optional: true}, "vlan_action": schema.StringAttribute{Required: true}, "vlan_id": schema.Int64Attribute{Optional: true}}}}}
}

// Invoke executes the action against the remote API.
func (r *AddMapRuleAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config AddMapRuleActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *AddMapRuleAction) invokeRemote(ctx context.Context, config *AddMapRuleActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/maps/{alias}/rules/{ruleType}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	reqPath = strings.ReplaceAll(reqPath, "{ruleType}", url.PathEscape(config.RuleType.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_add_map_rule", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_add_map_rule", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_add_map_rule", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_add_map_rule", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_map_rule", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_map_rule", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_map_rule", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_map_rule", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_map_rule", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_map_rule", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_map_rule", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_add_map_rule", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_map_rule", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *AddMapRuleAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
