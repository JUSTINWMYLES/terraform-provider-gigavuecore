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
	_ datasource.DataSource              = (*GetUpgradeSummaryDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetUpgradeSummaryDataSource)(nil)
)

// GetUpgradeSummaryDataSource is the generated Terraform data source implementation.
type GetUpgradeSummaryDataSource struct {
	client *client.Client
}

// GetUpgradeSummaryDataSourceModel describes the data source state shape.
type GetUpgradeSummaryDataSourceModel struct {
	ActivateInfo         types.String  `tfsdk:"activate_info" json:"activateInfo"`
	ActivateStage        types.String  `tfsdk:"activate_stage" json:"activateStage"`
	ChassisOperStatus    types.String  `tfsdk:"chassis_oper_status" json:"chassisOperStatus"`
	ClusterId            types.String  `tfsdk:"cluster_id" json:"clusterId"`
	Context              types.Object  `tfsdk:"context"`
	DeviceUpgradeSummary types.Dynamic `tfsdk:"device_upgrade_summary" json:"deviceUpgradeSummary"`
	FetchInfo            types.String  `tfsdk:"fetch_info" json:"fetchInfo"`
	FetchStage           types.String  `tfsdk:"fetch_stage" json:"fetchStage"`
	HealthState          types.String  `tfsdk:"health_state" json:"healthState"`
	Hostname             types.String  `tfsdk:"hostname"`
	InstallInfo          types.String  `tfsdk:"install_info" json:"installInfo"`
	InstallStage         types.String  `tfsdk:"install_stage" json:"installStage"`
	LastExecutedTask     types.String  `tfsdk:"last_executed_task" json:"lastExecutedTask"`
	Licensed             types.String  `tfsdk:"licensed"`
	Model                types.String  `tfsdk:"model"`
	Page                 types.String  `tfsdk:"page"`
	Role                 types.String  `tfsdk:"role"`
	SerialNumber         types.String  `tfsdk:"serial_number" json:"serialNumber"`
	Sort                 types.String  `tfsdk:"sort"`
	SwVersion            types.String  `tfsdk:"sw_version" json:"swVersion"`
	Tags                 types.String  `tfsdk:"tags"`
	TaskStatus           types.String  `tfsdk:"task_status" json:"taskStatus"`
	UbootVersion         types.String  `tfsdk:"uboot_version" json:"ubootVersion"`
	VerifyInfo           types.String  `tfsdk:"verify_info" json:"verifyInfo"`
	VerifyStage          types.String  `tfsdk:"verify_stage" json:"verifyStage"`
}

// NewGetUpgradeSummaryDataSource returns a new instance of the generated data source.
func NewGetUpgradeSummaryDataSource() datasource.DataSource {
	return &GetUpgradeSummaryDataSource{}
}

// Metadata returns the data source type name.
func (d *GetUpgradeSummaryDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_upgrade_summary"
}

// Schema returns the data source schema.
func (d *GetUpgradeSummaryDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get Upgrade Summary", Attributes: map[string]schema.Attribute{"activate_info": schema.StringAttribute{Optional: true}, "activate_stage": schema.StringAttribute{Optional: true}, "chassis_oper_status": schema.StringAttribute{Optional: true}, "cluster_id": schema.StringAttribute{Optional: true}, "context": schema.SingleNestedAttribute{MarkdownDescription: "Gigamon query result context", Computed: true, Attributes: map[string]schema.Attribute{"page_no": schema.Int64Attribute{MarkdownDescription: "page number of the returned result set", Computed: true}, "page_size": schema.Int64Attribute{MarkdownDescription: "page size of the returned result set", Computed: true}, "sort": schema.ListAttribute{MarkdownDescription: "sorting info of the returned result set. list of fields in the array indicate sorting order", Computed: true, ElementType: types.StringType}, "total_items": schema.Int64Attribute{MarkdownDescription: "total number of items in the queried entity type", Computed: true}}}, "device_upgrade_summary": schema.DynamicAttribute{Computed: true}, "fetch_info": schema.StringAttribute{Optional: true}, "fetch_stage": schema.StringAttribute{Optional: true}, "health_state": schema.StringAttribute{Optional: true}, "hostname": schema.StringAttribute{Optional: true}, "install_info": schema.StringAttribute{Optional: true}, "install_stage": schema.StringAttribute{Optional: true}, "last_executed_task": schema.StringAttribute{Optional: true}, "licensed": schema.StringAttribute{Optional: true}, "model": schema.StringAttribute{Optional: true}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "role": schema.StringAttribute{Optional: true}, "serial_number": schema.StringAttribute{Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}, "sw_version": schema.StringAttribute{Optional: true}, "tags": schema.StringAttribute{Optional: true}, "task_status": schema.StringAttribute{Optional: true}, "uboot_version": schema.StringAttribute{Optional: true}, "verify_info": schema.StringAttribute{Optional: true}, "verify_stage": schema.StringAttribute{Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetUpgradeSummaryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetUpgradeSummaryDataSourceModel
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
func (d *GetUpgradeSummaryDataSource) readRemote(ctx context.Context, config *GetUpgradeSummaryDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/device/upgrade/orchestration/upgradeSummary"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_summary", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.ClusterId.IsNull() {
		query.Set("clusterId", config.ClusterId.ValueString())
	}
	if !config.Hostname.IsNull() {
		query.Set("hostname", config.Hostname.ValueString())
	}
	if !config.Model.IsNull() {
		query.Set("model", config.Model.ValueString())
	}
	if !config.SwVersion.IsNull() {
		query.Set("swVersion", config.SwVersion.ValueString())
	}
	if !config.HealthState.IsNull() {
		query.Set("healthState", config.HealthState.ValueString())
	}
	if !config.Licensed.IsNull() {
		query.Set("licensed", config.Licensed.ValueString())
	}
	if !config.Role.IsNull() {
		query.Set("role", config.Role.ValueString())
	}
	if !config.UbootVersion.IsNull() {
		query.Set("ubootVersion", config.UbootVersion.ValueString())
	}
	if !config.SerialNumber.IsNull() {
		query.Set("serialNumber", config.SerialNumber.ValueString())
	}
	if !config.FetchStage.IsNull() {
		query.Set("fetchStage", config.FetchStage.ValueString())
	}
	if !config.InstallStage.IsNull() {
		query.Set("installStage", config.InstallStage.ValueString())
	}
	if !config.ActivateStage.IsNull() {
		query.Set("activateStage", config.ActivateStage.ValueString())
	}
	if !config.VerifyStage.IsNull() {
		query.Set("verifyStage", config.VerifyStage.ValueString())
	}
	if !config.FetchInfo.IsNull() {
		query.Set("fetchInfo", config.FetchInfo.ValueString())
	}
	if !config.InstallInfo.IsNull() {
		query.Set("installInfo", config.InstallInfo.ValueString())
	}
	if !config.ActivateInfo.IsNull() {
		query.Set("activateInfo", config.ActivateInfo.ValueString())
	}
	if !config.VerifyInfo.IsNull() {
		query.Set("verifyInfo", config.VerifyInfo.ValueString())
	}
	if !config.TaskStatus.IsNull() {
		query.Set("taskStatus", config.TaskStatus.ValueString())
	}
	if !config.ChassisOperStatus.IsNull() {
		query.Set("chassisOperStatus", config.ChassisOperStatus.ValueString())
	}
	if !config.Tags.IsNull() {
		query.Set("tags", config.Tags.ValueString())
	}
	if !config.LastExecutedTask.IsNull() {
		query.Set("lastExecutedTask", config.LastExecutedTask.ValueString())
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_summary", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_summary", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_summary", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_summary", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_summary", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_summary", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_summary", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_summary", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_summary", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_summary", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_summary", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_summary", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_summary", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetUpgradeSummaryDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
