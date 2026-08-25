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
	_ datasource.DataSource              = (*GetTrafficPolicyGraphTunnelInterfaceMappingsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetTrafficPolicyGraphTunnelInterfaceMappingsDataSource)(nil)
)

// GetTrafficPolicyGraphTunnelInterfaceMappingsDataSource is the generated Terraform data source implementation.
type GetTrafficPolicyGraphTunnelInterfaceMappingsDataSource struct {
	client *client.Client
}

// GetTrafficPolicyGraphTunnelInterfaceMappingsDataSourceModel describes the data source state shape.
type GetTrafficPolicyGraphTunnelInterfaceMappingsDataSourceModel struct {
	Alias                                       types.String `tfsdk:"alias"`
	Context                                     types.Object `tfsdk:"context"`
	TrafficPolicyGraphEndpointInterfaceMappings types.List   `tfsdk:"traffic_policy_graph_endpoint_interface_mappings" json:"trafficPolicyGraphEndpointInterfaceMappings"`
}

// NewGetTrafficPolicyGraphTunnelInterfaceMappingsDataSource returns a new instance of the generated data source.
func NewGetTrafficPolicyGraphTunnelInterfaceMappingsDataSource() datasource.DataSource {
	return &GetTrafficPolicyGraphTunnelInterfaceMappingsDataSource{}
}

// Metadata returns the data source type name.
func (d *GetTrafficPolicyGraphTunnelInterfaceMappingsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_traffic_policy_graph_tunnel_interface_mappings"
}

// Schema returns the data source schema.
func (d *GetTrafficPolicyGraphTunnelInterfaceMappingsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Find TrafficPolicyGraph Tunnel Interface Mappings by alias", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "alias of the target TrafficPolicyGraph", Required: true}, "context": schema.SingleNestedAttribute{MarkdownDescription: "Gigamon query result context", Computed: true, Attributes: map[string]schema.Attribute{"page_no": schema.Int64Attribute{MarkdownDescription: "page number of the returned result set", Computed: true}, "page_size": schema.Int64Attribute{MarkdownDescription: "page size of the returned result set", Computed: true}, "sort": schema.ListAttribute{MarkdownDescription: "sorting info of the returned result set. list of fields in the array indicate sorting order", Computed: true, ElementType: types.StringType}, "total_items": schema.Int64Attribute{MarkdownDescription: "total number of items in the queried entity type", Computed: true}}}, "traffic_policy_graph_endpoint_interface_mappings": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"monitoring_session_endpoint_iface_mappings": schema.SingleNestedAttribute{MarkdownDescription: "Monitoring Session Endpoint Info", Computed: true, Attributes: map[string]schema.Attribute{"monitoring_session_id": schema.StringAttribute{Computed: true}, "vseries_endpoint_iface_mappings": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"endpoint_iface_mappings": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"endpoint_id": schema.StringAttribute{Computed: true}, "iface": schema.StringAttribute{Computed: true}}}}, "vseries_node_ids": schema.ListAttribute{MarkdownDescription: "List of V Series Node Ids", Computed: true, ElementType: types.StringType}}}}}}}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetTrafficPolicyGraphTunnelInterfaceMappingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetTrafficPolicyGraphTunnelInterfaceMappingsDataSourceModel
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
func (d *GetTrafficPolicyGraphTunnelInterfaceMappingsDataSource) readRemote(ctx context.Context, config *GetTrafficPolicyGraphTunnelInterfaceMappingsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/trafficPolicyGraph/{alias}/tunnelInterfaceMappings"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_policy_graph_tunnel_interface_mappings", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_policy_graph_tunnel_interface_mappings", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_policy_graph_tunnel_interface_mappings", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_policy_graph_tunnel_interface_mappings", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_policy_graph_tunnel_interface_mappings", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_policy_graph_tunnel_interface_mappings", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_policy_graph_tunnel_interface_mappings", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_policy_graph_tunnel_interface_mappings", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_policy_graph_tunnel_interface_mappings", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_policy_graph_tunnel_interface_mappings", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_policy_graph_tunnel_interface_mappings", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_policy_graph_tunnel_interface_mappings", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_policy_graph_tunnel_interface_mappings", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_policy_graph_tunnel_interface_mappings", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetTrafficPolicyGraphTunnelInterfaceMappingsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
