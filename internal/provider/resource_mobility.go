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
	"strconv"
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
	_ resource.Resource                = (*MobilityResource)(nil)
	_ resource.ResourceWithIdentity    = (*MobilityResource)(nil)
	_ resource.ResourceWithImportState = (*MobilityResource)(nil)
	_ resource.ResourceWithConfigure   = (*MobilityResource)(nil)
)

// MobilityResource is the generated Terraform managed resource implementation.
type MobilityResource struct {
	client *client.Client
}

// MobilityResourceModel describes the Terraform state and plan shape for MobilityResource.
type MobilityResourceModel struct {
	Deployed           types.Bool     `tfsdk:"deployed"`
	HealthState        types.String   `tfsdk:"health_state" json:"healthState"`
	HealthStateReasons types.List     `tfsdk:"health_state_reasons" json:"healthStateReasons"`
	SiteName           types.String   `tfsdk:"site_name" json:"siteName"`
	SiteTag            types.String   `tfsdk:"site_tag" json:"siteTag"`
	Sites              types.Dynamic  `tfsdk:"sites"`
	SolutionAlias      types.String   `tfsdk:"solution_alias" json:"solutionAlias"`
	SolutionType       types.String   `tfsdk:"solution_type" json:"solutionType"`
	Tags               types.List     `tfsdk:"tags"`
	TrafficPolicies    types.Object   `tfsdk:"traffic_policies" json:"trafficPolicies"`
	Timeouts           timeouts.Value `tfsdk:"timeouts"`
}

// Metadata returns the resource type name.
func (r *MobilityResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_mobility"
}

// Schema returns the Terraform schema for this resource.
func (r *MobilityResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load mobility solutions by alias", Attributes: map[string]schema.Attribute{"deployed": schema.BoolAttribute{MarkdownDescription: "Mobility solution with gtpNodes, cpNodes and upNodes matching the requested deployed state is/are returned along with site configurations and trafficPolicies", Optional: true, Computed: true}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("green", "yellow", "orange", "red")}}, "health_state_reasons": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Computed: true}}}}, "site_name": schema.StringAttribute{MarkdownDescription: "mobility solution with requested site name is returned along with trafficPolicies", Optional: true}, "site_tag": schema.StringAttribute{MarkdownDescription: "Mobility solution with sites matching the requested tag values is/are returned along with trafficPolicies", Optional: true}, "sites": schema.DynamicAttribute{Required: true}, "solution_alias": schema.StringAttribute{MarkdownDescription: "Alias of the solution", Required: true}, "solution_type": schema.StringAttribute{MarkdownDescription: "Type of the solution", Required: true, Validators: []validator.String{stringvalidator.OneOf("NON_CUPS", "CUPS", "MIXED")}}, "tags": schema.ListNestedAttribute{MarkdownDescription: "RBAC Tags", Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"tag_key": schema.StringAttribute{MarkdownDescription: "Name of the tag", Required: true}, "tag_values": schema.ListAttribute{MarkdownDescription: "All possible values of the tag", Required: true, ElementType: types.StringType}}}}, "traffic_policies": schema.SingleNestedAttribute{MarkdownDescription: "Global forwarding policies for the processing nodes", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"for5_g": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"gtp_flow_timeout": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 6000)}}, "gtp_persistence": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Gtp Persistence Parameters", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{MarkdownDescription: "GTP Persistence Status", Optional: true, Computed: true}, "file_age_timeout": schema.Int64Attribute{MarkdownDescription: "GTP Persistence File Age Timeout(mins)", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(10, 1440)}}, "interval": schema.Int64Attribute{MarkdownDescription: "GTP Persistence Interval(mins)", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(10, 1440)}}, "restart_age_time": schema.Int64Attribute{MarkdownDescription: "GTP Persistence Restart Age Time(mins)", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(10, 1440)}}}}, "load_balancing": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"app_type": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.OneOf("flow5g")}}, "hashing_key": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.OneOf("supi", "pei", "gpsi")}}}}, "overlap_mode": schema.BoolAttribute{MarkdownDescription: "When enabled flow-filtering cannot be configured", Optional: true, Computed: true}, "sampling": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"flow_maps": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Alias of the sampling map", Required: true}, "comment": schema.StringAttribute{Optional: true, Computed: true}, "rules": schema.ListNestedAttribute{Required: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"comment": schema.StringAttribute{Optional: true, Computed: true}, "control_plane_percentage": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.Between(0, 100)}}, "dnn": schema.StringAttribute{MarkdownDescription: "Access Point Name pattern.  Alphanumeric, '.', '-', and '*' allowed.", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.LengthBetween(1, 128), stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9.*-]*$"), "value must match pattern \"^[a-zA-Z0-9.*-]*$\"")}}, "gpsi": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|[0-9]{6,15})|([0-9]{0,14}[*])$"), "value must match pattern \"^([*]|[0-9]{6,15})|([0-9]{0,14}[*])$\"")}}, "nas_5_qi": schema.StringAttribute{MarkdownDescription: "5G QoS Indicator", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|[0-9]{1,3})$"), "value must match pattern \"^([*]|[0-9]{1,3})$\"")}}, "nci": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|([A-Fa-f0-9]{9})|([A-Fa-f0-9]{4,8}[*]))$"), "value must match pattern \"^([*]|([A-Fa-f0-9]{9})|([A-Fa-f0-9]{4,8}[*]))$\"")}}, "nsiid": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the SD value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|([0-9]{1,3})|([0-9]{1,3}[.][a-fA-F0-9]{1,6})|([0-9]{1,3}[.][a-fA-F0-9]{0,5}[*]))$"), "value must match pattern \"^([*]|([0-9]{1,3})|([0-9]{1,3}[.][a-fA-F0-9]{1,6})|([0-9]{1,3}[.][a-fA-F0-9]{0,5}[*]))$\"")}}, "pei": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|[0-9]{6,15})|([0-9]{0,14}[*])$"), "value must match pattern \"^([*]|[0-9]{6,15})|([0-9]{0,14}[*])$\"")}}, "plmn_id": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the MNC value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|([0-9]{3}[.][0-9]{2,3})|([0-9]{3}[.][0-9]{0,2}[*]))$"), "value must match pattern \"^([*]|([0-9]{3}[.][0-9]{2,3})|([0-9]{3}[.][0-9]{0,2}[*]))$\"")}}, "supi": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|[0-9]{6,15})|([0-9]{0,14}[*])$"), "value must match pattern \"^([*]|[0-9]{6,15})|([0-9]{0,14}[*])$\"")}}, "tac": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|([0-9a-fA-F]{6})|([0-9a-fA-F]{2,5}[*]))$"), "value must match pattern \"^([*]|([0-9a-fA-F]{6})|([0-9a-fA-F]{2,5}[*]))$\"")}}, "user_plane_percentage": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.Between(0, 100)}}}}}, "source_group_id": schema.StringAttribute{MarkdownDescription: "if provided a vport with this sourceGroupId would be created else default vport would be used", Optional: true, Computed: true}, "tags": schema.ListNestedAttribute{MarkdownDescription: "User defined tags (Aggregation tags)", Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"tag_key": schema.StringAttribute{MarkdownDescription: "Name of the tag", Required: true}, "tag_values": schema.ListAttribute{MarkdownDescription: "All possible values of the tag", Required: true, ElementType: types.StringType}}}}, "tool": schema.StringAttribute{MarkdownDescription: "Alias of referenced tool", Required: true}}}}}}, "whitelisting": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"flow_maps": schema.ListNestedAttribute{Required: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Alias of the whitelist map", Required: true}, "comment": schema.StringAttribute{Optional: true, Computed: true}, "rules": schema.ListNestedAttribute{Required: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"dnn": schema.StringAttribute{MarkdownDescription: "Domain Network Name pattern.  Alphanumeric, '.', '-', and '*' allowed.", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.LengthBetween(1, 128), stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9.*-]*$"), "value must match pattern \"^[a-zA-Z0-9.*-]*$\"")}}, "type": schema.StringAttribute{MarkdownDescription: "Set 5G WL-DB lookup type", Optional: true, Computed: true}, "whitelist_databases": schema.ListAttribute{MarkdownDescription: "Attach whitelist databases to the map", Optional: true, Computed: true, ElementType: types.StringType}}}}, "source_group_id": schema.StringAttribute{MarkdownDescription: "if provided a vport with this sourceGroupId would be created else default vport would be used", Optional: true, Computed: true}, "tags": schema.ListNestedAttribute{MarkdownDescription: "User defined tags (Aggregation tags)", Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"tag_key": schema.StringAttribute{MarkdownDescription: "Name of the tag", Required: true}, "tag_values": schema.ListAttribute{MarkdownDescription: "All possible values of the tag", Required: true, ElementType: types.StringType}}}}, "tool": schema.StringAttribute{MarkdownDescription: "Alias of referenced tool", Required: true}}}}, "multi_whitelists": schema.ListAttribute{MarkdownDescription: "Alias/Aliases of referenced GTP Whitelists.Use either whiteListAlias or multiWhitelists.", Optional: true, Computed: true, ElementType: types.StringType}, "white_list_alias": schema.StringAttribute{MarkdownDescription: "Alias of referenced GTP Whitelist.Use either whiteListAlias or multiWhitelists.", Optional: true, Computed: true}}}}}, "for_lte": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"gtp_flow_timeout": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 6000)}}, "gtp_persistence": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Gtp Persistence Parameters", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{MarkdownDescription: "GTP Persistence Status", Optional: true, Computed: true}, "file_age_timeout": schema.Int64Attribute{MarkdownDescription: "GTP Persistence File Age Timeout(mins)", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(10, 1440)}}, "interval": schema.Int64Attribute{MarkdownDescription: "GTP Persistence Interval(mins)", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(10, 1440)}}, "restart_age_time": schema.Int64Attribute{MarkdownDescription: "GTP Persistence Restart Age Time(mins)", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(10, 1440)}}}}, "load_balancing_lte": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"app_type": schema.StringAttribute{Computed: true}, "hashing_key": schema.StringAttribute{Computed: true}}}, "overlap_mode": schema.BoolAttribute{MarkdownDescription: "When enabled flow-filtering cannot be configured", Optional: true, Computed: true}, "sampling": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"flow_maps": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Alias of the sampling map", Required: true}, "comment": schema.StringAttribute{Optional: true, Computed: true}, "rules": schema.ListNestedAttribute{Required: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"apn": schema.StringAttribute{MarkdownDescription: "Access Point Name pattern.  Alphanumeric, '.', '-', and '*' allowed.", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.LengthBetween(1, 128), stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9.*-]*$"), "value must match pattern \"^[a-zA-Z0-9.*-]*$\"")}}, "comment": schema.StringAttribute{Optional: true, Computed: true}, "control_plane_percentage": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.Between(0, 100)}}, "eci": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|([A-Fa-f0-9]{4,9})|([A-Fa-f0-9]{3,8}[*]))$"), "value must match pattern \"^([*]|([A-Fa-f0-9]{4,9})|([A-Fa-f0-9]{3,8}[*]))$\"")}}, "imei": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|[0-9]{6,16})|([0-9]{0,14}[*])$"), "value must match pattern \"^([*]|[0-9]{6,16})|([0-9]{0,14}[*])$\"")}}, "imsi": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|[0-9]{6,15})|([0-9]{0,14}[*])$"), "value must match pattern \"^([*]|[0-9]{6,15})|([0-9]{0,14}[*])$\"")}}, "interface": schema.StringAttribute{MarkdownDescription: "interface type. Mutually exclusive with version", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("Gn", "S5", "S10", "S11", "S2b", "S11U", "S1U", "S5S8U", "N3", "N9")}}, "msisdn": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|[0-9]{6,15})|([0-9]{0,14}[*])$"), "value must match pattern \"^([*]|[0-9]{6,15})|([0-9]{0,14}[*])$\"")}}, "nas_5_qi": schema.StringAttribute{MarkdownDescription: "5G QoS Indicator", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|[0-9]{1,3})$"), "value must match pattern \"^([*]|[0-9]{1,3})$\"")}}, "nci": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|([A-Fa-f0-9]{9})|([A-Fa-f0-9]{4,8}[*]))$"), "value must match pattern \"^([*]|([A-Fa-f0-9]{9})|([A-Fa-f0-9]{4,8}[*]))$\"")}}, "periodic_recalc": schema.BoolAttribute{MarkdownDescription: "Enable Periodic Recalc for rotational sampling. Map look up in the data path based on this flag", Optional: true, Computed: true}, "plmn_id": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the MNC value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|([0-9]{3}[.][0-9]{2,3})|([0-9]{3}[.][0-9]{0,2}[*]))$"), "value must match pattern \"^([*]|([0-9]{3}[.][0-9]{2,3})|([0-9]{3}[.][0-9]{0,2}[*]))$\"")}}, "qci": schema.Int64Attribute{MarkdownDescription: "QoS Class Indicator", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 255)}}, "snssai": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the SD value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|([0-9]{1,3})|([0-9]{1,3}[.][a-fA-F0-9]{1,6})|([0-9]{1,3}[.][a-fA-F0-9]{0,5}[*]))$"), "value must match pattern \"^([*]|([0-9]{1,3})|([0-9]{1,3}[.][a-fA-F0-9]{1,6})|([0-9]{1,3}[.][a-fA-F0-9]{0,5}[*]))$\"")}}, "tac": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|([0-9a-fA-F]{4})|([0-9a-fA-F]{2,3}[*]))$"), "value must match pattern \"^([*]|([0-9a-fA-F]{4})|([0-9a-fA-F]{2,3}[*]))$\"")}}, "tac_5_g": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|([0-9a-fA-F]{6})|([0-9a-fA-F]{2,5}[*]))$"), "value must match pattern \"^([*]|([0-9a-fA-F]{6})|([0-9a-fA-F]{2,5}[*]))$\"")}}, "user_plane_percentage": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.Between(0, 100)}}, "version": schema.StringAttribute{MarkdownDescription: "mutually exclusive with interface", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("any", "v1", "v2")}}}}}, "source_group_id": schema.StringAttribute{MarkdownDescription: "if provided a vport with this sourceGroupId would be created else default vport would be used", Optional: true, Computed: true}, "tags": schema.ListNestedAttribute{MarkdownDescription: "User defined tags (Aggregation tags)", Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"tag_key": schema.StringAttribute{MarkdownDescription: "Name of the tag", Required: true}, "tag_values": schema.ListAttribute{MarkdownDescription: "All possible values of the tag", Required: true, ElementType: types.StringType}}}}, "tool": schema.StringAttribute{MarkdownDescription: "Alias of referenced tool", Required: true}}}}}}, "whitelisting": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"flow_maps": schema.ListNestedAttribute{Required: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Alias of the whitelist map", Required: true}, "comment": schema.StringAttribute{Optional: true, Computed: true}, "rules": schema.ListNestedAttribute{Required: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"apn": schema.StringAttribute{MarkdownDescription: "Access Point Name pattern.  Alphanumeric, '.', '-', and '*' allowed.", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.LengthBetween(1, 128), stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9.*-]*$"), "value must match pattern \"^[a-zA-Z0-9.*-]*$\"")}}, "interface": schema.StringAttribute{MarkdownDescription: "interface type. Mutually exclusive with version. required till H 5.6", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("Gn", "S5", "S10", "S11", "S2b", "S11U", "S1U", "S5S8U", "N3", "N9")}}, "type": schema.StringAttribute{MarkdownDescription: "Set GTP WL-DB lookup type", Optional: true, Computed: true}, "version": schema.StringAttribute{MarkdownDescription: "mutually exclusive with interface", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("v1", "v2")}}, "whitelist_databases": schema.ListAttribute{MarkdownDescription: "Attach whitelist databases to the map", Optional: true, Computed: true, ElementType: types.StringType}}}}, "source_group_id": schema.StringAttribute{MarkdownDescription: "if provided a vport with this sourceGroupId would be created else default vport would be used", Optional: true, Computed: true}, "tool": schema.StringAttribute{MarkdownDescription: "Alias of referenced tool", Required: true}}}}, "multi_whitelists": schema.ListAttribute{MarkdownDescription: "Alias/Aliases of referenced GTP Whitelists.Use either whiteListAlias or multiWhitelists.", Optional: true, Computed: true, ElementType: types.StringType}, "white_list_alias": schema.StringAttribute{MarkdownDescription: "Alias of referenced GTP Whitelist.Use either whiteListAlias or multiWhitelists.", Optional: true, Computed: true}}}}}, "for_non_cups_lte": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"flowfiltering": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"flow_maps": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Alias of the flow-filtering map", Required: true}, "comment": schema.StringAttribute{Optional: true, Computed: true}, "drop_rules": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"imei": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|[0-9]{6,16})|([0-9]{0,14}[*])$"), "value must match pattern \"^([*]|[0-9]{6,16})|([0-9]{0,14}[*])$\"")}}, "imsi": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|[0-9]{6,15})|([0-9]{0,14}[*])$"), "value must match pattern \"^([*]|[0-9]{6,15})|([0-9]{0,14}[*])$\"")}}, "interface": schema.StringAttribute{MarkdownDescription: "interface type. Mutually exclusive with version", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("Gn", "S5", "S10", "S11", "S2b", "S11U")}}, "msisdn": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|[0-9]{6,15})|([0-9]{0,14}[*])$"), "value must match pattern \"^([*]|[0-9]{6,15})|([0-9]{0,14}[*])$\"")}}, "version": schema.StringAttribute{MarkdownDescription: "mutually exclusive with interface", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("v1", "v2")}}}}}, "pass_rules": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"imei": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|[0-9]{6,16})|([0-9]{0,14}[*])$"), "value must match pattern \"^([*]|[0-9]{6,16})|([0-9]{0,14}[*])$\"")}}, "imsi": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|[0-9]{6,15})|([0-9]{0,14}[*])$"), "value must match pattern \"^([*]|[0-9]{6,15})|([0-9]{0,14}[*])$\"")}}, "interface": schema.StringAttribute{MarkdownDescription: "interface type. Mutually exclusive with version", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("Gn", "S5", "S10", "S11", "S2b", "S11U")}}, "msisdn": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|[0-9]{6,15})|([0-9]{0,14}[*])$"), "value must match pattern \"^([*]|[0-9]{6,15})|([0-9]{0,14}[*])$\"")}}, "version": schema.StringAttribute{MarkdownDescription: "mutually exclusive with interface", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("v1", "v2")}}}}}, "source_group_id": schema.StringAttribute{MarkdownDescription: "if provided a vport with this sourceGroupId would be created else default vport would be used", Optional: true, Computed: true}, "tags": schema.ListNestedAttribute{MarkdownDescription: "User defined tags (Aggregation tags)", Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"tag_key": schema.StringAttribute{MarkdownDescription: "Name of the tag", Required: true}, "tag_values": schema.ListAttribute{MarkdownDescription: "All possible values of the tag", Required: true, ElementType: types.StringType}}}}, "tool": schema.StringAttribute{MarkdownDescription: "Alias of referenced tool", Required: true}}}}}}, "gtp_flow_timeout": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 6000)}}, "gtp_persistence": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Gtp Persistence Parameters", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{MarkdownDescription: "GTP Persistence Status", Optional: true, Computed: true}, "file_age_timeout": schema.Int64Attribute{MarkdownDescription: "GTP Persistence File Age Timeout(mins)", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(10, 1440)}}, "interval": schema.Int64Attribute{MarkdownDescription: "GTP Persistence Interval(mins)", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(10, 1440)}}, "restart_age_time": schema.Int64Attribute{MarkdownDescription: "GTP Persistence Restart Age Time(mins)", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(10, 1440)}}}}, "load_balancing": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"app_type": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.OneOf("gtp")}}, "hashing_key": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.OneOf("imsi", "imei", "msisdn")}}}}, "overlap_mode": schema.BoolAttribute{MarkdownDescription: "When enabled flow-filtering cannot be configured", Optional: true, Computed: true}, "sampling": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"flow_maps": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Alias of the sampling map", Required: true}, "comment": schema.StringAttribute{Optional: true, Computed: true}, "rules": schema.ListNestedAttribute{Required: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"apn": schema.StringAttribute{MarkdownDescription: "Access Point Name pattern.  Alphanumeric, '.', '-', and '*' allowed.", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.LengthBetween(1, 128), stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9.*-]*$"), "value must match pattern \"^[a-zA-Z0-9.*-]*$\"")}}, "comment": schema.StringAttribute{Optional: true, Computed: true}, "eci": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|([A-Fa-f0-9]{4,9})|([A-Fa-f0-9]{3,8}[*]))$"), "value must match pattern \"^([*]|([A-Fa-f0-9]{4,9})|([A-Fa-f0-9]{3,8}[*]))$\"")}}, "gtp_sample_percentage": schema.Int64Attribute{Required: true, Validators: []validator.Int64{int64validator.Between(0, 100)}}, "imei": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|[0-9]{6,16})|([0-9]{0,14}[*])$"), "value must match pattern \"^([*]|[0-9]{6,16})|([0-9]{0,14}[*])$\"")}}, "imsi": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|[0-9]{6,15})|([0-9]{0,14}[*])$"), "value must match pattern \"^([*]|[0-9]{6,15})|([0-9]{0,14}[*])$\"")}}, "interface": schema.StringAttribute{MarkdownDescription: "interface type. Mutually exclusive with version", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("Gn", "S5", "S10", "S11", "S2b", "S11U")}}, "msisdn": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|[0-9]{6,15})|([0-9]{0,14}[*])$"), "value must match pattern \"^([*]|[0-9]{6,15})|([0-9]{0,14}[*])$\"")}}, "periodic_recalc": schema.BoolAttribute{MarkdownDescription: "Enable Periodic Recalc for rotational sampling. Map look up in the data path based on this flag", Optional: true, Computed: true}, "plmn_id": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the MNC value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|([0-9]{3}[.][0-9]{2,3})|([0-9]{3}[.][0-9]{0,2}[*]))$"), "value must match pattern \"^([*]|([0-9]{3}[.][0-9]{2,3})|([0-9]{3}[.][0-9]{0,2}[*]))$\"")}}, "qci": schema.Int64Attribute{MarkdownDescription: "QoS Class Indicator", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 255)}}, "tac": schema.StringAttribute{MarkdownDescription: "If '*' is added at the end of the value, it is treated as prefix", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^([*]|([0-9a-fA-F]{4})|([0-9a-fA-F]{2,3}[*]))$"), "value must match pattern \"^([*]|([0-9a-fA-F]{4})|([0-9a-fA-F]{2,3}[*]))$\"")}}, "version": schema.StringAttribute{MarkdownDescription: "mutually exclusive with interface", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("any", "v1", "v2")}}}}}, "source_group_id": schema.StringAttribute{MarkdownDescription: "if provided a vport with this sourceGroupId would be created else default vport would be used", Optional: true, Computed: true}, "tags": schema.ListNestedAttribute{MarkdownDescription: "User defined tags (Aggregation tags)", Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"tag_key": schema.StringAttribute{MarkdownDescription: "Name of the tag", Required: true}, "tag_values": schema.ListAttribute{MarkdownDescription: "All possible values of the tag", Required: true, ElementType: types.StringType}}}}, "tool": schema.StringAttribute{MarkdownDescription: "Alias of referenced tool", Required: true}}}}}}, "whitelisting": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"flow_maps": schema.ListNestedAttribute{Required: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Alias of the whitelist map", Required: true}, "comment": schema.StringAttribute{Optional: true, Computed: true}, "rules": schema.ListNestedAttribute{Required: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"apn": schema.StringAttribute{MarkdownDescription: "Access Point Name pattern.  Alphanumeric, '.', '-', and '*' allowed.", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.LengthBetween(1, 128), stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9.*-]*$"), "value must match pattern \"^[a-zA-Z0-9.*-]*$\"")}}, "interface": schema.StringAttribute{MarkdownDescription: "interface type. Mutually exclusive with version. required till H 5.6", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("Gn", "S5", "S10", "S11", "S2b", "S11U", "S1U", "S5S8U", "N3", "N9")}}, "type": schema.StringAttribute{MarkdownDescription: "Set GTP WL-DB lookup type", Optional: true, Computed: true}, "version": schema.StringAttribute{MarkdownDescription: "mutually exclusive with interface", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("v1", "v2")}}, "whitelist_databases": schema.ListAttribute{MarkdownDescription: "Attach whitelist databases to the map", Optional: true, Computed: true, ElementType: types.StringType}}}}, "source_group_id": schema.StringAttribute{MarkdownDescription: "if provided a vport with this sourceGroupId would be created else default vport would be used", Optional: true, Computed: true}, "tags": schema.ListNestedAttribute{MarkdownDescription: "User defined tags (Aggregation tags)", Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"tag_key": schema.StringAttribute{MarkdownDescription: "Name of the tag", Required: true}, "tag_values": schema.ListAttribute{MarkdownDescription: "All possible values of the tag", Required: true, ElementType: types.StringType}}}}, "tool": schema.StringAttribute{MarkdownDescription: "Alias of referenced tool", Required: true}}}}, "multi_whitelists": schema.ListAttribute{MarkdownDescription: "Alias/Aliases of referenced GTP Whitelists.Use either whiteListAlias or multiWhitelists.", Optional: true, Computed: true, ElementType: types.StringType}, "white_list_alias": schema.StringAttribute{MarkdownDescription: "Alias of referenced GTP Whitelist.Use either whiteListAlias or multiWhitelists.", Optional: true, Computed: true}}}}}}}}, Blocks: map[string]schema.Block{"timeouts": timeouts.Block(ctx, timeouts.Opts{Create: true, Read: true, Update: true, Delete: true})}}
}

// IdentitySchema returns the resource identity schema shared with the paired list resource.
func (r *MobilityResource) IdentitySchema(_ context.Context, _ resource.IdentitySchemaRequest, resp *resource.IdentitySchemaResponse) {
	resp.IdentitySchema = identityschema.Schema{Attributes: map[string]identityschema.Attribute{"solution_alias": identityschema.StringAttribute{RequiredForImport: true, Description: "Alias of the solution"}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *MobilityResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan MobilityResourceModel
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
func (r *MobilityResource) createRemote(ctx context.Context, plan *MobilityResourceModel, resp *resource.CreateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_mobility", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_mobility", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/intent/mobility"
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_mobility", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_mobility", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 207) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error creating gigavuecore_mobility", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error creating gigavuecore_mobility", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error creating gigavuecore_mobility", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error creating gigavuecore_mobility", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error creating gigavuecore_mobility", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error creating gigavuecore_mobility", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error creating gigavuecore_mobility", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error creating gigavuecore_mobility", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error creating gigavuecore_mobility", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error creating gigavuecore_mobility", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_mobility", fmt.Sprintf("Could not map response to state: %s", err))
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
			resp.Diagnostics.AddError("Error creating gigavuecore_mobility", "The create response did not contain an identifier and no Location header was returned, so the resource cannot be tracked in state.")
			return
		}
	}
}

// Read refreshes the Terraform state with the latest remote values.
func (r *MobilityResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state MobilityResourceModel
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
func (r *MobilityResource) readRemote(ctx context.Context, state *MobilityResourceModel, resp *resource.ReadResponse) (removed bool) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/intent/mobility/{solutionAlias}"
	reqPath = strings.ReplaceAll(reqPath, "{solutionAlias}", url.PathEscape(state.SolutionAlias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_mobility", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !state.SiteName.IsNull() {
		query.Set("siteName", state.SiteName.ValueString())
	}
	if !state.HealthState.IsNull() {
		query.Set("healthState", state.HealthState.ValueString())
	}
	if !state.Deployed.IsNull() {
		query.Set("deployed", strconv.FormatBool(state.Deployed.ValueBool()))
	}
	if !state.SiteTag.IsNull() {
		query.Set("siteTag", state.SiteTag.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_mobility", fmt.Sprintf("Could not send request: %s", err))
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
			resp.Diagnostics.AddError("Error reading gigavuecore_mobility", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_mobility", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_mobility", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_mobility", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_mobility", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_mobility", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_mobility", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_mobility", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_mobility", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_mobility", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&state, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_mobility", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
	return
}

// Update modifies the remote resource to match the desired plan.
func (r *MobilityResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan MobilityResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state MobilityResourceModel
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
func (r *MobilityResource) updateRemote(ctx context.Context, plan *MobilityResourceModel, resp *resource.UpdateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_mobility", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_mobility", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/intent/mobility/{solutionAlias}"
	reqPath = strings.ReplaceAll(reqPath, "{solutionAlias}", url.PathEscape(plan.SolutionAlias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_mobility", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_mobility", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 207) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error updating gigavuecore_mobility", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error updating gigavuecore_mobility", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error updating gigavuecore_mobility", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error updating gigavuecore_mobility", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error updating gigavuecore_mobility", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error updating gigavuecore_mobility", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error updating gigavuecore_mobility", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error updating gigavuecore_mobility", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error updating gigavuecore_mobility", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error updating gigavuecore_mobility", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_mobility", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Delete destroys the remote resource.
func (r *MobilityResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state MobilityResourceModel
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
func (r *MobilityResource) deleteRemote(ctx context.Context, state *MobilityResourceModel, resp *resource.DeleteResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/intent/mobility/{solutionAlias}"
	reqPath = strings.ReplaceAll(reqPath, "{solutionAlias}", url.PathEscape(state.SolutionAlias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_mobility", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_mobility", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		return
	}
	if !(httpResp.StatusCode == 207) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error deleting gigavuecore_mobility", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error deleting gigavuecore_mobility", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error deleting gigavuecore_mobility", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error deleting gigavuecore_mobility", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error deleting gigavuecore_mobility", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error deleting gigavuecore_mobility", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error deleting gigavuecore_mobility", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error deleting gigavuecore_mobility", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error deleting gigavuecore_mobility", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *MobilityResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *MobilityResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("solution_alias"), req.ID)...)
}
