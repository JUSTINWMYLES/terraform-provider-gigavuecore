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
	_ datasource.DataSource              = (*GetCryptoCaDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetCryptoCaDataSource)(nil)
)

// GetCryptoCaDataSource is the generated Terraform data source implementation.
type GetCryptoCaDataSource struct {
	client *client.Client
}

// GetCryptoCaDataSourceModel describes the data source state shape.
type GetCryptoCaDataSourceModel struct {
	Algorithm   types.String `tfsdk:"algorithm"`
	ClusterId   types.String `tfsdk:"cluster_id" json:"clusterId"`
	Comment     types.String `tfsdk:"comment"`
	IssuerName  types.String `tfsdk:"issuer_name" json:"issuerName"`
	Items       types.List   `tfsdk:"items"`
	Name        types.String `tfsdk:"name"`
	Page        types.String `tfsdk:"page"`
	Sort        types.String `tfsdk:"sort"`
	SubjectName types.String `tfsdk:"subject_name" json:"subjectName"`
	ValidFrom   types.String `tfsdk:"valid_from" json:"validFrom"`
	ValidTill   types.String `tfsdk:"valid_till" json:"validTill"`
}

// NewGetCryptoCaDataSource returns a new instance of the generated data source.
func NewGetCryptoCaDataSource() datasource.DataSource {
	return &GetCryptoCaDataSource{}
}

// Metadata returns the data source type name.
func (d *GetCryptoCaDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_crypto_ca"
}

// Schema returns the data source schema.
func (d *GetCryptoCaDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "get a list of configured trusted certificate authorities (CA List)", Attributes: map[string]schema.Attribute{"algorithm": schema.StringAttribute{MarkdownDescription: "Certificate Public key algorithm (filter param)", Optional: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "comment": schema.StringAttribute{MarkdownDescription: "Certificate description (filter param)", Optional: true}, "issuer_name": schema.StringAttribute{MarkdownDescription: "Issuer's common Name (filter param)", Optional: true}, "items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"comment": schema.StringAttribute{Computed: true}, "default_cert": schema.BoolAttribute{MarkdownDescription: "If true, this certificate is the default server identity", Computed: true}, "issuer": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"common_name": schema.StringAttribute{MarkdownDescription: "issuer and subject common name (e.g. a domain name)", Computed: true}, "country": schema.StringAttribute{MarkdownDescription: "two-alphanumeric-character country code", Computed: true}, "email": schema.StringAttribute{Computed: true}, "locality": schema.StringAttribute{Computed: true}, "org_name": schema.StringAttribute{Computed: true}, "org_unit": schema.StringAttribute{Computed: true}, "state": schema.StringAttribute{MarkdownDescription: "state or provence name", Computed: true}}}, "name": schema.StringAttribute{Computed: true}, "pem": schema.StringAttribute{Computed: true}, "private_key": schema.BoolAttribute{Computed: true}, "serial_number": schema.StringAttribute{MarkdownDescription: "lower-case hexadecimal serial number not prefixed with '0x'", Computed: true}, "sha1_fingerprint": schema.StringAttribute{Computed: true}, "signature_algorithm": schema.StringAttribute{Computed: true}, "subject": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"common_name": schema.StringAttribute{MarkdownDescription: "issuer and subject common name (e.g. a domain name)", Computed: true}, "country": schema.StringAttribute{MarkdownDescription: "two-alphanumeric-character country code", Computed: true}, "email": schema.StringAttribute{Computed: true}, "locality": schema.StringAttribute{Computed: true}, "org_name": schema.StringAttribute{Computed: true}, "org_unit": schema.StringAttribute{Computed: true}, "state": schema.StringAttribute{MarkdownDescription: "state or provence name", Computed: true}}}, "subject_public_key_algorithm": schema.StringAttribute{Computed: true}, "subject_public_key_length": schema.Int64Attribute{Computed: true}, "validity": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"not_after": schema.StringAttribute{MarkdownDescription: "date and time when certificate stops being valid ([rfc3339](https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14))", Computed: true}, "not_before": schema.StringAttribute{MarkdownDescription: "date and time when certificate starts being valid ([rfc3339](https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14))", Computed: true}}}, "version": schema.StringAttribute{Computed: true}}}}, "name": schema.StringAttribute{MarkdownDescription: "Target Certificate name (filter param)", Optional: true}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}, "subject_name": schema.StringAttribute{MarkdownDescription: "Subject's common Name (filter param)", Optional: true}, "valid_from": schema.StringAttribute{MarkdownDescription: "Certificate Validity details (filter param)", Optional: true}, "valid_till": schema.StringAttribute{MarkdownDescription: "Certificate Validity details (filter param)", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetCryptoCaDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetCryptoCaDataSourceModel
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
func (d *GetCryptoCaDataSource) readListRemote(ctx context.Context, config *GetCryptoCaDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/ssl/nodeCertificate/trustedCA"
	params := url.Values{}
	params.Set("clusterId", config.ClusterId.ValueString())
	if !config.Name.IsNull() {
		params.Set("name", config.Name.ValueString())
	}
	if !config.Comment.IsNull() {
		params.Set("comment", config.Comment.ValueString())
	}
	if !config.SubjectName.IsNull() {
		params.Set("subjectName", config.SubjectName.ValueString())
	}
	if !config.IssuerName.IsNull() {
		params.Set("issuerName", config.IssuerName.ValueString())
	}
	if !config.Algorithm.IsNull() {
		params.Set("algorithm", config.Algorithm.ValueString())
	}
	if !config.ValidFrom.IsNull() {
		params.Set("validFrom", config.ValidFrom.ValueString())
	}
	if !config.ValidTill.IsNull() {
		params.Set("validTill", config.ValidTill.ValueString())
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_crypto_ca", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_crypto_ca", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["caList"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_crypto_ca", fmt.Sprintf("Could not decode list page: missing %q array", "caList"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_crypto_ca", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetCryptoCaDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
