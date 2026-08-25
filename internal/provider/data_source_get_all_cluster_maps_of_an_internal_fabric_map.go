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
	_ datasource.DataSource              = (*GetAllClusterMapsOfAnInternalFabricMapDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAllClusterMapsOfAnInternalFabricMapDataSource)(nil)
)

// GetAllClusterMapsOfAnInternalFabricMapDataSource is the generated Terraform data source implementation.
type GetAllClusterMapsOfAnInternalFabricMapDataSource struct {
	client *client.Client
}

// GetAllClusterMapsOfAnInternalFabricMapDataSourceModel describes the data source state shape.
type GetAllClusterMapsOfAnInternalFabricMapDataSourceModel struct {
	Alias       types.String  `tfsdk:"alias"`
	ClusterMaps types.Dynamic `tfsdk:"cluster_maps" json:"clusterMaps"`
	Context     types.Object  `tfsdk:"context"`
	IfmAlias    types.String  `tfsdk:"ifm_alias" json:"ifmAlias"`
}

// NewGetAllClusterMapsOfAnInternalFabricMapDataSource returns a new instance of the generated data source.
func NewGetAllClusterMapsOfAnInternalFabricMapDataSource() datasource.DataSource {
	return &GetAllClusterMapsOfAnInternalFabricMapDataSource{}
}

// Metadata returns the data source type name.
func (d *GetAllClusterMapsOfAnInternalFabricMapDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_cluster_maps_of_an_internal_fabric_map"
}

// Schema returns the data source schema.
func (d *GetAllClusterMapsOfAnInternalFabricMapDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get all cluster maps of an internal fabric map", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "alias of the fabric map", Required: true}, "cluster_maps": schema.DynamicAttribute{Computed: true}, "context": schema.SingleNestedAttribute{MarkdownDescription: "Gigamon query result context", Computed: true, Attributes: map[string]schema.Attribute{"page_no": schema.Int64Attribute{MarkdownDescription: "page number of the returned result set", Computed: true}, "page_size": schema.Int64Attribute{MarkdownDescription: "page size of the returned result set", Computed: true}, "sort": schema.ListAttribute{MarkdownDescription: "sorting info of the returned result set. list of fields in the array indicate sorting order", Computed: true, ElementType: types.StringType}, "total_items": schema.Int64Attribute{MarkdownDescription: "total number of items in the queried entity type", Computed: true}}}, "ifm_alias": schema.StringAttribute{MarkdownDescription: "alias of the internal fabric map", Required: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetAllClusterMapsOfAnInternalFabricMapDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAllClusterMapsOfAnInternalFabricMapDataSourceModel
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
func (d *GetAllClusterMapsOfAnInternalFabricMapDataSource) readRemote(ctx context.Context, config *GetAllClusterMapsOfAnInternalFabricMapDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/fabricMaps/{alias}/internalFabricMaps/{ifmAlias}/clusterMaps"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	reqPath = strings.ReplaceAll(reqPath, "{ifmAlias}", url.PathEscape(config.IfmAlias.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_maps_of_an_internal_fabric_map", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_maps_of_an_internal_fabric_map", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_maps_of_an_internal_fabric_map", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_maps_of_an_internal_fabric_map", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_maps_of_an_internal_fabric_map", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_maps_of_an_internal_fabric_map", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_maps_of_an_internal_fabric_map", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_maps_of_an_internal_fabric_map", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_maps_of_an_internal_fabric_map", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_maps_of_an_internal_fabric_map", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_maps_of_an_internal_fabric_map", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_maps_of_an_internal_fabric_map", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_maps_of_an_internal_fabric_map", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_maps_of_an_internal_fabric_map", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetAllClusterMapsOfAnInternalFabricMapDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
