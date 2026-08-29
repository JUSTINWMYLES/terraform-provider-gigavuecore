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
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*GetAppsInLastNPeriodsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAppsInLastNPeriodsDataSource)(nil)
)

// GetAppsInLastNPeriodsDataSource is the generated Terraform data source implementation.
type GetAppsInLastNPeriodsDataSource struct {
	client *client.Client
}

// GetAppsInLastNPeriodsDataSourceModel describes the data source state shape.
type GetAppsInLastNPeriodsDataSourceModel struct {
	Items types.List  `tfsdk:"items"`
	N     types.Int64 `tfsdk:"n"`
}

// NewGetAppsInLastNPeriodsDataSource returns a new instance of the generated data source.
func NewGetAppsInLastNPeriodsDataSource() datasource.DataSource {
	return &GetAppsInLastNPeriodsDataSource{}
}

// Metadata returns the data source type name.
func (d *GetAppsInLastNPeriodsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_apps_in_last_n_periods"
}

// Schema returns the data source schema.
func (d *GetAppsInLastNPeriodsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Gives the set of apps for which usage has been recorded in at least one of the last 'n' periods. If 'n' is unspecified or not positive, all periods are considered", Attributes: map[string]schema.Attribute{"items": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "n": schema.Int64Attribute{MarkdownDescription: "Number of periods", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetAppsInLastNPeriodsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAppsInLastNPeriodsDataSourceModel
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
func (d *GetAppsInLastNPeriodsDataSource) readListRemote(ctx context.Context, config *GetAppsInLastNPeriodsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/licensing/volumes/appsInLastNPeriods"
	params := url.Values{}
	if !config.N.IsNull() {
		params.Set("n", strconv.FormatInt(config.N.ValueInt64(), 10))
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_apps_in_last_n_periods", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageItems := []any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageItems); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_apps_in_last_n_periods", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_apps_in_last_n_periods", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetAppsInLastNPeriodsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
