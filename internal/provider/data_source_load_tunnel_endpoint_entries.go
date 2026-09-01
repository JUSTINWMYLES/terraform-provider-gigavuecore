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
	_ datasource.DataSource              = (*LoadTunnelEndpointEntriesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadTunnelEndpointEntriesDataSource)(nil)
)

// LoadTunnelEndpointEntriesDataSource is the generated Terraform data source implementation.
type LoadTunnelEndpointEntriesDataSource struct {
	client *client.Client
}

// LoadTunnelEndpointEntriesDataSourceModel describes the data source state shape.
type LoadTunnelEndpointEntriesDataSourceModel struct {
	ClusterId  types.String `tfsdk:"cluster_id" json:"clusterId"`
	Items      types.List   `tfsdk:"items"`
	Page       types.String `tfsdk:"page"`
	Sort       types.String `tfsdk:"sort"`
	TunnelType types.String `tfsdk:"tunnel_type" json:"tunnelType"`
}

// NewLoadTunnelEndpointEntriesDataSource returns a new instance of the generated data source.
func NewLoadTunnelEndpointEntriesDataSource() datasource.DataSource {
	return &LoadTunnelEndpointEntriesDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadTunnelEndpointEntriesDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_tunnel_endpoint_entries"
}

// Schema returns the data source schema.
func (d *LoadTunnelEndpointEntriesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load known Tunnel Endpoints", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Cluster ID to filter by", Required: true}, "items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"tunnel_endpoint": schema.SingleNestedAttribute{MarkdownDescription: "Tunnel Endpoint configuration", Computed: true, Attributes: map[string]schema.Attribute{"address": schema.StringAttribute{MarkdownDescription: "Tunnel Endpoint IP Address on the TunneledPort", Computed: true}, "gre_key": schema.Int64Attribute{MarkdownDescription: "only applicable for 'l2gre' type", Computed: true}, "port": schema.Int64Attribute{MarkdownDescription: "Tunnel Endpoint Port on the TunneledPort. Only applicable and required for 'gmip' type", Computed: true}, "src_port": schema.Int64Attribute{MarkdownDescription: "Tunnel Source Port (the port on the remote/client side). Only applicable and required for 'gmip' type", Computed: true}, "type": schema.StringAttribute{Computed: true}}}, "tunneled_port": schema.SingleNestedAttribute{MarkdownDescription: "Tunnel Endpoint's servicing Tunneled Port information", Computed: true, Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "ID of a cluster the tunneled port is a member of", Computed: true}, "device_box_id": schema.StringAttribute{MarkdownDescription: "BoxId of the device the tunneled port is a member of", Computed: true}, "device_id": schema.StringAttribute{MarkdownDescription: "ID of the device the tunneled port is a member of", Computed: true}, "device_model": schema.StringAttribute{MarkdownDescription: "model of the device hosting the tunneled port", Computed: true}, "port_id": schema.StringAttribute{MarkdownDescription: "Tunneled port ID on a Gigamon Chassis", Computed: true}}}}}}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned. Defaults to (1:100) to prevent reading entire DB by accident", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}, "tunnel_type": schema.StringAttribute{MarkdownDescription: "Tunnel type to filter by", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadTunnelEndpointEntriesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadTunnelEndpointEntriesDataSourceModel
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
func (d *LoadTunnelEndpointEntriesDataSource) readListRemote(ctx context.Context, config *LoadTunnelEndpointEntriesDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/tunnelEndpoints"
	params := url.Values{}
	params.Set("clusterId", config.ClusterId.ValueString())
	if !config.TunnelType.IsNull() {
		params.Set("tunnelType", config.TunnelType.ValueString())
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
		resp.Diagnostics.AddError("Error reading gigavuecore_load_tunnel_endpoint_entries", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_tunnel_endpoint_entries", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["tunnelEndpoints"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_tunnel_endpoint_entries", fmt.Sprintf("Could not decode list page: missing %q array", "tunnelEndpoints"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_tunnel_endpoint_entries", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadTunnelEndpointEntriesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
