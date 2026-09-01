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
	_ datasource.DataSource              = (*LoadMobilityControlNodeDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadMobilityControlNodeDataSource)(nil)
)

// LoadMobilityControlNodeDataSource is the generated Terraform data source implementation.
type LoadMobilityControlNodeDataSource struct {
	client *client.Client
}

// LoadMobilityControlNodeDataSourceModel describes the data source state shape.
type LoadMobilityControlNodeDataSourceModel struct {
	AdditionalToolPorts             types.List    `tfsdk:"additional_tool_ports" json:"additionalToolPorts"`
	Alias                           types.String  `tfsdk:"alias"`
	App5GHttp2Ports                 types.List    `tfsdk:"app5_g_http2_ports" json:"app5gHTTP2Ports"`
	AppTcp                          types.Object  `tfsdk:"app_tcp" json:"appTcp"`
	CollectorTools                  types.List    `tfsdk:"collector_tools" json:"collectorTools"`
	ConfigStatus                    types.String  `tfsdk:"config_status" json:"configStatus"`
	ConfigStatusReasons             types.String  `tfsdk:"config_status_reasons" json:"configStatusReasons"`
	ControlMetadataIpInterfaceAlias types.String  `tfsdk:"control_metadata_ip_interface_alias" json:"controlMetadataIpInterfaceAlias"`
	CpnodeAlias                     types.String  `tfsdk:"cpnode_alias" json:"cpnodeAlias"`
	Deployed                        types.Bool    `tfsdk:"deployed"`
	GtpControlSample                types.Bool    `tfsdk:"gtp_control_sample" json:"gtpControlSample"`
	GtpRandomSampling               types.Object  `tfsdk:"gtp_random_sampling" json:"gtpRandomSampling"`
	HealthState                     types.String  `tfsdk:"health_state" json:"healthState"`
	HealthStateReasons              types.List    `tfsdk:"health_state_reasons" json:"healthStateReasons"`
	IpInterfaceAlias                types.String  `tfsdk:"ip_interface_alias" json:"ipInterfaceAlias"`
	Location                        types.Object  `tfsdk:"location"`
	NodeOverrideNetworkPorts        types.List    `tfsdk:"node_override_network_ports" json:"nodeOverrideNetworkPorts"`
	NodeType                        types.String  `tfsdk:"node_type" json:"nodeType"`
	NumberOf5GSessions              types.Int64   `tfsdk:"number_of5_g_sessions" json:"numberOf5gSessions"`
	NumberOfLteSessions             types.Int64   `tfsdk:"number_of_lte_sessions" json:"numberOfLteSessions"`
	SolutionAlias                   types.String  `tfsdk:"solution_alias" json:"solutionAlias"`
	Tags                            types.List    `tfsdk:"tags"`
	TrafficSources                  types.Dynamic `tfsdk:"traffic_sources" json:"trafficSources"`
}

// NewLoadMobilityControlNodeDataSource returns a new instance of the generated data source.
func NewLoadMobilityControlNodeDataSource() datasource.DataSource {
	return &LoadMobilityControlNodeDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadMobilityControlNodeDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_mobility_control_node"
}

// Schema returns the data source schema.
func (d *LoadMobilityControlNodeDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load mobility intent control node by alias", Attributes: map[string]schema.Attribute{"additional_tool_ports": schema.ListAttribute{MarkdownDescription: "Additional Tool ports for the Control node which takes a copy of traffic along with vport that is added into the first level maps. Its a list of ports of format cluster:port", Computed: true, ElementType: types.StringType}, "alias": schema.StringAttribute{MarkdownDescription: "Alias of the Control Node", Computed: true}, "app5_g_http2_ports": schema.ListAttribute{MarkdownDescription: "List of TCP ports", Computed: true, ElementType: types.Int64Type}, "app_tcp": schema.SingleNestedAttribute{MarkdownDescription: "TCP Loadbalancing properties for Control 5G node (PCPN_5G)", Computed: true, Attributes: map[string]schema.Attribute{"application": schema.StringAttribute{MarkdownDescription: "To choose the action on Unknown Application Data", Computed: true}, "load_balance": schema.BoolAttribute{MarkdownDescription: "When true it enables TCP loadbalancing on the Tool Ports", Computed: true}, "tcp_control": schema.StringAttribute{MarkdownDescription: "To choose the action on TCP Control messages", Computed: true}}}, "collector_tools": schema.ListAttribute{MarkdownDescription: "Collector Tool ports. Its a list of ports of format cluster:port", Computed: true, ElementType: types.StringType}, "config_status": schema.StringAttribute{Computed: true}, "config_status_reasons": schema.StringAttribute{Computed: true}, "control_metadata_ip_interface_alias": schema.StringAttribute{MarkdownDescription: "Alias of the ip interface used for exporting control json records", Computed: true}, "cpnode_alias": schema.StringAttribute{MarkdownDescription: "Alias of the control node", Required: true}, "deployed": schema.BoolAttribute{MarkdownDescription: "True when the Control node attempted for deployment", Computed: true}, "gtp_control_sample": schema.BoolAttribute{MarkdownDescription: "When enabled, GTP Control plane traffic for the subscribers contained within the subscriber sample defined in GTP flowsample maps will be sent to tools. When disabled, all GTP Control plane traffic for all subscribers in GTP flowsample maps will be sent to tools.", Computed: true}, "gtp_random_sampling": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Gtp Random Sampling Parameters", Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{MarkdownDescription: "When enabled, sampling of subscriber's sessions happens in random fashion", Computed: true}, "interval": schema.Int64Attribute{MarkdownDescription: "Rotation Interval in multiples of 12 (hrs)", Computed: true}}}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "health_state_reasons": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Computed: true}}}}, "ip_interface_alias": schema.StringAttribute{MarkdownDescription: "Alias of the ip interface used by Control Node", Computed: true}, "location": schema.SingleNestedAttribute{MarkdownDescription: "Location of the engine port", Computed: true, Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{Computed: true}, "engine_ports": schema.ListAttribute{Computed: true, ElementType: types.StringType}}}, "node_override_network_ports": schema.ListAttribute{MarkdownDescription: "Network ports for the Control node. Its a list of ports of format cluster:port", Computed: true, ElementType: types.StringType}, "node_type": schema.StringAttribute{MarkdownDescription: "Type of the Control Node", Computed: true}, "number_of5_g_sessions": schema.Int64Attribute{MarkdownDescription: "Number of 5G sessions to allocate for Control 5G Node (PCPN_5G)", Computed: true}, "number_of_lte_sessions": schema.Int64Attribute{MarkdownDescription: "Number of LTE sessions to allocate for Control LTE Node (PCPN_LTE)", Computed: true}, "solution_alias": schema.StringAttribute{MarkdownDescription: "Alias of the mobility solution", Required: true}, "tags": schema.ListNestedAttribute{MarkdownDescription: "User defined tags (Aggregation tags)", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"tag_key": schema.StringAttribute{MarkdownDescription: "Name of the tag", Computed: true}, "tag_values": schema.ListAttribute{MarkdownDescription: "All possible values of the tag", Computed: true, ElementType: types.StringType}}}}, "traffic_sources": schema.DynamicAttribute{MarkdownDescription: "List of all traffic sources for the Control node", Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadMobilityControlNodeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadMobilityControlNodeDataSourceModel
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
func (d *LoadMobilityControlNodeDataSource) readRemote(ctx context.Context, config *LoadMobilityControlNodeDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/intent/mobility/{solutionAlias}/cpnode/{cpnodeAlias}"
	reqPath = strings.ReplaceAll(reqPath, "{solutionAlias}", url.PathEscape(config.SolutionAlias.ValueString()))
	reqPath = strings.ReplaceAll(reqPath, "{cpnodeAlias}", url.PathEscape(config.CpnodeAlias.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_control_node", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_control_node", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_control_node", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_control_node", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_control_node", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_control_node", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_control_node", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_control_node", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_control_node", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_control_node", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_control_node", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_control_node", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_control_node", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_control_node", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadMobilityControlNodeDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
