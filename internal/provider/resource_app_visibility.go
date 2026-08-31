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
	_ resource.Resource                = (*AppVisibilityResource)(nil)
	_ resource.ResourceWithIdentity    = (*AppVisibilityResource)(nil)
	_ resource.ResourceWithImportState = (*AppVisibilityResource)(nil)
	_ resource.ResourceWithConfigure   = (*AppVisibilityResource)(nil)
)

// AppVisibilityResource is the generated Terraform managed resource implementation.
type AppVisibilityResource struct {
	client *client.Client
}

// AppVisibilityResourceModel describes the Terraform state and plan shape for AppVisibilityResource.
type AppVisibilityResourceModel struct {
	AppEnvId                  types.String   `tfsdk:"app_env_id" json:"appEnvId"`
	AppExportConfig           types.Object   `tfsdk:"app_export_config" json:"appExportConfig"`
	AppExporterConfig         types.Object   `tfsdk:"app_exporter_config" json:"appExporterConfig"`
	AppFilterConfig           types.Object   `tfsdk:"app_filter_config" json:"appFilterConfig"`
	ConfigStatus              types.String   `tfsdk:"config_status" json:"configStatus"`
	ConfigStatusReasons       types.String   `tfsdk:"config_status_reasons" json:"configStatusReasons"`
	ConnId                    types.String   `tfsdk:"conn_id" json:"connId"`
	Dedup                     types.Object   `tfsdk:"dedup"`
	DistributeTraffic         types.Bool     `tfsdk:"distribute_traffic" json:"distributeTraffic"`
	DynamicScaleUnit          types.Bool     `tfsdk:"dynamic_scale_unit" json:"dynamicScaleUnit"`
	EnvId                     types.String   `tfsdk:"env_id" json:"envId"`
	HealthState               types.String   `tfsdk:"health_state" json:"healthState"`
	HealthStateReasons        types.List     `tfsdk:"health_state_reasons" json:"healthStateReasons"`
	IngressTrafficConfigs     types.List     `tfsdk:"ingress_traffic_configs" json:"ingressTrafficConfigs"`
	MonitorSolutionConfig     types.Object   `tfsdk:"monitor_solution_config" json:"monitorSolutionConfig"`
	ScaleUnit                 types.Int64    `tfsdk:"scale_unit" json:"scaleUnit"`
	SolutionAlias             types.String   `tfsdk:"solution_alias" json:"solutionAlias"`
	SolutionCreatedTimestamp  types.String   `tfsdk:"solution_created_timestamp" json:"solutionCreatedTimestamp"`
	SolutionDesc              types.String   `tfsdk:"solution_desc" json:"solutionDesc"`
	SolutionModifiedTimestamp types.String   `tfsdk:"solution_modified_timestamp" json:"solutionModifiedTimestamp"`
	TrafficPolicyGraphAlias   types.String   `tfsdk:"traffic_policy_graph_alias" json:"trafficPolicyGraphAlias"`
	Timeouts                  timeouts.Value `tfsdk:"timeouts"`
}

// Metadata returns the resource type name.
func (r *AppVisibilityResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_app_visibility"
}

// Schema returns the Terraform schema for this resource.
func (r *AppVisibilityResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load an Application Intelligence Solution by alias", Attributes: map[string]schema.Attribute{"app_env_id": schema.StringAttribute{MarkdownDescription: "ID of the Application Environment", Computed: true}, "app_export_config": schema.SingleNestedAttribute{MarkdownDescription: "Application Metadata Configuration", Optional: true, Attributes: map[string]schema.Attribute{"cache_config": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"advance_hash": schema.BoolAttribute{MarkdownDescription: "When scheduling flows for processing, use external encapsulate packet header if disable, use inner packet header if enable.", Optional: true}, "alias": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "description": schema.StringAttribute{Optional: true}, "dpi_inject_limit": schema.Int64Attribute{Optional: true}, "event": schema.StringAttribute{Optional: true, Validators: []validator.String{stringvalidator.OneOf("txnEnd", "none")}}, "exporters": schema.ListAttribute{MarkdownDescription: "alias of metadata exporters to attach this cache", Optional: true, ElementType: types.StringType}, "flow_behavior": schema.StringAttribute{MarkdownDescription: "direction for flow identification", Optional: true, Validators: []validator.String{stringvalidator.OneOf("unidir", "bidir")}}, "match": schema.SingleNestedAttribute{MarkdownDescription: "match criteria for record generation", Optional: true, Attributes: map[string]schema.Attribute{"datalink": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"mac_dst": schema.BoolAttribute{Optional: true}, "mac_src": schema.BoolAttribute{Optional: true}, "vlan": schema.BoolAttribute{Optional: true}}}, "interface": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"in_name_width": schema.Int64Attribute{Optional: true, Validators: []validator.Int64{int64validator.Between(1, 32)}}, "in_physical_width": schema.Int64Attribute{Optional: true, Validators: []validator.Int64{int64validator.OneOf(2, 4)}}}}, "ip": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"version": schema.BoolAttribute{Optional: true}}}, "ipv4": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"destination": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"prefix_min_mask": schema.StringAttribute{MarkdownDescription: "ipv4 destination prefix minimum-mask - netmask or mask length", Optional: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}}}, "dscp": schema.BoolAttribute{Optional: true}, "fragmentation": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"flags": schema.BoolAttribute{Optional: true}, "offset": schema.BoolAttribute{Optional: true}}}, "header_len": schema.BoolAttribute{Optional: true}, "option_map": schema.BoolAttribute{Optional: true}, "precedence": schema.BoolAttribute{Optional: true}, "protocol": schema.BoolAttribute{Optional: true}, "section": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"header_size": schema.Int64Attribute{Optional: true, Validators: []validator.Int64{int64validator.Between(1, 128)}}, "payload_size": schema.Int64Attribute{Optional: true, Validators: []validator.Int64{int64validator.Between(1, 128)}}}}, "source": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"prefix_min_mask": schema.StringAttribute{MarkdownDescription: "ipv4 source prefix minimum-mask - netmask or mask length", Optional: true}}}, "tos": schema.BoolAttribute{Optional: true}, "total_length": schema.BoolAttribute{Optional: true}, "ttl": schema.BoolAttribute{Optional: true}}}, "ipv6": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"destination": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"prefix_min_mask": schema.StringAttribute{Optional: true}}}, "dscp": schema.BoolAttribute{Optional: true}, "extension_map": schema.BoolAttribute{Optional: true}, "flow_label": schema.BoolAttribute{Optional: true}, "fragmentation": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"flags": schema.BoolAttribute{Optional: true}, "offset": schema.BoolAttribute{Optional: true}}}, "hop_limit": schema.BoolAttribute{Optional: true}, "length": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"header": schema.BoolAttribute{Optional: true}, "payload": schema.BoolAttribute{Optional: true}, "total": schema.BoolAttribute{Optional: true}}}, "next_header": schema.BoolAttribute{Optional: true}, "precedence": schema.BoolAttribute{Optional: true}, "section": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"header_size": schema.Int64Attribute{Optional: true, Validators: []validator.Int64{int64validator.Between(1, 128)}}, "payload_size": schema.Int64Attribute{Optional: true, Validators: []validator.Int64{int64validator.Between(1, 128)}}}}, "source": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"prefix_min_mask": schema.StringAttribute{MarkdownDescription: "ipv6 source prefix minimum-mask - netmask or mask length", Optional: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}}}, "traffic_class": schema.BoolAttribute{Optional: true}}}, "transport": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"dst_port": schema.BoolAttribute{Optional: true}, "icmp": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"ipv4_code": schema.BoolAttribute{Optional: true}, "ipv4_type": schema.BoolAttribute{Optional: true}, "ipv6_code": schema.BoolAttribute{Optional: true}, "ipv6_type": schema.BoolAttribute{Optional: true}}}, "src_port": schema.BoolAttribute{Optional: true}, "tcp": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"ack_number": schema.BoolAttribute{Optional: true}, "dst_port": schema.BoolAttribute{Optional: true}, "flags": schema.BoolAttribute{Optional: true}, "header_len": schema.BoolAttribute{Optional: true}, "seq_number": schema.BoolAttribute{Optional: true}, "src_port": schema.BoolAttribute{Optional: true}, "urgent_ptr": schema.BoolAttribute{Optional: true}, "window_size": schema.BoolAttribute{Optional: true}}}, "udp": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"dst_port": schema.BoolAttribute{Optional: true}, "msg_len": schema.BoolAttribute{Optional: true}, "src_port": schema.BoolAttribute{Optional: true}}}}}}}, "multi_collect": schema.BoolAttribute{MarkdownDescription: "Collect all attributes as it is discovered when enable. It will export the same record once when disable.", Optional: true}, "network_profiles": schema.ListAttribute{MarkdownDescription: "alias of metadata network profiles to attach this cache", Optional: true, ElementType: types.StringType}, "observation_domain_id": schema.Int64Attribute{Optional: true}, "sampling": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"mode": schema.StringAttribute{Optional: true, Validators: []validator.String{stringvalidator.OneOf("multiRate", "noSampling", "singleRate")}}, "single_sampling_rate": schema.Int64Attribute{MarkdownDescription: "Packet interval window size. Valid values: 10-16000 (in packets)", Optional: true, Validators: []validator.Int64{int64validator.Between(10, 16000)}}}}, "size": schema.SingleNestedAttribute{MarkdownDescription: "size of the flows", Optional: true, Attributes: map[string]schema.Attribute{"flows": schema.Int64Attribute{MarkdownDescription: "size of flows in millions", Optional: true, Validators: []validator.Int64{int64validator.Between(1, 4)}}}}, "timeout": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"idle": schema.Int64Attribute{MarkdownDescription: "idle timeout in seconds. max value 7days. default 30 min", Optional: true, Validators: []validator.Int64{int64validator.Between(1, 604800)}}}}}}, "destination_configs": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"application_names": schema.ListNestedAttribute{Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"attributes": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"name": schema.StringAttribute{MarkdownDescription: "attribute name", Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "value": schema.StringAttribute{MarkdownDescription: "application's attribute value", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}}}}, "is_user_defined": schema.BoolAttribute{MarkdownDescription: "Default value is false. Must be set to true only for App Intel solution while configuring user defined apps", Optional: true, Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "application name", Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}}}}, "destination_name": schema.StringAttribute{Optional: true, Computed: true}, "export_meta_app_profile": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "application profile alias", Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "application_id": schema.BoolAttribute{MarkdownDescription: "only valid with 'export' type", Optional: true}, "applications": schema.ListNestedAttribute{MarkdownDescription: "application and attributes.", Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"attributes": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"name": schema.StringAttribute{MarkdownDescription: "attribute name", Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "value": schema.StringAttribute{MarkdownDescription: "application's attribute value", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}}}}, "is_user_defined": schema.BoolAttribute{MarkdownDescription: "Default value is false. Must be set to true only for App Intel solution while configuring user defined apps", Optional: true, Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "application name", Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}}}}, "counter": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"bytes": schema.BoolAttribute{Optional: true}, "bytes_long": schema.BoolAttribute{Optional: true}, "inner_byte": schema.BoolAttribute{Optional: true}, "inner_byte_long": schema.BoolAttribute{Optional: true}, "packets": schema.BoolAttribute{Optional: true}, "packets_long": schema.BoolAttribute{Optional: true}}}, "datalink": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"mac_dst": schema.BoolAttribute{Optional: true}, "mac_src": schema.BoolAttribute{Optional: true}, "vlan": schema.BoolAttribute{Optional: true}}}, "description": schema.StringAttribute{Optional: true}, "flow": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"end_reason": schema.BoolAttribute{Optional: true}}}, "gtpu": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"qfi": schema.BoolAttribute{Optional: true}, "teid": schema.BoolAttribute{Optional: true}}}, "interface": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"in_name_width": schema.Int64Attribute{Optional: true, Validators: []validator.Int64{int64validator.Between(1, 32)}}, "in_physical_width": schema.Int64Attribute{Optional: true, Validators: []validator.Int64{int64validator.OneOf(2, 4)}}, "out_physical_width": schema.Int64Attribute{Optional: true, Validators: []validator.Int64{int64validator.OneOf(2, 4)}}}}, "ip": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"version": schema.BoolAttribute{Optional: true}}}, "ipv4": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"destination": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"prefix_min_mask": schema.StringAttribute{Optional: true}}}, "dscp": schema.BoolAttribute{Optional: true}, "fragmentation": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"flags": schema.BoolAttribute{Optional: true}, "offset": schema.BoolAttribute{Optional: true}}}, "header_len": schema.BoolAttribute{Optional: true}, "option_map": schema.BoolAttribute{Optional: true}, "precedence": schema.BoolAttribute{Optional: true}, "protocol": schema.BoolAttribute{Optional: true}, "section": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"header_size": schema.Int64Attribute{Optional: true, Validators: []validator.Int64{int64validator.Between(1, 128)}}, "payload_size": schema.Int64Attribute{Optional: true, Validators: []validator.Int64{int64validator.Between(1, 128)}}}}, "source": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"prefix_min_mask": schema.StringAttribute{MarkdownDescription: "ipv4 source prefix minimum-mask - netmask or mask length", Optional: true}}}, "tos": schema.BoolAttribute{Optional: true}, "total_length": schema.BoolAttribute{Optional: true}, "ttl": schema.BoolAttribute{Optional: true}}}, "ipv6": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"destination": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"prefix_min_mask": schema.StringAttribute{Optional: true}}}, "dscp": schema.BoolAttribute{Optional: true}, "extension_map": schema.BoolAttribute{Optional: true}, "flow_label": schema.BoolAttribute{Optional: true}, "fragmentation": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"flags": schema.BoolAttribute{Optional: true}, "offset": schema.BoolAttribute{Optional: true}}}, "hop_limit": schema.BoolAttribute{Optional: true}, "length": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"header": schema.BoolAttribute{Optional: true}, "payload": schema.BoolAttribute{Optional: true}, "total": schema.BoolAttribute{Optional: true}}}, "next_header": schema.BoolAttribute{Optional: true}, "precedence": schema.BoolAttribute{Optional: true}, "section": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"header_size": schema.Int64Attribute{Optional: true, Validators: []validator.Int64{int64validator.Between(1, 128)}}, "payload_size": schema.Int64Attribute{Optional: true, Validators: []validator.Int64{int64validator.Between(1, 128)}}}}, "source": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"prefix_min_mask": schema.StringAttribute{Optional: true}}}, "traffic_class": schema.BoolAttribute{Optional: true}}}, "outer_ipv4": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"destination": schema.BoolAttribute{Optional: true}, "source": schema.BoolAttribute{Optional: true}}}, "outer_ipv6": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"destination": schema.BoolAttribute{Optional: true}, "source": schema.BoolAttribute{Optional: true}}}, "timestamp": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"flow_end_msec": schema.BoolAttribute{Optional: true}, "flow_endsec": schema.BoolAttribute{Optional: true}, "flow_start_msec": schema.BoolAttribute{Optional: true}, "flow_startsec": schema.BoolAttribute{Optional: true}, "sys_up_time_first": schema.BoolAttribute{Optional: true}, "sys_up_time_last": schema.BoolAttribute{Optional: true}}}, "transport": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"dst_port": schema.BoolAttribute{Optional: true}, "icmp": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"ipv4_code": schema.BoolAttribute{Optional: true}, "ipv4_type": schema.BoolAttribute{Optional: true}, "ipv6_code": schema.BoolAttribute{Optional: true}, "ipv6_type": schema.BoolAttribute{Optional: true}}}, "src_port": schema.BoolAttribute{Optional: true}, "tcp": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"ack_number": schema.BoolAttribute{Optional: true}, "dst_port": schema.BoolAttribute{Optional: true}, "flags": schema.BoolAttribute{Optional: true}, "header_len": schema.BoolAttribute{Optional: true}, "seq_number": schema.BoolAttribute{Optional: true}, "src_port": schema.BoolAttribute{Optional: true}, "urgent_ptr": schema.BoolAttribute{Optional: true}, "window_size": schema.BoolAttribute{Optional: true}}}, "udp": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"dst_port": schema.BoolAttribute{Optional: true}, "msg_len": schema.BoolAttribute{Optional: true}, "src_port": schema.BoolAttribute{Optional: true}}}}}, "type": schema.StringAttribute{Optional: true, Validators: []validator.String{stringvalidator.OneOf("export", "filter")}}}}, "export_meta_app_profile_alias": schema.StringAttribute{Optional: true}, "exporter_alias": schema.StringAttribute{Optional: true, Computed: true}, "exporter_config": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "application_profiles": schema.ListAttribute{MarkdownDescription: "application profile aliases to attach to the exporter", Optional: true, ElementType: types.StringType}, "cef": schema.SingleNestedAttribute{MarkdownDescription: "cef attributes", Optional: true, Attributes: map[string]schema.Attribute{"active_timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Optional: true, Validators: []validator.Int64{int64validator.Between(1, 604800)}}, "inactive_timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Optional: true, Validators: []validator.Int64{int64validator.Between(1, 604800)}}}}, "description": schema.StringAttribute{Optional: true}, "destination": schema.SingleNestedAttribute{MarkdownDescription: "destination attributes", Optional: true, Attributes: map[string]schema.Attribute{"dscp": schema.Int64Attribute{Optional: true, Validators: []validator.Int64{int64validator.Between(0, 63)}}, "ipv4_address": schema.StringAttribute{MarkdownDescription: "ipv4 address", Optional: true}, "l4_port_dst": schema.Int64Attribute{Optional: true, Validators: []validator.Int64{int64validator.Between(1, 65535)}}, "l4_port_src": schema.Int64Attribute{Optional: true, Validators: []validator.Int64{int64validator.Between(1, 65535)}}, "l4_protocol": schema.StringAttribute{Optional: true, Validators: []validator.String{stringvalidator.OneOf("udp", "tcp")}}, "ttl": schema.Int64Attribute{Optional: true, Validators: []validator.Int64{int64validator.Between(1, 255)}}}}, "max_pkt_size": schema.Int64Attribute{Optional: true}, "mobility_sam": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"encoding": schema.StringAttribute{Optional: true}, "encoding_format": schema.StringAttribute{Optional: true, Validators: []validator.String{stringvalidator.OneOf("hierarchy", "flat")}}, "event_enable": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"modify": schema.BoolAttribute{Optional: true}, "update": schema.BoolAttribute{Optional: true}}}, "trigger": schema.StringAttribute{Optional: true}}}, "monitor": schema.SingleNestedAttribute{MarkdownDescription: "monitor attributes", Optional: true, Attributes: map[string]schema.Attribute{"timeout": schema.Int64Attribute{MarkdownDescription: "how often to export in seconds", Optional: true, Validators: []validator.Int64{int64validator.Between(60, 900)}}}}, "netflow": schema.SingleNestedAttribute{MarkdownDescription: "netflow attributes", Optional: true, Attributes: map[string]schema.Attribute{"active_timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Optional: true, Validators: []validator.Int64{int64validator.Between(1, 604800)}}, "inactive_timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Optional: true, Validators: []validator.Int64{int64validator.Between(1, 604800)}}, "template_refresh": schema.Int64Attribute{MarkdownDescription: "template refresh interval in seconds", Optional: true, Validators: []validator.Int64{int64validator.Between(1, 216000)}}, "template_type": schema.StringAttribute{Optional: true, Validators: []validator.String{stringvalidator.OneOf("cohesive", "segregated")}}, "version": schema.StringAttribute{Optional: true, Validators: []validator.String{stringvalidator.OneOf("v5", "v9", "ipfix")}}}}, "snmp": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{MarkdownDescription: "snmp reverse lookup enable/disable", Optional: true, Computed: true}}}, "source": schema.SingleNestedAttribute{MarkdownDescription: "source tunnel port", Optional: true, Attributes: map[string]schema.Attribute{"ip_interface": schema.StringAttribute{Optional: true}}}, "type": schema.StringAttribute{Optional: true, Validators: []validator.String{stringvalidator.OneOf("cef", "netflow", "monitor")}}}}, "iface": schema.StringAttribute{MarkdownDescription: "Interface Mapping for Egress Tunnel", Optional: true, Computed: true}}}}}}, "app_exporter_config": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"app_instance_name": schema.StringAttribute{MarkdownDescription: "The name of the AppMetadata app instance", Computed: true}, "app_metadata_tiered_batch_id": schema.StringAttribute{MarkdownDescription: "ID of the appExporterConfig tiered batch", Computed: true}, "cache_config_alias": schema.StringAttribute{MarkdownDescription: "Alias of the cacheConfig", Computed: true}, "destination_configs": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"destination_name": schema.StringAttribute{MarkdownDescription: "Alias of the destination tool", Computed: true}, "exporter_alias": schema.StringAttribute{MarkdownDescription: "Alias of the exporter", Computed: true}, "iface": schema.StringAttribute{MarkdownDescription: "Interface Mapping for Egress Tunnel", Computed: true}, "template_name": schema.StringAttribute{MarkdownDescription: "Name of the App Profile Template Used", Computed: true}}}}}}, "app_filter_config": schema.SingleNestedAttribute{MarkdownDescription: "Application Filtering Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"app_filter_tiered_batch_id": schema.StringAttribute{MarkdownDescription: "ID of the appFilterConfig tiered batch", Computed: true}, "egress_traffic_configs": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"drop_application_names": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"attributes": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"name": schema.StringAttribute{MarkdownDescription: "attribute name", Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "value": schema.StringAttribute{MarkdownDescription: "application's attribute value", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}}}}, "is_user_defined": schema.BoolAttribute{MarkdownDescription: "Default value is false. Must be set to true only for App Intel solution while configuring user defined apps", Optional: true, Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "application name", Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}}}}, "pass_application_names": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"attributes": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"name": schema.StringAttribute{MarkdownDescription: "attribute name", Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "value": schema.StringAttribute{MarkdownDescription: "application's attribute value", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}}}}, "is_user_defined": schema.BoolAttribute{MarkdownDescription: "Default value is false. Must be set to true only for App Intel solution while configuring user defined apps", Optional: true, Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "application name", Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}}}}, "priority": schema.Int64Attribute{Computed: true}, "tunnel_aliases": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType}}}}, "map_alias": schema.StringAttribute{MarkdownDescription: "Alias of the appFiltering map", Computed: true}, "sapf_profile_alias": schema.StringAttribute{MarkdownDescription: "Alias of the SAPF Profile", Computed: true}}}, "config_status": schema.StringAttribute{MarkdownDescription: "configuration status", Computed: true}, "config_status_reasons": schema.StringAttribute{MarkdownDescription: "configuration status reasons", Computed: true}, "conn_id": schema.StringAttribute{MarkdownDescription: "ID of the Connection Domain", Required: true}, "dedup": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"app_config_id": schema.StringAttribute{Computed: true}, "app_instance_name": schema.StringAttribute{Computed: true}, "dedup_tiered_batch_id": schema.StringAttribute{Computed: true}, "enabled": schema.BoolAttribute{Computed: true}, "gs_params_name": schema.StringAttribute{Computed: true}}}, "distribute_traffic": schema.BoolAttribute{MarkdownDescription: "Indicates Traffic Distribution enabled or not", Optional: true}, "dynamic_scale_unit": schema.BoolAttribute{MarkdownDescription: "When set as true, FM automatically sets the optimal scaleUnit", Optional: true, Computed: true}, "env_id": schema.StringAttribute{MarkdownDescription: "ID of the Environment", Required: true}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "health_state_reasons": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Computed: true}}}}, "ingress_traffic_configs": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"source_selector_aliases": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType}, "src_raw_end_point_interfaces": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType}, "tunnel_aliases": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType}, "tunnel_interface_mappings": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"iface": schema.StringAttribute{Optional: true, Computed: true}, "tunnel_alias": schema.StringAttribute{Optional: true, Computed: true}}}}}}}, "monitor_solution_config": schema.SingleNestedAttribute{MarkdownDescription: "Application Monitoring configuration", Required: true, Attributes: map[string]schema.Attribute{"app_instance_config_id": schema.StringAttribute{Computed: true}, "app_instance_name": schema.StringAttribute{Computed: true}, "app_viz_tiered_batch_id": schema.StringAttribute{MarkdownDescription: "ID of the monitoringSolutionConfig tiered batch", Computed: true}, "destination_config": schema.SingleNestedAttribute{Required: true, Attributes: map[string]schema.Attribute{"destination_name": schema.StringAttribute{Computed: true}, "exporter_alias": schema.StringAttribute{Computed: true}, "iface": schema.StringAttribute{MarkdownDescription: "Interface Mapping for Egress Tunnel", Optional: true, Computed: true}}}, "gs_param_config_id": schema.StringAttribute{MarkdownDescription: "Gs param name on which the solution is applied", Computed: true}, "mgmt_interface": schema.StringAttribute{MarkdownDescription: "Pass this value if you would like to use the management interface of the Node to export App Monitoring metadata to FM", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("internal")}}}}, "scale_unit": schema.Int64Attribute{MarkdownDescription: "Number of units of memory needed for each Application", Optional: true, Computed: true}, "solution_alias": schema.StringAttribute{MarkdownDescription: "Alias of the solution", Required: true}, "solution_created_timestamp": schema.StringAttribute{Computed: true}, "solution_desc": schema.StringAttribute{MarkdownDescription: "Description of the solution", Optional: true, Computed: true}, "solution_modified_timestamp": schema.StringAttribute{Computed: true}, "traffic_policy_graph_alias": schema.StringAttribute{MarkdownDescription: "Alias of the Traffic Policy Graph", Computed: true}}, Blocks: map[string]schema.Block{"timeouts": timeouts.Block(ctx, timeouts.Opts{Create: true, Read: true, Update: true, Delete: true})}}
}

// IdentitySchema returns the resource identity schema shared with the paired list resource.
func (r *AppVisibilityResource) IdentitySchema(_ context.Context, _ resource.IdentitySchemaRequest, resp *resource.IdentitySchemaResponse) {
	resp.IdentitySchema = identityschema.Schema{Attributes: map[string]identityschema.Attribute{"solution_alias": identityschema.StringAttribute{RequiredForImport: true, Description: "Alias of the solution"}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *AppVisibilityResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan AppVisibilityResourceModel
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
	resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("solution_alias"), plan.SolutionAlias)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// createRemote performs the create HTTP exchange and decodes the response into plan. Extracted from Create so the request/response logic is unit-testable without a tfsdk.Plan.
func (r *AppVisibilityResource) createRemote(ctx context.Context, plan *AppVisibilityResourceModel, resp *resource.CreateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_app_visibility", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_app_visibility", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/intent/appVisibility"
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_app_visibility", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_app_visibility", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error creating gigavuecore_app_visibility", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error creating gigavuecore_app_visibility", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error creating gigavuecore_app_visibility", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error creating gigavuecore_app_visibility", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error creating gigavuecore_app_visibility", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error creating gigavuecore_app_visibility", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error creating gigavuecore_app_visibility", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error creating gigavuecore_app_visibility", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error creating gigavuecore_app_visibility", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error creating gigavuecore_app_visibility", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_app_visibility", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
	if plan.SolutionAlias.IsNull() || plan.SolutionAlias.IsUnknown() {
		loc := httpResp.Header.Get("Location")
		if loc != "" {
			loc = strings.TrimRight(loc, "/")
			i := strings.LastIndex(loc, "/")
			if i >= 0 {
				loc = loc[i+1:]
			}
			plan.SolutionAlias = types.StringValue(loc)
		} else {
			resp.Diagnostics.AddError("Error creating gigavuecore_app_visibility", "The create response did not contain an identifier and no Location header was returned, so the resource cannot be tracked in state.")
			return
		}
	}
}

// Read refreshes the Terraform state with the latest remote values.
func (r *AppVisibilityResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state AppVisibilityResourceModel
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
	resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("solution_alias"), state.SolutionAlias)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// readRemote performs the read HTTP exchange and decodes the response into state, returning removed=true when the API reports 404. Extracted from Read so the request/response logic is unit-testable without a tfsdk.State.
func (r *AppVisibilityResource) readRemote(ctx context.Context, state *AppVisibilityResourceModel, resp *resource.ReadResponse) (removed bool) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/intent/appVisibility/{solutionAlias}"
	reqPath = strings.ReplaceAll(reqPath, "{solutionAlias}", url.PathEscape(state.SolutionAlias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_app_visibility", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_app_visibility", fmt.Sprintf("Could not send request: %s", err))
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
			resp.Diagnostics.AddError("Error reading gigavuecore_app_visibility", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_app_visibility", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_app_visibility", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_app_visibility", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_app_visibility", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_app_visibility", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_app_visibility", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_app_visibility", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_app_visibility", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_app_visibility", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["appsVisibilitySolution"]; ok {
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
		resp.Diagnostics.AddError("Error reading gigavuecore_app_visibility", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
	return
}

// Update modifies the remote resource to match the desired plan.
func (r *AppVisibilityResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan AppVisibilityResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state AppVisibilityResourceModel
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
	resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("solution_alias"), plan.SolutionAlias)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// updateRemote performs the update HTTP exchange and decodes the response into plan. Extracted from Update so the request/response logic is unit-testable without a tfsdk.Plan.
func (r *AppVisibilityResource) updateRemote(ctx context.Context, plan *AppVisibilityResourceModel, resp *resource.UpdateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_app_visibility", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_app_visibility", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/intent/appVisibility/{solutionAlias}"
	reqPath = strings.ReplaceAll(reqPath, "{solutionAlias}", url.PathEscape(plan.SolutionAlias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodPatch, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_app_visibility", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_app_visibility", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error updating gigavuecore_app_visibility", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error updating gigavuecore_app_visibility", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error updating gigavuecore_app_visibility", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error updating gigavuecore_app_visibility", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error updating gigavuecore_app_visibility", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error updating gigavuecore_app_visibility", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error updating gigavuecore_app_visibility", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error updating gigavuecore_app_visibility", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error updating gigavuecore_app_visibility", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error updating gigavuecore_app_visibility", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_app_visibility", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Delete destroys the remote resource.
func (r *AppVisibilityResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state AppVisibilityResourceModel
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
func (r *AppVisibilityResource) deleteRemote(ctx context.Context, state *AppVisibilityResourceModel, resp *resource.DeleteResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/intent/appVisibility/{solutionAlias}"
	reqPath = strings.ReplaceAll(reqPath, "{solutionAlias}", url.PathEscape(state.SolutionAlias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_app_visibility", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_app_visibility", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		return
	}
	if !(httpResp.StatusCode == 204) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error deleting gigavuecore_app_visibility", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error deleting gigavuecore_app_visibility", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error deleting gigavuecore_app_visibility", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error deleting gigavuecore_app_visibility", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error deleting gigavuecore_app_visibility", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error deleting gigavuecore_app_visibility", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error deleting gigavuecore_app_visibility", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error deleting gigavuecore_app_visibility", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error deleting gigavuecore_app_visibility", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *AppVisibilityResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *AppVisibilityResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("solution_alias"), req.ID)...)
}
