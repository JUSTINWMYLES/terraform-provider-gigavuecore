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
	_ datasource.DataSource              = (*GetUpgradeJobInfoDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetUpgradeJobInfoDataSource)(nil)
)

// GetUpgradeJobInfoDataSource is the generated Terraform data source implementation.
type GetUpgradeJobInfoDataSource struct {
	client *client.Client
}

// GetUpgradeJobInfoDataSourceModel describes the data source state shape.
type GetUpgradeJobInfoDataSourceModel struct {
	CreatedBy            types.String `tfsdk:"created_by" json:"createdBy"`
	CreatedTime          types.String `tfsdk:"created_time" json:"createdTime"`
	Description          types.String `tfsdk:"description"`
	DeviceUpgradeContext types.List   `tfsdk:"device_upgrade_context" json:"deviceUpgradeContext"`
	EndDate              types.String `tfsdk:"end_date" json:"endDate"`
	ImageServer          types.String `tfsdk:"image_server" json:"imageServer"`
	LastUpdateTime       types.String `tfsdk:"last_update_time" json:"lastUpdateTime"`
	ScheduledTime        types.String `tfsdk:"scheduled_time" json:"scheduledTime"`
	StartDate            types.String `tfsdk:"start_date" json:"startDate"`
	TaskId               types.String `tfsdk:"task_id" json:"taskId"`
	TaskName             types.String `tfsdk:"task_name" json:"taskName"`
	TaskStatus           types.String `tfsdk:"task_status" json:"taskStatus"`
	TaskType             types.String `tfsdk:"task_type" json:"taskType"`
	TimeLeft             types.String `tfsdk:"time_left" json:"timeLeft"`
}

// NewGetUpgradeJobInfoDataSource returns a new instance of the generated data source.
func NewGetUpgradeJobInfoDataSource() datasource.DataSource {
	return &GetUpgradeJobInfoDataSource{}
}

// Metadata returns the data source type name.
func (d *GetUpgradeJobInfoDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_upgrade_job_info"
}

// Schema returns the data source schema.
func (d *GetUpgradeJobInfoDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get Upgrade Job Info", Attributes: map[string]schema.Attribute{"created_by": schema.StringAttribute{MarkdownDescription: "Created By", Computed: true}, "created_time": schema.StringAttribute{MarkdownDescription: "Created Time", Computed: true}, "description": schema.StringAttribute{MarkdownDescription: "Description", Computed: true}, "device_upgrade_context": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"cluster_name": schema.StringAttribute{MarkdownDescription: "Cluster Name", Computed: true}, "config_backup": schema.BoolAttribute{MarkdownDescription: "Config Backup", Computed: true}, "device_ip": schema.StringAttribute{MarkdownDescription: "Device IP", Computed: true}, "file_path": schema.StringAttribute{MarkdownDescription: "File Path", Computed: true}, "file_type": schema.StringAttribute{MarkdownDescription: "File Type", Computed: true}, "hostname": schema.StringAttribute{MarkdownDescription: "Host Name", Computed: true}, "model": schema.StringAttribute{MarkdownDescription: "Model", Computed: true}}}}, "end_date": schema.StringAttribute{MarkdownDescription: "End Date", Computed: true}, "image_server": schema.StringAttribute{MarkdownDescription: "Image Server", Computed: true}, "last_update_time": schema.StringAttribute{MarkdownDescription: "Last Update Time", Computed: true}, "scheduled_time": schema.StringAttribute{MarkdownDescription: "Scheduled Time", Computed: true}, "start_date": schema.StringAttribute{MarkdownDescription: "Start Date", Computed: true}, "task_id": schema.StringAttribute{MarkdownDescription: "Task ID", Required: true}, "task_name": schema.StringAttribute{MarkdownDescription: "Task Name", Computed: true}, "task_status": schema.StringAttribute{MarkdownDescription: "Task Status", Computed: true}, "task_type": schema.StringAttribute{MarkdownDescription: "Task Type", Computed: true}, "time_left": schema.StringAttribute{MarkdownDescription: "Time Left", Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetUpgradeJobInfoDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetUpgradeJobInfoDataSourceModel
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
func (d *GetUpgradeJobInfoDataSource) readRemote(ctx context.Context, config *GetUpgradeJobInfoDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/device/upgrade/orchestration/upgradeJobInfo"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_job_info", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("taskId", config.TaskId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_job_info", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_job_info", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_job_info", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_job_info", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_job_info", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_job_info", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_job_info", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_job_info", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_job_info", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_job_info", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_job_info", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_job_info", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_job_info", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetUpgradeJobInfoDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
