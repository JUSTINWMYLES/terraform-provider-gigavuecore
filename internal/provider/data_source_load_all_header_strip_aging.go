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
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*LoadAllHeaderStripAgingDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadAllHeaderStripAgingDataSource)(nil)
)

// LoadAllHeaderStripAgingDataSource is the generated Terraform data source implementation.
type LoadAllHeaderStripAgingDataSource struct {
	client *client.Client
}

// LoadAllHeaderStripAgingDataSourceModel describes the data source state shape.
type LoadAllHeaderStripAgingDataSourceModel struct {
	BoxId types.String `tfsdk:"box_id" json:"boxId"`
	Items types.List   `tfsdk:"items"`
	Page  types.String `tfsdk:"page"`
	Sort  types.String `tfsdk:"sort"`
}

// NewLoadAllHeaderStripAgingDataSource returns a new instance of the generated data source.
func NewLoadAllHeaderStripAgingDataSource() datasource.DataSource {
	return &LoadAllHeaderStripAgingDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadAllHeaderStripAgingDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_all_header_strip_aging"
}

// Schema returns the data source schema.
func (d *LoadAllHeaderStripAgingDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load header strip aging for all boxes", Attributes: map[string]schema.Attribute{"box_id": schema.StringAttribute{MarkdownDescription: "device box id. valid range 1 - 64.", Optional: true}, "items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"aging_interval": schema.Int64Attribute{MarkdownDescription: "Interval in sec. Valid range 300-1000000. Enter 0 to disable.", Computed: true}, "box_id": schema.StringAttribute{MarkdownDescription: "device box id. valid range 1 - 64. all is applicable only for post request.", Computed: true}, "dst_port": schema.Int64Attribute{MarkdownDescription: "L4 destination port number.Valid value is between 0 to 65535.", Computed: true}, "protocol_type": schema.StringAttribute{MarkdownDescription: "protocol type", Computed: true}}}}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadAllHeaderStripAgingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadAllHeaderStripAgingDataSourceModel
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
func (d *LoadAllHeaderStripAgingDataSource) readListRemote(ctx context.Context, config *LoadAllHeaderStripAgingDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/headerStripAging"
	params := url.Values{}
	if !config.BoxId.IsNull() {
		params.Set("boxId", config.BoxId.ValueString())
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
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_header_strip_aging", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_header_strip_aging", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["headerStripsAgingDef"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_header_strip_aging", fmt.Sprintf("Could not decode list page: missing %q array", "headerStripsAgingDef"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_header_strip_aging", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadAllHeaderStripAgingDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
