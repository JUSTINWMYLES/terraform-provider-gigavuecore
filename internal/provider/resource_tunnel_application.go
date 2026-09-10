package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	regexp "regexp"
	"strings"
	"time"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	int64validator "github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	setvalidator "github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	stringvalidator "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	path "github.com/hashicorp/terraform-plugin-framework/path"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	identityschema "github.com/hashicorp/terraform-plugin-framework/resource/identityschema"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	validator "github.com/hashicorp/terraform-plugin-framework/schema/validator"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertions.
var (
	_ resource.Resource                = (*TunnelApplicationResource)(nil)
	_ resource.ResourceWithIdentity    = (*TunnelApplicationResource)(nil)
	_ resource.ResourceWithImportState = (*TunnelApplicationResource)(nil)
	_ resource.ResourceWithConfigure   = (*TunnelApplicationResource)(nil)
)

// TunnelApplicationResource is the generated Terraform managed resource implementation.
type TunnelApplicationResource struct {
	client *client.Client
}

// TunnelApplicationResourceModel describes the Terraform state and plan shape for TunnelApplicationResource.
type TunnelApplicationResourceModel struct {
	Alias               types.String                    `tfsdk:"alias"`
	ClusterId           types.String                    `tfsdk:"cluster_id" json:"clusterId"`
	ConfigStatus        types.String                    `tfsdk:"config_status" json:"configStatus"`
	ConfigStatusReasons types.List                      `tfsdk:"config_status_reasons" json:"configStatusReasons"`
	Decap               types.Dynamic                   `tfsdk:"decap"`
	Description         types.String                    `tfsdk:"description"`
	Dscp                types.Int64                     `tfsdk:"dscp"`
	Encap               types.Object                    `tfsdk:"encap"`
	Gsgroup             types.String                    `tfsdk:"gsgroup"`
	IpInterfaceIn       types.String                    `tfsdk:"ip_interface_in" json:"ipInterfaceIn"`
	IpInterfaceOut      types.String                    `tfsdk:"ip_interface_out" json:"ipInterfaceOut"`
	LocalKeyAlias       types.String                    `tfsdk:"local_key_alias" json:"localKeyAlias"`
	RemoteKeyAlias      types.String                    `tfsdk:"remote_key_alias" json:"remoteKeyAlias"`
	SslProfile          types.Object                    `tfsdk:"ssl_profile" json:"sslProfile"`
	SslProfileAlias     types.String                    `tfsdk:"ssl_profile_alias" json:"sslProfileAlias"`
	TcpProfile          types.Object                    `tfsdk:"tcp_profile" json:"tcpProfile"`
	TcpProfileAlias     types.String                    `tfsdk:"tcp_profile_alias" json:"tcpProfileAlias"`
	TrafficDir          types.String                    `tfsdk:"traffic_dir" json:"trafficDir"`
	Ttl                 types.Int64                     `tfsdk:"ttl"`
	TunnelType          types.String                    `tfsdk:"tunnel_type" json:"tunnelType"`
	Timeouts            *TunnelApplicationTimeoutsModel `tfsdk:"timeouts"`
}

// TunnelApplicationTimeoutsModel describes the per-operation timeout configuration (in seconds) for the gigavuecore_tunnel_application resource.
type TunnelApplicationTimeoutsModel struct {
	Create types.Int64 `tfsdk:"create"`
	Read   types.Int64 `tfsdk:"read"`
	Update types.Int64 `tfsdk:"update"`
	Delete types.Int64 `tfsdk:"delete"`
}

// Metadata returns the resource type name.
func (r *TunnelApplicationResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_tunnel_application"
}

// Schema returns the Terraform schema for this resource.
func (r *TunnelApplicationResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Create Tunnel Apps", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{Optional: true, Computed: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "config_status": schema.StringAttribute{Optional: true, Computed: true}, "config_status_reasons": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType}, "decap": schema.DynamicAttribute{Optional: true, Computed: true}, "description": schema.StringAttribute{Optional: true, Computed: true}, "dscp": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 63)}}, "encap": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"additonalgsops": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART Applications for a GSOP. At least one GsApp must be defined for the config to be valid", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"apf": schema.SingleNestedAttribute{MarkdownDescription: "Only applicable for H-series", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("enabled")}}}}, "dedup": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("enabled")}}}}, "diameter_whitelist": schema.SingleNestedAttribute{MarkdownDescription: "Only applicable for H-series", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("enabled")}}}}, "flow_filter": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Flow Filter' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"type": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.OneOf("gtp")}}}}, "flow_sampling": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Flow Sampling' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"type": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.OneOf("ip", "gtp", "sip", "diameter")}}}}, "gseries_header_add": schema.SingleNestedAttribute{MarkdownDescription: "Only applicable for G-series", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"types": schema.SetAttribute{Required: true, ElementType: types.StringType, Validators: []validator.Set{setvalidator.ValueStringsAre(stringvalidator.OneOf("srcid", "timestamp"))}}}}, "gseries_header_remove": schema.SingleNestedAttribute{MarkdownDescription: "Only applicable for G-series", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("enabled")}}}}, "gseries_load_balance": schema.SingleNestedAttribute{MarkdownDescription: "Only applicable for G-series per-rule GSOP", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"fixed_offset": schema.SingleNestedAttribute{MarkdownDescription: "Per-rule Fixed Offset Load Balancing config for G-seres devices", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"hash": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.OneOf("checksum", "xor", "crc")}}, "length": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.AtLeast(1)}}, "offset": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.AtLeast(0)}}}}, "variable_offset": schema.SingleNestedAttribute{MarkdownDescription: "Per-rule Variable Offset Load Balancing config for G-seres devices", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"end_delim": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "hash": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.OneOf("checksum", "xor", "crc")}}, "start_delim": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "start_field": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}}}}}, "gseries_pattern_match": schema.SingleNestedAttribute{MarkdownDescription: "Only applicable for G-series per-rule GSOP", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"fixed_offset": schema.SingleNestedAttribute{MarkdownDescription: "Per-rule Fixed Offset Pattern Match config for G-seres devices", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"length": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.AtLeast(1)}}, "offset": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.AtLeast(0)}}}}, "variable_offset": schema.SingleNestedAttribute{MarkdownDescription: "Per-rule Variable Offset Pattern Match config for G-seres devices", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"end_delim": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "start_delim": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}}}}}, "gtp_whitelist": schema.SingleNestedAttribute{MarkdownDescription: "Only applicable for H-series", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("enabled")}}}}, "header_add": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Add Header' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"vlan": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.Between(0, 4095)}}}}, "header_remove": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Remove Header' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"ah1": schema.StringAttribute{MarkdownDescription: "only valid and required for 'generic'. First anchor header, from which the header to be stripped is occurred.", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("none", "eth", "vlan", "mpls", "ipv4", "ipv6")}}, "ah2": schema.StringAttribute{MarkdownDescription: "only valid and required for 'generic'. next anchor header.", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("none", "eth", "vlan", "mpls", "ipv4", "ipv6", "tcp", "udp", "any")}}, "custom_len": schema.Int64Attribute{MarkdownDescription: "only valid for 'generic'. length of unknown header.", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 1500)}}, "erspan_flow_id": schema.Int64Attribute{MarkdownDescription: "only applicable for 'erspan', specifies the flow id to strip. Value of 0 will strip all flow ids", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 1023)}}, "fp_dst_switch_id": schema.Int64Attribute{MarkdownDescription: "Only valid and required for 'fabricPath', 12 bit destination switch id", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 4095)}}, "fp_src_switch_id": schema.Int64Attribute{MarkdownDescription: "Only valid and required for 'fabricPath', 12 bit source switch id", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 4095)}}, "header_count": schema.Int64Attribute{MarkdownDescription: "only valid for 'generic'. Number of headers to be stripped.", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 32)}}, "offset": schema.StringAttribute{MarkdownDescription: "only valid and required for 'generic'.'start': strip from start of ah1, 'end': strip from end of ah1 or any other offset in the range of length of ah1, 'offsetRangeValue' is required if offset is 'offsetRange'.", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("start", "end", "offsetRange")}}, "offset_range_value": schema.Int64Attribute{MarkdownDescription: "only valid and required when offset is 'offsetRange', integer within range of size of header", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 1500)}}, "protocol": schema.StringAttribute{MarkdownDescription: "'gre' and 'fabricPath' are only applicable for H-series", Required: true, Validators: []validator.String{stringvalidator.OneOf("gtp", "isl", "mpls", "mplsPlusVlan", "vlan", "vntag", "vxlan", "gre", "fabricPath", "fm6000Ts", "erspan", "generic")}}, "timestamp_format": schema.StringAttribute{MarkdownDescription: "Timestamp format. Only valid and required for 'fm6000Ts'", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("gigasmart", "x12Ts", "none")}}, "vlan_header": schema.StringAttribute{MarkdownDescription: "Only valid when protocol is 'vlan'. Specifies the target vlan header to strip. Defaults to 'outer'", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("all", "outer")}}, "vxlan_id": schema.Int64Attribute{MarkdownDescription: "24-bit value. Only valid when protocol is 'vxlan'. Specifies the vxlan id to strip. Value of '0' will strip all vxlan ids", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 16777215)}}}}, "icap": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART ICAP Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"icap_profile": schema.StringAttribute{MarkdownDescription: "Alias of referenced ICAP Profile", Required: true}}}, "inline_ssl": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART Inline SSL Profile Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"inline_ssl_profile": schema.StringAttribute{MarkdownDescription: "Alias of referenced Inline SSL Profile", Required: true}}}, "load_balance": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Load Balancing' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enhanced": schema.SingleNestedAttribute{MarkdownDescription: "Private class. Enhanced part of the GigaSMART 'Load Balancing' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"elb_alias": schema.StringAttribute{MarkdownDescription: "elb app alias", Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}}}, "stateful": schema.SingleNestedAttribute{MarkdownDescription: "Private class. Stateful part of the GigaSMART 'Load Balancing' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"app_type": schema.StringAttribute{MarkdownDescription: "'sapf' option is added in release H 4.3.01. 'sip' and 'tunnel' added in release H 5.1. 'diameter' is added in release H 5.5", Required: true, Validators: []validator.String{stringvalidator.OneOf("gtp", "sapf", "sip", "tunnel", "diameter")}}, "diameter_key_hash_type": schema.StringAttribute{MarkdownDescription: "required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyHash'", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("sessionId", "userName", "commandCode")}}, "diameter_key_multi_hash_type": schema.ListNestedAttribute{MarkdownDescription: "required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyMultiHash'", Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"avp_codevalue": schema.Int64Attribute{MarkdownDescription: "required when 'key' == 'avpCode'", Optional: true, Computed: true}, "key": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.OneOf("sessionId", "userName", "avpCode", "commandCode", "endToEnd", "hopByHop", "applicationId")}}}}}, "gtp_key_hash_type": schema.StringAttribute{MarkdownDescription: "required when 'appType' == 'gtp' and 'lbType' == 'gtpKeyHash'. ignored otherwise", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("imsi", "imei", "msisdn")}}, "lb_type": schema.StringAttribute{MarkdownDescription: "'gtpKeyHash' is only applicable when 'appType' == 'gtp'. 'sipKeyHash' is only applicable when 'appType' == 'sip'", Required: true, Validators: []validator.String{stringvalidator.OneOf("leastBw", "leastPktRate", "leastConn", "leastTotalTraffic", "roundRobin", "wtLeastBw", "wtLeastPktRate", "wtLeastConn", "wtLeastTotalTraffic", "wtRoundRobin", "wtImsi", "wtSupi", "gtpKeyHash", "sipKeyHash", "diameterKeyHash", "diameterKeyMultiHash")}}, "sip_key_hash_type": schema.StringAttribute{MarkdownDescription: "required when 'appType' == 'sip' and 'lbType' == 'sipKeyHash'. ignored otherwise", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("callerId")}}}}, "stateless": schema.SingleNestedAttribute{MarkdownDescription: "Private class. Stateless part of the GigaSMART 'Load Balancing' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"field_location": schema.StringAttribute{MarkdownDescription: "Ignored when 'hashFields' == 'gtpuTeid'. Required otherwise", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("inner", "outer")}}, "hash_fields": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.OneOf("ipOnly", "ipAndPort", "fiveTuple", "gtpuTeid")}}}}}}, "masking": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Masking' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"content_type": schema.StringAttribute{MarkdownDescription: "content type that will trigger masking, only valid and required for protocol 'sip'", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("message_cpim")}}, "length": schema.Int64Attribute{MarkdownDescription: "max length is 9600 on H-series; 9000 on G-series, not applicable for protocol 'sip' valid and required otherwise", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 9600)}}, "offset": schema.Int64Attribute{MarkdownDescription: "min offset is 0 for 'none'; 1 for protocol-referenced offsets, not applicable for protocol 'sip' valid and required otherwise", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 9000)}}, "pattern": schema.StringAttribute{MarkdownDescription: "1-byte hex mask, not applicable for protocol 'sip' valid and required otherwise", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^(0x)?[0-9a-fA-F]{1,2}$"), "value must match pattern \"^(0x)?[0-9a-fA-F]{1,2}$\"")}}, "protocol": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.OneOf("none", "ipv4", "ipv6", "udp", "tcp", "ftp-data", "https", "ssh", "gtp", "gtp-ipv4", "gtp-udp", "gtp-tcp", "sip")}}}}, "metadata_export": schema.SingleNestedAttribute{MarkdownDescription: "Only applicable for H-series", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"cache": schema.StringAttribute{MarkdownDescription: "metadata cache alias. cache should have exporters defined, and the exporters should have applicationProfiles defined", Optional: true, Computed: true}}}, "netflow": schema.SingleNestedAttribute{MarkdownDescription: "Only applicable for H-series", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("enabled")}}}}, "sa_apf": schema.SingleNestedAttribute{MarkdownDescription: "Deprecated (functionality will not be supported from 5.4 release onwards); Only applicable for H-series", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("enabled")}}}}, "sip_whitelist": schema.SingleNestedAttribute{MarkdownDescription: "Only applicable for H-series", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("enabled")}}}}, "slicing": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Slicing' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enhanced": schema.StringAttribute{MarkdownDescription: "enhanced-slicing apps alias", Optional: true, Computed: true}, "offset": schema.Int64Attribute{MarkdownDescription: "min offset is 64 for 'none'; 4 for protocol-referenced offsets. required property till H 5.6", Required: true, Validators: []validator.Int64{int64validator.Between(4, 9000)}}, "protocol": schema.StringAttribute{MarkdownDescription: "required property till H 5.6", Required: true, Validators: []validator.String{stringvalidator.OneOf("none", "ipv4", "ipv6", "udp", "tcp", "ftp-data", "https", "ssh", "gtp", "gtp-ipv4", "gtp-udp", "gtp-tcp")}}}}, "ssl_decrypt": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'SSL Decrypt' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"in_port": schema.Int64Attribute{MarkdownDescription: "Port number of 0 represents 'any' port", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 65535)}}, "out_port": schema.Int64Attribute{MarkdownDescription: "Port number of 0 represents 'auto' port", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 65535)}}}}, "trailer_add": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Add Trailer' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"types": schema.SetAttribute{MarkdownDescription: "'crc' is not applicable for G-series", Required: true, ElementType: types.StringType, Validators: []validator.Set{setvalidator.ValueStringsAre(stringvalidator.OneOf("crc", "srcid"))}}}}, "trailer_remove": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("enabled")}}}}, "tunnel_decap": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Decapsulate Tunnel' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"custom": schema.SingleNestedAttribute{MarkdownDescription: "only applicable for 'custom', in which case it is required.", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"port_dst": schema.Int64Attribute{MarkdownDescription: "when specified as 0, no validation will be done in the packet.", Required: true, Validators: []validator.Int64{int64validator.Between(0, 65535)}}, "port_src": schema.Int64Attribute{MarkdownDescription: "when specified as 0, no validation will be done in the packet.", Required: true, Validators: []validator.Int64{int64validator.Between(0, 65535)}}}}, "erspan_flow_id": schema.Int64Attribute{MarkdownDescription: "only applicable for 'erspan', A Flow ID of 0 decapsulates all ERSPAN tunnel traffic regardless of Flow ID", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 1023)}}, "gmip_port": schema.Int64Attribute{MarkdownDescription: "only applicable for 'gmip', in which case it is required. Specifies the UDP port on which the Tunnel Network port on the receiving GigaVUE H Series is listening. Must match the configuration of the portdst configured on the sending end of the tunnel", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 65535)}}, "l2_gre_key": schema.Int64Attribute{MarkdownDescription: "only applicable for 'l2gre', in which case it is required.", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 4294967295)}}, "tls_pcapng": schema.SingleNestedAttribute{MarkdownDescription: "only applicable for 'tls-pcapng', in which case it is required", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"decap_key": schema.StringAttribute{Optional: true, Computed: true}, "listener": schema.StringAttribute{Optional: true, Computed: true}}}, "type": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.OneOf("gmip", "erspan", "l2gre", "custom", "vxlan", "tls-pcapng")}}, "vxlan": schema.SingleNestedAttribute{MarkdownDescription: "only applicable for 'vxlan', in which case it is required.", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"port_dst": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.Between(1, 65535)}}, "port_src": schema.Int64Attribute{MarkdownDescription: "when specified as 0, no validation will be done in the packet.", Required: true, Validators: []validator.Int64{int64validator.Between(0, 65535)}}, "vni": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.Between(0, 16777215)}}}}}}, "tunnel_encap": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Encapsulate Tunnel' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"gmip_config": schema.SingleNestedAttribute{MarkdownDescription: "Configuration for GMIP Tunnel Encapsulate GigaSmaprt App", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"dscp": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 63)}}, "dst_ip": schema.StringAttribute{MarkdownDescription: "IP Destination. IPv4 or IPv6.", Required: true}, "dst_port": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.Between(0, 65535)}}, "flow_label": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 1048575)}}, "prec": schema.Int64Attribute{MarkdownDescription: "decimal Precedence value from 0-7 to be used in the ToS byte of the outer headers of tunneled packets", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 7)}}, "src_port": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.Between(0, 65535)}}, "ttl": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 255)}}}}, "l2_gre_config": schema.SingleNestedAttribute{MarkdownDescription: "Configuration for l2GRE Tunnel Encapsulate GigaSmaprt App", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"dscp": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 63)}}, "dst_ip": schema.StringAttribute{MarkdownDescription: "ip destination, mutually exclusive with 'pgDst'. IPv4 or IPv6.", Optional: true, Computed: true}, "flow_label": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 1048575)}}, "key": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.Between(0, 4294967295)}}, "pg_dst": schema.StringAttribute{MarkdownDescription: "port group destination alias, mutually exclusive with 'dstIp'", Optional: true, Computed: true}, "prec": schema.Int64Attribute{MarkdownDescription: "decimal Precedence value from 0-7 to be used in the ToS byte of the outer headers of tunneled packets", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 7)}}, "session_field": schema.StringAttribute{MarkdownDescription: "required with stateful loadBalance when 'appType' is 'tunnel'", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("fiveTupleIpv4", "fiveTupleIpv6", "fiveTupleAny", "threeTupleIpv4", "threeTupleIpv6", "threeTupleAny", "ipv4Only", "ipv6Only", "ipAny")}}, "session_pos": schema.StringAttribute{MarkdownDescription: "required if 'sessionField' is specified", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("inner", "outer")}}, "ttl": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 255)}}}}, "tls_pcapng": schema.SingleNestedAttribute{MarkdownDescription: "only applicable for 'tls-pcapng', in which case it is required", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"exporter": schema.StringAttribute{Optional: true, Computed: true}, "exporter_group": schema.StringAttribute{Optional: true, Computed: true}}}, "type": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.OneOf("gmip", "l2gre", "vxlan", "tls-pcapng")}}, "vxlan_config": schema.SingleNestedAttribute{MarkdownDescription: "Configuration for Vxlan Tunnel Encapsulate GigaSmaprt App", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"dscp": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 63)}}, "dst_ip": schema.StringAttribute{MarkdownDescription: "IP Destination. IPv4 or IPv6.", Required: true}, "dst_port": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.OneOf(4789, 8472, 48879)}}, "src_port": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.Between(0, 65535)}}, "ttl": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 255)}}, "vni": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.Between(1, 16777215)}}}}}}}}, "export_configs": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"exporter_alias": schema.StringAttribute{Optional: true, Computed: true}, "remote_application_port": schema.Int64Attribute{Optional: true, Computed: true}, "remote_ip": schema.StringAttribute{Optional: true, Computed: true}, "source_application_port": schema.Int64Attribute{Optional: true, Computed: true}}}}, "exporter_group_alias": schema.StringAttribute{Optional: true, Computed: true}, "gsop_alias": schema.StringAttribute{Optional: true, Computed: true}, "map_alias": schema.StringAttribute{Optional: true, Computed: true}, "rules": schema.SingleNestedAttribute{MarkdownDescription: "Map Rules Container. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"drop_rules": schema.DynamicAttribute{Optional: true, Computed: true}, "pass_rules": schema.DynamicAttribute{Optional: true, Computed: true}}}, "source_port": schema.ListAttribute{MarkdownDescription: "source port for encap traffic , it would be on the same cluster as GsEngine", Optional: true, Computed: true, ElementType: types.StringType}}}, "gsgroup": schema.StringAttribute{Optional: true, Computed: true}, "ip_interface_in": schema.StringAttribute{MarkdownDescription: "ipInterface for Decap source", Optional: true, Computed: true}, "ip_interface_out": schema.StringAttribute{MarkdownDescription: "ipInterface for Encap destination", Optional: true, Computed: true}, "local_key_alias": schema.StringAttribute{Optional: true, Computed: true}, "remote_key_alias": schema.StringAttribute{Optional: true, Computed: true}, "ssl_profile": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"cipher": schema.StringAttribute{Optional: true, Computed: true}, "mtls": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("enable", "disable")}}, "version": schema.StringAttribute{Optional: true, Computed: true}}}, "ssl_profile_alias": schema.StringAttribute{Optional: true, Computed: true}, "tcp_profile": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"keep_alive_timer": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(30, 7200)}}, "selective_ack": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("enable", "disable")}}, "syn_retries": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 6)}}}}, "tcp_profile_alias": schema.StringAttribute{Optional: true, Computed: true}, "traffic_dir": schema.StringAttribute{MarkdownDescription: "IN for Decap , OUT for Encap", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("IN", "OUT")}}, "ttl": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 255)}}, "tunnel_type": schema.StringAttribute{Optional: true, Computed: true}}, Blocks: map[string]schema.Block{"timeouts": schema.SingleNestedBlock{Attributes: map[string]schema.Attribute{"create": schema.Int64Attribute{MarkdownDescription: "Maximum time in seconds for the create operation. Defaults to 1200.", Optional: true}, "read": schema.Int64Attribute{MarkdownDescription: "Maximum time in seconds for the read operation. Defaults to 1200.", Optional: true}, "update": schema.Int64Attribute{MarkdownDescription: "Maximum time in seconds for the update operation. Defaults to 1200.", Optional: true}, "delete": schema.Int64Attribute{MarkdownDescription: "Maximum time in seconds for the delete operation. Defaults to 1200.", Optional: true}}, MarkdownDescription: "Per-operation timeouts in seconds."}}}
}

// IdentitySchema returns the resource identity schema shared with the paired list resource.
func (r *TunnelApplicationResource) IdentitySchema(_ context.Context, _ resource.IdentitySchemaRequest, resp *resource.IdentitySchemaResponse) {
	resp.IdentitySchema = identityschema.Schema{Attributes: map[string]identityschema.Attribute{"alias": identityschema.StringAttribute{RequiredForImport: true}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *TunnelApplicationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan TunnelApplicationResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	timeout := 20 * time.Minute
	if plan.Timeouts != nil && !plan.Timeouts.Create.IsNull() && !plan.Timeouts.Create.IsUnknown() {
		timeout = time.Duration(plan.Timeouts.Create.ValueInt64()) * time.Second
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	r.createRemote(ctx, &plan, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("alias"), plan.Alias)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// createRemote performs the create HTTP exchange and decodes the response into plan. Extracted from Create so the request/response logic is unit-testable without a tfsdk.Plan.
func (r *TunnelApplicationResource) createRemote(ctx context.Context, plan *TunnelApplicationResourceModel, resp *resource.CreateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_tunnel_application", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_tunnel_application", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/apps/tunnelApplication"
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_tunnel_application", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", plan.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_tunnel_application", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error creating gigavuecore_tunnel_application", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error creating gigavuecore_tunnel_application", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error creating gigavuecore_tunnel_application", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error creating gigavuecore_tunnel_application", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error creating gigavuecore_tunnel_application", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error creating gigavuecore_tunnel_application", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error creating gigavuecore_tunnel_application", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error creating gigavuecore_tunnel_application", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error creating gigavuecore_tunnel_application", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error creating gigavuecore_tunnel_application", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_tunnel_application", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
	if plan.Alias.IsNull() || plan.Alias.IsUnknown() {
		loc := httpResp.Header.Get("Location")
		if loc != "" {
			loc = strings.TrimRight(loc, "/")
			i := strings.LastIndex(loc, "/")
			if i >= 0 {
				loc = loc[i+1:]
			}
			plan.Alias = types.StringValue(loc)
		} else {
			resp.Diagnostics.AddError("Error creating gigavuecore_tunnel_application", "The create response did not contain an identifier and no Location header was returned, so the resource cannot be tracked in state.")
			return
		}
	}
}

// Read refreshes the Terraform state with the latest remote values.
func (r *TunnelApplicationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state TunnelApplicationResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	timeout := 20 * time.Minute
	if state.Timeouts != nil && !state.Timeouts.Read.IsNull() && !state.Timeouts.Read.IsUnknown() {
		timeout = time.Duration(state.Timeouts.Read.ValueInt64()) * time.Second
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	if r.readRemote(ctx, &state, resp) {
		resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("alias"), state.Alias)...)
		resp.State.RemoveResource(ctx)
		return
	}
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("alias"), state.Alias)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// readRemote performs the read HTTP exchange and decodes the response into state, returning removed=true when the API reports 404. Extracted from Read so the request/response logic is unit-testable without a tfsdk.State.
func (r *TunnelApplicationResource) readRemote(ctx context.Context, state *TunnelApplicationResourceModel, resp *resource.ReadResponse) (removed bool) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/tunnelApplication/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(state.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_tunnel_application", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", state.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_tunnel_application", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		removed = true
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_tunnel_application", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_tunnel_application", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_tunnel_application", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_tunnel_application", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_tunnel_application", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_tunnel_application", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_tunnel_application", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_tunnel_application", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_tunnel_application", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_tunnel_application", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["tunnelApplication"]; ok {
		if m, ok := v.(map[string]any); ok {
			data = m
		} else if arr, ok := v.([]any); ok && len(arr) > 0 {
			if m, ok := arr[0].(map[string]any); ok {
				data = m
			}
		}
	}
	err = applyJSONToModel(&state, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_tunnel_application", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
	return
}

// Update modifies the remote resource to match the desired plan.
func (r *TunnelApplicationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan TunnelApplicationResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state TunnelApplicationResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	timeout := 20 * time.Minute
	if plan.Timeouts != nil && !plan.Timeouts.Update.IsNull() && !plan.Timeouts.Update.IsUnknown() {
		timeout = time.Duration(plan.Timeouts.Update.ValueInt64()) * time.Second
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	if plan.Alias.IsNull() || plan.Alias.IsUnknown() {
		if !state.Alias.IsNull() && !state.Alias.IsUnknown() {
			plan.Alias = state.Alias
		}
	}
	preserveStateIntoPlan(&plan, &state)
	r.updateRemote(ctx, &plan, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("alias"), plan.Alias)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// updateRemote performs the update HTTP exchange and decodes the response into plan. Extracted from Update so the request/response logic is unit-testable without a tfsdk.Plan.
func (r *TunnelApplicationResource) updateRemote(ctx context.Context, plan *TunnelApplicationResourceModel, resp *resource.UpdateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_tunnel_application", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_tunnel_application", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/apps/tunnelApplication/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(plan.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodPatch, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_tunnel_application", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", plan.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_tunnel_application", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error updating gigavuecore_tunnel_application", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error updating gigavuecore_tunnel_application", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error updating gigavuecore_tunnel_application", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error updating gigavuecore_tunnel_application", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error updating gigavuecore_tunnel_application", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error updating gigavuecore_tunnel_application", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error updating gigavuecore_tunnel_application", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error updating gigavuecore_tunnel_application", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error updating gigavuecore_tunnel_application", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error updating gigavuecore_tunnel_application", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_tunnel_application", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Delete destroys the remote resource.
func (r *TunnelApplicationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state TunnelApplicationResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	timeout := 20 * time.Minute
	if state.Timeouts != nil && !state.Timeouts.Delete.IsNull() && !state.Timeouts.Delete.IsUnknown() {
		timeout = time.Duration(state.Timeouts.Delete.ValueInt64()) * time.Second
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	r.deleteRemote(ctx, &state, resp)
}

// deleteRemote performs the delete HTTP exchange, treating a 404 as already deleted. Extracted from Delete so the request/response logic is unit-testable without a tfsdk.State.
func (r *TunnelApplicationResource) deleteRemote(ctx context.Context, state *TunnelApplicationResourceModel, resp *resource.DeleteResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/tunnelApplication/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(state.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_tunnel_application", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", state.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_tunnel_application", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		return
	}
	if !(httpResp.StatusCode == 204) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error deleting gigavuecore_tunnel_application", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error deleting gigavuecore_tunnel_application", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error deleting gigavuecore_tunnel_application", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error deleting gigavuecore_tunnel_application", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error deleting gigavuecore_tunnel_application", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error deleting gigavuecore_tunnel_application", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error deleting gigavuecore_tunnel_application", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error deleting gigavuecore_tunnel_application", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error deleting gigavuecore_tunnel_application", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *TunnelApplicationResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData))
		return
	}
	r.client = c
}

// ImportState imports an existing remote resource into Terraform state.
func (r *TunnelApplicationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	tunnelApplicationImportIDParts := strings.Split(req.ID, "/")
	if len(tunnelApplicationImportIDParts) != 2 {
		resp.Diagnostics.AddError("Unexpected Import Identifier", fmt.Sprintf("Expected import identifier with format \"{alias}/{cluster_id}\". Got %q.", req.ID))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("alias"), tunnelApplicationImportIDParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("cluster_id"), tunnelApplicationImportIDParts[1])...)
}
