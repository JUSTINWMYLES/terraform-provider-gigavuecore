package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*GetPeriodVolumeDetailedByDateDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetPeriodVolumeDetailedByDateDataSource)(nil)
)

// GetPeriodVolumeDetailedByDateDataSource is the generated Terraform data source implementation.
type GetPeriodVolumeDetailedByDateDataSource struct {
	client *client.Client
}

// GetPeriodVolumeDetailedByDateDataSourceModel describes the data source state shape.
type GetPeriodVolumeDetailedByDateDataSourceModel struct {
	AppVolume         types.List    `tfsdk:"app_volume" json:"appVolume"`
	BestUnit          types.String  `tfsdk:"best_unit" json:"bestUnit"`
	BundleVolume      types.List    `tfsdk:"bundle_volume" json:"bundleVolume"`
	Date              types.Int64   `tfsdk:"date"`
	MaxDailyAllowance types.Float64 `tfsdk:"max_daily_allowance" json:"maxDailyAllowance"`
	MaxDailyUsage     types.Float64 `tfsdk:"max_daily_usage" json:"maxDailyUsage"`
	PeriodVolume      types.Object  `tfsdk:"period_volume" json:"periodVolume"`
}

// NewGetPeriodVolumeDetailedByDateDataSource returns a new instance of the generated data source.
func NewGetPeriodVolumeDetailedByDateDataSource() datasource.DataSource {
	return &GetPeriodVolumeDetailedByDateDataSource{}
}

// Metadata returns the data source type name.
func (d *GetPeriodVolumeDetailedByDateDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_period_volume_detailed_by_date"
}

// Schema returns the data source schema.
func (d *GetPeriodVolumeDetailedByDateDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Gives a detailed period volume by date", Attributes: map[string]schema.Attribute{"app_volume": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"app": schema.StringAttribute{MarkdownDescription: "Name of the app of this app volume", Computed: true}, "day_volumes": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"allowance": schema.Float64Attribute{MarkdownDescription: "Allowance of this day volume in bytes", Computed: true}, "overage": schema.Float64Attribute{MarkdownDescription: "Overage on this day in bytes, usually difference between allowance and usage if positive, else 0", Computed: true}, "usage": schema.Float64Attribute{MarkdownDescription: "Usage of this day volume in bytes", Computed: true}, "usage95_p": schema.Float64Attribute{MarkdownDescription: "95th percentile usage on this day calculated with a lookback period of 90 days", Computed: true}}}}, "total_allowance": schema.Float64Attribute{MarkdownDescription: "Total allowance of this app volume in bytes", Computed: true}, "total_overage": schema.Float64Attribute{MarkdownDescription: "Total overage of this app volume in bytes", Computed: true}, "total_usage": schema.Float64Attribute{MarkdownDescription: "Total usage of this app volume in bytes", Computed: true}}}}, "best_unit": schema.StringAttribute{MarkdownDescription: "Best unit to represent the volume (TeraMegaBytes, GigaBytes,..)", Computed: true}, "bundle_volume": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"bundle": schema.StringAttribute{MarkdownDescription: "Name of the bundle of this bundle volume", Computed: true}, "day_volumes": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"allowance": schema.Float64Attribute{MarkdownDescription: "Allowance of this day volume in bytes", Computed: true}, "overage": schema.Float64Attribute{MarkdownDescription: "Overage on this day in bytes, usually difference between allowance and usage if positive, else 0", Computed: true}, "usage": schema.Float64Attribute{MarkdownDescription: "Usage of this day volume in bytes", Computed: true}, "usage95_p": schema.Float64Attribute{MarkdownDescription: "95th percentile usage on this day calculated with a lookback period of 90 days", Computed: true}}}}, "days_overage": schema.Int64Attribute{MarkdownDescription: "Number of days in the period where licensed volume limit was exceeded", Computed: true}, "max_daily_allowance": schema.Float64Attribute{MarkdownDescription: "Max daily allowance of this bundle volume in bytes", Computed: true}, "max_daily_usage": schema.Float64Attribute{MarkdownDescription: "Maximum daily usage of this bundle volume in bytes", Computed: true}, "total_allowance": schema.Float64Attribute{MarkdownDescription: "Total bundle allowance in bytes", Computed: true}, "total_overage": schema.Float64Attribute{MarkdownDescription: "Total bundle overage in bytes", Computed: true}, "total_usage": schema.Float64Attribute{MarkdownDescription: "Total bundle usage in bytes", Computed: true}}}}, "date": schema.Int64Attribute{MarkdownDescription: "Date of detailed period volume to retrieve in format YYYYMMDD", Required: true}, "max_daily_allowance": schema.Float64Attribute{MarkdownDescription: "Max daily allowance for the period in bytes", Computed: true}, "max_daily_usage": schema.Float64Attribute{MarkdownDescription: "Max daily usage for the period in bytes", Computed: true}, "period_volume": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"apps": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "bundles": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "date_codes": schema.ListAttribute{Computed: true, ElementType: types.Float64Type}, "days_completed": schema.Float64Attribute{MarkdownDescription: "Number of days completed in the current period", Computed: true}, "days_over95_p": schema.Float64Attribute{MarkdownDescription: "Days where 95th percentile usage exceeded allowance, during the period", Computed: true}, "days_overage": schema.Float64Attribute{MarkdownDescription: "Number of days with usage exceeding allowance during the current period", Computed: true}, "period": schema.StringAttribute{MarkdownDescription: "Period from start to end dates formatted", Computed: true}, "pos_ids": schema.ListAttribute{Computed: true, ElementType: types.StringType}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetPeriodVolumeDetailedByDateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetPeriodVolumeDetailedByDateDataSourceModel
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
func (d *GetPeriodVolumeDetailedByDateDataSource) readRemote(ctx context.Context, config *GetPeriodVolumeDetailedByDateDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/licensing/volumes/periodDetailed/{date}"
	reqPath = strings.ReplaceAll(reqPath, "{date}", url.PathEscape(strconv.FormatInt(config.Date.ValueInt64(), 10)))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_period_volume_detailed_by_date", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_period_volume_detailed_by_date", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_period_volume_detailed_by_date", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_period_volume_detailed_by_date", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_period_volume_detailed_by_date", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_period_volume_detailed_by_date", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_period_volume_detailed_by_date", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_period_volume_detailed_by_date", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_period_volume_detailed_by_date", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_period_volume_detailed_by_date", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_period_volume_detailed_by_date", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_period_volume_detailed_by_date", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_period_volume_detailed_by_date", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_period_volume_detailed_by_date", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetPeriodVolumeDetailedByDateDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
