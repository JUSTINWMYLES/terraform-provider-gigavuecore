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
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*GetLastNPeriodsSummaryWithUnitDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetLastNPeriodsSummaryWithUnitDataSource)(nil)
)

// GetLastNPeriodsSummaryWithUnitDataSource is the generated Terraform data source implementation.
type GetLastNPeriodsSummaryWithUnitDataSource struct {
	client *client.Client
}

// GetLastNPeriodsSummaryWithUnitDataSourceModel describes the data source state shape.
type GetLastNPeriodsSummaryWithUnitDataSourceModel struct {
	LowestUnit types.String `tfsdk:"lowest_unit" json:"lowestUnit"`
	N          types.Int64  `tfsdk:"n"`
	Summaries  types.List   `tfsdk:"summaries"`
}

// NewGetLastNPeriodsSummaryWithUnitDataSource returns a new instance of the generated data source.
func NewGetLastNPeriodsSummaryWithUnitDataSource() datasource.DataSource {
	return &GetLastNPeriodsSummaryWithUnitDataSource{}
}

// Metadata returns the data source type name.
func (d *GetLastNPeriodsSummaryWithUnitDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_last_n_periods_summary_with_unit"
}

// Schema returns the data source schema.
func (d *GetLastNPeriodsSummaryWithUnitDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Gives the summary of total allowance, total usage, total overage for each app and bundle for each period, for the last 'n' periods, along with lowest of the best units for each period summary; for bundles, the maximum daily allowance and usage and the number of days of overage are also provided", Attributes: map[string]schema.Attribute{"lowest_unit": schema.StringAttribute{Computed: true}, "n": schema.Int64Attribute{MarkdownDescription: "Number of periods", Required: true}, "summaries": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"app_volume": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"app": schema.StringAttribute{MarkdownDescription: "Name of the app of this app volume", Computed: true}, "total_allowance": schema.Float64Attribute{MarkdownDescription: "Total allowance of this app volume in bytes", Computed: true}, "total_overage": schema.Float64Attribute{MarkdownDescription: "Total overage of this app volume in bytes", Computed: true}, "total_usage": schema.Float64Attribute{MarkdownDescription: "Total usage of this app volume in bytes", Computed: true}}}}, "best_unit": schema.StringAttribute{MarkdownDescription: "Best unit to represent the volume (TeraBytes, GigaBytes,..)", Computed: true}, "bundle_volume": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"bundle": schema.StringAttribute{MarkdownDescription: "Name of the license bundle", Computed: true}, "days_overage": schema.Int64Attribute{MarkdownDescription: "Number of days in the period where licensed volume allowance was exceeded", Computed: true}, "max_daily_allowance": schema.Float64Attribute{MarkdownDescription: "Max daily allowance of this bundle in bytes", Computed: true}, "max_daily_usage": schema.Float64Attribute{MarkdownDescription: "Maximum daily usage of this bundle in bytes", Computed: true}, "total_allowance": schema.Float64Attribute{MarkdownDescription: "Total bundle allowance for the period in bytes", Computed: true}, "total_overage": schema.Float64Attribute{MarkdownDescription: "Total bundle overage for the period in bytes", Computed: true}, "total_usage": schema.Float64Attribute{MarkdownDescription: "Total bundle usage for the period in bytes", Computed: true}}}}, "period_volume": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"apps": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "bundles": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "date_codes": schema.ListAttribute{Computed: true, ElementType: types.Float64Type}, "days_completed": schema.Float64Attribute{MarkdownDescription: "Number of days completed in the current period", Computed: true}, "days_over95_p": schema.Float64Attribute{MarkdownDescription: "Days where 95th percentile usage exceeded allowance, during the period", Computed: true}, "days_overage": schema.Float64Attribute{MarkdownDescription: "Number of days with usage exceeding allowance during the current period", Computed: true}, "period": schema.StringAttribute{MarkdownDescription: "Period from start to end dates formatted", Computed: true}, "pos_ids": schema.ListAttribute{Computed: true, ElementType: types.StringType}}}}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetLastNPeriodsSummaryWithUnitDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetLastNPeriodsSummaryWithUnitDataSourceModel
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
func (d *GetLastNPeriodsSummaryWithUnitDataSource) readRemote(ctx context.Context, config *GetLastNPeriodsSummaryWithUnitDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/licensing/volumes/lastNPeriodsSummaryWithUnit"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_last_n_periods_summary_with_unit", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("n", strconv.FormatInt(config.N.ValueInt64(), 10))
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_last_n_periods_summary_with_unit", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_last_n_periods_summary_with_unit", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_last_n_periods_summary_with_unit", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_last_n_periods_summary_with_unit", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_last_n_periods_summary_with_unit", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_last_n_periods_summary_with_unit", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_last_n_periods_summary_with_unit", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_last_n_periods_summary_with_unit", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_last_n_periods_summary_with_unit", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_last_n_periods_summary_with_unit", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_last_n_periods_summary_with_unit", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_last_n_periods_summary_with_unit", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_last_n_periods_summary_with_unit", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetLastNPeriodsSummaryWithUnitDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
