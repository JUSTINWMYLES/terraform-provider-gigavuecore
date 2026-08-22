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
	_ datasource.DataSource              = (*LoadAllPortConfigDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadAllPortConfigDataSource)(nil)
)

// LoadAllPortConfigDataSource is the generated Terraform data source implementation.
type LoadAllPortConfigDataSource struct {
	client *client.Client
}

// LoadAllPortConfigDataSourceModel describes the data source state shape.
type LoadAllPortConfigDataSourceModel struct {
	ClusterId   types.String `tfsdk:"cluster_id" json:"clusterId"`
	Context     types.Object `tfsdk:"context"`
	Page        types.String `tfsdk:"page"`
	PortConfigs types.List   `tfsdk:"port_configs" json:"portConfigs"`
	Sort        types.String `tfsdk:"sort"`
}

// NewLoadAllPortConfigDataSource returns a new instance of the generated data source.
func NewLoadAllPortConfigDataSource() datasource.DataSource {
	return &LoadAllPortConfigDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadAllPortConfigDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_all_port_config"
}

// Schema returns the data source schema.
func (d *LoadAllPortConfigDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load all PortConfigs", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "context": schema.SingleNestedAttribute{MarkdownDescription: "Gigamon query result context", Computed: true, Attributes: map[string]schema.Attribute{"page_no": schema.Int64Attribute{MarkdownDescription: "page number of the returned result set", Computed: true}, "page_size": schema.Int64Attribute{MarkdownDescription: "page size of the returned result set", Computed: true}, "sort": schema.ListAttribute{MarkdownDescription: "sorting info of the returned result set. list of fields in the array indicate sorting order", Computed: true, ElementType: types.StringType}, "total_items": schema.Int64Attribute{MarkdownDescription: "total number of items in the queried entity type", Computed: true}}}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "port_configs": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"access_roles": schema.SingleNestedAttribute{MarkdownDescription: "Port Access RBAC definitions", Computed: true, Attributes: map[string]schema.Attribute{"level1": schema.ListAttribute{MarkdownDescription: "list of roles with 'Level 1' access privileges. Level 1 - Read-only access. Can view port configuration and statistics", Computed: true, ElementType: types.StringType}, "level2": schema.ListAttribute{MarkdownDescription: "list of roles with 'Level 2' access privileges. Level 2 - Level 1 plus the capability to configure port-lock, lock-share, and all traffic objects except port-pair", Computed: true, ElementType: types.StringType}, "level3": schema.ListAttribute{MarkdownDescription: "list of roles with 'Level 3' access privileges. Level 3 - Level 2 plus the capability to configure port params (such as administrative status of the port, speed, duplex, and auto-negotiation), as well as port-pair", Computed: true, ElementType: types.StringType}, "level4": schema.ListAttribute{MarkdownDescription: "list of roles with 'Level 4' access privileges. Level 4 - Level 3 plus the capability to change the port type", Computed: true, ElementType: types.StringType}}}, "alarm_thresholds": schema.SingleNestedAttribute{MarkdownDescription: "Port Alarm Thresholds definitions", Computed: true, Attributes: map[string]schema.Attribute{"alarm_buffer_threshold_rx": schema.Int64Attribute{MarkdownDescription: "in percent's. Alarm is reported every time threshold is crossed. Value of 0 disables the threshold", Computed: true}, "alarm_buffer_threshold_tx": schema.Int64Attribute{MarkdownDescription: "in percent's. Alarm is reported every time threshold is crossed. Value of 0 disables the threshold", Computed: true}, "alarm_threshold": schema.Int64Attribute{MarkdownDescription: "in percent's. Alarm is reported every time threshold is crossed. The threshold must be exceeded for at least six consecutive seconds. Value of 0 disables the threshold", Computed: true}, "alarm_threshold_low": schema.Int64Attribute{MarkdownDescription: "in percents. Alarm is reported every time this low threshold is crossed. It must be under this threshold for at least six consecutive seconds. Value of 0 disables the threshold", Computed: true}}}, "fec": schema.StringAttribute{MarkdownDescription: "enable/disable forward error correction", Computed: true}, "gdp": schema.BoolAttribute{MarkdownDescription: "enable/disable GDP packets on port", Computed: true}, "header_strip": schema.StringAttribute{MarkdownDescription: "protocol type", Computed: true}, "ingress_vlan_tag": schema.Int64Attribute{MarkdownDescription: "Ingress port VLAN tag, with valid numbers between 2 and 4000 that is added to a packet. Setting the value to 0 disables VLAN tagging. The port must be a network port", Computed: true}, "l2_gre_id": schema.Int64Attribute{MarkdownDescription: "Value of 0 disables the l2greId", Computed: true}, "licensed": schema.BoolAttribute{MarkdownDescription: "for TA series indicates whether a port is licensed. Defaults to 'true'", Computed: true}, "lock": schema.SingleNestedAttribute{MarkdownDescription: "Port Locking definitions", Computed: true, Attributes: map[string]schema.Attribute{"description": schema.StringAttribute{MarkdownDescription: "Optional lock description string", Computed: true}, "locking_user": schema.StringAttribute{MarkdownDescription: "User account name this lock belongs to. Lock is used to restrict use of the port by only the specified user account. Users with Default/Operator privileges can only lock ports to which their account has been granted access under their own user account name only. Administrators can lock any port in the system for any existing user account name", Computed: true}, "shared_with": schema.ListAttribute{MarkdownDescription: "Used to share a locked port with other user accounts. Users with Default/Operator privileges can only share a lock if the port is locked under their account name. Administrators can lock a port for another user account", Computed: true, ElementType: types.StringType}}}, "mpls_advanced": schema.SingleNestedAttribute{MarkdownDescription: "Select a combination of Mpls-Advanced options", Computed: true, Attributes: map[string]schema.Attribute{"mpls_advanced_opt": schema.ListNestedAttribute{MarkdownDescription: "List of Mpls-Advanced Options", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"adv_opt": schema.StringAttribute{MarkdownDescription: "Mpls-Advanced Options", Computed: true}}}}}}, "neighbor_discovery": schema.StringAttribute{MarkdownDescription: "Configures port neighbor discovery options", Computed: true}, "port_id": schema.StringAttribute{MarkdownDescription: "device port id", Computed: true}, "ptp": schema.SingleNestedAttribute{MarkdownDescription: "Port PTP configurations", Computed: true, Attributes: map[string]schema.Attribute{"announce_interval": schema.Int64Attribute{MarkdownDescription: "Configures the interval between PTP announce messages on an interface.\nThe range for the PTP announcement interval is from -2 to 4 log seconds.\nFor the domain 0 as well as domain 24 to 43 the interval limit will be of -3 to 4", Computed: true}, "delay_request_interval": schema.Int64Attribute{MarkdownDescription: "Configures the minimum interval allowed between PTP delay messages when the port is in the source state.\nThe range is from log(-7) to log(5) seconds, where log(-1) = 1 frame per second.\nFor the domain 0 the interval limit is -4 to 5 and for the range of domain 24 to 43 the interval limit will be of -7 to 4.", Computed: true}, "enable": schema.BoolAttribute{MarkdownDescription: "Enable PTP on the port", Computed: true}, "local_priority": schema.Int64Attribute{Computed: true}, "role": schema.StringAttribute{MarkdownDescription: "deprecated: use roleAlias", Computed: true}, "role_alias": schema.StringAttribute{Computed: true}, "sync_interval": schema.Int64Attribute{MarkdownDescription: "Configures the interval between PTP synchronization messages on an interface.\nThe range is from log(-7) to log(1) seconds.\nFor domain 0 the interval limit is -4 to 1 and for the range of domain 24 to 43 the interval limit will be of -7 to 4. ", Computed: true}, "timestamp": schema.SingleNestedAttribute{MarkdownDescription: "Port PTP timestamp configurations", Computed: true, Attributes: map[string]schema.Attribute{"egress": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"insert": schema.BoolAttribute{MarkdownDescription: "When set, will insert timestamp for packets on the port", Computed: true}, "source_id": schema.Int64Attribute{MarkdownDescription: "Source identifier of Packet Time Stamp record header", Computed: true}}}, "ingress": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"insert": schema.BoolAttribute{MarkdownDescription: "When set, will insert timestamp for packets on the port", Computed: true}, "source_id": schema.Int64Attribute{MarkdownDescription: "Source identifier of Packet Time Stamp record header", Computed: true}}}}}, "vlan": schema.Int64Attribute{MarkdownDescription: "configures vlan on the PTP configured port", Computed: true}}}, "share": schema.SingleNestedAttribute{MarkdownDescription: "Port Sharing definitions", Computed: true, Attributes: map[string]schema.Attribute{"tool_share_roles": schema.ListAttribute{MarkdownDescription: "Only applicable to tool ports. Used to designate a tool port as available for tool-to-tool pass-alls (tool-mirrors) with the specified roles", Computed: true, ElementType: types.StringType}}}, "tag_protocol_id": schema.StringAttribute{MarkdownDescription: "When Ingress port VLAN tag is added , this protocol Id will be added which egress out the traffic of tool port ", Computed: true}, "taptx": schema.StringAttribute{MarkdownDescription: "Opens or closes the copper tap port relay. Active closes the port relay. Passive opens the port relay", Computed: true}, "timestamp": schema.SingleNestedAttribute{MarkdownDescription: "Timestamping definitions for GigaPORT-X12-TS ports", Computed: true, Attributes: map[string]schema.Attribute{"append_ingress": schema.BoolAttribute{MarkdownDescription: "Used to add a timestamp to ingress packets. Applicable to GigaPORT-X12-TS card ports x1..x12 when configured as network ports", Computed: true}, "source_id": schema.Int64Attribute{MarkdownDescription: "Used to specify a custom source-id to be included in the timestamp appended by the GigaPORT-X12-TS. Applicable to GigaPORT-X12-TS card ports x1..x12 when configured as network ports. The timestamp always includes a source-id field. If this custom value is not explicitly specified, the GigaPORT-X12-TS generates one automatically using the following formula: (<box-id> * 2048) + (<slot-id> * 256) + <port-number>", Computed: true}, "strip_egress": schema.BoolAttribute{MarkdownDescription: "Used to strip timestamps from egress packets. Use this argument to strip timestamps from egress packets. Important: this option should only be enabled to packets with time stamps appended. This function will strip the last 14 bytes of each packet regardless of whether a timestamp has been added. Applicable to GigaPORT-X12-TS card ports x9..x12 when configured as tool ports", Computed: true}}}, "vxlan_id": schema.Int64Attribute{MarkdownDescription: "Value of 0 disables the vxlanId", Computed: true}}}}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadAllPortConfigDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadAllPortConfigDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	d.readRemote(ctx, &config, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	if config.Page.IsNull() {
		resp.Diagnostics.AddWarning("Single-page result", fmt.Sprintf("This data source reads a single page of a paginated API endpoint and does not aggregate results across pages. The \"page\" argument is unset, so the default page is returned; set it to retrieve a different page."))
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}

// readRemote performs the read HTTP exchange and decodes the response into config. Extracted from Read so the request/response logic is unit-testable without a tfsdk.Config.
func (d *LoadAllPortConfigDataSource) readRemote(ctx context.Context, config *LoadAllPortConfigDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/portConfig/portConfigs"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_port_config", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	if !config.Page.IsNull() {
		query.Set("page", config.Page.ValueString())
	}
	if !config.Sort.IsNull() {
		query.Set("sort", config.Sort.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_port_config", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_port_config", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_port_config", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_port_config", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_port_config", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_port_config", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_port_config", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_port_config", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_port_config", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_load_all_port_config", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_port_config", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_port_config", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_port_config", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadAllPortConfigDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
