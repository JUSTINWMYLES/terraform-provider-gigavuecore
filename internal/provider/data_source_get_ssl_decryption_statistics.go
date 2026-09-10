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
	_ datasource.DataSource              = (*GetSslDecryptionStatisticsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetSslDecryptionStatisticsDataSource)(nil)
)

// GetSslDecryptionStatisticsDataSource is the generated Terraform data source implementation.
type GetSslDecryptionStatisticsDataSource struct {
	client *client.Client
}

// GetSslDecryptionStatisticsDataSourceModel describes the data source state shape.
type GetSslDecryptionStatisticsDataSourceModel struct {
	Alias                       types.String `tfsdk:"alias"`
	ClusterId                   types.String `tfsdk:"cluster_id" json:"clusterId"`
	Gsgroup                     types.String `tfsdk:"gsgroup"`
	Hostname                    types.String `tfsdk:"hostname"`
	SslDecryptionSessionDetails types.List   `tfsdk:"ssl_decryption_session_details" json:"sslDecryptionSessionDetails"`
}

// NewGetSslDecryptionStatisticsDataSource returns a new instance of the generated data source.
func NewGetSslDecryptionStatisticsDataSource() datasource.DataSource {
	return &GetSslDecryptionStatisticsDataSource{}
}

// Metadata returns the data source type name.
func (d *GetSslDecryptionStatisticsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_ssl_decryption_statistics"
}

// Schema returns the data source schema.
func (d *GetSslDecryptionStatisticsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load Ssl Decryption Report Statistics", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "alias of the target GS Group", Required: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "gsgroup": schema.StringAttribute{MarkdownDescription: "alias of gsgroup", Computed: true}, "hostname": schema.StringAttribute{MarkdownDescription: "match based on hostname", Optional: true}, "ssl_decryption_session_details": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"cipher_suite": schema.StringAttribute{Computed: true}, "client_ip": schema.StringAttribute{Computed: true}, "client_port": schema.Int64Attribute{Computed: true}, "decryption_status": schema.StringAttribute{Computed: true}, "duration": schema.Int64Attribute{Computed: true}, "fingerprint": schema.StringAttribute{Computed: true}, "first_error": schema.Int64Attribute{Computed: true}, "first_error_reason": schema.StringAttribute{Computed: true}, "in_pkts": schema.Int64Attribute{Computed: true}, "out_pkts": schema.Int64Attribute{Computed: true}, "server_ip": schema.StringAttribute{Computed: true}, "server_port": schema.Int64Attribute{Computed: true}, "sni": schema.StringAttribute{Computed: true}, "start_time": schema.StringAttribute{Computed: true}, "version": schema.StringAttribute{Computed: true}}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetSslDecryptionStatisticsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetSslDecryptionStatisticsDataSourceModel
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
func (d *GetSslDecryptionStatisticsDataSource) readRemote(ctx context.Context, config *GetSslDecryptionStatisticsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/gsGroups/{alias}/flowOpsReport/sslDecryption/statistics"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_statistics", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	if !config.Hostname.IsNull() {
		query.Set("hostname", config.Hostname.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_statistics", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_statistics", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_statistics", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_statistics", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_statistics", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_statistics", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_statistics", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_statistics", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_statistics", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_statistics", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_statistics", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_statistics", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["sslDecryptionReportStats"]; ok {
		if m, ok := v.(map[string]any); ok {
			data = m
		} else if arr, ok := v.([]any); ok && len(arr) > 0 {
			if m, ok := arr[0].(map[string]any); ok {
				data = m
			}
		}
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_statistics", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetSslDecryptionStatisticsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
