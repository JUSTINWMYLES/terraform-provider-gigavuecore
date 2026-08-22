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
	_ datasource.DataSource              = (*QueryDbForTrafficFlowsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*QueryDbForTrafficFlowsDataSource)(nil)
)

// QueryDbForTrafficFlowsDataSource is the generated Terraform data source implementation.
type QueryDbForTrafficFlowsDataSource struct {
	client *client.Client
}

// QueryDbForTrafficFlowsDataSourceModel describes the data source state shape.
type QueryDbForTrafficFlowsDataSourceModel struct {
	EndTime   types.String `tfsdk:"end_time" json:"endTime"`
	Map       types.String `tfsdk:"map"`
	MaxPoints types.String `tfsdk:"max_points" json:"maxPoints"`
	MaxSeries types.String `tfsdk:"max_series" json:"maxSeries"`
	Metric    types.String `tfsdk:"metric"`
	Rate      types.Bool   `tfsdk:"rate"`
	Series    types.List   `tfsdk:"series"`
	Since     types.String `tfsdk:"since"`
	SortBy    types.String `tfsdk:"sort_by" json:"sortBy"`
	StartTime types.String `tfsdk:"start_time" json:"startTime"`
	Type      types.String `tfsdk:"type"`
}

// NewQueryDbForTrafficFlowsDataSource returns a new instance of the generated data source.
func NewQueryDbForTrafficFlowsDataSource() datasource.DataSource {
	return &QueryDbForTrafficFlowsDataSource{}
}

// Metadata returns the data source type name.
func (d *QueryDbForTrafficFlowsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_query_db_for_traffic_flows"
}

// Schema returns the data source schema.
func (d *QueryDbForTrafficFlowsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Query time series traffic flow statistics", Attributes: map[string]schema.Attribute{"end_time": schema.StringAttribute{MarkdownDescription: "End Time in ISO 8601 format", Computed: true, Optional: true}, "map": schema.StringAttribute{MarkdownDescription: "Traffic flow alias (required)", Required: true}, "max_points": schema.StringAttribute{MarkdownDescription: "Maximum number of data points to return", Optional: true}, "max_series": schema.StringAttribute{MarkdownDescription: "Maximum number of series to return", Optional: true}, "metric": schema.StringAttribute{MarkdownDescription: "Time Series metric type", Computed: true, Optional: true}, "rate": schema.BoolAttribute{MarkdownDescription: "Return results as rate (true) or absolute values (false)", Optional: true}, "series": schema.ListNestedAttribute{MarkdownDescription: "time series", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"context": schema.ListNestedAttribute{MarkdownDescription: "Time Series query filtering values", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"name": schema.StringAttribute{MarkdownDescription: "Time-series filtering/grouping tag", Computed: true}, "value": schema.StringAttribute{MarkdownDescription: "tag value", Computed: true}}}}, "data_points": schema.ListNestedAttribute{MarkdownDescription: "Time Series timestamped data points", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"ts_utc": schema.Int64Attribute{MarkdownDescription: "Timestamp in UTC milliseconds", Computed: true}, "value": schema.Int64Attribute{MarkdownDescription: "time series value", Computed: true}}}}}}}, "since": schema.StringAttribute{MarkdownDescription: "Relative time window (e.g., 1h, 24h) to query from", Optional: true}, "sort_by": schema.StringAttribute{MarkdownDescription: "Field to sort the results by", Optional: true}, "start_time": schema.StringAttribute{MarkdownDescription: "Start Time in ISO 8601 format", Computed: true, Optional: true}, "type": schema.StringAttribute{MarkdownDescription: "Time Series domain type", Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *QueryDbForTrafficFlowsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config QueryDbForTrafficFlowsDataSourceModel
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
func (d *QueryDbForTrafficFlowsDataSource) readRemote(ctx context.Context, config *QueryDbForTrafficFlowsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/trending/traffic-flows/timeSeries"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_query_db_for_traffic_flows", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.Since.IsNull() {
		query.Set("since", config.Since.ValueString())
	}
	if !config.StartTime.IsNull() {
		query.Set("startTime", config.StartTime.ValueString())
	}
	if !config.EndTime.IsNull() {
		query.Set("endTime", config.EndTime.ValueString())
	}
	if !config.Metric.IsNull() {
		query.Set("metric", config.Metric.ValueString())
	}
	query.Set("map", config.Map.ValueString())
	if !config.MaxPoints.IsNull() {
		query.Set("maxPoints", config.MaxPoints.ValueString())
	}
	if !config.SortBy.IsNull() {
		query.Set("sortBy", config.SortBy.ValueString())
	}
	if !config.MaxSeries.IsNull() {
		query.Set("maxSeries", config.MaxSeries.ValueString())
	}
	if !config.Rate.IsNull() {
		query.Set("rate", strconv.FormatBool(config.Rate.ValueBool()))
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_query_db_for_traffic_flows", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_query_db_for_traffic_flows", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_query_db_for_traffic_flows", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_query_db_for_traffic_flows", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_query_db_for_traffic_flows", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_query_db_for_traffic_flows", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_query_db_for_traffic_flows", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_query_db_for_traffic_flows", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_query_db_for_traffic_flows", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_query_db_for_traffic_flows", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_query_db_for_traffic_flows", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_query_db_for_traffic_flows", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["timeSeries"]; ok {
		if m, ok := v.(map[string]any); ok {
			data = m
		}
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_query_db_for_traffic_flows", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *QueryDbForTrafficFlowsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
