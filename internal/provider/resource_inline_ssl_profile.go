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
	_ resource.Resource                = (*InlineSslProfileResource)(nil)
	_ resource.ResourceWithIdentity    = (*InlineSslProfileResource)(nil)
	_ resource.ResourceWithImportState = (*InlineSslProfileResource)(nil)
	_ resource.ResourceWithConfigure   = (*InlineSslProfileResource)(nil)
)

// InlineSslProfileResource is the generated Terraform managed resource implementation.
type InlineSslProfileResource struct {
	client *client.Client
}

// InlineSslProfileResourceModel describes the Terraform state and plan shape for InlineSslProfileResource.
type InlineSslProfileResourceModel struct {
	Alias                   types.String   `tfsdk:"alias"`
	Certificate             types.Object   `tfsdk:"certificate"`
	ClusterId               types.String   `tfsdk:"cluster_id" json:"clusterId"`
	Decrypt                 types.Object   `tfsdk:"decrypt"`
	DefaultAction           types.String   `tfsdk:"default_action" json:"defaultAction"`
	HighAvail               types.Object   `tfsdk:"high_avail" json:"highAvail"`
	InboundToolEarlyInspect types.Object   `tfsdk:"inbound_tool_early_inspect" json:"inboundToolEarlyInspect"`
	KeyMap                  types.List     `tfsdk:"key_map" json:"keyMap"`
	Monitor                 types.String   `tfsdk:"monitor"`
	NetworkGroup            types.Object   `tfsdk:"network_group" json:"networkGroup"`
	NoDecrypt               types.Object   `tfsdk:"no_decrypt" json:"noDecrypt"`
	NonSslTcp               types.Object   `tfsdk:"non_ssl_tcp" json:"nonSslTcp"`
	OneArm                  types.String   `tfsdk:"one_arm" json:"oneArm"`
	ResilientInline         types.Object   `tfsdk:"resilient_inline" json:"resilientInline"`
	Rules                   types.Dynamic  `tfsdk:"rules"`
	SplitProxy              types.Object   `tfsdk:"split_proxy" json:"splitProxy"`
	StartTls                types.Object   `tfsdk:"start_tls" json:"startTls"`
	Tcp                     types.Object   `tfsdk:"tcp"`
	Tool                    types.Object   `tfsdk:"tool"`
	ToolL3                  types.Object   `tfsdk:"tool_l3" json:"toolL3"`
	UrlCache                types.Object   `tfsdk:"url_cache" json:"urlCache"`
	Timeouts                timeouts.Value `tfsdk:"timeouts"`
}

// Metadata returns the resource type name.
func (r *InlineSslProfileResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_inline_ssl_profile"
}

// Schema returns the Terraform schema for this resource.
func (r *InlineSslProfileResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load inline SSL profile", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Alias of the inline SSL profile", Required: true}, "certificate": schema.SingleNestedAttribute{MarkdownDescription: "SSL profile certificate configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"expired": schema.StringAttribute{MarkdownDescription: "SSL profile on expired certificate", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("decrypt", "drop")}}, "invalid": schema.StringAttribute{MarkdownDescription: "SSL profile on invalid certificate", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("decrypt", "drop")}}, "revocation": schema.SingleNestedAttribute{MarkdownDescription: "SSL profile certificate revocation configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"crl": schema.SingleNestedAttribute{MarkdownDescription: "SSL profile certificate revocation configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"defer": schema.Int64Attribute{MarkdownDescription: "timeout in seconds", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(20, 100)}}, "enabled": schema.BoolAttribute{Optional: true, Computed: true}, "fail": schema.StringAttribute{MarkdownDescription: "only applicable when enabled", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("soft", "hard")}}}}, "ocsp": schema.SingleNestedAttribute{MarkdownDescription: "SSL profile certificate revocation configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"defer": schema.Int64Attribute{MarkdownDescription: "timeout in seconds", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(20, 100)}}, "enabled": schema.BoolAttribute{Optional: true, Computed: true}, "fail": schema.StringAttribute{MarkdownDescription: "only applicable when enabled", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("soft", "hard")}}}}}}, "self_signed": schema.StringAttribute{MarkdownDescription: "SSL profile on self-signed certificate", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("decrypt", "drop")}}, "unknown_ca": schema.StringAttribute{MarkdownDescription: "SSL profile on unknown CA certificate", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("decrypt", "drop")}}}}, "cluster_id": schema.StringAttribute{MarkdownDescription: "id of the defining cluster", Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "decrypt": schema.SingleNestedAttribute{MarkdownDescription: "SSL profile configuration on decrypt action", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"tcp": schema.SingleNestedAttribute{MarkdownDescription: "SSL decryption TCP control", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"inactive_timeout": schema.Int64Attribute{MarkdownDescription: "SSL decryption TCP inactive timeout (in minutes)", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(2, 1440)}}, "port_map": schema.SingleNestedAttribute{MarkdownDescription: "SSL decryption port map", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"default_out_port": schema.Int64Attribute{MarkdownDescription: "egress port for decryption port map. 0 is disabled.", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 65536)}}, "ports": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"in_port": schema.Int64Attribute{MarkdownDescription: "ingress port for decryption port map", Required: true, Validators: []validator.Int64{int64validator.Between(1, 65536)}}, "out_port": schema.Int64Attribute{MarkdownDescription: "egress port for decryption port map", Required: true, Validators: []validator.Int64{int64validator.Between(1, 65536)}}, "rule_id": schema.Int64Attribute{Optional: true, Computed: true}}}}}}}}, "tool_bypass": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enable": schema.BoolAttribute{Optional: true, Computed: true}}}}}, "default_action": schema.StringAttribute{MarkdownDescription: "Action to take if none of the profile rules match", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("decrypt", "no-decrypt")}}, "high_avail": schema.SingleNestedAttribute{MarkdownDescription: "SSL profile configuration on high availability", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"active_standby": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enable": schema.BoolAttribute{Optional: true, Computed: true}}}}}, "inbound_tool_early_inspect": schema.SingleNestedAttribute{MarkdownDescription: "SSL profile configuration for InboundToolEarlyInspect", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"connection_timeout": schema.Int64Attribute{MarkdownDescription: "connection timeout timeout in seconds.", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 10)}}, "mode": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enable": schema.BoolAttribute{Optional: true, Computed: true}}}}}, "key_map": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"hostname": schema.StringAttribute{MarkdownDescription: "hostname or IP address", Required: true}, "key": schema.StringAttribute{MarkdownDescription: "SSL key alias", Required: true}, "rule_id": schema.Int64Attribute{Optional: true, Computed: true}}}}, "monitor": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("enable", "disable", "inline")}}, "network_group": schema.SingleNestedAttribute{MarkdownDescription: "SSL Profile configuration for multiple entry in network groups", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"multiple_entry": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enable": schema.BoolAttribute{Optional: true, Computed: true}}}}}, "no_decrypt": schema.SingleNestedAttribute{MarkdownDescription: "SSL profile configuration on no-decrypt action", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"tool_bypass": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enable": schema.BoolAttribute{Optional: true, Computed: true}}}}}, "non_ssl_tcp": schema.SingleNestedAttribute{MarkdownDescription: "SSL profile configuration on TCP proxy action", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"tool_bypass": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enable": schema.BoolAttribute{Optional: true, Computed: true}}}}}, "one_arm": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("enable", "disable")}}, "resilient_inline": schema.SingleNestedAttribute{MarkdownDescription: "SSL profile configuration for Resilient Inline Arrangement", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"mode": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enable": schema.BoolAttribute{Optional: true, Computed: true}}}}}, "rules": schema.DynamicAttribute{MarkdownDescription: "inline SSL profile rules", Optional: true, Computed: true}, "split_proxy": schema.SingleNestedAttribute{MarkdownDescription: "SSL profile configuration for Split Proxy", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"mode": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enable": schema.BoolAttribute{Optional: true, Computed: true}}}, "server_non_pfs_ciphers": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enable": schema.BoolAttribute{Optional: true, Computed: true}}}}}, "start_tls": schema.SingleNestedAttribute{MarkdownDescription: "SSL profile configuration on start TLS action", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"l4_port": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.Int64Type}}}, "tcp": schema.SingleNestedAttribute{MarkdownDescription: "SSL profile configuration on TCP", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"delayed_ack": schema.BoolAttribute{MarkdownDescription: "enable/disable TCP delayed ACK", Optional: true, Computed: true}, "syn_retries": schema.Int64Attribute{MarkdownDescription: "TCP Sync retries", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 12)}}, "timewait_timeout": schema.Int64Attribute{MarkdownDescription: "TCP Wait Timeout value", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 300)}}}}, "tool": schema.SingleNestedAttribute{MarkdownDescription: "SSL Profile configuration for Tools", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"early_engage": schema.BoolAttribute{MarkdownDescription: "enable/disable tool early engage", Optional: true, Computed: true}, "fail_action": schema.StringAttribute{MarkdownDescription: "Action to take if the tool fails", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("fail-open", "fail-close")}}}}, "tool_l3": schema.SingleNestedAttribute{MarkdownDescription: "SSL profile configuration for Layer 3 ISSL", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"cache_server_cert_timeout": schema.Int64Attribute{MarkdownDescription: "cache server timeout in seconds.", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 1440)}}, "http2_downgrade": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enable": schema.BoolAttribute{Optional: true, Computed: true}}}, "nat_pat": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enable": schema.BoolAttribute{Optional: true, Computed: true}}}}}, "url_cache": schema.SingleNestedAttribute{MarkdownDescription: "SSL profile configuration on url-cache", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"miss_action": schema.StringAttribute{MarkdownDescription: "The action to take if local URL category resolution misses", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("decrypt", "no-decrypt", "defer")}}, "timeout": schema.Int64Attribute{MarkdownDescription: "defer timeout in seconds. Only applicable for missAction 'defer'", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 10)}}}}}, Blocks: map[string]schema.Block{"timeouts": timeouts.Block(ctx, timeouts.Opts{Create: true, Read: true, Update: true, Delete: true})}}
}

// IdentitySchema returns the resource identity schema shared with the paired list resource.
func (r *InlineSslProfileResource) IdentitySchema(_ context.Context, _ resource.IdentitySchemaRequest, resp *resource.IdentitySchemaResponse) {
	resp.IdentitySchema = identityschema.Schema{Attributes: map[string]identityschema.Attribute{"alias": identityschema.StringAttribute{RequiredForImport: true, Description: "Alias of the inline SSL profile"}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *InlineSslProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InlineSslProfileResourceModel
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
func (r *InlineSslProfileResource) createRemote(ctx context.Context, plan *InlineSslProfileResourceModel, resp *resource.CreateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_inline_ssl_profile", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_inline_ssl_profile", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/apps/inlineSsl/profiles"
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_inline_ssl_profile", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", plan.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_inline_ssl_profile", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error creating gigavuecore_inline_ssl_profile", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error creating gigavuecore_inline_ssl_profile", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error creating gigavuecore_inline_ssl_profile", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error creating gigavuecore_inline_ssl_profile", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error creating gigavuecore_inline_ssl_profile", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error creating gigavuecore_inline_ssl_profile", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error creating gigavuecore_inline_ssl_profile", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error creating gigavuecore_inline_ssl_profile", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error creating gigavuecore_inline_ssl_profile", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error creating gigavuecore_inline_ssl_profile", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_inline_ssl_profile", fmt.Sprintf("Could not map response to state: %s", err))
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
			resp.Diagnostics.AddError("Error creating gigavuecore_inline_ssl_profile", "The create response did not contain an identifier and no Location header was returned, so the resource cannot be tracked in state.")
			return
		}
	}
}

// Read refreshes the Terraform state with the latest remote values.
func (r *InlineSslProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InlineSslProfileResourceModel
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
func (r *InlineSslProfileResource) readRemote(ctx context.Context, state *InlineSslProfileResourceModel, resp *resource.ReadResponse) (removed bool) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/inlineSsl/profiles/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(state.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_inline_ssl_profile", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", state.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_inline_ssl_profile", fmt.Sprintf("Could not send request: %s", err))
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
			resp.Diagnostics.AddError("Error reading gigavuecore_inline_ssl_profile", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_inline_ssl_profile", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_inline_ssl_profile", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_inline_ssl_profile", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_inline_ssl_profile", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_inline_ssl_profile", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_inline_ssl_profile", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_inline_ssl_profile", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_inline_ssl_profile", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_inline_ssl_profile", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["inlineSslProfile"]; ok {
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
		resp.Diagnostics.AddError("Error reading gigavuecore_inline_ssl_profile", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
	return
}

// Update modifies the remote resource to match the desired plan.
func (r *InlineSslProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan InlineSslProfileResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state InlineSslProfileResourceModel
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
func (r *InlineSslProfileResource) updateRemote(ctx context.Context, plan *InlineSslProfileResourceModel, resp *resource.UpdateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_inline_ssl_profile", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_inline_ssl_profile", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/apps/inlineSsl/profiles/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(plan.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_inline_ssl_profile", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", plan.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_inline_ssl_profile", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error updating gigavuecore_inline_ssl_profile", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error updating gigavuecore_inline_ssl_profile", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error updating gigavuecore_inline_ssl_profile", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error updating gigavuecore_inline_ssl_profile", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error updating gigavuecore_inline_ssl_profile", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error updating gigavuecore_inline_ssl_profile", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error updating gigavuecore_inline_ssl_profile", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error updating gigavuecore_inline_ssl_profile", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error updating gigavuecore_inline_ssl_profile", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error updating gigavuecore_inline_ssl_profile", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_inline_ssl_profile", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Delete destroys the remote resource.
func (r *InlineSslProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InlineSslProfileResourceModel
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
func (r *InlineSslProfileResource) deleteRemote(ctx context.Context, state *InlineSslProfileResourceModel, resp *resource.DeleteResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/inlineSsl/profiles/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(state.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_inline_ssl_profile", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", state.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_inline_ssl_profile", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		return
	}
	if !(httpResp.StatusCode == 204) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error deleting gigavuecore_inline_ssl_profile", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error deleting gigavuecore_inline_ssl_profile", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error deleting gigavuecore_inline_ssl_profile", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error deleting gigavuecore_inline_ssl_profile", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error deleting gigavuecore_inline_ssl_profile", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error deleting gigavuecore_inline_ssl_profile", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error deleting gigavuecore_inline_ssl_profile", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error deleting gigavuecore_inline_ssl_profile", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error deleting gigavuecore_inline_ssl_profile", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *InlineSslProfileResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *InlineSslProfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	inlineSslProfileImportIDParts := strings.Split(req.ID, ":")
	if len(inlineSslProfileImportIDParts) != 2 {
		resp.Diagnostics.AddError("Unexpected Import Identifier", fmt.Sprintf("Expected import identifier with format \"{alias}:{cluster_id}\". Got %q.", req.ID))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("alias"), inlineSslProfileImportIDParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("cluster_id"), inlineSslProfileImportIDParts[1])...)
}
