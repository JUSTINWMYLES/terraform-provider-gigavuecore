package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*GetManagedClustersDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetManagedClustersDataSource)(nil)
)

// GetManagedClustersDataSource is the generated Terraform data source implementation.
type GetManagedClustersDataSource struct {
	client *client.Client
}

// GetManagedClustersDataSourceModel describes the data source state shape.
type GetManagedClustersDataSourceModel struct {
	ClusterId types.String `tfsdk:"cluster_id" json:"clusterId"`
	Items     types.List   `tfsdk:"items"`
	Page      types.String `tfsdk:"page"`
	Sort      types.String `tfsdk:"sort"`
}

// NewGetManagedClustersDataSource returns a new instance of the generated data source.
func NewGetManagedClustersDataSource() datasource.DataSource {
	return &GetManagedClustersDataSource{}
}

// Metadata returns the data source type name.
func (d *GetManagedClustersDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_managed_clusters"
}

// Schema returns the data source schema.
func (d *GetManagedClustersDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "get cluster-grouped managed device list", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "if provided, only requested cluster is returned", Optional: true}, "items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "for H-series clusters only. Identifies the configured cluster id", Computed: true}, "cluster_vip": schema.StringAttribute{MarkdownDescription: "for H-series clusters only. Identifies the configured cluster VIP address", Computed: true}, "family": schema.StringAttribute{MarkdownDescription: "identifies whether this is an H-series cluster or a G-series stack", Computed: true}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "health_state_reasons": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Computed: true}}}}, "leader_id": schema.StringAttribute{MarkdownDescription: "id of stack's leader device", Computed: true}, "master_id": schema.StringAttribute{MarkdownDescription: "id of stack's master device. (deprecated: use leaderId)", Computed: true}, "members": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"box_id": schema.StringAttribute{Computed: true}, "chassis_oper_status": schema.StringAttribute{MarkdownDescription: "Describes operational state available in Chassis information. 'up': node is up; 'down': node is down; 'left': node is left from the cluster; 'not-reachable': node is not reachable from the cluster' . Applicable only for H-series devices", Computed: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "applicable only for clustered H-series devices", Computed: true}, "cluster_leader": schema.StringAttribute{MarkdownDescription: "address of stack/cluster leader device", Computed: true}, "cluster_master": schema.StringAttribute{MarkdownDescription: "address of stack/cluster master device. (deprecated: use clusterLeader)", Computed: true}, "cluster_mode": schema.StringAttribute{MarkdownDescription: "standalone: not a member of any cluster/stack; master: cluster/stack master; standby: for H-series only; slave: non-master member of an H-series cluster or a slave in a G-series stack; vfex: visibility fabric extender of an H-Series cluster. (deprecated: use clusterModeAlias)", Computed: true}, "cluster_mode_alias": schema.StringAttribute{MarkdownDescription: "standalone: not a member of any cluster/stack; leader: cluster/stack leader; standby: for H-series only; member: non-leader member of an H-series cluster or a member in a G-series stack; vfex: visibility fabric extender of an H-Series cluster", Computed: true}, "cluster_state": schema.StringAttribute{MarkdownDescription: "Describes the cluster membership state. Applicable only for clustered H-series devices", Computed: true}, "cluster_vip": schema.StringAttribute{MarkdownDescription: "applicable only for clustered H-series devices", Computed: true}, "device_id": schema.StringAttribute{MarkdownDescription: "unique ID representing this managed Gigamon device", Computed: true}, "device_ip": schema.StringAttribute{MarkdownDescription: "device management IP Address", Computed: true}, "device_ips": schema.ListAttribute{MarkdownDescription: "list of all the IPs on the device", Computed: true, ElementType: types.StringType}, "disc_outcome": schema.StringAttribute{MarkdownDescription: "status of FM-to-device communication channel", Computed: true}, "disconnected": schema.BoolAttribute{MarkdownDescription: "Status of connection between FM and Device", Computed: true}, "dns_name": schema.StringAttribute{MarkdownDescription: "device DNS name, as registered in DNS server", Computed: true}, "failure_desc": schema.StringAttribute{MarkdownDescription: "reason for device discovery failure", Computed: true}, "family": schema.StringAttribute{Computed: true}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "health_state_reasons": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Computed: true}}}}, "hostname": schema.StringAttribute{MarkdownDescription: "device configured hostname", Computed: true}, "leader_pref": schema.Int64Attribute{MarkdownDescription: "leader election preference rank", Computed: true}, "licensed": schema.BoolAttribute{Computed: true}, "master_pref": schema.Int64Attribute{MarkdownDescription: "master election preference rank. (deprecated: use leaderPref)", Computed: true}, "model": schema.StringAttribute{MarkdownDescription: "Gigamon physical device models", Computed: true}, "nat_ip": schema.StringAttribute{MarkdownDescription: "Device NAT IP Address. Applicable for only devices behind NAT set up", Computed: true}, "oper_status": schema.StringAttribute{MarkdownDescription: "Describes device operational state. 'operational': fully configurable; 'safe': no configuration possible; 'limited': limited configuration possible. Applicable only for H-series devices", Computed: true}, "sw_build_number": schema.StringAttribute{Computed: true}, "sw_version": schema.StringAttribute{Computed: true}, "topo_node_id": schema.StringAttribute{MarkdownDescription: "unique topology node ID. Generated by FM server", Computed: true}}}}}}}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetManagedClustersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetManagedClustersDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	d.readListRemote(ctx, &config, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}

// readListRemote performs the paginated read HTTP exchange and decodes the response array into config. Extracted from Read so the request/response logic is unit-testable without a tfsdk.Config.
func (d *GetManagedClustersDataSource) readListRemote(ctx context.Context, config *GetManagedClustersDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/nodes"
	params := url.Values{}
	if !config.ClusterId.IsNull() {
		params.Set("clusterId", config.ClusterId.ValueString())
	}
	if !config.Page.IsNull() {
		params.Set("page", config.Page.ValueString())
	}
	if !config.Sort.IsNull() {
		params.Set("sort", config.Sort.ValueString())
	}
	var nextURL string
	fetch := func(ctx context.Context, p url.Values) (*http.Response, error) {
		httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
		if err != nil {
			return nil, err
		}
		if nextURL != "" {
			parsed, perr := url.Parse(nextURL)
			if perr != nil {
				return nil, perr
			}
			httpReq.URL = parsed
		} else {
			httpReq.URL.RawQuery = p.Encode()
		}
		return d.client.Do(httpReq)
	}
	pages, err := client.ListAllPages(ctx, params, fetch, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_managed_clusters", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_managed_clusters", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["clusters"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_managed_clusters", fmt.Sprintf("Could not decode list page: missing %q array", "clusters"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_managed_clusters", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetManagedClustersDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
