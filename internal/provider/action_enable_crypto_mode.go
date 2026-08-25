package provider

import (
	"context"
	"fmt"
	"net/http"
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*EnableCryptoModeAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*EnableCryptoModeAction)(nil)

// EnableCryptoModeAction is the generated Terraform action implementation.
type EnableCryptoModeAction struct {
	client *client.Client
}

// EnableCryptoModeActionModel describes the action configuration shape.
type EnableCryptoModeActionModel struct {
}

// NewEnableCryptoModeAction returns a new instance of the generated action.
func NewEnableCryptoModeAction() action.Action {
	return &EnableCryptoModeAction{}
}

// Metadata returns the action type name.
func (r *EnableCryptoModeAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_enable_crypto_mode"
}

// Schema returns the action schema.
func (r *EnableCryptoModeAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Enable crypto mode."}
}

// Invoke executes the action against the remote API.
func (r *EnableCryptoModeAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config EnableCryptoModeActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *EnableCryptoModeAction) invokeRemote(ctx context.Context, config *EnableCryptoModeActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/system/crypto/cryptoMode/enable"
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_enable_crypto_mode", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_enable_crypto_mode", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_enable_crypto_mode", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_enable_crypto_mode", "Access Denied. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_enable_crypto_mode", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_enable_crypto_mode", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_enable_crypto_mode", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_enable_crypto_mode", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_enable_crypto_mode", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *EnableCryptoModeAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
