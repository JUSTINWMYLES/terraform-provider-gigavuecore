package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*RedefineSecurityConfigAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*RedefineSecurityConfigAction)(nil)

// RedefineSecurityConfigAction is the generated Terraform action implementation.
type RedefineSecurityConfigAction struct {
	client *client.Client
}

// RedefineSecurityConfigActionModel describes the action configuration shape.
type RedefineSecurityConfigActionModel struct {
	AllowBlankPassword   types.Bool   `tfsdk:"allow_blank_password" json:"allowBlankPassword"`
	ClusterId            types.String `tfsdk:"cluster_id" json:"clusterId"`
	FipsEnabled          types.Bool   `tfsdk:"fips_enabled" json:"fipsEnabled"`
	FipsModeState        types.Bool   `tfsdk:"fips_mode_state" json:"fipsModeState"`
	MinPasswordLen       types.Int64  `tfsdk:"min_password_len" json:"minPasswordLen"`
	SecureCrypto         types.Bool   `tfsdk:"secure_crypto" json:"secureCrypto"`
	SecureCryptoEnforced types.Bool   `tfsdk:"secure_crypto_enforced" json:"secureCryptoEnforced"`
	SecurePasswords      types.Bool   `tfsdk:"secure_passwords" json:"securePasswords"`
}

// NewRedefineSecurityConfigAction returns a new instance of the generated action.
func NewRedefineSecurityConfigAction() action.Action {
	return &RedefineSecurityConfigAction{}
}

// Metadata returns the action type name.
func (r *RedefineSecurityConfigAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_redefine_security_config"
}

// Schema returns the action schema.
func (r *RedefineSecurityConfigAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Redefine Security config", Attributes: map[string]schema.Attribute{"allow_blank_password": schema.BoolAttribute{Optional: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "id of the defining cluster", Optional: true}, "fips_enabled": schema.BoolAttribute{MarkdownDescription: "enable/disable fips mode(pending fips mode), system reload is necessary to activate fips mode", Optional: true}, "fips_mode_state": schema.BoolAttribute{MarkdownDescription: "current fips mode", Optional: true}, "min_password_len": schema.Int64Attribute{Optional: true}, "secure_crypto": schema.BoolAttribute{MarkdownDescription: "Secure crypto mode. Value takes effect after reload", Optional: true}, "secure_crypto_enforced": schema.BoolAttribute{MarkdownDescription: "Secure crypto mode enforced", Optional: true}, "secure_passwords": schema.BoolAttribute{MarkdownDescription: "Secure passwords mode", Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *RedefineSecurityConfigAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config RedefineSecurityConfigActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *RedefineSecurityConfigAction) invokeRemote(ctx context.Context, config *RedefineSecurityConfigActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/system/security"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_security_config", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_security_config", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_security_config", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_security_config", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_security_config", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_security_config", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_security_config", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_security_config", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_security_config", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_security_config", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_security_config", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_security_config", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_security_config", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *RedefineSecurityConfigAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
