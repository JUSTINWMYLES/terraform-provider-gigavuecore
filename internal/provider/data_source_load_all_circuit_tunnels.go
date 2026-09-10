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
	_ datasource.DataSource              = (*LoadAllCircuitTunnelsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadAllCircuitTunnelsDataSource)(nil)
)

// LoadAllCircuitTunnelsDataSource is the generated Terraform data source implementation.
type LoadAllCircuitTunnelsDataSource struct {
	client *client.Client
}

// LoadAllCircuitTunnelsDataSourceModel describes the data source state shape.
type LoadAllCircuitTunnelsDataSourceModel struct {
	ClusterId types.String `tfsdk:"cluster_id" json:"clusterId"`
	Items     types.List   `tfsdk:"items"`
	Mode      types.String `tfsdk:"mode"`
	Page      types.String `tfsdk:"page"`
	Sort      types.String `tfsdk:"sort"`
	Type      types.String `tfsdk:"type"`
}

// NewLoadAllCircuitTunnelsDataSource returns a new instance of the generated data source.
func NewLoadAllCircuitTunnelsDataSource() datasource.DataSource {
	return &LoadAllCircuitTunnelsDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadAllCircuitTunnelsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_all_circuit_tunnels"
}

// Schema returns the data source schema.
func (d *LoadAllCircuitTunnelsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load all Circuit Tunnels", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "tunnel name", Computed: true}, "attach": schema.ListAttribute{MarkdownDescription: "ipInterface if type is vxlan. ports or gigastreams if type is circuit and is valid only with decap mode.", Computed: true, ElementType: types.StringType}, "circuit_ids": schema.ListAttribute{MarkdownDescription: "circuit ids, valid and required if type is circuit", Computed: true, ElementType: types.Int64Type}, "cluster_id": schema.StringAttribute{MarkdownDescription: "id of the cluster in which this circuit tunnel is created", Computed: true}, "comment": schema.StringAttribute{Computed: true}, "dip_address": schema.StringAttribute{Computed: true}, "l4_src_port": schema.Int64Attribute{Computed: true}, "mode": schema.StringAttribute{MarkdownDescription: "tunnel mode", Computed: true}, "type": schema.StringAttribute{Computed: true}}}}, "mode": schema.StringAttribute{MarkdownDescription: "Filter circuit tunnels by mode", Optional: true}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}, "type": schema.StringAttribute{MarkdownDescription: "Filter circuit tunnels by type", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadAllCircuitTunnelsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadAllCircuitTunnelsDataSourceModel
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
func (d *LoadAllCircuitTunnelsDataSource) readListRemote(ctx context.Context, config *LoadAllCircuitTunnelsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/circuitTunnels"
	params := url.Values{}
	params.Set("clusterId", config.ClusterId.ValueString())
	if !config.Mode.IsNull() {
		params.Set("mode", config.Mode.ValueString())
	}
	if !config.Type.IsNull() {
		params.Set("type", config.Type.ValueString())
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
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_circuit_tunnels", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_circuit_tunnels", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["circuitTunnels"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_circuit_tunnels", fmt.Sprintf("Could not decode list page: missing %q array", "circuitTunnels"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_circuit_tunnels", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadAllCircuitTunnelsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
