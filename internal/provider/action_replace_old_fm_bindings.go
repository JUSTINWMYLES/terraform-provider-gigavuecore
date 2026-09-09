package provider

import (
	"context"
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
var _ action.Action = (*ReplaceOldFmBindingsAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*ReplaceOldFmBindingsAction)(nil)

// ReplaceOldFmBindingsAction is the generated Terraform action implementation.
type ReplaceOldFmBindingsAction struct {
	client *client.Client
}

// ReplaceOldFmBindingsActionModel describes the action configuration shape.
type ReplaceOldFmBindingsActionModel struct {
	OldFmVmac types.String `tfsdk:"old_fm_vmac"`
}

// NewReplaceOldFmBindingsAction returns a new instance of the generated action.
func NewReplaceOldFmBindingsAction() action.Action {
	return &ReplaceOldFmBindingsAction{}
}

// Metadata returns the action type name.
func (r *ReplaceOldFmBindingsAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_replace_old_fm_bindings"
}

// Schema returns the action schema.
func (r *ReplaceOldFmBindingsAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Replaces old FM's bindings on a new FM", Attributes: map[string]schema.Attribute{"old_fm_vmac": schema.StringAttribute{MarkdownDescription: "The VMAC of old FM", Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *ReplaceOldFmBindingsAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config ReplaceOldFmBindingsActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *ReplaceOldFmBindingsAction) invokeRemote(ctx context.Context, config *ReplaceOldFmBindingsActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/licensing/activation/bindings/replaceOldFm"
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_replace_old_fm_bindings", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("oldFmVmac", config.OldFmVmac.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_replace_old_fm_bindings", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_replace_old_fm_bindings", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_replace_old_fm_bindings", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_replace_old_fm_bindings", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_replace_old_fm_bindings", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_replace_old_fm_bindings", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_replace_old_fm_bindings", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_replace_old_fm_bindings", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_replace_old_fm_bindings", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_replace_old_fm_bindings", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *ReplaceOldFmBindingsAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
