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
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
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
	Items types.List `tfsdk:"items"`
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
	resp.Schema = schema.Schema{MarkdownDescription: "Get all Apps Proxy Server Profile", Attributes: map[string]schema.Attribute{"items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{Computed: true}, "auth_type": schema.StringAttribute{Computed: true}, "comment": schema.StringAttribute{Computed: true}, "password": schema.StringAttribute{Computed: true, Sensitive: true}, "periodic_ping": schema.StringAttribute{Computed: true}, "periodic_ping_failure_retry": schema.Int64Attribute{Computed: true}, "periodic_ping_interval": schema.Int64Attribute{Computed: true}, "periodic_ping_type": schema.StringAttribute{Computed: true}, "port": schema.Int64Attribute{Computed: true}, "protocol": schema.StringAttribute{Computed: true}, "proxy_address": schema.StringAttribute{Computed: true}, "ssl_apps": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"cluster_name": schema.ListAttribute{MarkdownDescription: "Cluster Name where the proxy deployed", Computed: true, ElementType: types.StringType}}}, "username": schema.StringAttribute{Computed: true}}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetAllProxyServerProfilesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAllProxyServerProfilesDataSourceModel
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
func (d *GetAllProxyServerProfilesDataSource) readListRemote(ctx context.Context, config *GetAllProxyServerProfilesDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/proxyServer/profiles"
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_proxy_server_profiles", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_proxy_server_profiles", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["appsProxyServerProfiles"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_proxy_server_profiles", fmt.Sprintf("Could not decode list page: missing %q array", "appsProxyServerProfiles"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
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
