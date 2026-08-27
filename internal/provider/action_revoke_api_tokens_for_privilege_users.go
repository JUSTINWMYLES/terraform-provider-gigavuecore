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
var _ action.Action = (*RevokeApiTokensForPrivilegeUsersAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*RevokeApiTokensForPrivilegeUsersAction)(nil)

// RevokeApiTokensForPrivilegeUsersAction is the generated Terraform action implementation.
type RevokeApiTokensForPrivilegeUsersAction struct {
	client *client.Client
}

// RevokeApiTokensForPrivilegeUsersActionModel describes the action configuration shape.
type RevokeApiTokensForPrivilegeUsersActionModel struct {
	Context                types.Dynamic `tfsdk:"context"`
	FmApiTokenUserEntities types.Dynamic `tfsdk:"fm_api_token_user_entities" json:"fmApiTokenUserEntities"`
}

// NewRevokeApiTokensForPrivilegeUsersAction returns a new instance of the generated action.
func NewRevokeApiTokensForPrivilegeUsersAction() action.Action {
	return &RevokeApiTokensForPrivilegeUsersAction{}
}

// Metadata returns the action type name.
func (r *RevokeApiTokensForPrivilegeUsersAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_revoke_api_tokens_for_privilege_users"
}

// Schema returns the action schema.
func (r *RevokeApiTokensForPrivilegeUsersAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Revoke other user token of FM Users. Users with FM Security Management role with write access can access API.", Attributes: map[string]schema.Attribute{"context": schema.DynamicAttribute{MarkdownDescription: "Gigamon query result context", Optional: true}, "fm_api_token_user_entities": schema.DynamicAttribute{Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *RevokeApiTokensForPrivilegeUsersAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config RevokeApiTokensForPrivilegeUsersActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *RevokeApiTokensForPrivilegeUsersAction) invokeRemote(ctx context.Context, config *RevokeApiTokensForPrivilegeUsersActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/tokens/manage/revoke"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_revoke_api_tokens_for_privilege_users", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_revoke_api_tokens_for_privilege_users", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_revoke_api_tokens_for_privilege_users", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_revoke_api_tokens_for_privilege_users", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_revoke_api_tokens_for_privilege_users", "Invalid request. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_revoke_api_tokens_for_privilege_users", "Access Denied. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_revoke_api_tokens_for_privilege_users", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_revoke_api_tokens_for_privilege_users", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_revoke_api_tokens_for_privilege_users", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_revoke_api_tokens_for_privilege_users", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *RevokeApiTokensForPrivilegeUsersAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
