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
	_ resource.Resource                = (*ApplicationProfileResource)(nil)
	_ resource.ResourceWithIdentity    = (*ApplicationProfileResource)(nil)
	_ resource.ResourceWithImportState = (*ApplicationProfileResource)(nil)
	_ resource.ResourceWithConfigure   = (*ApplicationProfileResource)(nil)
)

// ApplicationProfileResource is the generated Terraform managed resource implementation.
type ApplicationProfileResource struct {
	client *client.Client
}

// ApplicationProfileResourceModel describes the Terraform state and plan shape for ApplicationProfileResource.
type ApplicationProfileResourceModel struct {
	Alias         types.String                     `tfsdk:"alias"`
	ApplicationId types.Bool                       `tfsdk:"application_id" json:"applicationId"`
	Applications  types.List                       `tfsdk:"applications"`
	Counter       types.Object                     `tfsdk:"counter"`
	Datalink      types.Object                     `tfsdk:"datalink"`
	Description   types.String                     `tfsdk:"description"`
	Flow          types.Object                     `tfsdk:"flow"`
	Gtpu          types.Object                     `tfsdk:"gtpu"`
	Interface     types.Object                     `tfsdk:"interface"`
	Ip            types.Object                     `tfsdk:"ip"`
	Ipv4          types.Object                     `tfsdk:"ipv4"`
	Ipv6          types.Object                     `tfsdk:"ipv6"`
	OuterIpv4     types.Object                     `tfsdk:"outer_ipv4" json:"outerIpv4"`
	OuterIpv6     types.Object                     `tfsdk:"outer_ipv6" json:"outerIpv6"`
	Timestamp     types.Object                     `tfsdk:"timestamp"`
	Transport     types.Object                     `tfsdk:"transport"`
	Type          types.String                     `tfsdk:"type"`
	Timeouts      *ApplicationProfileTimeoutsModel `tfsdk:"timeouts"`
}

// ApplicationProfileTimeoutsModel describes the per-operation timeout configuration (in seconds) for the gigavuecore_application_profile resource.
type ApplicationProfileTimeoutsModel struct {
	Create types.Int64 `tfsdk:"create"`
	Read   types.Int64 `tfsdk:"read"`
	Update types.Int64 `tfsdk:"update"`
	Delete types.Int64 `tfsdk:"delete"`
}

// Metadata returns the resource type name.
func (r *ApplicationProfileResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_application_profile"
}

// Schema returns the Terraform schema for this resource.
func (r *ApplicationProfileResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Create a Application Profile", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "application profile alias", Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "application_id": schema.BoolAttribute{MarkdownDescription: "only valid with 'export' type", Optional: true, Computed: true}, "applications": schema.ListNestedAttribute{MarkdownDescription: "application and attributes.", Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"attributes": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"name": schema.StringAttribute{MarkdownDescription: "attribute name", Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "value": schema.StringAttribute{MarkdownDescription: "application's attribute value", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}}}}, "is_user_defined": schema.BoolAttribute{MarkdownDescription: "Default value is false. Must be set to true only for App Intel solution while configuring user defined apps", Optional: true, Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "application name", Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}}}}, "counter": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"bytes": schema.BoolAttribute{Optional: true, Computed: true}, "bytes_long": schema.BoolAttribute{Optional: true, Computed: true}, "inner_byte": schema.BoolAttribute{Optional: true, Computed: true}, "inner_byte_long": schema.BoolAttribute{Optional: true, Computed: true}, "packets": schema.BoolAttribute{Optional: true, Computed: true}, "packets_long": schema.BoolAttribute{Optional: true, Computed: true}}}, "datalink": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"mac_dst": schema.BoolAttribute{Optional: true, Computed: true}, "mac_src": schema.BoolAttribute{Optional: true, Computed: true}, "vlan": schema.BoolAttribute{Optional: true, Computed: true}}}, "description": schema.StringAttribute{Optional: true, Computed: true}, "flow": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"end_reason": schema.BoolAttribute{Optional: true, Computed: true}}}, "gtpu": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"qfi": schema.BoolAttribute{Optional: true, Computed: true}, "teid": schema.BoolAttribute{Optional: true, Computed: true}}}, "interface": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"in_name_width": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 32)}}, "in_physical_width": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.OneOf(2, 4)}}, "out_physical_width": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.OneOf(2, 4)}}}}, "ip": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"version": schema.BoolAttribute{Optional: true, Computed: true}}}, "ipv4": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"destination": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"prefix_min_mask": schema.StringAttribute{Optional: true, Computed: true}}}, "dscp": schema.BoolAttribute{Optional: true, Computed: true}, "fragmentation": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"flags": schema.BoolAttribute{Optional: true, Computed: true}, "offset": schema.BoolAttribute{Optional: true, Computed: true}}}, "header_len": schema.BoolAttribute{Optional: true, Computed: true}, "option_map": schema.BoolAttribute{Optional: true, Computed: true}, "precedence": schema.BoolAttribute{Optional: true, Computed: true}, "protocol": schema.BoolAttribute{Optional: true, Computed: true}, "section": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"header_size": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 128)}}, "payload_size": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 128)}}}}, "source": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"prefix_min_mask": schema.StringAttribute{MarkdownDescription: "ipv4 source prefix minimum-mask - netmask or mask length", Optional: true, Computed: true}}}, "tos": schema.BoolAttribute{Optional: true, Computed: true}, "total_length": schema.BoolAttribute{Optional: true, Computed: true}, "ttl": schema.BoolAttribute{Optional: true, Computed: true}}}, "ipv6": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"destination": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"prefix_min_mask": schema.StringAttribute{Optional: true, Computed: true}}}, "dscp": schema.BoolAttribute{Optional: true, Computed: true}, "extension_map": schema.BoolAttribute{Optional: true, Computed: true}, "flow_label": schema.BoolAttribute{Optional: true, Computed: true}, "fragmentation": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"flags": schema.BoolAttribute{Optional: true, Computed: true}, "offset": schema.BoolAttribute{Optional: true, Computed: true}}}, "hop_limit": schema.BoolAttribute{Optional: true, Computed: true}, "length": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"header": schema.BoolAttribute{Optional: true, Computed: true}, "payload": schema.BoolAttribute{Optional: true, Computed: true}, "total": schema.BoolAttribute{Optional: true, Computed: true}}}, "next_header": schema.BoolAttribute{Optional: true, Computed: true}, "precedence": schema.BoolAttribute{Optional: true, Computed: true}, "section": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"header_size": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 128)}}, "payload_size": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 128)}}}}, "source": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"prefix_min_mask": schema.StringAttribute{Optional: true, Computed: true}}}, "traffic_class": schema.BoolAttribute{Optional: true, Computed: true}}}, "outer_ipv4": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"destination": schema.BoolAttribute{Optional: true, Computed: true}, "source": schema.BoolAttribute{Optional: true, Computed: true}}}, "outer_ipv6": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"destination": schema.BoolAttribute{Optional: true, Computed: true}, "source": schema.BoolAttribute{Optional: true, Computed: true}}}, "timestamp": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"flow_end_msec": schema.BoolAttribute{Optional: true, Computed: true}, "flow_endsec": schema.BoolAttribute{Optional: true, Computed: true}, "flow_start_msec": schema.BoolAttribute{Optional: true, Computed: true}, "flow_startsec": schema.BoolAttribute{Optional: true, Computed: true}, "sys_up_time_first": schema.BoolAttribute{Optional: true, Computed: true}, "sys_up_time_last": schema.BoolAttribute{Optional: true, Computed: true}}}, "transport": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"dst_port": schema.BoolAttribute{Optional: true, Computed: true}, "icmp": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"ipv4_code": schema.BoolAttribute{Optional: true, Computed: true}, "ipv4_type": schema.BoolAttribute{Optional: true, Computed: true}, "ipv6_code": schema.BoolAttribute{Optional: true, Computed: true}, "ipv6_type": schema.BoolAttribute{Optional: true, Computed: true}}}, "src_port": schema.BoolAttribute{Optional: true, Computed: true}, "tcp": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"ack_number": schema.BoolAttribute{Optional: true, Computed: true}, "dst_port": schema.BoolAttribute{Optional: true, Computed: true}, "flags": schema.BoolAttribute{Optional: true, Computed: true}, "header_len": schema.BoolAttribute{Optional: true, Computed: true}, "seq_number": schema.BoolAttribute{Optional: true, Computed: true}, "src_port": schema.BoolAttribute{Optional: true, Computed: true}, "urgent_ptr": schema.BoolAttribute{Optional: true, Computed: true}, "window_size": schema.BoolAttribute{Optional: true, Computed: true}}}, "udp": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"dst_port": schema.BoolAttribute{Optional: true, Computed: true}, "msg_len": schema.BoolAttribute{Optional: true, Computed: true}, "src_port": schema.BoolAttribute{Optional: true, Computed: true}}}}}, "type": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("export", "filter")}}}, Blocks: map[string]schema.Block{"timeouts": schema.SingleNestedBlock{Attributes: map[string]schema.Attribute{"create": schema.Int64Attribute{MarkdownDescription: "Maximum time in seconds for the create operation. Defaults to 1200.", Optional: true}, "read": schema.Int64Attribute{MarkdownDescription: "Maximum time in seconds for the read operation. Defaults to 1200.", Optional: true}, "update": schema.Int64Attribute{MarkdownDescription: "Maximum time in seconds for the update operation. Defaults to 1200.", Optional: true}, "delete": schema.Int64Attribute{MarkdownDescription: "Maximum time in seconds for the delete operation. Defaults to 1200.", Optional: true}}, MarkdownDescription: "Per-operation timeouts in seconds."}}}
}

// IdentitySchema returns the resource identity schema shared with the paired list resource.
func (r *ApplicationProfileResource) IdentitySchema(_ context.Context, _ resource.IdentitySchemaRequest, resp *resource.IdentitySchemaResponse) {
	resp.IdentitySchema = identityschema.Schema{Attributes: map[string]identityschema.Attribute{"alias": identityschema.StringAttribute{RequiredForImport: true, Description: "application profile alias"}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *ApplicationProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ApplicationProfileResourceModel
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
func (r *ApplicationProfileResource) createRemote(ctx context.Context, plan *ApplicationProfileResourceModel, resp *resource.CreateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_application_profile", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_application_profile", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/apps/metadata/applicationProfiles"
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_application_profile", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_application_profile", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error creating gigavuecore_application_profile", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error creating gigavuecore_application_profile", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error creating gigavuecore_application_profile", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error creating gigavuecore_application_profile", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error creating gigavuecore_application_profile", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error creating gigavuecore_application_profile", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error creating gigavuecore_application_profile", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error creating gigavuecore_application_profile", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error creating gigavuecore_application_profile", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error creating gigavuecore_application_profile", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_application_profile", fmt.Sprintf("Could not map response to state: %s", err))
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
			resp.Diagnostics.AddError("Error creating gigavuecore_application_profile", "The create response did not contain an identifier and no Location header was returned, so the resource cannot be tracked in state.")
			return
		}
	}
}

// Read refreshes the Terraform state with the latest remote values.
func (r *ApplicationProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ApplicationProfileResourceModel
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
func (r *ApplicationProfileResource) readRemote(ctx context.Context, state *ApplicationProfileResourceModel, resp *resource.ReadResponse) (removed bool) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/metadata/applicationProfiles/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(state.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_application_profile", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_application_profile", fmt.Sprintf("Could not send request: %s", err))
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
			resp.Diagnostics.AddError("Error reading gigavuecore_application_profile", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_application_profile", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_application_profile", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_application_profile", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_application_profile", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_application_profile", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_application_profile", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_application_profile", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_application_profile", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_application_profile", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["applicationProfile"]; ok {
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
		resp.Diagnostics.AddError("Error reading gigavuecore_application_profile", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
	return
}

// Update modifies the remote resource to match the desired plan.
func (r *ApplicationProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ApplicationProfileResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state ApplicationProfileResourceModel
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
func (r *ApplicationProfileResource) updateRemote(ctx context.Context, plan *ApplicationProfileResourceModel, resp *resource.UpdateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_application_profile", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_application_profile", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/apps/metadata/applicationProfiles/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(plan.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_application_profile", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_application_profile", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error updating gigavuecore_application_profile", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error updating gigavuecore_application_profile", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error updating gigavuecore_application_profile", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error updating gigavuecore_application_profile", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error updating gigavuecore_application_profile", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error updating gigavuecore_application_profile", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error updating gigavuecore_application_profile", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error updating gigavuecore_application_profile", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error updating gigavuecore_application_profile", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error updating gigavuecore_application_profile", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_application_profile", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Delete destroys the remote resource.
func (r *ApplicationProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ApplicationProfileResourceModel
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
func (r *ApplicationProfileResource) deleteRemote(ctx context.Context, state *ApplicationProfileResourceModel, resp *resource.DeleteResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/metadata/applicationProfiles/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(state.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_application_profile", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_application_profile", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		return
	}
	if !(httpResp.StatusCode == 204) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error deleting gigavuecore_application_profile", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error deleting gigavuecore_application_profile", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error deleting gigavuecore_application_profile", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error deleting gigavuecore_application_profile", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error deleting gigavuecore_application_profile", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error deleting gigavuecore_application_profile", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error deleting gigavuecore_application_profile", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error deleting gigavuecore_application_profile", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error deleting gigavuecore_application_profile", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *ApplicationProfileResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *ApplicationProfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("alias"), req.ID)...)
}
