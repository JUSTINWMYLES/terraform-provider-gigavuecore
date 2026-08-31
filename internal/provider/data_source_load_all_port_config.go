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
	_ datasource.DataSource              = (*LoadAllPortConfigDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadAllPortConfigDataSource)(nil)
)

// LoadAllPortConfigDataSource is the generated Terraform data source implementation.
type LoadAllPortConfigDataSource struct {
	client *client.Client
}

// LoadAllPortConfigDataSourceModel describes the data source state shape.
type LoadAllPortConfigDataSourceModel struct {
	ClusterId types.String `tfsdk:"cluster_id" json:"clusterId"`
	Items     types.List   `tfsdk:"items"`
	Page      types.String `tfsdk:"page"`
	Sort      types.String `tfsdk:"sort"`
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
	resp.Schema = schema.Schema{MarkdownDescription: "Load all PortConfigs", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"access_roles": schema.SingleNestedAttribute{MarkdownDescription: "Port Access RBAC definitions", Computed: true, Attributes: map[string]schema.Attribute{"level1": schema.ListAttribute{MarkdownDescription: "list of roles with 'Level 1' access privileges. Level 1 - Read-only access. Can view port configuration and statistics", Computed: true, ElementType: types.StringType}, "level2": schema.ListAttribute{MarkdownDescription: "list of roles with 'Level 2' access privileges. Level 2 - Level 1 plus the capability to configure port-lock, lock-share, and all traffic objects except port-pair", Computed: true, ElementType: types.StringType}, "level3": schema.ListAttribute{MarkdownDescription: "list of roles with 'Level 3' access privileges. Level 3 - Level 2 plus the capability to configure port params (such as administrative status of the port, speed, duplex, and auto-negotiation), as well as port-pair", Computed: true, ElementType: types.StringType}, "level4": schema.ListAttribute{MarkdownDescription: "list of roles with 'Level 4' access privileges. Level 4 - Level 3 plus the capability to change the port type", Computed: true, ElementType: types.StringType}}}, "alarm_thresholds": schema.SingleNestedAttribute{MarkdownDescription: "Port Alarm Thresholds definitions", Computed: true, Attributes: map[string]schema.Attribute{"alarm_buffer_threshold_rx": schema.Int64Attribute{MarkdownDescription: "in percent's. Alarm is reported every time threshold is crossed. Value of 0 disables the threshold", Computed: true}, "alarm_buffer_threshold_tx": schema.Int64Attribute{MarkdownDescription: "in percent's. Alarm is reported every time threshold is crossed. Value of 0 disables the threshold", Computed: true}, "alarm_threshold": schema.Int64Attribute{MarkdownDescription: "in percent's. Alarm is reported every time threshold is crossed. The threshold must be exceeded for at least six consecutive seconds. Value of 0 disables the threshold", Computed: true}, "alarm_threshold_low": schema.Int64Attribute{MarkdownDescription: "in percents. Alarm is reported every time this low threshold is crossed. It must be under this threshold for at least six consecutive seconds. Value of 0 disables the threshold", Computed: true}}}, "fec": schema.StringAttribute{MarkdownDescription: "enable/disable forward error correction", Computed: true}, "gdp": schema.BoolAttribute{MarkdownDescription: "enable/disable GDP packets on port", Computed: true}, "header_strip": schema.StringAttribute{MarkdownDescription: "protocol type", Computed: true}, "ingress_vlan_tag": schema.Int64Attribute{MarkdownDescription: "Ingress port VLAN tag, with valid numbers between 2 and 4000 that is added to a packet. Setting the value to 0 disables VLAN tagging. The port must be a network port", Computed: true}, "l2_gre_id": schema.Int64Attribute{MarkdownDescription: "Value of 0 disables the l2greId", Computed: true}, "licensed": schema.BoolAttribute{MarkdownDescription: "for TA series indicates whether a port is licensed. Defaults to 'true'", Computed: true}, "lock": schema.SingleNestedAttribute{MarkdownDescription: "Port Locking definitions", Computed: true, Attributes: map[string]schema.Attribute{"description": schema.StringAttribute{MarkdownDescription: "Optional lock description string", Computed: true}, "locking_user": schema.StringAttribute{MarkdownDescription: "User account name this lock belongs to. Lock is used to restrict use of the port by only the specified user account. Users with Default/Operator privileges can only lock ports to which their account has been granted access under their own user account name only. Administrators can lock any port in the system for any existing user account name", Computed: true}, "shared_with": schema.ListAttribute{MarkdownDescription: "Used to share a locked port with other user accounts. Users with Default/Operator privileges can only share a lock if the port is locked under their account name. Administrators can lock a port for another user account", Computed: true, ElementType: types.StringType}}}, "mpls_advanced": schema.SingleNestedAttribute{MarkdownDescription: "Select a combination of Mpls-Advanced options", Computed: true, Attributes: map[string]schema.Attribute{"mpls_advanced_opt": schema.ListNestedAttribute{MarkdownDescription: "List of Mpls-Advanced Options", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"adv_opt": schema.StringAttribute{MarkdownDescription: "Mpls-Advanced Options", Computed: true}}}}}}, "neighbor_discovery": schema.StringAttribute{MarkdownDescription: "Configures port neighbor discovery options", Computed: true}, "port_id": schema.StringAttribute{MarkdownDescription: "device port id", Computed: true}, "ptp": schema.SingleNestedAttribute{MarkdownDescription: "Port PTP configurations", Computed: true, Attributes: map[string]schema.Attribute{"announce_interval": schema.Int64Attribute{MarkdownDescription: "Configures the interval between PTP announce messages on an interface.\nThe range for the PTP announcement interval is from -2 to 4 log seconds.\nFor the domain 0 as well as domain 24 to 43 the interval limit will be of -3 to 4", Computed: true}, "delay_request_interval": schema.Int64Attribute{MarkdownDescription: "Configures the minimum interval allowed between PTP delay messages when the port is in the source state.\nThe range is from log(-7) to log(5) seconds, where log(-1) = 1 frame per second.\nFor the domain 0 the interval limit is -4 to 5 and for the range of domain 24 to 43 the interval limit will be of -7 to 4.", Computed: true}, "enable": schema.BoolAttribute{MarkdownDescription: "Enable PTP on the port", Computed: true}, "local_priority": schema.Int64Attribute{Computed: true}, "role": schema.StringAttribute{MarkdownDescription: "deprecated: use roleAlias", Computed: true}, "role_alias": schema.StringAttribute{Computed: true}, "sync_interval": schema.Int64Attribute{MarkdownDescription: "Configures the interval between PTP synchronization messages on an interface.\nThe range is from log(-7) to log(1) seconds.\nFor domain 0 the interval limit is -4 to 1 and for the range of domain 24 to 43 the interval limit will be of -7 to 4. ", Computed: true}, "timestamp": schema.SingleNestedAttribute{MarkdownDescription: "Port PTP timestamp configurations", Computed: true, Attributes: map[string]schema.Attribute{"egress": schema.SingleNestedAttribute{MarkdownDescription: "Tx (Egress) settings", Computed: true, Attributes: map[string]schema.Attribute{"insert": schema.BoolAttribute{MarkdownDescription: "When set, will insert timestamp for packets on the port", Computed: true}, "source_id": schema.Int64Attribute{MarkdownDescription: "Source identifier of Packet Time Stamp record header", Computed: true}}}, "ingress": schema.SingleNestedAttribute{MarkdownDescription: "Rx (Ingress) settings", Computed: true, Attributes: map[string]schema.Attribute{"insert": schema.BoolAttribute{MarkdownDescription: "When set, will insert timestamp for packets on the port", Computed: true}, "source_id": schema.Int64Attribute{MarkdownDescription: "Source identifier of Packet Time Stamp record header", Computed: true}}}}}, "vlan": schema.Int64Attribute{MarkdownDescription: "configures vlan on the PTP configured port", Computed: true}}}, "share": schema.SingleNestedAttribute{MarkdownDescription: "Port Sharing definitions", Computed: true, Attributes: map[string]schema.Attribute{"tool_share_roles": schema.ListAttribute{MarkdownDescription: "Only applicable to tool ports. Used to designate a tool port as available for tool-to-tool pass-alls (tool-mirrors) with the specified roles", Computed: true, ElementType: types.StringType}}}, "tag_protocol_id": schema.StringAttribute{MarkdownDescription: "When Ingress port VLAN tag is added , this protocol Id will be added which egress out the traffic of tool port ", Computed: true}, "taptx": schema.StringAttribute{MarkdownDescription: "Opens or closes the copper tap port relay. Active closes the port relay. Passive opens the port relay", Computed: true}, "timestamp": schema.SingleNestedAttribute{MarkdownDescription: "Timestamping definitions for GigaPORT-X12-TS ports", Computed: true, Attributes: map[string]schema.Attribute{"append_ingress": schema.BoolAttribute{MarkdownDescription: "Used to add a timestamp to ingress packets. Applicable to GigaPORT-X12-TS card ports x1..x12 when configured as network ports", Computed: true}, "source_id": schema.Int64Attribute{MarkdownDescription: "Used to specify a custom source-id to be included in the timestamp appended by the GigaPORT-X12-TS. Applicable to GigaPORT-X12-TS card ports x1..x12 when configured as network ports. The timestamp always includes a source-id field. If this custom value is not explicitly specified, the GigaPORT-X12-TS generates one automatically using the following formula: (<box-id> * 2048) + (<slot-id> * 256) + <port-number>", Computed: true}, "strip_egress": schema.BoolAttribute{MarkdownDescription: "Used to strip timestamps from egress packets. Use this argument to strip timestamps from egress packets. Important: this option should only be enabled to packets with time stamps appended. This function will strip the last 14 bytes of each packet regardless of whether a timestamp has been added. Applicable to GigaPORT-X12-TS card ports x9..x12 when configured as tool ports", Computed: true}}}, "vxlan_id": schema.Int64Attribute{MarkdownDescription: "Value of 0 disables the vxlanId", Computed: true}}}}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadAllPortConfigDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadAllPortConfigDataSourceModel
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
func (d *LoadAllPortConfigDataSource) readListRemote(ctx context.Context, config *LoadAllPortConfigDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/portConfig/portConfigs"
	params := url.Values{}
	params.Set("clusterId", config.ClusterId.ValueString())
	if !config.Page.IsNull() {
		params.Set("page", config.Page.ValueString())
	}
	if !config.Sort.IsNull() {
		params.Set("sort", config.Sort.ValueString())
	}
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
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_port_config", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_port_config", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["portConfigs"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_port_config", fmt.Sprintf("Could not decode list page: missing %q array", "portConfigs"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
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
