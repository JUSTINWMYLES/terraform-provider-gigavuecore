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
	_ datasource.DataSource              = (*LoadAllAlarmsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadAllAlarmsDataSource)(nil)
)

// LoadAllAlarmsDataSource is the generated Terraform data source implementation.
type LoadAllAlarmsDataSource struct {
	client *client.Client
}

// LoadAllAlarmsDataSourceModel describes the data source state shape.
type LoadAllAlarmsDataSourceModel struct {
	Acknowledged     types.String `tfsdk:"acknowledged"`
	Acknowledgedby   types.String `tfsdk:"acknowledgedby"`
	Alarms           types.List   `tfsdk:"alarms"`
	Alias            types.String `tfsdk:"alias"`
	ClusterId        types.String `tfsdk:"cluster_id" json:"clusterId"`
	Context          types.Object `tfsdk:"context"`
	DeviceIp         types.String `tfsdk:"device_ip" json:"deviceIp"`
	EndTime          types.String `tfsdk:"end_time" json:"endTime"`
	Hostname         types.String `tfsdk:"hostname"`
	Page             types.String `tfsdk:"page"`
	ResourceId       types.String `tfsdk:"resource_id" json:"resourceId"`
	ResourceType     types.String `tfsdk:"resource_type" json:"resourceType"`
	Severity         types.String `tfsdk:"severity"`
	Sort             types.String `tfsdk:"sort"`
	StartTime        types.String `tfsdk:"start_time" json:"startTime"`
	Suppressed       types.String `tfsdk:"suppressed"`
	Type             types.String `tfsdk:"type"`
	Unacknowledgedby types.String `tfsdk:"unacknowledgedby"`
}

// NewLoadAllAlarmsDataSource returns a new instance of the generated data source.
func NewLoadAllAlarmsDataSource() datasource.DataSource {
	return &LoadAllAlarmsDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadAllAlarmsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_all_alarms"
}

// Schema returns the data source schema.
func (d *LoadAllAlarmsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load All Alarms", Attributes: map[string]schema.Attribute{"acknowledged": schema.StringAttribute{MarkdownDescription: "filter by acknowledged alarms", Optional: true}, "acknowledgedby": schema.StringAttribute{MarkdownDescription: "filter alarms that are acknowledged by a specific user", Optional: true}, "alarms": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"acknowledged": schema.BoolAttribute{MarkdownDescription: "True if the alarm is acknowledged by the user", Computed: true}, "acknowledged_by": schema.StringAttribute{MarkdownDescription: "User who acknowledged the alarm", Computed: true}, "acknowledged_ts": schema.StringAttribute{MarkdownDescription: "Alarm acknowledged timestamp in ISO 8601 format", Computed: true}, "alias": schema.StringAttribute{MarkdownDescription: "Alias of the Resource ID", Computed: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Cluster ID of the device", Computed: true}, "comment": schema.StringAttribute{MarkdownDescription: "User comments", Computed: true}, "description": schema.StringAttribute{MarkdownDescription: "Alarm description", Computed: true}, "device_ip": schema.StringAttribute{MarkdownDescription: "IP address of the device", Computed: true}, "hostname": schema.StringAttribute{MarkdownDescription: "Hostname of the Resource", Computed: true}, "resource_id": schema.StringAttribute{MarkdownDescription: "Resource ID of the alarm", Computed: true}, "resource_type": schema.StringAttribute{MarkdownDescription: "Resource Type of the alarm", Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Severity of the alarm", Computed: true}, "suppressed": schema.BoolAttribute{MarkdownDescription: "True if the alarm is suppressed", Computed: true}, "ts": schema.StringAttribute{MarkdownDescription: "Alarm timestamp in ISO 8601 format", Computed: true}, "type": schema.StringAttribute{MarkdownDescription: "Alarm Type identifier", Computed: true}, "unacknowledged_by": schema.StringAttribute{MarkdownDescription: "User who unacknowledged the alarm", Computed: true}, "unacknowledged_ts": schema.StringAttribute{MarkdownDescription: "Alarm unacknowledged timestamp in ISO 8601 format", Computed: true}}}}, "alias": schema.StringAttribute{MarkdownDescription: " resource alias name to filter by", Optional: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Cluster ID to filter by", Optional: true}, "context": schema.SingleNestedAttribute{MarkdownDescription: "Gigamon query result context", Computed: true, Attributes: map[string]schema.Attribute{"page_no": schema.Int64Attribute{MarkdownDescription: "page number of the returned result set", Computed: true}, "page_size": schema.Int64Attribute{MarkdownDescription: "page size of the returned result set", Computed: true}, "sort": schema.ListAttribute{MarkdownDescription: "sorting info of the returned result set. list of fields in the array indicate sorting order", Computed: true, ElementType: types.StringType}, "total_items": schema.Int64Attribute{MarkdownDescription: "total number of items in the queried entity type", Computed: true}}}, "device_ip": schema.StringAttribute{MarkdownDescription: "IP address of Device to filter by", Optional: true}, "end_time": schema.StringAttribute{MarkdownDescription: "Filter End timestamps to include reports ending by. In ISO-8601 date format 'yyyy-MM-ddTHH:mm:ssZ'", Optional: true}, "hostname": schema.StringAttribute{MarkdownDescription: "hostname to filter by", Optional: true}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned. Defaults to (1:30) to prevent reading entire DB by accident", Optional: true}, "resource_id": schema.StringAttribute{MarkdownDescription: "Affected entity to filter by", Optional: true}, "resource_type": schema.StringAttribute{MarkdownDescription: "Affected entity type to filter by", Optional: true}, "severity": schema.StringAttribute{MarkdownDescription: "Alarm severity to filter by", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is DESC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}, "start_time": schema.StringAttribute{MarkdownDescription: "Filter Start timestamps to include reports starting from. In ISO-8601 date format 'yyyy-MM-ddTHH:mm:ssZ'", Optional: true}, "suppressed": schema.StringAttribute{MarkdownDescription: "filter by suppressed alarms", Optional: true}, "type": schema.StringAttribute{MarkdownDescription: "Alarm Type to filter by", Optional: true}, "unacknowledgedby": schema.StringAttribute{MarkdownDescription: "filter alarms that are unacknowledged by a specific user", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadAllAlarmsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadAllAlarmsDataSourceModel
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
func (d *LoadAllAlarmsDataSource) readRemote(ctx context.Context, config *LoadAllAlarmsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/alarms"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_alarms", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.Page.IsNull() {
		query.Set("page", config.Page.ValueString())
	}
	if !config.Sort.IsNull() {
		query.Set("sort", config.Sort.ValueString())
	}
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
	if !config.Suppressed.IsNull() {
		query.Set("suppressed", config.Suppressed.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_alarms", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_alarms", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_alarms", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_alarms", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_alarms", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_alarms", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_alarms", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_alarms", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_alarms", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_load_all_alarms", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_alarms", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_alarms", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_alarms", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadAllAlarmsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
