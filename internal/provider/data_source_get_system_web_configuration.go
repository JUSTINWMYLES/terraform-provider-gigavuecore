package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*GetSystemWebConfigurationDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetSystemWebConfigurationDataSource)(nil)
)

// GetSystemWebConfigurationDataSource is the generated Terraform data source implementation.
type GetSystemWebConfigurationDataSource struct {
	client *client.Client
}

// GetSystemWebConfigurationDataSourceModel describes the data source state shape.
type GetSystemWebConfigurationDataSourceModel struct {
	AutoLogoutTimeout   types.Int64  `tfsdk:"auto_logout_timeout" json:"autoLogoutTimeout"`
	ClusterId           types.String `tfsdk:"cluster_id" json:"clusterId"`
	ClusterName         types.String `tfsdk:"cluster_name" json:"clusterName"`
	EnableHttps         types.Bool   `tfsdk:"enable_https" json:"enableHttps"`
	HttpPort            types.Int64  `tfsdk:"http_port" json:"httpPort"`
	HttpsPort           types.Int64  `tfsdk:"https_port" json:"httpsPort"`
	ServerSslMinVersion types.String `tfsdk:"server_ssl_min_version" json:"serverSslMinVersion"`
	SessionRenewal      types.Int64  `tfsdk:"session_renewal" json:"sessionRenewal"`
	SessionTimeout      types.Int64  `tfsdk:"session_timeout" json:"sessionTimeout"`
}

// NewGetSystemWebConfigurationDataSource returns a new instance of the generated data source.
func NewGetSystemWebConfigurationDataSource() datasource.DataSource {
	return &GetSystemWebConfigurationDataSource{}
}

// Metadata returns the data source type name.
func (d *GetSystemWebConfigurationDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_system_web_configuration"
}

// Schema returns the data source schema.
func (d *GetSystemWebConfigurationDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "get system web configuration", Attributes: map[string]schema.Attribute{"auto_logout_timeout": schema.Int64Attribute{MarkdownDescription: "Auto Logout Timeout in seconds", Computed: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster Id", Optional: true}, "cluster_name": schema.StringAttribute{MarkdownDescription: "cluster name", Computed: true}, "enable_https": schema.BoolAttribute{MarkdownDescription: "enableHttps ", Computed: true}, "http_port": schema.Int64Attribute{MarkdownDescription: "Http Port Number", Computed: true}, "https_port": schema.Int64Attribute{MarkdownDescription: "Https Port Number", Computed: true}, "server_ssl_min_version": schema.StringAttribute{MarkdownDescription: "SSL/TLS minimum version supported", Computed: true}, "session_renewal": schema.Int64Attribute{MarkdownDescription: "Session renewal in seconds", Computed: true}, "session_timeout": schema.Int64Attribute{MarkdownDescription: "Session renewal in seconds", Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetSystemWebConfigurationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetSystemWebConfigurationDataSourceModel
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
func (d *GetSystemWebConfigurationDataSource) readRemote(ctx context.Context, config *GetSystemWebConfigurationDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/system/web/config"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_system_web_configuration", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.ClusterId.IsNull() {
		query.Set("clusterId", config.ClusterId.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_system_web_configuration", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_system_web_configuration", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_system_web_configuration", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_system_web_configuration", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_system_web_configuration", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_system_web_configuration", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_system_web_configuration", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_system_web_configuration", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_system_web_configuration", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_system_web_configuration", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_system_web_configuration", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_system_web_configuration", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_system_web_configuration", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetSystemWebConfigurationDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
