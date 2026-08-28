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
	_ datasource.DataSource              = (*LoadAllGsopsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadAllGsopsDataSource)(nil)
)

// LoadAllGsopsDataSource is the generated Terraform data source implementation.
type LoadAllGsopsDataSource struct {
	client *client.Client
}

// LoadAllGsopsDataSourceModel describes the data source state shape.
type LoadAllGsopsDataSourceModel struct {
	ClusterId types.String `tfsdk:"cluster_id" json:"clusterId"`
	Items     types.List   `tfsdk:"items"`
	Page      types.String `tfsdk:"page"`
	Sort      types.String `tfsdk:"sort"`
}

// NewLoadAllGsopsDataSource returns a new instance of the generated data source.
func NewLoadAllGsopsDataSource() datasource.DataSource {
	return &LoadAllGsopsDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadAllGsopsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_all_gsops"
}

// Schema returns the data source schema.
func (d *LoadAllGsopsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load all GSOPs", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{Computed: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "id of the defining cluster", Computed: true}, "gs_apps": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART Applications for a GSOP. At least one GsApp must be defined for the config to be valid", Computed: true, Attributes: map[string]schema.Attribute{"apf": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Computed: true}}}, "dedup": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Computed: true}}}, "diameter_whitelist": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Computed: true}}}, "flow_filter": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Flow Filter' Application Configuration", Computed: true, Attributes: map[string]schema.Attribute{"type": schema.StringAttribute{Computed: true}}}, "flow_sampling": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Flow Sampling' Application Configuration", Computed: true, Attributes: map[string]schema.Attribute{"type": schema.StringAttribute{Computed: true}}}, "gseries_header_add": schema.SingleNestedAttribute{MarkdownDescription: "Only applicable for G-series", Computed: true, Attributes: map[string]schema.Attribute{"types": schema.SetAttribute{Computed: true, ElementType: types.StringType}}}, "gseries_header_remove": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Computed: true}}}, "gseries_load_balance": schema.SingleNestedAttribute{MarkdownDescription: "Only applicable for G-series per-rule GSOP", Computed: true, Attributes: map[string]schema.Attribute{"fixed_offset": schema.SingleNestedAttribute{MarkdownDescription: "Per-rule Fixed Offset Load Balancing config for G-seres devices", Computed: true, Attributes: map[string]schema.Attribute{"hash": schema.StringAttribute{Computed: true}, "length": schema.Int64Attribute{Computed: true}, "offset": schema.Int64Attribute{Computed: true}}}, "variable_offset": schema.SingleNestedAttribute{MarkdownDescription: "Per-rule Variable Offset Load Balancing config for G-seres devices", Computed: true, Attributes: map[string]schema.Attribute{"end_delim": schema.StringAttribute{Computed: true}, "hash": schema.StringAttribute{Computed: true}, "start_delim": schema.StringAttribute{Computed: true}, "start_field": schema.StringAttribute{Computed: true}}}}}, "gseries_pattern_match": schema.SingleNestedAttribute{MarkdownDescription: "Only applicable for G-series per-rule GSOP", Computed: true, Attributes: map[string]schema.Attribute{"fixed_offset": schema.SingleNestedAttribute{MarkdownDescription: "Per-rule Fixed Offset Pattern Match config for G-seres devices", Computed: true, Attributes: map[string]schema.Attribute{"length": schema.Int64Attribute{Computed: true}, "offset": schema.Int64Attribute{Computed: true}}}, "variable_offset": schema.SingleNestedAttribute{MarkdownDescription: "Per-rule Variable Offset Pattern Match config for G-seres devices", Computed: true, Attributes: map[string]schema.Attribute{"end_delim": schema.StringAttribute{Computed: true}, "start_delim": schema.StringAttribute{Computed: true}}}}}, "gtp_whitelist": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Computed: true}}}, "header_add": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Add Header' Application Configuration", Computed: true, Attributes: map[string]schema.Attribute{"vlan": schema.Int64Attribute{Computed: true}}}, "header_remove": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Remove Header' Application Configuration", Computed: true, Attributes: map[string]schema.Attribute{"ah1": schema.StringAttribute{MarkdownDescription: "only valid and required for 'generic'. First anchor header, from which the header to be stripped is occurred.", Computed: true}, "ah2": schema.StringAttribute{MarkdownDescription: "only valid and required for 'generic'. next anchor header.", Computed: true}, "custom_len": schema.Int64Attribute{MarkdownDescription: "only valid for 'generic'. length of unknown header.", Computed: true}, "erspan_flow_id": schema.Int64Attribute{MarkdownDescription: "only applicable for 'erspan', specifies the flow id to strip. Value of 0 will strip all flow ids", Computed: true}, "fp_dst_switch_id": schema.Int64Attribute{MarkdownDescription: "Only valid and required for 'fabricPath', 12 bit destination switch id", Computed: true}, "fp_src_switch_id": schema.Int64Attribute{MarkdownDescription: "Only valid and required for 'fabricPath', 12 bit source switch id", Computed: true}, "header_count": schema.Int64Attribute{MarkdownDescription: "only valid for 'generic'. Number of headers to be stripped.", Computed: true}, "offset": schema.StringAttribute{MarkdownDescription: "only valid and required for 'generic'.'start': strip from start of ah1, 'end': strip from end of ah1 or any other offset in the range of length of ah1, 'offsetRangeValue' is required if offset is 'offsetRange'.", Computed: true}, "offset_range_value": schema.Int64Attribute{MarkdownDescription: "only valid and required when offset is 'offsetRange', integer within range of size of header", Computed: true}, "protocol": schema.StringAttribute{MarkdownDescription: "'gre' and 'fabricPath' are only applicable for H-series", Computed: true}, "timestamp_format": schema.StringAttribute{MarkdownDescription: "Timestamp format. Only valid and required for 'fm6000Ts'", Computed: true}, "vlan_header": schema.StringAttribute{MarkdownDescription: "Only valid when protocol is 'vlan'. Specifies the target vlan header to strip. Defaults to 'outer'", Computed: true}, "vxlan_id": schema.Int64Attribute{MarkdownDescription: "24-bit value. Only valid when protocol is 'vxlan'. Specifies the vxlan id to strip. Value of '0' will strip all vxlan ids", Computed: true}}}, "icap": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART ICAP Configuration", Computed: true, Attributes: map[string]schema.Attribute{"icap_profile": schema.StringAttribute{MarkdownDescription: "Alias of referenced ICAP Profile", Computed: true}}}, "inline_ssl": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART Inline SSL Profile Configuration", Computed: true, Attributes: map[string]schema.Attribute{"inline_ssl_profile": schema.StringAttribute{MarkdownDescription: "Alias of referenced Inline SSL Profile", Computed: true}}}, "load_balance": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Load Balancing' Application Configuration", Computed: true, Attributes: map[string]schema.Attribute{"enhanced": schema.SingleNestedAttribute{MarkdownDescription: "Private class. Enhanced part of the GigaSMART 'Load Balancing' Application Configuration", Computed: true, Attributes: map[string]schema.Attribute{"elb_alias": schema.StringAttribute{MarkdownDescription: "elb app alias", Computed: true}}}, "stateful": schema.SingleNestedAttribute{MarkdownDescription: "Private class. Stateful part of the GigaSMART 'Load Balancing' Application Configuration", Computed: true, Attributes: map[string]schema.Attribute{"app_type": schema.StringAttribute{MarkdownDescription: "'sapf' option is added in release H 4.3.01. 'sip' and 'tunnel' added in release H 5.1. 'diameter' is added in release H 5.5", Computed: true}, "diameter_key_hash_type": schema.StringAttribute{MarkdownDescription: "required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyHash'", Computed: true}, "diameter_key_multi_hash_type": schema.ListNestedAttribute{MarkdownDescription: "required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyMultiHash'", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"avp_codevalue": schema.Int64Attribute{MarkdownDescription: "required when 'key' == 'avpCode'", Computed: true}, "key": schema.StringAttribute{Computed: true}}}}, "gtp_key_hash_type": schema.StringAttribute{MarkdownDescription: "required when 'appType' == 'gtp' and 'lbType' == 'gtpKeyHash'. ignored otherwise", Computed: true}, "lb_type": schema.StringAttribute{MarkdownDescription: "'gtpKeyHash' is only applicable when 'appType' == 'gtp'. 'sipKeyHash' is only applicable when 'appType' == 'sip'", Computed: true}, "sip_key_hash_type": schema.StringAttribute{MarkdownDescription: "required when 'appType' == 'sip' and 'lbType' == 'sipKeyHash'. ignored otherwise", Computed: true}}}, "stateless": schema.SingleNestedAttribute{MarkdownDescription: "Private class. Stateless part of the GigaSMART 'Load Balancing' Application Configuration", Computed: true, Attributes: map[string]schema.Attribute{"field_location": schema.StringAttribute{MarkdownDescription: "Ignored when 'hashFields' == 'gtpuTeid'. Required otherwise", Computed: true}, "hash_fields": schema.StringAttribute{Computed: true}}}}}, "masking": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Masking' Application Configuration", Computed: true, Attributes: map[string]schema.Attribute{"content_type": schema.StringAttribute{MarkdownDescription: "content type that will trigger masking, only valid and required for protocol 'sip'", Computed: true}, "length": schema.Int64Attribute{MarkdownDescription: "max length is 9600 on H-series; 9000 on G-series, not applicable for protocol 'sip' valid and required otherwise", Computed: true}, "offset": schema.Int64Attribute{MarkdownDescription: "min offset is 0 for 'none'; 1 for protocol-referenced offsets, not applicable for protocol 'sip' valid and required otherwise", Computed: true}, "pattern": schema.StringAttribute{MarkdownDescription: "1-byte hex mask, not applicable for protocol 'sip' valid and required otherwise", Computed: true}, "protocol": schema.StringAttribute{Computed: true}}}, "metadata_export": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"cache": schema.StringAttribute{MarkdownDescription: "metadata cache alias. cache should have exporters defined, and the exporters should have applicationProfiles defined", Computed: true}}}, "netflow": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Computed: true}}}, "sa_apf": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Computed: true}}}, "sip_whitelist": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Computed: true}}}, "slicing": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Slicing' Application Configuration", Computed: true, Attributes: map[string]schema.Attribute{"enhanced": schema.StringAttribute{MarkdownDescription: "enhanced-slicing apps alias", Computed: true}, "offset": schema.Int64Attribute{MarkdownDescription: "min offset is 64 for 'none'; 4 for protocol-referenced offsets. required property till H 5.6", Computed: true}, "protocol": schema.StringAttribute{MarkdownDescription: "required property till H 5.6", Computed: true}}}, "ssl_decrypt": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'SSL Decrypt' Application Configuration", Computed: true, Attributes: map[string]schema.Attribute{"in_port": schema.Int64Attribute{MarkdownDescription: "Port number of 0 represents 'any' port", Computed: true}, "out_port": schema.Int64Attribute{MarkdownDescription: "Port number of 0 represents 'auto' port", Computed: true}}}, "trailer_add": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Add Trailer' Application Configuration", Computed: true, Attributes: map[string]schema.Attribute{"types": schema.SetAttribute{MarkdownDescription: "'crc' is not applicable for G-series", Computed: true, ElementType: types.StringType}}}, "trailer_remove": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Computed: true}}}, "tunnel_decap": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Decapsulate Tunnel' Application Configuration", Computed: true, Attributes: map[string]schema.Attribute{"custom": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"port_dst": schema.Int64Attribute{MarkdownDescription: "when specified as 0, no validation will be done in the packet.", Computed: true}, "port_src": schema.Int64Attribute{MarkdownDescription: "when specified as 0, no validation will be done in the packet.", Computed: true}}}, "erspan_flow_id": schema.Int64Attribute{MarkdownDescription: "only applicable for 'erspan', A Flow ID of 0 decapsulates all ERSPAN tunnel traffic regardless of Flow ID", Computed: true}, "gmip_port": schema.Int64Attribute{MarkdownDescription: "only applicable for 'gmip', in which case it is required. Specifies the UDP port on which the Tunnel Network port on the receiving GigaVUE H Series is listening. Must match the configuration of the portdst configured on the sending end of the tunnel", Computed: true}, "l2_gre_key": schema.Int64Attribute{MarkdownDescription: "only applicable for 'l2gre', in which case it is required.", Computed: true}, "tls_pcapng": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"decap_key": schema.StringAttribute{Computed: true}, "listener": schema.StringAttribute{Computed: true}}}, "type": schema.StringAttribute{Computed: true}, "vxlan": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"port_dst": schema.Int64Attribute{Computed: true}, "port_src": schema.Int64Attribute{MarkdownDescription: "when specified as 0, no validation will be done in the packet.", Computed: true}, "vni": schema.Int64Attribute{Computed: true}}}}}, "tunnel_encap": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Encapsulate Tunnel' Application Configuration", Computed: true, Attributes: map[string]schema.Attribute{"gmip_config": schema.SingleNestedAttribute{MarkdownDescription: "Configuration for GMIP Tunnel Encapsulate GigaSmaprt App", Computed: true, Attributes: map[string]schema.Attribute{"dscp": schema.Int64Attribute{Computed: true}, "dst_ip": schema.StringAttribute{MarkdownDescription: "IP Destination. IPv4 or IPv6.", Computed: true}, "dst_port": schema.Int64Attribute{Computed: true}, "flow_label": schema.Int64Attribute{Computed: true}, "prec": schema.Int64Attribute{MarkdownDescription: "decimal Precedence value from 0-7 to be used in the ToS byte of the outer headers of tunneled packets", Computed: true}, "src_port": schema.Int64Attribute{Computed: true}, "ttl": schema.Int64Attribute{Computed: true}}}, "l2_gre_config": schema.SingleNestedAttribute{MarkdownDescription: "Configuration for l2GRE Tunnel Encapsulate GigaSmaprt App", Computed: true, Attributes: map[string]schema.Attribute{"dscp": schema.Int64Attribute{Computed: true}, "dst_ip": schema.StringAttribute{MarkdownDescription: "ip destination, mutually exclusive with 'pgDst'. IPv4 or IPv6.", Computed: true}, "flow_label": schema.Int64Attribute{Computed: true}, "key": schema.Int64Attribute{Computed: true}, "pg_dst": schema.StringAttribute{MarkdownDescription: "port group destination alias, mutually exclusive with 'dstIp'", Computed: true}, "prec": schema.Int64Attribute{MarkdownDescription: "decimal Precedence value from 0-7 to be used in the ToS byte of the outer headers of tunneled packets", Computed: true}, "session_field": schema.StringAttribute{MarkdownDescription: "required with stateful loadBalance when 'appType' is 'tunnel'", Computed: true}, "session_pos": schema.StringAttribute{MarkdownDescription: "required if 'sessionField' is specified", Computed: true}, "ttl": schema.Int64Attribute{Computed: true}}}, "tls_pcapng": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"exporter": schema.StringAttribute{Computed: true}, "exporter_group": schema.StringAttribute{Computed: true}}}, "type": schema.StringAttribute{Computed: true}, "vxlan_config": schema.SingleNestedAttribute{MarkdownDescription: "Configuration for Vxlan Tunnel Encapsulate GigaSmaprt App", Computed: true, Attributes: map[string]schema.Attribute{"dscp": schema.Int64Attribute{Computed: true}, "dst_ip": schema.StringAttribute{MarkdownDescription: "IP Destination. IPv4 or IPv6.", Computed: true}, "dst_port": schema.Int64Attribute{Computed: true}, "src_port": schema.Int64Attribute{Computed: true}, "ttl": schema.Int64Attribute{Computed: true}, "vni": schema.Int64Attribute{Computed: true}}}}}}}, "gs_group": schema.StringAttribute{MarkdownDescription: "Alias of referenced managing GsGroup", Computed: true}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "health_state_reasons": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Computed: true}}}}}}}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadAllGsopsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadAllGsopsDataSourceModel
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
func (d *LoadAllGsopsDataSource) readListRemote(ctx context.Context, config *LoadAllGsopsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/gsops"
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
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_gsops", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_gsops", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["gsops"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_gsops", fmt.Sprintf("Could not decode list page: missing %q array", "gsops"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_gsops", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadAllGsopsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
