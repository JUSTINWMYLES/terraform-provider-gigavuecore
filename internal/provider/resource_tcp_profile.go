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
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	timeouts "github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	path "github.com/hashicorp/terraform-plugin-framework/path"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertions.
var (
	_ resource.Resource                = (*TcpProfileResource)(nil)
	_ resource.ResourceWithImportState = (*TcpProfileResource)(nil)
	_ resource.ResourceWithConfigure   = (*TcpProfileResource)(nil)
)

// TcpProfileResource is the generated Terraform managed resource implementation.
type TcpProfileResource struct {
	client *client.Client
}

// TcpProfileResourceModel describes the Terraform state and plan shape for TcpProfileResource.
type TcpProfileResourceModel struct {
	Alias          types.String   `tfsdk:"alias"`
	ClusterId      types.String   `tfsdk:"cluster_id" json:"clusterId"`
	KeepAliveTimer types.Int64    `tfsdk:"keep_alive_timer" json:"keepAliveTimer"`
	SelectiveAck   types.String   `tfsdk:"selective_ack" json:"selectiveAck"`
	SynRetries     types.Int64    `tfsdk:"syn_retries" json:"synRetries"`
	Timeouts       timeouts.Value `tfsdk:"timeouts"`
}

// Metadata returns the resource type name.
func (r *TcpProfileResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_tcp_profile"
}

// Schema returns the Terraform schema for this resource.
func (r *TcpProfileResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get Apps TCP Profile for given alias", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{Required: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "keep_alive_timer": schema.Int64Attribute{Optional: true, Computed: true}, "selective_ack": schema.StringAttribute{Optional: true, Computed: true}, "syn_retries": schema.Int64Attribute{Optional: true, Computed: true}}, Blocks: map[string]schema.Block{"timeouts": timeouts.Block(ctx, timeouts.Opts{Create: true, Read: true, Update: true, Delete: true})}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *TcpProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan TcpProfileResourceModel
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
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// createRemote performs the create HTTP exchange and decodes the response into plan. Extracted from Create so the request/response logic is unit-testable without a tfsdk.Plan.
func (r *TcpProfileResource) createRemote(ctx context.Context, plan *TcpProfileResourceModel, resp *resource.CreateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_tcp_profile", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_tcp_profile", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/apps/tcp/profile"
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_tcp_profile", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", plan.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_tcp_profile", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error creating gigavuecore_tcp_profile", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error creating gigavuecore_tcp_profile", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error creating gigavuecore_tcp_profile", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error creating gigavuecore_tcp_profile", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error creating gigavuecore_tcp_profile", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error creating gigavuecore_tcp_profile", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error creating gigavuecore_tcp_profile", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error creating gigavuecore_tcp_profile", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error creating gigavuecore_tcp_profile", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error creating gigavuecore_tcp_profile", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_tcp_profile", fmt.Sprintf("Could not map response to state: %s", err))
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
			resp.Diagnostics.AddError("Error creating gigavuecore_tcp_profile", "The create response did not contain an identifier and no Location header was returned, so the resource cannot be tracked in state.")
			return
		}
	}
}

// Read refreshes the Terraform state with the latest remote values.
func (r *TcpProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state TcpProfileResourceModel
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
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// readRemote performs the read HTTP exchange and decodes the response into state, returning removed=true when the API reports 404. Extracted from Read so the request/response logic is unit-testable without a tfsdk.State.
func (r *TcpProfileResource) readRemote(ctx context.Context, state *TcpProfileResourceModel, resp *resource.ReadResponse) (removed bool) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/tcp/profile/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(state.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_tcp_profile", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", state.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_tcp_profile", fmt.Sprintf("Could not send request: %s", err))
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
			resp.Diagnostics.AddError("Error reading gigavuecore_tcp_profile", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_tcp_profile", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_tcp_profile", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_tcp_profile", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_tcp_profile", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_tcp_profile", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_tcp_profile", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_tcp_profile", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_tcp_profile", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_tcp_profile", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["appsTcpProfile"]; ok {
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
		resp.Diagnostics.AddError("Error reading gigavuecore_tcp_profile", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
	return
}

// Update modifies the remote resource to match the desired plan.
func (r *TcpProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan TcpProfileResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state TcpProfileResourceModel
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
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// updateRemote performs the update HTTP exchange and decodes the response into plan. Extracted from Update so the request/response logic is unit-testable without a tfsdk.Plan.
func (r *TcpProfileResource) updateRemote(ctx context.Context, plan *TcpProfileResourceModel, resp *resource.UpdateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_tcp_profile", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_tcp_profile", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/apps/tcp/profile/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(plan.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodPatch, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_tcp_profile", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", plan.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_tcp_profile", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error updating gigavuecore_tcp_profile", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error updating gigavuecore_tcp_profile", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error updating gigavuecore_tcp_profile", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error updating gigavuecore_tcp_profile", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error updating gigavuecore_tcp_profile", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error updating gigavuecore_tcp_profile", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error updating gigavuecore_tcp_profile", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error updating gigavuecore_tcp_profile", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error updating gigavuecore_tcp_profile", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error updating gigavuecore_tcp_profile", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_tcp_profile", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Delete destroys the remote resource.
func (r *TcpProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state TcpProfileResourceModel
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
func (r *TcpProfileResource) deleteRemote(ctx context.Context, state *TcpProfileResourceModel, resp *resource.DeleteResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/tcp/profile/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(state.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_tcp_profile", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", state.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_tcp_profile", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		return
	}
	if !(httpResp.StatusCode == 204) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error deleting gigavuecore_tcp_profile", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error deleting gigavuecore_tcp_profile", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error deleting gigavuecore_tcp_profile", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error deleting gigavuecore_tcp_profile", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error deleting gigavuecore_tcp_profile", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error deleting gigavuecore_tcp_profile", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error deleting gigavuecore_tcp_profile", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error deleting gigavuecore_tcp_profile", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error deleting gigavuecore_tcp_profile", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *TcpProfileResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *TcpProfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("alias"), req.ID)...)
}
