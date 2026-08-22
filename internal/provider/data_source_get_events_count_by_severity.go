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
	_ datasource.DataSource              = (*GetEventsCountBySeverityDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetEventsCountBySeverityDataSource)(nil)
)

// GetEventsCountBySeverityDataSource is the generated Terraform data source implementation.
type GetEventsCountBySeverityDataSource struct {
	client *client.Client
}

// GetEventsCountBySeverityDataSourceModel describes the data source state shape.
type GetEventsCountBySeverityDataSourceModel struct {
	AggGroups types.List   `tfsdk:"agg_groups" json:"aggGroups"`
	EndTime   types.String `tfsdk:"end_time" json:"endTime"`
	GroupBy   types.String `tfsdk:"group_by" json:"groupBy"`
	StartTime types.String `tfsdk:"start_time" json:"startTime"`
}

// NewGetEventsCountBySeverityDataSource returns a new instance of the generated data source.
func NewGetEventsCountBySeverityDataSource() datasource.DataSource {
	return &GetEventsCountBySeverityDataSource{}
}

// Metadata returns the data source type name.
func (d *GetEventsCountBySeverityDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_events_count_by_severity"
}

// Schema returns the data source schema.
func (d *GetEventsCountBySeverityDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get Counts of Events Grouped By Severity", Attributes: map[string]schema.Attribute{"agg_groups": schema.ListNestedAttribute{MarkdownDescription: "list of aggregate counts per groupBy type", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"count_": schema.Int64Attribute{MarkdownDescription: "the number of events in the aggregation group", Computed: true}, "group_name": schema.StringAttribute{MarkdownDescription: "Name of the group. For grouping based on severity, this field will have a value from the set ['info', 'minor', 'major', 'critical']", Computed: true}}}}, "end_time": schema.StringAttribute{MarkdownDescription: "Optionally specify the end time to filter the records by. In ISO 8601 format", Optional: true}, "group_by": schema.StringAttribute{MarkdownDescription: "aggregation grouping category", Required: true}, "start_time": schema.StringAttribute{MarkdownDescription: "Optionally specify the start time to filter the records by. In ISO 8601 format", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetEventsCountBySeverityDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetEventsCountBySeverityDataSourceModel
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
func (d *GetEventsCountBySeverityDataSource) readRemote(ctx context.Context, config *GetEventsCountBySeverityDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/events/stats/aggregate"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_events_count_by_severity", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.StartTime.IsNull() {
		query.Set("startTime", config.StartTime.ValueString())
	}
	if !config.EndTime.IsNull() {
		query.Set("endTime", config.EndTime.ValueString())
	}
	query.Set("groupBy", config.GroupBy.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_events_count_by_severity", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_events_count_by_severity", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_events_count_by_severity", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_events_count_by_severity", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_events_count_by_severity", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_events_count_by_severity", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_events_count_by_severity", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_events_count_by_severity", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_events_count_by_severity", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_events_count_by_severity", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_events_count_by_severity", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_events_count_by_severity", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["summary"]; ok {
		if m, ok := v.(map[string]any); ok {
			data = m
		}
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_events_count_by_severity", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetEventsCountBySeverityDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
