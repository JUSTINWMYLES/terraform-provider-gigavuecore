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
var _ action.Action = (*AddMapFlowSampleDiameterRuleAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*AddMapFlowSampleDiameterRuleAction)(nil)

// AddMapFlowSampleDiameterRuleAction is the generated Terraform action implementation.
type AddMapFlowSampleDiameterRuleAction struct {
	client *client.Client
}

// AddMapFlowSampleDiameterRuleActionModel describes the action configuration shape.
type AddMapFlowSampleDiameterRuleActionModel struct {
	Alias      types.String  `tfsdk:"alias"`
	Diameter   types.Dynamic `tfsdk:"diameter"`
	Interface  types.String  `tfsdk:"interface"`
	Percentage types.Int64   `tfsdk:"percentage"`
	RuleId     types.Int64   `tfsdk:"rule_id" json:"ruleId"`
}

// NewAddMapFlowSampleDiameterRuleAction returns a new instance of the generated action.
func NewAddMapFlowSampleDiameterRuleAction() action.Action {
	return &AddMapFlowSampleDiameterRuleAction{}
}

// Metadata returns the action type name.
func (r *AddMapFlowSampleDiameterRuleAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_add_map_flow_sample_diameter_rule"
}

// Schema returns the action schema.
func (r *AddMapFlowSampleDiameterRuleAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Add new flowSampleDiameterRule to a 'secondLevel/flowSampleDiameter' map", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "alias of the target map", Required: true}, "diameter": schema.DynamicAttribute{MarkdownDescription: "Map Flow Sample Diameter Rule Definition", Required: true}, "interface": schema.StringAttribute{MarkdownDescription: "interface type", Required: true}, "percentage": schema.Int64Attribute{Required: true}, "rule_id": schema.Int64Attribute{Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *AddMapFlowSampleDiameterRuleAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config AddMapFlowSampleDiameterRuleActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *AddMapFlowSampleDiameterRuleAction) invokeRemote(ctx context.Context, config *AddMapFlowSampleDiameterRuleActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/maps/{alias}/flowSampleDiameterRules/pass"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_add_map_flow_sample_diameter_rule", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_add_map_flow_sample_diameter_rule", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_add_map_flow_sample_diameter_rule", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_add_map_flow_sample_diameter_rule", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_map_flow_sample_diameter_rule", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_map_flow_sample_diameter_rule", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_map_flow_sample_diameter_rule", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_map_flow_sample_diameter_rule", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_map_flow_sample_diameter_rule", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_map_flow_sample_diameter_rule", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_map_flow_sample_diameter_rule", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_add_map_flow_sample_diameter_rule", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_map_flow_sample_diameter_rule", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *AddMapFlowSampleDiameterRuleAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
