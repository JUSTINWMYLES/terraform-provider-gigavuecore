package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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
	_ datasource.DataSource              = (*GetLastNPeriodsSummaryVolumesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetLastNPeriodsSummaryVolumesDataSource)(nil)
)

// GetLastNPeriodsSummaryVolumesDataSource is the generated Terraform data source implementation.
type GetLastNPeriodsSummaryVolumesDataSource struct {
	client *client.Client
}

// GetLastNPeriodsSummaryVolumesDataSourceModel describes the data source state shape.
type GetLastNPeriodsSummaryVolumesDataSourceModel struct {
	Items types.List  `tfsdk:"items"`
	Num   types.Int64 `tfsdk:"num"`
}

// NewGetLastNPeriodsSummaryVolumesDataSource returns a new instance of the generated data source.
func NewGetLastNPeriodsSummaryVolumesDataSource() datasource.DataSource {
	return &GetLastNPeriodsSummaryVolumesDataSource{}
}

// Metadata returns the data source type name.
func (d *GetLastNPeriodsSummaryVolumesDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_last_n_periods_summary_volumes"
}

// Schema returns the data source schema.
func (d *GetLastNPeriodsSummaryVolumesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Gives the summary of total Allowance, totalUsage, totalOverage for all apps for each period in the set of periods determined by num", Attributes: map[string]schema.Attribute{"items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"app_volume": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"app": schema.StringAttribute{MarkdownDescription: "Name of the app of this app volume", Computed: true}, "total_allowance": schema.Float64Attribute{MarkdownDescription: "Total allowance of this app volume in bytes", Computed: true}, "total_overage": schema.Float64Attribute{MarkdownDescription: "Total overage of this app volume in bytes", Computed: true}, "total_usage": schema.Float64Attribute{MarkdownDescription: "Total usage of this app volume in bytes", Computed: true}}}}, "best_unit": schema.StringAttribute{MarkdownDescription: "Best unit to represent the volume (TeraBytes, GigaBytes,..)", Computed: true}, "bundle_volume": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"bundle": schema.StringAttribute{MarkdownDescription: "Name of the license bundle", Computed: true}, "days_overage": schema.Int64Attribute{MarkdownDescription: "Number of days in the period where licensed volume allowance was exceeded", Computed: true}, "max_daily_allowance": schema.Float64Attribute{MarkdownDescription: "Max daily allowance of this bundle in bytes", Computed: true}, "max_daily_usage": schema.Float64Attribute{MarkdownDescription: "Maximum daily usage of this bundle in bytes", Computed: true}, "total_allowance": schema.Float64Attribute{MarkdownDescription: "Total bundle allowance for the period in bytes", Computed: true}, "total_overage": schema.Float64Attribute{MarkdownDescription: "Total bundle overage for the period in bytes", Computed: true}, "total_usage": schema.Float64Attribute{MarkdownDescription: "Total bundle usage for the period in bytes", Computed: true}}}}, "period_volume": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"apps": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "bundles": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "date_codes": schema.ListAttribute{Computed: true, ElementType: types.Float64Type}, "days_completed": schema.Float64Attribute{MarkdownDescription: "Number of days completed in the current period", Computed: true}, "days_over95_p": schema.Float64Attribute{MarkdownDescription: "Days where 95th percentile usage exceeded allowance, during the period", Computed: true}, "days_overage": schema.Float64Attribute{MarkdownDescription: "Number of days with usage exceeding allowance during the current period", Computed: true}, "period": schema.StringAttribute{MarkdownDescription: "Period from start to end dates formatted", Computed: true}, "pos_ids": schema.ListAttribute{Computed: true, ElementType: types.StringType}}}}}}, "num": schema.Int64Attribute{MarkdownDescription: "Number of periods", Required: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetLastNPeriodsSummaryVolumesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetLastNPeriodsSummaryVolumesDataSourceModel
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
func (d *GetLastNPeriodsSummaryVolumesDataSource) readListRemote(ctx context.Context, config *GetLastNPeriodsSummaryVolumesDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/licensing/volumes/lastNPeriodsSummary"
	params := url.Values{}
	params.Set("num", strconv.FormatInt(config.Num.ValueInt64(), 10))
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_last_n_periods_summary_volumes", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_last_n_periods_summary_volumes", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["periods"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_last_n_periods_summary_volumes", fmt.Sprintf("Could not decode list page: missing %q array", "periods"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_last_n_periods_summary_volumes", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetLastNPeriodsSummaryVolumesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
