package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*LoadFabricPathDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadFabricPathDataSource)(nil)
)

// LoadFabricPathDataSource is the generated Terraform data source implementation.
type LoadFabricPathDataSource struct {
	client *client.Client
}

// LoadFabricPathDataSourceModel describes the data source state shape.
type LoadFabricPathDataSourceModel struct {
	Details      types.Bool   `tfsdk:"details"`
	DstClusterId types.String `tfsdk:"dst_cluster_id" json:"dstClusterId"`
	Items        types.List   `tfsdk:"items"`
	SrcClusterId types.String `tfsdk:"src_cluster_id" json:"srcClusterId"`
}

// NewLoadFabricPathDataSource returns a new instance of the generated data source.
func NewLoadFabricPathDataSource() datasource.DataSource {
	return &LoadFabricPathDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadFabricPathDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_fabric_path"
}

// Schema returns the data source schema.
func (d *LoadFabricPathDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load fabricPath details for the given srcClusterID and dstClusterID", Attributes: map[string]schema.Attribute{"details": schema.BoolAttribute{MarkdownDescription: "load comprehensive response on Fabric Path, including detailed information on topology link endpoints", Optional: true}, "dst_cluster_id": schema.StringAttribute{MarkdownDescription: "destination Cluster ID", Required: true}, "items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Composing from srcClusterUUID and dstClusterUUID strings, i.e., srcClusterUUID-to-dstClusterUUID", Computed: true}, "config_state": schema.StringAttribute{MarkdownDescription: "Configuration state of the fabric path", Computed: true}, "dst_cluster_uuid": schema.StringAttribute{MarkdownDescription: "All dstVertices must be from same dstClusterUUID", Computed: true}, "dst_vertices": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "gigastream alias. Uniquely identifies a gigastream within a cluster", Computed: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "id of the defining cluster", Computed: true}, "comment": schema.StringAttribute{Computed: true}, "drop_weight": schema.Int64Attribute{MarkdownDescription: "relative weight for dropping the traffic", Computed: true}, "failover_status": schema.StringAttribute{MarkdownDescription: "Failover Status", Computed: true}, "hash_size": schema.Int64Attribute{MarkdownDescription: "Hash bucket size", Computed: true}, "hash_tool_port": schema.ListNestedAttribute{MarkdownDescription: "Hash bucket id to tool port mapping", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"hash_bucket_ids": schema.ListAttribute{MarkdownDescription: "hash bucket id or range", Computed: true, ElementType: types.Int64Type}, "tool_ports": schema.ListAttribute{MarkdownDescription: "tool port(s) mapped to hashBucketIds", Computed: true, ElementType: types.StringType}}}}, "hash_type": schema.StringAttribute{Computed: true}, "hash_weights": schema.ListAttribute{MarkdownDescription: "hashWeights for 'ports'.If included, the list size must match the size of the 'ports' list", Computed: true, ElementType: types.Int64Type}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "health_state_reasons": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Computed: true}}}}, "ports": schema.ListAttribute{MarkdownDescription: "list of the ports to combine into a gigastream", Computed: true, ElementType: types.StringType}, "threshold_level": schema.StringAttribute{MarkdownDescription: "Threshold level", Computed: true}, "variance_threshold": schema.StringAttribute{MarkdownDescription: "Variance threshold percentage", Computed: true}}}}, "gfp_id": schema.Int64Attribute{MarkdownDescription: "Can be a hash number from alias string", Computed: true}, "operation_state": schema.StringAttribute{MarkdownDescription: "Operation state of the fabric path", Computed: true}, "path_type": schema.StringAttribute{MarkdownDescription: "Type for Fabric Path", Computed: true}, "referenced": schema.StringAttribute{MarkdownDescription: "Reference boolean of fabric path", Computed: true}, "src_cluster_uuid": schema.StringAttribute{MarkdownDescription: "All srcVertices must be from same srcClusterUUID", Computed: true}, "src_vertices": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "gigastream alias. Uniquely identifies a gigastream within a cluster", Computed: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "id of the defining cluster", Computed: true}, "comment": schema.StringAttribute{Computed: true}, "drop_weight": schema.Int64Attribute{MarkdownDescription: "relative weight for dropping the traffic", Computed: true}, "failover_status": schema.StringAttribute{MarkdownDescription: "Failover Status", Computed: true}, "hash_size": schema.Int64Attribute{MarkdownDescription: "Hash bucket size", Computed: true}, "hash_tool_port": schema.ListNestedAttribute{MarkdownDescription: "Hash bucket id to tool port mapping", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"hash_bucket_ids": schema.ListAttribute{MarkdownDescription: "hash bucket id or range", Computed: true, ElementType: types.Int64Type}, "tool_ports": schema.ListAttribute{MarkdownDescription: "tool port(s) mapped to hashBucketIds", Computed: true, ElementType: types.StringType}}}}, "hash_type": schema.StringAttribute{Computed: true}, "hash_weights": schema.ListAttribute{MarkdownDescription: "hashWeights for 'ports'.If included, the list size must match the size of the 'ports' list", Computed: true, ElementType: types.Int64Type}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "health_state_reasons": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Computed: true}}}}, "ports": schema.ListAttribute{MarkdownDescription: "list of the ports to combine into a gigastream", Computed: true, ElementType: types.StringType}, "threshold_level": schema.StringAttribute{MarkdownDescription: "Threshold level", Computed: true}, "variance_threshold": schema.StringAttribute{MarkdownDescription: "Variance threshold percentage", Computed: true}}}}, "topology_links": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"aggregated_link": schema.BoolAttribute{Computed: true}, "comment": schema.StringAttribute{Computed: true}, "discovery_state": schema.StringAttribute{MarkdownDescription: "specifies the discovery state of the topo link", Computed: true}, "endpoint1": schema.SingleNestedAttribute{MarkdownDescription: "node link endpoint", Computed: true, Attributes: map[string]schema.Attribute{"comment": schema.StringAttribute{Computed: true}, "component_alias": schema.StringAttribute{MarkdownDescription: "For 'gigamon': this is either portId (e.g. '1/1/x1') or Gigastream alias (e.g. 'gs1'). For 'lldp'/'cdp': this is the 'portId' as reported by neighbor discovery. For 'manual': this is user provided or may be omitted", Computed: true}, "component_ports": schema.ListAttribute{MarkdownDescription: "For the component type gigastream, this specifies the array of port Ids that forms the logical group", Computed: true, ElementType: types.StringType}, "component_type": schema.StringAttribute{MarkdownDescription: "Type of the endpoint component. This is either 'Port' or 'GigaStream'", Computed: true}, "node_alias": schema.StringAttribute{MarkdownDescription: "Node alias. References user-assigned topology node alias. Internal use only", Computed: true}, "port": schema.StringAttribute{MarkdownDescription: "For 'gigamon': this is the portId (e.g. '1/1/x1'). For 'lldp'/'cdp': this is the 'portId' as reported by neighbor discovery. For 'manual': this is user provided or may be omitted", Computed: true}, "topo_node_id": schema.StringAttribute{MarkdownDescription: "Topology graph node Id. References one of the Topology Nodes in the graph", Computed: true}, "topo_node_type": schema.StringAttribute{MarkdownDescription: "topology node type", Computed: true}}}, "endpoint2": schema.SingleNestedAttribute{MarkdownDescription: "node link endpoint", Computed: true, Attributes: map[string]schema.Attribute{"comment": schema.StringAttribute{Computed: true}, "component_alias": schema.StringAttribute{MarkdownDescription: "For 'gigamon': this is either portId (e.g. '1/1/x1') or Gigastream alias (e.g. 'gs1'). For 'lldp'/'cdp': this is the 'portId' as reported by neighbor discovery. For 'manual': this is user provided or may be omitted", Computed: true}, "component_ports": schema.ListAttribute{MarkdownDescription: "For the component type gigastream, this specifies the array of port Ids that forms the logical group", Computed: true, ElementType: types.StringType}, "component_type": schema.StringAttribute{MarkdownDescription: "Type of the endpoint component. This is either 'Port' or 'GigaStream'", Computed: true}, "node_alias": schema.StringAttribute{MarkdownDescription: "Node alias. References user-assigned topology node alias. Internal use only", Computed: true}, "port": schema.StringAttribute{MarkdownDescription: "For 'gigamon': this is the portId (e.g. '1/1/x1'). For 'lldp'/'cdp': this is the 'portId' as reported by neighbor discovery. For 'manual': this is user provided or may be omitted", Computed: true}, "topo_node_id": schema.StringAttribute{MarkdownDescription: "Topology graph node Id. References one of the Topology Nodes in the graph", Computed: true}, "topo_node_type": schema.StringAttribute{MarkdownDescription: "topology node type", Computed: true}}}, "giga_link_type": schema.StringAttribute{MarkdownDescription: "specifies whether this is a gigastream-based or port-based link. Only applicable for topoNodeType of 'gigamon'", Computed: true}, "link_alias": schema.StringAttribute{MarkdownDescription: "alias of the link", Computed: true}, "link_state": schema.StringAttribute{MarkdownDescription: "specifies health state", Computed: true}, "manual_override": schema.BoolAttribute{MarkdownDescription: "Specifies whether manual augmentation exist. Only applicable for topoLinkType of 'cdp' or 'lldp'", Computed: true}, "topo_link_id": schema.StringAttribute{MarkdownDescription: "unique Id representing a topology graph link. Generated by FM server", Computed: true}, "topo_link_type": schema.StringAttribute{MarkdownDescription: "indicates how link became a pert of this topology graph", Computed: true}}}}, "version": schema.Int64Attribute{MarkdownDescription: "Incremental version field to indicate the fabric path format", Computed: true}}}}, "src_cluster_id": schema.StringAttribute{MarkdownDescription: "source Cluster ID", Required: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadFabricPathDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadFabricPathDataSourceModel
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
func (d *LoadFabricPathDataSource) readListRemote(ctx context.Context, config *LoadFabricPathDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/topology/fabricPath"
	params := url.Values{}
	params.Set("srcClusterId", config.SrcClusterId.ValueString())
	params.Set("dstClusterId", config.DstClusterId.ValueString())
	if !config.Details.IsNull() {
		params.Set("details", strconv.FormatBool(config.Details.ValueBool()))
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
		resp.Diagnostics.AddError("Error reading gigavuecore_load_fabric_path", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_fabric_path", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["gigaTopologyFabricPathsDetails"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_fabric_path", fmt.Sprintf("Could not decode list page: missing %q array", "gigaTopologyFabricPathsDetails"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_fabric_path", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadFabricPathDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
