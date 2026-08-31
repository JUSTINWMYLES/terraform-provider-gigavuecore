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
	_ datasource.DataSource              = (*LoadFmLicensesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadFmLicensesDataSource)(nil)
)

// LoadFmLicensesDataSource is the generated Terraform data source implementation.
type LoadFmLicensesDataSource struct {
	client *client.Client
}

// LoadFmLicensesDataSourceModel describes the data source state shape.
type LoadFmLicensesDataSourceModel struct {
	Items types.List `tfsdk:"items"`
}

// NewLoadFmLicensesDataSource returns a new instance of the generated data source.
func NewLoadFmLicensesDataSource() datasource.DataSource {
	return &LoadFmLicensesDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadFmLicensesDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_fm_licenses"
}

// Schema returns the data source schema.
func (d *LoadFmLicensesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load FM Licenses", Attributes: map[string]schema.Attribute{"items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"active": schema.BoolAttribute{MarkdownDescription: "license is valid, matched target FM instance and within the valid time period", Computed: true}, "description": schema.StringAttribute{MarkdownDescription: "Description of the SKU code this license is issued for", Computed: true}, "end_date": schema.StringAttribute{MarkdownDescription: "End date in ISO 8601 format. If omitted, license never expires", Computed: true}, "inactive_reason": schema.StringAttribute{MarkdownDescription: "If license is inactive, this field contains the reason", Computed: true}, "license_key": schema.StringAttribute{Computed: true}, "revoked": schema.BoolAttribute{MarkdownDescription: "indicates whether the license key is revoked", Computed: true}, "start_date": schema.StringAttribute{MarkdownDescription: "Start date in ISO 8601 format. If omitted, license is effective from the date of issue", Computed: true}, "target_fm_id": schema.StringAttribute{MarkdownDescription: "identifies FM instance this license is issued for", Computed: true}, "valid": schema.BoolAttribute{MarkdownDescription: "license is well-formed, not revoked and matches the target FM instance", Computed: true}, "well_formed": schema.BoolAttribute{MarkdownDescription: "indicates that FM is able to parse and interpret the license key string", Computed: true}}}}}}
}

// Read fetches remote state into the data source model.
func (d *LoadFmLicensesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadFmLicensesDataSourceModel
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
func (d *LoadFmLicensesDataSource) readListRemote(ctx context.Context, config *LoadFmLicensesDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/licensing/fm/licenses"
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
		resp.Diagnostics.AddError("Error reading gigavuecore_load_fm_licenses", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_fm_licenses", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["licenses"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_fm_licenses", fmt.Sprintf("Could not decode list page: missing %q array", "licenses"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_fm_licenses", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadFmLicensesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
