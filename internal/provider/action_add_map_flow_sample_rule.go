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
var _ action.Action = (*AddMapFlowSampleRuleAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*AddMapFlowSampleRuleAction)(nil)

// AddMapFlowSampleRuleAction is the generated Terraform action implementation.
type AddMapFlowSampleRuleAction struct {
	client *client.Client
}

// AddMapFlowSampleRuleActionModel describes the action configuration shape.
type AddMapFlowSampleRuleActionModel struct {
	Alias          types.String `tfsdk:"alias"`
	ClusterId      types.String `tfsdk:"cluster_id"`
	Comment        types.String `tfsdk:"comment"`
	Gtp            types.Object `tfsdk:"gtp"`
	Percentage     types.Int64  `tfsdk:"percentage"`
	PeriodicRecalc types.Bool   `tfsdk:"periodic_recalc" json:"periodicRecalc"`
	Priority       types.Int64  `tfsdk:"priority"`
	RuleId         types.Int64  `tfsdk:"rule_id" json:"ruleId"`
}

// NewAddMapFlowSampleRuleAction returns a new instance of the generated action.
func NewAddMapFlowSampleRuleAction() action.Action {
	return &AddMapFlowSampleRuleAction{}
}

// Metadata returns the action type name.
func (r *AddMapFlowSampleRuleAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_add_map_flow_sample_rule"
}

// Schema returns the action schema.
func (r *AddMapFlowSampleRuleAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Add new flowSampleRule to a 'secondLevel/flowSample' map", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "alias of the target map", Required: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "comment": schema.StringAttribute{Optional: true}, "gtp": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Sample Rule GTP match Definition. Private class", Required: true, Attributes: map[string]schema.Attribute{"apn": schema.StringAttribute{MarkdownDescription: "Access Point Name pattern.  Alphanumeric, '.', '-', and '*' allowed.", Optional: true}, "eci": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true}, "imei": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true}, "imsi": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true}, "interface": schema.StringAttribute{MarkdownDescription: "interface type. Mutually exclusive with version", Optional: true}, "msisdn": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true}, "nas_5_qi": schema.StringAttribute{MarkdownDescription: "5G QoS Indicator. Valid 5qi value <1 - 255>", Optional: true}, "nci": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true}, "plmn_id": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the MNC value, it is treated as prefix", Optional: true}, "qci": schema.Int64Attribute{MarkdownDescription: "QoS Class Indicator", Optional: true}, "snssai": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the SD value, it is treated as prefix", Optional: true}, "tac": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true}, "tac_5_g": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true}, "version": schema.StringAttribute{MarkdownDescription: "mutually exclusive with interface", Optional: true}}}, "percentage": schema.Int64Attribute{Required: true}, "periodic_recalc": schema.BoolAttribute{MarkdownDescription: "Enable Periodic Recalc for rotational sampling. Map look up in the data path based on this flag", Optional: true}, "priority": schema.Int64Attribute{MarkdownDescription: "maximum value is equal to the number of rules upon completion of the request", Optional: true}, "rule_id": schema.Int64Attribute{Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *AddMapFlowSampleRuleAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config AddMapFlowSampleRuleActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *AddMapFlowSampleRuleAction) invokeRemote(ctx context.Context, config *AddMapFlowSampleRuleActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/maps/{alias}/flowSampleRules/pass"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_add_map_flow_sample_rule", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_add_map_flow_sample_rule", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_add_map_flow_sample_rule", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_add_map_flow_sample_rule", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_map_flow_sample_rule", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_map_flow_sample_rule", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_map_flow_sample_rule", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_map_flow_sample_rule", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_map_flow_sample_rule", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_map_flow_sample_rule", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_map_flow_sample_rule", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_add_map_flow_sample_rule", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_map_flow_sample_rule", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *AddMapFlowSampleRuleAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
