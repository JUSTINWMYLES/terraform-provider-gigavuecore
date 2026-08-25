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
	_ datasource.DataSource              = (*GetAcmeCertificateDetailsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAcmeCertificateDetailsDataSource)(nil)
)

// GetAcmeCertificateDetailsDataSource is the generated Terraform data source implementation.
type GetAcmeCertificateDetailsDataSource struct {
	client *client.Client
}

// GetAcmeCertificateDetailsDataSourceModel describes the data source state shape.
type GetAcmeCertificateDetailsDataSourceModel struct {
	Items types.List `tfsdk:"items"`
}

// NewGetAcmeCertificateDetailsDataSource returns a new instance of the generated data source.
func NewGetAcmeCertificateDetailsDataSource() datasource.DataSource {
	return &GetAcmeCertificateDetailsDataSource{}
}

// Metadata returns the data source type name.
func (d *GetAcmeCertificateDetailsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_acme_certificate_details"
}

// Schema returns the data source schema.
func (d *GetAcmeCertificateDetailsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "get a list of acme certificateDetails", Attributes: map[string]schema.Attribute{"items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"cert_last_request_status": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"acme_server_alias": schema.StringAttribute{MarkdownDescription: "only for issue acmeServerAlias details will be shown", Computed: true}, "cert_status": schema.StringAttribute{Computed: true}, "domain": schema.StringAttribute{MarkdownDescription: "domain name can be FM IP or FQDN", Computed: true}, "request_type": schema.StringAttribute{Computed: true}}}, "domain": schema.StringAttribute{MarkdownDescription: "domain name will be FM IP or FQDN", Computed: true}, "issued_cert_details": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"acme_server_alias": schema.StringAttribute{Computed: true}, "algorithm": schema.StringAttribute{Computed: true}, "expiry": schema.StringAttribute{Computed: true}, "first_issued": schema.StringAttribute{MarkdownDescription: "Date of the first issued certificate for FM", Computed: true}, "last_failed_renew": schema.StringAttribute{Computed: true}, "last_success_renew": schema.StringAttribute{Computed: true}, "next_renew": schema.StringAttribute{Computed: true}, "renew_days": schema.StringAttribute{MarkdownDescription: "only user Configured renewal days is shown", Computed: true}, "status": schema.StringAttribute{Computed: true}, "task_group_id": schema.StringAttribute{Computed: true}}}, "task_id": schema.StringAttribute{Computed: true}}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetAcmeCertificateDetailsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAcmeCertificateDetailsDataSourceModel
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
func (d *GetAcmeCertificateDetailsDataSource) readListRemote(ctx context.Context, config *GetAcmeCertificateDetailsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/acme/certificate"
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_acme_certificate_details", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_acme_certificate_details", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["acmeCertificate"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_acme_certificate_details", fmt.Sprintf("Could not decode list page: missing %q array", "acmeCertificate"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_acme_certificate_details", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetAcmeCertificateDetailsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
