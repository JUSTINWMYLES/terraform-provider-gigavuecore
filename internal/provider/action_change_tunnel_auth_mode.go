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
var _ action.Action = (*ChangeTunnelAuthModeAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*ChangeTunnelAuthModeAction)(nil)

// ChangeTunnelAuthModeAction is the generated Terraform action implementation.
type ChangeTunnelAuthModeAction struct {
	client *client.Client
}

// ChangeTunnelAuthModeActionModel describes the action configuration shape.
type ChangeTunnelAuthModeActionModel struct {
	TunnelAuthMode types.String `tfsdk:"tunnel_auth_mode" json:"tunnelAuthMode"`
}

// NewChangeTunnelAuthModeAction returns a new instance of the generated action.
func NewChangeTunnelAuthModeAction() action.Action {
	return &ChangeTunnelAuthModeAction{}
}

// Metadata returns the action type name.
func (r *ChangeTunnelAuthModeAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_change_tunnel_auth_mode"
}

// Schema returns the action schema.
func (r *ChangeTunnelAuthModeAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Change FMHA tunnel auth mode", Attributes: map[string]schema.Attribute{"tunnel_auth_mode": schema.StringAttribute{MarkdownDescription: "Auth mode for setting up FMHA tunnel. PSK is Pre Shared Key. PKI is Public Key Infrastructure(Certificate based)", Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *ChangeTunnelAuthModeAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config ChangeTunnelAuthModeActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *ChangeTunnelAuthModeAction) invokeRemote(ctx context.Context, config *ChangeTunnelAuthModeActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/fmHa/tunnel"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_change_tunnel_auth_mode", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_change_tunnel_auth_mode", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_change_tunnel_auth_mode", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_change_tunnel_auth_mode", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_change_tunnel_auth_mode", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_change_tunnel_auth_mode", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_change_tunnel_auth_mode", "Access Denied. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_change_tunnel_auth_mode", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_change_tunnel_auth_mode", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_change_tunnel_auth_mode", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_change_tunnel_auth_mode", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *ChangeTunnelAuthModeAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
