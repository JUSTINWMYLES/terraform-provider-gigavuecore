package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*ReplaceInlineSslProfileAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*ReplaceInlineSslProfileAction)(nil)

// ReplaceInlineSslProfileAction is the generated Terraform action implementation.
type ReplaceInlineSslProfileAction struct {
	client *client.Client
}

// ReplaceInlineSslProfileActionModel describes the action configuration shape.
type ReplaceInlineSslProfileActionModel struct {
	Alias                   types.String  `tfsdk:"alias"`
	BodyAlias               types.String  `tfsdk:"body_alias" json:"alias"`
	BodyClusterId           types.String  `tfsdk:"body_cluster_id" json:"clusterId"`
	Certificate             types.Dynamic `tfsdk:"certificate"`
	ClusterId               types.String  `tfsdk:"cluster_id"`
	Decrypt                 types.Dynamic `tfsdk:"decrypt"`
	DefaultAction           types.String  `tfsdk:"default_action" json:"defaultAction"`
	HighAvail               types.Dynamic `tfsdk:"high_avail" json:"highAvail"`
	InboundToolEarlyInspect types.Dynamic `tfsdk:"inbound_tool_early_inspect" json:"inboundToolEarlyInspect"`
	KeyMap                  types.Dynamic `tfsdk:"key_map" json:"keyMap"`
	Monitor                 types.String  `tfsdk:"monitor"`
	NetworkGroup            types.Dynamic `tfsdk:"network_group" json:"networkGroup"`
	NoDecrypt               types.Dynamic `tfsdk:"no_decrypt" json:"noDecrypt"`
	NonSslTcp               types.Dynamic `tfsdk:"non_ssl_tcp" json:"nonSslTcp"`
	OneArm                  types.String  `tfsdk:"one_arm" json:"oneArm"`
	ResilientInline         types.Dynamic `tfsdk:"resilient_inline" json:"resilientInline"`
	Rules                   types.Dynamic `tfsdk:"rules"`
	SplitProxy              types.Dynamic `tfsdk:"split_proxy" json:"splitProxy"`
	StartTls                types.Dynamic `tfsdk:"start_tls" json:"startTls"`
	Tcp                     types.Dynamic `tfsdk:"tcp"`
	Tool                    types.Dynamic `tfsdk:"tool"`
	ToolL3                  types.Dynamic `tfsdk:"tool_l3" json:"toolL3"`
	UrlCache                types.Dynamic `tfsdk:"url_cache" json:"urlCache"`
}

// NewReplaceInlineSslProfileAction returns a new instance of the generated action.
func NewReplaceInlineSslProfileAction() action.Action {
	return &ReplaceInlineSslProfileAction{}
}

// Metadata returns the action type name.
func (r *ReplaceInlineSslProfileAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_replace_inline_ssl_profile"
}

// Schema returns the action schema.
func (r *ReplaceInlineSslProfileAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Replace inline SSL profile", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Alias of the inline SSL profile", Required: true}, "body_alias": schema.StringAttribute{MarkdownDescription: "Alias of the inline SSL profile", Required: true}, "body_cluster_id": schema.StringAttribute{MarkdownDescription: "id of the defining cluster", Optional: true}, "certificate": schema.DynamicAttribute{Optional: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "decrypt": schema.DynamicAttribute{Optional: true}, "default_action": schema.StringAttribute{MarkdownDescription: "Action to take if none of the profile rules match", Optional: true}, "high_avail": schema.DynamicAttribute{Optional: true}, "inbound_tool_early_inspect": schema.DynamicAttribute{Optional: true}, "key_map": schema.DynamicAttribute{Optional: true}, "monitor": schema.StringAttribute{Optional: true}, "network_group": schema.DynamicAttribute{Optional: true}, "no_decrypt": schema.DynamicAttribute{Optional: true}, "non_ssl_tcp": schema.DynamicAttribute{Optional: true}, "one_arm": schema.StringAttribute{Optional: true}, "resilient_inline": schema.DynamicAttribute{Optional: true}, "rules": schema.DynamicAttribute{MarkdownDescription: "inline SSL profile rules", Optional: true}, "split_proxy": schema.DynamicAttribute{Optional: true}, "start_tls": schema.DynamicAttribute{Optional: true}, "tcp": schema.DynamicAttribute{Optional: true}, "tool": schema.DynamicAttribute{Optional: true}, "tool_l3": schema.DynamicAttribute{Optional: true}, "url_cache": schema.DynamicAttribute{Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *ReplaceInlineSslProfileAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config ReplaceInlineSslProfileActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *ReplaceInlineSslProfileAction) invokeRemote(ctx context.Context, config *ReplaceInlineSslProfileActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/inlineSsl/profiles/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_replace_inline_ssl_profile", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_replace_inline_ssl_profile", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_replace_inline_ssl_profile", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.BodyClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_replace_inline_ssl_profile", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_replace_inline_ssl_profile", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_replace_inline_ssl_profile", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_replace_inline_ssl_profile", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_replace_inline_ssl_profile", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_replace_inline_ssl_profile", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_replace_inline_ssl_profile", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_replace_inline_ssl_profile", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_replace_inline_ssl_profile", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_replace_inline_ssl_profile", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *ReplaceInlineSslProfileAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Action Configure Type", fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData))
		return
	}
	r.client = c
}
