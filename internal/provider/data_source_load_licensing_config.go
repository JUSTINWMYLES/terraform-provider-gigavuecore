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
	_ datasource.DataSource              = (*LoadLicensingConfigDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadLicensingConfigDataSource)(nil)
)

// LoadLicensingConfigDataSource is the generated Terraform data source implementation.
type LoadLicensingConfigDataSource struct {
	client *client.Client
}

// LoadLicensingConfigDataSourceModel describes the data source state shape.
type LoadLicensingConfigDataSourceModel struct {
	ExpiredLicensesExist              types.Bool   `tfsdk:"expired_licenses_exist" json:"expiredLicensesExist"`
	ExpiryAlertsLastSentOn            types.Int64  `tfsdk:"expiry_alerts_last_sent_on" json:"expiryAlertsLastSentOn"`
	LicenseExpiryAlertEnabled         types.Bool   `tfsdk:"license_expiry_alert_enabled" json:"licenseExpiryAlertEnabled"`
	MandatoryUpcomingRenewalsReceiver types.String `tfsdk:"mandatory_upcoming_renewals_receiver" json:"mandatoryUpcomingRenewalsReceiver"`
	MandatoryVblReportReceiver        types.String `tfsdk:"mandatory_vbl_report_receiver" json:"mandatoryVblReportReceiver"`
	MonthsPerPeriod                   types.Int64  `tfsdk:"months_per_period" json:"monthsPerPeriod"`
	OtherUpcomingRenewalsReceivers    types.List   `tfsdk:"other_upcoming_renewals_receivers" json:"otherUpcomingRenewalsReceivers"`
	OtherVblReportReceivers           types.List   `tfsdk:"other_vbl_report_receivers" json:"otherVblReportReceivers"`
	VolumeUsageAlertEnabled           types.Bool   `tfsdk:"volume_usage_alert_enabled" json:"volumeUsageAlertEnabled"`
	VolumeUsageAlertThreshold         types.Int64  `tfsdk:"volume_usage_alert_threshold" json:"volumeUsageAlertThreshold"`
}

// NewLoadLicensingConfigDataSource returns a new instance of the generated data source.
func NewLoadLicensingConfigDataSource() datasource.DataSource {
	return &LoadLicensingConfigDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadLicensingConfigDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_licensing_config"
}

// Schema returns the data source schema.
func (d *LoadLicensingConfigDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load Licensing Config", Attributes: map[string]schema.Attribute{"expired_licenses_exist": schema.BoolAttribute{Computed: true}, "expiry_alerts_last_sent_on": schema.Int64Attribute{Computed: true}, "license_expiry_alert_enabled": schema.BoolAttribute{Computed: true}, "mandatory_upcoming_renewals_receiver": schema.StringAttribute{Computed: true}, "mandatory_vbl_report_receiver": schema.StringAttribute{Computed: true}, "months_per_period": schema.Int64Attribute{Computed: true}, "other_upcoming_renewals_receivers": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "other_vbl_report_receivers": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "volume_usage_alert_enabled": schema.BoolAttribute{Computed: true}, "volume_usage_alert_threshold": schema.Int64Attribute{Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadLicensingConfigDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadLicensingConfigDataSourceModel
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
func (d *LoadLicensingConfigDataSource) readRemote(ctx context.Context, config *LoadLicensingConfigDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/licensing/config"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_licensing_config", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_licensing_config", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_licensing_config", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_licensing_config", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_licensing_config", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_licensing_config", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_licensing_config", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_licensing_config", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_licensing_config", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_licensing_config", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_load_licensing_config", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_load_licensing_config", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_licensing_config", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_licensing_config", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadLicensingConfigDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
