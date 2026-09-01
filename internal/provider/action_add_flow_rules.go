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
var _ action.Action = (*AddFlowRulesAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*AddFlowRulesAction)(nil)

// AddFlowRulesAction is the generated Terraform action implementation.
type AddFlowRulesAction struct {
	client *client.Client
}

// AddFlowRulesActionModel describes the action configuration shape.
type AddFlowRulesActionModel struct {
	Alias                       types.String `tfsdk:"alias"`
	FlowAlias                   types.String `tfsdk:"flow_alias"`
	FlowRules                   types.Object `tfsdk:"flow_rules" json:"flowRules"`
	FlowSample5GOverlapRules    types.Object `tfsdk:"flow_sample5_g_overlap_rules" json:"flowSample5gOverlapRules"`
	FlowSample5GRules           types.Object `tfsdk:"flow_sample5_g_rules" json:"flowSample5gRules"`
	FlowSampleDiameterRules     types.Object `tfsdk:"flow_sample_diameter_rules" json:"flowSampleDiameterRules"`
	FlowSampleOverlapRules      types.Object `tfsdk:"flow_sample_overlap_rules" json:"flowSampleOverlapRules"`
	FlowSampleRules             types.Object `tfsdk:"flow_sample_rules" json:"flowSampleRules"`
	FlowSampleSipRules          types.Object `tfsdk:"flow_sample_sip_rules" json:"flowSampleSipRules"`
	FlowWhitelist5GOverlapRules types.Object `tfsdk:"flow_whitelist5_g_overlap_rules" json:"flowWhitelist5gOverlapRules"`
	FlowWhitelist5GRules        types.Object `tfsdk:"flow_whitelist5_g_rules" json:"flowWhitelist5gRules"`
	FlowWhitelistOverlapRules   types.Object `tfsdk:"flow_whitelist_overlap_rules" json:"flowWhitelistOverlapRules"`
	FlowWhitelistRules          types.Object `tfsdk:"flow_whitelist_rules" json:"flowWhitelistRules"`
	GsRules                     types.Object `tfsdk:"gs_rules" json:"gsRules"`
	RuleType                    types.String `tfsdk:"rule_type"`
	SubFlowAlias                types.String `tfsdk:"sub_flow_alias"`
}

// NewAddFlowRulesAction returns a new instance of the generated action.
func NewAddFlowRulesAction() action.Action {
	return &AddFlowRulesAction{}
}

// Metadata returns the action type name.
func (r *AddFlowRulesAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_add_flow_rules"
}

// Schema returns the action schema.
func (r *AddFlowRulesAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Add rules within a flow", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Traffic Flows alias or ID", Required: true}, "flow_alias": schema.StringAttribute{MarkdownDescription: "Flow alias or ID", Required: true}, "flow_rules": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Rules Container. Private class", Optional: true, Attributes: map[string]schema.Attribute{"drop_rules": schema.SetNestedAttribute{Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"gtp": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Rule GTP match Definition. Private class", Required: true, Attributes: map[string]schema.Attribute{"imei": schema.StringAttribute{MarkdownDescription: "mutually exclusive with 'imsi' and 'msisdn'. If '*' is added at the end of the value, it is treated as prefix", Optional: true}, "imsi": schema.StringAttribute{MarkdownDescription: "mutually exclusive with 'imei' and 'msisdn'. If '*' is added at the end of the value, it is treated as prefix", Optional: true}, "interface": schema.StringAttribute{MarkdownDescription: "interface type. Mutually exclusive with version", Optional: true}, "msisdn": schema.StringAttribute{MarkdownDescription: "mutually exclusive with 'imsi' and 'imei'. If '*' is added at the end of the value, it is treated as prefix", Optional: true}, "version": schema.StringAttribute{MarkdownDescription: "mutually exclusive with interface", Optional: true}}}, "rule_id": schema.Int64Attribute{Required: true}}}}, "pass_rules": schema.SetNestedAttribute{Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"gtp": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Rule GTP match Definition. Private class", Required: true, Attributes: map[string]schema.Attribute{"imei": schema.StringAttribute{MarkdownDescription: "mutually exclusive with 'imsi' and 'msisdn'. If '*' is added at the end of the value, it is treated as prefix", Optional: true}, "imsi": schema.StringAttribute{MarkdownDescription: "mutually exclusive with 'imei' and 'msisdn'. If '*' is added at the end of the value, it is treated as prefix", Optional: true}, "interface": schema.StringAttribute{MarkdownDescription: "interface type. Mutually exclusive with version", Optional: true}, "msisdn": schema.StringAttribute{MarkdownDescription: "mutually exclusive with 'imsi' and 'imei'. If '*' is added at the end of the value, it is treated as prefix", Optional: true}, "version": schema.StringAttribute{MarkdownDescription: "mutually exclusive with interface", Optional: true}}}, "rule_id": schema.Int64Attribute{Required: true}}}}}}, "flow_sample5_g_overlap_rules": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Sample 5g Overlap Rules Container. Private class", Optional: true, Attributes: map[string]schema.Attribute{"pass_rules": schema.SetNestedAttribute{Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"comment": schema.StringAttribute{Optional: true}, "flow5_g": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Sample Overlap Rule 5g match Definition. Private class", Required: true, Attributes: map[string]schema.Attribute{"dnn": schema.StringAttribute{MarkdownDescription: "Access Point Name pattern.  Alphanumeric, '.', '-', and '*' allowed.", Optional: true}, "gpsi": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true}, "nas_5_qi": schema.StringAttribute{MarkdownDescription: "5G QoS Indicator. Valid 5qi value <1 - 255>", Optional: true}, "nci": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true}, "nsiid": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the SD value, it is treated as prefix", Optional: true}, "pei": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true}, "plmn_id": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the MNC value, it is treated as prefix", Optional: true}, "supi": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true}, "tac": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true}}}, "percentage": schema.Int64Attribute{Required: true}, "rule_id": schema.Int64Attribute{Required: true}}}}}}, "flow_sample5_g_rules": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Sample 5g Rules Container. Private class", Optional: true, Attributes: map[string]schema.Attribute{"pass_rules": schema.SetNestedAttribute{Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"comment": schema.StringAttribute{Optional: true}, "flow5_g": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Sample Rule 5g match Definition. Private class", Required: true, Attributes: map[string]schema.Attribute{"dnn": schema.StringAttribute{MarkdownDescription: "Access Point Name pattern.  Alphanumeric, '.', '-', and '*' allowed.", Optional: true}, "gpsi": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true}, "nci": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true}, "nsiid": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the SD value, it is treated as prefix", Optional: true}, "pei": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true}, "plmn_id": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the MNC value, it is treated as prefix", Optional: true}, "supi": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true}, "tac": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true}}}, "percentage": schema.Int64Attribute{Required: true}, "priority": schema.Int64Attribute{Optional: true}, "rule_id": schema.Int64Attribute{Required: true}}}}}}, "flow_sample_diameter_rules": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Sample Diameter Rules Container. Private class", Optional: true, Attributes: map[string]schema.Attribute{"pass_rules": schema.SetNestedAttribute{Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"diameter": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Sample Diameter Rule Definition", Required: true, Attributes: map[string]schema.Attribute{"user_name": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true}}}, "interface": schema.StringAttribute{MarkdownDescription: "interface type", Required: true}, "percentage": schema.Int64Attribute{Required: true}, "rule_id": schema.Int64Attribute{Required: true}}}}}}, "flow_sample_overlap_rules": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Sample Overlap Rules Container. Private class", Optional: true, Attributes: map[string]schema.Attribute{"pass_rules": schema.SetNestedAttribute{Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"comment": schema.StringAttribute{Optional: true}, "gtp": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Sample Rule GTP match Definition. Private class", Required: true, Attributes: map[string]schema.Attribute{"apn": schema.StringAttribute{MarkdownDescription: "Access Point Name pattern.  Alphanumeric, '.', '-', and '*' allowed.", Optional: true}, "eci": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true}, "imei": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true}, "imsi": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true}, "interface": schema.StringAttribute{MarkdownDescription: "interface type. Mutually exclusive with version", Optional: true}, "msisdn": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true}, "nas_5_qi": schema.StringAttribute{MarkdownDescription: "5G QoS Indicator. Valid 5qi value <1 - 255>", Optional: true}, "nci": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true}, "plmn_id": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the MNC value, it is treated as prefix", Optional: true}, "qci": schema.Int64Attribute{MarkdownDescription: "QoS Class Indicator", Optional: true}, "snssai": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the SD value, it is treated as prefix", Optional: true}, "tac": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true}, "tac_5_g": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true}, "version": schema.StringAttribute{MarkdownDescription: "mutually exclusive with interface", Optional: true}}}, "percentage": schema.Int64Attribute{Required: true}, "periodic_recalc": schema.BoolAttribute{MarkdownDescription: "Enable Periodic Recalc for rotational sampling. Map look up in the data path based on this flag", Optional: true}, "priority": schema.Int64Attribute{MarkdownDescription: "maximum value is equal to the number of rules upon completion of the request", Optional: true}, "rule_id": schema.Int64Attribute{Required: true}}}}}}, "flow_sample_rules": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Sample Rules Container. Private class", Optional: true, Attributes: map[string]schema.Attribute{"pass_rules": schema.SetNestedAttribute{Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"comment": schema.StringAttribute{Optional: true}, "gtp": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Sample Rule GTP match Definition. Private class", Required: true, Attributes: map[string]schema.Attribute{"apn": schema.StringAttribute{MarkdownDescription: "Access Point Name pattern.  Alphanumeric, '.', '-', and '*' allowed.", Optional: true}, "eci": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true}, "imei": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true}, "imsi": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true}, "interface": schema.StringAttribute{MarkdownDescription: "interface type. Mutually exclusive with version", Optional: true}, "msisdn": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true}, "nas_5_qi": schema.StringAttribute{MarkdownDescription: "5G QoS Indicator. Valid 5qi value <1 - 255>", Optional: true}, "nci": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true}, "plmn_id": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the MNC value, it is treated as prefix", Optional: true}, "qci": schema.Int64Attribute{MarkdownDescription: "QoS Class Indicator", Optional: true}, "snssai": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the SD value, it is treated as prefix", Optional: true}, "tac": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true}, "tac_5_g": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true}, "version": schema.StringAttribute{MarkdownDescription: "mutually exclusive with interface", Optional: true}}}, "percentage": schema.Int64Attribute{Required: true}, "periodic_recalc": schema.BoolAttribute{MarkdownDescription: "Enable Periodic Recalc for rotational sampling. Map look up in the data path based on this flag", Optional: true}, "priority": schema.Int64Attribute{MarkdownDescription: "maximum value is equal to the number of rules upon completion of the request", Optional: true}, "rule_id": schema.Int64Attribute{Required: true}}}}}}, "flow_sample_sip_rules": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Sample Sip Rules Container. Private class", Optional: true, Attributes: map[string]schema.Attribute{"pass_rules": schema.SetNestedAttribute{Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"percentage": schema.Int64Attribute{Required: true}, "rule_id": schema.Int64Attribute{Required: true}, "sip": schema.SingleNestedAttribute{Required: true, Attributes: map[string]schema.Attribute{"callee_id": schema.StringAttribute{MarkdownDescription: "sip callee id", Optional: true}, "callee_id_range": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"max_value": schema.StringAttribute{Required: true}, "value": schema.StringAttribute{Required: true}}}, "caller_id": schema.StringAttribute{MarkdownDescription: "sip caller id", Optional: true}, "caller_id_range": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"max_value": schema.StringAttribute{Required: true}, "value": schema.StringAttribute{Required: true}}}, "id_range": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"max_value": schema.StringAttribute{Required: true}, "value": schema.StringAttribute{Required: true}}}}}}}}}}, "flow_whitelist5_g_overlap_rules": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Whitelist 5g Overlap Rule match Definition. Private class", Optional: true, Attributes: map[string]schema.Attribute{"dnn": schema.StringAttribute{MarkdownDescription: "Domain Network Name pattern.  Alphanumeric, '.', '-', and '*' allowed.", Optional: true}, "type": schema.StringAttribute{MarkdownDescription: "Set 5G WL-DB lookup type", Optional: true}}}, "flow_whitelist5_g_rules": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Whitelist 5g Rule GTP match Definition. Private class", Optional: true, Attributes: map[string]schema.Attribute{"dnn": schema.StringAttribute{MarkdownDescription: "Domain Network Name pattern.  Alphanumeric, '.', '-', and '*' allowed.", Optional: true}, "type": schema.StringAttribute{MarkdownDescription: "Set 5G WL-DB lookup type", Optional: true}, "whitelist_databases": schema.ListAttribute{MarkdownDescription: "Attach whitelist databases to the map", Optional: true, ElementType: types.StringType}}}, "flow_whitelist_overlap_rules": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Whitelist Overlap Rules Container. Private class", Optional: true, Attributes: map[string]schema.Attribute{"pass_rules": schema.SetNestedAttribute{Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"flow5_g": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Whitelist 5g Rule GTP match Definition. Private class", Optional: true, Attributes: map[string]schema.Attribute{"dnn": schema.StringAttribute{MarkdownDescription: "Domain Network Name pattern.  Alphanumeric, '.', '-', and '*' allowed.", Optional: true}, "type": schema.StringAttribute{MarkdownDescription: "Set 5G WL-DB lookup type", Optional: true}, "whitelist_databases": schema.ListAttribute{MarkdownDescription: "Attach whitelist databases to the map", Optional: true, ElementType: types.StringType}}}, "gtp": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Whitelist Rule GTP match Definition. Private class", Optional: true, Attributes: map[string]schema.Attribute{"apn": schema.StringAttribute{MarkdownDescription: "Access Point Name pattern.  Alphanumeric, '.', '-', and '*' allowed.", Optional: true}, "interface": schema.StringAttribute{MarkdownDescription: "interface type. Mutually exclusive with version. required till H 5.6", Optional: true}, "type": schema.StringAttribute{MarkdownDescription: "Set GTP WL-DB lookup type", Optional: true}, "version": schema.StringAttribute{MarkdownDescription: "mutually exclusive with interface", Optional: true}, "whitelist_databases": schema.ListAttribute{MarkdownDescription: "Attach whitelist databases to the map", Optional: true, ElementType: types.StringType}}}, "rule_id": schema.Int64Attribute{Required: true}, "sip": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Whitelist Rule Sip match definition", Optional: true, Attributes: map[string]schema.Attribute{"type": schema.StringAttribute{MarkdownDescription: "all:Whitelist based on caller/callee/source/destination IP address, bothAddr: Whitelist source/destination IP address, bothId:Whitelist Caller/Callee Id's, calleeId: Whitelist Callee ID, callerId: Whitelist Caller ID, destIp: Whitelist based on destination IP address, srcIp: Whitelist based on Source IP address", Optional: true}}}}}}}}, "flow_whitelist_rules": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Whitelist Rules Container. Private class", Optional: true, Attributes: map[string]schema.Attribute{"pass_rules": schema.SetNestedAttribute{Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"flow5_g": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Whitelist 5g Rule GTP match Definition. Private class", Optional: true, Attributes: map[string]schema.Attribute{"dnn": schema.StringAttribute{MarkdownDescription: "Domain Network Name pattern.  Alphanumeric, '.', '-', and '*' allowed.", Optional: true}, "type": schema.StringAttribute{MarkdownDescription: "Set 5G WL-DB lookup type", Optional: true}, "whitelist_databases": schema.ListAttribute{MarkdownDescription: "Attach whitelist databases to the map", Optional: true, ElementType: types.StringType}}}, "gtp": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Whitelist Rule GTP match Definition. Private class", Optional: true, Attributes: map[string]schema.Attribute{"apn": schema.StringAttribute{MarkdownDescription: "Access Point Name pattern.  Alphanumeric, '.', '-', and '*' allowed.", Optional: true}, "interface": schema.StringAttribute{MarkdownDescription: "interface type. Mutually exclusive with version. required till H 5.6", Optional: true}, "type": schema.StringAttribute{MarkdownDescription: "Set GTP WL-DB lookup type", Optional: true}, "version": schema.StringAttribute{MarkdownDescription: "mutually exclusive with interface", Optional: true}, "whitelist_databases": schema.ListAttribute{MarkdownDescription: "Attach whitelist databases to the map", Optional: true, ElementType: types.StringType}}}, "rule_id": schema.Int64Attribute{Required: true}, "sip": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Whitelist Rule Sip match definition", Optional: true, Attributes: map[string]schema.Attribute{"type": schema.StringAttribute{MarkdownDescription: "all:Whitelist based on caller/callee/source/destination IP address, bothAddr: Whitelist source/destination IP address, bothId:Whitelist Caller/Callee Id's, calleeId: Whitelist Callee ID, callerId: Whitelist Caller ID, destIp: Whitelist based on destination IP address, srcIp: Whitelist based on Source IP address", Optional: true}}}}}}}}, "gs_rules": schema.SingleNestedAttribute{MarkdownDescription: "Map GigaSMART Rules Container. Private class", Optional: true, Attributes: map[string]schema.Attribute{"drop_rules": schema.DynamicAttribute{Optional: true}, "pass_rules": schema.DynamicAttribute{Optional: true}}}, "rule_type": schema.StringAttribute{MarkdownDescription: "Type of rule", Required: true}, "sub_flow_alias": schema.StringAttribute{MarkdownDescription: "Sub-flow alias or ID", Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *AddFlowRulesAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config AddFlowRulesActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *AddFlowRulesAction) invokeRemote(ctx context.Context, config *AddFlowRulesActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/traffic-flows/{alias}/flow/{flowAlias}/subFlow/{subFlowAlias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	reqPath = strings.ReplaceAll(reqPath, "{flowAlias}", url.PathEscape(config.FlowAlias.ValueString()))
	reqPath = strings.ReplaceAll(reqPath, "{subFlowAlias}", url.PathEscape(config.SubFlowAlias.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_add_flow_rules", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_add_flow_rules", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_add_flow_rules", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("ruleType", config.RuleType.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_add_flow_rules", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_flow_rules", "Invalid request")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_flow_rules", "Not Authenticated")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_flow_rules", "Access Denied")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_flow_rules", "Entity Not Found")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_flow_rules", "Conflict")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_flow_rules", "Internal Server Error")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_add_flow_rules", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_flow_rules", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *AddFlowRulesAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
