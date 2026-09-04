package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*DeployPolicyAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*DeployPolicyAction)(nil)

// DeployPolicyAction is the generated Terraform action implementation.
type DeployPolicyAction struct {
	client *client.Client
}

// DeployPolicyActionModel describes the action configuration shape.
type DeployPolicyActionModel struct {
	Comment            types.String  `tfsdk:"comment"`
	Deployed           types.Bool    `tfsdk:"deployed"`
	DeploymentError    types.String  `tfsdk:"deployment_error" json:"deploymentError"`
	DeploymentPercent  types.String  `tfsdk:"deployment_percent" json:"deploymentPercent"`
	DestPortTimestamp  types.String  `tfsdk:"dest_port_timestamp" json:"destPortTimestamp"`
	HealthState        types.String  `tfsdk:"health_state" json:"healthState"`
	HealthStateReasons types.List    `tfsdk:"health_state_reasons" json:"healthStateReasons"`
	Name               types.String  `tfsdk:"name"`
	PolicyId           types.String  `tfsdk:"policy_id" json:"policyId"`
	PolicyTimestamp    types.String  `tfsdk:"policy_timestamp" json:"policyTimestamp"`
	Priority           types.Bool    `tfsdk:"priority"`
	Rules              types.Dynamic `tfsdk:"rules"`
	SrcPortTimestamp   types.String  `tfsdk:"src_port_timestamp" json:"srcPortTimestamp"`
	SrcPortsInfo       types.Object  `tfsdk:"src_ports_info" json:"srcPortsInfo"`
	Tags               types.List    `tfsdk:"tags"`
}

// NewDeployPolicyAction returns a new instance of the generated action.
func NewDeployPolicyAction() action.Action {
	return &DeployPolicyAction{}
}

// Metadata returns the action type name.
func (r *DeployPolicyAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_deploy_policy"
}

// Schema returns the action schema.
func (r *DeployPolicyAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Deploy a new policy", Attributes: map[string]schema.Attribute{"comment": schema.StringAttribute{Optional: true}, "deployed": schema.BoolAttribute{MarkdownDescription: "policy is deployed or not", Optional: true}, "deployment_error": schema.StringAttribute{MarkdownDescription: "policy deployment error message", Optional: true}, "deployment_percent": schema.StringAttribute{MarkdownDescription: "policy deployment percentage", Optional: true}, "dest_port_timestamp": schema.StringAttribute{Optional: true}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Optional: true}, "health_state_reasons": schema.ListNestedAttribute{Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Optional: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Optional: true}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Optional: true}}}}, "name": schema.StringAttribute{MarkdownDescription: "policy name", Required: true}, "policy_id": schema.StringAttribute{MarkdownDescription: "generated unique policy ID", Optional: true}, "policy_timestamp": schema.StringAttribute{Optional: true}, "priority": schema.BoolAttribute{Optional: true}, "rules": schema.DynamicAttribute{Optional: true}, "src_port_timestamp": schema.StringAttribute{Optional: true}, "src_ports_info": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"comment": schema.StringAttribute{Optional: true}, "ports": schema.ListAttribute{Optional: true, ElementType: types.StringType}, "template_ids": schema.ListAttribute{Optional: true, ElementType: types.StringType}}}, "tags": schema.ListNestedAttribute{Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"tag_key": schema.StringAttribute{MarkdownDescription: "Name of the tag", Required: true}, "tag_values": schema.ListAttribute{MarkdownDescription: "All possible values of the tag", Required: true, ElementType: types.StringType}}}}}}
}

// Invoke executes the action against the remote API.
func (r *DeployPolicyAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config DeployPolicyActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *DeployPolicyAction) invokeRemote(ctx context.Context, config *DeployPolicyActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/intent/policies/deploy"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_deploy_policy", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_deploy_policy", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_deploy_policy", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_deploy_policy", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		apiErr, err := client.NewAPIError(httpResp)
		if err != nil {
			resp.Diagnostics.AddError("Error invoking gigavuecore_deploy_policy", fmt.Sprintf("Could not read error response: %s", err))
			return
		}
		resp.Diagnostics.AddError("Error invoking gigavuecore_deploy_policy", apiErr.Error())
		return
	}
}

// Configure stores the API client supplied by the provider.
func (r *DeployPolicyAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
