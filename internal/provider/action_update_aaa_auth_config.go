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
var _ action.Action = (*UpdateAaaAuthConfigAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*UpdateAaaAuthConfigAction)(nil)

// UpdateAaaAuthConfigAction is the generated Terraform action implementation.
type UpdateAaaAuthConfigAction struct {
	client *client.Client
}

// UpdateAaaAuthConfigActionModel describes the action configuration shape.
type UpdateAaaAuthConfigActionModel struct {
	AuthSequence         types.Set    `tfsdk:"auth_sequence" json:"authSequence"`
	ClusterId            types.String `tfsdk:"cluster_id"`
	ExternalLoginMapping types.Object `tfsdk:"external_login_mapping" json:"externalLoginMapping"`
	NonLocalUsers        types.Object `tfsdk:"non_local_users" json:"nonLocalUsers"`
	PasswordExpiration   types.Object `tfsdk:"password_expiration" json:"passwordExpiration"`
	SshdMaxSessions      types.Int64  `tfsdk:"sshd_max_sessions" json:"sshdMaxSessions"`
	UserLockout          types.Object `tfsdk:"user_lockout" json:"userLockout"`
}

// NewUpdateAaaAuthConfigAction returns a new instance of the generated action.
func NewUpdateAaaAuthConfigAction() action.Action {
	return &UpdateAaaAuthConfigAction{}
}

// Metadata returns the action type name.
func (r *UpdateAaaAuthConfigAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_update_aaa_auth_config"
}

// Schema returns the action schema.
func (r *UpdateAaaAuthConfigAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Update Node Authentication config", Attributes: map[string]schema.Attribute{"auth_sequence": schema.SetAttribute{MarkdownDescription: "Authentication methods order for user login. Valid values are 'local', 'ldap', 'radius', 'tacacs'", Required: true, ElementType: types.StringType}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "external_login_mapping": schema.SingleNestedAttribute{MarkdownDescription: "Specifies how externally authenticated logins (RADIUS, TACACS+, or LDAP) are mapped to local accounts. Private class", Optional: true, Attributes: map[string]schema.Attribute{"default_local_user": schema.StringAttribute{MarkdownDescription: "Specifies the account to which externally authenticated logins are mapped when map order is set to remote-first (if there is no matching local account) or local-only", Optional: true}, "user_map_order": schema.StringAttribute{MarkdownDescription: "Specifies how externally authenticated logins (RADIUS, TACACS+, or LDAP) are mapped to local accounts", Optional: true}}}, "non_local_users": schema.SingleNestedAttribute{MarkdownDescription: "Settings for treating usernames that are not recognized as real accounts (not a locally configured account). Private class", Optional: true, Attributes: map[string]schema.Attribute{"hash_username": schema.BoolAttribute{MarkdownDescription: "Apply a hash function to the username and store the hashed result for usernames that are not recognized as real accounts (not a locally configured account)", Optional: true}, "track_auth_failures": schema.BoolAttribute{MarkdownDescription: "Enables tracking authentication failures for usernames that are not recognized as real accounts (not a locally configured account)", Optional: true}}}, "password_expiration": schema.SingleNestedAttribute{MarkdownDescription: "Node Password Expiration config. Private class", Optional: true, Attributes: map[string]schema.Attribute{"duration": schema.Int64Attribute{MarkdownDescription: "the number of days password is valid", Optional: true}, "enabled": schema.BoolAttribute{MarkdownDescription: "Enables password expiration", Optional: true}}}, "sshd_max_sessions": schema.Int64Attribute{MarkdownDescription: "Maximum concurrent session that can be logged into devices", Optional: true}, "user_lockout": schema.SingleNestedAttribute{MarkdownDescription: "Node AAA user lockout settings. Private class", Optional: true, Attributes: map[string]schema.Attribute{"enable_admin_lockout": schema.BoolAttribute{MarkdownDescription: "Overrides the global settings for tracking and lockouts for the 'admin' account. Disabling means that the admin user will never be locked out, though their authentication failure history will still be tracked if tracking is enabled overall. This applies only to the single account with the username 'admin'. It does not apply to any other users with administrative privileges", Optional: true}, "enable_lockout": schema.BoolAttribute{MarkdownDescription: "Enables or disables locking out of user accounts based on authentication failures", Optional: true}, "lock_time": schema.Int64Attribute{MarkdownDescription: "In seconds. Specifies that no logins are permitted for this number of seconds following any login failure (not counting failures caused by the lockout mechanism, or the lockTime itself). This is not based on the number of consecutive failures. If both 'unlockTime' and 'lockTime' are set, the 'unlockTime' must be greater than the 'lockTime'", Optional: true}, "max_fail": schema.Int64Attribute{MarkdownDescription: "Sets the maximum number of consecutive authentication failures (attempts) permitted for a user account before the account is locked. After this number of failures, the account is locked and subsequent attempts are not permitted", Optional: true}, "track_auth_failures": schema.BoolAttribute{MarkdownDescription: "Enables or disables tracking of authentication failures. Can be used for informational purposes of for user account lockout", Optional: true}, "unlock_time": schema.Int64Attribute{MarkdownDescription: "In seconds. Specifies that if a user account is locked due to authentication failures, another login attempt will be permitted if this number of seconds has elapsed since the last login failure. That does not count failures caused by the lockout mechanism itself. A user must have been permitted to attempt to login, and then failed. After this interval has elapsed, the account does not become unlocked, nor does its history reset. It simply permits one more login attempt even if the account is locked", Optional: true}}}}}
}

// Invoke executes the action against the remote API.
func (r *UpdateAaaAuthConfigAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UpdateAaaAuthConfigActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *UpdateAaaAuthConfigAction) invokeRemote(ctx context.Context, config *UpdateAaaAuthConfigActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/system/aaa/auth"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_aaa_auth_config", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_aaa_auth_config", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPatch, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_aaa_auth_config", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_aaa_auth_config", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_aaa_auth_config", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_aaa_auth_config", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_aaa_auth_config", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_aaa_auth_config", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_aaa_auth_config", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_aaa_auth_config", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_aaa_auth_config", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_update_aaa_auth_config", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_aaa_auth_config", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *UpdateAaaAuthConfigAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
