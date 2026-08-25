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
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*LoadMobilitySamNodeDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadMobilitySamNodeDataSource)(nil)
)

// LoadMobilitySamNodeDataSource is the generated Terraform data source implementation.
type LoadMobilitySamNodeDataSource struct {
	client *client.Client
}

// LoadMobilitySamNodeDataSourceModel describes the data source state shape.
type LoadMobilitySamNodeDataSourceModel struct {
	Alias                      types.String  `tfsdk:"alias"`
	AppProfileConfig           types.Object  `tfsdk:"app_profile_config" json:"appProfileConfig"`
	ConfigStatus               types.String  `tfsdk:"config_status" json:"configStatus"`
	ConfigStatusReasons        types.String  `tfsdk:"config_status_reasons" json:"configStatusReasons"`
	ControlPlaneSetting        types.Object  `tfsdk:"control_plane_setting" json:"controlPlaneSetting"`
	Deployed                   types.Bool    `tfsdk:"deployed"`
	DeploymentDetails          types.List    `tfsdk:"deployment_details" json:"deploymentDetails"`
	EngineMetaDataCacheConfigs types.List    `tfsdk:"engine_meta_data_cache_configs" json:"engineMetaDataCacheConfigs"`
	EngineSourceMappings       types.List    `tfsdk:"engine_source_mappings" json:"engineSourceMappings"`
	ExporterConfig             types.Object  `tfsdk:"exporter_config" json:"exporterConfig"`
	HealthState                types.String  `tfsdk:"health_state" json:"healthState"`
	HealthStateReasons         types.List    `tfsdk:"health_state_reasons" json:"healthStateReasons"`
	IpInterfaceAlias           types.String  `tfsdk:"ip_interface_alias" json:"ipInterfaceAlias"`
	Location                   types.Object  `tfsdk:"location"`
	NodeOverrideNetworkPorts   types.List    `tfsdk:"node_override_network_ports" json:"nodeOverrideNetworkPorts"`
	NodeType                   types.String  `tfsdk:"node_type" json:"nodeType"`
	ParamConfigs               types.List    `tfsdk:"param_configs" json:"paramConfigs"`
	SamnodeAlias               types.String  `tfsdk:"samnode_alias" json:"samnodeAlias"`
	SmafDetails                types.List    `tfsdk:"smaf_details" json:"SMAFDetails"`
	SolutionAlias              types.String  `tfsdk:"solution_alias" json:"solutionAlias"`
	Tags                       types.List    `tfsdk:"tags"`
	TrafficSources             types.Dynamic `tfsdk:"traffic_sources" json:"trafficSources"`
}

// NewLoadMobilitySamNodeDataSource returns a new instance of the generated data source.
func NewLoadMobilitySamNodeDataSource() datasource.DataSource {
	return &LoadMobilitySamNodeDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadMobilitySamNodeDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_mobility_sam_node"
}

// Schema returns the data source schema.
func (d *LoadMobilitySamNodeDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load mobility intent sam node by alias", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Alias of the SAM Exporter node", Computed: true}, "app_profile_config": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"application_id": schema.BoolAttribute{Computed: true}, "applications": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"attributes": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"name": schema.StringAttribute{MarkdownDescription: "attribute name", Computed: true}, "value": schema.StringAttribute{MarkdownDescription: "application's attribute value", Computed: true}}}}, "is_user_defined": schema.BoolAttribute{MarkdownDescription: "Default value is false. Must be set to true only for App Intel solution while configuring user defined apps", Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "application name", Computed: true}}}, "counter": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"bytes": schema.BoolAttribute{Computed: true}, "bytes_long": schema.BoolAttribute{Computed: true}, "packets": schema.BoolAttribute{Computed: true}, "packets_long": schema.BoolAttribute{Computed: true}}}, "flow": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"end_reason": schema.BoolAttribute{Computed: true}}}, "gtpu": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"qfi": schema.BoolAttribute{Computed: true}, "teid": schema.BoolAttribute{Computed: true}}}, "inner_ipv4": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"destination": schema.BoolAttribute{Computed: true}, "protocol": schema.BoolAttribute{Computed: true}, "source": schema.BoolAttribute{Computed: true}}}, "inner_ipv6": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"destination": schema.BoolAttribute{Computed: true}, "next_header": schema.BoolAttribute{Computed: true}, "source": schema.BoolAttribute{Computed: true}}}, "outer_ipv4": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"destination": schema.BoolAttribute{Computed: true}, "source": schema.BoolAttribute{Computed: true}}}, "outer_ipv6": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"destination": schema.BoolAttribute{Computed: true}, "source": schema.BoolAttribute{Computed: true}}}, "timestamp": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"flow_end_msec": schema.BoolAttribute{Computed: true}, "flow_endsec": schema.BoolAttribute{Computed: true}, "flow_start_msec": schema.BoolAttribute{Computed: true}, "flow_startsec": schema.BoolAttribute{Computed: true}}}, "transport": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"dst_port": schema.BoolAttribute{Computed: true}, "src_port": schema.BoolAttribute{Computed: true}}}}}, "config_status": schema.StringAttribute{Computed: true}, "config_status_reasons": schema.StringAttribute{Computed: true}, "control_plane_setting": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"encoding": schema.StringAttribute{Computed: true}, "encoding_format": schema.StringAttribute{Computed: true}, "event_enable": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"modify": schema.BoolAttribute{Computed: true}, "update": schema.BoolAttribute{Computed: true}}}, "trigger": schema.StringAttribute{Computed: true}}}, "deployed": schema.BoolAttribute{MarkdownDescription: "True when the SAM node is attempted for deployment", Computed: true}, "deployment_details": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"engine_port": schema.StringAttribute{Computed: true}, "sam_node_alias": schema.StringAttribute{Computed: true}}}}, "engine_meta_data_cache_configs": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"engine_port": schema.StringAttribute{Computed: true}, "event": schema.StringAttribute{Computed: true}, "flow_behavior": schema.StringAttribute{Computed: true}, "flows_size": schema.Int64Attribute{Computed: true}, "idle_timeout": schema.Int64Attribute{Computed: true}, "match": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"ipv4": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"destination": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"prefix_min_mask": schema.StringAttribute{Computed: true}}}, "protocol": schema.BoolAttribute{Computed: true}, "source": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"prefix_min_mask": schema.StringAttribute{Computed: true}}}}}, "ipv6": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"destination": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"prefix_min_mask": schema.StringAttribute{Computed: true}}}, "next_header": schema.BoolAttribute{Computed: true}, "source": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"prefix_min_mask": schema.StringAttribute{Computed: true}}}}}, "transport": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"dst_port": schema.BoolAttribute{Computed: true}, "src_port": schema.BoolAttribute{Computed: true}}}}}, "observation_domain_id": schema.Int64Attribute{Computed: true}}}}, "engine_source_mappings": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"engine_port": schema.StringAttribute{Computed: true}, "network_source": schema.StringAttribute{Computed: true}}}}, "exporter_config": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"active_timeout": schema.StringAttribute{Computed: true}, "inactive_timeout": schema.Int64Attribute{Computed: true}, "record_type": schema.StringAttribute{Computed: true}}}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "health_state_reasons": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Computed: true}}}}, "ip_interface_alias": schema.StringAttribute{MarkdownDescription: "Alias of the ip interface used by SAM Exporter Node", Computed: true}, "location": schema.SingleNestedAttribute{MarkdownDescription: "Location of the engine port", Computed: true, Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{Computed: true}, "engine_ports": schema.ListAttribute{Computed: true, ElementType: types.StringType}}}, "node_override_network_ports": schema.ListAttribute{MarkdownDescription: "Network ports for the SAM Exporter node. Its a list of ports of format cluster:port", Computed: true, ElementType: types.StringType}, "node_type": schema.StringAttribute{MarkdownDescription: "Type of the SAM Exporter Node", Computed: true}, "param_configs": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"engine_port": schema.StringAttribute{Computed: true}, "resource_metadata": schema.Int64Attribute{Computed: true}}}}, "samnode_alias": schema.StringAttribute{MarkdownDescription: "Alias of the sam node", Required: true}, "smaf_details": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"control_application": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"interface_address": schema.StringAttribute{Computed: true}, "port": schema.Int64Attribute{Computed: true}, "protocol": schema.StringAttribute{Computed: true}}}, "management_address": schema.StringAttribute{Computed: true}, "user_application": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"interface_address": schema.StringAttribute{Computed: true}, "port": schema.Int64Attribute{Computed: true}, "protocol": schema.StringAttribute{Computed: true}}}}}}, "solution_alias": schema.StringAttribute{MarkdownDescription: "Alias of the mobility solution", Required: true}, "tags": schema.ListNestedAttribute{MarkdownDescription: "User defined tags (Aggregation tags)", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"tag_key": schema.StringAttribute{MarkdownDescription: "Name of the tag", Computed: true}, "tag_values": schema.ListAttribute{MarkdownDescription: "All possible values of the tag", Computed: true, ElementType: types.StringType}}}}, "traffic_sources": schema.DynamicAttribute{MarkdownDescription: "List of all traffic sources for the SAM Exporter node", Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadMobilitySamNodeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadMobilitySamNodeDataSourceModel
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
func (d *LoadMobilitySamNodeDataSource) readRemote(ctx context.Context, config *LoadMobilitySamNodeDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/intent/mobility/{solutionAlias}/samnode/{samnodeAlias}"
	reqPath = strings.ReplaceAll(reqPath, "{solutionAlias}", url.PathEscape(config.SolutionAlias.ValueString()))
	reqPath = strings.ReplaceAll(reqPath, "{samnodeAlias}", url.PathEscape(config.SamnodeAlias.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_sam_node", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_sam_node", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_sam_node", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_sam_node", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_sam_node", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_sam_node", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_sam_node", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_sam_node", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_sam_node", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_sam_node", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_sam_node", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_sam_node", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_sam_node", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_mobility_sam_node", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadMobilitySamNodeDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
