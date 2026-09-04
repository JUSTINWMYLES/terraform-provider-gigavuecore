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
var _ action.Action = (*FmServerAaaConfigAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*FmServerAaaConfigAction)(nil)

// FmServerAaaConfigAction is the generated Terraform action implementation.
type FmServerAaaConfigAction struct {
	client *client.Client
}

// FmServerAaaConfigActionModel describes the action configuration shape.
type FmServerAaaConfigActionModel struct {
	AccountingMethod     types.String `tfsdk:"accounting_method" json:"accountingMethod"`
	AuthMethod           types.String `tfsdk:"auth_method" json:"authMethod"`
	DefaultLoginAttempts types.Int64  `tfsdk:"default_login_attempts" json:"defaultLoginAttempts"`
	DefaultUserGroup     types.String `tfsdk:"default_user_group" json:"defaultUserGroup"`
	UserLockEnabled      types.Bool   `tfsdk:"user_lock_enabled" json:"userLockEnabled"`
}

// NewFmServerAaaConfigAction returns a new instance of the generated action.
func NewFmServerAaaConfigAction() action.Action {
	return &FmServerAaaConfigAction{}
}

// Metadata returns the action type name.
func (r *FmServerAaaConfigAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_fm_server_aaa_config"
}

// Schema returns the action schema.
func (r *FmServerAaaConfigAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "fmServer AAA Config", Attributes: map[string]schema.Attribute{"accounting_method": schema.StringAttribute{MarkdownDescription: "Accounting Method", Optional: true}, "auth_method": schema.StringAttribute{MarkdownDescription: "Auth Method", Required: true}, "default_login_attempts": schema.Int64Attribute{MarkdownDescription: "DefaultLoginAttempts will only be applicable for local auth type", Optional: true}, "default_user_group": schema.StringAttribute{MarkdownDescription: "Default UserGroup", Required: true}, "user_lock_enabled": schema.BoolAttribute{MarkdownDescription: "UserLockEnabled will only be applicable for local auth type", Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *FmServerAaaConfigAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config FmServerAaaConfigActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *FmServerAaaConfigAction) invokeRemote(ctx context.Context, config *FmServerAaaConfigActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/sys/aaa"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_fm_server_aaa_config", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_fm_server_aaa_config", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_fm_server_aaa_config", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_fm_server_aaa_config", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_fm_server_aaa_config", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_fm_server_aaa_config", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_fm_server_aaa_config", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_fm_server_aaa_config", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_fm_server_aaa_config", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_fm_server_aaa_config", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_fm_server_aaa_config", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_fm_server_aaa_config", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_fm_server_aaa_config", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *FmServerAaaConfigAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
