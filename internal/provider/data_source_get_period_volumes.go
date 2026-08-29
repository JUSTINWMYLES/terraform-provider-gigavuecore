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
	_ datasource.DataSource              = (*GetPeriodVolumesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetPeriodVolumesDataSource)(nil)
)

// GetPeriodVolumesDataSource is the generated Terraform data source implementation.
type GetPeriodVolumesDataSource struct {
	client *client.Client
}

// GetPeriodVolumesDataSourceModel describes the data source state shape.
type GetPeriodVolumesDataSourceModel struct {
	Items types.List `tfsdk:"items"`
}

// NewGetPeriodVolumesDataSource returns a new instance of the generated data source.
func NewGetPeriodVolumesDataSource() datasource.DataSource {
	return &GetPeriodVolumesDataSource{}
}

// Metadata returns the data source type name.
func (d *GetPeriodVolumesDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_period_volumes"
}

// Schema returns the data source schema.
func (d *GetPeriodVolumesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Gives a list of all period volumes", Attributes: map[string]schema.Attribute{"items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"apps": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "bundles": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "date_codes": schema.ListAttribute{Computed: true, ElementType: types.Float64Type}, "days_completed": schema.Float64Attribute{MarkdownDescription: "Number of days completed in the current period", Computed: true}, "days_over95_p": schema.Float64Attribute{MarkdownDescription: "Days where 95th percentile usage exceeded allowance, during the period", Computed: true}, "days_overage": schema.Float64Attribute{MarkdownDescription: "Number of days with usage exceeding allowance during the current period", Computed: true}, "period": schema.StringAttribute{MarkdownDescription: "Period from start to end dates formatted", Computed: true}, "pos_ids": schema.ListAttribute{Computed: true, ElementType: types.StringType}}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetPeriodVolumesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetPeriodVolumesDataSourceModel
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
func (d *GetPeriodVolumesDataSource) readListRemote(ctx context.Context, config *GetPeriodVolumesDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/licensing/volumes/period"
	params := url.Values{}
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_period_volumes", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_period_volumes", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["volumes"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_period_volumes", fmt.Sprintf("Could not decode list page: missing %q array", "volumes"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_period_volumes", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetPeriodVolumesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
