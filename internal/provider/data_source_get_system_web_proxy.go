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
	_ datasource.DataSource              = (*GetSystemWebProxyDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetSystemWebProxyDataSource)(nil)
)

// GetSystemWebProxyDataSource is the generated Terraform data source implementation.
type GetSystemWebProxyDataSource struct {
	client *client.Client
}

// GetSystemWebProxyDataSourceModel describes the data source state shape.
type GetSystemWebProxyDataSourceModel struct {
	AuthType     types.String `tfsdk:"auth_type" json:"authType"`
	ClusterId    types.String `tfsdk:"cluster_id" json:"clusterId"`
	ClusterName  types.String `tfsdk:"cluster_name" json:"clusterName"`
	Password     types.String `tfsdk:"password"`
	ProxyAddress types.String `tfsdk:"proxy_address" json:"proxyAddress"`
	ProxyPort    types.Int64  `tfsdk:"proxy_port" json:"proxyPort"`
	Username     types.String `tfsdk:"username"`
}

// NewGetSystemWebProxyDataSource returns a new instance of the generated data source.
func NewGetSystemWebProxyDataSource() datasource.DataSource {
	return &GetSystemWebProxyDataSource{}
}

// Metadata returns the data source type name.
func (d *GetSystemWebProxyDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_system_web_proxy"
}

// Schema returns the data source schema.
func (d *GetSystemWebProxyDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "get system web proxy", Attributes: map[string]schema.Attribute{"auth_type": schema.StringAttribute{MarkdownDescription: "Auth type", Computed: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster Id", Optional: true}, "cluster_name": schema.StringAttribute{MarkdownDescription: "cluster name", Computed: true}, "password": schema.StringAttribute{MarkdownDescription: "Proxy Password", Computed: true, Sensitive: true}, "proxy_address": schema.StringAttribute{MarkdownDescription: "Proxy address", Computed: true}, "proxy_port": schema.Int64Attribute{MarkdownDescription: "Proxy port", Computed: true}, "username": schema.StringAttribute{MarkdownDescription: "Proxy Username", Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetSystemWebProxyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetSystemWebProxyDataSourceModel
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
func (d *GetSystemWebProxyDataSource) readRemote(ctx context.Context, config *GetSystemWebProxyDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/system/web/proxy"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_system_web_proxy", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.ClusterId.IsNull() {
		query.Set("clusterId", config.ClusterId.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_system_web_proxy", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_system_web_proxy", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_system_web_proxy", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_system_web_proxy", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_system_web_proxy", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_system_web_proxy", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_system_web_proxy", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_system_web_proxy", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_system_web_proxy", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_system_web_proxy", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_system_web_proxy", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_system_web_proxy", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_system_web_proxy", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetSystemWebProxyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
