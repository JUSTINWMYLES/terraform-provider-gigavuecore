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
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	timeouts "github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	int64validator "github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	listvalidator "github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	stringvalidator "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	path "github.com/hashicorp/terraform-plugin-framework/path"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	validator "github.com/hashicorp/terraform-plugin-framework/schema/validator"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertions.
var (
	_ resource.Resource                = (*MapResource)(nil)
	_ resource.ResourceWithImportState = (*MapResource)(nil)
	_ resource.ResourceWithConfigure   = (*MapResource)(nil)
)

// MapResource is the generated Terraform managed resource implementation.
type MapResource struct {
	client *client.Client
}

// MapResourceModel describes the Terraform state and plan shape for MapResource.
type MapResourceModel struct {
	Alias                       types.String   `tfsdk:"alias"`
	ApRules                     types.Object   `tfsdk:"ap_rules" json:"apRules"`
	ClusterId                   types.String   `tfsdk:"cluster_id" json:"clusterId"`
	Comment                     types.String   `tfsdk:"comment"`
	DstPorts                    types.List     `tfsdk:"dst_ports" json:"dstPorts"`
	EgressGigastream            types.List     `tfsdk:"egress_gigastream" json:"egressGigastream"`
	Enable                      types.Bool     `tfsdk:"enable"`
	EncapTunnel                 types.String   `tfsdk:"encap_tunnel" json:"encapTunnel"`
	FlexInline                  types.Object   `tfsdk:"flex_inline" json:"flexInline"`
	FlexInlineFailover          types.String   `tfsdk:"flex_inline_failover" json:"flexInlineFailover"`
	FlexInlineVlanId            types.Int64    `tfsdk:"flex_inline_vlan_id" json:"flexInlineVlanId"`
	FlowRules                   types.Object   `tfsdk:"flow_rules" json:"flowRules"`
	FlowSample5GOverlapRules    types.Object   `tfsdk:"flow_sample5_g_overlap_rules" json:"flowSample5gOverlapRules"`
	FlowSample5GRules           types.Object   `tfsdk:"flow_sample5_g_rules" json:"flowSample5gRules"`
	FlowSampleDiameterRules     types.Object   `tfsdk:"flow_sample_diameter_rules" json:"flowSampleDiameterRules"`
	FlowSampleOverlapRules      types.Object   `tfsdk:"flow_sample_overlap_rules" json:"flowSampleOverlapRules"`
	FlowSampleRules             types.Object   `tfsdk:"flow_sample_rules" json:"flowSampleRules"`
	FlowSampleSipRules          types.Object   `tfsdk:"flow_sample_sip_rules" json:"flowSampleSipRules"`
	FlowWhitelist5GOverlapRules types.Object   `tfsdk:"flow_whitelist5_g_overlap_rules" json:"flowWhitelist5gOverlapRules"`
	FlowWhitelist5GRules        types.Object   `tfsdk:"flow_whitelist5_g_rules" json:"flowWhitelist5gRules"`
	FlowWhitelistOverlapRules   types.Object   `tfsdk:"flow_whitelist_overlap_rules" json:"flowWhitelistOverlapRules"`
	FlowWhitelistRules          types.Object   `tfsdk:"flow_whitelist_rules" json:"flowWhitelistRules"`
	Fstype                      types.Object   `tfsdk:"fstype"`
	GsRules                     types.Object   `tfsdk:"gs_rules" json:"gsRules"`
	Gsop                        types.String   `tfsdk:"gsop"`
	HealthState                 types.String   `tfsdk:"health_state" json:"healthState"`
	HealthStateReasons          types.List     `tfsdk:"health_state_reasons" json:"healthStateReasons"`
	InlineTrafficPath           types.String   `tfsdk:"inline_traffic_path" json:"inlineTrafficPath"`
	InlineTrafficType           types.String   `tfsdk:"inline_traffic_type" json:"inlineTrafficType"`
	IpRewrite                   types.Object   `tfsdk:"ip_rewrite" json:"ipRewrite"`
	ModTime                     types.Int64    `tfsdk:"mod_time" json:"modTime"`
	NullDstPort                 types.Bool     `tfsdk:"null_dst_port" json:"nullDstPort"`
	Order                       types.Int64    `tfsdk:"order"`
	Rewrite                     types.Object   `tfsdk:"rewrite"`
	Roles                       types.Object   `tfsdk:"roles"`
	RuleMatching                types.String   `tfsdk:"rule_matching" json:"ruleMatching"`
	Rules                       types.Object   `tfsdk:"rules"`
	RxClusterPorts              types.List     `tfsdk:"rx_cluster_ports" json:"rxClusterPorts"`
	SrcPorts                    types.List     `tfsdk:"src_ports" json:"srcPorts"`
	SubType                     types.String   `tfsdk:"sub_type" json:"subType"`
	TrafficType                 types.String   `tfsdk:"traffic_type" json:"trafficType"`
	TxClusterPorts              types.List     `tfsdk:"tx_cluster_ports" json:"txClusterPorts"`
	Type                        types.String   `tfsdk:"type"`
	UpdatedTime                 types.Float64  `tfsdk:"updated_time" json:"updatedTime"`
	VlanTag                     types.Object   `tfsdk:"vlan_tag" json:"vlanTag"`
	Timeouts                    timeouts.Value `tfsdk:"timeouts"`
}

// Metadata returns the resource type name.
func (r *MapResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_map"
}

// Schema returns the Terraform schema for this resource.
func (r *MapResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Find map by alias", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "unique map alias", Required: true}, "ap_rules": schema.SingleNestedAttribute{MarkdownDescription: "pass and drop application profile rules", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"drop_rules": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"application_profile": schema.StringAttribute{MarkdownDescription: "application profile alias", Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "rule_id": schema.Int64Attribute{MarkdownDescription: "application profile rule Id, should not have same id as gsRules", Required: true, Validators: []validator.Int64{int64validator.AtLeast(1)}}}}}, "pass_rules": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"application_profile": schema.StringAttribute{MarkdownDescription: "application profile alias", Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "rule_id": schema.Int64Attribute{MarkdownDescription: "application profile rule Id, should not have same id as gsRules", Required: true, Validators: []validator.Int64{int64validator.AtLeast(1)}}}}}}}, "cluster_id": schema.StringAttribute{MarkdownDescription: "id of the defining cluster", Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "comment": schema.StringAttribute{Optional: true, Computed: true}, "dst_ports": schema.ListAttribute{MarkdownDescription: "List of the 'to' ports. Invalid if 'nullDstPort' is true'. Only port number is supported", Required: true, ElementType: types.StringType}, "egress_gigastream": schema.ListAttribute{MarkdownDescription: "Applicable to fabric map only. It is to specify gigastream(s) to use as egress when going across cluster. The format of array member is clusterID:gigastreamAlias", Optional: true, Computed: true, ElementType: types.StringType}, "enable": schema.BoolAttribute{MarkdownDescription: "enable/disable map, applicable only to first level maps", Optional: true, Computed: true}, "encap_tunnel": schema.StringAttribute{MarkdownDescription: "tunnel alias", Optional: true, Computed: true}, "flex_inline": schema.SingleNestedAttribute{MarkdownDescription: "When 'oobCopy' is configured at least one of 'aToB' or 'bToA' must be configured", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"a_to_b": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"ib_pathway": schema.StringAttribute{MarkdownDescription: "ibPathway alias. Only applicable when type is 'ibPathway'", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "tools": schema.ListAttribute{MarkdownDescription: "ordered list of inline tools or vports. Only applicable when 'type' is 'tools'", Optional: true, Computed: true, ElementType: types.StringType}, "type": schema.StringAttribute{MarkdownDescription: "when set to 'tools', the list of the processing inline tools have to be provided. when set to 'ibPathway', ibPathway alias should be provided ", Required: true, Validators: []validator.String{stringvalidator.OneOf("bypass", "tools", "reverse", "same", "ibPathway")}}}}, "b_to_a": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"ib_pathway": schema.StringAttribute{MarkdownDescription: "ibPathway alias. Only applicable when type is 'ibPathway'", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "tools": schema.ListAttribute{MarkdownDescription: "ordered list of inline tools or vports. Only applicable when 'type' is 'tools'", Optional: true, Computed: true, ElementType: types.StringType}, "type": schema.StringAttribute{MarkdownDescription: "when set to 'tools', the list of the processing inline tools have to be provided. when set to 'ibPathway', ibPathway alias should be provided ", Required: true, Validators: []validator.String{stringvalidator.OneOf("bypass", "tools", "reverse", "same", "ibPathway")}}}}, "oob_copy": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"direction": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("aToB", "bToA", "both")}}, "dst_ports": schema.ListAttribute{MarkdownDescription: "list of destination tool ports", Required: true, ElementType: types.StringType, Validators: []validator.List{listvalidator.SizeAtLeast(1)}}, "src_ports": schema.ListAttribute{MarkdownDescription: "inline network or an item from a-to-b and b-to-a lists", Required: true, ElementType: types.StringType, Validators: []validator.List{listvalidator.SizeAtLeast(1)}}, "tag": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"type": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.OneOf("none", "asInline", "original")}}}}}}}, "svt_mode": schema.BoolAttribute{Optional: true, Computed: true}, "svt_tag": schema.Int64Attribute{MarkdownDescription: "only applicable when svtMode is enabled", Optional: true, Computed: true}, "tag": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"tag_protocol_id": schema.StringAttribute{MarkdownDescription: "When tool VLAN tag is added , this protocol Id will be added which egress out the traffic", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("0x8100", "0x9100", "0x88a8")}}, "type": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.OneOf("auto", "vlan")}}, "vlan_id": schema.Int64Attribute{MarkdownDescription: "only applicable when type is 'vlan'", Optional: true, Computed: true}}}}}, "flex_inline_failover": schema.StringAttribute{MarkdownDescription: "only valid for flexInline maps", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("bypass", "originalMap")}}, "flex_inline_vlan_id": schema.Int64Attribute{MarkdownDescription: "VLAN ID carried in the VLAN tag of packets coming from the inline-network port(s). valid and applicable only with 'flexInline' map type, 'srcPorts' should be either inline-network or ib-pathway or vport", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 4094)}}, "flow_rules": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Rules Container. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"drop_rules": schema.SetNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"gtp": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Rule GTP match Definition. Private class", Required: true, Attributes: map[string]schema.Attribute{"imei": schema.StringAttribute{MarkdownDescription: "mutually exclusive with 'imsi' and 'msisdn'. If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([0-9]{6,16})|([0-9]{0,14}[*])$"), "value must match pattern \"^([0-9]{6,16})|([0-9]{0,14}[*])$\"")}}, "imsi": schema.StringAttribute{MarkdownDescription: "mutually exclusive with 'imei' and 'msisdn'. If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([0-9]{6,15})|([0-9]{0,14}[*])$"), "value must match pattern \"^([0-9]{6,15})|([0-9]{0,14}[*])$\"")}}, "interface": schema.StringAttribute{MarkdownDescription: "interface type. Mutually exclusive with version", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("Gn", "S5", "S10", "S11")}}, "msisdn": schema.StringAttribute{MarkdownDescription: "mutually exclusive with 'imsi' and 'imei'. If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([0-9]{6,15})|([0-9]{0,14}[*])$"), "value must match pattern \"^([0-9]{6,15})|([0-9]{0,14}[*])$\"")}}, "version": schema.StringAttribute{MarkdownDescription: "mutually exclusive with interface", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("any", "v1", "v2")}}}}, "rule_id": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.AtLeast(1)}}}}}, "pass_rules": schema.SetNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"gtp": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Rule GTP match Definition. Private class", Required: true, Attributes: map[string]schema.Attribute{"imei": schema.StringAttribute{MarkdownDescription: "mutually exclusive with 'imsi' and 'msisdn'. If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([0-9]{6,16})|([0-9]{0,14}[*])$"), "value must match pattern \"^([0-9]{6,16})|([0-9]{0,14}[*])$\"")}}, "imsi": schema.StringAttribute{MarkdownDescription: "mutually exclusive with 'imei' and 'msisdn'. If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([0-9]{6,15})|([0-9]{0,14}[*])$"), "value must match pattern \"^([0-9]{6,15})|([0-9]{0,14}[*])$\"")}}, "interface": schema.StringAttribute{MarkdownDescription: "interface type. Mutually exclusive with version", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("Gn", "S5", "S10", "S11")}}, "msisdn": schema.StringAttribute{MarkdownDescription: "mutually exclusive with 'imsi' and 'imei'. If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([0-9]{6,15})|([0-9]{0,14}[*])$"), "value must match pattern \"^([0-9]{6,15})|([0-9]{0,14}[*])$\"")}}, "version": schema.StringAttribute{MarkdownDescription: "mutually exclusive with interface", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("any", "v1", "v2")}}}}, "rule_id": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.AtLeast(1)}}}}}}}, "flow_sample5_g_overlap_rules": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Sample 5g Overlap Rules Container. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"pass_rules": schema.SetNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"comment": schema.StringAttribute{Optional: true, Computed: true}, "flow5_g": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Sample Overlap Rule 5g match Definition. Private class", Required: true, Attributes: map[string]schema.Attribute{"dnn": schema.StringAttribute{MarkdownDescription: "Access Point Name pattern.  Alphanumeric, '.', '-', and '*' allowed.", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.LengthBetween(1, 128), stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9.*-]*$"), "value must match pattern \"^[a-zA-Z0-9.*-]*$\"")}}, "gpsi": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([0-9]{6,15})|([0-9]{0,14}[*])$"), "value must match pattern \"^([0-9]{6,15})|([0-9]{0,14}[*])$\"")}}, "nas_5_qi": schema.StringAttribute{MarkdownDescription: "5G QoS Indicator. Valid 5qi value <1 - 255>", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|[0-9]{1,3})$"), "value must match pattern \"^([*]|[0-9]{1,3})$\"")}}, "nci": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|([A-Fa-f0-9]{9})|([A-Fa-f0-9]{4,8}[*]))$"), "value must match pattern \"^([*]|([A-Fa-f0-9]{9})|([A-Fa-f0-9]{4,8}[*]))$\"")}}, "nsiid": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the SD value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|([0-9]{1,3})|([0-9]{1,3}[.][a-fA-F0-9]{1,6})|([0-9]{1,3}[.][a-fA-F0-9]{0,5}[*]))$"), "value must match pattern \"^([*]|([0-9]{1,3})|([0-9]{1,3}[.][a-fA-F0-9]{1,6})|([0-9]{1,3}[.][a-fA-F0-9]{0,5}[*]))$\"")}}, "pei": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([0-9]{6,15})|([0-9]{0,14}[*])$"), "value must match pattern \"^([0-9]{6,15})|([0-9]{0,14}[*])$\"")}}, "plmn_id": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the MNC value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|([0-9]{3}[.][0-9]{2,3})|([0-9]{3}[.][0-9]{0,2}[*]))$"), "value must match pattern \"^([*]|([0-9]{3}[.][0-9]{2,3})|([0-9]{3}[.][0-9]{0,2}[*]))$\"")}}, "supi": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([0-9]{6,15})|([0-9]{0,14}[*])$"), "value must match pattern \"^([0-9]{6,15})|([0-9]{0,14}[*])$\"")}}, "tac": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|([0-9a-fA-F]{6})|([0-9a-fA-F]{2,5}[*]))$"), "value must match pattern \"^([*]|([0-9a-fA-F]{6})|([0-9a-fA-F]{2,5}[*]))$\"")}}}}, "percentage": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.Between(0, 100)}}, "rule_id": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.Between(1, 2147483647)}}}}}}}, "flow_sample5_g_rules": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Sample 5g Rules Container. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"pass_rules": schema.SetNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"comment": schema.StringAttribute{Optional: true, Computed: true}, "flow5_g": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Sample Rule 5g match Definition. Private class", Required: true, Attributes: map[string]schema.Attribute{"dnn": schema.StringAttribute{MarkdownDescription: "Access Point Name pattern.  Alphanumeric, '.', '-', and '*' allowed.", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.LengthBetween(1, 128), stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9.*-]*$"), "value must match pattern \"^[a-zA-Z0-9.*-]*$\"")}}, "gpsi": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([0-9]{6,15})|([0-9]{0,14}[*])$"), "value must match pattern \"^([0-9]{6,15})|([0-9]{0,14}[*])$\"")}}, "nci": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^(([A-Fa-f0-9]{9})|([A-Fa-f0-9]{4,8}[*]))$"), "value must match pattern \"^(([A-Fa-f0-9]{9})|([A-Fa-f0-9]{4,8}[*]))$\"")}}, "nsiid": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the SD value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^(([0-9]{1,3})|([0-9]{1,3}[.][a-fA-F0-9]{1,6})|([0-9]{1,3}[.][a-fA-F0-9]{0,5}[*]))$"), "value must match pattern \"^(([0-9]{1,3})|([0-9]{1,3}[.][a-fA-F0-9]{1,6})|([0-9]{1,3}[.][a-fA-F0-9]{0,5}[*]))$\"")}}, "pei": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([0-9]{6,16})|([0-9]{0,14}[*])$"), "value must match pattern \"^([0-9]{6,16})|([0-9]{0,14}[*])$\"")}}, "plmn_id": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the MNC value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^(([0-9]{3}[.][0-9]{2,3})|([0-9]{3}[.][0-9]{0,2}[*]))$"), "value must match pattern \"^(([0-9]{3}[.][0-9]{2,3})|([0-9]{3}[.][0-9]{0,2}[*]))$\"")}}, "supi": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([0-9]{6,15})|([0-9]{0,14}[*])$"), "value must match pattern \"^([0-9]{6,15})|([0-9]{0,14}[*])$\"")}}, "tac": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^(([0-9a-fA-F]{6})|([0-9a-fA-F]{0,5}[*]))$"), "value must match pattern \"^(([0-9a-fA-F]{6})|([0-9a-fA-F]{0,5}[*]))$\"")}}}}, "percentage": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.Between(0, 100)}}, "priority": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.AtLeast(1)}}, "rule_id": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.AtLeast(1)}}}}}}}, "flow_sample_diameter_rules": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Sample Diameter Rules Container. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"pass_rules": schema.SetNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"diameter": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Sample Diameter Rule Definition", Required: true, Attributes: map[string]schema.Attribute{"user_name": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([0-9]{6,15})|([0-9]{0,14}[*])$"), "value must match pattern \"^([0-9]{6,15})|([0-9]{0,14}[*])$\"")}}}}, "interface": schema.StringAttribute{MarkdownDescription: "interface type", Required: true, Validators: []validator.String{stringvalidator.OneOf("s6a")}}, "percentage": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.Between(0, 100)}}, "rule_id": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.AtLeast(1)}}}}}}}, "flow_sample_overlap_rules": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Sample Overlap Rules Container. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"pass_rules": schema.SetNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"comment": schema.StringAttribute{Optional: true, Computed: true}, "gtp": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Sample Rule GTP match Definition. Private class", Required: true, Attributes: map[string]schema.Attribute{"apn": schema.StringAttribute{MarkdownDescription: "Access Point Name pattern.  Alphanumeric, '.', '-', and '*' allowed.", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.LengthBetween(1, 128), stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9.*-]*$"), "value must match pattern \"^[a-zA-Z0-9.*-]*$\"")}}, "eci": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^(([A-Fa-f0-9]{7,8})|([A-Fa-f0-9]{4,7}[*]))$"), "value must match pattern \"^(([A-Fa-f0-9]{7,8})|([A-Fa-f0-9]{4,7}[*]))$\"")}}, "imei": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([0-9]{6,16})|([0-9]{0,14}[*])$"), "value must match pattern \"^([0-9]{6,16})|([0-9]{0,14}[*])$\"")}}, "imsi": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([0-9]{6,15})|([0-9]{0,14}[*])$"), "value must match pattern \"^([0-9]{6,15})|([0-9]{0,14}[*])$\"")}}, "interface": schema.StringAttribute{MarkdownDescription: "interface type. Mutually exclusive with version", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("Gn", "S5", "S10", "S11", "S2b", "S11U", "S1U", "S5S8U", "N3", "N9")}}, "msisdn": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([0-9]{6,15})|([0-9]{0,14}[*])$"), "value must match pattern \"^([0-9]{6,15})|([0-9]{0,14}[*])$\"")}}, "nas_5_qi": schema.StringAttribute{MarkdownDescription: "5G QoS Indicator. Valid 5qi value <1 - 255>", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|[0-9]{1,3})$"), "value must match pattern \"^([*]|[0-9]{1,3})$\"")}}, "nci": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|([A-Fa-f0-9]{9})|([A-Fa-f0-9]{4,8}[*]))$"), "value must match pattern \"^([*]|([A-Fa-f0-9]{9})|([A-Fa-f0-9]{4,8}[*]))$\"")}}, "plmn_id": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the MNC value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^(([0-9]{3}[.][0-9]{2,3})|([0-9]{3}[.][0-9]{0,2}[*]))$"), "value must match pattern \"^(([0-9]{3}[.][0-9]{2,3})|([0-9]{3}[.][0-9]{0,2}[*]))$\"")}}, "qci": schema.Int64Attribute{MarkdownDescription: "QoS Class Indicator", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 255)}}, "snssai": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the SD value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|([0-9]{1,3})|([0-9]{1,3}[.][a-fA-F0-9]{1,6})|([0-9]{1,3}[.][a-fA-F0-9]{0,5}[*]))$"), "value must match pattern \"^([*]|([0-9]{1,3})|([0-9]{1,3}[.][a-fA-F0-9]{1,6})|([0-9]{1,3}[.][a-fA-F0-9]{0,5}[*]))$\"")}}, "tac": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^(([0-9a-fA-F]{4})|([0-9a-fA-F]{0,3}[*]))$"), "value must match pattern \"^(([0-9a-fA-F]{4})|([0-9a-fA-F]{0,3}[*]))$\"")}}, "tac_5_g": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|([0-9a-fA-F]{6})|([0-9a-fA-F]{2,5}[*]))$"), "value must match pattern \"^([*]|([0-9a-fA-F]{6})|([0-9a-fA-F]{2,5}[*]))$\"")}}, "version": schema.StringAttribute{MarkdownDescription: "mutually exclusive with interface", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("any", "v1", "v2")}}}}, "percentage": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.Between(0, 100)}}, "periodic_recalc": schema.BoolAttribute{MarkdownDescription: "Enable Periodic Recalc for rotational sampling. Map look up in the data path based on this flag", Optional: true, Computed: true}, "priority": schema.Int64Attribute{MarkdownDescription: "maximum value is equal to the number of rules upon completion of the request", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.AtLeast(1)}}, "rule_id": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.AtLeast(1)}}}}}}}, "flow_sample_rules": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Sample Rules Container. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"pass_rules": schema.SetNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"comment": schema.StringAttribute{Optional: true, Computed: true}, "gtp": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Sample Rule GTP match Definition. Private class", Required: true, Attributes: map[string]schema.Attribute{"apn": schema.StringAttribute{MarkdownDescription: "Access Point Name pattern.  Alphanumeric, '.', '-', and '*' allowed.", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.LengthBetween(1, 128), stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9.*-]*$"), "value must match pattern \"^[a-zA-Z0-9.*-]*$\"")}}, "eci": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^(([A-Fa-f0-9]{7,8})|([A-Fa-f0-9]{4,7}[*]))$"), "value must match pattern \"^(([A-Fa-f0-9]{7,8})|([A-Fa-f0-9]{4,7}[*]))$\"")}}, "imei": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([0-9]{6,16})|([0-9]{0,14}[*])$"), "value must match pattern \"^([0-9]{6,16})|([0-9]{0,14}[*])$\"")}}, "imsi": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([0-9]{6,15})|([0-9]{0,14}[*])$"), "value must match pattern \"^([0-9]{6,15})|([0-9]{0,14}[*])$\"")}}, "interface": schema.StringAttribute{MarkdownDescription: "interface type. Mutually exclusive with version", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("Gn", "S5", "S10", "S11", "S2b", "S11U", "S1U", "S5S8U", "N3", "N9")}}, "msisdn": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([0-9]{6,15})|([0-9]{0,14}[*])$"), "value must match pattern \"^([0-9]{6,15})|([0-9]{0,14}[*])$\"")}}, "nas_5_qi": schema.StringAttribute{MarkdownDescription: "5G QoS Indicator. Valid 5qi value <1 - 255>", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|[0-9]{1,3})$"), "value must match pattern \"^([*]|[0-9]{1,3})$\"")}}, "nci": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|([A-Fa-f0-9]{9})|([A-Fa-f0-9]{4,8}[*]))$"), "value must match pattern \"^([*]|([A-Fa-f0-9]{9})|([A-Fa-f0-9]{4,8}[*]))$\"")}}, "plmn_id": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the MNC value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^(([0-9]{3}[.][0-9]{2,3})|([0-9]{3}[.][0-9]{0,2}[*]))$"), "value must match pattern \"^(([0-9]{3}[.][0-9]{2,3})|([0-9]{3}[.][0-9]{0,2}[*]))$\"")}}, "qci": schema.Int64Attribute{MarkdownDescription: "QoS Class Indicator", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 255)}}, "snssai": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the SD value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|([0-9]{1,3})|([0-9]{1,3}[.][a-fA-F0-9]{1,6})|([0-9]{1,3}[.][a-fA-F0-9]{0,5}[*]))$"), "value must match pattern \"^([*]|([0-9]{1,3})|([0-9]{1,3}[.][a-fA-F0-9]{1,6})|([0-9]{1,3}[.][a-fA-F0-9]{0,5}[*]))$\"")}}, "tac": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^(([0-9a-fA-F]{4})|([0-9a-fA-F]{0,3}[*]))$"), "value must match pattern \"^(([0-9a-fA-F]{4})|([0-9a-fA-F]{0,3}[*]))$\"")}}, "tac_5_g": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|([0-9a-fA-F]{6})|([0-9a-fA-F]{2,5}[*]))$"), "value must match pattern \"^([*]|([0-9a-fA-F]{6})|([0-9a-fA-F]{2,5}[*]))$\"")}}, "version": schema.StringAttribute{MarkdownDescription: "mutually exclusive with interface", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("any", "v1", "v2")}}}}, "percentage": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.Between(0, 100)}}, "periodic_recalc": schema.BoolAttribute{MarkdownDescription: "Enable Periodic Recalc for rotational sampling. Map look up in the data path based on this flag", Optional: true, Computed: true}, "priority": schema.Int64Attribute{MarkdownDescription: "maximum value is equal to the number of rules upon completion of the request", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.AtLeast(1)}}, "rule_id": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.AtLeast(1)}}}}}}}, "flow_sample_sip_rules": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Sample Sip Rules Container. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"pass_rules": schema.SetNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"percentage": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.Between(0, 100)}}, "rule_id": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.AtLeast(1)}}, "sip": schema.SingleNestedAttribute{Required: true, Attributes: map[string]schema.Attribute{"callee_id": schema.StringAttribute{MarkdownDescription: "sip callee id", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9@#$*-_.!~()&=+,;?/]{1,64}$"), "value must match pattern \"^[a-zA-Z0-9@#$*-_.!~()&=+,;?/]{1,64}$\"")}}, "callee_id_range": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"max_value": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9@#$*-_.!~()&=+,;?/]{1,64}$"), "value must match pattern \"^[a-zA-Z0-9@#$*-_.!~()&=+,;?/]{1,64}$\"")}}, "value": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9@#$*-_.!~()&=+,;?/]{1,64}$"), "value must match pattern \"^[a-zA-Z0-9@#$*-_.!~()&=+,;?/]{1,64}$\"")}}}}, "caller_id": schema.StringAttribute{MarkdownDescription: "sip caller id", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9@#$*-_.!~()&=+,;?/]{1,64}$"), "value must match pattern \"^[a-zA-Z0-9@#$*-_.!~()&=+,;?/]{1,64}$\"")}}, "caller_id_range": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"max_value": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9@#$*-_.!~()&=+,;?/]{1,64}$"), "value must match pattern \"^[a-zA-Z0-9@#$*-_.!~()&=+,;?/]{1,64}$\"")}}, "value": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9@#$*-_.!~()&=+,;?/]{1,64}$"), "value must match pattern \"^[a-zA-Z0-9@#$*-_.!~()&=+,;?/]{1,64}$\"")}}}}, "id_range": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"max_value": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9@#$*-_.!~()&=+,;?/]{1,64}$"), "value must match pattern \"^[a-zA-Z0-9@#$*-_.!~()&=+,;?/]{1,64}$\"")}}, "value": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9@#$*-_.!~()&=+,;?/]{1,64}$"), "value must match pattern \"^[a-zA-Z0-9@#$*-_.!~()&=+,;?/]{1,64}$\"")}}}}}}}}}}}, "flow_whitelist5_g_overlap_rules": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Whitelist 5g Overlap Rule match Definition. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"dnn": schema.StringAttribute{MarkdownDescription: "Domain Network Name pattern.  Alphanumeric, '.', '-', and '*' allowed.", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.LengthBetween(1, 128), stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9.*-]*$"), "value must match pattern \"^[a-zA-Z0-9.*-]*$\"")}}, "type": schema.StringAttribute{MarkdownDescription: "Set 5G WL-DB lookup type", Optional: true, Computed: true}}}, "flow_whitelist5_g_rules": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Whitelist 5g Rule GTP match Definition. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"dnn": schema.StringAttribute{MarkdownDescription: "Domain Network Name pattern.  Alphanumeric, '.', '-', and '*' allowed.", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.LengthBetween(1, 128), stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9.*-]*$"), "value must match pattern \"^[a-zA-Z0-9.*-]*$\"")}}, "type": schema.StringAttribute{MarkdownDescription: "Set 5G WL-DB lookup type", Optional: true, Computed: true}, "whitelist_databases": schema.ListAttribute{MarkdownDescription: "Attach whitelist databases to the map", Optional: true, Computed: true, ElementType: types.StringType}}}, "flow_whitelist_overlap_rules": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Whitelist Overlap Rules Container. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"pass_rules": schema.SetNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"flow5_g": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Whitelist 5g Rule GTP match Definition. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"dnn": schema.StringAttribute{MarkdownDescription: "Domain Network Name pattern.  Alphanumeric, '.', '-', and '*' allowed.", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.LengthBetween(1, 128), stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9.*-]*$"), "value must match pattern \"^[a-zA-Z0-9.*-]*$\"")}}, "type": schema.StringAttribute{MarkdownDescription: "Set 5G WL-DB lookup type", Optional: true, Computed: true}, "whitelist_databases": schema.ListAttribute{MarkdownDescription: "Attach whitelist databases to the map", Optional: true, Computed: true, ElementType: types.StringType}}}, "gtp": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Whitelist Rule GTP match Definition. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"apn": schema.StringAttribute{MarkdownDescription: "Access Point Name pattern.  Alphanumeric, '.', '-', and '*' allowed.", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.LengthBetween(1, 128), stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9.*-]*$"), "value must match pattern \"^[a-zA-Z0-9.*-]*$\"")}}, "interface": schema.StringAttribute{MarkdownDescription: "interface type. Mutually exclusive with version. required till H 5.6", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("Gn", "S5", "S10", "S11", "S2b", "S11U", "S1U", "S5S8U", "N3", "N9")}}, "type": schema.StringAttribute{MarkdownDescription: "Set GTP WL-DB lookup type", Optional: true, Computed: true}, "version": schema.StringAttribute{MarkdownDescription: "mutually exclusive with interface", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("v1", "v2")}}, "whitelist_databases": schema.ListAttribute{MarkdownDescription: "Attach whitelist databases to the map", Optional: true, Computed: true, ElementType: types.StringType}}}, "rule_id": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.AtLeast(1)}}, "sip": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Whitelist Rule Sip match definition", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"type": schema.StringAttribute{MarkdownDescription: "all:Whitelist based on caller/callee/source/destination IP address, bothAddr: Whitelist source/destination IP address, bothId:Whitelist Caller/Callee Id's, calleeId: Whitelist Callee ID, callerId: Whitelist Caller ID, destIp: Whitelist based on destination IP address, srcIp: Whitelist based on Source IP address", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("all", "bothAddr", "bothId", "calleeId", "callerId", "destIp", "srcIp")}}}}}}}}}, "flow_whitelist_rules": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Whitelist Rules Container. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"pass_rules": schema.SetNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"flow5_g": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Whitelist 5g Rule GTP match Definition. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"dnn": schema.StringAttribute{MarkdownDescription: "Domain Network Name pattern.  Alphanumeric, '.', '-', and '*' allowed.", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.LengthBetween(1, 128), stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9.*-]*$"), "value must match pattern \"^[a-zA-Z0-9.*-]*$\"")}}, "type": schema.StringAttribute{MarkdownDescription: "Set 5G WL-DB lookup type", Optional: true, Computed: true}, "whitelist_databases": schema.ListAttribute{MarkdownDescription: "Attach whitelist databases to the map", Optional: true, Computed: true, ElementType: types.StringType}}}, "gtp": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Whitelist Rule GTP match Definition. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"apn": schema.StringAttribute{MarkdownDescription: "Access Point Name pattern.  Alphanumeric, '.', '-', and '*' allowed.", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.LengthBetween(1, 128), stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9.*-]*$"), "value must match pattern \"^[a-zA-Z0-9.*-]*$\"")}}, "interface": schema.StringAttribute{MarkdownDescription: "interface type. Mutually exclusive with version. required till H 5.6", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("Gn", "S5", "S10", "S11", "S2b", "S11U", "S1U", "S5S8U", "N3", "N9")}}, "type": schema.StringAttribute{MarkdownDescription: "Set GTP WL-DB lookup type", Optional: true, Computed: true}, "version": schema.StringAttribute{MarkdownDescription: "mutually exclusive with interface", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("v1", "v2")}}, "whitelist_databases": schema.ListAttribute{MarkdownDescription: "Attach whitelist databases to the map", Optional: true, Computed: true, ElementType: types.StringType}}}, "rule_id": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.AtLeast(1)}}, "sip": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Whitelist Rule Sip match definition", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"type": schema.StringAttribute{MarkdownDescription: "all:Whitelist based on caller/callee/source/destination IP address, bothAddr: Whitelist source/destination IP address, bothId:Whitelist Caller/Callee Id's, calleeId: Whitelist Callee ID, callerId: Whitelist Caller ID, destIp: Whitelist based on destination IP address, srcIp: Whitelist based on Source IP address", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("all", "bothAddr", "bothId", "calleeId", "callerId", "destIp", "srcIp")}}}}}}}}}, "fstype": schema.SingleNestedAttribute{MarkdownDescription: "Type of Mobility Flowsampling & properties", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"offset": schema.Int64Attribute{MarkdownDescription: "Offset for Mobility Rotational Flowsampling", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 99)}}, "timer": schema.Int64Attribute{MarkdownDescription: "Timer for Mobility Rotational Flowsampling in minutes", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(15, 45000)}}, "type": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("_default", "rotational")}}}}, "gs_rules": schema.SingleNestedAttribute{MarkdownDescription: "Map GigaSMART Rules Container. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"drop_rules": schema.DynamicAttribute{Optional: true, Computed: true}, "pass_rules": schema.DynamicAttribute{Optional: true, Computed: true}}}, "gsop": schema.StringAttribute{MarkdownDescription: "Alias of referenced GSOP. Applicable to 'regular' and 'secondLevel' map types", Optional: true, Computed: true}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("green", "yellow", "orange", "red")}}, "health_state_reasons": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Optional: true, Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("green", "yellow", "orange", "red")}}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("PORT_LOW_UTIL", "PORT_HIGH_UTIL", "PORT_PACKET_DROP", "PORT_PACKET_ERROR", "GIGASTREAM_IMBALANCE")}}}}}, "inline_traffic_path": schema.StringAttribute{MarkdownDescription: "Only applicable for 'inline' map types, in which case defaults to 'normal'. When set to 'bypass', the 'dstPorts' must be empty", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("normal", "bypass", "oneArm", "natPat")}}, "inline_traffic_type": schema.StringAttribute{MarkdownDescription: "Only applicable for 'inline/passAll' map types, in which case defaults to 'symmetric'. For 'asymmetric' maps, 'srcPort' must be of type 'inline-net'", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("symmetric", "asymmetric")}}, "ip_rewrite": schema.SingleNestedAttribute{MarkdownDescription: "IpRewrite options on the packets", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"dst_ip": schema.StringAttribute{Optional: true, Computed: true}, "src_ip": schema.StringAttribute{Optional: true, Computed: true}}}, "mod_time": schema.Int64Attribute{MarkdownDescription: "Last modification time of the map in milliseconds since the epoch", Computed: true}, "null_dst_port": schema.BoolAttribute{MarkdownDescription: "enabled when the dstPort is null", Optional: true, Computed: true}, "order": schema.Int64Attribute{MarkdownDescription: "relative order within per-source port map chain", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.AtLeast(0)}}, "rewrite": schema.SingleNestedAttribute{MarkdownDescription: "Rewrite options on the packets", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"dst_mac": schema.StringAttribute{Optional: true, Computed: true}, "src_mac": schema.StringAttribute{Optional: true, Computed: true}}}, "roles": schema.SingleNestedAttribute{MarkdownDescription: "Map Roles Container. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"editors": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType}, "listeners": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType}, "owners": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType}, "viewers": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType}}}, "rule_matching": schema.StringAttribute{MarkdownDescription: "If set to 'blacklist', packet are passed when no 'drop' rules are matched. This field is only valid for 'regular/byRule' map types. (maps into 'no-rule-match' CLI command)", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("normal", "blacklist")}}, "rules": schema.SingleNestedAttribute{MarkdownDescription: "Map Rules Container. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"drop_rules": schema.DynamicAttribute{Optional: true, Computed: true}, "pass_rules": schema.DynamicAttribute{Optional: true, Computed: true}}}, "rx_cluster_ports": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "src_ports": schema.ListAttribute{MarkdownDescription: "list of the 'from' ports. Only port number is supported", Required: true, ElementType: types.StringType, Validators: []validator.List{listvalidator.SizeAtLeast(1)}}, "sub_type": schema.StringAttribute{MarkdownDescription: "'byRule' is applicable to all map types; 'collector' is applicable to 'regular', 'inline', 'flexInline' and 'secondLevel' maps; 'passAll' is applicable to 'regular' and 'inline' maps; 'flowFilter', 'flowSample', 'flowWhitelist', 'flowSampleSip', 'flowWhitelistSip', 'flowSampleDiameter', 'flowWhitelistDiameter', 'flowSampleOverlap', 'flowWhitelistOverlap', 'flowSample5g','flowWhitelist5g','flowSample5gOverlap' and 'flowWhitelist5gOverlap' are applicable to 'secondLevel' maps", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("byRule", "passAll", "collector", "flowFilter", "flowSample", "flowWhitelist", "flowSampleOverlap", "flowWhitelistOverlap", "ingressToVp", "egressFromVp", "oobFromVp", "flowSampleSip", "flowWhitelistSip", "flowSampleDiameter", "flowWhitelistDiameter", "flowSample5g", "flowWhitelist5g", "flowSample5gOverlap", "flowWhitelist5gOverlap", "networkEnd", "inlinePair")}}, "traffic_type": schema.StringAttribute{MarkdownDescription: "Only applicable for 'firstLevel/byRule' map types, in which case defaults to 'user'", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("control", "user")}}, "tx_cluster_ports": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "type": schema.StringAttribute{MarkdownDescription: "'regular' maps are from network/hybrid ports to tool/hybrid/gigastream; 'inline' maps are from inline ports to inline/tool/hybrid/gigastream; 'firstLevel' are from network/hybrid ports to vPorts/tool/hybrid/gigastream; 'secondLevel' maps are from vPorts to tool/hybrid/gigastream; 'inlineFirstLevel is from inline network to vport; 'inlineSecondLevel' is from vport to inline tool; 'flexInline' is from inline network to tools; 'transitLevel' is from vport to vport", Required: true, Validators: []validator.String{stringvalidator.OneOf("regular", "inline", "firstLevel", "secondLevel", "inlineFirstLevel", "inlineSecondLevel", "flexInline", "transitLevel")}}, "updated_time": schema.Float64Attribute{MarkdownDescription: "Last Updated time of the fabric map", Optional: true, Computed: true}, "vlan_tag": schema.SingleNestedAttribute{MarkdownDescription: "This field is only valid for 'regular/byRule', 'collector' map types", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"tag_protocol_id": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("0x8100", "0x9100", "0x88a8")}}, "vlan_action": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.OneOf("add", "modify", "strip")}}, "vlan_id": schema.Int64Attribute{Optional: true, Computed: true}}}}, Blocks: map[string]schema.Block{"timeouts": timeouts.Block(ctx, timeouts.Opts{Create: true, Read: true, Update: true, Delete: true})}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *MapResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan MapResourceModel
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
func (r *MapResource) createRemote(ctx context.Context, plan *MapResourceModel, resp *resource.CreateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_map", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_map", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/maps"
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_map", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", plan.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_map", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error creating gigavuecore_map", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error creating gigavuecore_map", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error creating gigavuecore_map", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error creating gigavuecore_map", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error creating gigavuecore_map", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error creating gigavuecore_map", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error creating gigavuecore_map", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error creating gigavuecore_map", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error creating gigavuecore_map", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error creating gigavuecore_map", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_map", fmt.Sprintf("Could not map response to state: %s", err))
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
			resp.Diagnostics.AddError("Error creating gigavuecore_map", "The create response did not contain an identifier and no Location header was returned, so the resource cannot be tracked in state.")
			return
		}
	}
}

// Read refreshes the Terraform state with the latest remote values.
func (r *MapResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state MapResourceModel
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
func (r *MapResource) readRemote(ctx context.Context, state *MapResourceModel, resp *resource.ReadResponse) (removed bool) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/maps/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(state.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_map", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", state.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_map", fmt.Sprintf("Could not send request: %s", err))
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
			resp.Diagnostics.AddError("Error reading gigavuecore_map", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_map", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_map", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_map", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_map", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_map", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_map", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_map", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_map", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_map", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["map"]; ok {
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
		resp.Diagnostics.AddError("Error reading gigavuecore_map", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
	return
}

// Update modifies the remote resource to match the desired plan.
func (r *MapResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan MapResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state MapResourceModel
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
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// updateRemote performs the update HTTP exchange and decodes the response into plan. Extracted from Update so the request/response logic is unit-testable without a tfsdk.Plan.
func (r *MapResource) updateRemote(ctx context.Context, plan *MapResourceModel, resp *resource.UpdateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_map", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_map", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/maps/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(plan.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_map", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", plan.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_map", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error updating gigavuecore_map", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error updating gigavuecore_map", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error updating gigavuecore_map", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error updating gigavuecore_map", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error updating gigavuecore_map", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error updating gigavuecore_map", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error updating gigavuecore_map", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error updating gigavuecore_map", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error updating gigavuecore_map", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error updating gigavuecore_map", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_map", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Delete destroys the remote resource.
func (r *MapResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state MapResourceModel
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
func (r *MapResource) deleteRemote(ctx context.Context, state *MapResourceModel, resp *resource.DeleteResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/maps/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(state.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_map", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", state.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_map", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		return
	}
	if !(httpResp.StatusCode == 204) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error deleting gigavuecore_map", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error deleting gigavuecore_map", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error deleting gigavuecore_map", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error deleting gigavuecore_map", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error deleting gigavuecore_map", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error deleting gigavuecore_map", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error deleting gigavuecore_map", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error deleting gigavuecore_map", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error deleting gigavuecore_map", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *MapResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *MapResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	mapImportIDParts := strings.Split(req.ID, "/")
	if len(mapImportIDParts) != 2 {
		resp.Diagnostics.AddError("Unexpected Import Identifier", fmt.Sprintf("Expected import identifier with format \"{alias}/{cluster_id}\". Got %q.", req.ID))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("alias"), mapImportIDParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("cluster_id"), mapImportIDParts[1])...)
}
