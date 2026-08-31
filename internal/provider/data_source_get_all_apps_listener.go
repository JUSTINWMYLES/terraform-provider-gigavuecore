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
	_ datasource.DataSource              = (*GetAllAppsListenerDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAllAppsListenerDataSource)(nil)
)

// GetAllAppsListenerDataSource is the generated Terraform data source implementation.
type GetAllAppsListenerDataSource struct {
	client *client.Client
}

// GetAllAppsListenerDataSourceModel describes the data source state shape.
type GetAllAppsListenerDataSourceModel struct {
	ClusterId types.String `tfsdk:"cluster_id" json:"clusterId"`
	Items     types.List   `tfsdk:"items"`
}

// NewGetAllAppsListenerDataSource returns a new instance of the generated data source.
func NewGetAllAppsListenerDataSource() datasource.DataSource {
	return &GetAllAppsListenerDataSource{}
}

// Metadata returns the data source type name.
func (d *GetAllAppsListenerDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_apps_listener"
}

// Schema returns the data source schema.
func (d *GetAllAppsListenerDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get all Apps Listener", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Alias of the listener", Computed: true}, "description": schema.StringAttribute{MarkdownDescription: "Comments if necessary", Computed: true}, "gs_group_associated": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "ip_interface": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "l3": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"dscp": schema.Int64Attribute{MarkdownDescription: "DSCP Value to use", Computed: true}, "protocol": schema.StringAttribute{MarkdownDescription: "Protocol used - ipv4 , ipv6 or both", Computed: true}, "ttl": schema.Int64Attribute{MarkdownDescription: "TTL Value to use", Computed: true}}}, "l4": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"port_list": schema.ListAttribute{MarkdownDescription: "Port lists to listen on", Computed: true, ElementType: types.Int64Type}, "protocol": schema.StringAttribute{MarkdownDescription: "Protocol used - TCP or UDP", Computed: true}}}, "mode": schema.StringAttribute{MarkdownDescription: "Listen to IP interface or promiscuous mode", Computed: true}, "ssl_profile": schema.StringAttribute{MarkdownDescription: "SSL profile", Computed: true}, "status": schema.StringAttribute{Computed: true}, "tags": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"tag_key": schema.StringAttribute{MarkdownDescription: "Name of the tag", Computed: true}, "tag_values": schema.ListAttribute{MarkdownDescription: "All possible values of the tag", Computed: true, ElementType: types.StringType}}}}, "tcp_profile": schema.StringAttribute{MarkdownDescription: "TCP profile", Computed: true}, "type": schema.StringAttribute{MarkdownDescription: "Type of Apps that listen", Computed: true}}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetAllAppsListenerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAllAppsListenerDataSourceModel
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
func (d *GetAllAppsListenerDataSource) readListRemote(ctx context.Context, config *GetAllAppsListenerDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/listener"
	params := url.Values{}
	params.Set("clusterId", config.ClusterId.ValueString())
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_apps_listener", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_apps_listener", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["appsListeners"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_apps_listener", fmt.Sprintf("Could not decode list page: missing %q array", "appsListeners"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_apps_listener", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetAllAppsListenerDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
