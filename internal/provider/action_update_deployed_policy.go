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
var _ action.Action = (*UpdateDeployedPolicyAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*UpdateDeployedPolicyAction)(nil)

// UpdateDeployedPolicyAction is the generated Terraform action implementation.
type UpdateDeployedPolicyAction struct {
	client *client.Client
}

// UpdateDeployedPolicyActionModel describes the action configuration shape.
type UpdateDeployedPolicyActionModel struct {
	BodyName      types.String  `tfsdk:"body_name" json:"name"`
	Comment       types.String  `tfsdk:"comment"`
	DestPorts     types.Dynamic `tfsdk:"dest_ports" json:"destPorts"`
	GigasmartInfo types.Dynamic `tfsdk:"gigasmart_info" json:"gigasmartInfo"`
	Name          types.String  `tfsdk:"name"`
	PolicyId      types.String  `tfsdk:"policy_id" json:"policyId"`
	Priority      types.String  `tfsdk:"priority"`
	RuleCriteria  types.Dynamic `tfsdk:"rule_criteria" json:"ruleCriteria"`
	RulesInfo     types.Dynamic `tfsdk:"rules_info" json:"rulesInfo"`
	SrcPorts      types.Dynamic `tfsdk:"src_ports" json:"srcPorts"`
	Tags          types.Dynamic `tfsdk:"tags"`
}

// NewUpdateDeployedPolicyAction returns a new instance of the generated action.
func NewUpdateDeployedPolicyAction() action.Action {
	return &UpdateDeployedPolicyAction{}
}

// Metadata returns the action type name.
func (r *UpdateDeployedPolicyAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_update_deployed_policy"
}

// Schema returns the action schema.
func (r *UpdateDeployedPolicyAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Update deployed policy", Attributes: map[string]schema.Attribute{"body_name": schema.StringAttribute{Optional: true}, "comment": schema.StringAttribute{Optional: true}, "dest_ports": schema.DynamicAttribute{Optional: true}, "gigasmart_info": schema.DynamicAttribute{Optional: true}, "name": schema.StringAttribute{MarkdownDescription: "policy name", Required: true}, "policy_id": schema.StringAttribute{Optional: true}, "priority": schema.StringAttribute{Optional: true}, "rule_criteria": schema.DynamicAttribute{Optional: true}, "rules_info": schema.DynamicAttribute{Optional: true}, "src_ports": schema.DynamicAttribute{Optional: true}, "tags": schema.DynamicAttribute{Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *UpdateDeployedPolicyAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UpdateDeployedPolicyActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *UpdateDeployedPolicyAction) invokeRemote(ctx context.Context, config *UpdateDeployedPolicyActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/intent/policies/{name}/update"
	reqPath = strings.ReplaceAll(reqPath, "{name}", url.PathEscape(config.Name.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_deployed_policy", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_deployed_policy", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_deployed_policy", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_deployed_policy", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		apiErr, err := client.NewAPIError(httpResp)
		if err != nil {
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_deployed_policy", fmt.Sprintf("Could not read error response: %s", err))
			return
		}
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_deployed_policy", apiErr.Error())
		return
	}
}

// Configure stores the API client supplied by the provider.
func (r *UpdateDeployedPolicyAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
