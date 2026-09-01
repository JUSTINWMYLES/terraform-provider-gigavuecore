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
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
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
	Items     types.List   `tfsdk:"items"`
	Page      types.String `tfsdk:"page"`
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
	resp.Schema = schema.Schema{MarkdownDescription: "Load all Device Ports", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"admin_status": schema.StringAttribute{Computed: true}, "alias": schema.StringAttribute{MarkdownDescription: "device port alias", Computed: true}, "auto_neg": schema.BoolAttribute{MarkdownDescription: "When auto-negotiation is enabled, duplex and speed settings are ignored (they are set via auto-negotiation). Auto-negotiation is always disabled for 40Gb and 100Gb ports. For 1Gb speeds over copper, auto-negotiation must be enabled, per the IEEE 802.3 specification.", Computed: true}, "breakout_mode": schema.StringAttribute{MarkdownDescription: "none = port not broken out; na = port not capable of breakout; 4x = 4x10G; 2q = 2x40G", Computed: true}, "cable_length": schema.StringAttribute{MarkdownDescription: "Attached cable length in meter", Computed: true}, "comment": schema.StringAttribute{Computed: true}, "config_speed": schema.StringAttribute{MarkdownDescription: "The configured line speed of a port. Only applicable for copper ports. Only applicable if auto-negotiation is off", Computed: true}, "duplex": schema.StringAttribute{MarkdownDescription: "Only applicable for 10M/100M operations. Only applicable if auto-negotiation is off. Duplex is always set to full for 10G ports. If this parameter is set explicitly on one end of the connection, it must also be set explicitly on the other end. Duplex mismatches will occur if the duplex setting are forced on one end of the connection while the other end attempts to auto-negotiate settings", Computed: true}, "force_link_up": schema.BoolAttribute{MarkdownDescription: "Forces connection on an optical port. Use this option when an optical GigaPORT tool port is connected to a legacy optical tool that does not transmit light; Available for optical 1Gb/10Gb tool ports; Not available for 10Gb-capable ports with a 1Gb SFP installed. 10Gb-capable optical tool ports only support force-linkup when a 10Gb SFP+ is installed.", Computed: true}, "force_link_up_status": schema.StringAttribute{MarkdownDescription: "Forcelinkup state for the port configured", Computed: true}, "gsparams": schema.SingleNestedAttribute{MarkdownDescription: "Gsparams Resource HW limit for engine port", Computed: true, Attributes: map[string]schema.Attribute{"bufferasf": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"initial": schema.Int64Attribute{Computed: true}, "max": schema.Int64Attribute{Computed: true}, "min": schema.Int64Attribute{Computed: true}}}, "metadata": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"initial": schema.Int64Attribute{Computed: true}, "max": schema.Int64Attribute{Computed: true}, "min": schema.Int64Attribute{Computed: true}}}}}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "health_state_reasons": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Computed: true}}}}, "host_name": schema.StringAttribute{MarkdownDescription: "device host name", Computed: true}, "is_signal_detected": schema.StringAttribute{Computed: true}, "last_state_transition_time": schema.Int64Attribute{MarkdownDescription: "Port state change timestamp in milliseconds", Computed: true}, "mtu": schema.Int64Attribute{Computed: true}, "oper_speed": schema.StringAttribute{MarkdownDescription: "Port operational runtime speed", Computed: true}, "oper_status": schema.StringAttribute{Computed: true}, "port_id": schema.StringAttribute{MarkdownDescription: "device port id", Computed: true}, "port_role": schema.StringAttribute{MarkdownDescription: "master/slave role of port. (deprecated: use portRoleAlias)", Computed: true}, "port_role_alias": schema.StringAttribute{MarkdownDescription: "source/receiver role of port", Computed: true}, "port_type": schema.StringAttribute{MarkdownDescription: "Configures port type for eligible ports based on the underlying card/chassis hardware type. 'gigasmart' port type can never be explicitly assigned", Computed: true}, "sfp": schema.SingleNestedAttribute{MarkdownDescription: "Port SFP details", Computed: true, Attributes: map[string]schema.Attribute{"sfp_type": schema.StringAttribute{Computed: true}, "vendor_name": schema.StringAttribute{Computed: true}, "vendor_pn": schema.StringAttribute{Computed: true}, "vendor_sn": schema.StringAttribute{Computed: true}}}, "ude": schema.SingleNestedAttribute{MarkdownDescription: "Unidirectional Ethernet", Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{MarkdownDescription: "Only applicable if 100g-bidi is detected", Computed: true}}}}}}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadGigaPortsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadGigaPortsDataSourceModel
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
func (d *LoadGigaPortsDataSource) readListRemote(ctx context.Context, config *LoadGigaPortsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/inventory/ports"
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
		resp.Diagnostics.AddError("Error reading gigavuecore_load_giga_ports", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_giga_ports", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["ports"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_giga_ports", fmt.Sprintf("Could not decode list page: missing %q array", "ports"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
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
