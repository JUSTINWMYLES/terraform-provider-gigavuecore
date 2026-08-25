package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*UpdateInlineSslConfigAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*UpdateInlineSslConfigAction)(nil)

// UpdateInlineSslConfigAction is the generated Terraform action implementation.
type UpdateInlineSslConfigAction struct {
	client *client.Client
}

// UpdateInlineSslConfigActionModel describes the action configuration shape.
type UpdateInlineSslConfigActionModel struct {
	BodyClusterId types.String  `tfsdk:"body_cluster_id" json:"clusterId"`
	Caching       types.Dynamic `tfsdk:"caching"`
	ClusterId     types.String  `tfsdk:"cluster_id"`
	DheCiphersuit types.String  `tfsdk:"dhe_ciphersuit" json:"dheCiphersuit"`
	Monitor       types.Dynamic `tfsdk:"monitor"`
	Resumption    types.Dynamic `tfsdk:"resumption"`
	SslVersions   types.Dynamic `tfsdk:"ssl_versions" json:"sslVersions"`
	StartTls      types.Dynamic `tfsdk:"start_tls" json:"startTls"`
}

// NewUpdateInlineSslConfigAction returns a new instance of the generated action.
func NewUpdateInlineSslConfigAction() action.Action {
	return &UpdateInlineSslConfigAction{}
}

// Metadata returns the action type name.
func (r *UpdateInlineSslConfigAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_update_inline_ssl_config"
}

// Schema returns the action schema.
func (r *UpdateInlineSslConfigAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Update inline SSL global configuration", Attributes: map[string]schema.Attribute{"body_cluster_id": schema.StringAttribute{MarkdownDescription: "id of the defining cluster", Optional: true}, "caching": schema.DynamicAttribute{Optional: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "dhe_ciphersuit": schema.StringAttribute{Optional: true}, "monitor": schema.DynamicAttribute{Optional: true}, "resumption": schema.DynamicAttribute{Optional: true}, "ssl_versions": schema.DynamicAttribute{Optional: true}, "start_tls": schema.DynamicAttribute{Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *UpdateInlineSslConfigAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UpdateInlineSslConfigActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *UpdateInlineSslConfigAction) invokeRemote(ctx context.Context, config *UpdateInlineSslConfigActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/inlineSsl"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_inline_ssl_config", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_inline_ssl_config", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPatch, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_inline_ssl_config", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.BodyClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_inline_ssl_config", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_inline_ssl_config", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_inline_ssl_config", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_inline_ssl_config", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_inline_ssl_config", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_inline_ssl_config", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_inline_ssl_config", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_inline_ssl_config", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_update_inline_ssl_config", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_inline_ssl_config", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *UpdateInlineSslConfigAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
