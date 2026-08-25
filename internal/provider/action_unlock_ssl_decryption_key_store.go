package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*UnlockSslDecryptionKeyStoreAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*UnlockSslDecryptionKeyStoreAction)(nil)

// UnlockSslDecryptionKeyStoreAction is the generated Terraform action implementation.
type UnlockSslDecryptionKeyStoreAction struct {
	client *client.Client
}

// UnlockSslDecryptionKeyStoreActionModel describes the action configuration shape.
type UnlockSslDecryptionKeyStoreActionModel struct {
	AutoLogin  types.Bool   `tfsdk:"auto_login"`
	ClusterId  types.String `tfsdk:"cluster_id"`
	KsPassword types.String `tfsdk:"ks_password" json:"ksPassword"`
}

// NewUnlockSslDecryptionKeyStoreAction returns a new instance of the generated action.
func NewUnlockSslDecryptionKeyStoreAction() action.Action {
	return &UnlockSslDecryptionKeyStoreAction{}
}

// Metadata returns the action type name.
func (r *UnlockSslDecryptionKeyStoreAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_unlock_ssl_decryption_key_store"
}

// Schema returns the action schema.
func (r *UnlockSslDecryptionKeyStoreAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Unlock keystore keychain", Attributes: map[string]schema.Attribute{"auto_login": schema.BoolAttribute{MarkdownDescription: "If enabled, On every node reboot FM would auto login into device with the password configured.", Optional: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "ks_password": schema.StringAttribute{MarkdownDescription: "Only valid for 'set' operations", Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *UnlockSslDecryptionKeyStoreAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UnlockSslDecryptionKeyStoreActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *UnlockSslDecryptionKeyStoreAction) invokeRemote(ctx context.Context, config *UnlockSslDecryptionKeyStoreActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/keystore/keychain/lock"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_unlock_ssl_decryption_key_store", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_unlock_ssl_decryption_key_store", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_unlock_ssl_decryption_key_store", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	if !config.AutoLogin.IsNull() {
		query.Set("autoLogin", strconv.FormatBool(config.AutoLogin.ValueBool()))
	}
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_unlock_ssl_decryption_key_store", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_unlock_ssl_decryption_key_store", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_unlock_ssl_decryption_key_store", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_unlock_ssl_decryption_key_store", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_unlock_ssl_decryption_key_store", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_unlock_ssl_decryption_key_store", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_unlock_ssl_decryption_key_store", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_unlock_ssl_decryption_key_store", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_unlock_ssl_decryption_key_store", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_unlock_ssl_decryption_key_store", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *UnlockSslDecryptionKeyStoreAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
