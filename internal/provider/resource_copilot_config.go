package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
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
	_ resource.Resource                = (*CopilotConfigResource)(nil)
	_ resource.ResourceWithImportState = (*CopilotConfigResource)(nil)
	_ resource.ResourceWithConfigure   = (*CopilotConfigResource)(nil)
)

// CopilotConfigResource is the generated Terraform managed resource implementation.
type CopilotConfigResource struct {
	client *client.Client
}

// CopilotConfigResourceModel describes the Terraform state and plan shape for CopilotConfigResource.
type CopilotConfigResourceModel struct {
	ApiKey           types.String   `tfsdk:"api_key" json:"apiKey"`
	BedrockConfig    types.Object   `tfsdk:"bedrock_config" json:"bedrockConfig"`
	ConnectionStatus types.Object   `tfsdk:"connection_status" json:"connectionStatus"`
	Enabled          types.Bool     `tfsdk:"enabled"`
	Id               types.String   `tfsdk:"id"`
	Provider         types.String   `tfsdk:"provider_" json:"provider"`
	ServerUrl        types.String   `tfsdk:"server_url" json:"serverUrl"`
	TestConnection   types.Bool     `tfsdk:"test_connection" json:"testConnection"`
	Timeouts         timeouts.Value `tfsdk:"timeouts"`
}

// Metadata returns the resource type name.
func (r *CopilotConfigResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_copilot_config"
}

// Schema returns the Terraform schema for this resource.
func (r *CopilotConfigResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load the copilot configuration", Attributes: map[string]schema.Attribute{"api_key": schema.StringAttribute{MarkdownDescription: "API key for Copilot", Optional: true, Computed: true, Sensitive: true}, "bedrock_config": schema.SingleNestedAttribute{MarkdownDescription: "Bedrock specific configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"access_key_id": schema.StringAttribute{MarkdownDescription: "AWS Access Key ID for Bedrock service", Required: true, Sensitive: true}, "region": schema.StringAttribute{MarkdownDescription: "AWS region for Bedrock service", Required: true}, "secret_access_key": schema.StringAttribute{MarkdownDescription: "AWS Secret Access Key for Bedrock service", Required: true, Sensitive: true}}}, "connection_status": schema.SingleNestedAttribute{MarkdownDescription: "Connection Status Domain Class", Computed: true, Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Copilot Connection status message", Computed: true}, "status": schema.StringAttribute{MarkdownDescription: "Copilot Connection status code reason string", Computed: true}, "status_code": schema.Int64Attribute{Computed: true}}}, "enabled": schema.BoolAttribute{MarkdownDescription: "Enable/Disable the copilot feature", Required: true}, "id": schema.StringAttribute{Required: true}, "provider_": schema.StringAttribute{MarkdownDescription: "Copilot provider name. Bedrock", Optional: true, Computed: true}, "server_url": schema.StringAttribute{MarkdownDescription: "Copilot VM URL", Required: true}, "test_connection": schema.BoolAttribute{MarkdownDescription: "To decide whether to do connection validation on update", Optional: true}}, Blocks: map[string]schema.Block{"timeouts": timeouts.Block(ctx, timeouts.Opts{Create: true, Read: true, Update: true, Delete: true})}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *CopilotConfigResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan CopilotConfigResourceModel
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
func (r *CopilotConfigResource) createRemote(ctx context.Context, plan *CopilotConfigResourceModel, resp *resource.CreateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_copilot_config", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_copilot_config", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/fm/copilot/config"
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_copilot_config", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !plan.TestConnection.IsNull() {
		query.Set("testConnection", strconv.FormatBool(plan.TestConnection.ValueBool()))
	}
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_copilot_config", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error creating gigavuecore_copilot_config", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error creating gigavuecore_copilot_config", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error creating gigavuecore_copilot_config", "Access Denied. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error creating gigavuecore_copilot_config", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error creating gigavuecore_copilot_config", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error creating gigavuecore_copilot_config", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error creating gigavuecore_copilot_config", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error creating gigavuecore_copilot_config", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_copilot_config", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
	if plan.Id.IsNull() || plan.Id.IsUnknown() {
		loc := httpResp.Header.Get("Location")
		if loc != "" {
			loc = strings.TrimRight(loc, "/")
			i := strings.LastIndex(loc, "/")
			if i >= 0 {
				loc = loc[i+1:]
			}
			plan.Id = types.StringValue(loc)
		} else {
			resp.Diagnostics.AddError("Error creating gigavuecore_copilot_config", "The create response did not contain an identifier and no Location header was returned, so the resource cannot be tracked in state.")
			return
		}
	}
}

// Read refreshes the Terraform state with the latest remote values.
func (r *CopilotConfigResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state CopilotConfigResourceModel
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
func (r *CopilotConfigResource) readRemote(ctx context.Context, state *CopilotConfigResourceModel, resp *resource.ReadResponse) (removed bool) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/fm/copilot/config"
	httpReq, err := r.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_copilot_config", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !state.TestConnection.IsNull() {
		query.Set("testConnection", strconv.FormatBool(state.TestConnection.ValueBool()))
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_copilot_config", fmt.Sprintf("Could not send request: %s", err))
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
			resp.Diagnostics.AddError("Error reading gigavuecore_copilot_config", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_copilot_config", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_copilot_config", "Access Denied. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_copilot_config", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_copilot_config", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_copilot_config", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_copilot_config", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_copilot_config", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&state, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_copilot_config", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
	return
}

// Update modifies the remote resource to match the desired plan.
func (r *CopilotConfigResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan CopilotConfigResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state CopilotConfigResourceModel
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
	if plan.Id.IsNull() || plan.Id.IsUnknown() {
		if !state.Id.IsNull() && !state.Id.IsUnknown() {
			plan.Id = state.Id
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
func (r *CopilotConfigResource) updateRemote(ctx context.Context, plan *CopilotConfigResourceModel, resp *resource.UpdateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_copilot_config", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_copilot_config", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/fm/copilot/config"
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_copilot_config", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !plan.TestConnection.IsNull() {
		query.Set("testConnection", strconv.FormatBool(plan.TestConnection.ValueBool()))
	}
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_copilot_config", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error updating gigavuecore_copilot_config", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error updating gigavuecore_copilot_config", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error updating gigavuecore_copilot_config", "Access Denied. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error updating gigavuecore_copilot_config", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error updating gigavuecore_copilot_config", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error updating gigavuecore_copilot_config", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error updating gigavuecore_copilot_config", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error updating gigavuecore_copilot_config", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_copilot_config", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Delete destroys the remote resource.
func (r *CopilotConfigResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state CopilotConfigResourceModel
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
func (r *CopilotConfigResource) deleteRemote(ctx context.Context, state *CopilotConfigResourceModel, resp *resource.DeleteResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/fm/copilot/config"
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_copilot_config", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_copilot_config", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		return
	}
	if !(httpResp.StatusCode == 204) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error deleting gigavuecore_copilot_config", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error deleting gigavuecore_copilot_config", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error deleting gigavuecore_copilot_config", "Access Denied. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error deleting gigavuecore_copilot_config", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error deleting gigavuecore_copilot_config", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error deleting gigavuecore_copilot_config", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error deleting gigavuecore_copilot_config", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *CopilotConfigResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *CopilotConfigResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
