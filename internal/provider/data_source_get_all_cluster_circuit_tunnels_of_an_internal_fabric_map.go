package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*GetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSource)(nil)
)

// GetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSource is the generated Terraform data source implementation.
type GetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSource struct {
	client *client.Client
}

// GetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSourceModel describes the data source state shape.
type GetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSourceModel struct {
	Alias       types.String `tfsdk:"alias"`
	ClusterMaps types.List   `tfsdk:"cluster_maps" json:"clusterMaps"`
	Context     types.Object `tfsdk:"context"`
	IfmAlias    types.String `tfsdk:"ifm_alias" json:"ifmAlias"`
	Mode        types.String `tfsdk:"mode"`
	Type        types.String `tfsdk:"type"`
}

// NewGetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSource returns a new instance of the generated data source.
func NewGetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSource() datasource.DataSource {
	return &GetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSource{}
}

// Metadata returns the data source type name.
func (d *GetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_cluster_circuit_tunnels_of_an_internal_fabric_map"
}

// Schema returns the data source schema.
func (d *GetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get all cluster circuit tunnel endpoints of an internal fabric map", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "alias of the fabric map", Required: true}, "cluster_maps": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"circuit_tunnel": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "tunnel name", Computed: true}, "attach": schema.ListAttribute{MarkdownDescription: "ipInterface if type is vxlan. ports or gigastreams if type is circuit and is valid only with decap mode.", Computed: true, ElementType: types.StringType}, "circuit_ids": schema.ListAttribute{MarkdownDescription: "circuit ids, valid and required if type is circuit", Computed: true, ElementType: types.Int64Type}, "cluster_id": schema.StringAttribute{MarkdownDescription: "id of the cluster in which this circuit tunnel is created", Computed: true}, "comment": schema.StringAttribute{Computed: true}, "dip_address": schema.StringAttribute{Computed: true}, "l4_src_port": schema.Int64Attribute{Computed: true}, "mode": schema.StringAttribute{MarkdownDescription: "tunnel mode", Computed: true}, "type": schema.StringAttribute{Computed: true}}}, "cluster_id": schema.StringAttribute{MarkdownDescription: "id of the cluster in which the circuit tunnel endpoint is created.", Computed: true}, "config_status": schema.StringAttribute{MarkdownDescription: "Configuration status of this circuit tunnel endpoint.", Computed: true}, "error_message": schema.StringAttribute{MarkdownDescription: "In case of configuration failure, this message provides details about the possible cause of the failure.", Computed: true}, "fabric_map_aliases": schema.ListAttribute{MarkdownDescription: "Aliases of the fabric maps that this circuit tunnel endpoint supports.", Computed: true, ElementType: types.StringType}}}}, "context": schema.SingleNestedAttribute{MarkdownDescription: "Gigamon query result context", Computed: true, Attributes: map[string]schema.Attribute{"page_no": schema.Int64Attribute{MarkdownDescription: "page number of the returned result set", Computed: true}, "page_size": schema.Int64Attribute{MarkdownDescription: "page size of the returned result set", Computed: true}, "sort": schema.ListAttribute{MarkdownDescription: "sorting info of the returned result set. list of fields in the array indicate sorting order", Computed: true, ElementType: types.StringType}, "total_items": schema.Int64Attribute{MarkdownDescription: "total number of items in the queried entity type", Computed: true}}}, "ifm_alias": schema.StringAttribute{MarkdownDescription: "alias of the internal fabric map", Required: true}, "mode": schema.StringAttribute{MarkdownDescription: "tunnel mode: 'encap' or 'decap'.", Optional: true}, "type": schema.StringAttribute{MarkdownDescription: "tunnel type: 'circuit' or 'vxlan'.", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSourceModel
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
func (d *GetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSource) readRemote(ctx context.Context, config *GetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/fabricMaps/{alias}/internalFabricMaps/{ifmAlias}/clusterCircuitTunnels"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	reqPath = strings.ReplaceAll(reqPath, "{ifmAlias}", url.PathEscape(config.IfmAlias.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_circuit_tunnels_of_an_internal_fabric_map", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.Mode.IsNull() {
		query.Set("mode", config.Mode.ValueString())
	}
	if !config.Type.IsNull() {
		query.Set("type", config.Type.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_circuit_tunnels_of_an_internal_fabric_map", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_circuit_tunnels_of_an_internal_fabric_map", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_circuit_tunnels_of_an_internal_fabric_map", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_circuit_tunnels_of_an_internal_fabric_map", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_circuit_tunnels_of_an_internal_fabric_map", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_circuit_tunnels_of_an_internal_fabric_map", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_circuit_tunnels_of_an_internal_fabric_map", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_circuit_tunnels_of_an_internal_fabric_map", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_circuit_tunnels_of_an_internal_fabric_map", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_circuit_tunnels_of_an_internal_fabric_map", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_circuit_tunnels_of_an_internal_fabric_map", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_circuit_tunnels_of_an_internal_fabric_map", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_circuit_tunnels_of_an_internal_fabric_map", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
