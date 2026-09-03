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
var _ action.Action = (*UpdateTrafficFlowsGlobalSettingsAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*UpdateTrafficFlowsGlobalSettingsAction)(nil)

// UpdateTrafficFlowsGlobalSettingsAction is the generated Terraform action implementation.
type UpdateTrafficFlowsGlobalSettingsAction struct {
	client *client.Client
}

// UpdateTrafficFlowsGlobalSettingsActionModel describes the action configuration shape.
type UpdateTrafficFlowsGlobalSettingsActionModel struct {
	Body types.Object `tfsdk:"body"`
}

// NewUpdateTrafficFlowsGlobalSettingsAction returns a new instance of the generated action.
func NewUpdateTrafficFlowsGlobalSettingsAction() action.Action {
	return &UpdateTrafficFlowsGlobalSettingsAction{}
}

// Metadata returns the action type name.
func (r *UpdateTrafficFlowsGlobalSettingsAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_update_traffic_flows_global_settings"
}

// Schema returns the action schema.
func (r *UpdateTrafficFlowsGlobalSettingsAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Update global settings for Traffic Flows", Attributes: map[string]schema.Attribute{"body": schema.SingleNestedAttribute{MarkdownDescription: "Settings for traffic flows, including migration and fabric resource configuration.", Required: true, Attributes: map[string]schema.Attribute{"auto_migrate": schema.BoolAttribute{Optional: true}, "fabric_resource": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"mode": schema.StringAttribute{MarkdownDescription: "['SHARE' or 'NOT_SHARE']: Resource sharing mode. SHARE: same resource can be shared by different fabric maps; NOT_SHARE: not shared. Default: NOT_SHARE.", Optional: true}, "scope": schema.StringAttribute{MarkdownDescription: "['GLOBAL']: Scope of resource pool. GLOBAL: only one global resource pool.", Optional: true}, "type": schema.StringAttribute{MarkdownDescription: "['L2CIRCUIT']: Resource type.", Optional: true}}}, "l2_circuit": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"vlan_ids": schema.StringAttribute{MarkdownDescription: "VLAN id ranges used for L2CIRCUIT resource type.", Required: true}}}}}}}
}

// Invoke executes the action against the remote API.
func (r *UpdateTrafficFlowsGlobalSettingsAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UpdateTrafficFlowsGlobalSettingsActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *UpdateTrafficFlowsGlobalSettingsAction) invokeRemote(ctx context.Context, config *UpdateTrafficFlowsGlobalSettingsActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/traffic-flows/global/settings"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_traffic_flows_global_settings", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_traffic_flows_global_settings", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPatch, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_traffic_flows_global_settings", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_traffic_flows_global_settings", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_traffic_flows_global_settings", "Invalid request")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_traffic_flows_global_settings", "Not Authenticated")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_traffic_flows_global_settings", "Access Denied")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_traffic_flows_global_settings", "Entity Not Found")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_traffic_flows_global_settings", "Conflict")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_traffic_flows_global_settings", "Internal Server Error")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_traffic_flows_global_settings", "Service Unavailable")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_update_traffic_flows_global_settings", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_traffic_flows_global_settings", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *UpdateTrafficFlowsGlobalSettingsAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
