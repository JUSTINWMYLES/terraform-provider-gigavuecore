package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*LoadTopologyDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadTopologyDataSource)(nil)
)

// LoadTopologyDataSource is the generated Terraform data source implementation.
type LoadTopologyDataSource struct {
	client *client.Client
}

// LoadTopologyDataSourceModel describes the data source state shape.
type LoadTopologyDataSourceModel struct {
	ClusterId    types.String  `tfsdk:"cluster_id" json:"clusterId"`
	Links        types.List    `tfsdk:"links"`
	Nodes        types.Dynamic `tfsdk:"nodes"`
	TopoNodeType types.String  `tfsdk:"topo_node_type" json:"topoNodeType"`
}

// NewLoadTopologyDataSource returns a new instance of the generated data source.
func NewLoadTopologyDataSource() datasource.DataSource {
	return &LoadTopologyDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadTopologyDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_topology"
}

// Schema returns the data source schema.
func (d *LoadTopologyDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load Classic Topology", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Cluster ID to filter by", Optional: true}, "links": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"comment": schema.StringAttribute{Computed: true}, "connections": schema.ListNestedAttribute{MarkdownDescription: "Actual port-ids which are part of the physical connection", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"port1": schema.StringAttribute{Computed: true}, "port2": schema.StringAttribute{Computed: true}}}}, "endpoint1": schema.SingleNestedAttribute{MarkdownDescription: "node link endpoint", Computed: true, Attributes: map[string]schema.Attribute{"comment": schema.StringAttribute{Computed: true}, "component_alias": schema.StringAttribute{MarkdownDescription: "For 'gigamon': this is either portId (e.g. '1/1/x1') or Gigastream alias (e.g. 'gs1'). For 'lldp'/'cdp': this is the 'portId' as reported by neighbor discovery. For 'manual': this is user provided or may be omitted", Computed: true}, "component_ports": schema.ListAttribute{MarkdownDescription: "For the component type gigastream, this specifies the array of port Ids that forms the logical group", Computed: true, ElementType: types.StringType}, "component_type": schema.StringAttribute{MarkdownDescription: "Type of the endpoint component. This is either 'Port' or 'GigaStream'", Computed: true}, "node_alias": schema.StringAttribute{MarkdownDescription: "Node alias. References user-assigned topology node alias. Internal use only", Computed: true}, "port": schema.StringAttribute{MarkdownDescription: "For 'gigamon': this is the portId (e.g. '1/1/x1'). For 'lldp'/'cdp': this is the 'portId' as reported by neighbor discovery. For 'manual': this is user provided or may be omitted", Computed: true}, "topo_node_id": schema.StringAttribute{MarkdownDescription: "Topology graph node Id. References one of the Topology Nodes in the graph", Computed: true}, "topo_node_type": schema.StringAttribute{MarkdownDescription: "topology node type", Computed: true}}}, "endpoint2": schema.SingleNestedAttribute{MarkdownDescription: "node link endpoint", Computed: true, Attributes: map[string]schema.Attribute{"comment": schema.StringAttribute{Computed: true}, "component_alias": schema.StringAttribute{MarkdownDescription: "For 'gigamon': this is either portId (e.g. '1/1/x1') or Gigastream alias (e.g. 'gs1'). For 'lldp'/'cdp': this is the 'portId' as reported by neighbor discovery. For 'manual': this is user provided or may be omitted", Computed: true}, "component_ports": schema.ListAttribute{MarkdownDescription: "For the component type gigastream, this specifies the array of port Ids that forms the logical group", Computed: true, ElementType: types.StringType}, "component_type": schema.StringAttribute{MarkdownDescription: "Type of the endpoint component. This is either 'Port' or 'GigaStream'", Computed: true}, "node_alias": schema.StringAttribute{MarkdownDescription: "Node alias. References user-assigned topology node alias. Internal use only", Computed: true}, "port": schema.StringAttribute{MarkdownDescription: "For 'gigamon': this is the portId (e.g. '1/1/x1'). For 'lldp'/'cdp': this is the 'portId' as reported by neighbor discovery. For 'manual': this is user provided or may be omitted", Computed: true}, "topo_node_id": schema.StringAttribute{MarkdownDescription: "Topology graph node Id. References one of the Topology Nodes in the graph", Computed: true}, "topo_node_type": schema.StringAttribute{MarkdownDescription: "topology node type", Computed: true}}}, "giga_link_type": schema.StringAttribute{MarkdownDescription: "specifies whether this is a gigastream-based or port-based link. Only applicable for topoNodeType of 'gigamon'", Computed: true}, "link_state": schema.StringAttribute{MarkdownDescription: "specifies health state", Computed: true}, "manual_override": schema.BoolAttribute{MarkdownDescription: "Specifies whether manual augmentation exist. Only applicable for topoLinkType of 'cdp' or 'lldp'", Computed: true}, "topo_link_id": schema.StringAttribute{MarkdownDescription: "unique Id representing a topology graph link. Generated by FM server", Computed: true}, "topo_link_type": schema.StringAttribute{MarkdownDescription: "indicates how link became a pert of this topology graph", Computed: true}}}}, "nodes": schema.DynamicAttribute{Computed: true}, "topo_node_type": schema.StringAttribute{MarkdownDescription: "Comma-separated list of topoNodeTypes to filter by. Valid values are ['gigamon', 'manual', 'lldp', 'cdp']", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadTopologyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadTopologyDataSourceModel
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
func (d *LoadTopologyDataSource) readRemote(ctx context.Context, config *LoadTopologyDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/topology"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_topology", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.ClusterId.IsNull() {
		query.Set("clusterId", config.ClusterId.ValueString())
	}
	if !config.TopoNodeType.IsNull() {
		query.Set("topoNodeType", config.TopoNodeType.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_topology", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_topology", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_topology", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_topology", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_topology", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_topology", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_topology", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_topology", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_topology", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_load_topology", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_load_topology", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_topology", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["gigaTopology"]; ok {
		if m, ok := v.(map[string]any); ok {
			data = m
		} else if arr, ok := v.([]any); ok && len(arr) > 0 {
			if m, ok := arr[0].(map[string]any); ok {
				data = m
			}
		}
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_topology", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadTopologyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
