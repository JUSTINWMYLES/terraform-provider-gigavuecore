package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*GetAlarmByIdDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAlarmByIdDataSource)(nil)
)

// GetAlarmByIdDataSource is the generated Terraform data source implementation.
type GetAlarmByIdDataSource struct {
	client *client.Client
}

// GetAlarmByIdDataSourceModel describes the data source state shape.
type GetAlarmByIdDataSourceModel struct {
	Acknowledged     types.Bool   `tfsdk:"acknowledged"`
	AcknowledgedBy   types.String `tfsdk:"acknowledged_by" json:"acknowledgedBy"`
	AcknowledgedTs   types.String `tfsdk:"acknowledged_ts" json:"acknowledgedTs"`
	AlarmId          types.String `tfsdk:"alarm_id" json:"alarmId"`
	Alias            types.String `tfsdk:"alias"`
	ClusterId        types.String `tfsdk:"cluster_id" json:"clusterId"`
	Comment          types.String `tfsdk:"comment"`
	Description      types.String `tfsdk:"description"`
	DeviceIp         types.String `tfsdk:"device_ip" json:"deviceIp"`
	Hostname         types.String `tfsdk:"hostname"`
	ResourceId       types.String `tfsdk:"resource_id" json:"resourceId"`
	ResourceType     types.String `tfsdk:"resource_type" json:"resourceType"`
	Severity         types.String `tfsdk:"severity"`
	Suppressed       types.Bool   `tfsdk:"suppressed"`
	Ts               types.String `tfsdk:"ts"`
	Type             types.String `tfsdk:"type"`
	UnacknowledgedBy types.String `tfsdk:"unacknowledged_by" json:"unacknowledgedBy"`
	UnacknowledgedTs types.String `tfsdk:"unacknowledged_ts" json:"unacknowledgedTs"`
}

// NewGetAlarmByIdDataSource returns a new instance of the generated data source.
func NewGetAlarmByIdDataSource() datasource.DataSource {
	return &GetAlarmByIdDataSource{}
}

// Metadata returns the data source type name.
func (d *GetAlarmByIdDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_alarm_by_id"
}

// Schema returns the data source schema.
func (d *GetAlarmByIdDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Find Alarm by ID", Attributes: map[string]schema.Attribute{"acknowledged": schema.BoolAttribute{MarkdownDescription: "True if the alarm is acknowledged by the user", Computed: true}, "acknowledged_by": schema.StringAttribute{MarkdownDescription: "User who acknowledged the alarm", Computed: true}, "acknowledged_ts": schema.StringAttribute{MarkdownDescription: "Alarm acknowledged timestamp in ISO 8601 format", Computed: true}, "alarm_id": schema.StringAttribute{MarkdownDescription: "ID of the target Alarm", Required: true}, "alias": schema.StringAttribute{MarkdownDescription: "Alias of the Resource ID", Computed: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Cluster ID of the device", Computed: true}, "comment": schema.StringAttribute{MarkdownDescription: "User comments", Computed: true}, "description": schema.StringAttribute{MarkdownDescription: "Alarm description", Computed: true}, "device_ip": schema.StringAttribute{MarkdownDescription: "IP address of the device", Computed: true}, "hostname": schema.StringAttribute{MarkdownDescription: "Hostname of the Resource", Computed: true}, "resource_id": schema.StringAttribute{MarkdownDescription: "Resource ID of the alarm", Computed: true}, "resource_type": schema.StringAttribute{MarkdownDescription: "Resource Type of the alarm", Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Severity of the alarm", Computed: true}, "suppressed": schema.BoolAttribute{MarkdownDescription: "True if the alarm is suppressed", Computed: true}, "ts": schema.StringAttribute{MarkdownDescription: "Alarm timestamp in ISO 8601 format", Computed: true}, "type": schema.StringAttribute{MarkdownDescription: "Alarm Type identifier", Computed: true}, "unacknowledged_by": schema.StringAttribute{MarkdownDescription: "User who unacknowledged the alarm", Computed: true}, "unacknowledged_ts": schema.StringAttribute{MarkdownDescription: "Alarm unacknowledged timestamp in ISO 8601 format", Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetAlarmByIdDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAlarmByIdDataSourceModel
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
func (d *GetAlarmByIdDataSource) readRemote(ctx context.Context, config *GetAlarmByIdDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/alarms/{alarmId}"
	reqPath = strings.ReplaceAll(reqPath, "{alarmId}", url.PathEscape(config.AlarmId.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_alarm_by_id", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_alarm_by_id", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_alarm_by_id", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_alarm_by_id", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_alarm_by_id", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_alarm_by_id", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_alarm_by_id", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_alarm_by_id", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_alarm_by_id", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_alarm_by_id", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_alarm_by_id", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_alarm_by_id", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_alarm_by_id", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["alarm"]; ok {
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_alarm_by_id", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetAlarmByIdDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
