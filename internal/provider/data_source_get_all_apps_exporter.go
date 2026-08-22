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
	_ datasource.DataSource              = (*GetAllAppsExporterDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAllAppsExporterDataSource)(nil)
)

// GetAllAppsExporterDataSource is the generated Terraform data source implementation.
type GetAllAppsExporterDataSource struct {
	client *client.Client
}

// GetAllAppsExporterDataSourceModel describes the data source state shape.
type GetAllAppsExporterDataSourceModel struct {
	AppsExporters types.List   `tfsdk:"apps_exporters" json:"appsExporters"`
	ClusterId     types.String `tfsdk:"cluster_id" json:"clusterId"`
	Context       types.Object `tfsdk:"context"`
}

// NewGetAllAppsExporterDataSource returns a new instance of the generated data source.
func NewGetAllAppsExporterDataSource() datasource.DataSource {
	return &GetAllAppsExporterDataSource{}
}

// Metadata returns the data source type name.
func (d *GetAllAppsExporterDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_apps_exporter"
}

// Schema returns the data source schema.
func (d *GetAllAppsExporterDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get all Apps Exporter", Attributes: map[string]schema.Attribute{"apps_exporters": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Alias of the exporter", Computed: true}, "description": schema.StringAttribute{MarkdownDescription: "Comments if necessary", Computed: true}, "destination": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"l3": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"ip": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"dscp": schema.Int64Attribute{MarkdownDescription: "DSCP Value to use", Computed: true}, "ttl": schema.Int64Attribute{MarkdownDescription: "TTL Value to use", Computed: true}, "ver4": schema.StringAttribute{MarkdownDescription: "IPv4 address. If not specified, it will be determined by the App based on other traffic correlation or external means (auto mode).", Computed: true}, "ver6": schema.StringAttribute{MarkdownDescription: "IPv6 address. If not specified, it will be determined by the App based on other traffic correlation or external means (auto mode).", Computed: true}}}, "protocol": schema.StringAttribute{MarkdownDescription: "Protocol used (when it's auto, it's determined by the App based on context or by discovery)", Computed: true}}}, "l4": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"port": schema.Int64Attribute{MarkdownDescription: "Base port used to export, port is optional for type:gtp-cups", Computed: true}, "protocol": schema.StringAttribute{MarkdownDescription: "Protocol used - TCP or UDP", Computed: true}}}}}, "gs_group_associated": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "source": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"interface": schema.StringAttribute{MarkdownDescription: "Alias of IP Interface", Computed: true}, "l4_port": schema.Int64Attribute{MarkdownDescription: "Base source port number to use for outgoing connections", Computed: true}}}, "ssl_profile": schema.StringAttribute{MarkdownDescription: "SSL profile alias", Computed: true}, "status": schema.StringAttribute{Computed: true}, "tags": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"tag_key": schema.StringAttribute{MarkdownDescription: "Name of the tag", Computed: true}, "tag_values": schema.ListAttribute{MarkdownDescription: "All possible values of the tag", Computed: true, ElementType: types.StringType}}}}, "tcp_profile": schema.StringAttribute{MarkdownDescription: "TCP profile alias", Computed: true}, "type": schema.StringAttribute{MarkdownDescription: "Type of Apps that export", Computed: true}}}}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "context": schema.SingleNestedAttribute{MarkdownDescription: "Gigamon query result context", Computed: true, Attributes: map[string]schema.Attribute{"page_no": schema.Int64Attribute{MarkdownDescription: "page number of the returned result set", Computed: true}, "page_size": schema.Int64Attribute{MarkdownDescription: "page size of the returned result set", Computed: true}, "sort": schema.ListAttribute{MarkdownDescription: "sorting info of the returned result set. list of fields in the array indicate sorting order", Computed: true, ElementType: types.StringType}, "total_items": schema.Int64Attribute{MarkdownDescription: "total number of items in the queried entity type", Computed: true}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetAllAppsExporterDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAllAppsExporterDataSourceModel
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
func (d *GetAllAppsExporterDataSource) readRemote(ctx context.Context, config *GetAllAppsExporterDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/exporter"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_apps_exporter", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_apps_exporter", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_apps_exporter", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_apps_exporter", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_apps_exporter", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_apps_exporter", "Entity Not Found. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_all_apps_exporter", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_apps_exporter", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_apps_exporter", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_apps_exporter", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetAllAppsExporterDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
