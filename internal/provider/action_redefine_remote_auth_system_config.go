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
var _ action.Action = (*RedefineRemoteAuthSystemConfigAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*RedefineRemoteAuthSystemConfigAction)(nil)

// RedefineRemoteAuthSystemConfigAction is the generated Terraform action implementation.
type RedefineRemoteAuthSystemConfigAction struct {
	client *client.Client
}

// RedefineRemoteAuthSystemConfigActionModel describes the action configuration shape.
type RedefineRemoteAuthSystemConfigActionModel struct {
	ClusterId    types.String `tfsdk:"cluster_id"`
	LdapConfig   types.Object `tfsdk:"ldap_config" json:"ldapConfig"`
	RadiusConfig types.Object `tfsdk:"radius_config" json:"radiusConfig"`
	TacacsConfig types.Object `tfsdk:"tacacs_config" json:"tacacsConfig"`
}

// NewRedefineRemoteAuthSystemConfigAction returns a new instance of the generated action.
func NewRedefineRemoteAuthSystemConfigAction() action.Action {
	return &RedefineRemoteAuthSystemConfigAction{}
}

// Metadata returns the action type name.
func (r *RedefineRemoteAuthSystemConfigAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_redefine_remote_auth_system_config"
}

// Schema returns the action schema.
func (r *RedefineRemoteAuthSystemConfigAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Redefine Remote Auth servers system configuration", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "ldap_config": schema.SingleNestedAttribute{MarkdownDescription: "System-level LDAP config", Optional: true, Attributes: map[string]schema.Attribute{"accept_user_roles": schema.BoolAttribute{MarkdownDescription: "Enables the GigaVUE H Series node to accept user roles assigned in the LDAP server", Optional: true}, "map_enable": schema.BoolAttribute{MarkdownDescription: "Enables mapping of remote base DN to local account", Optional: true}, "referrals": schema.BoolAttribute{MarkdownDescription: "Toggles support for LDAP referrals", Optional: true}, "remote_map_table": schema.ListNestedAttribute{Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"local_account_name": schema.StringAttribute{MarkdownDescription: "Specifies local account to which remote base-dn needs to be mapped.", Optional: true}, "remote_base_dn": schema.StringAttribute{MarkdownDescription: "Specifies the base-dn of the remote user group to be mapped to local account", Optional: true}}}}, "server_config_defaults": schema.SingleNestedAttribute{MarkdownDescription: "Remote LDAP Server default config", Optional: true, Attributes: map[string]schema.Attribute{"base_dn": schema.StringAttribute{MarkdownDescription: "Identifies the base distinguished name (location) of the user information in the LDAP server's schema", Optional: true}, "bind_dn": schema.StringAttribute{MarkdownDescription: "Specifies the distinguished name (dn) on the LDAP server with which to bind. By default, this is left empty for anonymous login", Optional: true}, "bind_pwd": schema.StringAttribute{MarkdownDescription: "Provides the credentials to be used for binding with the LDAP server. If bind-dn is left undefined for anonymous login (the default), bind-password should be left undefined, too", Optional: true}, "bind_timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Optional: true}, "group_attribute": schema.StringAttribute{MarkdownDescription: "Use this argument to specify the name of the attribute to check for group membership. If you specify a value for group-dn, the attribute you name here will be checked to see whether it contains the user's distinguished name as one of the values in the LDAP server", Optional: true}, "group_dn": schema.StringAttribute{MarkdownDescription: "You can use this option to require membership in the named group-dn for successful login to the H Series node", Optional: true}, "login_attribute": schema.StringAttribute{MarkdownDescription: "Specify the name of the LDAP attribute containing the login name", Optional: true}, "port": schema.Int64Attribute{Optional: true}, "search_scope": schema.StringAttribute{MarkdownDescription: "Specifies the search scope for the user under the base distinguished name", Optional: true}, "search_timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Optional: true}, "ssl": schema.SingleNestedAttribute{MarkdownDescription: "Remote LDAP Server SSL config", Optional: true, Attributes: map[string]schema.Attribute{"ca_list": schema.StringAttribute{MarkdownDescription: "LDAP to use a supplemental CA list.", Optional: true}, "cert_verify": schema.BoolAttribute{MarkdownDescription: "Enable LDAP SSL/TLS certificate verification", Optional: true}, "mode": schema.StringAttribute{MarkdownDescription: "LDAP either in SSL or TLS", Optional: true}, "server_port": schema.Int64Attribute{MarkdownDescription: "LDAP SSL port number", Optional: true}}}, "version": schema.StringAttribute{Optional: true}}}}}, "radius_config": schema.SingleNestedAttribute{MarkdownDescription: "System-level RADIUS config", Optional: true, Attributes: map[string]schema.Attribute{"accept_user_roles": schema.BoolAttribute{MarkdownDescription: "Enables the GigaVUE H Series node to accept user roles assigned in the RADIUS server", Optional: true}, "server_config_defaults": schema.SingleNestedAttribute{MarkdownDescription: "Remote RADIUS Server default config", Optional: true, Attributes: map[string]schema.Attribute{"retries": schema.Int64Attribute{MarkdownDescription: "value of 0 disables retries", Optional: true}, "secret_key": schema.StringAttribute{Optional: true}, "timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Optional: true}}}}}, "tacacs_config": schema.SingleNestedAttribute{MarkdownDescription: "System-level TACACS+ config", Optional: true, Attributes: map[string]schema.Attribute{"accept_user_roles": schema.BoolAttribute{MarkdownDescription: "Enables the GigaVUE H Series node to accept user roles assigned in the TACACS+ server", Optional: true}, "server_config_defaults": schema.SingleNestedAttribute{MarkdownDescription: "Remote TACACS+ Server default config", Optional: true, Attributes: map[string]schema.Attribute{"retries": schema.Int64Attribute{MarkdownDescription: "value of 0 disables retries", Optional: true}, "secret_key": schema.StringAttribute{Optional: true}, "service": schema.StringAttribute{MarkdownDescription: "Specify which authorization service will be used for TACACS", Optional: true}, "timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Optional: true}}}}}}}
}

// Invoke executes the action against the remote API.
func (r *RedefineRemoteAuthSystemConfigAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config RedefineRemoteAuthSystemConfigActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *RedefineRemoteAuthSystemConfigAction) invokeRemote(ctx context.Context, config *RedefineRemoteAuthSystemConfigActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/system/aaa/auth/servers"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_remote_auth_system_config", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_remote_auth_system_config", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_remote_auth_system_config", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_remote_auth_system_config", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_remote_auth_system_config", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_remote_auth_system_config", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_remote_auth_system_config", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_remote_auth_system_config", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_remote_auth_system_config", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_remote_auth_system_config", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_remote_auth_system_config", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_remote_auth_system_config", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_remote_auth_system_config", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *RedefineRemoteAuthSystemConfigAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
