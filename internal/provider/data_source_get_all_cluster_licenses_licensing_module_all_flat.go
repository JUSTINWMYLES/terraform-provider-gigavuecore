package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	_ datasource.DataSource              = (*GetAllClusterLicensesLicensingModuleAllFlatDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAllClusterLicensesLicensingModuleAllFlatDataSource)(nil)
)

// GetAllClusterLicensesLicensingModuleAllFlatDataSource is the generated Terraform data source implementation.
type GetAllClusterLicensesLicensingModuleAllFlatDataSource struct {
	client *client.Client
}

// GetAllClusterLicensesLicensingModuleAllFlatDataSourceModel describes the data source state shape.
type GetAllClusterLicensesLicensingModuleAllFlatDataSourceModel struct {
	BoxId               types.String `tfsdk:"box_id" json:"boxId"`
	ChassisSerialNumber types.String `tfsdk:"chassis_serial_number" json:"chassisSerialNumber"`
	ClusterName         types.String `tfsdk:"cluster_name" json:"clusterName"`
	DeviceSwVersion     types.String `tfsdk:"device_sw_version" json:"deviceSwVersion"`
	EndDate             types.Int64  `tfsdk:"end_date" json:"endDate"`
	Feature             types.List   `tfsdk:"feature"`
	Floated             types.String `tfsdk:"floated"`
	GracePeriod         types.Int64  `tfsdk:"grace_period" json:"gracePeriod"`
	HostName            types.String `tfsdk:"host_name" json:"hostName"`
	HwType              types.String `tfsdk:"hw_type" json:"hwType"`
	LicenseStatus       types.String `tfsdk:"license_status" json:"licenseStatus"`
	LicenseType         types.String `tfsdk:"license_type" json:"licenseType"`
	Model               types.String `tfsdk:"model"`
	NodeId              types.String `tfsdk:"node_id" json:"nodeId"`
	NotifPeriod         types.Int64  `tfsdk:"notif_period" json:"notifPeriod"`
	Page                types.String `tfsdk:"page"`
	SerialNumber        types.String `tfsdk:"serial_number" json:"serialNumber"`
	SlotId              types.String `tfsdk:"slot_id" json:"slotId"`
	Sort                types.String `tfsdk:"sort"`
	StartDate           types.Int64  `tfsdk:"start_date" json:"startDate"`
}

// NewGetAllClusterLicensesLicensingModuleAllFlatDataSource returns a new instance of the generated data source.
func NewGetAllClusterLicensesLicensingModuleAllFlatDataSource() datasource.DataSource {
	return &GetAllClusterLicensesLicensingModuleAllFlatDataSource{}
}

// Metadata returns the data source type name.
func (d *GetAllClusterLicensesLicensingModuleAllFlatDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_cluster_licenses_licensing_module_all_flat"
}

// Schema returns the data source schema.
func (d *GetAllClusterLicensesLicensingModuleAllFlatDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get all clusters with all the chassis and cards information and licenses installed in List View", Attributes: map[string]schema.Attribute{"box_id": schema.StringAttribute{Computed: true}, "chassis_serial_number": schema.StringAttribute{Computed: true}, "cluster_name": schema.StringAttribute{MarkdownDescription: "Get all clusters by Cluster Name", Computed: true, Optional: true}, "device_sw_version": schema.StringAttribute{Computed: true}, "end_date": schema.Int64Attribute{MarkdownDescription: "Get all clusters by EndDate", Computed: true, Optional: true}, "feature": schema.ListAttribute{MarkdownDescription: "Get all clusters by Feature", Computed: true, Optional: true, ElementType: types.StringType}, "floated": schema.StringAttribute{Computed: true}, "grace_period": schema.Int64Attribute{Computed: true}, "host_name": schema.StringAttribute{MarkdownDescription: "Get all clusters by Host Name", Computed: true, Optional: true}, "hw_type": schema.StringAttribute{Computed: true}, "license_status": schema.StringAttribute{Computed: true}, "license_type": schema.StringAttribute{Computed: true}, "model": schema.StringAttribute{Computed: true}, "node_id": schema.StringAttribute{Computed: true}, "notif_period": schema.Int64Attribute{Computed: true}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "serial_number": schema.StringAttribute{MarkdownDescription: "Get all clusters by SerialNumber", Computed: true, Optional: true}, "slot_id": schema.StringAttribute{MarkdownDescription: "Get all clusters by SlotId", Computed: true, Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}, "start_date": schema.Int64Attribute{MarkdownDescription: "Get all clusters by StartDate", Computed: true, Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetAllClusterLicensesLicensingModuleAllFlatDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAllClusterLicensesLicensingModuleAllFlatDataSourceModel
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
func (d *GetAllClusterLicensesLicensingModuleAllFlatDataSource) readRemote(ctx context.Context, config *GetAllClusterLicensesLicensingModuleAllFlatDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/licensing/module/all/flat"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_licenses_licensing_module_all_flat", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.ClusterName.IsNull() {
		query.Set("clusterName", config.ClusterName.ValueString())
	}
	if !config.HostName.IsNull() {
		query.Set("hostName", config.HostName.ValueString())
	}
	if !config.SlotId.IsNull() {
		query.Set("slotId", config.SlotId.ValueString())
	}
	if !config.Feature.IsNull() {
		for _, elem := range config.Feature.Elements() {
			query.Add("feature", elem.(types.String).ValueString())
		}
	}
	if !config.SerialNumber.IsNull() {
		query.Set("serialNumber", config.SerialNumber.ValueString())
	}
	if !config.StartDate.IsNull() {
		query.Set("startDate", strconv.FormatInt(config.StartDate.ValueInt64(), 10))
	}
	if !config.EndDate.IsNull() {
		query.Set("endDate", strconv.FormatInt(config.EndDate.ValueInt64(), 10))
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_licenses_licensing_module_all_flat", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_licenses_licensing_module_all_flat", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_licenses_licensing_module_all_flat", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_licenses_licensing_module_all_flat", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_licenses_licensing_module_all_flat", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_licenses_licensing_module_all_flat", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_licenses_licensing_module_all_flat", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_licenses_licensing_module_all_flat", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_licenses_licensing_module_all_flat", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_licenses_licensing_module_all_flat", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_licenses_licensing_module_all_flat", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_licenses_licensing_module_all_flat", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_licenses_licensing_module_all_flat", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetAllClusterLicensesLicensingModuleAllFlatDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
