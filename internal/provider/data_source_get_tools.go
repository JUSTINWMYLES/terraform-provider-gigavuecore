package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*GetToolsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetToolsDataSource)(nil)
)

// GetToolsDataSource is the generated Terraform data source implementation.
type GetToolsDataSource struct {
	client *client.Client
}

// GetToolsDataSourceModel describes the data source state shape.
type GetToolsDataSourceModel struct {
	Alias     types.String `tfsdk:"alias"`
	NodeAlias types.String `tfsdk:"node_alias" json:"nodeAlias"`
	ToolsInfo types.List   `tfsdk:"tools_info" json:"toolsInfo"`
	Type      types.String `tfsdk:"type"`
}

// NewGetToolsDataSource returns a new instance of the generated data source.
func NewGetToolsDataSource() datasource.DataSource {
	return &GetToolsDataSource{}
}

// Metadata returns the data source type name.
func (d *GetToolsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_tools"
}

// Schema returns the data source schema.
func (d *GetToolsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get Tools", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{Computed: true}, "node_alias": schema.StringAttribute{MarkdownDescription: "manual node alias", Optional: true}, "tools_info": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"comment": schema.StringAttribute{Computed: true}, "compression_ratio": schema.Float64Attribute{Computed: true}, "giga_streams": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "gigastream alias. Uniquely identifies a gigastream within a cluster", Computed: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "id of the defining cluster", Computed: true}, "comment": schema.StringAttribute{Computed: true}, "drop_weight": schema.Int64Attribute{MarkdownDescription: "relative weight for dropping the traffic", Computed: true}, "failover_status": schema.StringAttribute{MarkdownDescription: "Failover Status", Computed: true}, "hash_size": schema.Int64Attribute{MarkdownDescription: "Hash bucket size", Computed: true}, "hash_tool_port": schema.ListNestedAttribute{MarkdownDescription: "Hash bucket id to tool port mapping", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"hash_bucket_ids": schema.ListAttribute{MarkdownDescription: "hash bucket id or range", Computed: true, ElementType: types.Int64Type}, "tool_ports": schema.ListAttribute{MarkdownDescription: "tool port(s) mapped to hashBucketIds", Computed: true, ElementType: types.StringType}}}}, "hash_type": schema.StringAttribute{Computed: true}, "hash_weights": schema.ListAttribute{MarkdownDescription: "hashWeights for 'ports'.If included, the list size must match the size of the 'ports' list", Computed: true, ElementType: types.Int64Type}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "health_state_reasons": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Computed: true}}}}, "ports": schema.ListAttribute{MarkdownDescription: "list of the ports to combine into a gigastream", Computed: true, ElementType: types.StringType}, "threshold_level": schema.StringAttribute{MarkdownDescription: "Threshold level", Computed: true}, "variance_threshold": schema.StringAttribute{MarkdownDescription: "Variance threshold percentage", Computed: true}}}}, "is_tool": schema.BoolAttribute{Computed: true}, "is_used_in_deployed_policy": schema.BoolAttribute{Computed: true}, "max_throughput": schema.Float64Attribute{Computed: true}, "model": schema.StringAttribute{Computed: true}, "node_alias": schema.StringAttribute{Computed: true}, "ports": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{Computed: true}, "cluster_id": schema.StringAttribute{Computed: true}, "device_ip": schema.StringAttribute{Computed: true}, "hostname": schema.StringAttribute{Computed: true}, "link_id": schema.StringAttribute{Computed: true}, "port_id": schema.StringAttribute{Computed: true}, "port_type": schema.StringAttribute{Computed: true}}}}, "topo_node_id": schema.StringAttribute{Computed: true}, "total_storage": schema.Float64Attribute{Computed: true}, "type": schema.StringAttribute{Computed: true}, "vendor": schema.StringAttribute{Computed: true}}}}, "type": schema.StringAttribute{Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetToolsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetToolsDataSourceModel
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
func (d *GetToolsDataSource) readRemote(ctx context.Context, config *GetToolsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/intent/policies/tools"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_tools", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.NodeAlias.IsNull() {
		query.Set("nodeAlias", config.NodeAlias.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_tools", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_tools", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		apiErr, err := client.NewAPIError(httpResp)
		if err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_tools", fmt.Sprintf("Could not read error response: %s", err))
			return
		}
		resp.Diagnostics.AddError("Error reading gigavuecore_get_tools", apiErr.Error())
		return
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_tools", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_tools", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetToolsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
