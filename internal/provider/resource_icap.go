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
	"time"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	timeouts "github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	int64validator "github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	stringvalidator "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	path "github.com/hashicorp/terraform-plugin-framework/path"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	identityschema "github.com/hashicorp/terraform-plugin-framework/resource/identityschema"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	validator "github.com/hashicorp/terraform-plugin-framework/schema/validator"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertions.
var (
	_ resource.Resource                = (*IcapResource)(nil)
	_ resource.ResourceWithIdentity    = (*IcapResource)(nil)
	_ resource.ResourceWithImportState = (*IcapResource)(nil)
	_ resource.ResourceWithConfigure   = (*IcapResource)(nil)
)

// IcapResource is the generated Terraform managed resource implementation.
type IcapResource struct {
	client *client.Client
}

// IcapResourceModel describes the Terraform state and plan shape for IcapResource.
type IcapResourceModel struct {
	Alias               types.String   `tfsdk:"alias"`
	ClusterId           types.String   `tfsdk:"cluster_id" json:"clusterId"`
	ConfigStatus        types.String   `tfsdk:"config_status" json:"configStatus"`
	ConfigStatusReasons types.List     `tfsdk:"config_status_reasons" json:"configStatusReasons"`
	GsEngines           types.List     `tfsdk:"gs_engines" json:"gsEngines"`
	GsGrpAlias          types.String   `tfsdk:"gs_grp_alias" json:"gsGrpAlias"`
	GsopAlias           types.String   `tfsdk:"gsop_alias" json:"gsopAlias"`
	HealthState         types.String   `tfsdk:"health_state" json:"healthState"`
	HealthStateReasons  types.List     `tfsdk:"health_state_reasons" json:"healthStateReasons"`
	IcapMapAlias        types.String   `tfsdk:"icap_map_alias" json:"icapMapAlias"`
	IcapProfileConfig   types.Object   `tfsdk:"icap_profile_config" json:"icapProfileConfig"`
	IcapServerGrpAlias  types.String   `tfsdk:"icap_server_grp_alias" json:"icapServerGrpAlias"`
	IcapServers         types.List     `tfsdk:"icap_servers" json:"icapServers"`
	IngAlias            types.String   `tfsdk:"ing_alias" json:"ingAlias"`
	InlineNetworks      types.List     `tfsdk:"inline_networks" json:"inlineNetworks"`
	IpInterfaceAlias    types.String   `tfsdk:"ip_interface_alias" json:"ipInterfaceAlias"`
	Timeouts            timeouts.Value `tfsdk:"timeouts"`
}

// Metadata returns the resource type name.
func (r *IcapResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_icap"
}

// Schema returns the Terraform schema for this resource.
func (r *IcapResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Create an icap client", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Icap Alias", Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "cluster_id": schema.StringAttribute{MarkdownDescription: "id of the defining cluster", Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "config_status": schema.StringAttribute{MarkdownDescription: "Configuration status of this ICAP Client app", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("SUCCESS", "PARTIAL_SUCCESS", "FAILED", "PENDING", "INCOMPLETE_DELETION", "DAMAGED")}}, "config_status_reasons": schema.ListAttribute{MarkdownDescription: "In case of configuration failure, this message provides details about the possible cause of the failure", Optional: true, Computed: true, ElementType: types.StringType}, "gs_engines": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType}, "gs_grp_alias": schema.StringAttribute{MarkdownDescription: "GSGroup Alias // User provided (or) Auto-generated", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "gsop_alias": schema.StringAttribute{MarkdownDescription: "GSOP Alias", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("green", "yellow", "orange", "red")}}, "health_state_reasons": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Optional: true, Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("green", "yellow", "orange", "red")}}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("PORT_LOW_UTIL", "PORT_HIGH_UTIL", "PORT_PACKET_DROP", "PORT_PACKET_ERROR", "GIGASTREAM_IMBALANCE")}}}}}, "icap_map_alias": schema.StringAttribute{MarkdownDescription: "Icap MAP Alias", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "icap_profile_config": schema.SingleNestedAttribute{MarkdownDescription: "ICAP Profile", Required: true, Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Icap Alias", Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "cluster_id": schema.StringAttribute{MarkdownDescription: "id of the defining cluster", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "exceed_action": schema.StringAttribute{MarkdownDescription: "Icap Profile action incase of Http request buffer exceeded", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("drop", "bypass")}}, "http_req_buf": schema.Int64Attribute{MarkdownDescription: "Icap Profile Http request buffer in KB", Required: true, Validators: []validator.Int64{int64validator.Between(10, 10000)}}, "inactivity_timeout": schema.Int64Attribute{MarkdownDescription: "Icap Inactivity timeout in minutes", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(2, 1440)}}, "preview": schema.Int64Attribute{MarkdownDescription: "Icap Preview bytes in KB", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 50)}}, "resp_mod": schema.StringAttribute{MarkdownDescription: "Icap Response Modification Enable|Disable", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("enable", "disable")}}, "resp_timeout": schema.Int64Attribute{MarkdownDescription: "Icap Server Response Timeout value in seconds", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(5, 60)}}, "resp_timeout_action": schema.StringAttribute{MarkdownDescription: "Response Timeout action Drop|Bypass", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("drop", "bypass")}}, "server_group": schema.StringAttribute{MarkdownDescription: "Icap Server Group Alias", Required: true}, "src_max_l4_port": schema.Int64Attribute{MarkdownDescription: "Icap Service Source l4 port maximum", Required: true, Validators: []validator.Int64{int64validator.Between(10000, 60000)}}, "src_min_l4_port": schema.Int64Attribute{MarkdownDescription: "Icap Service Source l4 port minimum", Required: true, Validators: []validator.Int64{int64validator.Between(10000, 60000)}}}}, "icap_server_grp_alias": schema.StringAttribute{MarkdownDescription: "Icap Server Group Alias // Auto-generated", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "icap_servers": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Icap Server Alias", Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "cluster_id": schema.StringAttribute{MarkdownDescription: "id of the defining cluster", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "comment": schema.StringAttribute{MarkdownDescription: "Icap Server Comment", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "l3_address": schema.StringAttribute{MarkdownDescription: "Icap Server IP Address", Required: true}, "l4_port": schema.Int64Attribute{MarkdownDescription: "Icap Server l4 Port", Required: true}, "options_service_url": schema.StringAttribute{MarkdownDescription: "Options Service URL", Optional: true, Computed: true}, "reqmod_service_url": schema.StringAttribute{MarkdownDescription: "Request Modification Service URL", Optional: true, Computed: true}, "respmod_service_url": schema.StringAttribute{MarkdownDescription: "Response Modification Service URL", Optional: true, Computed: true}}}}, "ing_alias": schema.StringAttribute{MarkdownDescription: "Inline network group Alias //Auto-generated", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "inline_networks": schema.ListAttribute{Required: true, ElementType: types.StringType}, "ip_interface_alias": schema.StringAttribute{MarkdownDescription: "IpInterface Alias", Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}}, Blocks: map[string]schema.Block{"timeouts": timeouts.Block(ctx, timeouts.Opts{Create: true, Read: true, Update: true, Delete: true})}}
}

// IdentitySchema returns the resource identity schema shared with the paired list resource.
func (r *IcapResource) IdentitySchema(_ context.Context, _ resource.IdentitySchemaRequest, resp *resource.IdentitySchemaResponse) {
	resp.IdentitySchema = identityschema.Schema{Attributes: map[string]identityschema.Attribute{"alias": identityschema.StringAttribute{RequiredForImport: true, Description: "Icap Alias"}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *IcapResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IcapResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	timeout, diags := plan.Timeouts.Create(ctx, 20*time.Minute)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	r.createRemote(ctx, &plan, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("alias"), plan.Alias)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// createRemote performs the create HTTP exchange and decodes the response into plan. Extracted from Create so the request/response logic is unit-testable without a tfsdk.Plan.
func (r *IcapResource) createRemote(ctx context.Context, plan *IcapResourceModel, resp *resource.CreateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_icap", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_icap", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/icap"
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_icap", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_icap", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error creating gigavuecore_icap", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error creating gigavuecore_icap", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error creating gigavuecore_icap", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error creating gigavuecore_icap", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error creating gigavuecore_icap", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error creating gigavuecore_icap", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error creating gigavuecore_icap", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error creating gigavuecore_icap", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error creating gigavuecore_icap", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error creating gigavuecore_icap", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_icap", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
	if plan.Alias.IsNull() || plan.Alias.IsUnknown() {
		loc := httpResp.Header.Get("Location")
		if loc != "" {
			loc = strings.TrimRight(loc, "/")
			i := strings.LastIndex(loc, "/")
			if i >= 0 {
				loc = loc[i+1:]
			}
			plan.Alias = types.StringValue(loc)
		} else {
			resp.Diagnostics.AddError("Error creating gigavuecore_icap", "The create response did not contain an identifier and no Location header was returned, so the resource cannot be tracked in state.")
			return
		}
	}
}

// Read refreshes the Terraform state with the latest remote values.
func (r *IcapResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IcapResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	timeout, diags := state.Timeouts.Read(ctx, 10*time.Minute)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	if r.readRemote(ctx, &state, resp) {
		resp.State.RemoveResource(ctx)
		return
	}
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("alias"), state.Alias)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// readRemote performs the read HTTP exchange and decodes the response into state, returning removed=true when the API reports 404. Extracted from Read so the request/response logic is unit-testable without a tfsdk.State.
func (r *IcapResource) readRemote(ctx context.Context, state *IcapResourceModel, resp *resource.ReadResponse) (removed bool) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/icap/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(state.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_icap", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !state.ClusterId.IsNull() {
		query.Set("clusterId", state.ClusterId.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_icap", fmt.Sprintf("Could not send request: %s", err))
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
			resp.Diagnostics.AddError("Error reading gigavuecore_icap", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_icap", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_icap", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_icap", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_icap", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_icap", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_icap", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_icap", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_icap", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_icap", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["gigaIcapClient"]; ok {
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
		resp.Diagnostics.AddError("Error reading gigavuecore_icap", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
	return
}

// Update modifies the remote resource to match the desired plan.
func (r *IcapResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan IcapResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state IcapResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	timeout, diags := plan.Timeouts.Update(ctx, 20*time.Minute)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
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
	resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("alias"), plan.Alias)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// updateRemote performs the update HTTP exchange and decodes the response into plan. Extracted from Update so the request/response logic is unit-testable without a tfsdk.Plan.
func (r *IcapResource) updateRemote(ctx context.Context, plan *IcapResourceModel, resp *resource.UpdateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_icap", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_icap", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/icap/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(plan.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_icap", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !plan.ClusterId.IsNull() {
		query.Set("clusterId", plan.ClusterId.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_icap", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error updating gigavuecore_icap", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error updating gigavuecore_icap", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error updating gigavuecore_icap", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error updating gigavuecore_icap", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error updating gigavuecore_icap", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error updating gigavuecore_icap", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error updating gigavuecore_icap", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error updating gigavuecore_icap", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error updating gigavuecore_icap", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error updating gigavuecore_icap", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_icap", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Delete destroys the remote resource.
func (r *IcapResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IcapResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	timeout, diags := state.Timeouts.Delete(ctx, 10*time.Minute)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	r.deleteRemote(ctx, &state, resp)
}

// deleteRemote performs the delete HTTP exchange, treating a 404 as already deleted. Extracted from Delete so the request/response logic is unit-testable without a tfsdk.State.
func (r *IcapResource) deleteRemote(ctx context.Context, state *IcapResourceModel, resp *resource.DeleteResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/icap/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(state.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_icap", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !state.ClusterId.IsNull() {
		query.Set("clusterId", state.ClusterId.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_icap", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		return
	}
	if !(httpResp.StatusCode == 204) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error deleting gigavuecore_icap", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error deleting gigavuecore_icap", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error deleting gigavuecore_icap", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error deleting gigavuecore_icap", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error deleting gigavuecore_icap", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error deleting gigavuecore_icap", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error deleting gigavuecore_icap", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error deleting gigavuecore_icap", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error deleting gigavuecore_icap", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *IcapResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *IcapResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("alias"), req.ID)...)
}
