package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
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
	_ datasource.DataSource              = (*LoadManagementInterfaceNeighborsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadManagementInterfaceNeighborsDataSource)(nil)
)

// LoadManagementInterfaceNeighborsDataSource is the generated Terraform data source implementation.
type LoadManagementInterfaceNeighborsDataSource struct {
	client *client.Client
}

// LoadManagementInterfaceNeighborsDataSourceModel describes the data source state shape.
type LoadManagementInterfaceNeighborsDataSourceModel struct {
	BoxId         types.Int64  `tfsdk:"box_id" json:"boxId"`
	CdpNeighbors  types.List   `tfsdk:"cdp_neighbors" json:"cdpNeighbors"`
	ClusterId     types.String `tfsdk:"cluster_id" json:"clusterId"`
	InterfaceName types.String `tfsdk:"interface_name" json:"interfaceName"`
	LldpNeighbors types.List   `tfsdk:"lldp_neighbors" json:"lldpNeighbors"`
}

// NewLoadManagementInterfaceNeighborsDataSource returns a new instance of the generated data source.
func NewLoadManagementInterfaceNeighborsDataSource() datasource.DataSource {
	return &LoadManagementInterfaceNeighborsDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadManagementInterfaceNeighborsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_management_interface_neighbors"
}

// Schema returns the data source schema.
func (d *LoadManagementInterfaceNeighborsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "since FM 5.8", Attributes: map[string]schema.Attribute{"box_id": schema.Int64Attribute{MarkdownDescription: "boxId", Required: true}, "cdp_neighbors": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"device_id": schema.StringAttribute{Computed: true}, "iface_addr": schema.StringAttribute{MarkdownDescription: "ipv4", Computed: true}, "last_update": schema.Int64Attribute{MarkdownDescription: "Timestamp of the last update. In UTC milliseconds", Computed: true}, "mgmt_addr": schema.StringAttribute{MarkdownDescription: "ipv4", Computed: true}, "net_prefix_addr": schema.StringAttribute{MarkdownDescription: "ipv4", Computed: true}, "net_prefix_mask": schema.StringAttribute{MarkdownDescription: "ipv4", Computed: true}, "platform": schema.StringAttribute{Computed: true}, "port_id": schema.StringAttribute{Computed: true}, "port_vlan_id": schema.Int64Attribute{Computed: true}, "sw_version": schema.StringAttribute{Computed: true}, "sys_cap_available": schema.Int64Attribute{Computed: true}, "ttl": schema.Int64Attribute{Computed: true}}}}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "interface_name": schema.StringAttribute{MarkdownDescription: "Management Interface Name", Required: true}, "lldp_neighbors": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"chassis_id": schema.StringAttribute{Computed: true}, "last_update": schema.Int64Attribute{MarkdownDescription: "Timestamp of the last update. In UTC milliseconds", Computed: true}, "link_agg_port_id": schema.Int64Attribute{Computed: true}, "link_agg_status": schema.Int64Attribute{Computed: true}, "mgmt_addr": schema.StringAttribute{MarkdownDescription: "ipv4 or ipv6", Computed: true}, "mgmt_vlan_id": schema.Int64Attribute{Computed: true}, "mtu": schema.Int64Attribute{Computed: true}, "port_descr": schema.StringAttribute{Computed: true}, "port_id": schema.StringAttribute{Computed: true}, "port_vlan_id": schema.Int64Attribute{Computed: true}, "sys_cap_available": schema.Int64Attribute{Computed: true}, "sys_cap_enabled": schema.Int64Attribute{Computed: true}, "sys_descr": schema.StringAttribute{Computed: true}, "sys_name": schema.StringAttribute{Computed: true}, "ttl": schema.Int64Attribute{Computed: true}, "vlan_name": schema.StringAttribute{Computed: true}}}}}}
}

// Read fetches remote state into the data source model.
func (d *LoadManagementInterfaceNeighborsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadManagementInterfaceNeighborsDataSourceModel
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
func (d *LoadManagementInterfaceNeighborsDataSource) readRemote(ctx context.Context, config *LoadManagementInterfaceNeighborsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/system/interfaces/mgmt/neighbors/{boxId}/{interfaceName}"
	reqPath = strings.ReplaceAll(reqPath, "{boxId}", url.PathEscape(strconv.FormatInt(config.BoxId.ValueInt64(), 10)))
	reqPath = strings.ReplaceAll(reqPath, "{interfaceName}", url.PathEscape(config.InterfaceName.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_management_interface_neighbors", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_management_interface_neighbors", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_management_interface_neighbors", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_management_interface_neighbors", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_management_interface_neighbors", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_management_interface_neighbors", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_management_interface_neighbors", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_management_interface_neighbors", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_management_interface_neighbors", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_management_interface_neighbors", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_load_management_interface_neighbors", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_load_management_interface_neighbors", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_management_interface_neighbors", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["mgmtInterfaceNeighbors"]; ok {
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
		resp.Diagnostics.AddError("Error reading gigavuecore_load_management_interface_neighbors", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadManagementInterfaceNeighborsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
