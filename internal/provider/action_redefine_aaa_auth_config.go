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
var _ action.Action = (*RedefineAaaAuthConfigAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*RedefineAaaAuthConfigAction)(nil)

// RedefineAaaAuthConfigAction is the generated Terraform action implementation.
type RedefineAaaAuthConfigAction struct {
	client *client.Client
}

// RedefineAaaAuthConfigActionModel describes the action configuration shape.
type RedefineAaaAuthConfigActionModel struct {
	AuthSequence         types.Set     `tfsdk:"auth_sequence" json:"authSequence"`
	ClusterId            types.String  `tfsdk:"cluster_id"`
	ExternalLoginMapping types.Dynamic `tfsdk:"external_login_mapping" json:"externalLoginMapping"`
	NonLocalUsers        types.Dynamic `tfsdk:"non_local_users" json:"nonLocalUsers"`
	PasswordExpiration   types.Dynamic `tfsdk:"password_expiration" json:"passwordExpiration"`
	SshdMaxSessions      types.Int64   `tfsdk:"sshd_max_sessions" json:"sshdMaxSessions"`
	UserLockout          types.Dynamic `tfsdk:"user_lockout" json:"userLockout"`
}

// NewRedefineAaaAuthConfigAction returns a new instance of the generated action.
func NewRedefineAaaAuthConfigAction() action.Action {
	return &RedefineAaaAuthConfigAction{}
}

// Metadata returns the action type name.
func (r *RedefineAaaAuthConfigAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_redefine_aaa_auth_config"
}

// Schema returns the action schema.
func (r *RedefineAaaAuthConfigAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Redefine Node Authentication config", Attributes: map[string]schema.Attribute{"auth_sequence": schema.SetAttribute{MarkdownDescription: "Authentication methods order for user login. Valid values are 'local', 'ldap', 'radius', 'tacacs'", Required: true, ElementType: types.StringType}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "external_login_mapping": schema.DynamicAttribute{MarkdownDescription: "Specifies how externally authenticated logins (RADIUS, TACACS+, or LDAP) are mapped to local accounts. Private class", Optional: true}, "non_local_users": schema.DynamicAttribute{MarkdownDescription: "Settings for treating usernames that are not recognized as real accounts (not a locally configured account). Private class", Optional: true}, "password_expiration": schema.DynamicAttribute{MarkdownDescription: "Node Password Expiration config. Private class", Optional: true}, "sshd_max_sessions": schema.Int64Attribute{MarkdownDescription: "Maximum concurrent session that can be logged into devices", Optional: true}, "user_lockout": schema.DynamicAttribute{MarkdownDescription: "Node AAA user lockout settings. Private class", Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *RedefineAaaAuthConfigAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config RedefineAaaAuthConfigActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *RedefineAaaAuthConfigAction) invokeRemote(ctx context.Context, config *RedefineAaaAuthConfigActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/system/aaa/auth"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_aaa_auth_config", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_aaa_auth_config", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_aaa_auth_config", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_aaa_auth_config", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_aaa_auth_config", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_aaa_auth_config", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_aaa_auth_config", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_aaa_auth_config", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_aaa_auth_config", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_aaa_auth_config", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_aaa_auth_config", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_aaa_auth_config", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_aaa_auth_config", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *RedefineAaaAuthConfigAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
