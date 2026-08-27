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
	_ datasource.DataSource              = (*LoadDeviceSyslogDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadDeviceSyslogDataSource)(nil)
)

// LoadDeviceSyslogDataSource is the generated Terraform data source implementation.
type LoadDeviceSyslogDataSource struct {
	client *client.Client
}

// LoadDeviceSyslogDataSourceModel describes the data source state shape.
type LoadDeviceSyslogDataSourceModel struct {
	Category           types.String `tfsdk:"category"`
	ClusterId          types.String `tfsdk:"cluster_id" json:"clusterId"`
	CustomerId         types.String `tfsdk:"customer_id" json:"customerId"`
	Description        types.String `tfsdk:"description"`
	DeviceIp           types.String `tfsdk:"device_ip" json:"deviceIp"`
	EndTime            types.String `tfsdk:"end_time" json:"endTime"`
	Hostname           types.String `tfsdk:"hostname"`
	Page               types.String `tfsdk:"page"`
	Pid                types.String `tfsdk:"pid"`
	Process            types.String `tfsdk:"process"`
	ResourceId         types.String `tfsdk:"resource_id" json:"resourceId"`
	ResourceType       types.String `tfsdk:"resource_type" json:"resourceType"`
	Severity           types.String `tfsdk:"severity"`
	Sort               types.String `tfsdk:"sort"`
	Source             types.String `tfsdk:"source"`
	StartTime          types.String `tfsdk:"start_time" json:"startTime"`
	StructuredData     types.String `tfsdk:"structured_data" json:"structuredData"`
	SyslogPriority     types.String `tfsdk:"syslog_priority" json:"syslogPriority"`
	SyslogVersion      types.String `tfsdk:"syslog_version" json:"syslogVersion"`
	TimestampUtcString types.String `tfsdk:"timestamp_utc_string" json:"timestampUtcString"`
	Type               types.String `tfsdk:"type"`
}

// NewLoadDeviceSyslogDataSource returns a new instance of the generated data source.
func NewLoadDeviceSyslogDataSource() datasource.DataSource {
	return &LoadDeviceSyslogDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadDeviceSyslogDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_device_syslog"
}

// Schema returns the data source schema.
func (d *LoadDeviceSyslogDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load Device Syslog", Attributes: map[string]schema.Attribute{"category": schema.StringAttribute{MarkdownDescription: "Category of the log", Computed: true, Optional: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Cluster ID of the device", Optional: true}, "customer_id": schema.StringAttribute{MarkdownDescription: "Customer Id of the log. Gigamon reserved customerId is 26866", Computed: true}, "description": schema.StringAttribute{MarkdownDescription: "Log description", Computed: true}, "device_ip": schema.StringAttribute{MarkdownDescription: "Device IP address", Computed: true, Optional: true}, "end_time": schema.StringAttribute{MarkdownDescription: "End time to filter by. In UTC, ISO 8601 format", Optional: true}, "hostname": schema.StringAttribute{MarkdownDescription: "Device host name", Computed: true, Optional: true}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned. Defaults to (1:100) to prevent reading entire DB by accident", Optional: true}, "pid": schema.StringAttribute{MarkdownDescription: "Process Id of the process generating the log", Computed: true}, "process": schema.StringAttribute{MarkdownDescription: "System process generating the log", Computed: true, Optional: true}, "resource_id": schema.StringAttribute{MarkdownDescription: "Affected Entity of the log", Computed: true, Optional: true}, "resource_type": schema.StringAttribute{MarkdownDescription: "Affected Entity Type of the log", Computed: true, Optional: true}, "severity": schema.StringAttribute{MarkdownDescription: "Severity of the event", Computed: true, Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}, "source": schema.StringAttribute{MarkdownDescription: "Log Source. Control Card or GigaSmart", Computed: true}, "start_time": schema.StringAttribute{MarkdownDescription: "Start time to filter by. In UTC, ISO 8601 format", Optional: true}, "structured_data": schema.StringAttribute{MarkdownDescription: "Additional information about the log. Example: link speed=10000, portID=1/1/x1 for linkChangeNotify event", Computed: true}, "syslog_priority": schema.StringAttribute{MarkdownDescription: "Syslog Priority as per RFC 5424", Computed: true, Optional: true}, "syslog_version": schema.StringAttribute{MarkdownDescription: "Syslog Version as per RFC 5424", Computed: true, Optional: true}, "timestamp_utc_string": schema.StringAttribute{MarkdownDescription: "Log timestamp in UTC represented in ISO 8601 format", Computed: true}, "type": schema.StringAttribute{MarkdownDescription: "Log Type identifier. Example linkChangeNotify", Computed: true, Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadDeviceSyslogDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadDeviceSyslogDataSourceModel
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
func (d *LoadDeviceSyslogDataSource) readRemote(ctx context.Context, config *LoadDeviceSyslogDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/system/syslog"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_device_syslog", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.ClusterId.IsNull() {
		query.Set("clusterId", config.ClusterId.ValueString())
	}
	if !config.DeviceIp.IsNull() {
		query.Set("deviceIp", config.DeviceIp.ValueString())
	}
	if !config.Hostname.IsNull() {
		query.Set("hostname", config.Hostname.ValueString())
	}
	if !config.Category.IsNull() {
		query.Set("category", config.Category.ValueString())
	}
	if !config.Type.IsNull() {
		query.Set("type", config.Type.ValueString())
	}
	if !config.Process.IsNull() {
		query.Set("process", config.Process.ValueString())
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
	if !config.SyslogPriority.IsNull() {
		query.Set("syslogPriority", config.SyslogPriority.ValueString())
	}
	if !config.SyslogVersion.IsNull() {
		query.Set("syslogVersion", config.SyslogVersion.ValueString())
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
		resp.Diagnostics.AddError("Error reading gigavuecore_load_device_syslog", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_device_syslog", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_device_syslog", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_device_syslog", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_device_syslog", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_device_syslog", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_device_syslog", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_device_syslog", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_device_syslog", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_load_device_syslog", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_load_device_syslog", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_device_syslog", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_device_syslog", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadDeviceSyslogDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
