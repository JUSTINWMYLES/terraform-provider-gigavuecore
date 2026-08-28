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
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	timeouts "github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	path "github.com/hashicorp/terraform-plugin-framework/path"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertions.
var (
	_ resource.Resource                = (*FabricMapResource)(nil)
	_ resource.ResourceWithImportState = (*FabricMapResource)(nil)
	_ resource.ResourceWithConfigure   = (*FabricMapResource)(nil)
)

// FabricMapResource is the generated Terraform managed resource implementation.
type FabricMapResource struct {
	client *client.Client
}

// FabricMapResourceModel describes the Terraform state and plan shape for FabricMapResource.
type FabricMapResourceModel struct {
	AfmMap             types.Object   `tfsdk:"afm_map" json:"afmMap"`
	Alias              types.String   `tfsdk:"alias"`
	ChildMapAliases    types.List     `tfsdk:"child_map_aliases" json:"childMapAliases"`
	ConfigStatus       types.String   `tfsdk:"config_status" json:"configStatus"`
	CreationTime       types.Float64  `tfsdk:"creation_time" json:"creationTime"`
	DecapAliases       types.List     `tfsdk:"decap_aliases" json:"decapAliases"`
	EncapAliases       types.List     `tfsdk:"encap_aliases" json:"encapAliases"`
	ErrorMessage       types.String   `tfsdk:"error_message" json:"errorMessage"`
	HealthState        types.String   `tfsdk:"health_state" json:"healthState"`
	HealthStateReasons types.List     `tfsdk:"health_state_reasons" json:"healthStateReasons"`
	ParentMapAliases   types.List     `tfsdk:"parent_map_aliases" json:"parentMapAliases"`
	PendingReason      types.String   `tfsdk:"pending_reason" json:"pendingReason"`
	UpdatedTime        types.Float64  `tfsdk:"updated_time" json:"updatedTime"`
	Timeouts           timeouts.Value `tfsdk:"timeouts"`
}

// Metadata returns the resource type name.
func (r *FabricMapResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_fabric_map"
}

// Schema returns the Terraform schema for this resource.
func (r *FabricMapResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get a user-defined fabric map by alias", Attributes: map[string]schema.Attribute{"afm_map": schema.SingleNestedAttribute{Required: true, Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "unique map alias", Required: true}, "ap_rules": schema.SingleNestedAttribute{MarkdownDescription: "pass and drop application profile rules", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"drop_rules": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"application_profile": schema.StringAttribute{MarkdownDescription: "application profile alias", Required: true}, "rule_id": schema.Int64Attribute{MarkdownDescription: "application profile rule Id, should not have same id as gsRules", Required: true}}}}, "pass_rules": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"application_profile": schema.StringAttribute{MarkdownDescription: "application profile alias", Required: true}, "rule_id": schema.Int64Attribute{MarkdownDescription: "application profile rule Id, should not have same id as gsRules", Required: true}}}}}}, "cluster_id": schema.StringAttribute{MarkdownDescription: "id of the defining cluster", Optional: true, Computed: true}, "comment": schema.StringAttribute{Optional: true, Computed: true}, "dst_ports": schema.ListAttribute{MarkdownDescription: "List of the 'to' ports. Invalid if 'nullDstPort' is true'. Only port number is supported", Required: true, ElementType: types.StringType}, "egress_gigastream": schema.ListAttribute{MarkdownDescription: "Applicable to fabric map only. It is to specify gigastream(s) to use as egress when going across cluster. The format of array member is clusterID:gigastreamAlias", Optional: true, Computed: true, ElementType: types.StringType}, "enable": schema.BoolAttribute{MarkdownDescription: "enable/disable map, applicable only to first level maps", Optional: true, Computed: true}, "encap_tunnel": schema.StringAttribute{MarkdownDescription: "tunnel alias", Optional: true, Computed: true}, "flex_inline": schema.SingleNestedAttribute{MarkdownDescription: "When 'oobCopy' is configured at least one of 'aToB' or 'bToA' must be configured", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"a_to_b": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"ib_pathway": schema.StringAttribute{MarkdownDescription: "ibPathway alias. Only applicable when type is 'ibPathway'", Optional: true, Computed: true}, "tools": schema.ListAttribute{MarkdownDescription: "ordered list of inline tools or vports. Only applicable when 'type' is 'tools'", Optional: true, Computed: true, ElementType: types.StringType}, "type": schema.StringAttribute{MarkdownDescription: "when set to 'tools', the list of the processing inline tools have to be provided. when set to 'ibPathway', ibPathway alias should be provided ", Required: true}}}, "b_to_a": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"ib_pathway": schema.StringAttribute{MarkdownDescription: "ibPathway alias. Only applicable when type is 'ibPathway'", Optional: true, Computed: true}, "tools": schema.ListAttribute{MarkdownDescription: "ordered list of inline tools or vports. Only applicable when 'type' is 'tools'", Optional: true, Computed: true, ElementType: types.StringType}, "type": schema.StringAttribute{MarkdownDescription: "when set to 'tools', the list of the processing inline tools have to be provided. when set to 'ibPathway', ibPathway alias should be provided ", Required: true}}}, "oob_copy": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"direction": schema.StringAttribute{Optional: true, Computed: true}, "dst_ports": schema.ListAttribute{MarkdownDescription: "list of destination tool ports", Required: true, ElementType: types.StringType}, "src_ports": schema.ListAttribute{MarkdownDescription: "inline network or an item from a-to-b and b-to-a lists", Required: true, ElementType: types.StringType}, "tag": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"type": schema.StringAttribute{Required: true}}}}}}, "svt_mode": schema.BoolAttribute{Optional: true, Computed: true}, "svt_tag": schema.Int64Attribute{MarkdownDescription: "only applicable when svtMode is enabled", Optional: true, Computed: true}, "tag": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"tag_protocol_id": schema.StringAttribute{MarkdownDescription: "When tool VLAN tag is added , this protocol Id will be added which egress out the traffic", Optional: true, Computed: true}, "type": schema.StringAttribute{Required: true}, "vlan_id": schema.Int64Attribute{MarkdownDescription: "only applicable when type is 'vlan'", Optional: true, Computed: true}}}}}, "flex_inline_failover": schema.StringAttribute{MarkdownDescription: "only valid for flexInline maps", Optional: true, Computed: true}, "flex_inline_vlan_id": schema.Int64Attribute{MarkdownDescription: "VLAN ID carried in the VLAN tag of packets coming from the inline-network port(s). valid and applicable only with 'flexInline' map type, 'srcPorts' should be either inline-network or ib-pathway or vport", Optional: true, Computed: true}, "flow_rules": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Rules Container. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"drop_rules": schema.SetNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"gtp": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Rule GTP match Definition. Private class", Required: true, Attributes: map[string]schema.Attribute{"imei": schema.StringAttribute{MarkdownDescription: "mutually exclusive with 'imsi' and 'msisdn'. If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true}, "imsi": schema.StringAttribute{MarkdownDescription: "mutually exclusive with 'imei' and 'msisdn'. If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true}, "interface": schema.StringAttribute{MarkdownDescription: "interface type. Mutually exclusive with version", Optional: true, Computed: true}, "msisdn": schema.StringAttribute{MarkdownDescription: "mutually exclusive with 'imsi' and 'imei'. If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true}, "version": schema.StringAttribute{MarkdownDescription: "mutually exclusive with interface", Optional: true, Computed: true}}}, "rule_id": schema.Int64Attribute{Required: true}}}}, "pass_rules": schema.SetNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"gtp": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Rule GTP match Definition. Private class", Required: true, Attributes: map[string]schema.Attribute{"imei": schema.StringAttribute{MarkdownDescription: "mutually exclusive with 'imsi' and 'msisdn'. If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true}, "imsi": schema.StringAttribute{MarkdownDescription: "mutually exclusive with 'imei' and 'msisdn'. If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true}, "interface": schema.StringAttribute{MarkdownDescription: "interface type. Mutually exclusive with version", Optional: true, Computed: true}, "msisdn": schema.StringAttribute{MarkdownDescription: "mutually exclusive with 'imsi' and 'imei'. If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true}, "version": schema.StringAttribute{MarkdownDescription: "mutually exclusive with interface", Optional: true, Computed: true}}}, "rule_id": schema.Int64Attribute{Required: true}}}}}}, "flow_sample5_g_overlap_rules": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Sample 5g Overlap Rules Container. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"pass_rules": schema.SetNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"comment": schema.StringAttribute{Optional: true, Computed: true}, "flow5_g": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Sample Overlap Rule 5g match Definition. Private class", Required: true, Attributes: map[string]schema.Attribute{"dnn": schema.StringAttribute{MarkdownDescription: "Access Point Name pattern.  Alphanumeric, '.', '-', and '*' allowed.", Optional: true, Computed: true}, "gpsi": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true}, "nas_5_qi": schema.StringAttribute{MarkdownDescription: "5G QoS Indicator. Valid 5qi value <1 - 255>", Optional: true, Computed: true}, "nci": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true}, "nsiid": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the SD value, it is treated as prefix", Optional: true, Computed: true}, "pei": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true}, "plmn_id": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the MNC value, it is treated as prefix", Optional: true, Computed: true}, "supi": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true}, "tac": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true}}}, "percentage": schema.Int64Attribute{Required: true}, "rule_id": schema.Int64Attribute{Required: true}}}}}}, "flow_sample5_g_rules": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Sample 5g Rules Container. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"pass_rules": schema.SetNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"comment": schema.StringAttribute{Optional: true, Computed: true}, "flow5_g": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Sample Rule 5g match Definition. Private class", Required: true, Attributes: map[string]schema.Attribute{"dnn": schema.StringAttribute{MarkdownDescription: "Access Point Name pattern.  Alphanumeric, '.', '-', and '*' allowed.", Optional: true, Computed: true}, "gpsi": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true}, "nci": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true}, "nsiid": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the SD value, it is treated as prefix", Optional: true, Computed: true}, "pei": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true}, "plmn_id": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the MNC value, it is treated as prefix", Optional: true, Computed: true}, "supi": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true}, "tac": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true}}}, "percentage": schema.Int64Attribute{Required: true}, "priority": schema.Int64Attribute{Optional: true, Computed: true}, "rule_id": schema.Int64Attribute{Required: true}}}}}}, "flow_sample_diameter_rules": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Sample Diameter Rules Container. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"pass_rules": schema.SetNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"diameter": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Sample Diameter Rule Definition", Required: true, Attributes: map[string]schema.Attribute{"user_name": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true}}}, "interface": schema.StringAttribute{MarkdownDescription: "interface type", Required: true}, "percentage": schema.Int64Attribute{Required: true}, "rule_id": schema.Int64Attribute{Required: true}}}}}}, "flow_sample_overlap_rules": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Sample Overlap Rules Container. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"pass_rules": schema.SetNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"comment": schema.StringAttribute{Optional: true, Computed: true}, "gtp": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Sample Rule GTP match Definition. Private class", Required: true, Attributes: map[string]schema.Attribute{"apn": schema.StringAttribute{MarkdownDescription: "Access Point Name pattern.  Alphanumeric, '.', '-', and '*' allowed.", Optional: true, Computed: true}, "eci": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true}, "imei": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true}, "imsi": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true}, "interface": schema.StringAttribute{MarkdownDescription: "interface type. Mutually exclusive with version", Optional: true, Computed: true}, "msisdn": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true}, "nas_5_qi": schema.StringAttribute{MarkdownDescription: "5G QoS Indicator. Valid 5qi value <1 - 255>", Optional: true, Computed: true}, "nci": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true}, "plmn_id": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the MNC value, it is treated as prefix", Optional: true, Computed: true}, "qci": schema.Int64Attribute{MarkdownDescription: "QoS Class Indicator", Optional: true, Computed: true}, "snssai": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the SD value, it is treated as prefix", Optional: true, Computed: true}, "tac": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true}, "tac_5_g": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true}, "version": schema.StringAttribute{MarkdownDescription: "mutually exclusive with interface", Optional: true, Computed: true}}}, "percentage": schema.Int64Attribute{Required: true}, "periodic_recalc": schema.BoolAttribute{MarkdownDescription: "Enable Periodic Recalc for rotational sampling. Map look up in the data path based on this flag", Optional: true, Computed: true}, "priority": schema.Int64Attribute{MarkdownDescription: "maximum value is equal to the number of rules upon completion of the request", Optional: true, Computed: true}, "rule_id": schema.Int64Attribute{Required: true}}}}}}, "flow_sample_rules": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Sample Rules Container. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"pass_rules": schema.SetNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"comment": schema.StringAttribute{Optional: true, Computed: true}, "gtp": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Sample Rule GTP match Definition. Private class", Required: true, Attributes: map[string]schema.Attribute{"apn": schema.StringAttribute{MarkdownDescription: "Access Point Name pattern.  Alphanumeric, '.', '-', and '*' allowed.", Optional: true, Computed: true}, "eci": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true}, "imei": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true}, "imsi": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true}, "interface": schema.StringAttribute{MarkdownDescription: "interface type. Mutually exclusive with version", Optional: true, Computed: true}, "msisdn": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true}, "nas_5_qi": schema.StringAttribute{MarkdownDescription: "5G QoS Indicator. Valid 5qi value <1 - 255>", Optional: true, Computed: true}, "nci": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true}, "plmn_id": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the MNC value, it is treated as prefix", Optional: true, Computed: true}, "qci": schema.Int64Attribute{MarkdownDescription: "QoS Class Indicator", Optional: true, Computed: true}, "snssai": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the SD value, it is treated as prefix", Optional: true, Computed: true}, "tac": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true}, "tac_5_g": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true}, "version": schema.StringAttribute{MarkdownDescription: "mutually exclusive with interface", Optional: true, Computed: true}}}, "percentage": schema.Int64Attribute{Required: true}, "periodic_recalc": schema.BoolAttribute{MarkdownDescription: "Enable Periodic Recalc for rotational sampling. Map look up in the data path based on this flag", Optional: true, Computed: true}, "priority": schema.Int64Attribute{MarkdownDescription: "maximum value is equal to the number of rules upon completion of the request", Optional: true, Computed: true}, "rule_id": schema.Int64Attribute{Required: true}}}}}}, "flow_sample_sip_rules": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Sample Sip Rules Container. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"pass_rules": schema.SetNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"percentage": schema.Int64Attribute{Required: true}, "rule_id": schema.Int64Attribute{Required: true}, "sip": schema.SingleNestedAttribute{Required: true, Attributes: map[string]schema.Attribute{"callee_id": schema.StringAttribute{MarkdownDescription: "sip callee id", Optional: true, Computed: true}, "callee_id_range": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"max_value": schema.StringAttribute{Required: true}, "value": schema.StringAttribute{Required: true}}}, "caller_id": schema.StringAttribute{MarkdownDescription: "sip caller id", Optional: true, Computed: true}, "caller_id_range": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"max_value": schema.StringAttribute{Required: true}, "value": schema.StringAttribute{Required: true}}}, "id_range": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"max_value": schema.StringAttribute{Required: true}, "value": schema.StringAttribute{Required: true}}}}}}}}}}, "flow_whitelist5_g_overlap_rules": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Whitelist 5g Overlap Rule match Definition. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"dnn": schema.StringAttribute{MarkdownDescription: "Domain Network Name pattern.  Alphanumeric, '.', '-', and '*' allowed.", Optional: true, Computed: true}, "type": schema.StringAttribute{MarkdownDescription: "Set 5G WL-DB lookup type", Optional: true, Computed: true}}}, "flow_whitelist5_g_rules": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Whitelist 5g Rule GTP match Definition. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"dnn": schema.StringAttribute{MarkdownDescription: "Domain Network Name pattern.  Alphanumeric, '.', '-', and '*' allowed.", Optional: true, Computed: true}, "type": schema.StringAttribute{MarkdownDescription: "Set 5G WL-DB lookup type", Optional: true, Computed: true}, "whitelist_databases": schema.ListAttribute{MarkdownDescription: "Attach whitelist databases to the map", Optional: true, Computed: true, ElementType: types.StringType}}}, "flow_whitelist_overlap_rules": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Whitelist Overlap Rules Container. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"pass_rules": schema.SetNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"flow5_g": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Whitelist 5g Rule GTP match Definition. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"dnn": schema.StringAttribute{MarkdownDescription: "Domain Network Name pattern.  Alphanumeric, '.', '-', and '*' allowed.", Optional: true, Computed: true}, "type": schema.StringAttribute{MarkdownDescription: "Set 5G WL-DB lookup type", Optional: true, Computed: true}, "whitelist_databases": schema.ListAttribute{MarkdownDescription: "Attach whitelist databases to the map", Optional: true, Computed: true, ElementType: types.StringType}}}, "gtp": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Whitelist Rule GTP match Definition. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"apn": schema.StringAttribute{MarkdownDescription: "Access Point Name pattern.  Alphanumeric, '.', '-', and '*' allowed.", Optional: true, Computed: true}, "interface": schema.StringAttribute{MarkdownDescription: "interface type. Mutually exclusive with version. required till H 5.6", Optional: true, Computed: true}, "type": schema.StringAttribute{MarkdownDescription: "Set GTP WL-DB lookup type", Optional: true, Computed: true}, "version": schema.StringAttribute{MarkdownDescription: "mutually exclusive with interface", Optional: true, Computed: true}, "whitelist_databases": schema.ListAttribute{MarkdownDescription: "Attach whitelist databases to the map", Optional: true, Computed: true, ElementType: types.StringType}}}, "rule_id": schema.Int64Attribute{Required: true}, "sip": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Whitelist Rule Sip match definition", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"type": schema.StringAttribute{MarkdownDescription: "all:Whitelist based on caller/callee/source/destination IP address, bothAddr: Whitelist source/destination IP address, bothId:Whitelist Caller/Callee Id's, calleeId: Whitelist Callee ID, callerId: Whitelist Caller ID, destIp: Whitelist based on destination IP address, srcIp: Whitelist based on Source IP address", Optional: true, Computed: true}}}}}}}}, "flow_whitelist_rules": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Whitelist Rules Container. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"pass_rules": schema.SetNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"flow5_g": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Whitelist 5g Rule GTP match Definition. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"dnn": schema.StringAttribute{MarkdownDescription: "Domain Network Name pattern.  Alphanumeric, '.', '-', and '*' allowed.", Optional: true, Computed: true}, "type": schema.StringAttribute{MarkdownDescription: "Set 5G WL-DB lookup type", Optional: true, Computed: true}, "whitelist_databases": schema.ListAttribute{MarkdownDescription: "Attach whitelist databases to the map", Optional: true, Computed: true, ElementType: types.StringType}}}, "gtp": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Whitelist Rule GTP match Definition. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"apn": schema.StringAttribute{MarkdownDescription: "Access Point Name pattern.  Alphanumeric, '.', '-', and '*' allowed.", Optional: true, Computed: true}, "interface": schema.StringAttribute{MarkdownDescription: "interface type. Mutually exclusive with version. required till H 5.6", Optional: true, Computed: true}, "type": schema.StringAttribute{MarkdownDescription: "Set GTP WL-DB lookup type", Optional: true, Computed: true}, "version": schema.StringAttribute{MarkdownDescription: "mutually exclusive with interface", Optional: true, Computed: true}, "whitelist_databases": schema.ListAttribute{MarkdownDescription: "Attach whitelist databases to the map", Optional: true, Computed: true, ElementType: types.StringType}}}, "rule_id": schema.Int64Attribute{Required: true}, "sip": schema.SingleNestedAttribute{MarkdownDescription: "Map Flow Whitelist Rule Sip match definition", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"type": schema.StringAttribute{MarkdownDescription: "all:Whitelist based on caller/callee/source/destination IP address, bothAddr: Whitelist source/destination IP address, bothId:Whitelist Caller/Callee Id's, calleeId: Whitelist Callee ID, callerId: Whitelist Caller ID, destIp: Whitelist based on destination IP address, srcIp: Whitelist based on Source IP address", Optional: true, Computed: true}}}}}}}}, "fstype": schema.SingleNestedAttribute{MarkdownDescription: "Type of Mobility Flowsampling & properties", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"offset": schema.Int64Attribute{MarkdownDescription: "Offset for Mobility Rotational Flowsampling", Optional: true, Computed: true}, "timer": schema.Int64Attribute{MarkdownDescription: "Timer for Mobility Rotational Flowsampling in minutes", Optional: true, Computed: true}, "type": schema.StringAttribute{Optional: true, Computed: true}}}, "gs_rules": schema.SingleNestedAttribute{MarkdownDescription: "Map GigaSMART Rules Container. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"drop_rules": schema.DynamicAttribute{Optional: true, Computed: true}, "pass_rules": schema.DynamicAttribute{Optional: true, Computed: true}}}, "gsop": schema.StringAttribute{MarkdownDescription: "Alias of referenced GSOP. Applicable to 'regular' and 'secondLevel' map types", Optional: true, Computed: true}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "health_state_reasons": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Computed: true}}}}, "inline_traffic_path": schema.StringAttribute{MarkdownDescription: "Only applicable for 'inline' map types, in which case defaults to 'normal'. When set to 'bypass', the 'dstPorts' must be empty", Optional: true, Computed: true}, "inline_traffic_type": schema.StringAttribute{MarkdownDescription: "Only applicable for 'inline/passAll' map types, in which case defaults to 'symmetric'. For 'asymmetric' maps, 'srcPort' must be of type 'inline-net'", Optional: true, Computed: true}, "ip_rewrite": schema.SingleNestedAttribute{MarkdownDescription: "IpRewrite options on the packets", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"dst_ip": schema.StringAttribute{Optional: true, Computed: true}, "src_ip": schema.StringAttribute{Optional: true, Computed: true}}}, "mod_time": schema.Int64Attribute{MarkdownDescription: "Last modification time of the map in milliseconds since the epoch", Optional: true, Computed: true}, "null_dst_port": schema.BoolAttribute{MarkdownDescription: "enabled when the dstPort is null", Optional: true, Computed: true}, "order": schema.Int64Attribute{MarkdownDescription: "relative order within per-source port map chain", Optional: true, Computed: true}, "rewrite": schema.SingleNestedAttribute{MarkdownDescription: "Rewrite options on the packets", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"dst_mac": schema.StringAttribute{Optional: true, Computed: true}, "src_mac": schema.StringAttribute{Optional: true, Computed: true}}}, "roles": schema.SingleNestedAttribute{MarkdownDescription: "Map Roles Container. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"editors": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType}, "listeners": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType}, "owners": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType}, "viewers": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType}}}, "rule_matching": schema.StringAttribute{MarkdownDescription: "If set to 'blacklist', packet are passed when no 'drop' rules are matched. This field is only valid for 'regular/byRule' map types. (maps into 'no-rule-match' CLI command)", Optional: true, Computed: true}, "rules": schema.SingleNestedAttribute{MarkdownDescription: "Map Rules Container. Private class", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"drop_rules": schema.DynamicAttribute{Optional: true, Computed: true}, "pass_rules": schema.DynamicAttribute{Optional: true, Computed: true}}}, "rx_cluster_ports": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType}, "src_ports": schema.ListAttribute{MarkdownDescription: "list of the 'from' ports. Only port number is supported", Required: true, ElementType: types.StringType}, "sub_type": schema.StringAttribute{MarkdownDescription: "'byRule' is applicable to all map types; 'collector' is applicable to 'regular', 'inline', 'flexInline' and 'secondLevel' maps; 'passAll' is applicable to 'regular' and 'inline' maps; 'flowFilter', 'flowSample', 'flowWhitelist', 'flowSampleSip', 'flowWhitelistSip', 'flowSampleDiameter', 'flowWhitelistDiameter', 'flowSampleOverlap', 'flowWhitelistOverlap', 'flowSample5g','flowWhitelist5g','flowSample5gOverlap' and 'flowWhitelist5gOverlap' are applicable to 'secondLevel' maps", Optional: true, Computed: true}, "traffic_type": schema.StringAttribute{MarkdownDescription: "Only applicable for 'firstLevel/byRule' map types, in which case defaults to 'user'", Optional: true, Computed: true}, "tx_cluster_ports": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType}, "type": schema.StringAttribute{MarkdownDescription: "'regular' maps are from network/hybrid ports to tool/hybrid/gigastream; 'inline' maps are from inline ports to inline/tool/hybrid/gigastream; 'firstLevel' are from network/hybrid ports to vPorts/tool/hybrid/gigastream; 'secondLevel' maps are from vPorts to tool/hybrid/gigastream; 'inlineFirstLevel is from inline network to vport; 'inlineSecondLevel' is from vport to inline tool; 'flexInline' is from inline network to tools; 'transitLevel' is from vport to vport", Required: true}, "updated_time": schema.Float64Attribute{MarkdownDescription: "Last Updated time of the fabric map", Computed: true}, "vlan_tag": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"tag_protocol_id": schema.StringAttribute{Optional: true, Computed: true}, "vlan_action": schema.StringAttribute{Required: true}, "vlan_id": schema.Int64Attribute{Optional: true, Computed: true}}}}}, "alias": schema.StringAttribute{Computed: true}, "child_map_aliases": schema.ListAttribute{MarkdownDescription: "Aliases of the fabric maps that support this fabric map. Applicable only to user-defined fabric maps. For internally generated fabric maps, there is no child maps. This list is empty.", Optional: true, Computed: true, ElementType: types.StringType}, "config_status": schema.StringAttribute{MarkdownDescription: "Configuration status of this fabric map.", Optional: true, Computed: true}, "creation_time": schema.Float64Attribute{MarkdownDescription: "Fabric map creation timestamp in UTC", Computed: true}, "decap_aliases": schema.ListAttribute{MarkdownDescription: "Aliases of the decap tunnel endpoints supporting this fabric map. Applicable only to internally generated fabric maps.", Optional: true, Computed: true, ElementType: types.StringType}, "encap_aliases": schema.ListAttribute{MarkdownDescription: "Aliases of the encap tunnel endpoints supporting this fabric map. Applicable only to internally generated fabric maps.", Optional: true, Computed: true, ElementType: types.StringType}, "error_message": schema.StringAttribute{MarkdownDescription: "In case of configuration failure, this message provides details about the possible cause of the failure.", Optional: true, Computed: true}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "health_state_reasons": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Computed: true}}}}, "parent_map_aliases": schema.ListAttribute{MarkdownDescription: "Aliases of the fabric maps that this fabric map supports. Applicable only to internally generated fabric maps that support user-defined fabric maps. For user-defined fabric maps, there is no parent maps. This list is empty.", Optional: true, Computed: true, ElementType: types.StringType}, "pending_reason": schema.StringAttribute{Computed: true}, "updated_time": schema.Float64Attribute{MarkdownDescription: "Fabric map updated timestamp in UTC", Computed: true}}, Blocks: map[string]schema.Block{"timeouts": timeouts.Block(ctx, timeouts.Opts{Create: true, Read: true, Update: true, Delete: true})}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *FabricMapResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan FabricMapResourceModel
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
func (r *FabricMapResource) createRemote(ctx context.Context, plan *FabricMapResourceModel, resp *resource.CreateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_fabric_map", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_fabric_map", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/fabricMaps"
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_fabric_map", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_fabric_map", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201 || httpResp.StatusCode == 202) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error creating gigavuecore_fabric_map", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error creating gigavuecore_fabric_map", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error creating gigavuecore_fabric_map", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error creating gigavuecore_fabric_map", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error creating gigavuecore_fabric_map", "Entity Already Exists. See errors payload for details")
			return
		case 417:
			resp.Diagnostics.AddError("Error creating gigavuecore_fabric_map", "Deployment failed. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error creating gigavuecore_fabric_map", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error creating gigavuecore_fabric_map", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error creating gigavuecore_fabric_map", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error creating gigavuecore_fabric_map", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error creating gigavuecore_fabric_map", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_fabric_map", fmt.Sprintf("Could not map response to state: %s", err))
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
			resp.Diagnostics.AddError("Error creating gigavuecore_fabric_map", "The create response did not contain an identifier and no Location header was returned, so the resource cannot be tracked in state.")
			return
		}
	}
}

// Read refreshes the Terraform state with the latest remote values.
func (r *FabricMapResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state FabricMapResourceModel
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
func (r *FabricMapResource) readRemote(ctx context.Context, state *FabricMapResourceModel, resp *resource.ReadResponse) (removed bool) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/fabricMaps/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(state.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_fabric_map", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_fabric_map", fmt.Sprintf("Could not send request: %s", err))
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
			resp.Diagnostics.AddError("Error reading gigavuecore_fabric_map", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_fabric_map", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_fabric_map", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_fabric_map", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_fabric_map", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_fabric_map", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_fabric_map", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_fabric_map", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_fabric_map", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_fabric_map", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["fabricMap"]; ok {
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
		resp.Diagnostics.AddError("Error reading gigavuecore_fabric_map", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
	return
}

// Update modifies the remote resource to match the desired plan.
func (r *FabricMapResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan FabricMapResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state FabricMapResourceModel
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
func (r *FabricMapResource) updateRemote(ctx context.Context, plan *FabricMapResourceModel, resp *resource.UpdateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_fabric_map", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_fabric_map", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/fabricMaps/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(plan.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_fabric_map", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_fabric_map", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error updating gigavuecore_fabric_map", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error updating gigavuecore_fabric_map", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error updating gigavuecore_fabric_map", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error updating gigavuecore_fabric_map", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error updating gigavuecore_fabric_map", "Entity Already Exists. See errors payload for details")
			return
		case 417:
			resp.Diagnostics.AddError("Error updating gigavuecore_fabric_map", "Deployment failed. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error updating gigavuecore_fabric_map", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error updating gigavuecore_fabric_map", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error updating gigavuecore_fabric_map", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error updating gigavuecore_fabric_map", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error updating gigavuecore_fabric_map", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_fabric_map", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Delete destroys the remote resource.
func (r *FabricMapResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state FabricMapResourceModel
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
func (r *FabricMapResource) deleteRemote(ctx context.Context, state *FabricMapResourceModel, resp *resource.DeleteResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/fabricMaps/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(state.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_fabric_map", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_fabric_map", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		return
	}
	if !(httpResp.StatusCode == 204) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error deleting gigavuecore_fabric_map", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error deleting gigavuecore_fabric_map", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error deleting gigavuecore_fabric_map", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error deleting gigavuecore_fabric_map", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error deleting gigavuecore_fabric_map", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error deleting gigavuecore_fabric_map", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error deleting gigavuecore_fabric_map", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error deleting gigavuecore_fabric_map", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error deleting gigavuecore_fabric_map", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *FabricMapResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *FabricMapResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("alias"), req.ID)...)
}
