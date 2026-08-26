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
	_ datasource.DataSource              = (*LoadClusterPortFilterCountersDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadClusterPortFilterCountersDataSource)(nil)
)

// LoadClusterPortFilterCountersDataSource is the generated Terraform data source implementation.
type LoadClusterPortFilterCountersDataSource struct {
	client *client.Client
}

// LoadClusterPortFilterCountersDataSourceModel describes the data source state shape.
type LoadClusterPortFilterCountersDataSourceModel struct {
	ClusterId       types.String `tfsdk:"cluster_id" json:"clusterId"`
	PortFilterStats types.Object `tfsdk:"port_filter_stats" json:"portFilterStats"`
	PortId          types.String `tfsdk:"port_id" json:"portId"`
	SysUpTime       types.Int64  `tfsdk:"sys_up_time" json:"sysUpTime"`
}

// NewLoadClusterPortFilterCountersDataSource returns a new instance of the generated data source.
func NewLoadClusterPortFilterCountersDataSource() datasource.DataSource {
	return &LoadClusterPortFilterCountersDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadClusterPortFilterCountersDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_cluster_port_filter_counters"
}

// Schema returns the data source schema.
func (d *LoadClusterPortFilterCountersDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load Cluster Port Filter Counters", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "port_filter_stats": schema.SingleNestedAttribute{MarkdownDescription: "Egress Port Filter counters", Computed: true, Attributes: map[string]schema.Attribute{"port": schema.StringAttribute{Computed: true}, "rules": schema.SingleNestedAttribute{MarkdownDescription: "Port Filter Rules Container", Computed: true, Attributes: map[string]schema.Attribute{"drop_rules": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"bytes": schema.Int64Attribute{Computed: true}, "filters": schema.ListNestedAttribute{MarkdownDescription: "list of the filters configured for rule", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"filter_type": schema.StringAttribute{Computed: true}, "filter_value": schema.StringAttribute{Computed: true}}}}, "pkts": schema.Int64Attribute{Computed: true}, "rule_id": schema.Int64Attribute{Computed: true}}}}, "pass_rules": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"bytes": schema.Int64Attribute{Computed: true}, "filters": schema.ListNestedAttribute{MarkdownDescription: "list of the filters configured for rule", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"filter_type": schema.StringAttribute{Computed: true}, "filter_value": schema.StringAttribute{Computed: true}}}}, "pkts": schema.Int64Attribute{Computed: true}, "rule_id": schema.Int64Attribute{Computed: true}}}}}}}}, "port_id": schema.StringAttribute{MarkdownDescription: "port id to filter by", Required: true}, "sys_up_time": schema.Int64Attribute{MarkdownDescription: "Device sysUpTime. The time (in hundredths of a second) since the network management portion of the system was last re-initialized", Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadClusterPortFilterCountersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadClusterPortFilterCountersDataSourceModel
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
func (d *LoadClusterPortFilterCountersDataSource) readRemote(ctx context.Context, config *LoadClusterPortFilterCountersDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/nodeCounters/portFilter"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_cluster_port_filter_counters", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	query.Set("portId", config.PortId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_cluster_port_filter_counters", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_cluster_port_filter_counters", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_cluster_port_filter_counters", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_cluster_port_filter_counters", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_cluster_port_filter_counters", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_cluster_port_filter_counters", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_cluster_port_filter_counters", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_cluster_port_filter_counters", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_cluster_port_filter_counters", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_load_cluster_port_filter_counters", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_load_cluster_port_filter_counters", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_cluster_port_filter_counters", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_cluster_port_filter_counters", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadClusterPortFilterCountersDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
