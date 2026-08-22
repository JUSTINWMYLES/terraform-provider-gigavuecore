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
	_ datasource.DataSource              = (*LoadGigaPortsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadGigaPortsDataSource)(nil)
)

// LoadGigaPortsDataSource is the generated Terraform data source implementation.
type LoadGigaPortsDataSource struct {
	client *client.Client
}

// LoadGigaPortsDataSourceModel describes the data source state shape.
type LoadGigaPortsDataSourceModel struct {
	ClusterId types.String `tfsdk:"cluster_id" json:"clusterId"`
	Context   types.Object `tfsdk:"context"`
	Page      types.String `tfsdk:"page"`
	Ports     types.List   `tfsdk:"ports"`
	Sort      types.String `tfsdk:"sort"`
}

// NewLoadGigaPortsDataSource returns a new instance of the generated data source.
func NewLoadGigaPortsDataSource() datasource.DataSource {
	return &LoadGigaPortsDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadGigaPortsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_giga_ports"
}

// Schema returns the data source schema.
func (d *LoadGigaPortsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load all Device Ports", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "context": schema.SingleNestedAttribute{MarkdownDescription: "Gigamon query result context", Computed: true, Attributes: map[string]schema.Attribute{"page_no": schema.Int64Attribute{MarkdownDescription: "page number of the returned result set", Computed: true}, "page_size": schema.Int64Attribute{MarkdownDescription: "page size of the returned result set", Computed: true}, "sort": schema.ListAttribute{MarkdownDescription: "sorting info of the returned result set. list of fields in the array indicate sorting order", Computed: true, ElementType: types.StringType}, "total_items": schema.Int64Attribute{MarkdownDescription: "total number of items in the queried entity type", Computed: true}}}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "ports": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"admin_status": schema.StringAttribute{Computed: true}, "alias": schema.StringAttribute{MarkdownDescription: "device port alias", Computed: true}, "auto_neg": schema.BoolAttribute{MarkdownDescription: "When auto-negotiation is enabled, duplex and speed settings are ignored (they are set via auto-negotiation). Auto-negotiation is always disabled for 40Gb and 100Gb ports. For 1Gb speeds over copper, auto-negotiation must be enabled, per the IEEE 802.3 specification.", Computed: true}, "breakout_mode": schema.StringAttribute{MarkdownDescription: "none = port not broken out; na = port not capable of breakout; 4x = 4x10G; 2q = 2x40G", Computed: true}, "cable_length": schema.StringAttribute{MarkdownDescription: "Attached cable length in meter", Computed: true}, "comment": schema.StringAttribute{Computed: true}, "config_speed": schema.StringAttribute{MarkdownDescription: "The configured line speed of a port. Only applicable for copper ports. Only applicable if auto-negotiation is off", Computed: true}, "duplex": schema.StringAttribute{MarkdownDescription: "Only applicable for 10M/100M operations. Only applicable if auto-negotiation is off. Duplex is always set to full for 10G ports. If this parameter is set explicitly on one end of the connection, it must also be set explicitly on the other end. Duplex mismatches will occur if the duplex setting are forced on one end of the connection while the other end attempts to auto-negotiate settings", Computed: true}, "force_link_up": schema.BoolAttribute{MarkdownDescription: "Forces connection on an optical port. Use this option when an optical GigaPORT tool port is connected to a legacy optical tool that does not transmit light; Available for optical 1Gb/10Gb tool ports; Not available for 10Gb-capable ports with a 1Gb SFP installed. 10Gb-capable optical tool ports only support force-linkup when a 10Gb SFP+ is installed.", Computed: true}, "force_link_up_status": schema.StringAttribute{MarkdownDescription: "Forcelinkup state for the port configured", Computed: true}, "gsparams": schema.SingleNestedAttribute{MarkdownDescription: "Gsparams Resource HW limit for engine port", Computed: true, Attributes: map[string]schema.Attribute{"bufferasf": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"initial": schema.Int64Attribute{Computed: true}, "max": schema.Int64Attribute{Computed: true}, "min": schema.Int64Attribute{Computed: true}}}, "metadata": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"initial": schema.Int64Attribute{Computed: true}, "max": schema.Int64Attribute{Computed: true}, "min": schema.Int64Attribute{Computed: true}}}}}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "health_state_reasons": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Computed: true}}}}, "host_name": schema.StringAttribute{MarkdownDescription: "device host name", Computed: true}, "is_signal_detected": schema.StringAttribute{Computed: true}, "last_state_transition_time": schema.Int64Attribute{MarkdownDescription: "Port state change timestamp in milliseconds", Computed: true}, "mtu": schema.Int64Attribute{Computed: true}, "oper_speed": schema.StringAttribute{MarkdownDescription: "Port operational runtime speed", Computed: true}, "oper_status": schema.StringAttribute{Computed: true}, "port_id": schema.StringAttribute{MarkdownDescription: "device port id", Computed: true}, "port_role": schema.StringAttribute{MarkdownDescription: "master/slave role of port. (deprecated: use portRoleAlias)", Computed: true}, "port_role_alias": schema.StringAttribute{MarkdownDescription: "source/receiver role of port", Computed: true}, "port_type": schema.StringAttribute{MarkdownDescription: "Configures port type for eligible ports based on the underlying card/chassis hardware type. 'gigasmart' port type can never be explicitly assigned", Computed: true}, "sfp": schema.SingleNestedAttribute{MarkdownDescription: "Port SFP details", Computed: true, Attributes: map[string]schema.Attribute{"sfp_type": schema.StringAttribute{Computed: true}, "vendor_name": schema.StringAttribute{Computed: true}, "vendor_pn": schema.StringAttribute{Computed: true}, "vendor_sn": schema.StringAttribute{Computed: true}}}, "ude": schema.SingleNestedAttribute{MarkdownDescription: "Unidirectional Ethernet", Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{MarkdownDescription: "Only applicable if 100g-bidi is detected", Computed: true}}}}}}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadGigaPortsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadGigaPortsDataSourceModel
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
func (d *LoadGigaPortsDataSource) readRemote(ctx context.Context, config *LoadGigaPortsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/inventory/ports"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_giga_ports", fmt.Sprintf("Could not build request: %s", err))
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
		resp.Diagnostics.AddError("Error reading gigavuecore_load_giga_ports", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_giga_ports", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_giga_ports", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_giga_ports", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_giga_ports", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_giga_ports", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_giga_ports", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_giga_ports", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_giga_ports", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_load_giga_ports", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_load_giga_ports", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_giga_ports", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_giga_ports", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadGigaPortsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
