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
	_ resource.Resource                = (*NetworkProfileResource)(nil)
	_ resource.ResourceWithIdentity    = (*NetworkProfileResource)(nil)
	_ resource.ResourceWithImportState = (*NetworkProfileResource)(nil)
	_ resource.ResourceWithConfigure   = (*NetworkProfileResource)(nil)
)

// NetworkProfileResource is the generated Terraform managed resource implementation.
type NetworkProfileResource struct {
	client *client.Client
}

// NetworkProfileResourceModel describes the Terraform state and plan shape for NetworkProfileResource.
type NetworkProfileResourceModel struct {
	Alias         types.String   `tfsdk:"alias"`
	ApplicationId types.Bool     `tfsdk:"application_id" json:"applicationId"`
	Applications  types.List     `tfsdk:"applications"`
	Counter       types.Object   `tfsdk:"counter"`
	Datalink      types.Object   `tfsdk:"datalink"`
	Description   types.String   `tfsdk:"description"`
	Flow          types.Object   `tfsdk:"flow"`
	Gtpu          types.Object   `tfsdk:"gtpu"`
	Interface     types.Object   `tfsdk:"interface"`
	Ip            types.Object   `tfsdk:"ip"`
	Ipv4          types.Object   `tfsdk:"ipv4"`
	Ipv4Subnets   types.List     `tfsdk:"ipv4_subnets" json:"ipv4Subnets"`
	Ipv6          types.Object   `tfsdk:"ipv6"`
	Ipv6Masks     types.List     `tfsdk:"ipv6_masks" json:"ipv6Masks"`
	OuterIpv4     types.Object   `tfsdk:"outer_ipv4" json:"outerIpv4"`
	OuterIpv6     types.Object   `tfsdk:"outer_ipv6" json:"outerIpv6"`
	Timestamp     types.Object   `tfsdk:"timestamp"`
	Transport     types.Object   `tfsdk:"transport"`
	Type          types.String   `tfsdk:"type"`
	Timeouts      timeouts.Value `tfsdk:"timeouts"`
}

// Metadata returns the resource type name.
func (r *NetworkProfileResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_network_profile"
}

// Schema returns the Terraform schema for this resource.
func (r *NetworkProfileResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Create a Network Profile", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "network profile alias", Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "application_id": schema.BoolAttribute{MarkdownDescription: "only valid with 'export' type", Optional: true}, "applications": schema.ListNestedAttribute{MarkdownDescription: "application and attributes.", Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"attributes": schema.ListNestedAttribute{Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"name": schema.StringAttribute{MarkdownDescription: "attribute name", Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "value": schema.StringAttribute{MarkdownDescription: "application's attribute value", Optional: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}}}}, "is_user_defined": schema.BoolAttribute{MarkdownDescription: "Default value is false. Must be set to true only for App Intel solution while configuring user defined apps", Optional: true}, "name": schema.StringAttribute{MarkdownDescription: "application name", Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}}}}, "counter": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"bytes": schema.BoolAttribute{Optional: true}, "bytes_long": schema.BoolAttribute{Optional: true}, "inner_byte": schema.BoolAttribute{Optional: true}, "inner_byte_long": schema.BoolAttribute{Optional: true}, "packets": schema.BoolAttribute{Optional: true}, "packets_long": schema.BoolAttribute{Optional: true}}}, "datalink": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"mac_dst": schema.BoolAttribute{Optional: true}, "mac_src": schema.BoolAttribute{Optional: true}, "vlan": schema.BoolAttribute{Optional: true}}}, "description": schema.StringAttribute{Optional: true, Computed: true}, "flow": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"end_reason": schema.BoolAttribute{Optional: true}}}, "gtpu": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"qfi": schema.BoolAttribute{Optional: true}, "teid": schema.BoolAttribute{Optional: true}}}, "interface": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"in_name_width": schema.Int64Attribute{Optional: true, Validators: []validator.Int64{int64validator.Between(1, 32)}}, "in_physical_width": schema.Int64Attribute{Optional: true, Validators: []validator.Int64{int64validator.OneOf(2, 4)}}, "out_physical_width": schema.Int64Attribute{Optional: true, Validators: []validator.Int64{int64validator.OneOf(2, 4)}}}}, "ip": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"version": schema.BoolAttribute{Optional: true}}}, "ipv4": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"destination": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"prefix_min_mask": schema.StringAttribute{Optional: true}}}, "dscp": schema.BoolAttribute{Optional: true}, "fragmentation": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"flags": schema.BoolAttribute{Optional: true}, "offset": schema.BoolAttribute{Optional: true}}}, "header_len": schema.BoolAttribute{Optional: true}, "option_map": schema.BoolAttribute{Optional: true}, "precedence": schema.BoolAttribute{Optional: true}, "protocol": schema.BoolAttribute{Optional: true}, "section": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"header_size": schema.Int64Attribute{Optional: true, Validators: []validator.Int64{int64validator.Between(1, 128)}}, "payload_size": schema.Int64Attribute{Optional: true, Validators: []validator.Int64{int64validator.Between(1, 128)}}}}, "source": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"prefix_min_mask": schema.StringAttribute{MarkdownDescription: "ipv4 source prefix minimum-mask - netmask or mask length", Optional: true}}}, "tos": schema.BoolAttribute{Optional: true}, "total_length": schema.BoolAttribute{Optional: true}, "ttl": schema.BoolAttribute{Optional: true}}}, "ipv4_subnets": schema.ListAttribute{MarkdownDescription: "ipv4 subnets, to identify client", Optional: true, Computed: true, ElementType: types.StringType}, "ipv6": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"destination": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"prefix_min_mask": schema.StringAttribute{Optional: true}}}, "dscp": schema.BoolAttribute{Optional: true}, "extension_map": schema.BoolAttribute{Optional: true}, "flow_label": schema.BoolAttribute{Optional: true}, "fragmentation": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"flags": schema.BoolAttribute{Optional: true}, "offset": schema.BoolAttribute{Optional: true}}}, "hop_limit": schema.BoolAttribute{Optional: true}, "length": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"header": schema.BoolAttribute{Optional: true}, "payload": schema.BoolAttribute{Optional: true}, "total": schema.BoolAttribute{Optional: true}}}, "next_header": schema.BoolAttribute{Optional: true}, "precedence": schema.BoolAttribute{Optional: true}, "section": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"header_size": schema.Int64Attribute{Optional: true, Validators: []validator.Int64{int64validator.Between(1, 128)}}, "payload_size": schema.Int64Attribute{Optional: true, Validators: []validator.Int64{int64validator.Between(1, 128)}}}}, "source": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"prefix_min_mask": schema.StringAttribute{Optional: true}}}, "traffic_class": schema.BoolAttribute{Optional: true}}}, "ipv6_masks": schema.ListAttribute{MarkdownDescription: "ipv6 subnets, to identify client", Optional: true, Computed: true, ElementType: types.StringType}, "outer_ipv4": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"destination": schema.BoolAttribute{Optional: true}, "source": schema.BoolAttribute{Optional: true}}}, "outer_ipv6": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"destination": schema.BoolAttribute{Optional: true}, "source": schema.BoolAttribute{Optional: true}}}, "timestamp": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"flow_end_msec": schema.BoolAttribute{Optional: true}, "flow_endsec": schema.BoolAttribute{Optional: true}, "flow_start_msec": schema.BoolAttribute{Optional: true}, "flow_startsec": schema.BoolAttribute{Optional: true}, "sys_up_time_first": schema.BoolAttribute{Optional: true}, "sys_up_time_last": schema.BoolAttribute{Optional: true}}}, "transport": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"dst_port": schema.BoolAttribute{Optional: true}, "icmp": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"ipv4_code": schema.BoolAttribute{Optional: true}, "ipv4_type": schema.BoolAttribute{Optional: true}, "ipv6_code": schema.BoolAttribute{Optional: true}, "ipv6_type": schema.BoolAttribute{Optional: true}}}, "src_port": schema.BoolAttribute{Optional: true}, "tcp": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"ack_number": schema.BoolAttribute{Optional: true}, "dst_port": schema.BoolAttribute{Optional: true}, "flags": schema.BoolAttribute{Optional: true}, "header_len": schema.BoolAttribute{Optional: true}, "seq_number": schema.BoolAttribute{Optional: true}, "src_port": schema.BoolAttribute{Optional: true}, "urgent_ptr": schema.BoolAttribute{Optional: true}, "window_size": schema.BoolAttribute{Optional: true}}}, "udp": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"dst_port": schema.BoolAttribute{Optional: true}, "msg_len": schema.BoolAttribute{Optional: true}, "src_port": schema.BoolAttribute{Optional: true}}}}}, "type": schema.StringAttribute{Optional: true, Validators: []validator.String{stringvalidator.OneOf("export", "filter")}}}, Blocks: map[string]schema.Block{"timeouts": timeouts.Block(ctx, timeouts.Opts{Create: true, Read: true, Update: true, Delete: true})}}
}

// IdentitySchema returns the resource identity schema shared with the paired list resource.
func (r *NetworkProfileResource) IdentitySchema(_ context.Context, _ resource.IdentitySchemaRequest, resp *resource.IdentitySchemaResponse) {
	resp.IdentitySchema = identityschema.Schema{Attributes: map[string]identityschema.Attribute{"alias": identityschema.StringAttribute{RequiredForImport: true, Description: "network profile alias"}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *NetworkProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan NetworkProfileResourceModel
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
func (r *NetworkProfileResource) createRemote(ctx context.Context, plan *NetworkProfileResourceModel, resp *resource.CreateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_network_profile", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_network_profile", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/apps/metadata/networkProfiles"
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_network_profile", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_network_profile", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error creating gigavuecore_network_profile", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error creating gigavuecore_network_profile", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error creating gigavuecore_network_profile", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error creating gigavuecore_network_profile", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error creating gigavuecore_network_profile", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error creating gigavuecore_network_profile", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error creating gigavuecore_network_profile", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error creating gigavuecore_network_profile", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error creating gigavuecore_network_profile", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error creating gigavuecore_network_profile", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_network_profile", fmt.Sprintf("Could not map response to state: %s", err))
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
			resp.Diagnostics.AddError("Error creating gigavuecore_network_profile", "The create response did not contain an identifier and no Location header was returned, so the resource cannot be tracked in state.")
			return
		}
	}
}

// Read refreshes the Terraform state with the latest remote values.
func (r *NetworkProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state NetworkProfileResourceModel
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
func (r *NetworkProfileResource) readRemote(ctx context.Context, state *NetworkProfileResourceModel, resp *resource.ReadResponse) (removed bool) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/metadata/networkProfiles/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(state.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_network_profile", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_network_profile", fmt.Sprintf("Could not send request: %s", err))
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
			resp.Diagnostics.AddError("Error reading gigavuecore_network_profile", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_network_profile", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_network_profile", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_network_profile", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_network_profile", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_network_profile", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_network_profile", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_network_profile", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_network_profile", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_network_profile", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["networkProfile"]; ok {
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
		resp.Diagnostics.AddError("Error reading gigavuecore_network_profile", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
	return
}

// Update modifies the remote resource to match the desired plan.
func (r *NetworkProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan NetworkProfileResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state NetworkProfileResourceModel
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
func (r *NetworkProfileResource) updateRemote(ctx context.Context, plan *NetworkProfileResourceModel, resp *resource.UpdateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_network_profile", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_network_profile", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/apps/metadata/networkProfiles/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(plan.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_network_profile", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_network_profile", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error updating gigavuecore_network_profile", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error updating gigavuecore_network_profile", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error updating gigavuecore_network_profile", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error updating gigavuecore_network_profile", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error updating gigavuecore_network_profile", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error updating gigavuecore_network_profile", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error updating gigavuecore_network_profile", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error updating gigavuecore_network_profile", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error updating gigavuecore_network_profile", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error updating gigavuecore_network_profile", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_network_profile", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Delete destroys the remote resource.
func (r *NetworkProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworkProfileResourceModel
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
func (r *NetworkProfileResource) deleteRemote(ctx context.Context, state *NetworkProfileResourceModel, resp *resource.DeleteResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/metadata/networkProfiles/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(state.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_network_profile", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_network_profile", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		return
	}
	if !(httpResp.StatusCode == 204) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error deleting gigavuecore_network_profile", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error deleting gigavuecore_network_profile", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error deleting gigavuecore_network_profile", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error deleting gigavuecore_network_profile", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error deleting gigavuecore_network_profile", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error deleting gigavuecore_network_profile", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error deleting gigavuecore_network_profile", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error deleting gigavuecore_network_profile", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error deleting gigavuecore_network_profile", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *NetworkProfileResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *NetworkProfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("alias"), req.ID)...)
}
