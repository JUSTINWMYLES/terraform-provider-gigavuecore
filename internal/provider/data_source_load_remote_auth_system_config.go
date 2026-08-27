package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*LoadRemoteAuthSystemConfigDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadRemoteAuthSystemConfigDataSource)(nil)
)

// LoadRemoteAuthSystemConfigDataSource is the generated Terraform data source implementation.
type LoadRemoteAuthSystemConfigDataSource struct {
	client *client.Client
}

// LoadRemoteAuthSystemConfigDataSourceModel describes the data source state shape.
type LoadRemoteAuthSystemConfigDataSourceModel struct {
	ClusterId    types.String `tfsdk:"cluster_id" json:"clusterId"`
	LdapConfig   types.Object `tfsdk:"ldap_config" json:"ldapConfig"`
	RadiusConfig types.Object `tfsdk:"radius_config" json:"radiusConfig"`
	TacacsConfig types.Object `tfsdk:"tacacs_config" json:"tacacsConfig"`
}

// NewLoadRemoteAuthSystemConfigDataSource returns a new instance of the generated data source.
func NewLoadRemoteAuthSystemConfigDataSource() datasource.DataSource {
	return &LoadRemoteAuthSystemConfigDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadRemoteAuthSystemConfigDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_remote_auth_system_config"
}

// Schema returns the data source schema.
func (d *LoadRemoteAuthSystemConfigDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load Remote Auth servers system configuration", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "ldap_config": schema.SingleNestedAttribute{MarkdownDescription: "System-level LDAP config", Computed: true, Attributes: map[string]schema.Attribute{"accept_user_roles": schema.BoolAttribute{MarkdownDescription: "Enables the GigaVUE H Series node to accept user roles assigned in the LDAP server", Computed: true}, "map_enable": schema.BoolAttribute{MarkdownDescription: "Enables mapping of remote base DN to local account", Computed: true}, "referrals": schema.BoolAttribute{MarkdownDescription: "Toggles support for LDAP referrals", Computed: true}, "remote_map_table": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"local_account_name": schema.StringAttribute{MarkdownDescription: "Specifies local account to which remote base-dn needs to be mapped.", Computed: true}, "remote_base_dn": schema.StringAttribute{MarkdownDescription: "Specifies the base-dn of the remote user group to be mapped to local account", Computed: true}}}}, "server_config_defaults": schema.SingleNestedAttribute{MarkdownDescription: "Remote LDAP Server default config", Computed: true, Attributes: map[string]schema.Attribute{"base_dn": schema.StringAttribute{MarkdownDescription: "Identifies the base distinguished name (location) of the user information in the LDAP server's schema", Computed: true}, "bind_dn": schema.StringAttribute{MarkdownDescription: "Specifies the distinguished name (dn) on the LDAP server with which to bind. By default, this is left empty for anonymous login", Computed: true}, "bind_pwd": schema.StringAttribute{MarkdownDescription: "Provides the credentials to be used for binding with the LDAP server. If bind-dn is left undefined for anonymous login (the default), bind-password should be left undefined, too", Computed: true}, "bind_timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Computed: true}, "group_attribute": schema.StringAttribute{MarkdownDescription: "Use this argument to specify the name of the attribute to check for group membership. If you specify a value for group-dn, the attribute you name here will be checked to see whether it contains the user's distinguished name as one of the values in the LDAP server", Computed: true}, "group_dn": schema.StringAttribute{MarkdownDescription: "You can use this option to require membership in the named group-dn for successful login to the H Series node", Computed: true}, "login_attribute": schema.StringAttribute{MarkdownDescription: "Specify the name of the LDAP attribute containing the login name", Computed: true}, "port": schema.Int64Attribute{Computed: true}, "search_scope": schema.StringAttribute{MarkdownDescription: "Specifies the search scope for the user under the base distinguished name", Computed: true}, "search_timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Computed: true}, "ssl": schema.SingleNestedAttribute{MarkdownDescription: "Remote LDAP Server SSL config", Computed: true, Attributes: map[string]schema.Attribute{"ca_list": schema.StringAttribute{MarkdownDescription: "LDAP to use a supplemental CA list.", Computed: true}, "cert_verify": schema.BoolAttribute{MarkdownDescription: "Enable LDAP SSL/TLS certificate verification", Computed: true}, "mode": schema.StringAttribute{MarkdownDescription: "LDAP either in SSL or TLS", Computed: true}, "server_port": schema.Int64Attribute{MarkdownDescription: "LDAP SSL port number", Computed: true}}}, "version": schema.StringAttribute{Computed: true}}}}}, "radius_config": schema.SingleNestedAttribute{MarkdownDescription: "System-level RADIUS config", Computed: true, Attributes: map[string]schema.Attribute{"accept_user_roles": schema.BoolAttribute{MarkdownDescription: "Enables the GigaVUE H Series node to accept user roles assigned in the RADIUS server", Computed: true}, "server_config_defaults": schema.SingleNestedAttribute{MarkdownDescription: "Remote RADIUS Server default config", Computed: true, Attributes: map[string]schema.Attribute{"retries": schema.Int64Attribute{MarkdownDescription: "value of 0 disables retries", Computed: true}, "secret_key": schema.StringAttribute{Computed: true, Sensitive: true}, "timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Computed: true}}}}}, "tacacs_config": schema.SingleNestedAttribute{MarkdownDescription: "System-level TACACS+ config", Computed: true, Attributes: map[string]schema.Attribute{"accept_user_roles": schema.BoolAttribute{MarkdownDescription: "Enables the GigaVUE H Series node to accept user roles assigned in the TACACS+ server", Computed: true}, "server_config_defaults": schema.SingleNestedAttribute{MarkdownDescription: "Remote TACACS+ Server default config", Computed: true, Attributes: map[string]schema.Attribute{"retries": schema.Int64Attribute{MarkdownDescription: "value of 0 disables retries", Computed: true}, "secret_key": schema.StringAttribute{Computed: true, Sensitive: true}, "service": schema.StringAttribute{MarkdownDescription: "Specify which authorization service will be used for TACACS", Computed: true}, "timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Computed: true}}}}}}}
}

// Read fetches remote state into the data source model.
func (d *LoadRemoteAuthSystemConfigDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadRemoteAuthSystemConfigDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	d.readRemote(ctx, &config, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}

// readRemote performs the read HTTP exchange and decodes the response into config. Extracted from Read so the request/response logic is unit-testable without a tfsdk.Config.
func (d *LoadRemoteAuthSystemConfigDataSource) readRemote(ctx context.Context, config *LoadRemoteAuthSystemConfigDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/system/aaa/auth/servers"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_remote_auth_system_config", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_remote_auth_system_config", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_remote_auth_system_config", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_remote_auth_system_config", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_remote_auth_system_config", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_remote_auth_system_config", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_remote_auth_system_config", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_remote_auth_system_config", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_remote_auth_system_config", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_remote_auth_system_config", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_load_remote_auth_system_config", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_load_remote_auth_system_config", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_remote_auth_system_config", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_remote_auth_system_config", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadRemoteAuthSystemConfigDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Data Source Configure Type", fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData))
		return
	}
	d.client = c
}
