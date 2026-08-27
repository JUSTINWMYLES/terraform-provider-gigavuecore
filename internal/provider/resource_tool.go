package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	path "github.com/hashicorp/terraform-plugin-framework/path"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertions.
var (
	_ resource.Resource                = (*ToolResource)(nil)
	_ resource.ResourceWithImportState = (*ToolResource)(nil)
	_ resource.ResourceWithConfigure   = (*ToolResource)(nil)
)

// ToolResource is the generated Terraform managed resource implementation.
type ToolResource struct {
	client *client.Client
}

// ToolResourceModel describes the Terraform state and plan shape for ToolResource.
type ToolResourceModel struct {
	Alias                   types.String `tfsdk:"alias"`
	ClusterId               types.String `tfsdk:"cluster_id" json:"clusterId"`
	CombinedHeartBeatStatus types.String `tfsdk:"combined_heart_beat_status" json:"combinedHeartBeatStatus"`
	Comment                 types.String `tfsdk:"comment"`
	Enabled                 types.Bool   `tfsdk:"enabled"`
	FailoverAction          types.String `tfsdk:"failover_action" json:"failoverAction"`
	FlexStatus              types.String `tfsdk:"flex_status" json:"flexStatus"`
	FlexTrafficPath         types.String `tfsdk:"flex_traffic_path" json:"flexTrafficPath"`
	HealthState             types.String `tfsdk:"health_state" json:"healthState"`
	HealthStateReasons      types.List   `tfsdk:"health_state_reasons" json:"healthStateReasons"`
	Heartbeat               types.Object `tfsdk:"heartbeat"`
	InlineToolType          types.String `tfsdk:"inline_tool_type" json:"inlineToolType"`
	NegativeHeartbeat       types.Object `tfsdk:"negative_heartbeat" json:"negativeHeartbeat"`
	OperationalState        types.String `tfsdk:"operational_state" json:"operationalState"`
	PortA                   types.String `tfsdk:"port_a" json:"portA"`
	PortAStatus             types.String `tfsdk:"port_a_status" json:"portAStatus"`
	PortB                   types.String `tfsdk:"port_b" json:"portB"`
	PortBStatus             types.String `tfsdk:"port_b_status" json:"portBStatus"`
	RecoveryMode            types.String `tfsdk:"recovery_mode" json:"recoveryMode"`
	Shared                  types.Bool   `tfsdk:"shared"`
	Timestamp               types.String `tfsdk:"timestamp"`
}

// Metadata returns the resource type name.
func (r *ToolResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_tool"
}

// Schema returns the Terraform schema for this resource.
func (r *ToolResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Find Inline Tool by alias", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Inline Tool alias. Unique within a cluster", Required: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "combined_heart_beat_status": schema.StringAttribute{MarkdownDescription: "combined Heartbeat status", Optional: true, Computed: true}, "comment": schema.StringAttribute{Optional: true, Computed: true}, "enabled": schema.BoolAttribute{Optional: true, Computed: true}, "failover_action": schema.StringAttribute{Optional: true, Computed: true}, "flex_status": schema.StringAttribute{Optional: true, Computed: true}, "flex_traffic_path": schema.StringAttribute{Optional: true, Computed: true}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Optional: true, Computed: true}, "health_state_reasons": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Optional: true, Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Optional: true, Computed: true}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Optional: true, Computed: true}}}}, "heartbeat": schema.SingleNestedAttribute{MarkdownDescription: "Embedded Heartbeat configuration for an Inline Tool. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{Required: true}, "ip_address_a": schema.StringAttribute{MarkdownDescription: "the destination IP address to be used in heartbeat packets send from side A to side B (the default is N.N.N.N where N is the port number within the chassis as shown on the face plate)", Required: true}, "ip_address_b": schema.StringAttribute{MarkdownDescription: "the destination IP address to be used in heartbeat packets send from side B to side A (the default is N.N.N.N where N is the port number within the chassis as shown on the face plate)", Required: true}, "profile": schema.StringAttribute{MarkdownDescription: "Alias of referenced Heartbeat Profile. the default is the heartbeat profile named 'default'", Required: true}, "status": schema.SingleNestedAttribute{MarkdownDescription: "Embedded Heartbeat configuration for an Inline Tool. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"heartbeat_passing": schema.StringAttribute{MarkdownDescription: "indicates whether heartbeat packets from portA are reaching portB and vice versa", Optional: true, Computed: true}, "stats_ato_b": schema.SingleNestedAttribute{MarkdownDescription: "Embedded Heartbeat configuration for an Inline Tool. Private class", Required: true, Attributes: map[string]schema.Attribute{"rx_count": schema.Int64Attribute{MarkdownDescription: "number of heartbeat packets received", Required: true}, "tx_count": schema.Int64Attribute{MarkdownDescription: "number of heartbeat packets sent", Required: true}}}, "stats_bto_a": schema.SingleNestedAttribute{MarkdownDescription: "Embedded Heartbeat configuration for an Inline Tool. Private class", Required: true, Attributes: map[string]schema.Attribute{"rx_count": schema.Int64Attribute{MarkdownDescription: "number of heartbeat packets received", Required: true}, "tx_count": schema.Int64Attribute{MarkdownDescription: "number of heartbeat packets sent", Required: true}}}}}}}, "inline_tool_type": schema.StringAttribute{MarkdownDescription: "inlineTool type", Optional: true, Computed: true}, "negative_heartbeat": schema.SingleNestedAttribute{MarkdownDescription: "Embedded Heartbeat configuration for an Inline Tool. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{Required: true}, "profile": schema.StringAttribute{MarkdownDescription: "Alias of referenced Negative Heartbeat Profile", Optional: true, Computed: true}, "status": schema.SingleNestedAttribute{MarkdownDescription: "Embedded Heartbeat configuration for an Inline Tool. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"heartbeat_passing": schema.StringAttribute{MarkdownDescription: "indicates whether heartbeat packets from portA are reaching portB and vice versa", Optional: true, Computed: true}, "stats_ato_b": schema.SingleNestedAttribute{MarkdownDescription: "Embedded Heartbeat configuration for an Inline Tool. Private class", Required: true, Attributes: map[string]schema.Attribute{"rx_count": schema.Int64Attribute{MarkdownDescription: "number of heartbeat packets received", Required: true}, "tx_count": schema.Int64Attribute{MarkdownDescription: "number of heartbeat packets sent", Required: true}}}, "stats_bto_a": schema.SingleNestedAttribute{MarkdownDescription: "Embedded Heartbeat configuration for an Inline Tool. Private class", Required: true, Attributes: map[string]schema.Attribute{"rx_count": schema.Int64Attribute{MarkdownDescription: "number of heartbeat packets received", Required: true}, "tx_count": schema.Int64Attribute{MarkdownDescription: "number of heartbeat packets sent", Required: true}}}}}}}, "operational_state": schema.StringAttribute{MarkdownDescription: "Operational State", Optional: true, Computed: true}, "port_a": schema.StringAttribute{MarkdownDescription: "portId of side A inline tool port", Required: true}, "port_a_status": schema.StringAttribute{MarkdownDescription: "port A status", Optional: true, Computed: true}, "port_b": schema.StringAttribute{MarkdownDescription: "portId of side B inline tool port", Required: true}, "port_b_status": schema.StringAttribute{MarkdownDescription: "port B status", Optional: true, Computed: true}, "recovery_mode": schema.StringAttribute{Optional: true, Computed: true}, "shared": schema.BoolAttribute{MarkdownDescription: "inline tool sharing mode", Optional: true, Computed: true}, "timestamp": schema.StringAttribute{MarkdownDescription: "date-time of stats collection in RFC 3339 format", Optional: true, Computed: true}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *ToolResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ToolResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.createRemote(ctx, &plan, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// createRemote performs the create HTTP exchange and decodes the response into plan. Extracted from Create so the request/response logic is unit-testable without a tfsdk.Plan.
func (r *ToolResource) createRemote(ctx context.Context, plan *ToolResourceModel, resp *resource.CreateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_tool", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_tool", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/inline/tools"
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_tool", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", plan.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_tool", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error creating gigavuecore_tool", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error creating gigavuecore_tool", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error creating gigavuecore_tool", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error creating gigavuecore_tool", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error creating gigavuecore_tool", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error creating gigavuecore_tool", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error creating gigavuecore_tool", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error creating gigavuecore_tool", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error creating gigavuecore_tool", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error creating gigavuecore_tool", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_tool", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
	if plan.Alias.IsNull() || plan.Alias.IsUnknown() {
		loc := httpResp.Header.Get("Location")
		if loc != "" {
			plan.Alias = types.StringValue(loc)
		} else {
			resp.Diagnostics.AddError("Error creating gigavuecore_tool", "The create response did not contain an identifier and no Location header was returned, so the resource cannot be tracked in state.")
			return
		}
	}
}

// Read refreshes the Terraform state with the latest remote values.
func (r *ToolResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ToolResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if r.readRemote(ctx, &state, resp) {
		resp.State.RemoveResource(ctx)
		return
	}
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// readRemote performs the read HTTP exchange and decodes the response into state, returning removed=true when the API reports 404. Extracted from Read so the request/response logic is unit-testable without a tfsdk.State.
func (r *ToolResource) readRemote(ctx context.Context, state *ToolResourceModel, resp *resource.ReadResponse) (removed bool) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/inline/tools/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(state.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_tool", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", state.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_tool", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		removed = true
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_tool", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_tool", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_tool", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_tool", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_tool", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_tool", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_tool", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_tool", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_tool", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_tool", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["inlineTool"]; ok {
		if m, ok := v.(map[string]any); ok {
			data = m
		} else if arr, ok := v.([]any); ok && len(arr) > 0 {
			if m, ok := arr[0].(map[string]any); ok {
				data = m
			}
		}
	}
	err = applyJSONToModel(&state, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_tool", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
	return
}

// Update modifies the remote resource to match the desired plan.
func (r *ToolResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ToolResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state ToolResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if plan.Alias.IsNull() || plan.Alias.IsUnknown() {
		if !state.Alias.IsNull() && !state.Alias.IsUnknown() {
			plan.Alias = state.Alias
		}
	}
	preserveStateIntoPlan(&plan, &state)
	r.updateRemote(ctx, &plan, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// updateRemote performs the update HTTP exchange and decodes the response into plan. Extracted from Update so the request/response logic is unit-testable without a tfsdk.Plan.
func (r *ToolResource) updateRemote(ctx context.Context, plan *ToolResourceModel, resp *resource.UpdateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_tool", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_tool", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/inline/tools/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(plan.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_tool", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", plan.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_tool", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error updating gigavuecore_tool", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error updating gigavuecore_tool", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error updating gigavuecore_tool", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error updating gigavuecore_tool", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error updating gigavuecore_tool", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error updating gigavuecore_tool", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error updating gigavuecore_tool", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error updating gigavuecore_tool", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error updating gigavuecore_tool", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error updating gigavuecore_tool", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_tool", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Delete destroys the remote resource.
func (r *ToolResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ToolResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.deleteRemote(ctx, &state, resp)
}

// deleteRemote performs the delete HTTP exchange, treating a 404 as already deleted. Extracted from Delete so the request/response logic is unit-testable without a tfsdk.State.
func (r *ToolResource) deleteRemote(ctx context.Context, state *ToolResourceModel, resp *resource.DeleteResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/inline/tools/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(state.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_tool", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", state.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_tool", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		return
	}
	if !(httpResp.StatusCode == 204) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error deleting gigavuecore_tool", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error deleting gigavuecore_tool", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error deleting gigavuecore_tool", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error deleting gigavuecore_tool", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error deleting gigavuecore_tool", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error deleting gigavuecore_tool", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error deleting gigavuecore_tool", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error deleting gigavuecore_tool", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error deleting gigavuecore_tool", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *ToolResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData))
		return
	}
	r.client = c
}

// ImportState imports an existing remote resource into Terraform state.
func (r *ToolResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("alias"), req.ID)...)
}
