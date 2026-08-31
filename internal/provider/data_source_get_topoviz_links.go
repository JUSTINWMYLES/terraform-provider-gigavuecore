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
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*GetTopovizLinksDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetTopovizLinksDataSource)(nil)
)

// GetTopovizLinksDataSource is the generated Terraform data source implementation.
type GetTopovizLinksDataSource struct {
	client *client.Client
}

// GetTopovizLinksDataSourceModel describes the data source state shape.
type GetTopovizLinksDataSourceModel struct {
	Endpoint1Alias    types.String `tfsdk:"endpoint1_alias" json:"endpoint1Alias"`
	Endpoint1HostName types.String `tfsdk:"endpoint1_host_name" json:"endpoint1HostName"`
	Endpoint2Alias    types.String `tfsdk:"endpoint2_alias" json:"endpoint2Alias"`
	Endpoint2HostName types.String `tfsdk:"endpoint2_host_name" json:"endpoint2HostName"`
	Items             types.List   `tfsdk:"items"`
	LinkSource        types.String `tfsdk:"link_source" json:"linkSource"`
	LinkSpeed         types.String `tfsdk:"link_speed" json:"linkSpeed"`
	LinkType          types.String `tfsdk:"link_type" json:"linkType"`
	Page              types.String `tfsdk:"page"`
	Sort              types.String `tfsdk:"sort"`
}

// NewGetTopovizLinksDataSource returns a new instance of the generated data source.
func NewGetTopovizLinksDataSource() datasource.DataSource {
	return &GetTopovizLinksDataSource{}
}

// Metadata returns the data source type name.
func (d *GetTopovizLinksDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_topoviz_links"
}

// Schema returns the data source schema.
func (d *GetTopovizLinksDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Loads all topology links", Attributes: map[string]schema.Attribute{"endpoint1_alias": schema.StringAttribute{MarkdownDescription: "Aliases of endpoint port/gigastream", Optional: true}, "endpoint1_host_name": schema.StringAttribute{MarkdownDescription: "Hostname of endpoint device", Optional: true}, "endpoint2_alias": schema.StringAttribute{MarkdownDescription: "Aliases of endpoint port/gigastream", Optional: true}, "endpoint2_host_name": schema.StringAttribute{MarkdownDescription: "Hostname of endpoint devices", Optional: true}, "items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"connections": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"port1": schema.StringAttribute{Computed: true}, "port1_reported": schema.BoolAttribute{Computed: true}, "port2": schema.StringAttribute{Computed: true}, "port2_reported": schema.BoolAttribute{Computed: true}}}}, "endpoint1": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "PortId if endpointType is port, alias if it is a gigastream", Computed: true}, "endpoint_type": schema.StringAttribute{Computed: true}, "port_type": schema.StringAttribute{Computed: true}, "ports": schema.ListAttribute{MarkdownDescription: "List of portIds in the gigastream", Computed: true, ElementType: types.StringType}}}}, "endpoint1_cluster_name": schema.StringAttribute{Computed: true}, "endpoint1_global_node_id": schema.StringAttribute{Computed: true}, "endpoint2": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "PortId if endpointType is port, alias if it is a gigastream", Computed: true}, "endpoint_type": schema.StringAttribute{Computed: true}, "port_type": schema.StringAttribute{Computed: true}, "ports": schema.ListAttribute{MarkdownDescription: "List of portIds in the gigastream", Computed: true, ElementType: types.StringType}}}}, "endpoint2_cluster_name": schema.StringAttribute{Computed: true}, "endpoint2_global_node_id": schema.StringAttribute{Computed: true}, "link_id": schema.StringAttribute{Computed: true}, "link_speed": schema.StringAttribute{Computed: true}, "link_type": schema.StringAttribute{Computed: true}}}}, "link_source": schema.StringAttribute{MarkdownDescription: "Link sources", Optional: true}, "link_speed": schema.StringAttribute{MarkdownDescription: "Link speeds", Optional: true}, "link_type": schema.StringAttribute{MarkdownDescription: "Types of links", Optional: true}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned. Defaults to (1:100) to prevent reading entire DB by accident", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetTopovizLinksDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetTopovizLinksDataSourceModel
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
func (d *GetTopovizLinksDataSource) readListRemote(ctx context.Context, config *GetTopovizLinksDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/topoviz/links"
	params := url.Values{}
	if !config.Endpoint1HostName.IsNull() {
		params.Set("endpoint1HostName", config.Endpoint1HostName.ValueString())
	}
	if !config.Endpoint2HostName.IsNull() {
		params.Set("endpoint2HostName", config.Endpoint2HostName.ValueString())
	}
	if !config.Endpoint1Alias.IsNull() {
		params.Set("endpoint1Alias", config.Endpoint1Alias.ValueString())
	}
	if !config.Endpoint2Alias.IsNull() {
		params.Set("endpoint2Alias", config.Endpoint2Alias.ValueString())
	}
	if !config.LinkType.IsNull() {
		params.Set("linkType", config.LinkType.ValueString())
	}
	if !config.LinkSource.IsNull() {
		params.Set("linkSource", config.LinkSource.ValueString())
	}
	if !config.LinkSpeed.IsNull() {
		params.Set("linkSpeed", config.LinkSpeed.ValueString())
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_topoviz_links", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageItems := []any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageItems); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_topoviz_links", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_topoviz_links", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetTopovizLinksDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
