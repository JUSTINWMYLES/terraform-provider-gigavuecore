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
var _ action.Action = (*UpdateBatteryOptimizationAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*UpdateBatteryOptimizationAction)(nil)

// UpdateBatteryOptimizationAction is the generated Terraform action implementation.
type UpdateBatteryOptimizationAction struct {
	client *client.Client
}

// UpdateBatteryOptimizationActionModel describes the action configuration shape.
type UpdateBatteryOptimizationActionModel struct {
	ClusterId      types.String `tfsdk:"cluster_id"`
	CpuHibernation types.Object `tfsdk:"cpu_hibernation" json:"cpuHibernation"`
	MonitorPort    types.List   `tfsdk:"monitor_port" json:"monitorPort"`
	UnusedPort     types.List   `tfsdk:"unused_port" json:"unusedPort"`
}

// NewUpdateBatteryOptimizationAction returns a new instance of the generated action.
func NewUpdateBatteryOptimizationAction() action.Action {
	return &UpdateBatteryOptimizationAction{}
}

// Metadata returns the action type name.
func (r *UpdateBatteryOptimizationAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_update_battery_optimization"
}

// Schema returns the action schema.
func (r *UpdateBatteryOptimizationAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "update battery optimization", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "cpu_hibernation": schema.SingleNestedAttribute{MarkdownDescription: "CPU hibernation battery optimization", Optional: true, Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{MarkdownDescription: "Enable/disable cpu hibernation battery optimization", Optional: true}, "sleep_in_mins": schema.Int64Attribute{MarkdownDescription: "Apply cpu hibernation battery optimization for this configured sleep time", Optional: true}, "threshold_level": schema.Int64Attribute{MarkdownDescription: "Apply cpu hibernation battery optimization at this battery level", Optional: true}}}, "monitor_port": schema.ListNestedAttribute{MarkdownDescription: "Monitor port off battery optimization", Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{MarkdownDescription: "Enable/disable port battery optimization", Optional: true}, "level": schema.StringAttribute{MarkdownDescription: "Apply port battery optimization at this battery level", Optional: true}, "port_group_id": schema.StringAttribute{MarkdownDescription: "Port group Id", Required: true}}}}, "unused_port": schema.ListNestedAttribute{MarkdownDescription: "Unused port off battery optimization", Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{MarkdownDescription: "Enable/disable port battery optimization", Optional: true}, "level": schema.StringAttribute{MarkdownDescription: "Apply port battery optimization at this battery level", Optional: true}, "port_group_id": schema.StringAttribute{MarkdownDescription: "Port group Id", Required: true}}}}}}
}

// Invoke executes the action against the remote API.
func (r *UpdateBatteryOptimizationAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UpdateBatteryOptimizationActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *UpdateBatteryOptimizationAction) invokeRemote(ctx context.Context, config *UpdateBatteryOptimizationActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/gtap/battery/optimization"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_battery_optimization", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_battery_optimization", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPatch, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_battery_optimization", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_battery_optimization", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_battery_optimization", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_battery_optimization", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_battery_optimization", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_battery_optimization", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_battery_optimization", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_battery_optimization", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_battery_optimization", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_update_battery_optimization", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_battery_optimization", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *UpdateBatteryOptimizationAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
