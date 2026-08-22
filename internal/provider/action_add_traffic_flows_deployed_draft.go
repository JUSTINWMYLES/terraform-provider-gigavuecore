package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*AddTrafficFlowsDeployedDraftAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*AddTrafficFlowsDeployedDraftAction)(nil)

// AddTrafficFlowsDeployedDraftAction is the generated Terraform action implementation.
type AddTrafficFlowsDeployedDraftAction struct {
	client *client.Client
}

// AddTrafficFlowsDeployedDraftActionModel describes the action configuration shape.
type AddTrafficFlowsDeployedDraftActionModel struct {
	Alias           types.String  `tfsdk:"alias"`
	Comment         types.String  `tfsdk:"comment"`
	DeploymentType  types.String  `tfsdk:"deployment_type" json:"deploymentType"`
	Enable          types.Bool    `tfsdk:"enable"`
	Flows           types.Dynamic `tfsdk:"flows"`
	HasDraft        types.Bool    `tfsdk:"has_draft" json:"hasDraft"`
	PriorityType    types.String  `tfsdk:"priority_type" json:"priorityType"`
	SourcesAndRules types.Dynamic `tfsdk:"sources_and_rules" json:"sourcesAndRules"`
	Tags            types.Dynamic `tfsdk:"tags"`
}

// NewAddTrafficFlowsDeployedDraftAction returns a new instance of the generated action.
func NewAddTrafficFlowsDeployedDraftAction() action.Action {
	return &AddTrafficFlowsDeployedDraftAction{}
}

// Metadata returns the action type name.
func (r *AddTrafficFlowsDeployedDraftAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_add_traffic_flows_deployed_draft"
}

// Schema returns the action schema.
func (r *AddTrafficFlowsDeployedDraftAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Add a new deployed traffic flows draft", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "unique map alias", Required: true}, "comment": schema.StringAttribute{Optional: true}, "deployment_type": schema.StringAttribute{MarkdownDescription: "'regular' maps are from network/hybrid ports to tool/hybrid/gigastream; 'inline' maps are from inline ports to inline/tool/hybrid/gigastream; 'firstLevel' are from network/hybrid ports to vPorts/tool/hybrid/gigastream; 'secondLevel' maps are from vPorts to tool/hybrid/gigastream; 'inlineFirstLevel is from inline network to vport; 'inlineSecondLevel' is from vport to inline tool; 'transitLevel' is from vport to vport", Required: true}, "enable": schema.BoolAttribute{MarkdownDescription: "enable/disable map, applicable only to first level maps", Optional: true}, "flows": schema.DynamicAttribute{Required: true}, "has_draft": schema.BoolAttribute{Optional: true}, "priority_type": schema.StringAttribute{MarkdownDescription: "Define the map priority to be LOWEST or HIGHEST. Default priority is LOWEST.", Optional: true}, "sources_and_rules": schema.DynamicAttribute{Required: true}, "tags": schema.DynamicAttribute{Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *AddTrafficFlowsDeployedDraftAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config AddTrafficFlowsDeployedDraftActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *AddTrafficFlowsDeployedDraftAction) invokeRemote(ctx context.Context, config *AddTrafficFlowsDeployedDraftActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/traffic-flows/deployedDraft/save"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_add_traffic_flows_deployed_draft", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_add_traffic_flows_deployed_draft", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_add_traffic_flows_deployed_draft", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_add_traffic_flows_deployed_draft", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_traffic_flows_deployed_draft", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_traffic_flows_deployed_draft", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_traffic_flows_deployed_draft", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_traffic_flows_deployed_draft", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_traffic_flows_deployed_draft", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_traffic_flows_deployed_draft", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_traffic_flows_deployed_draft", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_add_traffic_flows_deployed_draft", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_traffic_flows_deployed_draft", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *AddTrafficFlowsDeployedDraftAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
