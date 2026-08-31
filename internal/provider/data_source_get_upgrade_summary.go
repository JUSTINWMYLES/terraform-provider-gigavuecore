package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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
	ActivateInfo      types.String  `tfsdk:"activate_info" json:"activateInfo"`
	ActivateStage     types.String  `tfsdk:"activate_stage" json:"activateStage"`
	ChassisOperStatus types.String  `tfsdk:"chassis_oper_status" json:"chassisOperStatus"`
	ClusterId         types.String  `tfsdk:"cluster_id" json:"clusterId"`
	FetchInfo         types.String  `tfsdk:"fetch_info" json:"fetchInfo"`
	FetchStage        types.String  `tfsdk:"fetch_stage" json:"fetchStage"`
	HealthState       types.String  `tfsdk:"health_state" json:"healthState"`
	Hostname          types.String  `tfsdk:"hostname"`
	InstallInfo       types.String  `tfsdk:"install_info" json:"installInfo"`
	InstallStage      types.String  `tfsdk:"install_stage" json:"installStage"`
	Items             types.Dynamic `tfsdk:"items"`
	LastExecutedTask  types.String  `tfsdk:"last_executed_task" json:"lastExecutedTask"`
	Licensed          types.String  `tfsdk:"licensed"`
	Model             types.String  `tfsdk:"model"`
	Page              types.String  `tfsdk:"page"`
	Role              types.String  `tfsdk:"role"`
	SerialNumber      types.String  `tfsdk:"serial_number" json:"serialNumber"`
	Sort              types.String  `tfsdk:"sort"`
	SwVersion         types.String  `tfsdk:"sw_version" json:"swVersion"`
	Tags              types.String  `tfsdk:"tags"`
	TaskStatus        types.String  `tfsdk:"task_status" json:"taskStatus"`
	UbootVersion      types.String  `tfsdk:"uboot_version" json:"ubootVersion"`
	VerifyInfo        types.String  `tfsdk:"verify_info" json:"verifyInfo"`
	VerifyStage       types.String  `tfsdk:"verify_stage" json:"verifyStage"`
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
	resp.Schema = schema.Schema{MarkdownDescription: "Get Upgrade Summary", Attributes: map[string]schema.Attribute{"activate_info": schema.StringAttribute{MarkdownDescription: "Activate Info", Optional: true}, "activate_stage": schema.StringAttribute{MarkdownDescription: "Activate Stage", Optional: true}, "chassis_oper_status": schema.StringAttribute{MarkdownDescription: "Chassis OperStatus", Optional: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Cluster ID", Optional: true}, "fetch_info": schema.StringAttribute{MarkdownDescription: "Fetch Info", Optional: true}, "fetch_stage": schema.StringAttribute{MarkdownDescription: "Fetch Stage", Optional: true}, "health_state": schema.StringAttribute{MarkdownDescription: "Health State", Optional: true}, "hostname": schema.StringAttribute{MarkdownDescription: "Host Name", Optional: true}, "install_info": schema.StringAttribute{MarkdownDescription: "Install Info", Optional: true}, "install_stage": schema.StringAttribute{MarkdownDescription: "Install Stage", Optional: true}, "items": schema.DynamicAttribute{Computed: true}, "last_executed_task": schema.StringAttribute{MarkdownDescription: "Last Executed Task", Optional: true}, "licensed": schema.StringAttribute{MarkdownDescription: "Licensed", Optional: true}, "model": schema.StringAttribute{MarkdownDescription: "Model", Optional: true}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "role": schema.StringAttribute{MarkdownDescription: "Role", Optional: true}, "serial_number": schema.StringAttribute{MarkdownDescription: "Serial Number", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}, "sw_version": schema.StringAttribute{MarkdownDescription: "Software Version", Optional: true}, "tags": schema.StringAttribute{MarkdownDescription: "Tags", Optional: true}, "task_status": schema.StringAttribute{MarkdownDescription: "Task Status", Optional: true}, "uboot_version": schema.StringAttribute{MarkdownDescription: "UBoot Version", Optional: true}, "verify_info": schema.StringAttribute{MarkdownDescription: "verifyInfo", Optional: true}, "verify_stage": schema.StringAttribute{MarkdownDescription: "Verify Stage", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetUpgradeSummaryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetUpgradeSummaryDataSourceModel
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
func (d *GetUpgradeSummaryDataSource) readListRemote(ctx context.Context, config *GetUpgradeSummaryDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/device/upgrade/orchestration/upgradeSummary"
	params := url.Values{}
	if !config.ClusterId.IsNull() {
		params.Set("clusterId", config.ClusterId.ValueString())
	}
	if !config.Hostname.IsNull() {
		params.Set("hostname", config.Hostname.ValueString())
	}
	if !config.Model.IsNull() {
		params.Set("model", config.Model.ValueString())
	}
	if !config.SwVersion.IsNull() {
		params.Set("swVersion", config.SwVersion.ValueString())
	}
	if !config.HealthState.IsNull() {
		params.Set("healthState", config.HealthState.ValueString())
	}
	if !config.Licensed.IsNull() {
		params.Set("licensed", config.Licensed.ValueString())
	}
	if !config.Role.IsNull() {
		params.Set("role", config.Role.ValueString())
	}
	if !config.UbootVersion.IsNull() {
		params.Set("ubootVersion", config.UbootVersion.ValueString())
	}
	if !config.SerialNumber.IsNull() {
		params.Set("serialNumber", config.SerialNumber.ValueString())
	}
	if !config.FetchStage.IsNull() {
		params.Set("fetchStage", config.FetchStage.ValueString())
	}
	if !config.InstallStage.IsNull() {
		params.Set("installStage", config.InstallStage.ValueString())
	}
	if !config.ActivateStage.IsNull() {
		params.Set("activateStage", config.ActivateStage.ValueString())
	}
	if !config.VerifyStage.IsNull() {
		params.Set("verifyStage", config.VerifyStage.ValueString())
	}
	if !config.FetchInfo.IsNull() {
		params.Set("fetchInfo", config.FetchInfo.ValueString())
	}
	if !config.InstallInfo.IsNull() {
		params.Set("installInfo", config.InstallInfo.ValueString())
	}
	if !config.ActivateInfo.IsNull() {
		params.Set("activateInfo", config.ActivateInfo.ValueString())
	}
	if !config.VerifyInfo.IsNull() {
		params.Set("verifyInfo", config.VerifyInfo.ValueString())
	}
	if !config.TaskStatus.IsNull() {
		params.Set("taskStatus", config.TaskStatus.ValueString())
	}
	if !config.ChassisOperStatus.IsNull() {
		params.Set("chassisOperStatus", config.ChassisOperStatus.ValueString())
	}
	if !config.Tags.IsNull() {
		params.Set("tags", config.Tags.ValueString())
	}
	if !config.LastExecutedTask.IsNull() {
		params.Set("lastExecutedTask", config.LastExecutedTask.ValueString())
	}
	if !config.Page.IsNull() {
		params.Set("page", config.Page.ValueString())
	}
	if !config.Sort.IsNull() {
		params.Set("sort", config.Sort.ValueString())
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_summary", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_summary", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["deviceUpgradeSummary"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_summary", fmt.Sprintf("Could not decode list page: missing %q array", "deviceUpgradeSummary"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
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
