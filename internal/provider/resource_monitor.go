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
	int64validator "github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	listvalidator "github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	setvalidator "github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
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
	_ resource.Resource                = (*MonitorResource)(nil)
	_ resource.ResourceWithIdentity    = (*MonitorResource)(nil)
	_ resource.ResourceWithImportState = (*MonitorResource)(nil)
	_ resource.ResourceWithConfigure   = (*MonitorResource)(nil)
)

// MonitorResource is the generated Terraform managed resource implementation.
type MonitorResource struct {
	client *client.Client
}

// MonitorResourceModel describes the Terraform state and plan shape for MonitorResource.
type MonitorResourceModel struct {
	Alias               types.String          `tfsdk:"alias"`
	Cache               types.Object          `tfsdk:"cache"`
	ClusterId           types.String          `tfsdk:"cluster_id" json:"clusterId"`
	Description         types.String          `tfsdk:"description"`
	Records             types.Set             `tfsdk:"records"`
	Sampling            types.Object          `tfsdk:"sampling"`
	SamplingSpace       types.Int64           `tfsdk:"sampling_space" json:"samplingSpace"`
	SslPortRestrictions types.Object          `tfsdk:"ssl_port_restrictions" json:"sslPortRestrictions"`
	Timeouts            *MonitorTimeoutsModel `tfsdk:"timeouts"`
}

// MonitorTimeoutsModel describes the per-operation timeout configuration (in seconds) for the gigavuecore_monitor resource.
type MonitorTimeoutsModel struct {
	Create types.Int64 `tfsdk:"create"`
	Read   types.Int64 `tfsdk:"read"`
	Update types.Int64 `tfsdk:"update"`
	Delete types.Int64 `tfsdk:"delete"`
}

// Metadata returns the resource type name.
func (r *MonitorResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_monitor"
}

// Schema returns the Terraform schema for this resource.
func (r *MonitorResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Create a new Netflow Monitor", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{Required: true}, "cache": schema.SingleNestedAttribute{MarkdownDescription: "Netflow Monitor Cache", Required: true, Attributes: map[string]schema.Attribute{"export_triggers": schema.SingleNestedAttribute{MarkdownDescription: "Netflow Monitor Cache Export Triggers", Required: true, Attributes: map[string]schema.Attribute{"event": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("txnEnd", "none")}}, "timeout_active": schema.Int64Attribute{MarkdownDescription: "in seconds. max value is 7 days. default is 30 min", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 604800)}}, "timeout_inactive": schema.Int64Attribute{MarkdownDescription: "in seconds. max value is 7 days. default is 15 sec", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 604800)}}}}, "type": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("normal")}}}}, "cluster_id": schema.StringAttribute{MarkdownDescription: "id of the defining cluster", Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "description": schema.StringAttribute{Optional: true, Computed: true}, "records": schema.SetAttribute{MarkdownDescription: "Aliases of referenced Netflow Records defined on the GsGroup. Up to 5 records.", Optional: true, Computed: true, ElementType: types.StringType, Validators: []validator.Set{setvalidator.SizeAtMost(5)}}, "sampling": schema.SingleNestedAttribute{MarkdownDescription: "monitor sampling", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"mode": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("multi-rate", "no-sampling", "single-rate")}}, "single_sampling_rate": schema.Int64Attribute{MarkdownDescription: "Packet interval window size. Valid values: 10-16000 (in packets)", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(10, 16000)}}}}, "sampling_space": schema.Int64Attribute{MarkdownDescription: "DEPRECATED: use 'sampling'", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 16000)}}, "ssl_port_restrictions": schema.SingleNestedAttribute{MarkdownDescription: "Port restrictions for Netflow/SSL sessions", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"ports": schema.ListAttribute{MarkdownDescription: "The list of TCP ports whose packets will be sent to the SSL module. Not valid if sslPort is 'all'. if 'wellKnownPorts' is selected for 'sslPorts' ports will be [993, 995, 465, 636, 563, 443]", Optional: true, Computed: true, ElementType: types.Int64Type, Validators: []validator.List{listvalidator.SizeAtMost(10)}}, "ssl_ports": schema.StringAttribute{MarkdownDescription: "Ports whose packets will be sent to the SSL module. 'all' ports; 'wellKnownPorts' is [993, 995, 465, 636, 563, 443]; or 'ports' - specify upto 10 ports", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("all", "wellKnownPorts", "ports")}}}}}, Blocks: map[string]schema.Block{"timeouts": schema.SingleNestedBlock{Attributes: map[string]schema.Attribute{"create": schema.Int64Attribute{MarkdownDescription: "Maximum time in seconds for the create operation. Defaults to 1200.", Optional: true}, "read": schema.Int64Attribute{MarkdownDescription: "Maximum time in seconds for the read operation. Defaults to 1200.", Optional: true}, "update": schema.Int64Attribute{MarkdownDescription: "Maximum time in seconds for the update operation. Defaults to 1200.", Optional: true}, "delete": schema.Int64Attribute{MarkdownDescription: "Maximum time in seconds for the delete operation. Defaults to 1200.", Optional: true}}, MarkdownDescription: "Per-operation timeouts in seconds."}}}
}

// IdentitySchema returns the resource identity schema shared with the paired list resource.
func (r *MonitorResource) IdentitySchema(_ context.Context, _ resource.IdentitySchemaRequest, resp *resource.IdentitySchemaResponse) {
	resp.IdentitySchema = identityschema.Schema{Attributes: map[string]identityschema.Attribute{"alias": identityschema.StringAttribute{RequiredForImport: true}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *MonitorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan MonitorResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	timeout := 20 * time.Minute
	if plan.Timeouts != nil && !plan.Timeouts.Create.IsNull() && !plan.Timeouts.Create.IsUnknown() {
		timeout = time.Duration(plan.Timeouts.Create.ValueInt64()) * time.Second
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
func (r *MonitorResource) createRemote(ctx context.Context, plan *MonitorResourceModel, resp *resource.CreateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_monitor", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_monitor", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/apps/netflow/monitors"
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_monitor", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", plan.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_monitor", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error creating gigavuecore_monitor", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error creating gigavuecore_monitor", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error creating gigavuecore_monitor", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error creating gigavuecore_monitor", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error creating gigavuecore_monitor", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error creating gigavuecore_monitor", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error creating gigavuecore_monitor", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error creating gigavuecore_monitor", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error creating gigavuecore_monitor", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error creating gigavuecore_monitor", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_monitor", fmt.Sprintf("Could not map response to state: %s", err))
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
			resp.Diagnostics.AddError("Error creating gigavuecore_monitor", "The create response did not contain an identifier and no Location header was returned, so the resource cannot be tracked in state.")
			return
		}
	}
}

// Read refreshes the Terraform state with the latest remote values.
func (r *MonitorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state MonitorResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	timeout := 20 * time.Minute
	if state.Timeouts != nil && !state.Timeouts.Read.IsNull() && !state.Timeouts.Read.IsUnknown() {
		timeout = time.Duration(state.Timeouts.Read.ValueInt64()) * time.Second
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
func (r *MonitorResource) readRemote(ctx context.Context, state *MonitorResourceModel, resp *resource.ReadResponse) (removed bool) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/netflow/monitors/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(state.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_monitor", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", state.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_monitor", fmt.Sprintf("Could not send request: %s", err))
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
			resp.Diagnostics.AddError("Error reading gigavuecore_monitor", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_monitor", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_monitor", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_monitor", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_monitor", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_monitor", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_monitor", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_monitor", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_monitor", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_monitor", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["nfMonitor"]; ok {
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
		resp.Diagnostics.AddError("Error reading gigavuecore_monitor", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
	return
}

// Update modifies the remote resource to match the desired plan.
func (r *MonitorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan MonitorResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state MonitorResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	timeout := 20 * time.Minute
	if plan.Timeouts != nil && !plan.Timeouts.Update.IsNull() && !plan.Timeouts.Update.IsUnknown() {
		timeout = time.Duration(plan.Timeouts.Update.ValueInt64()) * time.Second
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
func (r *MonitorResource) updateRemote(ctx context.Context, plan *MonitorResourceModel, resp *resource.UpdateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_monitor", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_monitor", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/apps/netflow/monitors/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(plan.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_monitor", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", plan.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_monitor", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error updating gigavuecore_monitor", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error updating gigavuecore_monitor", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error updating gigavuecore_monitor", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error updating gigavuecore_monitor", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error updating gigavuecore_monitor", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error updating gigavuecore_monitor", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error updating gigavuecore_monitor", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error updating gigavuecore_monitor", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error updating gigavuecore_monitor", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error updating gigavuecore_monitor", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_monitor", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Delete destroys the remote resource.
func (r *MonitorResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state MonitorResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	timeout := 20 * time.Minute
	if state.Timeouts != nil && !state.Timeouts.Delete.IsNull() && !state.Timeouts.Delete.IsUnknown() {
		timeout = time.Duration(state.Timeouts.Delete.ValueInt64()) * time.Second
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	r.deleteRemote(ctx, &state, resp)
}

// deleteRemote performs the delete HTTP exchange, treating a 404 as already deleted. Extracted from Delete so the request/response logic is unit-testable without a tfsdk.State.
func (r *MonitorResource) deleteRemote(ctx context.Context, state *MonitorResourceModel, resp *resource.DeleteResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/netflow/monitors/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(state.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_monitor", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", state.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_monitor", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		return
	}
	if !(httpResp.StatusCode == 204) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error deleting gigavuecore_monitor", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error deleting gigavuecore_monitor", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error deleting gigavuecore_monitor", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error deleting gigavuecore_monitor", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error deleting gigavuecore_monitor", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error deleting gigavuecore_monitor", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error deleting gigavuecore_monitor", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error deleting gigavuecore_monitor", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error deleting gigavuecore_monitor", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *MonitorResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *MonitorResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	monitorImportIDParts := strings.Split(req.ID, "/")
	if len(monitorImportIDParts) != 2 {
		resp.Diagnostics.AddError("Unexpected Import Identifier", fmt.Sprintf("Expected import identifier with format \"{alias}/{cluster_id}\". Got %q.", req.ID))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("alias"), monitorImportIDParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("cluster_id"), monitorImportIDParts[1])...)
}
