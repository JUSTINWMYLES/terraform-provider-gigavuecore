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
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	path "github.com/hashicorp/terraform-plugin-framework/path"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertions.
var (
	_ resource.Resource                = (*SolutionResource)(nil)
	_ resource.ResourceWithImportState = (*SolutionResource)(nil)
	_ resource.ResourceWithConfigure   = (*SolutionResource)(nil)
)

// SolutionResource is the generated Terraform managed resource implementation.
type SolutionResource struct {
	client *client.Client
}

// SolutionResourceModel describes the Terraform state and plan shape for SolutionResource.
type SolutionResourceModel struct {
	AppExportConfig                types.Object  `tfsdk:"app_export_config" json:"appExportConfig"`
	AppFilterConfig                types.Object  `tfsdk:"app_filter_config" json:"appFilterConfig"`
	AssociatedMonitorSolutionAlias types.String  `tfsdk:"associated_monitor_solution_alias" json:"associatedMonitorSolutionAlias"`
	ClusterId                      types.String  `tfsdk:"cluster_id" json:"clusterId"`
	ConfigStatus                   types.String  `tfsdk:"config_status" json:"configStatus"`
	DeleteMonitorSol               types.Bool    `tfsdk:"delete_monitor_sol" json:"deleteMonitorSol"`
	EgressMapAliasesToDelete       types.List    `tfsdk:"egress_map_aliases_to_delete" json:"egressMapAliasesToDelete"`
	ExporterAliasesToDelete        types.List    `tfsdk:"exporter_aliases_to_delete" json:"exporterAliasesToDelete"`
	HealthState                    types.String  `tfsdk:"health_state" json:"healthState"`
	HealthStateReasons             types.List    `tfsdk:"health_state_reasons" json:"healthStateReasons"`
	IngressMapAliasesToDelete      types.List    `tfsdk:"ingress_map_aliases_to_delete" json:"ingressMapAliasesToDelete"`
	IngressTrafficConfigs          types.Dynamic `tfsdk:"ingress_traffic_configs" json:"ingressTrafficConfigs"`
	MonitorSolutionConfig          types.Object  `tfsdk:"monitor_solution_config" json:"monitorSolutionConfig"`
	SolutionAlias                  types.String  `tfsdk:"solution_alias" json:"solutionAlias"`
	SolutionDesc                   types.String  `tfsdk:"solution_desc" json:"solutionDesc"`
	SolutionStatus                 types.String  `tfsdk:"solution_status" json:"solutionStatus"`
	SolutionType                   types.String  `tfsdk:"solution_type" json:"solutionType"`
}

// Metadata returns the resource type name.
func (r *SolutionResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_solution"
}

// Schema returns the Terraform schema for this resource.
func (r *SolutionResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load Apps Visibility Solutions by alias", Attributes: map[string]schema.Attribute{"app_export_config": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"cache_config": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"advance_hash": schema.BoolAttribute{MarkdownDescription: "When scheduling flows for processing, use external encapsulate packet header if disable, use inner packet header if enable.", Optional: true, Computed: true}, "alias": schema.StringAttribute{Required: true}, "description": schema.StringAttribute{Optional: true, Computed: true}, "dpi_inject_limit": schema.Int64Attribute{Optional: true, Computed: true}, "event": schema.StringAttribute{Optional: true, Computed: true}, "exporters": schema.ListAttribute{MarkdownDescription: "alias of metadata exporters to attach this cache", Optional: true, Computed: true, ElementType: types.StringType}, "flow_behavior": schema.StringAttribute{MarkdownDescription: "direction for flow identification", Optional: true, Computed: true}, "match": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"datalink": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"mac_dst": schema.BoolAttribute{Optional: true, Computed: true}, "mac_src": schema.BoolAttribute{Optional: true, Computed: true}, "vlan": schema.BoolAttribute{Optional: true, Computed: true}}}, "interface": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"in_name_width": schema.Int64Attribute{Optional: true, Computed: true}, "in_physical_width": schema.Int64Attribute{Optional: true, Computed: true}}}, "ip": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"version": schema.BoolAttribute{Optional: true, Computed: true}}}, "ipv4": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"destination": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"prefix_min_mask": schema.StringAttribute{MarkdownDescription: "ipv4 destination prefix minimum-mask - netmask or mask length", Optional: true, Computed: true}}}, "dscp": schema.BoolAttribute{Optional: true, Computed: true}, "fragmentation": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"flags": schema.BoolAttribute{Optional: true, Computed: true}, "offset": schema.BoolAttribute{Optional: true, Computed: true}}}, "header_len": schema.BoolAttribute{Optional: true, Computed: true}, "option_map": schema.BoolAttribute{Optional: true, Computed: true}, "precedence": schema.BoolAttribute{Optional: true, Computed: true}, "protocol": schema.BoolAttribute{Optional: true, Computed: true}, "section": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"header_size": schema.Int64Attribute{Optional: true, Computed: true}, "payload_size": schema.Int64Attribute{Optional: true, Computed: true}}}, "source": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"prefix_min_mask": schema.StringAttribute{MarkdownDescription: "ipv4 source prefix minimum-mask - netmask or mask length", Optional: true, Computed: true}}}, "tos": schema.BoolAttribute{Optional: true, Computed: true}, "total_length": schema.BoolAttribute{Optional: true, Computed: true}, "ttl": schema.BoolAttribute{Optional: true, Computed: true}}}, "ipv6": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"destination": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"prefix_min_mask": schema.StringAttribute{Optional: true, Computed: true}}}, "dscp": schema.BoolAttribute{Optional: true, Computed: true}, "extension_map": schema.BoolAttribute{Optional: true, Computed: true}, "flow_label": schema.BoolAttribute{Optional: true, Computed: true}, "fragmentation": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"flags": schema.BoolAttribute{Optional: true, Computed: true}, "offset": schema.BoolAttribute{Optional: true, Computed: true}}}, "hop_limit": schema.BoolAttribute{Optional: true, Computed: true}, "length": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"header": schema.BoolAttribute{Optional: true, Computed: true}, "payload": schema.BoolAttribute{Optional: true, Computed: true}, "total": schema.BoolAttribute{Optional: true, Computed: true}}}, "next_header": schema.BoolAttribute{Optional: true, Computed: true}, "precedence": schema.BoolAttribute{Optional: true, Computed: true}, "section": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"header_size": schema.Int64Attribute{Optional: true, Computed: true}, "payload_size": schema.Int64Attribute{Optional: true, Computed: true}}}, "source": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"prefix_min_mask": schema.StringAttribute{MarkdownDescription: "ipv6 source prefix minimum-mask - netmask or mask length", Optional: true, Computed: true}}}, "traffic_class": schema.BoolAttribute{Optional: true, Computed: true}}}, "transport": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"dst_port": schema.BoolAttribute{Optional: true, Computed: true}, "icmp": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"ipv4_code": schema.BoolAttribute{Optional: true, Computed: true}, "ipv4_type": schema.BoolAttribute{Optional: true, Computed: true}, "ipv6_code": schema.BoolAttribute{Optional: true, Computed: true}, "ipv6_type": schema.BoolAttribute{Optional: true, Computed: true}}}, "src_port": schema.BoolAttribute{Optional: true, Computed: true}, "tcp": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"ack_number": schema.BoolAttribute{Optional: true, Computed: true}, "dst_port": schema.BoolAttribute{Optional: true, Computed: true}, "flags": schema.BoolAttribute{Optional: true, Computed: true}, "header_len": schema.BoolAttribute{Optional: true, Computed: true}, "seq_number": schema.BoolAttribute{Optional: true, Computed: true}, "src_port": schema.BoolAttribute{Optional: true, Computed: true}, "urgent_ptr": schema.BoolAttribute{Optional: true, Computed: true}, "window_size": schema.BoolAttribute{Optional: true, Computed: true}}}, "udp": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"dst_port": schema.BoolAttribute{Optional: true, Computed: true}, "msg_len": schema.BoolAttribute{Optional: true, Computed: true}, "src_port": schema.BoolAttribute{Optional: true, Computed: true}}}}}}}, "multi_collect": schema.BoolAttribute{MarkdownDescription: "Collect all attributes as it is discovered when enable. It will export the same record once when disable.", Optional: true, Computed: true}, "network_profiles": schema.ListAttribute{MarkdownDescription: "alias of metadata network profiles to attach this cache", Optional: true, Computed: true, ElementType: types.StringType}, "observation_domain_id": schema.Int64Attribute{Optional: true, Computed: true}, "sampling": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"mode": schema.StringAttribute{Optional: true, Computed: true}, "single_sampling_rate": schema.Int64Attribute{MarkdownDescription: "Packet interval window size. Valid values: 10-16000 (in packets)", Optional: true, Computed: true}}}, "size": schema.SingleNestedAttribute{MarkdownDescription: "size of the flows", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"flows": schema.Int64Attribute{MarkdownDescription: "size of flows in millions", Optional: true, Computed: true}}}, "timeout": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"idle": schema.Int64Attribute{MarkdownDescription: "idle timeout in seconds. max value 7days. default 30 min", Optional: true, Computed: true}}}}}, "cache_config_alias": schema.StringAttribute{Optional: true, Computed: true}, "destination_configs": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"application_names": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"attributes": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"name": schema.StringAttribute{MarkdownDescription: "attribute name", Required: true}, "value": schema.StringAttribute{MarkdownDescription: "application's attribute value", Optional: true, Computed: true}}}}, "is_user_defined": schema.BoolAttribute{MarkdownDescription: "Default value is false. Must be set to true only for App Intel solution while configuring user defined apps", Optional: true, Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "application name", Required: true}}}}, "destination_name": schema.StringAttribute{Optional: true, Computed: true}, "export_ip_interface": schema.StringAttribute{Optional: true, Computed: true}, "export_meta_app_profile": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "application profile alias", Required: true}, "application_id": schema.BoolAttribute{MarkdownDescription: "only valid with 'export' type", Optional: true, Computed: true}, "applications": schema.ListNestedAttribute{MarkdownDescription: "application and attributes.", Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"attributes": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"name": schema.StringAttribute{MarkdownDescription: "attribute name", Required: true}, "value": schema.StringAttribute{MarkdownDescription: "application's attribute value", Optional: true, Computed: true}}}}, "is_user_defined": schema.BoolAttribute{MarkdownDescription: "Default value is false. Must be set to true only for App Intel solution while configuring user defined apps", Optional: true, Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "application name", Required: true}}}}, "counter": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"bytes": schema.BoolAttribute{Optional: true, Computed: true}, "bytes_long": schema.BoolAttribute{Optional: true, Computed: true}, "inner_byte": schema.BoolAttribute{Optional: true, Computed: true}, "inner_byte_long": schema.BoolAttribute{Optional: true, Computed: true}, "packets": schema.BoolAttribute{Optional: true, Computed: true}, "packets_long": schema.BoolAttribute{Optional: true, Computed: true}}}, "datalink": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"mac_dst": schema.BoolAttribute{Optional: true, Computed: true}, "mac_src": schema.BoolAttribute{Optional: true, Computed: true}, "vlan": schema.BoolAttribute{Optional: true, Computed: true}}}, "description": schema.StringAttribute{Optional: true, Computed: true}, "flow": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"end_reason": schema.BoolAttribute{Optional: true, Computed: true}}}, "gtpu": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"qfi": schema.BoolAttribute{Optional: true, Computed: true}, "teid": schema.BoolAttribute{Optional: true, Computed: true}}}, "interface": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"in_name_width": schema.Int64Attribute{Optional: true, Computed: true}, "in_physical_width": schema.Int64Attribute{Optional: true, Computed: true}, "out_physical_width": schema.Int64Attribute{Optional: true, Computed: true}}}, "ip": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"version": schema.BoolAttribute{Optional: true, Computed: true}}}, "ipv4": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"destination": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"prefix_min_mask": schema.StringAttribute{Optional: true, Computed: true}}}, "dscp": schema.BoolAttribute{Optional: true, Computed: true}, "fragmentation": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"flags": schema.BoolAttribute{Optional: true, Computed: true}, "offset": schema.BoolAttribute{Optional: true, Computed: true}}}, "header_len": schema.BoolAttribute{Optional: true, Computed: true}, "option_map": schema.BoolAttribute{Optional: true, Computed: true}, "precedence": schema.BoolAttribute{Optional: true, Computed: true}, "protocol": schema.BoolAttribute{Optional: true, Computed: true}, "section": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"header_size": schema.Int64Attribute{Optional: true, Computed: true}, "payload_size": schema.Int64Attribute{Optional: true, Computed: true}}}, "source": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"prefix_min_mask": schema.StringAttribute{MarkdownDescription: "ipv4 source prefix minimum-mask - netmask or mask length", Optional: true, Computed: true}}}, "tos": schema.BoolAttribute{Optional: true, Computed: true}, "total_length": schema.BoolAttribute{Optional: true, Computed: true}, "ttl": schema.BoolAttribute{Optional: true, Computed: true}}}, "ipv6": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"destination": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"prefix_min_mask": schema.StringAttribute{Optional: true, Computed: true}}}, "dscp": schema.BoolAttribute{Optional: true, Computed: true}, "extension_map": schema.BoolAttribute{Optional: true, Computed: true}, "flow_label": schema.BoolAttribute{Optional: true, Computed: true}, "fragmentation": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"flags": schema.BoolAttribute{Optional: true, Computed: true}, "offset": schema.BoolAttribute{Optional: true, Computed: true}}}, "hop_limit": schema.BoolAttribute{Optional: true, Computed: true}, "length": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"header": schema.BoolAttribute{Optional: true, Computed: true}, "payload": schema.BoolAttribute{Optional: true, Computed: true}, "total": schema.BoolAttribute{Optional: true, Computed: true}}}, "next_header": schema.BoolAttribute{Optional: true, Computed: true}, "precedence": schema.BoolAttribute{Optional: true, Computed: true}, "section": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"header_size": schema.Int64Attribute{Optional: true, Computed: true}, "payload_size": schema.Int64Attribute{Optional: true, Computed: true}}}, "source": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"prefix_min_mask": schema.StringAttribute{Optional: true, Computed: true}}}, "traffic_class": schema.BoolAttribute{Optional: true, Computed: true}}}, "outer_ipv4": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"destination": schema.BoolAttribute{Optional: true, Computed: true}, "source": schema.BoolAttribute{Optional: true, Computed: true}}}, "outer_ipv6": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"destination": schema.BoolAttribute{Optional: true, Computed: true}, "source": schema.BoolAttribute{Optional: true, Computed: true}}}, "timestamp": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"flow_end_msec": schema.BoolAttribute{Optional: true, Computed: true}, "flow_endsec": schema.BoolAttribute{Optional: true, Computed: true}, "flow_start_msec": schema.BoolAttribute{Optional: true, Computed: true}, "flow_startsec": schema.BoolAttribute{Optional: true, Computed: true}, "sys_up_time_first": schema.BoolAttribute{Optional: true, Computed: true}, "sys_up_time_last": schema.BoolAttribute{Optional: true, Computed: true}}}, "transport": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"dst_port": schema.BoolAttribute{Optional: true, Computed: true}, "icmp": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"ipv4_code": schema.BoolAttribute{Optional: true, Computed: true}, "ipv4_type": schema.BoolAttribute{Optional: true, Computed: true}, "ipv6_code": schema.BoolAttribute{Optional: true, Computed: true}, "ipv6_type": schema.BoolAttribute{Optional: true, Computed: true}}}, "src_port": schema.BoolAttribute{Optional: true, Computed: true}, "tcp": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"ack_number": schema.BoolAttribute{Optional: true, Computed: true}, "dst_port": schema.BoolAttribute{Optional: true, Computed: true}, "flags": schema.BoolAttribute{Optional: true, Computed: true}, "header_len": schema.BoolAttribute{Optional: true, Computed: true}, "seq_number": schema.BoolAttribute{Optional: true, Computed: true}, "src_port": schema.BoolAttribute{Optional: true, Computed: true}, "urgent_ptr": schema.BoolAttribute{Optional: true, Computed: true}, "window_size": schema.BoolAttribute{Optional: true, Computed: true}}}, "udp": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"dst_port": schema.BoolAttribute{Optional: true, Computed: true}, "msg_len": schema.BoolAttribute{Optional: true, Computed: true}, "src_port": schema.BoolAttribute{Optional: true, Computed: true}}}}}, "type": schema.StringAttribute{Optional: true, Computed: true}}}, "export_meta_app_profile_alias": schema.StringAttribute{Optional: true, Computed: true}, "exporter_alias": schema.StringAttribute{Optional: true, Computed: true}, "exporter_config": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{Required: true}, "application_profiles": schema.ListAttribute{MarkdownDescription: "application profile aliases to attach to the exporter", Optional: true, Computed: true, ElementType: types.StringType}, "cef": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"active_timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Optional: true, Computed: true}, "inactive_timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Optional: true, Computed: true}}}, "description": schema.StringAttribute{Optional: true, Computed: true}, "destination": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"dscp": schema.Int64Attribute{Optional: true, Computed: true}, "ipv4_address": schema.StringAttribute{MarkdownDescription: "ipv4 address", Optional: true, Computed: true}, "l4_port_dst": schema.Int64Attribute{Optional: true, Computed: true}, "l4_port_src": schema.Int64Attribute{Optional: true, Computed: true}, "l4_protocol": schema.StringAttribute{Optional: true, Computed: true}, "ttl": schema.Int64Attribute{Optional: true, Computed: true}}}, "max_pkt_size": schema.Int64Attribute{Optional: true, Computed: true}, "mobility_sam": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"encoding": schema.StringAttribute{Optional: true, Computed: true}, "encoding_format": schema.StringAttribute{Optional: true, Computed: true}, "event_enable": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"modify": schema.BoolAttribute{Optional: true, Computed: true}, "update": schema.BoolAttribute{Optional: true, Computed: true}}}, "trigger": schema.StringAttribute{Optional: true, Computed: true}}}, "monitor": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"timeout": schema.Int64Attribute{MarkdownDescription: "how often to export in seconds", Optional: true, Computed: true}}}, "netflow": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"active_timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Optional: true, Computed: true}, "inactive_timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Optional: true, Computed: true}, "template_refresh": schema.Int64Attribute{MarkdownDescription: "template refresh interval in seconds", Optional: true, Computed: true}, "template_type": schema.StringAttribute{Optional: true, Computed: true}, "version": schema.StringAttribute{Optional: true, Computed: true}}}, "snmp": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{MarkdownDescription: "snmp reverse lookup enable/disable", Optional: true, Computed: true}}}, "source": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"ip_interface": schema.StringAttribute{Optional: true, Computed: true}}}, "type": schema.StringAttribute{Optional: true, Computed: true}}}}}}, "gsop_alias": schema.StringAttribute{Optional: true, Computed: true}, "gsop_config": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART Operation", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{Required: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "id of the defining cluster", Optional: true, Computed: true}, "gs_apps": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART Applications for a GSOP. At least one GsApp must be defined for the config to be valid", Required: true, Attributes: map[string]schema.Attribute{"apf": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Optional: true, Computed: true}}}, "dedup": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Optional: true, Computed: true}}}, "diameter_whitelist": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Optional: true, Computed: true}}}, "flow_filter": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Flow Filter' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"type": schema.StringAttribute{Required: true}}}, "flow_sampling": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Flow Sampling' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"type": schema.StringAttribute{Required: true}}}, "gseries_header_add": schema.SingleNestedAttribute{MarkdownDescription: "Only applicable for G-series", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"types": schema.SetAttribute{Required: true, ElementType: types.StringType}}}, "gseries_header_remove": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Optional: true, Computed: true}}}, "gseries_load_balance": schema.SingleNestedAttribute{MarkdownDescription: "Only applicable for G-series per-rule GSOP", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"fixed_offset": schema.SingleNestedAttribute{MarkdownDescription: "Per-rule Fixed Offset Load Balancing config for G-seres devices", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"hash": schema.StringAttribute{Required: true}, "length": schema.Int64Attribute{Required: true}, "offset": schema.Int64Attribute{Required: true}}}, "variable_offset": schema.SingleNestedAttribute{MarkdownDescription: "Per-rule Variable Offset Load Balancing config for G-seres devices", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"end_delim": schema.StringAttribute{Required: true}, "hash": schema.StringAttribute{Required: true}, "start_delim": schema.StringAttribute{Required: true}, "start_field": schema.StringAttribute{Required: true}}}}}, "gseries_pattern_match": schema.SingleNestedAttribute{MarkdownDescription: "Only applicable for G-series per-rule GSOP", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"fixed_offset": schema.SingleNestedAttribute{MarkdownDescription: "Per-rule Fixed Offset Pattern Match config for G-seres devices", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"length": schema.Int64Attribute{Required: true}, "offset": schema.Int64Attribute{Required: true}}}, "variable_offset": schema.SingleNestedAttribute{MarkdownDescription: "Per-rule Variable Offset Pattern Match config for G-seres devices", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"end_delim": schema.StringAttribute{Required: true}, "start_delim": schema.StringAttribute{Required: true}}}}}, "gtp_whitelist": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Optional: true, Computed: true}}}, "header_add": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Add Header' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"vlan": schema.Int64Attribute{Required: true}}}, "header_remove": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Remove Header' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"ah1": schema.StringAttribute{MarkdownDescription: "only valid and required for 'generic'. First anchor header, from which the header to be stripped is occurred.", Optional: true, Computed: true}, "ah2": schema.StringAttribute{MarkdownDescription: "only valid and required for 'generic'. next anchor header.", Optional: true, Computed: true}, "custom_len": schema.Int64Attribute{MarkdownDescription: "only valid for 'generic'. length of unknown header.", Optional: true, Computed: true}, "erspan_flow_id": schema.Int64Attribute{MarkdownDescription: "only applicable for 'erspan', specifies the flow id to strip. Value of 0 will strip all flow ids", Optional: true, Computed: true}, "fp_dst_switch_id": schema.Int64Attribute{MarkdownDescription: "Only valid and required for 'fabricPath', 12 bit destination switch id", Optional: true, Computed: true}, "fp_src_switch_id": schema.Int64Attribute{MarkdownDescription: "Only valid and required for 'fabricPath', 12 bit source switch id", Optional: true, Computed: true}, "header_count": schema.Int64Attribute{MarkdownDescription: "only valid for 'generic'. Number of headers to be stripped.", Optional: true, Computed: true}, "offset": schema.StringAttribute{MarkdownDescription: "only valid and required for 'generic'.'start': strip from start of ah1, 'end': strip from end of ah1 or any other offset in the range of length of ah1, 'offsetRangeValue' is required if offset is 'offsetRange'.", Optional: true, Computed: true}, "offset_range_value": schema.Int64Attribute{MarkdownDescription: "only valid and required when offset is 'offsetRange', integer within range of size of header", Optional: true, Computed: true}, "protocol": schema.StringAttribute{MarkdownDescription: "'gre' and 'fabricPath' are only applicable for H-series", Required: true}, "timestamp_format": schema.StringAttribute{MarkdownDescription: "Timestamp format. Only valid and required for 'fm6000Ts'", Optional: true, Computed: true}, "vlan_header": schema.StringAttribute{MarkdownDescription: "Only valid when protocol is 'vlan'. Specifies the target vlan header to strip. Defaults to 'outer'", Optional: true, Computed: true}, "vxlan_id": schema.Int64Attribute{MarkdownDescription: "24-bit value. Only valid when protocol is 'vxlan'. Specifies the vxlan id to strip. Value of '0' will strip all vxlan ids", Optional: true, Computed: true}}}, "icap": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART ICAP Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"icap_profile": schema.StringAttribute{MarkdownDescription: "Alias of referenced ICAP Profile", Required: true}}}, "inline_ssl": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART Inline SSL Profile Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"inline_ssl_profile": schema.StringAttribute{MarkdownDescription: "Alias of referenced Inline SSL Profile", Required: true}}}, "load_balance": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Load Balancing' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enhanced": schema.SingleNestedAttribute{MarkdownDescription: "Private class. Enhanced part of the GigaSMART 'Load Balancing' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"elb_alias": schema.StringAttribute{MarkdownDescription: "elb app alias", Required: true}}}, "stateful": schema.SingleNestedAttribute{MarkdownDescription: "Private class. Stateful part of the GigaSMART 'Load Balancing' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"app_type": schema.StringAttribute{MarkdownDescription: "'sapf' option is added in release H 4.3.01. 'sip' and 'tunnel' added in release H 5.1. 'diameter' is added in release H 5.5", Required: true}, "diameter_key_hash_type": schema.StringAttribute{MarkdownDescription: "required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyHash'", Optional: true, Computed: true}, "diameter_key_multi_hash_type": schema.ListNestedAttribute{MarkdownDescription: "required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyMultiHash'", Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"avp_codevalue": schema.Int64Attribute{MarkdownDescription: "required when 'key' == 'avpCode'", Optional: true, Computed: true}, "key": schema.StringAttribute{Required: true}}}}, "gtp_key_hash_type": schema.StringAttribute{MarkdownDescription: "required when 'appType' == 'gtp' and 'lbType' == 'gtpKeyHash'. ignored otherwise", Optional: true, Computed: true}, "lb_type": schema.StringAttribute{MarkdownDescription: "'gtpKeyHash' is only applicable when 'appType' == 'gtp'. 'sipKeyHash' is only applicable when 'appType' == 'sip'", Required: true}, "sip_key_hash_type": schema.StringAttribute{MarkdownDescription: "required when 'appType' == 'sip' and 'lbType' == 'sipKeyHash'. ignored otherwise", Optional: true, Computed: true}}}, "stateless": schema.SingleNestedAttribute{MarkdownDescription: "Private class. Stateless part of the GigaSMART 'Load Balancing' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"field_location": schema.StringAttribute{MarkdownDescription: "Ignored when 'hashFields' == 'gtpuTeid'. Required otherwise", Optional: true, Computed: true}, "hash_fields": schema.StringAttribute{Required: true}}}}}, "masking": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Masking' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"content_type": schema.StringAttribute{MarkdownDescription: "content type that will trigger masking, only valid and required for protocol 'sip'", Optional: true, Computed: true}, "length": schema.Int64Attribute{MarkdownDescription: "max length is 9600 on H-series; 9000 on G-series, not applicable for protocol 'sip' valid and required otherwise", Optional: true, Computed: true}, "offset": schema.Int64Attribute{MarkdownDescription: "min offset is 0 for 'none'; 1 for protocol-referenced offsets, not applicable for protocol 'sip' valid and required otherwise", Optional: true, Computed: true}, "pattern": schema.StringAttribute{MarkdownDescription: "1-byte hex mask, not applicable for protocol 'sip' valid and required otherwise", Optional: true, Computed: true}, "protocol": schema.StringAttribute{Required: true}}}, "metadata_export": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"cache": schema.StringAttribute{MarkdownDescription: "metadata cache alias. cache should have exporters defined, and the exporters should have applicationProfiles defined", Optional: true, Computed: true}}}, "netflow": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Optional: true, Computed: true}}}, "sa_apf": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Optional: true, Computed: true}}}, "sip_whitelist": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Optional: true, Computed: true}}}, "slicing": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Slicing' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enhanced": schema.StringAttribute{MarkdownDescription: "enhanced-slicing apps alias", Optional: true, Computed: true}, "offset": schema.Int64Attribute{MarkdownDescription: "min offset is 64 for 'none'; 4 for protocol-referenced offsets. required property till H 5.6", Required: true}, "protocol": schema.StringAttribute{MarkdownDescription: "required property till H 5.6", Required: true}}}, "ssl_decrypt": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'SSL Decrypt' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"in_port": schema.Int64Attribute{MarkdownDescription: "Port number of 0 represents 'any' port", Optional: true, Computed: true}, "out_port": schema.Int64Attribute{MarkdownDescription: "Port number of 0 represents 'auto' port", Optional: true, Computed: true}}}, "trailer_add": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Add Trailer' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"types": schema.SetAttribute{MarkdownDescription: "'crc' is not applicable for G-series", Required: true, ElementType: types.StringType}}}, "trailer_remove": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.StringAttribute{Optional: true, Computed: true}}}, "tunnel_decap": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Decapsulate Tunnel' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"custom": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"port_dst": schema.Int64Attribute{MarkdownDescription: "when specified as 0, no validation will be done in the packet.", Required: true}, "port_src": schema.Int64Attribute{MarkdownDescription: "when specified as 0, no validation will be done in the packet.", Required: true}}}, "erspan_flow_id": schema.Int64Attribute{MarkdownDescription: "only applicable for 'erspan', A Flow ID of 0 decapsulates all ERSPAN tunnel traffic regardless of Flow ID", Optional: true, Computed: true}, "gmip_port": schema.Int64Attribute{MarkdownDescription: "only applicable for 'gmip', in which case it is required. Specifies the UDP port on which the Tunnel Network port on the receiving GigaVUE H Series is listening. Must match the configuration of the portdst configured on the sending end of the tunnel", Optional: true, Computed: true}, "l2_gre_key": schema.Int64Attribute{MarkdownDescription: "only applicable for 'l2gre', in which case it is required.", Optional: true, Computed: true}, "tls_pcapng": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"decap_key": schema.StringAttribute{Optional: true, Computed: true}, "listener": schema.StringAttribute{Optional: true, Computed: true}}}, "type": schema.StringAttribute{Required: true}, "vxlan": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"port_dst": schema.Int64Attribute{Required: true}, "port_src": schema.Int64Attribute{MarkdownDescription: "when specified as 0, no validation will be done in the packet.", Required: true}, "vni": schema.Int64Attribute{Required: true}}}}}, "tunnel_encap": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART 'Encapsulate Tunnel' Application Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"gmip_config": schema.SingleNestedAttribute{MarkdownDescription: "Configuration for GMIP Tunnel Encapsulate GigaSmaprt App", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"dscp": schema.Int64Attribute{Optional: true, Computed: true}, "dst_ip": schema.StringAttribute{MarkdownDescription: "IP Destination. IPv4 or IPv6.", Required: true}, "dst_port": schema.Int64Attribute{Required: true}, "flow_label": schema.Int64Attribute{Optional: true, Computed: true}, "prec": schema.Int64Attribute{MarkdownDescription: "decimal Precedence value from 0-7 to be used in the ToS byte of the outer headers of tunneled packets", Optional: true, Computed: true}, "src_port": schema.Int64Attribute{Required: true}, "ttl": schema.Int64Attribute{Optional: true, Computed: true}}}, "l2_gre_config": schema.SingleNestedAttribute{MarkdownDescription: "Configuration for l2GRE Tunnel Encapsulate GigaSmaprt App", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"dscp": schema.Int64Attribute{Optional: true, Computed: true}, "dst_ip": schema.StringAttribute{MarkdownDescription: "ip destination, mutually exclusive with 'pgDst'. IPv4 or IPv6.", Optional: true, Computed: true}, "flow_label": schema.Int64Attribute{Optional: true, Computed: true}, "key": schema.Int64Attribute{Required: true}, "pg_dst": schema.StringAttribute{MarkdownDescription: "port group destination alias, mutually exclusive with 'dstIp'", Optional: true, Computed: true}, "prec": schema.Int64Attribute{MarkdownDescription: "decimal Precedence value from 0-7 to be used in the ToS byte of the outer headers of tunneled packets", Optional: true, Computed: true}, "session_field": schema.StringAttribute{MarkdownDescription: "required with stateful loadBalance when 'appType' is 'tunnel'", Optional: true, Computed: true}, "session_pos": schema.StringAttribute{MarkdownDescription: "required if 'sessionField' is specified", Optional: true, Computed: true}, "ttl": schema.Int64Attribute{Optional: true, Computed: true}}}, "tls_pcapng": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"exporter": schema.StringAttribute{Optional: true, Computed: true}, "exporter_group": schema.StringAttribute{Optional: true, Computed: true}}}, "type": schema.StringAttribute{Required: true}, "vxlan_config": schema.SingleNestedAttribute{MarkdownDescription: "Configuration for Vxlan Tunnel Encapsulate GigaSmaprt App", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"dscp": schema.Int64Attribute{Optional: true, Computed: true}, "dst_ip": schema.StringAttribute{MarkdownDescription: "IP Destination. IPv4 or IPv6.", Required: true}, "dst_port": schema.Int64Attribute{Required: true}, "src_port": schema.Int64Attribute{Required: true}, "ttl": schema.Int64Attribute{Optional: true, Computed: true}, "vni": schema.Int64Attribute{Required: true}}}}}}}, "gs_group": schema.StringAttribute{MarkdownDescription: "Alias of referenced managing GsGroup", Required: true}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Optional: true, Computed: true}, "health_state_reasons": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Optional: true, Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Optional: true, Computed: true}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Optional: true, Computed: true}}}}}}}}, "app_filter_config": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"egress_traffic_config": schema.DynamicAttribute{Optional: true, Computed: true}, "sapf_profile": schema.SingleNestedAttribute{MarkdownDescription: "Session-Aware APF", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{Required: true}, "bidi": schema.BoolAttribute{MarkdownDescription: "include reverse traffic for the selected sessionFields. Not applicable for 'mplsLabel', 'gtpuTeid', 'vlanId'", Optional: true, Computed: true}, "buffering": schema.SingleNestedAttribute{MarkdownDescription: "Session-Aware APF Buffereing settings", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"buffer_count_before_match": schema.Int64Attribute{MarkdownDescription: "Maximum number of packets BSAPF will buffer per session before APF match", Optional: true, Computed: true}, "enabled": schema.BoolAttribute{Required: true}, "protocol": schema.StringAttribute{MarkdownDescription: "When buffering is enabled, changing protocol requires reboot before new value takes effect.", Optional: true, Computed: true}}}, "cluster_id": schema.StringAttribute{MarkdownDescription: "id of the defining cluster", Optional: true, Computed: true}, "packet_count": schema.Int64Attribute{MarkdownDescription: "For each session, forward this number of packets. Valid range is 2-100; 0 disables packetCount.", Optional: true, Computed: true}, "session_fields": schema.SetNestedAttribute{MarkdownDescription: "A sessionField cannot contain overlapping session attribute, position pairs (e.g. ipv4-5tuple pos 1 and ipv4-src pos 1 is not allowed). Max number of session fields is 20", Required: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"pos": schema.Int64Attribute{MarkdownDescription: "Value of 1 also an alias for 'outer'. Value of 2 also an alias for 'inner'. Not applicable for 'gtpuTeid'. For 'fiveTuple' 'outer' is not supported", Optional: true, Computed: true}, "type": schema.StringAttribute{Required: true}}}}, "timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Optional: true, Computed: true}}}, "sapf_profile_alias": schema.StringAttribute{Optional: true, Computed: true}}}, "associated_monitor_solution_alias": schema.StringAttribute{MarkdownDescription: "monitor object alias", Optional: true, Computed: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "cluster id for which solution is getting created", Optional: true, Computed: true}, "config_status": schema.StringAttribute{MarkdownDescription: "configuration status", Optional: true, Computed: true}, "delete_monitor_sol": schema.BoolAttribute{Optional: true, Computed: true}, "egress_map_aliases_to_delete": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType}, "exporter_aliases_to_delete": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "health_state_reasons": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Computed: true}}}}, "ingress_map_aliases_to_delete": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType}, "ingress_traffic_configs": schema.DynamicAttribute{MarkdownDescription: "ingress traffic configuration", Optional: true, Computed: true}, "monitor_solution_config": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"action": schema.StringAttribute{Optional: true, Computed: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "cluster id", Optional: true, Computed: true}, "config_status": schema.StringAttribute{Optional: true, Computed: true}, "error_message": schema.StringAttribute{Optional: true, Computed: true}, "exporter_alias": schema.StringAttribute{MarkdownDescription: "exporter alias associated with monitor", Optional: true, Computed: true}, "exporter_config": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{Required: true}, "application_profiles": schema.ListAttribute{MarkdownDescription: "application profile aliases to attach to the exporter", Optional: true, Computed: true, ElementType: types.StringType}, "cef": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"active_timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Optional: true, Computed: true}, "inactive_timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Optional: true, Computed: true}}}, "description": schema.StringAttribute{Optional: true, Computed: true}, "destination": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"dscp": schema.Int64Attribute{Optional: true, Computed: true}, "ipv4_address": schema.StringAttribute{MarkdownDescription: "ipv4 address", Optional: true, Computed: true}, "l4_port_dst": schema.Int64Attribute{Optional: true, Computed: true}, "l4_port_src": schema.Int64Attribute{Optional: true, Computed: true}, "l4_protocol": schema.StringAttribute{Optional: true, Computed: true}, "ttl": schema.Int64Attribute{Optional: true, Computed: true}}}, "max_pkt_size": schema.Int64Attribute{Optional: true, Computed: true}, "mobility_sam": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"encoding": schema.StringAttribute{Optional: true, Computed: true}, "encoding_format": schema.StringAttribute{Optional: true, Computed: true}, "event_enable": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"modify": schema.BoolAttribute{Optional: true, Computed: true}, "update": schema.BoolAttribute{Optional: true, Computed: true}}}, "trigger": schema.StringAttribute{Optional: true, Computed: true}}}, "monitor": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"timeout": schema.Int64Attribute{MarkdownDescription: "how often to export in seconds", Optional: true, Computed: true}}}, "netflow": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"active_timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Optional: true, Computed: true}, "inactive_timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Optional: true, Computed: true}, "template_refresh": schema.Int64Attribute{MarkdownDescription: "template refresh interval in seconds", Optional: true, Computed: true}, "template_type": schema.StringAttribute{Optional: true, Computed: true}, "version": schema.StringAttribute{Optional: true, Computed: true}}}, "snmp": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{MarkdownDescription: "snmp reverse lookup enable/disable", Optional: true, Computed: true}}}, "source": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"ip_interface": schema.StringAttribute{Optional: true, Computed: true}}}, "type": schema.StringAttribute{Optional: true, Computed: true}}}, "gs_group": schema.StringAttribute{MarkdownDescription: "gsgroup alias associated with monitor", Optional: true, Computed: true}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "health_state_reasons": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Computed: true}}}}, "monitor_ip_interface": schema.StringAttribute{MarkdownDescription: "ip interface alias associated with monitor", Optional: true, Computed: true}, "ref_sols": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{Optional: true, Computed: true}, "associated_map": schema.StringAttribute{Optional: true, Computed: true}, "type": schema.StringAttribute{Optional: true, Computed: true}}}}, "solution_alias": schema.StringAttribute{MarkdownDescription: "solution alias", Optional: true, Computed: true}, "vport_alias": schema.StringAttribute{MarkdownDescription: "vport alias associated with monitor", Optional: true, Computed: true}, "vport_config": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART vPort", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{Required: true}, "deferred_binding": schema.BoolAttribute{MarkdownDescription: "enable/disable deferred-binding", Optional: true, Computed: true}, "fail_over_action": schema.StringAttribute{Optional: true, Computed: true}, "gs_group": schema.StringAttribute{MarkdownDescription: "Alias of referenced managing GsGroup", Required: true}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "health_state_reasons": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Computed: true}}}}, "inline_status": schema.StringAttribute{Optional: true, Computed: true}, "inner_traffic_path": schema.StringAttribute{MarkdownDescription: "Similar to inline-network traffic-path, applicable for inner map", Optional: true, Computed: true}, "metadata_monitoring": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"action": schema.StringAttribute{MarkdownDescription: "metadata monitoring action", Optional: true, Computed: true}, "exporters": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType}}}, "mode": schema.StringAttribute{Optional: true, Computed: true}, "outer_traffic_path": schema.StringAttribute{MarkdownDescription: "Similar to inline-tool/inline-tool-group flex-traffic path, applicable for outer map", Optional: true, Computed: true}, "sa_apf_profile": schema.StringAttribute{MarkdownDescription: "ASF session profile", Optional: true, Computed: true}}}}}, "solution_alias": schema.StringAttribute{MarkdownDescription: "user defined application visibility solution alias", Optional: true, Computed: true}, "solution_desc": schema.StringAttribute{MarkdownDescription: "user defined application visibility solution description", Optional: true, Computed: true}, "solution_status": schema.StringAttribute{MarkdownDescription: "solution status", Optional: true, Computed: true}, "solution_type": schema.StringAttribute{MarkdownDescription: "user intent solution type", Optional: true, Computed: true}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *SolutionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SolutionResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.createRemote(ctx, &plan, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// createRemote performs the create HTTP exchange and decodes the response into plan. Extracted from Create so the request/response logic is unit-testable without a tfsdk.Plan.
func (r *SolutionResource) createRemote(ctx context.Context, plan *SolutionResourceModel, resp *resource.CreateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_solution", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_solution", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/appsVisibility/solutions"
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_solution", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_solution", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error creating gigavuecore_solution", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error creating gigavuecore_solution", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error creating gigavuecore_solution", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error creating gigavuecore_solution", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error creating gigavuecore_solution", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error creating gigavuecore_solution", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error creating gigavuecore_solution", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error creating gigavuecore_solution", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error creating gigavuecore_solution", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error creating gigavuecore_solution", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_solution", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
	if plan.SolutionAlias.IsNull() || plan.SolutionAlias.IsUnknown() {
		loc := httpResp.Header.Get("Location")
		if loc != "" {
			plan.SolutionAlias = types.StringValue(loc)
		} else {
			resp.Diagnostics.AddError("Error creating gigavuecore_solution", "The create response did not contain an identifier and no Location header was returned, so the resource cannot be tracked in state.")
			return
		}
	}
}

// Read refreshes the Terraform state with the latest remote values.
func (r *SolutionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SolutionResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
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
func (r *SolutionResource) readRemote(ctx context.Context, state *SolutionResourceModel, resp *resource.ReadResponse) (removed bool) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/appsVisibility/solutions/{solutionAlias}"
	reqPath = strings.ReplaceAll(reqPath, "{solutionAlias}", url.PathEscape(state.SolutionAlias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_solution", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_solution", fmt.Sprintf("Could not send request: %s", err))
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
			resp.Diagnostics.AddError("Error reading gigavuecore_solution", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_solution", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_solution", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_solution", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_solution", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_solution", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_solution", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_solution", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_solution", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_solution", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["appVzbilitySolution"]; ok {
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
		resp.Diagnostics.AddError("Error reading gigavuecore_solution", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
	return
}

// Update modifies the remote resource to match the desired plan.
func (r *SolutionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan SolutionResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state SolutionResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if plan.SolutionAlias.IsNull() || plan.SolutionAlias.IsUnknown() {
		if !state.SolutionAlias.IsNull() && !state.SolutionAlias.IsUnknown() {
			plan.SolutionAlias = state.SolutionAlias
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
func (r *SolutionResource) updateRemote(ctx context.Context, plan *SolutionResourceModel, resp *resource.UpdateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_solution", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_solution", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/appsVisibility/solutions/{solutionAlias}"
	reqPath = strings.ReplaceAll(reqPath, "{solutionAlias}", url.PathEscape(plan.SolutionAlias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodPatch, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_solution", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_solution", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error updating gigavuecore_solution", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error updating gigavuecore_solution", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error updating gigavuecore_solution", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error updating gigavuecore_solution", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error updating gigavuecore_solution", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error updating gigavuecore_solution", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error updating gigavuecore_solution", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error updating gigavuecore_solution", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error updating gigavuecore_solution", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error updating gigavuecore_solution", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_solution", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Delete destroys the remote resource.
func (r *SolutionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SolutionResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.deleteRemote(ctx, &state, resp)
}

// deleteRemote performs the delete HTTP exchange, treating a 404 as already deleted. Extracted from Delete so the request/response logic is unit-testable without a tfsdk.State.
func (r *SolutionResource) deleteRemote(ctx context.Context, state *SolutionResourceModel, resp *resource.DeleteResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/appsVisibility/solutions/{solutionAlias}"
	reqPath = strings.ReplaceAll(reqPath, "{solutionAlias}", url.PathEscape(state.SolutionAlias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_solution", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_solution", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		return
	}
	if !(httpResp.StatusCode == 204) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error deleting gigavuecore_solution", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error deleting gigavuecore_solution", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error deleting gigavuecore_solution", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error deleting gigavuecore_solution", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error deleting gigavuecore_solution", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error deleting gigavuecore_solution", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error deleting gigavuecore_solution", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error deleting gigavuecore_solution", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error deleting gigavuecore_solution", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *SolutionResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *SolutionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("solution_alias"), req.ID)...)
}
