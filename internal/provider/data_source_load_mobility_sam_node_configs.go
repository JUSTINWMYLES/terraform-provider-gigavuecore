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
	_ datasource.DataSource              = (*LoadMobilitySamNodeConfigsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadMobilitySamNodeConfigsDataSource)(nil)
)

// LoadMobilitySamNodeConfigsDataSource is the generated Terraform data source implementation.
type LoadMobilitySamNodeConfigsDataSource struct {
	client *client.Client
}

// LoadMobilitySamNodeConfigsDataSourceModel describes the data source state shape.
type LoadMobilitySamNodeConfigsDataSourceModel struct {
	Alias                   types.String  `tfsdk:"alias"`
	EngineLevelAppVzConfigs types.Dynamic `tfsdk:"engine_level_app_vz_configs" json:"engineLevelAppVzConfigs"`
	IpInterfaceConfig       types.Object  `tfsdk:"ip_interface_config" json:"ipInterfaceConfig"`
	Location                types.Object  `tfsdk:"location"`
	SamnodeAlias            types.String  `tfsdk:"samnode_alias" json:"samnodeAlias"`
	SiteName                types.String  `tfsdk:"site_name" json:"siteName"`
	SolutionAlias           types.String  `tfsdk:"solution_alias" json:"solutionAlias"`
}

// NewLoadMobilitySamNodeConfigsDataSource returns a new instance of the generated data source.
func NewLoadMobilitySamNodeConfigsDataSource() datasource.DataSource {
	return &LoadMobilitySamNodeConfigsDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadMobilitySamNodeConfigsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_mobility_sam_node_configs"
}

// Schema returns the data source schema.
func (d *LoadMobilitySamNodeConfigsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load mobility sam node configurations by alias", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Alias of the Sam node", Computed: true}, "engine_level_app_vz_configs": schema.DynamicAttribute{Computed: true}, "ip_interface_config": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "ip interface name", Computed: true}, "attach": schema.ListAttribute{MarkdownDescription: "network ports ,tool ports or circuit ports", Computed: true, ElementType: types.StringType}, "comment": schema.StringAttribute{Computed: true}, "gateway": schema.StringAttribute{MarkdownDescription: "gateway ipv4 or ipv6 address", Computed: true}, "gs_groups": schema.ListAttribute{MarkdownDescription: "Gs Groups associated with the IP Interface", Computed: true, ElementType: types.StringType}, "hw_address": schema.StringAttribute{Computed: true}, "ip_address": schema.StringAttribute{MarkdownDescription: "ipv4/ipv6 address", Computed: true}, "ip_mask": schema.StringAttribute{MarkdownDescription: "ipAddress netmask required with ipAddress", Computed: true}, "ip_type": schema.StringAttribute{Computed: true}, "mtu": schema.Int64Attribute{Computed: true}, "netflow_exporters": schema.ListAttribute{MarkdownDescription: "Netflow Exporters associated with the IP Interface", Computed: true, ElementType: types.StringType}, "tags": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"tag_key": schema.StringAttribute{MarkdownDescription: "Name of the tag", Computed: true}, "tag_values": schema.ListAttribute{MarkdownDescription: "All possible values of the tag", Computed: true, ElementType: types.StringType}}}}}}, "location": schema.SingleNestedAttribute{MarkdownDescription: "Location of the engine port", Computed: true, Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{Computed: true}, "engine_ports": schema.ListAttribute{Computed: true, ElementType: types.StringType}}}, "samnode_alias": schema.StringAttribute{MarkdownDescription: "Alias of the sam node", Required: true}, "site_name": schema.StringAttribute{MarkdownDescription: "Name of the site where the node is deployed", Computed: true}, "solution_alias": schema.StringAttribute{MarkdownDescription: "Alias of the configured solution", Required: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadMobilitySamNodeConfigsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadMobilitySamNodeConfigsDataSourceModel
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
func (d *LoadMobilitySamNodeConfigsDataSource) readRemote(ctx context.Context, config *LoadMobilitySamNodeConfigsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/intent/mobility/{solutionAlias}/samnode/{samnodeAlias}?configObject=True"
	reqPath = strings.ReplaceAll(reqPath, "{solutionAlias}", url.PathEscape(config.SolutionAlias.ValueString()))
	reqPath = strings.ReplaceAll(reqPath, "{samnodeAlias}", url.PathEscape(config.SamnodeAlias.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_sam_node_configs", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_sam_node_configs", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_sam_node_configs", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_sam_node_configs", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_sam_node_configs", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_sam_node_configs", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_sam_node_configs", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_sam_node_configs", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_sam_node_configs", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_sam_node_configs", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_sam_node_configs", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_sam_node_configs", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_sam_node_configs", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_sam_node_configs", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadMobilitySamNodeConfigsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
