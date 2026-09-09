package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*LoadSystemAcmeCertificateDetailsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadSystemAcmeCertificateDetailsDataSource)(nil)
)

// LoadSystemAcmeCertificateDetailsDataSource is the generated Terraform data source implementation.
type LoadSystemAcmeCertificateDetailsDataSource struct {
	client *client.Client
}

// LoadSystemAcmeCertificateDetailsDataSourceModel describes the data source state shape.
type LoadSystemAcmeCertificateDetailsDataSourceModel struct {
	BoxId           types.String `tfsdk:"box_id" json:"boxId"`
	ClusterId       types.String `tfsdk:"cluster_id" json:"clusterId"`
	ClusterName     types.String `tfsdk:"cluster_name" json:"clusterName"`
	IssuedCert      types.Object `tfsdk:"issued_cert" json:"issuedCert"`
	LastAcmeRequest types.Object `tfsdk:"last_acme_request" json:"lastAcmeRequest"`
}

// NewLoadSystemAcmeCertificateDetailsDataSource returns a new instance of the generated data source.
func NewLoadSystemAcmeCertificateDetailsDataSource() datasource.DataSource {
	return &LoadSystemAcmeCertificateDetailsDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadSystemAcmeCertificateDetailsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_system_acme_certificate_details"
}

// Schema returns the data source schema.
func (d *LoadSystemAcmeCertificateDetailsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "get deviceacme certificate details", Attributes: map[string]schema.Attribute{"box_id": schema.StringAttribute{Computed: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "cluster_name": schema.StringAttribute{Computed: true}, "issued_cert": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"acme_ca_url": schema.StringAttribute{Computed: true}, "acme_service": schema.StringAttribute{Computed: true}, "algorithm": schema.StringAttribute{Computed: true}, "cert_name": schema.StringAttribute{Computed: true}, "domain": schema.StringAttribute{Computed: true}, "expiry": schema.StringAttribute{Computed: true}, "first_issued": schema.StringAttribute{Computed: true}, "last_failed_renew": schema.StringAttribute{Computed: true}, "last_success_renew": schema.StringAttribute{Computed: true}, "next_renew": schema.StringAttribute{Computed: true}, "renew_days": schema.Int64Attribute{MarkdownDescription: "default will be 1/3rd of certificate validity period", Computed: true}, "status": schema.StringAttribute{Computed: true}}}, "last_acme_request": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"acme_ca_url": schema.StringAttribute{Computed: true}, "domain_name": schema.StringAttribute{Computed: true}, "status": schema.StringAttribute{Computed: true}, "type": schema.StringAttribute{Computed: true}}}}}
}

// Read fetches remote state into the data source model.
func (d *LoadSystemAcmeCertificateDetailsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadSystemAcmeCertificateDetailsDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	d.readRemote(ctx, &config, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}

// readRemote performs the read HTTP exchange and decodes the response into config. Extracted from Read so the request/response logic is unit-testable without a tfsdk.Config.
func (d *LoadSystemAcmeCertificateDetailsDataSource) readRemote(ctx context.Context, config *LoadSystemAcmeCertificateDetailsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/acme/certificate/nodes/{clusterId}"
	reqPath = strings.ReplaceAll(reqPath, "{clusterId}", url.PathEscape(config.ClusterId.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_system_acme_certificate_details", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_system_acme_certificate_details", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_system_acme_certificate_details", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_system_acme_certificate_details", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_system_acme_certificate_details", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_system_acme_certificate_details", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_system_acme_certificate_details", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_system_acme_certificate_details", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_system_acme_certificate_details", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_system_acme_certificate_details", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_load_system_acme_certificate_details", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_load_system_acme_certificate_details", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_system_acme_certificate_details", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_system_acme_certificate_details", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadSystemAcmeCertificateDetailsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
