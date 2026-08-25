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
	_ datasource.DataSource              = (*GetAllProxyServerProfilesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAllProxyServerProfilesDataSource)(nil)
)

// GetAllProxyServerProfilesDataSource is the generated Terraform data source implementation.
type GetAllProxyServerProfilesDataSource struct {
	client *client.Client
}

// GetAllProxyServerProfilesDataSourceModel describes the data source state shape.
type GetAllProxyServerProfilesDataSourceModel struct {
	AppsProxyServerProfiles types.List   `tfsdk:"apps_proxy_server_profiles" json:"appsProxyServerProfiles"`
	Context                 types.Object `tfsdk:"context"`
}

// NewGetAllProxyServerProfilesDataSource returns a new instance of the generated data source.
func NewGetAllProxyServerProfilesDataSource() datasource.DataSource {
	return &GetAllProxyServerProfilesDataSource{}
}

// Metadata returns the data source type name.
func (d *GetAllProxyServerProfilesDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_proxy_server_profiles"
}

// Schema returns the data source schema.
func (d *GetAllProxyServerProfilesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get all Apps Proxy Server Profile", Attributes: map[string]schema.Attribute{"apps_proxy_server_profiles": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{Computed: true}, "auth_type": schema.StringAttribute{Computed: true}, "comment": schema.StringAttribute{Computed: true}, "password": schema.StringAttribute{Computed: true, Sensitive: true}, "periodic_ping": schema.StringAttribute{Computed: true}, "periodic_ping_failure_retry": schema.Int64Attribute{Computed: true}, "periodic_ping_interval": schema.Int64Attribute{Computed: true}, "periodic_ping_type": schema.StringAttribute{Computed: true}, "port": schema.Int64Attribute{Computed: true}, "protocol": schema.StringAttribute{Computed: true}, "proxy_address": schema.StringAttribute{Computed: true}, "ssl_apps": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"cluster_name": schema.ListAttribute{MarkdownDescription: "Cluster Name where the proxy deployed", Computed: true, ElementType: types.StringType}}}, "username": schema.StringAttribute{Computed: true}}}}, "context": schema.SingleNestedAttribute{MarkdownDescription: "Gigamon query result context", Computed: true, Attributes: map[string]schema.Attribute{"page_no": schema.Int64Attribute{MarkdownDescription: "page number of the returned result set", Computed: true}, "page_size": schema.Int64Attribute{MarkdownDescription: "page size of the returned result set", Computed: true}, "sort": schema.ListAttribute{MarkdownDescription: "sorting info of the returned result set. list of fields in the array indicate sorting order", Computed: true, ElementType: types.StringType}, "total_items": schema.Int64Attribute{MarkdownDescription: "total number of items in the queried entity type", Computed: true}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetAllProxyServerProfilesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAllProxyServerProfilesDataSourceModel
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
func (d *GetAllProxyServerProfilesDataSource) readRemote(ctx context.Context, config *GetAllProxyServerProfilesDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/proxyServer/profiles"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_proxy_server_profiles", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_proxy_server_profiles", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_proxy_server_profiles", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_proxy_server_profiles", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_proxy_server_profiles", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_proxy_server_profiles", "Entity Not Found. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_all_proxy_server_profiles", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_proxy_server_profiles", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_proxy_server_profiles", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_proxy_server_profiles", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetAllProxyServerProfilesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
