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
	_ datasource.DataSource              = (*GetManagedDevicesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetManagedDevicesDataSource)(nil)
)

// GetManagedDevicesDataSource is the generated Terraform data source implementation.
type GetManagedDevicesDataSource struct {
	client *client.Client
}

// GetManagedDevicesDataSourceModel describes the data source state shape.
type GetManagedDevicesDataSourceModel struct {
	Context types.Object `tfsdk:"context"`
	NodeId  types.String `tfsdk:"node_id" json:"nodeId"`
	Nodes   types.List   `tfsdk:"nodes"`
	Page    types.String `tfsdk:"page"`
	Sort    types.String `tfsdk:"sort"`
}

// NewGetManagedDevicesDataSource returns a new instance of the generated data source.
func NewGetManagedDevicesDataSource() datasource.DataSource {
	return &GetManagedDevicesDataSource{}
}

// Metadata returns the data source type name.
func (d *GetManagedDevicesDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_managed_devices"
}

// Schema returns the data source schema.
func (d *GetManagedDevicesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "get flat/ungroupped managed device list", Attributes: map[string]schema.Attribute{"context": schema.SingleNestedAttribute{MarkdownDescription: "Gigamon query result context", Computed: true, Attributes: map[string]schema.Attribute{"page_no": schema.Int64Attribute{MarkdownDescription: "page number of the returned result set", Computed: true}, "page_size": schema.Int64Attribute{MarkdownDescription: "page size of the returned result set", Computed: true}, "sort": schema.ListAttribute{MarkdownDescription: "sorting info of the returned result set. list of fields in the array indicate sorting order", Computed: true, ElementType: types.StringType}, "total_items": schema.Int64Attribute{MarkdownDescription: "total number of items in the queried entity type", Computed: true}}}, "node_id": schema.StringAttribute{MarkdownDescription: "if provided, only requested node is returned", Optional: true}, "nodes": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"box_id": schema.StringAttribute{Computed: true}, "chassis_oper_status": schema.StringAttribute{MarkdownDescription: "Describes operational state available in Chassis information. 'up': node is up; 'down': node is down; 'left': node is left from the cluster; 'not-reachable': node is not reachable from the cluster' . Applicable only for H-series devices", Computed: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "applicable only for clustered H-series devices", Computed: true}, "cluster_leader": schema.StringAttribute{MarkdownDescription: "address of stack/cluster leader device", Computed: true}, "cluster_master": schema.StringAttribute{MarkdownDescription: "address of stack/cluster master device. (deprecated: use clusterLeader)", Computed: true}, "cluster_mode": schema.StringAttribute{MarkdownDescription: "standalone: not a member of any cluster/stack; master: cluster/stack master; standby: for H-series only; slave: non-master member of an H-series cluster or a slave in a G-series stack; vfex: visibility fabric extender of an H-Series cluster. (deprecated: use clusterModeAlias)", Computed: true}, "cluster_mode_alias": schema.StringAttribute{MarkdownDescription: "standalone: not a member of any cluster/stack; leader: cluster/stack leader; standby: for H-series only; member: non-leader member of an H-series cluster or a member in a G-series stack; vfex: visibility fabric extender of an H-Series cluster", Computed: true}, "cluster_state": schema.StringAttribute{MarkdownDescription: "Describes the cluster membership state. Applicable only for clustered H-series devices", Computed: true}, "cluster_vip": schema.StringAttribute{MarkdownDescription: "applicable only for clustered H-series devices", Computed: true}, "device_id": schema.StringAttribute{MarkdownDescription: "unique ID representing this managed Gigamon device", Computed: true}, "device_ip": schema.StringAttribute{MarkdownDescription: "device management IP Address", Computed: true}, "device_ips": schema.ListAttribute{MarkdownDescription: "list of all the IPs on the device", Computed: true, ElementType: types.StringType}, "disc_outcome": schema.StringAttribute{MarkdownDescription: "status of FM-to-device communication channel", Computed: true}, "disconnected": schema.BoolAttribute{MarkdownDescription: "Status of connection between FM and Device", Computed: true}, "dns_name": schema.StringAttribute{MarkdownDescription: "device DNS name, as registered in DNS server", Computed: true}, "failure_desc": schema.StringAttribute{MarkdownDescription: "reason for device discovery failure", Computed: true}, "family": schema.StringAttribute{Computed: true}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "health_state_reasons": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Computed: true}}}}, "hostname": schema.StringAttribute{MarkdownDescription: "device configured hostname", Computed: true}, "leader_pref": schema.Int64Attribute{MarkdownDescription: "leader election preference rank", Computed: true}, "licensed": schema.BoolAttribute{Computed: true}, "master_pref": schema.Int64Attribute{MarkdownDescription: "master election preference rank. (deprecated: use leaderPref)", Computed: true}, "model": schema.StringAttribute{MarkdownDescription: "Gigamon physical device models", Computed: true}, "nat_ip": schema.StringAttribute{MarkdownDescription: "Device NAT IP Address. Applicable for only devices behind NAT set up", Computed: true}, "oper_status": schema.StringAttribute{MarkdownDescription: "Describes device operational state. 'operational': fully configurable; 'safe': no configuration possible; 'limited': limited configuration possible. Applicable only for H-series devices", Computed: true}, "sw_build_number": schema.StringAttribute{Computed: true}, "sw_version": schema.StringAttribute{Computed: true}, "topo_node_id": schema.StringAttribute{MarkdownDescription: "unique topology node ID. Generated by FM server", Computed: true}}}}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetManagedDevicesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetManagedDevicesDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	d.readRemote(ctx, &config, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	if config.Page.IsNull() {
		resp.Diagnostics.AddWarning("Single-page result", fmt.Sprintf("This data source reads a single page of a paginated API endpoint and does not aggregate results across pages. The \"page\" argument is unset, so the default page is returned; set it to retrieve a different page."))
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}

// readRemote performs the read HTTP exchange and decodes the response into config. Extracted from Read so the request/response logic is unit-testable without a tfsdk.Config.
func (d *GetManagedDevicesDataSource) readRemote(ctx context.Context, config *GetManagedDevicesDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/nodes/flat"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_managed_devices", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.NodeId.IsNull() {
		query.Set("nodeId", config.NodeId.ValueString())
	}
	if !config.Page.IsNull() {
		query.Set("page", config.Page.ValueString())
	}
	if !config.Sort.IsNull() {
		query.Set("sort", config.Sort.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_managed_devices", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_managed_devices", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_managed_devices", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_managed_devices", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_managed_devices", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_managed_devices", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_managed_devices", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_managed_devices", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_managed_devices", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_managed_devices", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_managed_devices", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_managed_devices", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_managed_devices", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetManagedDevicesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
