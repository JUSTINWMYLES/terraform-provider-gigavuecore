package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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
	Alias types.String `tfsdk:"alias"`
	Items types.List   `tfsdk:"items"`
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
	resp.Schema = schema.Schema{MarkdownDescription: "Find TrafficPolicyGraph Tunnel Interface Mappings by alias", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "alias of the target TrafficPolicyGraph", Required: true}, "items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"monitoring_session_endpoint_iface_mappings": schema.SingleNestedAttribute{MarkdownDescription: "Monitoring Session Endpoint Info", Computed: true, Attributes: map[string]schema.Attribute{"monitoring_session_id": schema.StringAttribute{Computed: true}, "vseries_endpoint_iface_mappings": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"endpoint_iface_mappings": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"endpoint_id": schema.StringAttribute{Computed: true}, "iface": schema.StringAttribute{Computed: true}}}}, "vseries_node_ids": schema.ListAttribute{MarkdownDescription: "List of V Series Node Ids", Computed: true, ElementType: types.StringType}}}}}}}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetTrafficPolicyGraphTunnelInterfaceMappingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetTrafficPolicyGraphTunnelInterfaceMappingsDataSourceModel
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
func (d *GetTrafficPolicyGraphTunnelInterfaceMappingsDataSource) readListRemote(ctx context.Context, config *GetTrafficPolicyGraphTunnelInterfaceMappingsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/trafficPolicyGraph/{alias}/tunnelInterfaceMappings"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	params := url.Values{}
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_policy_graph_tunnel_interface_mappings", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_policy_graph_tunnel_interface_mappings", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["trafficPolicyGraphEndpointInterfaceMappings"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_policy_graph_tunnel_interface_mappings", fmt.Sprintf("Could not decode list page: missing %q array", "trafficPolicyGraphEndpointInterfaceMappings"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
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
