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
	_ resource.Resource                = (*DiameterWhitelistResource)(nil)
	_ resource.ResourceWithImportState = (*DiameterWhitelistResource)(nil)
	_ resource.ResourceWithConfigure   = (*DiameterWhitelistResource)(nil)
)

// DiameterWhitelistResource is the generated Terraform managed resource implementation.
type DiameterWhitelistResource struct {
	client *client.Client
}

// DiameterWhitelistResourceModel describes the Terraform state and plan shape for DiameterWhitelistResource.
type DiameterWhitelistResourceModel struct {
	Alias         types.String   `tfsdk:"alias"`
	Page          types.String   `tfsdk:"page"`
	Sort          types.String   `tfsdk:"sort"`
	UserNameCount types.Int64    `tfsdk:"user_name_count" json:"userNameCount"`
	Timeouts      timeouts.Value `tfsdk:"timeouts"`
}

// Metadata returns the resource type name.
func (r *DiameterWhitelistResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_diameter_whitelist"
}

// Schema returns the Terraform schema for this resource.
func (r *DiameterWhitelistResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load available Diameter Whitelists", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{Required: true}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}, "user_name_count": schema.Int64Attribute{MarkdownDescription: "Number of username entries in diameter-whitelist", Computed: true}}, Blocks: map[string]schema.Block{"timeouts": timeouts.Block(ctx, timeouts.Opts{Create: true, Read: true, Update: true, Delete: true})}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *DiameterWhitelistResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan DiameterWhitelistResourceModel
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
func (r *DiameterWhitelistResource) createRemote(ctx context.Context, plan *DiameterWhitelistResourceModel, resp *resource.CreateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_diameter_whitelist", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_diameter_whitelist", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/apps/diameter/whitelists"
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_diameter_whitelist", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_diameter_whitelist", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error creating gigavuecore_diameter_whitelist", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error creating gigavuecore_diameter_whitelist", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error creating gigavuecore_diameter_whitelist", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error creating gigavuecore_diameter_whitelist", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error creating gigavuecore_diameter_whitelist", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error creating gigavuecore_diameter_whitelist", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error creating gigavuecore_diameter_whitelist", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error creating gigavuecore_diameter_whitelist", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error creating gigavuecore_diameter_whitelist", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error creating gigavuecore_diameter_whitelist", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_diameter_whitelist", fmt.Sprintf("Could not map response to state: %s", err))
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
			resp.Diagnostics.AddError("Error creating gigavuecore_diameter_whitelist", "The create response did not contain an identifier and no Location header was returned, so the resource cannot be tracked in state.")
			return
		}
	}
}

// Read refreshes the Terraform state with the latest remote values.
func (r *DiameterWhitelistResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state DiameterWhitelistResourceModel
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
func (r *DiameterWhitelistResource) readRemote(ctx context.Context, state *DiameterWhitelistResourceModel, resp *resource.ReadResponse) (removed bool) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/diameter/whitelists"
	httpReq, err := r.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_diameter_whitelist", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !state.Page.IsNull() {
		query.Set("page", state.Page.ValueString())
	}
	if !state.Sort.IsNull() {
		query.Set("sort", state.Sort.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_diameter_whitelist", fmt.Sprintf("Could not send request: %s", err))
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
			resp.Diagnostics.AddError("Error reading gigavuecore_diameter_whitelist", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_diameter_whitelist", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_diameter_whitelist", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_diameter_whitelist", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_diameter_whitelist", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_diameter_whitelist", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_diameter_whitelist", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_diameter_whitelist", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_diameter_whitelist", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_diameter_whitelist", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["diameterWhitelists"]; ok {
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
		resp.Diagnostics.AddError("Error reading gigavuecore_diameter_whitelist", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
	return
}

// Update modifies the remote resource to match the desired plan.
func (r *DiameterWhitelistResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan DiameterWhitelistResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state DiameterWhitelistResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if plan.Alias.IsNull() || plan.Alias.IsUnknown() {
		if !state.Alias.IsNull() && !state.Alias.IsUnknown() {
			plan.Alias = state.Alias
		}
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Update is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Delete destroys the remote resource.
func (r *DiameterWhitelistResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state DiameterWhitelistResourceModel
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
func (r *DiameterWhitelistResource) deleteRemote(ctx context.Context, state *DiameterWhitelistResourceModel, resp *resource.DeleteResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/diameter/whitelists/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(state.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_diameter_whitelist", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_diameter_whitelist", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		return
	}
	if !(httpResp.StatusCode == 204) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error deleting gigavuecore_diameter_whitelist", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error deleting gigavuecore_diameter_whitelist", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error deleting gigavuecore_diameter_whitelist", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error deleting gigavuecore_diameter_whitelist", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error deleting gigavuecore_diameter_whitelist", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error deleting gigavuecore_diameter_whitelist", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error deleting gigavuecore_diameter_whitelist", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error deleting gigavuecore_diameter_whitelist", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error deleting gigavuecore_diameter_whitelist", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *DiameterWhitelistResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *DiameterWhitelistResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("alias"), req.ID)...)
}
