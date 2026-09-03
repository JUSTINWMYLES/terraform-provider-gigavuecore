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
	_ datasource.DataSource              = (*LoadAllNetflowMonitorsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadAllNetflowMonitorsDataSource)(nil)
)

// LoadAllNetflowMonitorsDataSource is the generated Terraform data source implementation.
type LoadAllNetflowMonitorsDataSource struct {
	client *client.Client
}

// LoadAllNetflowMonitorsDataSourceModel describes the data source state shape.
type LoadAllNetflowMonitorsDataSourceModel struct {
	ClusterId types.String `tfsdk:"cluster_id" json:"clusterId"`
	Items     types.List   `tfsdk:"items"`
	Page      types.String `tfsdk:"page"`
	Sort      types.String `tfsdk:"sort"`
}

// NewLoadAllNetflowMonitorsDataSource returns a new instance of the generated data source.
func NewLoadAllNetflowMonitorsDataSource() datasource.DataSource {
	return &LoadAllNetflowMonitorsDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadAllNetflowMonitorsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_all_netflow_monitors"
}

// Schema returns the data source schema.
func (d *LoadAllNetflowMonitorsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load all Netwflow Monitors", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{Computed: true}, "cache": schema.SingleNestedAttribute{MarkdownDescription: "Netflow Monitor Cache", Computed: true, Attributes: map[string]schema.Attribute{"export_triggers": schema.SingleNestedAttribute{MarkdownDescription: "Netflow Monitor Cache Export Triggers", Computed: true, Attributes: map[string]schema.Attribute{"event": schema.StringAttribute{Computed: true}, "timeout_active": schema.Int64Attribute{MarkdownDescription: "in seconds. max value is 7 days. default is 30 min", Computed: true}, "timeout_inactive": schema.Int64Attribute{MarkdownDescription: "in seconds. max value is 7 days. default is 15 sec", Computed: true}}}, "type": schema.StringAttribute{Computed: true}}}, "cluster_id": schema.StringAttribute{MarkdownDescription: "id of the defining cluster", Computed: true}, "description": schema.StringAttribute{Computed: true}, "records": schema.SetAttribute{MarkdownDescription: "Aliases of referenced Netflow Records defined on the GsGroup. Up to 5 records.", Computed: true, ElementType: types.StringType}, "sampling": schema.SingleNestedAttribute{MarkdownDescription: "monitor sampling", Computed: true, Attributes: map[string]schema.Attribute{"mode": schema.StringAttribute{Computed: true}, "single_sampling_rate": schema.Int64Attribute{MarkdownDescription: "Packet interval window size. Valid values: 10-16000 (in packets)", Computed: true}}}, "sampling_space": schema.Int64Attribute{MarkdownDescription: "DEPRECATED: use 'sampling'", Computed: true}, "ssl_port_restrictions": schema.SingleNestedAttribute{MarkdownDescription: "Port restrictions for Netflow/SSL sessions", Computed: true, Attributes: map[string]schema.Attribute{"ports": schema.ListAttribute{MarkdownDescription: "The list of TCP ports whose packets will be sent to the SSL module. Not valid if sslPort is 'all'. if 'wellKnownPorts' is selected for 'sslPorts' ports will be [993, 995, 465, 636, 563, 443]", Computed: true, ElementType: types.Int64Type}, "ssl_ports": schema.StringAttribute{MarkdownDescription: "Ports whose packets will be sent to the SSL module. 'all' ports; 'wellKnownPorts' is [993, 995, 465, 636, 563, 443]; or 'ports' - specify upto 10 ports", Computed: true}}}}}}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadAllNetflowMonitorsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadAllNetflowMonitorsDataSourceModel
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
func (d *LoadAllNetflowMonitorsDataSource) readListRemote(ctx context.Context, config *LoadAllNetflowMonitorsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/netflow/monitors"
	params := url.Values{}
	params.Set("clusterId", config.ClusterId.ValueString())
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
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_netflow_monitors", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_netflow_monitors", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["nfMonitors"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_netflow_monitors", fmt.Sprintf("Could not decode list page: missing %q array", "nfMonitors"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_netflow_monitors", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadAllNetflowMonitorsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
