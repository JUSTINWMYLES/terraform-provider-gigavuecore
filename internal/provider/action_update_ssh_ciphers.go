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
var _ action.Action = (*UpdateSshCiphersAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*UpdateSshCiphersAction)(nil)

// UpdateSshCiphersAction is the generated Terraform action implementation.
type UpdateSshCiphersAction struct {
	client *client.Client
}

// UpdateSshCiphersActionModel describes the action configuration shape.
type UpdateSshCiphersActionModel struct {
	ClientCiphers types.List `tfsdk:"client_ciphers" json:"clientCiphers"`
	ClientHostkey types.List `tfsdk:"client_hostkey" json:"clientHostkey"`
	ClientKex     types.List `tfsdk:"client_kex" json:"clientKex"`
	ClientMacs    types.List `tfsdk:"client_macs" json:"clientMacs"`
	ServerCiphers types.List `tfsdk:"server_ciphers" json:"serverCiphers"`
	ServerHostkey types.List `tfsdk:"server_hostkey" json:"serverHostkey"`
	ServerKex     types.List `tfsdk:"server_kex" json:"serverKex"`
	ServerMacs    types.List `tfsdk:"server_macs" json:"serverMacs"`
}

// NewUpdateSshCiphersAction returns a new instance of the generated action.
func NewUpdateSshCiphersAction() action.Action {
	return &UpdateSshCiphersAction{}
}

// Metadata returns the action type name.
func (r *UpdateSshCiphersAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_update_ssh_ciphers"
}

// Schema returns the action schema.
func (r *UpdateSshCiphersAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Update all ciphers, Kex, Macs and HostKey", Attributes: map[string]schema.Attribute{"client_ciphers": schema.ListAttribute{Required: true, ElementType: types.StringType}, "client_hostkey": schema.ListAttribute{Required: true, ElementType: types.StringType}, "client_kex": schema.ListAttribute{Required: true, ElementType: types.StringType}, "client_macs": schema.ListAttribute{Required: true, ElementType: types.StringType}, "server_ciphers": schema.ListAttribute{Required: true, ElementType: types.StringType}, "server_hostkey": schema.ListAttribute{Required: true, ElementType: types.StringType}, "server_kex": schema.ListAttribute{Required: true, ElementType: types.StringType}, "server_macs": schema.ListAttribute{Required: true, ElementType: types.StringType}}}
}

// Invoke executes the action against the remote API.
func (r *UpdateSshCiphersAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UpdateSshCiphersActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *UpdateSshCiphersAction) invokeRemote(ctx context.Context, config *UpdateSshCiphersActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/system/ssh/config"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_ssh_ciphers", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_ssh_ciphers", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_ssh_ciphers", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_ssh_ciphers", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_ssh_ciphers", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_ssh_ciphers", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_ssh_ciphers", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_ssh_ciphers", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_ssh_ciphers", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_ssh_ciphers", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_ssh_ciphers", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_update_ssh_ciphers", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_ssh_ciphers", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *UpdateSshCiphersAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
