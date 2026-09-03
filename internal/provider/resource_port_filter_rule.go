package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	timeouts "github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	int64validator "github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	path "github.com/hashicorp/terraform-plugin-framework/path"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	validator "github.com/hashicorp/terraform-plugin-framework/schema/validator"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertions.
var (
	_ resource.Resource                = (*PortFilterRuleResource)(nil)
	_ resource.ResourceWithImportState = (*PortFilterRuleResource)(nil)
	_ resource.ResourceWithConfigure   = (*PortFilterRuleResource)(nil)
)

// PortFilterRuleResource is the generated Terraform managed resource implementation.
type PortFilterRuleResource struct {
	client *client.Client
}

// PortFilterRuleResourceModel describes the Terraform state and plan shape for PortFilterRuleResource.
type PortFilterRuleResourceModel struct {
	ClusterId types.String   `tfsdk:"cluster_id" json:"clusterId"`
	Comment   types.String   `tfsdk:"comment"`
	Matches   types.Dynamic  `tfsdk:"matches"`
	PortId    types.String   `tfsdk:"port_id" json:"portId"`
	RuleId    types.Int64    `tfsdk:"rule_id" json:"ruleId"`
	RuleType  types.String   `tfsdk:"rule_type" json:"ruleType"`
	Timeouts  timeouts.Value `tfsdk:"timeouts"`
}

// Metadata returns the resource type name.
func (r *PortFilterRuleResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_port_filter_rule"
}

// Schema returns the Terraform schema for this resource.
func (r *PortFilterRuleResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Add new filtering rule to a Port", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "comment": schema.StringAttribute{Optional: true, Computed: true}, "matches": schema.DynamicAttribute{MarkdownDescription: "Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched", Required: true}, "port_id": schema.StringAttribute{MarkdownDescription: "id of the target device Port (format: boxId_slotId_port, example: 1_1_c1)", Required: true}, "rule_id": schema.Int64Attribute{MarkdownDescription: "id of the rule to update", Required: true, Validators: []validator.Int64{int64validator.AtLeast(1)}}, "rule_type": schema.StringAttribute{MarkdownDescription: "filter rule type", Required: true}}, Blocks: map[string]schema.Block{"timeouts": timeouts.Block(ctx, timeouts.Opts{Create: true, Read: true, Update: true, Delete: true})}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *PortFilterRuleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan PortFilterRuleResourceModel
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
func (r *PortFilterRuleResource) createRemote(ctx context.Context, plan *PortFilterRuleResourceModel, resp *resource.CreateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_port_filter_rule", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	delete(body, "portId")
	delete(body, "ruleType")
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_port_filter_rule", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/portConfig/portFilters/{portId}/rules/{ruleType}"
	reqPath = strings.ReplaceAll(reqPath, "{portId}", url.PathEscape(plan.PortId.ValueString()))
	reqPath = strings.ReplaceAll(reqPath, "{ruleType}", url.PathEscape(plan.RuleType.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_port_filter_rule", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", plan.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_port_filter_rule", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error creating gigavuecore_port_filter_rule", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error creating gigavuecore_port_filter_rule", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error creating gigavuecore_port_filter_rule", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error creating gigavuecore_port_filter_rule", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error creating gigavuecore_port_filter_rule", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error creating gigavuecore_port_filter_rule", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error creating gigavuecore_port_filter_rule", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error creating gigavuecore_port_filter_rule", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error creating gigavuecore_port_filter_rule", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error creating gigavuecore_port_filter_rule", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_port_filter_rule", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
	if plan.RuleId.IsNull() || plan.RuleId.IsUnknown() {
		resp.Diagnostics.AddError("Error creating gigavuecore_port_filter_rule", "The create response did not contain an identifier for this resource.")
		return
	}
}

// Read refreshes the Terraform state with the latest remote values.
func (r *PortFilterRuleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state PortFilterRuleResourceModel
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
func (r *PortFilterRuleResource) readRemote(ctx context.Context, state *PortFilterRuleResourceModel, resp *resource.ReadResponse) (removed bool) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/portConfig/portFilters/{portId}"
	reqPath = strings.ReplaceAll(reqPath, "{portId}", url.PathEscape(state.PortId.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_port_filter_rule", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", state.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_port_filter_rule", fmt.Sprintf("Could not send request: %s", err))
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
			resp.Diagnostics.AddError("Error reading gigavuecore_port_filter_rule", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_port_filter_rule", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_port_filter_rule", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_port_filter_rule", "Entity Not Found. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_port_filter_rule", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_port_filter_rule", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_port_filter_rule", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["portFilter"]; ok {
		if m, ok := v.(map[string]any); ok {
			data = m
		}
	}
	var arr []any
	if v, ok := data["rules"]; ok {
		if m, ok := v.(map[string]any); ok {
			data = m
			for _, val := range data {
				if a, ok := val.([]any); ok {
					arr = append(arr, a...)
				}
			}
		}
	}
	var match map[string]any
	if state.RuleId.IsNull() {
		resp.Diagnostics.AddWarning("Error reading gigavuecore_port_filter_rule", "The identifier attribute \"rule_id\" is null in state, so the matching collection element cannot be identified. Falling back to the first element; this may track the wrong instance when the collection has more than one.")
	} else {
		want := fmt.Sprint(state.RuleId.ValueInt64())
		for _, item := range arr {
			if m, ok := item.(map[string]any); ok {
				if idVal, ok := m["ruleId"]; ok && fmt.Sprint(idVal) == want {
					match = m
					break
				}
			} else {
				continue
			}
		}
		if match == nil {
			removed = true
			return
		}
	}
	if match == nil && len(arr) > 0 {
		if m, ok := arr[0].(map[string]any); ok {
			match = m
		}
	}
	if match != nil {
		data = match
	}
	err = applyJSONToModel(&state, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_port_filter_rule", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
	return
}

// Update modifies the remote resource to match the desired plan.
func (r *PortFilterRuleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan PortFilterRuleResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state PortFilterRuleResourceModel
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
	if plan.RuleId.IsNull() || plan.RuleId.IsUnknown() {
		if !state.RuleId.IsNull() && !state.RuleId.IsUnknown() {
			plan.RuleId = state.RuleId
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
func (r *PortFilterRuleResource) updateRemote(ctx context.Context, plan *PortFilterRuleResourceModel, resp *resource.UpdateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_port_filter_rule", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	delete(body, "portId")
	delete(body, "ruleType")
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_port_filter_rule", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/portConfig/portFilters/{portId}/rules/{ruleType}/{ruleId}"
	reqPath = strings.ReplaceAll(reqPath, "{portId}", url.PathEscape(plan.PortId.ValueString()))
	reqPath = strings.ReplaceAll(reqPath, "{ruleType}", url.PathEscape(plan.RuleType.ValueString()))
	reqPath = strings.ReplaceAll(reqPath, "{ruleId}", url.PathEscape(strconv.FormatInt(plan.RuleId.ValueInt64(), 10)))
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_port_filter_rule", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", plan.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_port_filter_rule", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error updating gigavuecore_port_filter_rule", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error updating gigavuecore_port_filter_rule", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error updating gigavuecore_port_filter_rule", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error updating gigavuecore_port_filter_rule", "Entity Not Found. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error updating gigavuecore_port_filter_rule", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error updating gigavuecore_port_filter_rule", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error updating gigavuecore_port_filter_rule", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_port_filter_rule", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Delete destroys the remote resource.
func (r *PortFilterRuleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state PortFilterRuleResourceModel
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
func (r *PortFilterRuleResource) deleteRemote(ctx context.Context, state *PortFilterRuleResourceModel, resp *resource.DeleteResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/portConfig/portFilters/{portId}/rules/{ruleId}"
	reqPath = strings.ReplaceAll(reqPath, "{portId}", url.PathEscape(state.PortId.ValueString()))
	reqPath = strings.ReplaceAll(reqPath, "{ruleId}", url.PathEscape(strconv.FormatInt(state.RuleId.ValueInt64(), 10)))
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_port_filter_rule", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", state.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_port_filter_rule", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		return
	}
	if !(httpResp.StatusCode == 204) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error deleting gigavuecore_port_filter_rule", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error deleting gigavuecore_port_filter_rule", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error deleting gigavuecore_port_filter_rule", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error deleting gigavuecore_port_filter_rule", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error deleting gigavuecore_port_filter_rule", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error deleting gigavuecore_port_filter_rule", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error deleting gigavuecore_port_filter_rule", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error deleting gigavuecore_port_filter_rule", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error deleting gigavuecore_port_filter_rule", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *PortFilterRuleResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *PortFilterRuleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	portFilterRuleImportIDParts := strings.Split(req.ID, ":")
	if len(portFilterRuleImportIDParts) != 2 {
		resp.Diagnostics.AddError("Unexpected Import Identifier", fmt.Sprintf("Expected import identifier with format \"{port_id}:{cluster_id}\". Got %q.", req.ID))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("port_id"), portFilterRuleImportIDParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("cluster_id"), portFilterRuleImportIDParts[1])...)
}
