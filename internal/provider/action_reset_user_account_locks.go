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
var _ action.Action = (*ResetUserAccountLocksAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*ResetUserAccountLocksAction)(nil)

// ResetUserAccountLocksAction is the generated Terraform action implementation.
type ResetUserAccountLocksAction struct {
	client *client.Client
}

// ResetUserAccountLocksActionModel describes the action configuration shape.
type ResetUserAccountLocksActionModel struct {
	ClearHistory types.Bool   `tfsdk:"clear_history" json:"clearHistory"`
	ClusterId    types.String `tfsdk:"cluster_id"`
	Unlock       types.Bool   `tfsdk:"unlock"`
	Username     types.String `tfsdk:"username"`
}

// NewResetUserAccountLocksAction returns a new instance of the generated action.
func NewResetUserAccountLocksAction() action.Action {
	return &ResetUserAccountLocksAction{}
}

// Metadata returns the action type name.
func (r *ResetUserAccountLocksAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_reset_user_account_locks"
}

// Schema returns the action schema.
func (r *ResetUserAccountLocksAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Clears history of login failures, and/or unlocks user account", Attributes: map[string]schema.Attribute{"clear_history": schema.BoolAttribute{MarkdownDescription: "If set to 'false', leave the history alone and only unlock the account. Therefore, one more login will be permitted, but the account could then become re-locked after another failure (if it was already over the threshold)", Optional: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "unlock": schema.BoolAttribute{MarkdownDescription: "If set to 'false', clear the history, but leave the account's lock alone. Therefore, if it was locked, it remains locked until further action is taken", Optional: true}, "username": schema.StringAttribute{MarkdownDescription: "account to reset lockout settings for. If left out, all user accounts are affected", Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *ResetUserAccountLocksAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config ResetUserAccountLocksActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *ResetUserAccountLocksAction) invokeRemote(ctx context.Context, config *ResetUserAccountLocksActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/system/aaa/auth/accountLock"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_reset_user_account_locks", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_reset_user_account_locks", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_reset_user_account_locks", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_reset_user_account_locks", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 204) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_reset_user_account_locks", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_reset_user_account_locks", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_reset_user_account_locks", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_reset_user_account_locks", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_reset_user_account_locks", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_reset_user_account_locks", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_reset_user_account_locks", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_reset_user_account_locks", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_reset_user_account_locks", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *ResetUserAccountLocksAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
