package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	timeouts "github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	int64validator "github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	listvalidator "github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	setvalidator "github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	stringvalidator "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	path "github.com/hashicorp/terraform-plugin-framework/path"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	validator "github.com/hashicorp/terraform-plugin-framework/schema/validator"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertions.
var (
	_ resource.Resource                = (*FmTemplateResource)(nil)
	_ resource.ResourceWithImportState = (*FmTemplateResource)(nil)
	_ resource.ResourceWithConfigure   = (*FmTemplateResource)(nil)
)

// FmTemplateResource is the generated Terraform managed resource implementation.
type FmTemplateResource struct {
	client *client.Client
}

// FmTemplateResourceModel describes the Terraform state and plan shape for FmTemplateResource.
type FmTemplateResourceModel struct {
	Config           types.Object   `tfsdk:"config"`
	ConfigLevel      types.String   `tfsdk:"config_level" json:"configLevel"`
	ConfigLevelValue types.List     `tfsdk:"config_level_value" json:"configLevelValue"`
	ConfigResource   types.Dynamic  `tfsdk:"config_resource" json:"configResource"`
	ConfigType       types.String   `tfsdk:"config_type" json:"configType"`
	Modifiable       types.Bool     `tfsdk:"modifiable"`
	RefCount         types.Int64    `tfsdk:"ref_count" json:"refCount"`
	RefObject        types.Object   `tfsdk:"ref_object" json:"refObject"`
	TemplateName     types.String   `tfsdk:"template_name" json:"templateName"`
	UpdateTime       types.String   `tfsdk:"update_time" json:"updateTime"`
	Timeouts         timeouts.Value `tfsdk:"timeouts"`
}

// Metadata returns the resource type name.
func (r *FmTemplateResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_fm_template"
}

// Schema returns the Terraform schema for this resource.
func (r *FmTemplateResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "new in FM 5.7", Attributes: map[string]schema.Attribute{"config": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"aaa_auth_config": schema.SingleNestedAttribute{MarkdownDescription: "FM Global AAA Authentication config to the device's", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"auth_sequence": schema.SetAttribute{MarkdownDescription: "Authentication methods order for user login. Valid values are 'local', 'ldap', 'radius', 'tacacs'", Required: true, ElementType: types.StringType, Validators: []validator.Set{setvalidator.SizeAtMost(4), setvalidator.ValueStringsAre(stringvalidator.OneOf("local", "ldap", "radius", "tacacs"))}}, "external_login_mapping": schema.SingleNestedAttribute{MarkdownDescription: "Specifies how externally authenticated logins (RADIUS, TACACS+, or LDAP) are mapped to local accounts. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"default_local_user": schema.StringAttribute{MarkdownDescription: "Specifies the account to which externally authenticated logins are mapped when map order is set to remote-first (if there is no matching local account) or local-only", Optional: true, Computed: true}, "user_map_order": schema.StringAttribute{MarkdownDescription: "Specifies how externally authenticated logins (RADIUS, TACACS+, or LDAP) are mapped to local accounts", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("localOnly", "remoteFirst", "remoteOnly")}}}}}}, "acme_certificate": schema.ListNestedAttribute{MarkdownDescription: "Available when ConfigType is ACME_TEMPLATE", Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"acme_server_url": schema.StringAttribute{Required: true}, "algorithm": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("rsa-2048", "rsa-4096", "ec-prime256v1", "ec-secp384r1")}}, "operation_type": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.OneOf("issue", "renew", "revoke", "clear")}}, "renew_days": schema.Int64Attribute{MarkdownDescription: "default will be 1/3rd of certificate validity period", Optional: true, Computed: true}}}}, "device_ssl_certificate_configs": schema.ListNestedAttribute{MarkdownDescription: "Available when ConfigType is SSL_CERTIFICATE_TEMPLATE", Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"issuer": schema.StringAttribute{MarkdownDescription: "issuer details of the certificate", Optional: true, Computed: true}, "not_after": schema.StringAttribute{MarkdownDescription: "date and time when certificate stops being valid ([rfc3339](https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14))", Optional: true, Computed: true}, "not_before": schema.StringAttribute{MarkdownDescription: "date and time when certificate starts being valid ([rfc3339](https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14))", Optional: true, Computed: true}, "operation_type": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.OneOf("add", "delete")}}, "signature_algorithm": schema.StringAttribute{Optional: true, Computed: true}, "subject": schema.StringAttribute{MarkdownDescription: "subject name of the certificate", Optional: true, Computed: true}, "trusted_ca": schema.SingleNestedAttribute{Required: true, Attributes: map[string]schema.Attribute{"name": schema.StringAttribute{MarkdownDescription: "name of the certificate", Required: true}}}, "upload_spec": schema.SingleNestedAttribute{Required: true, Attributes: map[string]schema.Attribute{"info": schema.SingleNestedAttribute{MarkdownDescription: "Certificate info", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"comment": schema.StringAttribute{MarkdownDescription: "a short description of the certificate", Optional: true, Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "name of the certificate", Required: true}, "passphrase": schema.StringAttribute{MarkdownDescription: "used to decrypt pkcs12 and private keys", Optional: true, Computed: true}, "type": schema.StringAttribute{MarkdownDescription: "type of the certificate", Required: true, Validators: []validator.String{stringvalidator.OneOf("privateKey", "certificate", "pkcs12")}}}}, "pem": schema.StringAttribute{MarkdownDescription: "contents of the certificate in pem format", Optional: true, Computed: true}}}}}}, "export_metadata_app_profile": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "application profile alias", Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "application_id": schema.BoolAttribute{MarkdownDescription: "only valid with 'export' type", Optional: true, Computed: true}, "applications": schema.ListNestedAttribute{MarkdownDescription: "application and attributes.", Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"attributes": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"name": schema.StringAttribute{MarkdownDescription: "attribute name", Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "value": schema.StringAttribute{MarkdownDescription: "application's attribute value", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}}}}, "is_user_defined": schema.BoolAttribute{MarkdownDescription: "Default value is false. Must be set to true only for App Intel solution while configuring user defined apps", Optional: true, Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "application name", Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}}}}, "counter": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"bytes": schema.BoolAttribute{Optional: true, Computed: true}, "bytes_long": schema.BoolAttribute{Optional: true, Computed: true}, "inner_byte": schema.BoolAttribute{Optional: true, Computed: true}, "inner_byte_long": schema.BoolAttribute{Optional: true, Computed: true}, "packets": schema.BoolAttribute{Optional: true, Computed: true}, "packets_long": schema.BoolAttribute{Optional: true, Computed: true}}}, "datalink": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"mac_dst": schema.BoolAttribute{Optional: true, Computed: true}, "mac_src": schema.BoolAttribute{Optional: true, Computed: true}, "vlan": schema.BoolAttribute{Optional: true, Computed: true}}}, "description": schema.StringAttribute{Optional: true, Computed: true}, "flow": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"end_reason": schema.BoolAttribute{Optional: true, Computed: true}}}, "gtpu": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"qfi": schema.BoolAttribute{Optional: true, Computed: true}, "teid": schema.BoolAttribute{Optional: true, Computed: true}}}, "interface": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"in_name_width": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 32)}}, "in_physical_width": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.OneOf(2, 4)}}, "out_physical_width": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.OneOf(2, 4)}}}}, "ip": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"version": schema.BoolAttribute{Optional: true, Computed: true}}}, "ipv4": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"destination": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"prefix_min_mask": schema.StringAttribute{Optional: true, Computed: true}}}, "dscp": schema.BoolAttribute{Optional: true, Computed: true}, "fragmentation": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"flags": schema.BoolAttribute{Optional: true, Computed: true}, "offset": schema.BoolAttribute{Optional: true, Computed: true}}}, "header_len": schema.BoolAttribute{Optional: true, Computed: true}, "option_map": schema.BoolAttribute{Optional: true, Computed: true}, "precedence": schema.BoolAttribute{Optional: true, Computed: true}, "protocol": schema.BoolAttribute{Optional: true, Computed: true}, "section": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"header_size": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 128)}}, "payload_size": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 128)}}}}, "source": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"prefix_min_mask": schema.StringAttribute{MarkdownDescription: "ipv4 source prefix minimum-mask - netmask or mask length", Optional: true, Computed: true}}}, "tos": schema.BoolAttribute{Optional: true, Computed: true}, "total_length": schema.BoolAttribute{Optional: true, Computed: true}, "ttl": schema.BoolAttribute{Optional: true, Computed: true}}}, "ipv6": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"destination": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"prefix_min_mask": schema.StringAttribute{Optional: true, Computed: true}}}, "dscp": schema.BoolAttribute{Optional: true, Computed: true}, "extension_map": schema.BoolAttribute{Optional: true, Computed: true}, "flow_label": schema.BoolAttribute{Optional: true, Computed: true}, "fragmentation": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"flags": schema.BoolAttribute{Optional: true, Computed: true}, "offset": schema.BoolAttribute{Optional: true, Computed: true}}}, "hop_limit": schema.BoolAttribute{Optional: true, Computed: true}, "length": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"header": schema.BoolAttribute{Optional: true, Computed: true}, "payload": schema.BoolAttribute{Optional: true, Computed: true}, "total": schema.BoolAttribute{Optional: true, Computed: true}}}, "next_header": schema.BoolAttribute{Optional: true, Computed: true}, "precedence": schema.BoolAttribute{Optional: true, Computed: true}, "section": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"header_size": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 128)}}, "payload_size": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 128)}}}}, "source": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"prefix_min_mask": schema.StringAttribute{Optional: true, Computed: true}}}, "traffic_class": schema.BoolAttribute{Optional: true, Computed: true}}}, "outer_ipv4": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"destination": schema.BoolAttribute{Optional: true, Computed: true}, "source": schema.BoolAttribute{Optional: true, Computed: true}}}, "outer_ipv6": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"destination": schema.BoolAttribute{Optional: true, Computed: true}, "source": schema.BoolAttribute{Optional: true, Computed: true}}}, "timestamp": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"flow_end_msec": schema.BoolAttribute{Optional: true, Computed: true}, "flow_endsec": schema.BoolAttribute{Optional: true, Computed: true}, "flow_start_msec": schema.BoolAttribute{Optional: true, Computed: true}, "flow_startsec": schema.BoolAttribute{Optional: true, Computed: true}, "sys_up_time_first": schema.BoolAttribute{Optional: true, Computed: true}, "sys_up_time_last": schema.BoolAttribute{Optional: true, Computed: true}}}, "transport": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"dst_port": schema.BoolAttribute{Optional: true, Computed: true}, "icmp": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"ipv4_code": schema.BoolAttribute{Optional: true, Computed: true}, "ipv4_type": schema.BoolAttribute{Optional: true, Computed: true}, "ipv6_code": schema.BoolAttribute{Optional: true, Computed: true}, "ipv6_type": schema.BoolAttribute{Optional: true, Computed: true}}}, "src_port": schema.BoolAttribute{Optional: true, Computed: true}, "tcp": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"ack_number": schema.BoolAttribute{Optional: true, Computed: true}, "dst_port": schema.BoolAttribute{Optional: true, Computed: true}, "flags": schema.BoolAttribute{Optional: true, Computed: true}, "header_len": schema.BoolAttribute{Optional: true, Computed: true}, "seq_number": schema.BoolAttribute{Optional: true, Computed: true}, "src_port": schema.BoolAttribute{Optional: true, Computed: true}, "urgent_ptr": schema.BoolAttribute{Optional: true, Computed: true}, "window_size": schema.BoolAttribute{Optional: true, Computed: true}}}, "udp": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"dst_port": schema.BoolAttribute{Optional: true, Computed: true}, "msg_len": schema.BoolAttribute{Optional: true, Computed: true}, "src_port": schema.BoolAttribute{Optional: true, Computed: true}}}}}, "type": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("export", "filter")}}}}, "giga_port_neighbors_discovery_config": schema.SingleNestedAttribute{MarkdownDescription: "To contain the port neighbour discovery template config", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"resource_configs": schema.DynamicAttribute{MarkdownDescription: "List of config for each port type", Optional: true, Computed: true}}}, "giga_stream_threshold_config": schema.SingleNestedAttribute{MarkdownDescription: "Gigastream threshold configs", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"giga_stream_type_thresholds": schema.ListNestedAttribute{MarkdownDescription: "GigaStream threshold config for port types", Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"type": schema.StringAttribute{MarkdownDescription: "Port type", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("all", "tool", "stack", "hybrid", "circuit")}}, "variance_threshold": schema.Float64Attribute{MarkdownDescription: "Threshold value in percentage. '-1' will disable threshold feature.", Optional: true, Computed: true}}}}}}, "giga_user_defined_application_config": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "pattern: ^[a-zA-Z0-9&'+_=|\\-]+$", Optional: true, Computed: true}, "app_id": schema.Int64Attribute{MarkdownDescription: "minimum: 16777216\nmaximum: 16777343 \n Ranges from 2^24 to 2^24 + 127. It needs to be unique", Optional: true, Computed: true}, "priority": schema.Int64Attribute{MarkdownDescription: "minimum: 1\nmaximum: 120 \nLower the priority higher the precedence", Optional: true, Computed: true}, "rules": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"rule": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"address": schema.StringAttribute{Optional: true, Computed: true}, "code": schema.StringAttribute{Optional: true, Computed: true}, "common_name": schema.StringAttribute{Optional: true, Computed: true}, "content": schema.StringAttribute{Optional: true, Computed: true}, "cts_cookie": schema.StringAttribute{Optional: true, Computed: true}, "cts_page_url": schema.StringAttribute{Optional: true, Computed: true}, "cts_referer": schema.StringAttribute{Optional: true, Computed: true}, "cts_server": schema.StringAttribute{Optional: true, Computed: true}, "cts_uri": schema.StringAttribute{Optional: true, Computed: true}, "cts_user_agent": schema.StringAttribute{Optional: true, Computed: true}, "dscp": schema.StringAttribute{Optional: true, Computed: true}, "mime_type": schema.StringAttribute{Optional: true, Computed: true}, "mindata": schema.Int64Attribute{Optional: true, Computed: true}, "port": schema.StringAttribute{Optional: true, Computed: true}, "resolv_name": schema.StringAttribute{Optional: true, Computed: true}, "stc_location": schema.StringAttribute{Optional: true, Computed: true}, "stc_server_agent": schema.StringAttribute{Optional: true, Computed: true}, "stc_subject_alt_name": schema.StringAttribute{Optional: true, Computed: true}, "stream": schema.StringAttribute{Optional: true, Computed: true}, "typeval": schema.StringAttribute{Optional: true, Computed: true}, "user_agent": schema.StringAttribute{Optional: true, Computed: true}}}}}}}}, "ldap_servers": schema.ListNestedAttribute{MarkdownDescription: "List of LDAP Servers. Available for ConfigType LDAP_SERVERS_TEMPLATE ChildType ldapServers", Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"order": schema.StringAttribute{MarkdownDescription: "The order in which the server is to be reached. 1 means server will be contacted first", Optional: true, Computed: true}, "server_address": schema.StringAttribute{MarkdownDescription: "ipv4 or ipv6 or hostname. Specifies address of the LDAP server where authentication requests will be sent", Required: true}}}}, "metadata_exporter": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "application_profiles": schema.ListAttribute{MarkdownDescription: "application profile aliases to attach to the exporter", Optional: true, Computed: true, ElementType: types.StringType}, "cef": schema.SingleNestedAttribute{MarkdownDescription: "cef attributes", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"active_timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 604800)}}, "inactive_timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 604800)}}}}, "description": schema.StringAttribute{Optional: true, Computed: true}, "destination": schema.SingleNestedAttribute{MarkdownDescription: "destination attributes", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"dscp": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 63)}}, "ipv4_address": schema.StringAttribute{MarkdownDescription: "ipv4 address", Optional: true, Computed: true}, "l4_port_dst": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 65535)}}, "l4_port_src": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 65535)}}, "l4_protocol": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("udp", "tcp")}}, "ttl": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 255)}}}}, "max_pkt_size": schema.Int64Attribute{Optional: true, Computed: true}, "mobility_sam": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"encoding": schema.StringAttribute{Optional: true, Computed: true}, "encoding_format": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("hierarchy", "flat")}}, "event_enable": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"modify": schema.BoolAttribute{Optional: true, Computed: true}, "update": schema.BoolAttribute{Optional: true, Computed: true}}}, "trigger": schema.StringAttribute{Optional: true, Computed: true}}}, "monitor": schema.SingleNestedAttribute{MarkdownDescription: "monitor attributes", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"timeout": schema.Int64Attribute{MarkdownDescription: "how often to export in seconds", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(60, 900)}}}}, "netflow": schema.SingleNestedAttribute{MarkdownDescription: "netflow attributes", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"active_timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 604800)}}, "inactive_timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 604800)}}, "template_refresh": schema.Int64Attribute{MarkdownDescription: "template refresh interval in seconds", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 216000)}}, "template_type": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("cohesive", "segregated")}}, "version": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("v5", "v9", "ipfix")}}}}, "snmp": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{MarkdownDescription: "snmp reverse lookup enable/disable", Optional: true, Computed: true}}}, "source": schema.SingleNestedAttribute{MarkdownDescription: "source tunnel port", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"ip_interface": schema.StringAttribute{Optional: true, Computed: true}}}, "type": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("cef", "netflow", "monitor")}}}}, "port_packet_threshold_config": schema.SingleNestedAttribute{MarkdownDescription: "Contains port threshold configs", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"drop_threshold": schema.SingleNestedAttribute{MarkdownDescription: "Drop threshold config for Rx and Tx", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"rx": schema.ListNestedAttribute{MarkdownDescription: "Threshold values for Rx packet drop", Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"count_": schema.Int64Attribute{MarkdownDescription: "Threshold value in number of packets out of total packets in a given interval. '-1' will disable threshold feature.", Optional: true, Computed: true}, "percent": schema.Float64Attribute{MarkdownDescription: "Threshold value in percent of packets out of total packets in a given interval. '-1' will disable threshold feature.", Optional: true, Computed: true}, "port_type": schema.StringAttribute{MarkdownDescription: "Port type", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("all", "network", "tool", "stack", "inline-net", "inline-tool", "hybrid", "gigasmart", "circuit")}}}}}, "tx": schema.ListNestedAttribute{MarkdownDescription: "Threshold values for Tx packet drop", Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"count_": schema.Int64Attribute{MarkdownDescription: "Threshold value in number of packets out of total packets in a given interval. '-1' will disable threshold feature.", Optional: true, Computed: true}, "percent": schema.Float64Attribute{MarkdownDescription: "Threshold value in percent of packets out of total packets in a given interval. '-1' will disable threshold feature.", Optional: true, Computed: true}, "port_type": schema.StringAttribute{MarkdownDescription: "Port type", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("all", "network", "tool", "stack", "inline-net", "inline-tool", "hybrid", "gigasmart", "circuit")}}}}}}}, "error_threshold": schema.SingleNestedAttribute{MarkdownDescription: "Error threshold config for Rx and Tx", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"rx": schema.ListNestedAttribute{MarkdownDescription: "Threshold values for Rx packet error", Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"count_": schema.Int64Attribute{MarkdownDescription: "Threshold value in number of packets out of total packets in a given interval. '-1' will disable threshold feature.", Optional: true, Computed: true}, "percent": schema.Float64Attribute{MarkdownDescription: "Threshold value in percent of packets out of total packets in a given interval. '-1' will disable threshold feature.", Optional: true, Computed: true}, "port_type": schema.StringAttribute{MarkdownDescription: "Port type", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("all", "network", "tool", "stack", "inline-net", "inline-tool", "hybrid", "circuit")}}}}}, "tx": schema.ListNestedAttribute{MarkdownDescription: "Threshold values for Tx packet error", Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"count_": schema.Int64Attribute{MarkdownDescription: "Threshold value in number of packets out of total packets in a given interval. '-1' will disable threshold feature.", Optional: true, Computed: true}, "percent": schema.Float64Attribute{MarkdownDescription: "Threshold value in percent of packets out of total packets in a given interval. '-1' will disable threshold feature.", Optional: true, Computed: true}, "port_type": schema.StringAttribute{MarkdownDescription: "Port type", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("all", "network", "tool", "stack", "inline-net", "inline-tool", "hybrid", "circuit")}}}}}}}}}, "proxy_server_profile": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{Required: true}, "auth_type": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.OneOf("none", "basic")}}, "comment": schema.StringAttribute{Optional: true, Computed: true}, "password": schema.StringAttribute{Optional: true, Computed: true, Sensitive: true}, "periodic_ping": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.OneOf("enable", "disable")}}, "periodic_ping_failure_retry": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 3)}}, "periodic_ping_interval": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 10)}}, "periodic_ping_type": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("http-connect")}}, "port": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.Between(1, 65535)}}, "protocol": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.OneOf("http")}}, "proxy_address": schema.StringAttribute{Required: true}, "ssl_apps": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"cluster_name": schema.ListAttribute{MarkdownDescription: "Cluster Name where the proxy deployed", Optional: true, Computed: true, ElementType: types.StringType}}}, "username": schema.StringAttribute{Optional: true, Computed: true}}}, "snmp_trap_event_configs": schema.ListNestedAttribute{MarkdownDescription: "Available when ConfigType is SNMPTRAPS", Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{Optional: true, Computed: true}, "notify_event": schema.StringAttribute{Optional: true, Computed: true}}}}, "snmp_v3_users_config": schema.SingleNestedAttribute{MarkdownDescription: "Contains list of FM's snmpv3 users", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"snmp_v3_user": schema.ListNestedAttribute{Required: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"auth_key": schema.StringAttribute{MarkdownDescription: "auth key/passphrase", Required: true}, "auth_protocol": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.OneOf("md5", "sha1", "sha256", "sha384", "sha512")}}, "min_sw_version": schema.StringAttribute{MarkdownDescription: "minimum software version of the device", Optional: true, Computed: true}, "previous_username": schema.StringAttribute{MarkdownDescription: "When the username is changed, the username prior changing is sent in this property. Based on this application handles the username update logic. Not found error is returned when this property is not sent during the username change", Optional: true, Computed: true}, "priv_key": schema.StringAttribute{MarkdownDescription: "privacy key/passphrase", Required: true}, "priv_protocol": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.OneOf("des", "aes128", "aes256")}}, "username": schema.StringAttribute{MarkdownDescription: "snmpv3 username", Required: true}}}}}}, "ssh_ciphers_config": schema.SingleNestedAttribute{MarkdownDescription: "System SSH Cipher", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"classic": schema.SingleNestedAttribute{MarkdownDescription: "System SSH Cipher", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"client_ciphers": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType, Validators: []validator.List{listvalidator.ValueStringsAre(stringvalidator.OneOf("default", "aes128_cbc", "aes128_ctr", "aes128_gcm", "aes192_ctr", "aes256_cbc", "aes256_ctr", "aes256_gcm"))}}, "client_hostkey": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType, Validators: []validator.List{listvalidator.ValueStringsAre(stringvalidator.OneOf("default", "ecdsa_sha2_nistp256", "ecdsa_sha2_nistp384", "ecdsa_sha2_nistp521", "rsa_sha2_512", "rsa_sha2_256"))}}, "client_kex": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType, Validators: []validator.List{listvalidator.ValueStringsAre(stringvalidator.OneOf("default", "ecdh_sha2_nistp256", "ecdh_sha2_nistp384", "ecdh_sha2_nistp521", "diffie_hellman_group14_sha256"))}}, "client_macs": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType, Validators: []validator.List{listvalidator.ValueStringsAre(stringvalidator.OneOf("default", "hmac_sha2_512", "hmac_sha2_256"))}}, "server_ciphers": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType, Validators: []validator.List{listvalidator.ValueStringsAre(stringvalidator.OneOf("default", "aes128_cbc", "aes128_ctr", "aes128_gcm", "aes192_ctr", "aes256_cbc", "aes256_ctr", "aes256_gcm"))}}, "server_hostkey": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType, Validators: []validator.List{listvalidator.ValueStringsAre(stringvalidator.OneOf("default", "ecdsa_sha2_nistp256", "ecdsa_sha2_nistp384", "ecdsa_sha2_nistp521", "rsa_sha2_512", "rsa_sha2_256"))}}, "server_kex": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType, Validators: []validator.List{listvalidator.ValueStringsAre(stringvalidator.OneOf("default", "ecdh_sha2_nistp256", "ecdh_sha2_nistp384", "ecdh_sha2_nistp521", "diffie_hellman_group14_sha256"))}}, "server_macs": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType, Validators: []validator.List{listvalidator.ValueStringsAre(stringvalidator.OneOf("default", "hmac_sha2_512", "hmac_sha2_256"))}}}}, "crypto": schema.SingleNestedAttribute{MarkdownDescription: "System SSH Cipher", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"client_ciphers": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType, Validators: []validator.List{listvalidator.ValueStringsAre(stringvalidator.OneOf("default", "aes128_cbc", "aes128_ctr", "aes128_gcm", "aes192_ctr", "aes256_cbc", "aes256_ctr", "aes256_gcm"))}}, "client_hostkey": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType, Validators: []validator.List{listvalidator.ValueStringsAre(stringvalidator.OneOf("default", "ecdsa_sha2_nistp256", "ecdsa_sha2_nistp384", "ecdsa_sha2_nistp521", "rsa_sha2_512", "rsa_sha2_256"))}}, "client_kex": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType, Validators: []validator.List{listvalidator.ValueStringsAre(stringvalidator.OneOf("default", "ecdh_sha2_nistp256", "ecdh_sha2_nistp384", "ecdh_sha2_nistp521", "diffie_hellman_group14_sha256"))}}, "client_macs": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType, Validators: []validator.List{listvalidator.ValueStringsAre(stringvalidator.OneOf("default", "hmac_sha2_512", "hmac_sha2_256"))}}, "server_ciphers": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType, Validators: []validator.List{listvalidator.ValueStringsAre(stringvalidator.OneOf("default", "aes128_cbc", "aes128_ctr", "aes128_gcm", "aes192_ctr", "aes256_cbc", "aes256_ctr", "aes256_gcm"))}}, "server_hostkey": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType, Validators: []validator.List{listvalidator.ValueStringsAre(stringvalidator.OneOf("default", "ecdsa_sha2_nistp256", "ecdsa_sha2_nistp384", "ecdsa_sha2_nistp521", "rsa_sha2_512", "rsa_sha2_256"))}}, "server_kex": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType, Validators: []validator.List{listvalidator.ValueStringsAre(stringvalidator.OneOf("default", "ecdh_sha2_nistp256", "ecdh_sha2_nistp384", "ecdh_sha2_nistp521", "diffie_hellman_group14_sha256"))}}, "server_macs": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType, Validators: []validator.List{listvalidator.ValueStringsAre(stringvalidator.OneOf("default", "hmac_sha2_512", "hmac_sha2_256"))}}}}, "fips": schema.SingleNestedAttribute{MarkdownDescription: "System SSH Cipher", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"client_ciphers": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType, Validators: []validator.List{listvalidator.ValueStringsAre(stringvalidator.OneOf("default", "aes128_cbc", "aes128_ctr", "aes128_gcm", "aes192_ctr", "aes256_cbc", "aes256_ctr", "aes256_gcm"))}}, "client_hostkey": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType, Validators: []validator.List{listvalidator.ValueStringsAre(stringvalidator.OneOf("default", "ecdsa_sha2_nistp256", "ecdsa_sha2_nistp384", "ecdsa_sha2_nistp521", "rsa_sha2_512", "rsa_sha2_256"))}}, "client_kex": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType, Validators: []validator.List{listvalidator.ValueStringsAre(stringvalidator.OneOf("default", "ecdh_sha2_nistp256", "ecdh_sha2_nistp384", "ecdh_sha2_nistp521", "diffie_hellman_group14_sha256"))}}, "client_macs": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType, Validators: []validator.List{listvalidator.ValueStringsAre(stringvalidator.OneOf("default", "hmac_sha2_512", "hmac_sha2_256"))}}, "server_ciphers": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType, Validators: []validator.List{listvalidator.ValueStringsAre(stringvalidator.OneOf("default", "aes128_cbc", "aes128_ctr", "aes128_gcm", "aes192_ctr", "aes256_cbc", "aes256_ctr", "aes256_gcm"))}}, "server_hostkey": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType, Validators: []validator.List{listvalidator.ValueStringsAre(stringvalidator.OneOf("default", "ecdsa_sha2_nistp256", "ecdsa_sha2_nistp384", "ecdsa_sha2_nistp521", "rsa_sha2_512", "rsa_sha2_256"))}}, "server_kex": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType, Validators: []validator.List{listvalidator.ValueStringsAre(stringvalidator.OneOf("default", "ecdh_sha2_nistp256", "ecdh_sha2_nistp384", "ecdh_sha2_nistp521", "diffie_hellman_group14_sha256"))}}, "server_macs": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType, Validators: []validator.List{listvalidator.ValueStringsAre(stringvalidator.OneOf("default", "hmac_sha2_512", "hmac_sha2_256"))}}}}}}}}, "config_level": schema.StringAttribute{MarkdownDescription: "Scope of the applied FM template", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("GLOBAL", "SITE", "CLUSTER", "METADATA_SOLUTION", "TAG", "APP_INTEL_SOLUTION")}}, "config_level_value": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType}, "config_resource": schema.DynamicAttribute{MarkdownDescription: "For a particular config type, there may be multiple templates with different levels. For example, cluster level the user can have more than one template and in same way in tag level there may be more than one template for each tag combination. In such cases, the configResource property can be used to contain the list of clusters associated to the template or the tag key values combination of the template. The data structure changes based on the level hence it is kept as type Object.", Optional: true, Computed: true}, "config_type": schema.StringAttribute{MarkdownDescription: "Configuration Type of the FM template", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("SNMPTRAPS", "METADATA_TEMPLATE", "PORT_NEIGHBOR_DISCOVERY", "ACME_TEMPLATE", "LDAP_SERVERS_TEMPLATE", "LDAP_SYSTEM_CONFIG_TEMPLATE", "AAA_AUTH_CONFIG_TEMPLATE", "SSL_CERTIFICATE_TEMPLATE", "PORT_PACKET_THRESHOLD", "GIGASTREAM_THRESHOLD", "USERDEFINED_APPLICATION_TEMPLATE", "SSH_CIPHERS_CONFIG", "FM_SNMPV3_USER", "PROXY_SERVER_PROFILE", "TLS_CIPHERS_CONFIG")}}, "modifiable": schema.BoolAttribute{Optional: true, Computed: true}, "ref_count": schema.Int64Attribute{Optional: true, Computed: true}, "ref_object": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"ref_object_type": schema.StringAttribute{MarkdownDescription: "Type of the reference object", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("PHYSICAL", "VIRTUAL")}}}}, "template_name": schema.StringAttribute{Optional: true, Computed: true}, "update_time": schema.StringAttribute{Optional: true, Computed: true}}, Blocks: map[string]schema.Block{"timeouts": timeouts.Block(ctx, timeouts.Opts{Create: true, Read: true, Update: true, Delete: true})}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *FmTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan FmTemplateResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	timeout, diags := plan.Timeouts.Create(ctx, 20*time.Minute)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	r.createRemote(ctx, &plan, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// createRemote performs the create HTTP exchange and decodes the response into plan. Extracted from Create so the request/response logic is unit-testable without a tfsdk.Plan.
func (r *FmTemplateResource) createRemote(ctx context.Context, plan *FmTemplateResourceModel, resp *resource.CreateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_fm_template", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_fm_template", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/fm/templates"
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_fm_template", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_fm_template", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error creating gigavuecore_fm_template", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error creating gigavuecore_fm_template", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error creating gigavuecore_fm_template", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error creating gigavuecore_fm_template", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error creating gigavuecore_fm_template", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error creating gigavuecore_fm_template", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error creating gigavuecore_fm_template", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error creating gigavuecore_fm_template", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error creating gigavuecore_fm_template", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error creating gigavuecore_fm_template", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_fm_template", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
	if plan.TemplateName.IsNull() || plan.TemplateName.IsUnknown() {
		loc := httpResp.Header.Get("Location")
		if loc != "" {
			loc = strings.TrimRight(loc, "/")
			i := strings.LastIndex(loc, "/")
			if i >= 0 {
				loc = loc[i+1:]
			}
			plan.TemplateName = types.StringValue(loc)
		} else {
			resp.Diagnostics.AddError("Error creating gigavuecore_fm_template", "The create response did not contain an identifier and no Location header was returned, so the resource cannot be tracked in state.")
			return
		}
	}
}

// Read refreshes the Terraform state with the latest remote values.
func (r *FmTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state FmTemplateResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	timeout, diags := state.Timeouts.Read(ctx, 10*time.Minute)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	if r.readRemote(ctx, &state, resp) {
		resp.State.RemoveResource(ctx)
		return
	}
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// readRemote performs the read HTTP exchange and decodes the response into state, returning removed=true when the API reports 404. Extracted from Read so the request/response logic is unit-testable without a tfsdk.State.
func (r *FmTemplateResource) readRemote(ctx context.Context, state *FmTemplateResourceModel, resp *resource.ReadResponse) (removed bool) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/fm/templates/{templateName}"
	reqPath = strings.ReplaceAll(reqPath, "{templateName}", url.PathEscape(state.TemplateName.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_fm_template", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_fm_template", fmt.Sprintf("Could not send request: %s", err))
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
			resp.Diagnostics.AddError("Error reading gigavuecore_fm_template", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_fm_template", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_fm_template", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_fm_template", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_fm_template", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_fm_template", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_fm_template", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_fm_template", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_fm_template", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_fm_template", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&state, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_fm_template", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
	return
}

// Update modifies the remote resource to match the desired plan.
func (r *FmTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan FmTemplateResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state FmTemplateResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	timeout, diags := plan.Timeouts.Update(ctx, 20*time.Minute)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	if plan.TemplateName.IsNull() || plan.TemplateName.IsUnknown() {
		if !state.TemplateName.IsNull() && !state.TemplateName.IsUnknown() {
			plan.TemplateName = state.TemplateName
		}
	}
	preserveStateIntoPlan(&plan, &state)
	r.updateRemote(ctx, &plan, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// updateRemote performs the update HTTP exchange and decodes the response into plan. Extracted from Update so the request/response logic is unit-testable without a tfsdk.Plan.
func (r *FmTemplateResource) updateRemote(ctx context.Context, plan *FmTemplateResourceModel, resp *resource.UpdateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_fm_template", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_fm_template", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/fm/templates/{templateName}"
	reqPath = strings.ReplaceAll(reqPath, "{templateName}", url.PathEscape(plan.TemplateName.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodPatch, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_fm_template", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_fm_template", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error updating gigavuecore_fm_template", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error updating gigavuecore_fm_template", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error updating gigavuecore_fm_template", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error updating gigavuecore_fm_template", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error updating gigavuecore_fm_template", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error updating gigavuecore_fm_template", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error updating gigavuecore_fm_template", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error updating gigavuecore_fm_template", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error updating gigavuecore_fm_template", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error updating gigavuecore_fm_template", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_fm_template", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Delete destroys the remote resource.
func (r *FmTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state FmTemplateResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	timeout, diags := state.Timeouts.Delete(ctx, 10*time.Minute)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	r.deleteRemote(ctx, &state, resp)
}

// deleteRemote performs the delete HTTP exchange, treating a 404 as already deleted. Extracted from Delete so the request/response logic is unit-testable without a tfsdk.State.
func (r *FmTemplateResource) deleteRemote(ctx context.Context, state *FmTemplateResourceModel, resp *resource.DeleteResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/fm/templates/{templateName}"
	reqPath = strings.ReplaceAll(reqPath, "{templateName}", url.PathEscape(state.TemplateName.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_fm_template", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_fm_template", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		return
	}
	if !(httpResp.StatusCode == 204) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error deleting gigavuecore_fm_template", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error deleting gigavuecore_fm_template", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error deleting gigavuecore_fm_template", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error deleting gigavuecore_fm_template", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error deleting gigavuecore_fm_template", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error deleting gigavuecore_fm_template", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error deleting gigavuecore_fm_template", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error deleting gigavuecore_fm_template", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error deleting gigavuecore_fm_template", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *FmTemplateResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *FmTemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("template_name"), req.ID)...)
}
