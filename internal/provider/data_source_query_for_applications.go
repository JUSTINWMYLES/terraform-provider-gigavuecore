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
	_ datasource.DataSource              = (*QueryForApplicationsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*QueryForApplicationsDataSource)(nil)
)

// QueryForApplicationsDataSource is the generated Terraform data source implementation.
type QueryForApplicationsDataSource struct {
	client *client.Client
}

// QueryForApplicationsDataSourceModel describes the data source state shape.
type QueryForApplicationsDataSourceModel struct {
	EndTime                       types.String `tfsdk:"end_time" json:"endTime"`
	FilterField                   types.String `tfsdk:"filter_field" json:"filterField"`
	GroupByClause                 types.String `tfsdk:"group_by_clause" json:"groupByClause"`
	Items                         types.List   `tfsdk:"items"`
	Metric                        types.String `tfsdk:"metric"`
	MonitorTags                   types.String `tfsdk:"monitor_tags" json:"monitor.tags"`
	OrderByDesc                   types.Bool   `tfsdk:"order_by_desc" json:"orderByDesc"`
	Rate                          types.Bool   `tfsdk:"rate"`
	SolutionCreatedTimestamp      types.String `tfsdk:"solution_created_timestamp" json:"solutionCreatedTimestamp"`
	SolutionMonitorIntervalInMins types.String `tfsdk:"solution_monitor_interval_in_mins" json:"solutionMonitorIntervalInMins"`
	StartTime                     types.String `tfsdk:"start_time" json:"startTime"`
	Tags                          types.String `tfsdk:"tags"`
	Time                          types.String `tfsdk:"time"`
	TimeUnit                      types.String `tfsdk:"time_unit" json:"timeUnit"`
	Top                           types.String `tfsdk:"top"`
}

// NewQueryForApplicationsDataSource returns a new instance of the generated data source.
func NewQueryForApplicationsDataSource() datasource.DataSource {
	return &QueryForApplicationsDataSource{}
}

// Metadata returns the data source type name.
func (d *QueryForApplicationsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_query_for_applications"
}

// Schema returns the data source schema.
func (d *QueryForApplicationsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Query application visibility time series statistics", Attributes: map[string]schema.Attribute{"end_time": schema.StringAttribute{MarkdownDescription: "End time for the query (ISO 8601 or epoch)", Optional: true}, "filter_field": schema.StringAttribute{MarkdownDescription: "Field to filter results by", Optional: true}, "group_by_clause": schema.StringAttribute{MarkdownDescription: "Group by clause for aggregation", Optional: true}, "items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"time_series": schema.SingleNestedAttribute{MarkdownDescription: "Fabric Map Time Series Group", Computed: true, Attributes: map[string]schema.Attribute{"end_time": schema.StringAttribute{MarkdownDescription: "End Time in ISO 8601 format", Computed: true}, "metric": schema.StringAttribute{MarkdownDescription: "Time Series metric type", Computed: true}, "series": schema.ListNestedAttribute{MarkdownDescription: "time series", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"context": schema.ListNestedAttribute{MarkdownDescription: "Time Series query filtering values", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"name": schema.StringAttribute{MarkdownDescription: "Time-series filtering/grouping tag", Computed: true}, "value": schema.StringAttribute{MarkdownDescription: "tag value", Computed: true}}}}, "data_points": schema.ListNestedAttribute{MarkdownDescription: "Time Series timestamped data points", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"ts_utc": schema.Int64Attribute{MarkdownDescription: "Timestamp in UTC milliseconds", Computed: true}, "value": schema.Int64Attribute{MarkdownDescription: "time series value", Computed: true}}}}}}}, "start_time": schema.StringAttribute{MarkdownDescription: "Start Time in ISO 8601 format", Computed: true}, "type": schema.StringAttribute{MarkdownDescription: "Time Series domain type", Computed: true}}}}}}, "metric": schema.StringAttribute{MarkdownDescription: "Metric to query (e.g., application bytes, flows)", Optional: true}, "monitor_tags": schema.StringAttribute{MarkdownDescription: "Monitor tags for filtering", Optional: true}, "order_by_desc": schema.BoolAttribute{MarkdownDescription: "Order results in descending order", Optional: true}, "rate": schema.BoolAttribute{MarkdownDescription: "Return results as rate (true) or absolute values (false)", Optional: true}, "solution_created_timestamp": schema.StringAttribute{MarkdownDescription: "Solution creation timestamp", Optional: true}, "solution_monitor_interval_in_mins": schema.StringAttribute{MarkdownDescription: "Solution monitor interval in minutes", Optional: true}, "start_time": schema.StringAttribute{MarkdownDescription: "Start time for the query (ISO 8601 or epoch)", Optional: true}, "tags": schema.StringAttribute{MarkdownDescription: "Tags to filter the query", Optional: true}, "time": schema.StringAttribute{MarkdownDescription: "Time window for the query", Optional: true}, "time_unit": schema.StringAttribute{MarkdownDescription: "Time unit for the query (e.g., minutes, hours)", Optional: true}, "top": schema.StringAttribute{MarkdownDescription: "Number of top results to return", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *QueryForApplicationsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config QueryForApplicationsDataSourceModel
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
func (d *QueryForApplicationsDataSource) readListRemote(ctx context.Context, config *QueryForApplicationsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/trending/traffic-flows/applicationVisibility/timeSeries"
	params := url.Values{}
	if !config.Metric.IsNull() {
		params.Set("metric", config.Metric.ValueString())
	}
	if !config.GroupByClause.IsNull() {
		params.Set("groupByClause", config.GroupByClause.ValueString())
	}
	if !config.FilterField.IsNull() {
		params.Set("filterField", config.FilterField.ValueString())
	}
	if !config.OrderByDesc.IsNull() {
		params.Set("orderByDesc", strconv.FormatBool(config.OrderByDesc.ValueBool()))
	}
	if !config.Tags.IsNull() {
		params.Set("tags", config.Tags.ValueString())
	}
	if !config.Rate.IsNull() {
		params.Set("rate", strconv.FormatBool(config.Rate.ValueBool()))
	}
	if !config.Time.IsNull() {
		params.Set("time", config.Time.ValueString())
	}
	if !config.SolutionMonitorIntervalInMins.IsNull() {
		params.Set("solutionMonitorIntervalInMins", config.SolutionMonitorIntervalInMins.ValueString())
	}
	if !config.Top.IsNull() {
		params.Set("top", config.Top.ValueString())
	}
	if !config.TimeUnit.IsNull() {
		params.Set("timeUnit", config.TimeUnit.ValueString())
	}
	if !config.StartTime.IsNull() {
		params.Set("startTime", config.StartTime.ValueString())
	}
	if !config.EndTime.IsNull() {
		params.Set("endTime", config.EndTime.ValueString())
	}
	if !config.SolutionCreatedTimestamp.IsNull() {
		params.Set("solutionCreatedTimestamp", config.SolutionCreatedTimestamp.ValueString())
	}
	if !config.MonitorTags.IsNull() {
		params.Set("monitor.tags", config.MonitorTags.ValueString())
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
		resp.Diagnostics.AddError("Error reading gigavuecore_query_for_applications", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageItems := []any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageItems); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_query_for_applications", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_query_for_applications", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *QueryForApplicationsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
