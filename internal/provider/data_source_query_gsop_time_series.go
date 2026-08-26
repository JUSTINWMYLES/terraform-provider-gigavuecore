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
	_ datasource.DataSource              = (*QueryGsopTimeSeriesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*QueryGsopTimeSeriesDataSource)(nil)
)

// QueryGsopTimeSeriesDataSource is the generated Terraform data source implementation.
type QueryGsopTimeSeriesDataSource struct {
	client *client.Client
}

// QueryGsopTimeSeriesDataSourceModel describes the data source state shape.
type QueryGsopTimeSeriesDataSourceModel struct {
	Cluster    types.String `tfsdk:"cluster"`
	EndTime    types.String `tfsdk:"end_time" json:"endTime"`
	GsGroup    types.String `tfsdk:"gs_group" json:"gsGroup"`
	Gsop       types.String `tfsdk:"gsop"`
	MaxPoints  types.Int64  `tfsdk:"max_points" json:"maxPoints"`
	MaxSeries  types.Int64  `tfsdk:"max_series" json:"maxSeries"`
	Metric     types.String `tfsdk:"metric"`
	Rate       types.Bool   `tfsdk:"rate"`
	Series     types.List   `tfsdk:"series"`
	Since      types.String `tfsdk:"since"`
	SortBy     types.String `tfsdk:"sort_by" json:"sortBy"`
	StartTime  types.String `tfsdk:"start_time" json:"startTime"`
	TrafficDir types.String `tfsdk:"traffic_dir" json:"trafficDir"`
	Type       types.String `tfsdk:"type"`
}

// NewQueryGsopTimeSeriesDataSource returns a new instance of the generated data source.
func NewQueryGsopTimeSeriesDataSource() datasource.DataSource {
	return &QueryGsopTimeSeriesDataSource{}
}

// Metadata returns the data source type name.
func (d *QueryGsopTimeSeriesDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_query_gsop_time_series"
}

// Schema returns the data source schema.
func (d *QueryGsopTimeSeriesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "GSOP Time Series Query interface", Attributes: map[string]schema.Attribute{"cluster": schema.StringAttribute{MarkdownDescription: "Cluster id to filter by. Value of '*' indicates a 'group-by-clusterId' request. Multiple values separated by '|' can be specified, in which case time series are filtered and grouped by the requested clusters", Optional: true}, "end_time": schema.StringAttribute{MarkdownDescription: "End Time in ISO 8601 format", Computed: true, Optional: true}, "gs_group": schema.StringAttribute{MarkdownDescription: "GsGroup alias to filter by. Value of '*' indicates a 'group-by-GsGroup' request. Multiple values separated by '|' can be specified, in which case time series are filtered and grouped by the requested GsGroup", Optional: true}, "gsop": schema.StringAttribute{MarkdownDescription: "GSOP alias to filter by. Value of '*' indicates a 'group-by-gsop' request. Multiple values separated by '|' can be specified, in which case time series are filtered and grouped by the requested GSOP", Optional: true}, "max_points": schema.Int64Attribute{MarkdownDescription: "Number of the data points to return for the requested time period. Useful for building charts. Default value of 0 will return every datapoint, which could be costly for extensive time periods and/or high data sampling rates", Optional: true}, "max_series": schema.Int64Attribute{MarkdownDescription: "In Conjunction with the 'sortBy', returns the first N elements of the sorted time series list. Useful for the 'top-N' queries", Optional: true}, "metric": schema.StringAttribute{MarkdownDescription: "Time Series metric type", Required: true}, "rate": schema.BoolAttribute{MarkdownDescription: "Controls whether rate calculation should be done on the stored raw data before results are returned", Optional: true}, "series": schema.ListNestedAttribute{MarkdownDescription: "time series", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"context": schema.ListNestedAttribute{MarkdownDescription: "Time Series query filtering values", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"name": schema.StringAttribute{MarkdownDescription: "Time-series filtering/grouping tag", Computed: true}, "value": schema.StringAttribute{MarkdownDescription: "tag value", Computed: true}}}}, "data_points": schema.ListNestedAttribute{MarkdownDescription: "Time Series timestamped data points", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"ts_utc": schema.Int64Attribute{MarkdownDescription: "Timestamp in UTC milliseconds", Computed: true}, "value": schema.Int64Attribute{MarkdownDescription: "time series value", Computed: true}}}}}}}, "since": schema.StringAttribute{MarkdownDescription: "Time Series start time reference in the past relative to 'now'. In the {number}-{timeUnit} format, where {timeUnit} is one of: ['minute', 'hour', 'day', 'week', 'month']. (Ex: '3-hour'). Mutually exclusive with and takes Precedence over the 'startTime' and 'endTime' attributes", Optional: true}, "sort_by": schema.StringAttribute{MarkdownDescription: "If specified, each resulted timeseries over the specified time period is averaged to a single value, and then sorted in the requested order. Useful for the 'top-N' queries", Optional: true}, "start_time": schema.StringAttribute{MarkdownDescription: "Start Time in ISO 8601 format", Computed: true, Optional: true}, "traffic_dir": schema.StringAttribute{MarkdownDescription: "Traffic direction to filter by. Value of '*' indicates a 'group-by-traffic-direction' request", Optional: true}, "type": schema.StringAttribute{MarkdownDescription: "Time Series domain type", Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *QueryGsopTimeSeriesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config QueryGsopTimeSeriesDataSourceModel
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
func (d *QueryGsopTimeSeriesDataSource) readRemote(ctx context.Context, config *QueryGsopTimeSeriesDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/trending/gsops/timeSeries"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_query_gsop_time_series", fmt.Sprintf("Could not build request: %s", err))
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
	query.Set("metric", config.Metric.ValueString())
	if !config.Rate.IsNull() {
		query.Set("rate", strconv.FormatBool(config.Rate.ValueBool()))
	}
	if !config.Cluster.IsNull() {
		query.Set("cluster", config.Cluster.ValueString())
	}
	if !config.GsGroup.IsNull() {
		query.Set("gsGroup", config.GsGroup.ValueString())
	}
	if !config.Gsop.IsNull() {
		query.Set("gsop", config.Gsop.ValueString())
	}
	if !config.TrafficDir.IsNull() {
		query.Set("trafficDir", config.TrafficDir.ValueString())
	}
	if !config.MaxPoints.IsNull() {
		query.Set("maxPoints", strconv.FormatInt(config.MaxPoints.ValueInt64(), 10))
	}
	if !config.SortBy.IsNull() {
		query.Set("sortBy", config.SortBy.ValueString())
	}
	if !config.MaxSeries.IsNull() {
		query.Set("maxSeries", strconv.FormatInt(config.MaxSeries.ValueInt64(), 10))
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_query_gsop_time_series", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_query_gsop_time_series", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_query_gsop_time_series", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_query_gsop_time_series", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_query_gsop_time_series", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_query_gsop_time_series", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_query_gsop_time_series", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_query_gsop_time_series", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_query_gsop_time_series", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_query_gsop_time_series", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_query_gsop_time_series", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_query_gsop_time_series", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["timeSeries"]; ok {
		if m, ok := v.(map[string]any); ok {
			data = m
		} else if arr, ok := v.([]any); ok && len(arr) > 0 {
			if m, ok := arr[0].(map[string]any); ok {
				data = m
			}
		}
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_query_gsop_time_series", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *QueryGsopTimeSeriesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
