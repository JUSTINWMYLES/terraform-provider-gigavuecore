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
	_ datasource.DataSource              = (*GetCorrelatedAlarmsCountByParameterDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetCorrelatedAlarmsCountByParameterDataSource)(nil)
)

// GetCorrelatedAlarmsCountByParameterDataSource is the generated Terraform data source implementation.
type GetCorrelatedAlarmsCountByParameterDataSource struct {
	client *client.Client
}

// GetCorrelatedAlarmsCountByParameterDataSourceModel describes the data source state shape.
type GetCorrelatedAlarmsCountByParameterDataSourceModel struct {
	Acknowledged     types.String `tfsdk:"acknowledged"`
	Acknowledgedby   types.String `tfsdk:"acknowledgedby"`
	AggGroups        types.List   `tfsdk:"agg_groups" json:"aggGroups"`
	Alias            types.String `tfsdk:"alias"`
	ClusterId        types.String `tfsdk:"cluster_id" json:"clusterId"`
	DeviceIp         types.String `tfsdk:"device_ip" json:"deviceIp"`
	EndTime          types.String `tfsdk:"end_time" json:"endTime"`
	GroupBy          types.String `tfsdk:"group_by" json:"groupBy"`
	Hostname         types.String `tfsdk:"hostname"`
	ResourceId       types.String `tfsdk:"resource_id" json:"resourceId"`
	ResourceType     types.String `tfsdk:"resource_type" json:"resourceType"`
	Severity         types.String `tfsdk:"severity"`
	StartTime        types.String `tfsdk:"start_time" json:"startTime"`
	Suppressed       types.String `tfsdk:"suppressed"`
	Type             types.String `tfsdk:"type"`
	Unacknowledgedby types.String `tfsdk:"unacknowledgedby"`
}

// NewGetCorrelatedAlarmsCountByParameterDataSource returns a new instance of the generated data source.
func NewGetCorrelatedAlarmsCountByParameterDataSource() datasource.DataSource {
	return &GetCorrelatedAlarmsCountByParameterDataSource{}
}

// Metadata returns the data source type name.
func (d *GetCorrelatedAlarmsCountByParameterDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_correlated_alarms_count_by_parameter"
}

// Schema returns the data source schema.
func (d *GetCorrelatedAlarmsCountByParameterDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get Counts of Correlated Alarms Grouped By The Requested Alarm Parameter", Attributes: map[string]schema.Attribute{"acknowledged": schema.StringAttribute{MarkdownDescription: "filter by acknowledged alarms", Optional: true}, "acknowledgedby": schema.StringAttribute{MarkdownDescription: "filter alarms that are acknowledged by a specific user", Optional: true}, "agg_groups": schema.ListNestedAttribute{MarkdownDescription: "list of aggregate counts per groupBy type", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"count_": schema.Int64Attribute{MarkdownDescription: "the number of alarms in the aggregation group", Computed: true}, "group_name": schema.StringAttribute{MarkdownDescription: "Name of the group. For grouping based on severity, this field will have a value from the set ['Clear', 'Warn', 'Minor', 'Major', 'Critical']", Computed: true}}}}, "alias": schema.StringAttribute{MarkdownDescription: "resource alias name to filter by", Optional: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Cluster ID to filter by", Optional: true}, "device_ip": schema.StringAttribute{MarkdownDescription: "IP address of Device to filter by", Optional: true}, "end_time": schema.StringAttribute{MarkdownDescription: "Optionally specify the end time to filter the records by. In ISO 8601 format 'yyyy-MM-ddTHH:mm:ssZ'", Optional: true}, "group_by": schema.StringAttribute{MarkdownDescription: "aggregation grouping category", Required: true}, "hostname": schema.StringAttribute{MarkdownDescription: "hostname to filter by", Optional: true}, "resource_id": schema.StringAttribute{MarkdownDescription: "Affected entity to filter by", Optional: true}, "resource_type": schema.StringAttribute{MarkdownDescription: "Affected entity type to filter by", Optional: true}, "severity": schema.StringAttribute{MarkdownDescription: "Alarm severity to filter by", Optional: true}, "start_time": schema.StringAttribute{MarkdownDescription: "Optionally specify the start time to filter the records by. In ISO 8601 format 'yyyy-MM-ddTHH:mm:ssZ'", Optional: true}, "suppressed": schema.StringAttribute{MarkdownDescription: "filter by suppressed alarms", Optional: true}, "type": schema.StringAttribute{MarkdownDescription: "Alarm Type to filter by", Optional: true}, "unacknowledgedby": schema.StringAttribute{MarkdownDescription: "filter alarms that are unacknowledged by a specific user", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetCorrelatedAlarmsCountByParameterDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetCorrelatedAlarmsCountByParameterDataSourceModel
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
func (d *GetCorrelatedAlarmsCountByParameterDataSource) readRemote(ctx context.Context, config *GetCorrelatedAlarmsCountByParameterDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/alarms/stats/aggregate/correlated"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_correlated_alarms_count_by_parameter", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.Type.IsNull() {
		query.Set("type", config.Type.ValueString())
	}
	if !config.Severity.IsNull() {
		query.Set("severity", config.Severity.ValueString())
	}
	if !config.StartTime.IsNull() {
		query.Set("startTime", config.StartTime.ValueString())
	}
	if !config.EndTime.IsNull() {
		query.Set("endTime", config.EndTime.ValueString())
	}
	if !config.ResourceType.IsNull() {
		query.Set("resourceType", config.ResourceType.ValueString())
	}
	if !config.ResourceId.IsNull() {
		query.Set("resourceId", config.ResourceId.ValueString())
	}
	if !config.ClusterId.IsNull() {
		query.Set("clusterId", config.ClusterId.ValueString())
	}
	if !config.Hostname.IsNull() {
		query.Set("hostname", config.Hostname.ValueString())
	}
	if !config.Alias.IsNull() {
		query.Set("alias", config.Alias.ValueString())
	}
	if !config.DeviceIp.IsNull() {
		query.Set("deviceIp", config.DeviceIp.ValueString())
	}
	if !config.Acknowledged.IsNull() {
		query.Set("acknowledged", config.Acknowledged.ValueString())
	}
	if !config.Acknowledgedby.IsNull() {
		query.Set("acknowledgedby", config.Acknowledgedby.ValueString())
	}
	if !config.Unacknowledgedby.IsNull() {
		query.Set("unacknowledgedby", config.Unacknowledgedby.ValueString())
	}
	query.Set("groupBy", config.GroupBy.ValueString())
	if !config.Suppressed.IsNull() {
		query.Set("suppressed", config.Suppressed.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_correlated_alarms_count_by_parameter", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_correlated_alarms_count_by_parameter", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_correlated_alarms_count_by_parameter", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_correlated_alarms_count_by_parameter", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_correlated_alarms_count_by_parameter", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_correlated_alarms_count_by_parameter", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_correlated_alarms_count_by_parameter", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_correlated_alarms_count_by_parameter", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_correlated_alarms_count_by_parameter", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_correlated_alarms_count_by_parameter", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_correlated_alarms_count_by_parameter", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_correlated_alarms_count_by_parameter", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["summary"]; ok {
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_correlated_alarms_count_by_parameter", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetCorrelatedAlarmsCountByParameterDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
