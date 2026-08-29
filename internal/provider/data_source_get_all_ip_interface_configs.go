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
	_ datasource.DataSource              = (*GetAllIpInterfaceConfigsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAllIpInterfaceConfigsDataSource)(nil)
)

// GetAllIpInterfaceConfigsDataSource is the generated Terraform data source implementation.
type GetAllIpInterfaceConfigsDataSource struct {
	client *client.Client
}

// GetAllIpInterfaceConfigsDataSourceModel describes the data source state shape.
type GetAllIpInterfaceConfigsDataSourceModel struct {
	Application        types.String `tfsdk:"application"`
	ClusterId          types.String `tfsdk:"cluster_id" json:"clusterId"`
	IpInterfaceConfigs types.List   `tfsdk:"ip_interface_configs" json:"ipInterfaceConfigs"`
	Tags               types.List   `tfsdk:"tags"`
}

// NewGetAllIpInterfaceConfigsDataSource returns a new instance of the generated data source.
func NewGetAllIpInterfaceConfigsDataSource() datasource.DataSource {
	return &GetAllIpInterfaceConfigsDataSource{}
}

// Metadata returns the data source type name.
func (d *GetAllIpInterfaceConfigsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_ip_interface_configs"
}

// Schema returns the data source schema.
func (d *GetAllIpInterfaceConfigsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load all ipInterface solution configurations", Attributes: map[string]schema.Attribute{"application": schema.StringAttribute{MarkdownDescription: "Gets ipInterface solution that supports the specified GigaSmart application", Optional: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Gets ipInterface solution configs for the provided cluster", Optional: true}, "ip_interface_configs": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Alias of the ip interface solution", Computed: true}, "applications": schema.ListAttribute{MarkdownDescription: "GigaSMART applications for which the ipInterface is used", Computed: true, ElementType: types.StringType}, "cluster_name": schema.StringAttribute{MarkdownDescription: "clusterId", Computed: true}, "config_status": schema.StringAttribute{Computed: true}, "config_status_reasons": schema.StringAttribute{Computed: true}, "gateway": schema.StringAttribute{MarkdownDescription: "gateway ipv4 or ipv6 address", Computed: true}, "interfaces": schema.ListAttribute{MarkdownDescription: "network ports ,tool ports or circuit ports", Computed: true, ElementType: types.StringType}, "ip_address": schema.StringAttribute{MarkdownDescription: "ipv4/ipv6 address", Computed: true}, "ip_mask": schema.StringAttribute{MarkdownDescription: "ipAddress netmask required with ipAddress", Computed: true}, "managed_status": schema.StringAttribute{MarkdownDescription: "managed status of the ip interface configuration. ACTIVE if the cluster is managed in this FM , else it will be marked as INACTIVE", Computed: true}, "mtu": schema.Int64Attribute{Computed: true}, "ref_count": schema.Int64Attribute{MarkdownDescription: "number of control or user nodes using this ip interface configuration", Computed: true}}}}, "tags": schema.ListNestedAttribute{MarkdownDescription: "RBAC Tags", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"tag_key": schema.StringAttribute{MarkdownDescription: "Name of the tag", Computed: true}, "tag_values": schema.ListAttribute{MarkdownDescription: "All possible values of the tag", Computed: true, ElementType: types.StringType}}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetAllIpInterfaceConfigsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAllIpInterfaceConfigsDataSourceModel
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
func (d *GetAllIpInterfaceConfigsDataSource) readRemote(ctx context.Context, config *GetAllIpInterfaceConfigsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/ipInterfaceConfigs"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_ip_interface_configs", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.ClusterId.IsNull() {
		query.Set("clusterId", config.ClusterId.ValueString())
	}
	if !config.Application.IsNull() {
		query.Set("application", config.Application.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_ip_interface_configs", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_ip_interface_configs", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_ip_interface_configs", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_ip_interface_configs", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_ip_interface_configs", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_ip_interface_configs", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_ip_interface_configs", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_ip_interface_configs", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_ip_interface_configs", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_all_ip_interface_configs", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_ip_interface_configs", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_ip_interface_configs", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_ip_interface_configs", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetAllIpInterfaceConfigsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
