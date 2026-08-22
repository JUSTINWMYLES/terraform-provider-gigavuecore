package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*GetAllTrafficFlowsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAllTrafficFlowsDataSource)(nil)
)

// GetAllTrafficFlowsDataSource is the generated Terraform data source implementation.
type GetAllTrafficFlowsDataSource struct {
	client *client.Client
}

// GetAllTrafficFlowsDataSourceModel describes the data source state shape.
type GetAllTrafficFlowsDataSourceModel struct {
	Alias        types.String  `tfsdk:"alias"`
	ConfigStatus types.String  `tfsdk:"config_status" json:"configStatus"`
	DstCluster   types.String  `tfsdk:"dst_cluster" json:"dstCluster"`
	DstPorts     types.String  `tfsdk:"dst_ports" json:"dstPorts"`
	HealthState  types.String  `tfsdk:"health_state" json:"healthState"`
	Page         types.String  `tfsdk:"page"`
	Select       types.String  `tfsdk:"select"`
	Sort         types.String  `tfsdk:"sort"`
	SrcCluster   types.String  `tfsdk:"src_cluster" json:"srcCluster"`
	SrcPorts     types.String  `tfsdk:"src_ports" json:"srcPorts"`
	Summary      types.Bool    `tfsdk:"summary"`
	Type         types.String  `tfsdk:"type"`
	Value        types.Dynamic `tfsdk:"value"`
}

// NewGetAllTrafficFlowsDataSource returns a new instance of the generated data source.
func NewGetAllTrafficFlowsDataSource() datasource.DataSource {
	return &GetAllTrafficFlowsDataSource{}
}

// Metadata returns the data source type name.
func (d *GetAllTrafficFlowsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_traffic_flows"
}

// Schema returns the data source schema.
func (d *GetAllTrafficFlowsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "List all Traffic Flows with optional filters", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Traffic Flows alias filter", Optional: true}, "config_status": schema.StringAttribute{MarkdownDescription: "Configuration status filter", Optional: true}, "dst_cluster": schema.StringAttribute{MarkdownDescription: "Destination cluster filter", Optional: true}, "dst_ports": schema.StringAttribute{MarkdownDescription: "Destination port(s) filter", Optional: true}, "health_state": schema.StringAttribute{MarkdownDescription: "Health state filter", Optional: true}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "select": schema.StringAttribute{MarkdownDescription: "Comma-separated list of fields to select", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC and default sort field is fabric map alias. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}, "src_cluster": schema.StringAttribute{MarkdownDescription: "Source cluster filter", Optional: true}, "src_ports": schema.StringAttribute{MarkdownDescription: "Source port(s) filter", Optional: true}, "summary": schema.BoolAttribute{MarkdownDescription: "Return summary data", Optional: true}, "type": schema.StringAttribute{MarkdownDescription: "Traffic Flows type filter", Optional: true}, "value": schema.DynamicAttribute{Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetAllTrafficFlowsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAllTrafficFlowsDataSourceModel
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
func (d *GetAllTrafficFlowsDataSource) readRemote(ctx context.Context, config *GetAllTrafficFlowsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/traffic-flows"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_traffic_flows", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.SrcPorts.IsNull() {
		query.Set("srcPorts", config.SrcPorts.ValueString())
	}
	if !config.DstPorts.IsNull() {
		query.Set("dstPorts", config.DstPorts.ValueString())
	}
	if !config.SrcCluster.IsNull() {
		query.Set("srcCluster", config.SrcCluster.ValueString())
	}
	if !config.DstCluster.IsNull() {
		query.Set("dstCluster", config.DstCluster.ValueString())
	}
	if !config.Alias.IsNull() {
		query.Set("alias", config.Alias.ValueString())
	}
	if !config.Type.IsNull() {
		query.Set("type", config.Type.ValueString())
	}
	if !config.Select.IsNull() {
		query.Set("select", config.Select.ValueString())
	}
	if !config.ConfigStatus.IsNull() {
		query.Set("configStatus", config.ConfigStatus.ValueString())
	}
	if !config.HealthState.IsNull() {
		query.Set("healthState", config.HealthState.ValueString())
	}
	if !config.Summary.IsNull() {
		query.Set("summary", strconv.FormatBool(config.Summary.ValueBool()))
	}
	if !config.Page.IsNull() {
		query.Set("page", config.Page.ValueString())
	}
	if !config.Sort.IsNull() {
		query.Set("sort", config.Sort.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_traffic_flows", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_traffic_flows", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_traffic_flows", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_traffic_flows", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_traffic_flows", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_traffic_flows", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_traffic_flows", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_traffic_flows", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_traffic_flows", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_all_traffic_flows", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_traffic_flows", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_traffic_flows", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_traffic_flows", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetAllTrafficFlowsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
