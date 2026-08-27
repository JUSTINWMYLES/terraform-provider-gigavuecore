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
	_ datasource.DataSource              = (*LoadAllSystemAcmeCertificateDetailsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadAllSystemAcmeCertificateDetailsDataSource)(nil)
)

// LoadAllSystemAcmeCertificateDetailsDataSource is the generated Terraform data source implementation.
type LoadAllSystemAcmeCertificateDetailsDataSource struct {
	client *client.Client
}

// LoadAllSystemAcmeCertificateDetailsDataSourceModel describes the data source state shape.
type LoadAllSystemAcmeCertificateDetailsDataSourceModel struct {
	Items types.List `tfsdk:"items"`
}

// NewLoadAllSystemAcmeCertificateDetailsDataSource returns a new instance of the generated data source.
func NewLoadAllSystemAcmeCertificateDetailsDataSource() datasource.DataSource {
	return &LoadAllSystemAcmeCertificateDetailsDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadAllSystemAcmeCertificateDetailsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_all_system_acme_certificate_details"
}

// Schema returns the data source schema.
func (d *LoadAllSystemAcmeCertificateDetailsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "get a list of deviceacme certificate details", Attributes: map[string]schema.Attribute{"items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"box_id": schema.StringAttribute{Computed: true}, "cluster_name": schema.StringAttribute{Computed: true}, "issued_cert": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"acme_ca_url": schema.StringAttribute{Computed: true}, "acme_service": schema.StringAttribute{Computed: true}, "algorithm": schema.StringAttribute{Computed: true}, "cert_name": schema.StringAttribute{Computed: true}, "domain": schema.StringAttribute{Computed: true}, "expiry": schema.StringAttribute{Computed: true}, "first_issued": schema.StringAttribute{Computed: true}, "last_failed_renew": schema.StringAttribute{Computed: true}, "last_success_renew": schema.StringAttribute{Computed: true}, "next_renew": schema.StringAttribute{Computed: true}, "renew_days": schema.Int64Attribute{MarkdownDescription: "default will be 1/3rd of certificate validity period", Computed: true}, "status": schema.StringAttribute{Computed: true}}}, "last_acme_request": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"acme_ca_url": schema.StringAttribute{Computed: true}, "domain_name": schema.StringAttribute{Computed: true}, "status": schema.StringAttribute{Computed: true}, "type": schema.StringAttribute{Computed: true}}}}}}}}
}

// Read fetches remote state into the data source model.
func (d *LoadAllSystemAcmeCertificateDetailsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadAllSystemAcmeCertificateDetailsDataSourceModel
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
func (d *LoadAllSystemAcmeCertificateDetailsDataSource) readListRemote(ctx context.Context, config *LoadAllSystemAcmeCertificateDetailsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/acme/certificate/nodes"
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
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_system_acme_certificate_details", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_system_acme_certificate_details", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["certificateInfo"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_system_acme_certificate_details", fmt.Sprintf("Could not decode list page: missing %q array", "certificateInfo"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_system_acme_certificate_details", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadAllSystemAcmeCertificateDetailsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
