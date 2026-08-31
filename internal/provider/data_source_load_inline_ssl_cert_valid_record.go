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
	_ datasource.DataSource              = (*LoadInlineSslCertValidRecordDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadInlineSslCertValidRecordDataSource)(nil)
)

// LoadInlineSslCertValidRecordDataSource is the generated Terraform data source implementation.
type LoadInlineSslCertValidRecordDataSource struct {
	client *client.Client
}

// LoadInlineSslCertValidRecordDataSourceModel describes the data source state shape.
type LoadInlineSslCertValidRecordDataSourceModel struct {
	Fingerprint types.String `tfsdk:"fingerprint"`
	Items       types.List   `tfsdk:"items"`
	Page        types.String `tfsdk:"page"`
	Sort        types.String `tfsdk:"sort"`
}

// NewLoadInlineSslCertValidRecordDataSource returns a new instance of the generated data source.
func NewLoadInlineSslCertValidRecordDataSource() datasource.DataSource {
	return &LoadInlineSslCertValidRecordDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadInlineSslCertValidRecordDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_inline_ssl_cert_valid_record"
}

// Schema returns the data source schema.
func (d *LoadInlineSslCertValidRecordDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get Inline SSL certificate validation record", Attributes: map[string]schema.Attribute{"fingerprint": schema.StringAttribute{MarkdownDescription: "The fingerprint of the record to lookup", Required: true}, "items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"cert_cn": schema.StringAttribute{Computed: true}, "expiry": schema.StringAttribute{Computed: true}, "revocation_status": schema.StringAttribute{Computed: true}, "sha1": schema.StringAttribute{Computed: true}}}}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadInlineSslCertValidRecordDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadInlineSslCertValidRecordDataSourceModel
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
func (d *LoadInlineSslCertValidRecordDataSource) readListRemote(ctx context.Context, config *LoadInlineSslCertValidRecordDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/inlineSsl/caching/certValidation/{fingerprint}"
	reqPath = strings.ReplaceAll(reqPath, "{fingerprint}", url.PathEscape(config.Fingerprint.ValueString()))
	params := url.Values{}
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
		resp.Diagnostics.AddError("Error reading gigavuecore_load_inline_ssl_cert_valid_record", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_inline_ssl_cert_valid_record", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["records"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_inline_ssl_cert_valid_record", fmt.Sprintf("Could not decode list page: missing %q array", "records"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_inline_ssl_cert_valid_record", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadInlineSslCertValidRecordDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
