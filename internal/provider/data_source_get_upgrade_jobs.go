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
	_ datasource.DataSource              = (*GetUpgradeJobsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetUpgradeJobsDataSource)(nil)
)

// GetUpgradeJobsDataSource is the generated Terraform data source implementation.
type GetUpgradeJobsDataSource struct {
	client *client.Client
}

// GetUpgradeJobsDataSourceModel describes the data source state shape.
type GetUpgradeJobsDataSourceModel struct {
	Context                     types.Object `tfsdk:"context"`
	CreatedBy                   types.String `tfsdk:"created_by" json:"createdBy"`
	DeviceUpgradeTaskInfoRecord types.List   `tfsdk:"device_upgrade_task_info_record" json:"deviceUpgradeTaskInfoRecord"`
	EndDate                     types.String `tfsdk:"end_date" json:"endDate"`
	Page                        types.String `tfsdk:"page"`
	ScheduledTime               types.String `tfsdk:"scheduled_time" json:"scheduledTime"`
	Sort                        types.String `tfsdk:"sort"`
	StartDate                   types.String `tfsdk:"start_date" json:"startDate"`
	TaskId                      types.String `tfsdk:"task_id" json:"taskId"`
	TaskName                    types.String `tfsdk:"task_name" json:"taskName"`
	TaskStatus                  types.String `tfsdk:"task_status" json:"taskStatus"`
	TimeLeft                    types.String `tfsdk:"time_left" json:"timeLeft"`
}

// NewGetUpgradeJobsDataSource returns a new instance of the generated data source.
func NewGetUpgradeJobsDataSource() datasource.DataSource {
	return &GetUpgradeJobsDataSource{}
}

// Metadata returns the data source type name.
func (d *GetUpgradeJobsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_upgrade_jobs"
}

// Schema returns the data source schema.
func (d *GetUpgradeJobsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get Upgrade Jobs", Attributes: map[string]schema.Attribute{"context": schema.SingleNestedAttribute{MarkdownDescription: "Gigamon query result context", Computed: true, Attributes: map[string]schema.Attribute{"page_no": schema.Int64Attribute{MarkdownDescription: "page number of the returned result set", Computed: true}, "page_size": schema.Int64Attribute{MarkdownDescription: "page size of the returned result set", Computed: true}, "sort": schema.ListAttribute{MarkdownDescription: "sorting info of the returned result set. list of fields in the array indicate sorting order", Computed: true, ElementType: types.StringType}, "total_items": schema.Int64Attribute{MarkdownDescription: "total number of items in the queried entity type", Computed: true}}}, "created_by": schema.StringAttribute{Optional: true}, "device_upgrade_task_info_record": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"created_by": schema.StringAttribute{MarkdownDescription: "Created By", Computed: true}, "created_time": schema.StringAttribute{MarkdownDescription: "Created Time", Computed: true}, "description": schema.StringAttribute{MarkdownDescription: "Description", Computed: true}, "end_date": schema.StringAttribute{MarkdownDescription: "End Date", Computed: true}, "scheduled_time": schema.StringAttribute{MarkdownDescription: "Scheduled Time", Computed: true}, "start_date": schema.StringAttribute{MarkdownDescription: "Start Date", Computed: true}, "task_id": schema.StringAttribute{MarkdownDescription: "Task ID", Computed: true}, "task_name": schema.StringAttribute{MarkdownDescription: "Task Name", Computed: true}, "task_status": schema.StringAttribute{MarkdownDescription: "Task Status", Computed: true}, "task_type": schema.StringAttribute{MarkdownDescription: "Task Type", Computed: true}, "time_left": schema.StringAttribute{MarkdownDescription: "Time Left", Computed: true}}}}, "end_date": schema.StringAttribute{Optional: true}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "scheduled_time": schema.StringAttribute{Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}, "start_date": schema.StringAttribute{Optional: true}, "task_id": schema.StringAttribute{Optional: true}, "task_name": schema.StringAttribute{Optional: true}, "task_status": schema.StringAttribute{Optional: true}, "time_left": schema.StringAttribute{Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetUpgradeJobsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetUpgradeJobsDataSourceModel
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
func (d *GetUpgradeJobsDataSource) readRemote(ctx context.Context, config *GetUpgradeJobsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/device/upgrade/orchestration/upgradeJobs"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_jobs", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.TaskId.IsNull() {
		query.Set("taskId", config.TaskId.ValueString())
	}
	if !config.TaskName.IsNull() {
		query.Set("taskName", config.TaskName.ValueString())
	}
	if !config.TaskStatus.IsNull() {
		query.Set("taskStatus", config.TaskStatus.ValueString())
	}
	if !config.StartDate.IsNull() {
		query.Set("startDate", config.StartDate.ValueString())
	}
	if !config.EndDate.IsNull() {
		query.Set("endDate", config.EndDate.ValueString())
	}
	if !config.ScheduledTime.IsNull() {
		query.Set("scheduledTime", config.ScheduledTime.ValueString())
	}
	if !config.CreatedBy.IsNull() {
		query.Set("createdBy", config.CreatedBy.ValueString())
	}
	if !config.TimeLeft.IsNull() {
		query.Set("timeLeft", config.TimeLeft.ValueString())
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_jobs", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_jobs", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_jobs", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_jobs", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_jobs", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_jobs", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_jobs", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_jobs", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_jobs", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_jobs", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_jobs", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_jobs", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_jobs", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetUpgradeJobsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
