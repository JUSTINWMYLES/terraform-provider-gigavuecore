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
	_ datasource.DataSource              = (*GetClusterConfigStateDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetClusterConfigStateDataSource)(nil)
)

// GetClusterConfigStateDataSource is the generated Terraform data source implementation.
type GetClusterConfigStateDataSource struct {
	client *client.Client
}

// GetClusterConfigStateDataSourceModel describes the data source state shape.
type GetClusterConfigStateDataSourceModel struct {
	ClusterEnabled        types.Bool   `tfsdk:"cluster_enabled" json:"clusterEnabled"`
	ClusterId             types.String `tfsdk:"cluster_id" json:"clusterId"`
	ClusterLeaderBoxId    types.Int64  `tfsdk:"cluster_leader_box_id" json:"clusterLeaderBoxId"`
	ClusterMasterBoxId    types.Int64  `tfsdk:"cluster_master_box_id" json:"clusterMasterBoxId"`
	ClusterMembers        types.List   `tfsdk:"cluster_members" json:"clusterMembers"`
	ClusterParams         types.Object `tfsdk:"cluster_params" json:"clusterParams"`
	GlobalClusterId       types.String `tfsdk:"global_cluster_id" json:"globalClusterId"`
	LeaderAddress         types.String `tfsdk:"leader_address" json:"leaderAddress"`
	LocalClusterRole      types.String `tfsdk:"local_cluster_role" json:"localClusterRole"`
	LocalClusterRoleAlias types.String `tfsdk:"local_cluster_role_alias" json:"localClusterRoleAlias"`
	LocalclusterState     types.String `tfsdk:"localcluster_state" json:"localclusterState"`
	MasterAddress         types.String `tfsdk:"master_address" json:"masterAddress"`
}

// NewGetClusterConfigStateDataSource returns a new instance of the generated data source.
func NewGetClusterConfigStateDataSource() datasource.DataSource {
	return &GetClusterConfigStateDataSource{}
}

// Metadata returns the data source type name.
func (d *GetClusterConfigStateDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_cluster_config_state"
}

// Schema returns the data source schema.
func (d *GetClusterConfigStateDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "get the cluster configuration state", Attributes: map[string]schema.Attribute{"cluster_enabled": schema.BoolAttribute{MarkdownDescription: "cluster enable or disable", Computed: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "if provided, the state and the configurations of the requested cluster will be returned", Required: true}, "cluster_leader_box_id": schema.Int64Attribute{MarkdownDescription: "cluster leader box id", Computed: true}, "cluster_master_box_id": schema.Int64Attribute{MarkdownDescription: "cluster master box id. (deprecated: use clusterLeaderBoxId)", Computed: true}, "cluster_members": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"boot_time": schema.StringAttribute{MarkdownDescription: "the time node booted up. In ISO-8601 date format 'yyyy-MM-dd'T'HH:mm:ssZ'", Computed: true}, "box_id": schema.Int64Attribute{MarkdownDescription: "box id of the cluster member", Computed: true}, "cc_sync_status": schema.StringAttribute{MarkdownDescription: "CC1/CC2 dynamic sync status", Computed: true}, "cluster_intf": schema.StringAttribute{MarkdownDescription: "cluster service interface", Computed: true}, "cluster_name": schema.StringAttribute{MarkdownDescription: "cluster name", Computed: true}, "cluster_node_id": schema.Int64Attribute{MarkdownDescription: "cluster node id", Computed: true}, "external_address": schema.StringAttribute{MarkdownDescription: "node external address", Computed: true}, "global_node_id": schema.StringAttribute{MarkdownDescription: "global node id", Computed: true}, "host_id": schema.StringAttribute{MarkdownDescription: "host id", Computed: true}, "hostname": schema.StringAttribute{MarkdownDescription: "host name of the member", Computed: true}, "internal_address": schema.StringAttribute{MarkdownDescription: "node internal address", Computed: true}, "internal_port": schema.Int64Attribute{MarkdownDescription: "node internal port", Computed: true}, "leader_discovery": schema.SingleNestedAttribute{MarkdownDescription: "provides leader discovery information", Computed: true, Attributes: map[string]schema.Attribute{"auto_discovery": schema.BoolAttribute{MarkdownDescription: "Cluster auto-discovery. Nodes with auto-discovery disabled cannot become leader", Computed: true}, "cluster_formation_timeout": schema.Int64Attribute{MarkdownDescription: "maximum expected time for cluster startup in seconds", Computed: true}, "primary_ip": schema.StringAttribute{MarkdownDescription: "Leader primary ip address. Only valid if 'autoDiscovery' is disabled", Computed: true}, "primary_port": schema.Int64Attribute{MarkdownDescription: "Leader primary port. Only valid if 'autoDiscovery' is disabled", Computed: true}, "secondary_ip": schema.StringAttribute{MarkdownDescription: "Leader secondary ip address. Only valid if 'autoDiscovery' is disabled", Computed: true}, "secondary_port": schema.Int64Attribute{MarkdownDescription: "Leader secondary port. Only valid if 'autoDiscovery' is disabled", Computed: true}, "static_discovery_timeout": schema.Int64Attribute{MarkdownDescription: "Set the cluster leader connection timeout in seconds. Maps to CLI 'connectTimeout'. Controls how long 'primaryIp' should be tried before trying to connect to 'secondaryIp' and vice versa. Only valid if 'autoDiscovery' is disabled", Computed: true}}}, "leader_preference": schema.Int64Attribute{MarkdownDescription: "leader election preference rank", Computed: true}, "master_discovery": schema.SingleNestedAttribute{MarkdownDescription: "provides master discovery information. (deprecated: use leaderDiscovery)", Computed: true, Attributes: map[string]schema.Attribute{"auto_discovery": schema.BoolAttribute{MarkdownDescription: "Cluster auto-discovery. Nodes with auto-discovery disabled cannot become master", Computed: true}, "cluster_formation_timeout": schema.Int64Attribute{MarkdownDescription: "maximum expected time for cluster startup in seconds", Computed: true}, "primary_ip": schema.StringAttribute{MarkdownDescription: "Master primary ip address. Only valid if 'autoDiscovery' is disabled", Computed: true}, "primary_port": schema.Int64Attribute{MarkdownDescription: "Master primary port. Only valid if 'autoDiscovery' is disabled", Computed: true}, "secondary_ip": schema.StringAttribute{MarkdownDescription: "Master secondary ip address. Only valid if 'autoDiscovery' is disabled", Computed: true}, "secondary_port": schema.Int64Attribute{MarkdownDescription: "Master secondary port. Only valid if 'autoDiscovery' is disabled", Computed: true}, "static_discovery_timeout": schema.Int64Attribute{MarkdownDescription: "Set the cluster master connection timeout in seconds. Maps to CLI 'connectTimeout'. Controls how long 'primaryIp' should be tried before trying to connect to 'secondaryIp' and vice versa. Only valid if 'autoDiscovery' is disabled", Computed: true}}}, "master_preference": schema.Int64Attribute{MarkdownDescription: "master election preference rank. (deprecated: use leaderPreference)", Computed: true}, "mgmt_address": schema.StringAttribute{MarkdownDescription: "device management address", Computed: true}, "model": schema.StringAttribute{MarkdownDescription: "Device model", Computed: true}, "node_cluster_role": schema.StringAttribute{MarkdownDescription: "standalone: not a member of any cluster; master: cluster master; standby: stand-by master; normal: non-master member; unknown: error, cannot find master. (deprecated: use nodeClusterRoleAlias)", Computed: true}, "node_cluster_role_alias": schema.StringAttribute{MarkdownDescription: "standalone: not a member of any cluster; leader: cluster leader; standby: stand-by leader; normal: non-leader member; unknown: error, cannot find leader", Computed: true}, "node_cluster_state": schema.StringAttribute{MarkdownDescription: "Operational status of the cluster member", Computed: true}, "oper_status": schema.StringAttribute{MarkdownDescription: "'operational': fully configurable; 'safe': no configuration possible; 'limited': limited configuration possible", Computed: true}, "recv_hb_from": schema.Int64Attribute{MarkdownDescription: "node id from which heartbeats received", Computed: true}, "send_hb_to": schema.Int64Attribute{MarkdownDescription: "node id to which heartbeats sent", Computed: true}, "serial_number": schema.StringAttribute{MarkdownDescription: "unique serial number for node chassis", Computed: true}, "sw_build_number": schema.StringAttribute{MarkdownDescription: "Device software build number", Computed: true}, "sw_version": schema.StringAttribute{MarkdownDescription: "Device software version", Computed: true}, "vip_intf": schema.StringAttribute{MarkdownDescription: "cluster master interface", Computed: true}}}}, "cluster_params": schema.SingleNestedAttribute{MarkdownDescription: "provides clustering parameters", Computed: true, Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "cluster id", Computed: true}, "cluster_vip": schema.StringAttribute{MarkdownDescription: "cluster leader virtual ip address", Computed: true}, "cluster_vip_mask_len": schema.Int64Attribute{MarkdownDescription: "cluster leader virtual ip mask length, valid and required when 'clusterVip' is specified", Computed: true}, "ip_protocol": schema.StringAttribute{MarkdownDescription: "Supported since 6.2", Computed: true}, "shared_secret": schema.StringAttribute{MarkdownDescription: "shared-secret used for message authentication, length 16 to 64", Computed: true, Sensitive: true}, "stacking_mode": schema.StringAttribute{MarkdownDescription: "stacking mode", Computed: true}}}, "global_cluster_id": schema.StringAttribute{MarkdownDescription: "global cluster id", Computed: true}, "leader_address": schema.StringAttribute{MarkdownDescription: "cluster leader management address", Computed: true}, "local_cluster_role": schema.StringAttribute{MarkdownDescription: "standalone: not a member of any cluster; master: cluster master; standby: stand-by master; normal: non-master member; unknown: error, cannot find master. (deprecated: use localClusterRoleAlias)", Computed: true}, "local_cluster_role_alias": schema.StringAttribute{MarkdownDescription: "standalone: not a member of any cluster; leader: cluster leader; standby: stand-by leader; normal: non-leader member; unknown: error, cannot find leader", Computed: true}, "localcluster_state": schema.StringAttribute{MarkdownDescription: "operational status of the cluster member", Computed: true}, "master_address": schema.StringAttribute{MarkdownDescription: "cluster master management address. (deprecated: use leaderAddress)", Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetClusterConfigStateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetClusterConfigStateDataSourceModel
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
func (d *GetClusterConfigStateDataSource) readRemote(ctx context.Context, config *GetClusterConfigStateDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/clusterConfig/clusteringState"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_cluster_config_state", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_cluster_config_state", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_cluster_config_state", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_cluster_config_state", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_cluster_config_state", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_cluster_config_state", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_cluster_config_state", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_cluster_config_state", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_cluster_config_state", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_cluster_config_state", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_cluster_config_state", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_cluster_config_state", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_cluster_config_state", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_cluster_config_state", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetClusterConfigStateDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
