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
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*LoadTopologyVizEndPointsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadTopologyVizEndPointsDataSource)(nil)
)

// LoadTopologyVizEndPointsDataSource is the generated Terraform data source implementation.
type LoadTopologyVizEndPointsDataSource struct {
	client *client.Client
}

// LoadTopologyVizEndPointsDataSourceModel describes the data source state shape.
type LoadTopologyVizEndPointsDataSourceModel struct {
	ClusterId  types.String `tfsdk:"cluster_id" json:"clusterId"`
	Hostname   types.String `tfsdk:"hostname"`
	IgnoreUsed types.Bool   `tfsdk:"ignore_used" json:"ignoreUsed"`
	Items      types.List   `tfsdk:"items"`
	Type       types.String `tfsdk:"type"`
}

// NewLoadTopologyVizEndPointsDataSource returns a new instance of the generated data source.
func NewLoadTopologyVizEndPointsDataSource() datasource.DataSource {
	return &LoadTopologyVizEndPointsDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadTopologyVizEndPointsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_topology_viz_end_points"
}

// Schema returns the data source schema.
func (d *LoadTopologyVizEndPointsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load ports and gigaStream information required for manual link creation", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Cluster Name", Required: true}, "hostname": schema.StringAttribute{MarkdownDescription: "Node hostname", Required: true}, "ignore_used": schema.BoolAttribute{MarkdownDescription: "Set this to 'true' to display endpoints that are not part of anyother link", Optional: true}, "items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "PortId if endpointType is port, alias if it is a gigastream", Computed: true}, "endpoint_type": schema.StringAttribute{Computed: true}, "port_type": schema.StringAttribute{Computed: true}, "ports": schema.ListAttribute{MarkdownDescription: "List of portIds in the gigastream", Computed: true, ElementType: types.StringType}}}}, "type": schema.StringAttribute{MarkdownDescription: "EndPoint type", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadTopologyVizEndPointsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadTopologyVizEndPointsDataSourceModel
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
func (d *LoadTopologyVizEndPointsDataSource) readListRemote(ctx context.Context, config *LoadTopologyVizEndPointsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/topoviz/endPoints"
	params := url.Values{}
	params.Set("clusterId", config.ClusterId.ValueString())
	params.Set("hostname", config.Hostname.ValueString())
	if !config.IgnoreUsed.IsNull() {
		params.Set("ignoreUsed", strconv.FormatBool(config.IgnoreUsed.ValueBool()))
	}
	if !config.Type.IsNull() {
		params.Set("type", config.Type.ValueString())
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
		resp.Diagnostics.AddError("Error reading gigavuecore_load_topology_viz_end_points", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageItems := []any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageItems); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_topology_viz_end_points", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_topology_viz_end_points", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadTopologyVizEndPointsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
