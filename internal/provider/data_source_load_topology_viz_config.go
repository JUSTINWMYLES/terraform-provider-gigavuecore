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
	_ datasource.DataSource              = (*LoadTopologyVizConfigDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadTopologyVizConfigDataSource)(nil)
)

// LoadTopologyVizConfigDataSource is the generated Terraform data source implementation.
type LoadTopologyVizConfigDataSource struct {
	client *client.Client
}

// LoadTopologyVizConfigDataSourceModel describes the data source state shape.
type LoadTopologyVizConfigDataSourceModel struct {
	AliasConfig               types.List   `tfsdk:"alias_config" json:"aliasConfig"`
	GdpLinkCacheTtlInSecs     types.Int64  `tfsdk:"gdp_link_cache_ttl_in_secs" json:"gdpLinkCacheTtlInSecs"`
	HierarchicalTags          types.List   `tfsdk:"hierarchical_tags" json:"hierarchicalTags"`
	LinkRepresentationConfigs types.List   `tfsdk:"link_representation_configs" json:"linkRepresentationConfigs"`
	PlacementTags             types.List   `tfsdk:"placement_tags" json:"placementTags"`
	SubGroupKeys              types.List   `tfsdk:"sub_group_keys" json:"subGroupKeys"`
	ToolsViewEnabled          types.Bool   `tfsdk:"tools_view_enabled" json:"toolsViewEnabled"`
	TopologyType              types.String `tfsdk:"topology_type" json:"topologyType"`
	Warn                      types.Bool   `tfsdk:"warn"`
}

// NewLoadTopologyVizConfigDataSource returns a new instance of the generated data source.
func NewLoadTopologyVizConfigDataSource() datasource.DataSource {
	return &LoadTopologyVizConfigDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadTopologyVizConfigDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_topology_viz_config"
}

// Schema returns the data source schema.
func (d *LoadTopologyVizConfigDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load FM global topology visualization configurations", Attributes: map[string]schema.Attribute{"alias_config": schema.ListNestedAttribute{MarkdownDescription: "Title of node groups in smart sankey.", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"key": schema.StringAttribute{Computed: true}, "value": schema.StringAttribute{Computed: true}}}}, "gdp_link_cache_ttl_in_secs": schema.Int64Attribute{MarkdownDescription: "When a link is not reported by the GDP, how long the Topology has to maintain the unreported link. If the link reports within the configured TTL time, no changes to the existing link. But the link doesn't report then it will be deleted from Topology maintained links", Computed: true}, "hierarchical_tags": schema.ListAttribute{MarkdownDescription: "List of topology hierarchical tag keys(Ids). This will allow the user to drill down from a top level hierarchy view to a lower level detailed view", Computed: true, ElementType: types.StringType}, "link_representation_configs": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"color_code": schema.StringAttribute{MarkdownDescription: "The color code of the link", Computed: true}, "link_format": schema.StringAttribute{MarkdownDescription: "The format of the link", Computed: true}, "link_type": schema.StringAttribute{MarkdownDescription: "Represents the type of the link", Computed: true}}}}, "placement_tags": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"name": schema.StringAttribute{MarkdownDescription: "Unique alias for the placement config. This will help the user to  give a meaning full name for placement tag config", Computed: true}, "tag_key": schema.StringAttribute{MarkdownDescription: "Unique placement tag key(Id)", Computed: true}, "tag_values": schema.ListAttribute{MarkdownDescription: "List of placement tag values associated with the specified tag key", Computed: true, ElementType: types.StringType}}}}, "sub_group_keys": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "tools_view_enabled": schema.BoolAttribute{MarkdownDescription: "shows tools information in topology view if it is set to 'true'. By default it will be false", Computed: true}, "topology_type": schema.StringAttribute{MarkdownDescription: "Topology type. Decided based on the number of managed nodes and tags configuration. 'TAG_BASED_SANKEY' if the topology tags are configured,'SMART_SANKEY' if the tags are not configured and number of managed nodes is lesser than or equal to 100, 'NOT_DEFINED' if the tags are not configured and number of managed nodes is greater than 100", Computed: true}, "warn": schema.BoolAttribute{MarkdownDescription: "'true' if the tags are not configured and no. of managed nodes is more than the warning threshold(55 nodes),'false' otherwise", Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadTopologyVizConfigDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadTopologyVizConfigDataSourceModel
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
func (d *LoadTopologyVizConfigDataSource) readRemote(ctx context.Context, config *LoadTopologyVizConfigDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/topoviz/config"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_topology_viz_config", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_topology_viz_config", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_topology_viz_config", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_topology_viz_config", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_topology_viz_config", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_topology_viz_config", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_topology_viz_config", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_topology_viz_config", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_topology_viz_config", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_topology_viz_config", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_load_topology_viz_config", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_load_topology_viz_config", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_topology_viz_config", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["topologyConfigs"]; ok {
		if m, ok := v.(map[string]any); ok {
			data = m
		}
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_topology_viz_config", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadTopologyVizConfigDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
