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
	_ datasource.DataSource              = (*LoadAllNetflowMonitorsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadAllNetflowMonitorsDataSource)(nil)
)

// LoadAllNetflowMonitorsDataSource is the generated Terraform data source implementation.
type LoadAllNetflowMonitorsDataSource struct {
	client *client.Client
}

// LoadAllNetflowMonitorsDataSourceModel describes the data source state shape.
type LoadAllNetflowMonitorsDataSourceModel struct {
	ClusterId  types.String `tfsdk:"cluster_id" json:"clusterId"`
	Context    types.Object `tfsdk:"context"`
	NfMonitors types.List   `tfsdk:"nf_monitors" json:"nfMonitors"`
	Page       types.String `tfsdk:"page"`
	Sort       types.String `tfsdk:"sort"`
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
	resp.Schema = schema.Schema{MarkdownDescription: "Load all Netwflow Monitors", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "context": schema.SingleNestedAttribute{MarkdownDescription: "Gigamon query result context", Computed: true, Attributes: map[string]schema.Attribute{"page_no": schema.Int64Attribute{MarkdownDescription: "page number of the returned result set", Computed: true}, "page_size": schema.Int64Attribute{MarkdownDescription: "page size of the returned result set", Computed: true}, "sort": schema.ListAttribute{MarkdownDescription: "sorting info of the returned result set. list of fields in the array indicate sorting order", Computed: true, ElementType: types.StringType}, "total_items": schema.Int64Attribute{MarkdownDescription: "total number of items in the queried entity type", Computed: true}}}, "nf_monitors": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{Computed: true}, "cache": schema.SingleNestedAttribute{MarkdownDescription: "Netflow Monitor Cache", Computed: true, Attributes: map[string]schema.Attribute{"export_triggers": schema.SingleNestedAttribute{MarkdownDescription: "Netflow Monitor Cache Export Triggers", Computed: true, Attributes: map[string]schema.Attribute{"event": schema.StringAttribute{Computed: true}, "timeout_active": schema.Int64Attribute{MarkdownDescription: "in seconds. max value is 7 days. default is 30 min", Computed: true}, "timeout_inactive": schema.Int64Attribute{MarkdownDescription: "in seconds. max value is 7 days. default is 15 sec", Computed: true}}}, "type": schema.StringAttribute{Computed: true}}}, "cluster_id": schema.StringAttribute{MarkdownDescription: "id of the defining cluster", Computed: true}, "description": schema.StringAttribute{Computed: true}, "records": schema.SetAttribute{MarkdownDescription: "Aliases of referenced Netflow Records defined on the GsGroup. Up to 5 records.", Computed: true, ElementType: types.StringType}, "sampling": schema.SingleNestedAttribute{MarkdownDescription: "monitor sampling", Computed: true, Attributes: map[string]schema.Attribute{"mode": schema.StringAttribute{Computed: true}, "single_sampling_rate": schema.Int64Attribute{MarkdownDescription: "Packet interval window size. Valid values: 10-16000 (in packets)", Computed: true}}}, "sampling_space": schema.Int64Attribute{MarkdownDescription: "DEPRECATED: use 'sampling'", Computed: true}, "ssl_port_restrictions": schema.SingleNestedAttribute{MarkdownDescription: "Port restrictions for Netflow/SSL sessions", Computed: true, Attributes: map[string]schema.Attribute{"ports": schema.ListAttribute{MarkdownDescription: "The list of TCP ports whose packets will be sent to the SSL module. Not valid if sslPort is 'all'. if 'wellKnownPorts' is selected for 'sslPorts' ports will be [993, 995, 465, 636, 563, 443]", Computed: true, ElementType: types.Int64Type}, "ssl_ports": schema.StringAttribute{MarkdownDescription: "Ports whose packets will be sent to the SSL module. 'all' ports; 'wellKnownPorts' is [993, 995, 465, 636, 563, 443]; or 'ports' - specify upto 10 ports", Computed: true}}}}}}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadAllNetflowMonitorsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadAllNetflowMonitorsDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	d.readRemote(ctx, &config, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	if config.Page.IsNull() {
		resp.Diagnostics.AddWarning("Single-page result", fmt.Sprintf("This data source reads a single page of a paginated API endpoint and does not aggregate results across pages. The \"page\" argument is unset, so the default page is returned; set it to retrieve a different page."))
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}

// readRemote performs the read HTTP exchange and decodes the response into config. Extracted from Read so the request/response logic is unit-testable without a tfsdk.Config.
func (d *LoadAllNetflowMonitorsDataSource) readRemote(ctx context.Context, config *LoadAllNetflowMonitorsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/netflow/monitors"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_netflow_monitors", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	if !config.Page.IsNull() {
		query.Set("page", config.Page.ValueString())
	}
	if !config.Sort.IsNull() {
		query.Set("sort", config.Sort.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_netflow_monitors", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_netflow_monitors", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_netflow_monitors", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_netflow_monitors", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_netflow_monitors", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_netflow_monitors", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_netflow_monitors", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_netflow_monitors", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_netflow_monitors", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_load_all_netflow_monitors", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_netflow_monitors", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_netflow_monitors", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
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
