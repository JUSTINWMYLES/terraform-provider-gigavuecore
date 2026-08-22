package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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
	_ datasource.DataSource              = (*GetPolicyNrtStatsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetPolicyNrtStatsDataSource)(nil)
)

// GetPolicyNrtStatsDataSource is the generated Terraform data source implementation.
type GetPolicyNrtStatsDataSource struct {
	client *client.Client
}

// GetPolicyNrtStatsDataSourceModel describes the data source state shape.
type GetPolicyNrtStatsDataSourceModel struct {
	Alias types.String  `tfsdk:"alias"`
	Items types.Dynamic `tfsdk:"items"`
}

// NewGetPolicyNrtStatsDataSource returns a new instance of the generated data source.
func NewGetPolicyNrtStatsDataSource() datasource.DataSource {
	return &GetPolicyNrtStatsDataSource{}
}

// Metadata returns the data source type name.
func (d *GetPolicyNrtStatsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_policy_nrt_stats"
}

// Schema returns the data source schema.
func (d *GetPolicyNrtStatsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get near real-time statistics for a traffic flow", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Traffic flow alias", Required: true}, "items": schema.DynamicAttribute{Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetPolicyNrtStatsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetPolicyNrtStatsDataSourceModel
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
func (d *GetPolicyNrtStatsDataSource) readListRemote(ctx context.Context, config *GetPolicyNrtStatsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/traffic-flows/nrtStats/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_policy_nrt_stats", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_policy_nrt_stats", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["fmNearRealTimeCachedStatisticsList"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_policy_nrt_stats", fmt.Sprintf("Could not decode list page: missing %q array", "fmNearRealTimeCachedStatisticsList"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_policy_nrt_stats", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetPolicyNrtStatsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
