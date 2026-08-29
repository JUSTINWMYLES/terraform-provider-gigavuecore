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
	_ datasource.DataSource              = (*LoadSnmpServerConfigDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadSnmpServerConfigDataSource)(nil)
)

// LoadSnmpServerConfigDataSource is the generated Terraform data source implementation.
type LoadSnmpServerConfigDataSource struct {
	client *client.Client
}

// LoadSnmpServerConfigDataSourceModel describes the data source state shape.
type LoadSnmpServerConfigDataSourceModel struct {
	ClusterId          types.String `tfsdk:"cluster_id" json:"clusterId"`
	CommunityConfig    types.Object `tfsdk:"community_config" json:"communityConfig"`
	NotifyConfig       types.Object `tfsdk:"notify_config" json:"notifyConfig"`
	SnmpThrottleConfig types.Object `tfsdk:"snmp_throttle_config" json:"snmpThrottleConfig"`
	SnmpV3Config       types.Object `tfsdk:"snmp_v3_config" json:"snmpV3Config"`
	SystemConfig       types.Object `tfsdk:"system_config" json:"systemConfig"`
}

// NewLoadSnmpServerConfigDataSource returns a new instance of the generated data source.
func NewLoadSnmpServerConfigDataSource() datasource.DataSource {
	return &LoadSnmpServerConfigDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadSnmpServerConfigDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_snmp_server_config"
}

// Schema returns the data source schema.
func (d *LoadSnmpServerConfigDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load Snmp Server config", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "community_config": schema.SingleNestedAttribute{MarkdownDescription: "SNMP Server Community Strings config", Computed: true, Attributes: map[string]schema.Attribute{"community_strings": schema.ListAttribute{MarkdownDescription: "community string[s] used to connect to this node using SNMP. The default value is 'public'. If 'enableMultiCommunity' is enabled, multiple community strings for the node are allowed", Computed: true, ElementType: types.StringType}, "enable_community_auth": schema.BoolAttribute{MarkdownDescription: "turn on community-based authentication for the system", Computed: true}, "enable_community_auth_v1": schema.BoolAttribute{MarkdownDescription: "turn on community-based authentication for the SNMP v1", Computed: true}, "enable_multi_community": schema.BoolAttribute{MarkdownDescription: "allow configuration of multiple communities", Computed: true}}}, "notify_config": schema.SingleNestedAttribute{MarkdownDescription: "Node SNMP Server Notification settings", Computed: true, Attributes: map[string]schema.Attribute{"notify_events": schema.SetAttribute{MarkdownDescription: "The set of notification event types", Computed: true, ElementType: types.StringType}, "notify_set": schema.StringAttribute{MarkdownDescription: "When 'notifySet' is 'select', notifyEvents represents the subset of notification event types activated for dispatch. When set to 'none', effectively disables SNMP notifications from the node", Computed: true}, "targets": schema.ListNestedAttribute{MarkdownDescription: "list of notification destinations", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{MarkdownDescription: "temporarily enable/disable the notification destination", Computed: true}, "host": schema.StringAttribute{MarkdownDescription: "ipv4 or ipv6 or domain name", Computed: true}, "notify_config": schema.SingleNestedAttribute{MarkdownDescription: "Notification Target configuration for specific Notification type (Trap/Inform)", Computed: true, Attributes: map[string]schema.Attribute{"auth_key": schema.StringAttribute{MarkdownDescription: "authentication password. required with 'v3user'", Computed: true}, "auth_protocol": schema.StringAttribute{MarkdownDescription: "authentication hash algorithm. required with 'v3user'", Computed: true}, "community": schema.StringAttribute{MarkdownDescription: "required when when 'version' is 'v2c'", Computed: true}, "engine_id": schema.StringAttribute{MarkdownDescription: "remote engineID. only valid with notifyType 'inform' and 'version' v3", Computed: true}, "port": schema.Int64Attribute{Computed: true}, "priv_key": schema.StringAttribute{MarkdownDescription: "privacy password", Computed: true}, "priv_protocol": schema.StringAttribute{MarkdownDescription: "privacy encryption", Computed: true}, "v3_user": schema.StringAttribute{MarkdownDescription: "required when when 'version' is 'v3'", Computed: true}, "version": schema.StringAttribute{MarkdownDescription: "SNMP version to use. v1 is only valid for traps. for v3, user name should be provided", Computed: true}}}, "notify_type": schema.StringAttribute{MarkdownDescription: "SNMP notification type to use", Computed: true}}}}}}, "snmp_throttle_config": schema.SingleNestedAttribute{MarkdownDescription: "SNMP Throttle Configuration", Computed: true, Attributes: map[string]schema.Attribute{"throttle_config_details": schema.ListNestedAttribute{MarkdownDescription: "list of SNMP Throttle Config Details on the node", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"interval": schema.Int64Attribute{MarkdownDescription: "Time interval at which the throttle should occur (in seconds)", Computed: true}, "notify_set": schema.StringAttribute{MarkdownDescription: "When 'notifySet' is 'select', throttleEvents represents the subset of event types for SNMP throttling. When set to 'none', effectively disables SNMP throttle from the node.", Computed: true}, "report_threshold": schema.Int64Attribute{MarkdownDescription: "Minimum count threshold to send the throttle report", Computed: true}, "throttle_events": schema.SetAttribute{MarkdownDescription: "The set of notification event types", Computed: true, ElementType: types.StringType}}}}}}, "snmp_v3_config": schema.SingleNestedAttribute{MarkdownDescription: "SNMP Server v3 config", Computed: true, Attributes: map[string]schema.Attribute{"snmp_v3_users": schema.ListNestedAttribute{MarkdownDescription: "list of SNMPv3 users on the node", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"auth_key": schema.StringAttribute{Computed: true}, "auth_protocol": schema.StringAttribute{Computed: true}, "enabled": schema.BoolAttribute{Computed: true}, "priv_key": schema.StringAttribute{Computed: true}, "priv_protocol": schema.StringAttribute{Computed: true}, "read_only": schema.BoolAttribute{MarkdownDescription: "This property is introduced for GUI's purpose to restrict the edit/delete on the SNMPv3 user used by FM. If this property doesn't exist, consider it false. And this property is ignored if it exists in the create/update spec.", Computed: true}, "username": schema.StringAttribute{Computed: true}}}}}}, "system_config": schema.SingleNestedAttribute{MarkdownDescription: "SNMP Server System level config", Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{MarkdownDescription: "Enables SNMP Server on the node", Computed: true}, "engine_id": schema.StringAttribute{MarkdownDescription: "local EngineID", Computed: true}, "port": schema.Int64Attribute{Computed: true}, "sys_contact": schema.StringAttribute{MarkdownDescription: "MIB-II 'sysContact'", Computed: true}, "sys_descr": schema.StringAttribute{MarkdownDescription: "MIB-II 'sysDescr'", Computed: true}, "sys_location": schema.StringAttribute{MarkdownDescription: "MIB-II 'sysLocation'", Computed: true}, "sys_name": schema.StringAttribute{MarkdownDescription: "MIB-II 'sysName'", Computed: true}}}}}
}

// Read fetches remote state into the data source model.
func (d *LoadSnmpServerConfigDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadSnmpServerConfigDataSourceModel
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
func (d *LoadSnmpServerConfigDataSource) readRemote(ctx context.Context, config *LoadSnmpServerConfigDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/system/snmp"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_snmp_server_config", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_snmp_server_config", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_snmp_server_config", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_snmp_server_config", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_snmp_server_config", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_snmp_server_config", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_snmp_server_config", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_snmp_server_config", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_snmp_server_config", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_snmp_server_config", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_load_snmp_server_config", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_load_snmp_server_config", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_snmp_server_config", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_snmp_server_config", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadSnmpServerConfigDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
