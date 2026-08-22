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
	_ datasource.DataSource              = (*GetAllAppsListenerDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAllAppsListenerDataSource)(nil)
)

// GetAllAppsListenerDataSource is the generated Terraform data source implementation.
type GetAllAppsListenerDataSource struct {
	client *client.Client
}

// GetAllAppsListenerDataSourceModel describes the data source state shape.
type GetAllAppsListenerDataSourceModel struct {
	AppsListeners types.List   `tfsdk:"apps_listeners" json:"appsListeners"`
	ClusterId     types.String `tfsdk:"cluster_id" json:"clusterId"`
	Context       types.Object `tfsdk:"context"`
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
	resp.Schema = schema.Schema{MarkdownDescription: "Get all Apps Listener", Attributes: map[string]schema.Attribute{"apps_listeners": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Alias of the listener", Computed: true}, "description": schema.StringAttribute{MarkdownDescription: "Comments if necessary", Computed: true}, "gs_group_associated": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "ip_interface": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "l3": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"dscp": schema.Int64Attribute{MarkdownDescription: "DSCP Value to use", Computed: true}, "protocol": schema.StringAttribute{MarkdownDescription: "Protocol used - ipv4 , ipv6 or both", Computed: true}, "ttl": schema.Int64Attribute{MarkdownDescription: "TTL Value to use", Computed: true}}}, "l4": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"port_list": schema.ListAttribute{MarkdownDescription: "Port lists to listen on", Computed: true, ElementType: types.Int64Type}, "protocol": schema.StringAttribute{MarkdownDescription: "Protocol used - TCP or UDP", Computed: true}}}, "mode": schema.StringAttribute{MarkdownDescription: "Listen to IP interface or promiscuous mode", Computed: true}, "ssl_profile": schema.StringAttribute{MarkdownDescription: "SSL profile", Computed: true}, "status": schema.StringAttribute{Computed: true}, "tags": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"tag_key": schema.StringAttribute{MarkdownDescription: "Name of the tag", Computed: true}, "tag_values": schema.ListAttribute{MarkdownDescription: "All possible values of the tag", Computed: true, ElementType: types.StringType}}}}, "tcp_profile": schema.StringAttribute{MarkdownDescription: "TCP profile", Computed: true}, "type": schema.StringAttribute{MarkdownDescription: "Type of Apps that listen", Computed: true}}}}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "context": schema.SingleNestedAttribute{MarkdownDescription: "Gigamon query result context", Computed: true, Attributes: map[string]schema.Attribute{"page_no": schema.Int64Attribute{MarkdownDescription: "page number of the returned result set", Computed: true}, "page_size": schema.Int64Attribute{MarkdownDescription: "page size of the returned result set", Computed: true}, "sort": schema.ListAttribute{MarkdownDescription: "sorting info of the returned result set. list of fields in the array indicate sorting order", Computed: true, ElementType: types.StringType}, "total_items": schema.Int64Attribute{MarkdownDescription: "total number of items in the queried entity type", Computed: true}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetAllAppsListenerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAllAppsListenerDataSourceModel
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
func (d *GetAllAppsListenerDataSource) readRemote(ctx context.Context, config *GetAllAppsListenerDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/listener"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_apps_listener", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_apps_listener", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_apps_listener", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_apps_listener", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_apps_listener", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_apps_listener", "Entity Not Found. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_all_apps_listener", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_apps_listener", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_apps_listener", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
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
