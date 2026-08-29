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
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*GetSslDecryptionSummaryDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetSslDecryptionSummaryDataSource)(nil)
)

// GetSslDecryptionSummaryDataSource is the generated Terraform data source implementation.
type GetSslDecryptionSummaryDataSource struct {
	client *client.Client
}

// GetSslDecryptionSummaryDataSourceModel describes the data source state shape.
type GetSslDecryptionSummaryDataSourceModel struct {
	Alias        types.String `tfsdk:"alias"`
	ClusterId    types.String `tfsdk:"cluster_id" json:"clusterId"`
	DeviceIp     types.String `tfsdk:"device_ip" json:"deviceIp"`
	DeviceMask   types.String `tfsdk:"device_mask" json:"deviceMask"`
	Gsgroup      types.String `tfsdk:"gsgroup"`
	SessionIds   types.Int64  `tfsdk:"session_ids" json:"sessionIds"`
	Ssl30Session types.Int64  `tfsdk:"ssl30_session" json:"ssl30Session"`
	Tickets      types.Int64  `tfsdk:"tickets"`
	Tls10Session types.Int64  `tfsdk:"tls10_session" json:"tls10Session"`
	Tls11Session types.Int64  `tfsdk:"tls11_session" json:"tls11Session"`
	Tls12Session types.Int64  `tfsdk:"tls12_session" json:"tls12Session"`
	TotalSession types.Int64  `tfsdk:"total_session" json:"totalSession"`
}

// NewGetSslDecryptionSummaryDataSource returns a new instance of the generated data source.
func NewGetSslDecryptionSummaryDataSource() datasource.DataSource {
	return &GetSslDecryptionSummaryDataSource{}
}

// Metadata returns the data source type name.
func (d *GetSslDecryptionSummaryDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_ssl_decryption_summary"
}

// Schema returns the data source schema.
func (d *GetSslDecryptionSummaryDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load Ssl Decryption Report Summary", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "alias of the target GS Group", Required: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "device_ip": schema.StringAttribute{MarkdownDescription: "deviceIp pattern based active flows", Optional: true}, "device_mask": schema.StringAttribute{MarkdownDescription: "deviceIpMask", Optional: true}, "gsgroup": schema.StringAttribute{MarkdownDescription: "alias of gsgroup", Computed: true}, "session_ids": schema.Int64Attribute{Computed: true}, "ssl30_session": schema.Int64Attribute{Computed: true}, "tickets": schema.Int64Attribute{Computed: true}, "tls10_session": schema.Int64Attribute{Computed: true}, "tls11_session": schema.Int64Attribute{Computed: true}, "tls12_session": schema.Int64Attribute{Computed: true}, "total_session": schema.Int64Attribute{Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetSslDecryptionSummaryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetSslDecryptionSummaryDataSourceModel
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
func (d *GetSslDecryptionSummaryDataSource) readRemote(ctx context.Context, config *GetSslDecryptionSummaryDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/gsGroups/{alias}/flowOpsReport/sslDecryption/summary"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_summary", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	if !config.DeviceIp.IsNull() {
		query.Set("deviceIp", config.DeviceIp.ValueString())
	}
	if !config.DeviceMask.IsNull() {
		query.Set("deviceMask", config.DeviceMask.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_summary", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_summary", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_summary", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_summary", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_summary", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_summary", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_summary", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_summary", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_summary", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_summary", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_summary", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_summary", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["sslDecryptionReportSummary"]; ok {
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_summary", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetSslDecryptionSummaryDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
