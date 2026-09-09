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
	int64validator "github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	listvalidator "github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
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
	_ resource.Resource                = (*InlineSslAppResource)(nil)
	_ resource.ResourceWithIdentity    = (*InlineSslAppResource)(nil)
	_ resource.ResourceWithImportState = (*InlineSslAppResource)(nil)
	_ resource.ResourceWithConfigure   = (*InlineSslAppResource)(nil)
)

// InlineSslAppResource is the generated Terraform managed resource implementation.
type InlineSslAppResource struct {
	client *client.Client
}

// InlineSslAppResourceModel describes the Terraform state and plan shape for InlineSslAppResource.
type InlineSslAppResourceModel struct {
	Alias               types.String               `tfsdk:"alias"`
	AppIntentConfigs    types.Object               `tfsdk:"app_intent_configs" json:"appIntentConfigs"`
	ClusterId           types.String               `tfsdk:"cluster_id" json:"clusterId"`
	ClusterName         types.String               `tfsdk:"cluster_name" json:"clusterName"`
	ConfigStatus        types.String               `tfsdk:"config_status" json:"configStatus"`
	ConfigStatusReasons types.List                 `tfsdk:"config_status_reasons" json:"configStatusReasons"`
	HealthState         types.String               `tfsdk:"health_state" json:"healthState"`
	HealthStateReasons  types.List                 `tfsdk:"health_state_reasons" json:"healthStateReasons"`
	MTls                types.String               `tfsdk:"m_tls" json:"mTLS"`
	RiaConfigs          types.List                 `tfsdk:"ria_configs" json:"riaConfigs"`
	RiaEnabled          types.String               `tfsdk:"ria_enabled" json:"riaEnabled"`
	Timeouts            *InlineSslAppTimeoutsModel `tfsdk:"timeouts"`
}

// InlineSslAppTimeoutsModel describes the per-operation timeout configuration (in seconds) for the gigavuecore_inline_ssl_app resource.
type InlineSslAppTimeoutsModel struct {
	Create types.Int64 `tfsdk:"create"`
	Read   types.Int64 `tfsdk:"read"`
	Update types.Int64 `tfsdk:"update"`
	Delete types.Int64 `tfsdk:"delete"`
}

// Metadata returns the resource type name.
func (r *InlineSslAppResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_inline_ssl_app"
}

// Schema returns the Terraform schema for this resource.
func (r *InlineSslAppResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Create an inline ssl app", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Alias of the inline ssl app", Optional: true, Computed: true}, "app_intent_configs": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"black_list_config": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"operation": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("ADD", "REPLACE", "DELETE")}}, "profile_list": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"file_source": schema.SingleNestedAttribute{MarkdownDescription: "Remote file source or destination", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"hostname": schema.StringAttribute{MarkdownDescription: "server address", Required: true}, "password": schema.StringAttribute{MarkdownDescription: "password to use for server login", Optional: true, Computed: true, Sensitive: true}, "path": schema.StringAttribute{MarkdownDescription: "file path on server", Required: true}, "protocol": schema.StringAttribute{MarkdownDescription: "Protocol to access server. HTTP and HTTPS only applicable for retrieving file", Required: true, Validators: []validator.String{stringvalidator.OneOf("scp", "sftp", "ftp", "tftp", "http", "https")}}, "username": schema.StringAttribute{MarkdownDescription: "user name to use for server login", Optional: true, Computed: true}}}, "list": schema.StringAttribute{MarkdownDescription: "The nodecryptlist or decryptlist . Mutually exclusive with 'fileSource'", Optional: true, Computed: true}}}}}, "global_default_configs": schema.SingleNestedAttribute{MarkdownDescription: "inline SSL configuration parameters", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"caching": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"persistence": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enable": schema.BoolAttribute{Optional: true, Computed: true}}}}}, "cluster_id": schema.StringAttribute{MarkdownDescription: "id of the defining cluster", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "dhe_ciphersuit": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("disable", "enable")}}, "monitor": schema.SingleNestedAttribute{MarkdownDescription: "This property is moved to ssl profile for device version >=5.7", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enable": schema.BoolAttribute{Optional: true, Computed: true}}}, "resumption": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"client": schema.SingleNestedAttribute{MarkdownDescription: "enable client initiated resumption (for debug purposes only)", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enable": schema.BoolAttribute{Optional: true, Computed: true}}}}}, "ssl_versions": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"connection_reset_action_for_max_version": schema.StringAttribute{MarkdownDescription: "Action to take to reset connection if TLS version is higher than configured", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("no-decrypt", "drop")}}, "connection_reset_action_for_min_version": schema.StringAttribute{MarkdownDescription: "Action to take to reset connection if TLS version is lower than configured", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("no-decrypt", "drop")}}, "max_version": schema.StringAttribute{MarkdownDescription: "maxVersion must be greater than minVersion", Required: true, Validators: []validator.String{stringvalidator.OneOf("sslv3", "tls1", "tls11", "tls12", "tls13")}}, "min_version": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.OneOf("sslv3", "tls1", "tls11", "tls12", "tls13")}}}}, "start_tls": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enable": schema.BoolAttribute{Optional: true, Computed: true}}}}}, "gs_engines": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType}, "gs_group_alias": schema.StringAttribute{MarkdownDescription: "Alias of the gs group auto created by ssl app", Optional: true, Computed: true}, "gs_group_param_configs": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"hsm_group": schema.StringAttribute{MarkdownDescription: "Alias of the hsmGroup", Optional: true, Computed: true}, "session_logging": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Session Logging Configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"interface": schema.StringAttribute{MarkdownDescription: "Associated IP Interface", Optional: true, Computed: true}, "log_level": schema.StringAttribute{MarkdownDescription: "Log Level", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("err", "warning", "notice", "info", "debug", "none")}}, "remote_syslog_ip": schema.StringAttribute{MarkdownDescription: "Remote Syslog IP", Optional: true, Computed: true}, "remote_syslog_port": schema.Int64Attribute{MarkdownDescription: "Remote Syslog Port Number", Optional: true, Computed: true}}}}}, "gsop_alias": schema.StringAttribute{MarkdownDescription: "Alias of the gsop auto created by ssl app", Optional: true, Computed: true}, "inline_ssl": schema.SingleNestedAttribute{MarkdownDescription: "Used to configure other GS apps in addition to Inline SSL on a HC1 box", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"standalone": schema.BoolAttribute{MarkdownDescription: "If enabled , behaves in a normal way. If disabled, can configure other GS apps along side Inline ssl with Inline ssl occupying only 50% of the memory", Optional: true, Computed: true}}}, "key_store_configs": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"deployment_type": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("Inbound", "Outbound", "Hybrid")}}, "inbound_keys": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"key_alias": schema.StringAttribute{MarkdownDescription: "Key to decrypt the traffic for the associated server/domain address", Optional: true, Computed: true}, "server_domain_alias": schema.StringAttribute{Optional: true, Computed: true}}}}, "outboundkeys": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"key_alias": schema.StringAttribute{Optional: true, Computed: true}, "signing_for": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("Primary", "Secondary")}}}}}}}, "m_tls_configs": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"primary_signing": schema.StringAttribute{MarkdownDescription: "Alias of the key from keyStore used for primary signing for mTLS", Optional: true, Computed: true}, "secondary_signing": schema.StringAttribute{MarkdownDescription: "Alias of the key from keyStore used for secondary signing for mTLS", Optional: true, Computed: true}, "trust_store": schema.StringAttribute{MarkdownDescription: "Alias of the client trust store", Optional: true, Computed: true}}}, "network_access_configs": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"network_access": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "id of the defining cluster", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "dhcp": schema.BoolAttribute{Optional: true, Computed: true}, "dns": schema.StringAttribute{MarkdownDescription: "Only valid and required when 'dhcp' is false", Optional: true, Computed: true}, "eport": schema.StringAttribute{MarkdownDescription: "GigaSMART engine port", Required: true}, "gateway": schema.StringAttribute{MarkdownDescription: "Only valid and required when 'dhcp' is false", Optional: true, Computed: true}, "hw_address": schema.StringAttribute{Optional: true, Computed: true}, "interface": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("eth2", "eth3")}}, "ip_address": schema.StringAttribute{MarkdownDescription: "Only valid and required when 'dhcp' is false", Optional: true, Computed: true}, "ip_mask": schema.StringAttribute{MarkdownDescription: "Only valid and required when 'dhcp' is false", Optional: true, Computed: true}, "mtu": schema.Int64Attribute{MarkdownDescription: "Only valid when 'dhcp' is false", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(68, 1500)}}, "proxy_server_profile": schema.StringAttribute{Optional: true, Computed: true}, "status": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("up", "down")}}, "vlan": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(20, 4094)}}}}, "operation": schema.StringAttribute{MarkdownDescription: "Adds or deletes(None) engine interface connectivity", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("Add", "None")}}}}}, "ssl_path": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{Optional: true, Computed: true}, "flex_inline_map": schema.SingleNestedAttribute{MarkdownDescription: "When 'oobCopy' is configured at least one of 'aToB' or 'bToA' must be configured", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"a_to_b": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"ib_pathway": schema.StringAttribute{MarkdownDescription: "ibPathway alias. Only applicable when type is 'ibPathway'", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "tools": schema.ListAttribute{MarkdownDescription: "ordered list of inline tools or vports. Only applicable when 'type' is 'tools'", Optional: true, Computed: true, ElementType: types.StringType}, "type": schema.StringAttribute{MarkdownDescription: "when set to 'tools', the list of the processing inline tools have to be provided. when set to 'ibPathway', ibPathway alias should be provided ", Required: true, Validators: []validator.String{stringvalidator.OneOf("bypass", "tools", "reverse", "same", "ibPathway")}}}}, "b_to_a": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"ib_pathway": schema.StringAttribute{MarkdownDescription: "ibPathway alias. Only applicable when type is 'ibPathway'", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "tools": schema.ListAttribute{MarkdownDescription: "ordered list of inline tools or vports. Only applicable when 'type' is 'tools'", Optional: true, Computed: true, ElementType: types.StringType}, "type": schema.StringAttribute{MarkdownDescription: "when set to 'tools', the list of the processing inline tools have to be provided. when set to 'ibPathway', ibPathway alias should be provided ", Required: true, Validators: []validator.String{stringvalidator.OneOf("bypass", "tools", "reverse", "same", "ibPathway")}}}}, "oob_copy": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"direction": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("aToB", "bToA", "both")}}, "dst_ports": schema.ListAttribute{MarkdownDescription: "list of destination tool ports", Required: true, ElementType: types.StringType, Validators: []validator.List{listvalidator.SizeAtLeast(1)}}, "src_ports": schema.ListAttribute{MarkdownDescription: "inline network or an item from a-to-b and b-to-a lists", Required: true, ElementType: types.StringType, Validators: []validator.List{listvalidator.SizeAtLeast(1)}}, "tag": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"type": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.OneOf("none", "asInline", "original")}}}}}}}, "svt_mode": schema.BoolAttribute{Optional: true, Computed: true}, "svt_tag": schema.Int64Attribute{MarkdownDescription: "only applicable when svtMode is enabled", Optional: true, Computed: true}, "tag": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"tag_protocol_id": schema.StringAttribute{MarkdownDescription: "When tool VLAN tag is added , this protocol Id will be added which egress out the traffic", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("0x8100", "0x9100", "0x88a8")}}, "type": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.OneOf("auto", "vlan")}}, "vlan_id": schema.Int64Attribute{MarkdownDescription: "only applicable when type is 'vlan'", Optional: true, Computed: true}}}}}}}}, "ssl_profile_alias": schema.StringAttribute{MarkdownDescription: "Alias of the ssl profile auto created by ssl app", Optional: true, Computed: true}, "ssl_profile_config": schema.SingleNestedAttribute{MarkdownDescription: "Flexible Inline SSL Profile", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Alias of the inline SSL profile", Required: true}, "certificate": schema.SingleNestedAttribute{MarkdownDescription: "SSL profile certificate configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"expired": schema.StringAttribute{MarkdownDescription: "SSL profile on expired certificate", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("decrypt", "drop")}}, "invalid": schema.StringAttribute{MarkdownDescription: "SSL profile on invalid certificate", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("decrypt", "drop")}}, "revocation": schema.SingleNestedAttribute{MarkdownDescription: "SSL profile certificate revocation configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"crl": schema.SingleNestedAttribute{MarkdownDescription: "SSL profile certificate revocation configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"defer": schema.Int64Attribute{MarkdownDescription: "timeout in seconds", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(20, 100)}}, "enabled": schema.BoolAttribute{Optional: true, Computed: true}, "fail": schema.StringAttribute{MarkdownDescription: "only applicable when enabled", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("soft", "hard")}}}}, "ocsp": schema.SingleNestedAttribute{MarkdownDescription: "SSL profile certificate revocation configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"defer": schema.Int64Attribute{MarkdownDescription: "timeout in seconds", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(20, 100)}}, "enabled": schema.BoolAttribute{Optional: true, Computed: true}, "fail": schema.StringAttribute{MarkdownDescription: "only applicable when enabled", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("soft", "hard")}}}}}}, "self_signed": schema.StringAttribute{MarkdownDescription: "SSL profile on self-signed certificate", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("decrypt", "drop")}}, "unknown_ca": schema.StringAttribute{MarkdownDescription: "SSL profile on unknown CA certificate", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("decrypt", "drop")}}}}, "client_auth": schema.SingleNestedAttribute{MarkdownDescription: "SSL profile certificate configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"expired": schema.StringAttribute{MarkdownDescription: "SSL profile on expired certificate", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("decrypt", "drop")}}, "invalid": schema.StringAttribute{MarkdownDescription: "SSL profile on invalid certificate", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("decrypt", "drop")}}, "revocation": schema.SingleNestedAttribute{MarkdownDescription: "SSL profile certificate revocation configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"crl": schema.SingleNestedAttribute{MarkdownDescription: "SSL profile certificate revocation configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"defer": schema.Int64Attribute{MarkdownDescription: "timeout in seconds", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(20, 100)}}, "enabled": schema.BoolAttribute{Optional: true, Computed: true}, "fail": schema.StringAttribute{MarkdownDescription: "only applicable when enabled", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("soft", "hard")}}}}, "ocsp": schema.SingleNestedAttribute{MarkdownDescription: "SSL profile certificate revocation configuration", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"defer": schema.Int64Attribute{MarkdownDescription: "timeout in seconds", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(20, 100)}}, "enabled": schema.BoolAttribute{Optional: true, Computed: true}, "fail": schema.StringAttribute{MarkdownDescription: "only applicable when enabled", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("soft", "hard")}}}}}}, "self_signed": schema.StringAttribute{MarkdownDescription: "SSL profile on self-signed certificate", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("decrypt", "drop")}}, "unknown_ca": schema.StringAttribute{MarkdownDescription: "SSL profile on unknown CA certificate", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("decrypt", "drop")}}}}, "cluster_id": schema.StringAttribute{MarkdownDescription: "id of the defining cluster", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "decrypt": schema.SingleNestedAttribute{MarkdownDescription: "SSL profile configuration on decrypt action", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"tcp": schema.SingleNestedAttribute{MarkdownDescription: "SSL decryption TCP control", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"inactive_timeout": schema.Int64Attribute{MarkdownDescription: "SSL decryption TCP inactive timeout (in minutes)", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(2, 1440)}}, "port_map": schema.SingleNestedAttribute{MarkdownDescription: "SSL decryption port map", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"default_out_port": schema.Int64Attribute{MarkdownDescription: "egress port for decryption port map. 0 is disabled.", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 65536)}}, "ports": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"in_port": schema.Int64Attribute{MarkdownDescription: "ingress port for decryption port map", Required: true, Validators: []validator.Int64{int64validator.Between(1, 65536)}}, "out_port": schema.Int64Attribute{MarkdownDescription: "egress port for decryption port map", Required: true, Validators: []validator.Int64{int64validator.Between(1, 65536)}}, "rule_id": schema.Int64Attribute{Optional: true, Computed: true}}}}}}}}, "tool_bypass": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enable": schema.BoolAttribute{Optional: true, Computed: true}}}}}, "default_action": schema.StringAttribute{MarkdownDescription: "Action to take if none of the profile rules match", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("decrypt", "no-decrypt")}}, "high_avail": schema.SingleNestedAttribute{MarkdownDescription: "SSL profile configuration on high availability", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"active_standby": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enable": schema.BoolAttribute{Optional: true, Computed: true}}}}}, "key_map": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"hostname": schema.StringAttribute{MarkdownDescription: "hostname or IP address", Required: true}, "key": schema.StringAttribute{MarkdownDescription: "SSL key alias", Required: true}, "rule_id": schema.Int64Attribute{Optional: true, Computed: true}}}}, "monitor": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("enable", "disable", "inline")}}, "network_group": schema.SingleNestedAttribute{MarkdownDescription: "SSL Profile configuration for multiple entry in network groups", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"multiple_entry": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enable": schema.BoolAttribute{Optional: true, Computed: true}}}}}, "no_decrypt": schema.SingleNestedAttribute{MarkdownDescription: "SSL profile configuration on no-decrypt action", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"tool_bypass": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enable": schema.BoolAttribute{Optional: true, Computed: true}}}}}, "non_ssl_tcp": schema.SingleNestedAttribute{MarkdownDescription: "SSL profile configuration on TCP proxy action", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"tool_bypass": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enable": schema.BoolAttribute{Optional: true, Computed: true}}}}}, "split_proxy": schema.SingleNestedAttribute{MarkdownDescription: "SSL profile configuration for Split Proxy", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"mode": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enable": schema.BoolAttribute{Optional: true, Computed: true}}}, "server_non_pfs_ciphers": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"enable": schema.BoolAttribute{Optional: true, Computed: true}}}}}, "start_tls": schema.SingleNestedAttribute{MarkdownDescription: "SSL profile configuration on start TLS action", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"l4_port": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.Int64Type}}}, "tcp": schema.SingleNestedAttribute{MarkdownDescription: "SSL profile configuration on TCP", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"delayed_ack": schema.BoolAttribute{MarkdownDescription: "enable/disable TCP delayed ACK", Optional: true, Computed: true}, "syn_retries": schema.Int64Attribute{MarkdownDescription: "TCP Sync retries", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 12)}}, "timewait_timeout": schema.Int64Attribute{MarkdownDescription: "TCP Wait Timeout value", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 300)}}}}, "tool": schema.SingleNestedAttribute{MarkdownDescription: "SSL Profile configuration for Tools", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"early_engage": schema.BoolAttribute{MarkdownDescription: "enable/disable tool early engage", Optional: true, Computed: true}, "fail_action": schema.StringAttribute{MarkdownDescription: "Action to take if the tool fails", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("fail-open", "fail-close")}}}}, "url_cache": schema.SingleNestedAttribute{MarkdownDescription: "SSL profile configuration on url-cache", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"miss_action": schema.StringAttribute{MarkdownDescription: "The action to take if local URL category resolution misses", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("decrypt", "no-decrypt", "defer")}}, "timeout": schema.Int64Attribute{MarkdownDescription: "defer timeout in seconds. Only applicable for missAction 'defer'", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 10)}}}}}}, "tag_protocol_id": schema.StringAttribute{MarkdownDescription: "When tool VLAN tag is added , this protocol Id will be added which egress out the traffic", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("0x8100", "0x9100", "0x88a8")}}, "trust_store_configs": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"trust_store_append_configs": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"file": schema.StringAttribute{MarkdownDescription: "The contents of the file. Mutually exclusive with 'fileSource'", Optional: true, Computed: true}, "file_source": schema.SingleNestedAttribute{MarkdownDescription: "Remote file source or destination", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"hostname": schema.StringAttribute{MarkdownDescription: "server address", Required: true}, "password": schema.StringAttribute{MarkdownDescription: "password to use for server login", Optional: true, Computed: true, Sensitive: true}, "path": schema.StringAttribute{MarkdownDescription: "file path on server", Required: true}, "protocol": schema.StringAttribute{MarkdownDescription: "Protocol to access server. HTTP and HTTPS only applicable for retrieving file", Required: true, Validators: []validator.String{stringvalidator.OneOf("scp", "sftp", "ftp", "tftp", "http", "https")}}, "username": schema.StringAttribute{MarkdownDescription: "user name to use for server login", Optional: true, Computed: true}}}}}, "trust_store_replace_configs": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"file": schema.StringAttribute{MarkdownDescription: "The contents of the file. Mutually exclusive with 'fileSource'", Optional: true, Computed: true}, "file_source": schema.SingleNestedAttribute{MarkdownDescription: "Remote file source or destination", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"hostname": schema.StringAttribute{MarkdownDescription: "server address", Required: true}, "password": schema.StringAttribute{MarkdownDescription: "password to use for server login", Optional: true, Computed: true, Sensitive: true}, "path": schema.StringAttribute{MarkdownDescription: "file path on server", Required: true}, "protocol": schema.StringAttribute{MarkdownDescription: "Protocol to access server. HTTP and HTTPS only applicable for retrieving file", Required: true, Validators: []validator.String{stringvalidator.OneOf("scp", "sftp", "ftp", "tftp", "http", "https")}}, "username": schema.StringAttribute{MarkdownDescription: "user name to use for server login", Optional: true, Computed: true}}}}}}}, "vlan_id": schema.Int64Attribute{MarkdownDescription: "Vlan id for proxy maps", Optional: true, Computed: true}, "vport_alias": schema.StringAttribute{MarkdownDescription: "Alias of the vport auto created by ssl app", Optional: true, Computed: true}, "vport_config": schema.SingleNestedAttribute{MarkdownDescription: "GigaSMART vPort", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{Required: true}, "deferred_binding": schema.BoolAttribute{MarkdownDescription: "enable/disable deferred-binding", Optional: true, Computed: true}, "fail_over_action": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("vport-bypass", "vport-drop", "network-bypass", "network-drop", "network-port-forced-down")}}, "gs_group": schema.StringAttribute{MarkdownDescription: "Alias of referenced managing GsGroup", Required: true}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("green", "yellow", "orange", "red")}}, "health_state_reasons": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Optional: true, Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("green", "yellow", "orange", "red")}}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("PORT_LOW_UTIL", "PORT_HIGH_UTIL", "PORT_PACKET_DROP", "PORT_PACKET_ERROR", "GIGASTREAM_IMBALANCE")}}}}}, "inline_status": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("up", "down")}}, "inner_traffic_path": schema.StringAttribute{MarkdownDescription: "Similar to inline-network traffic-path, applicable for inner map", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("to-inline-tool", "bypass", "drop", "monitor")}}, "metadata_monitoring": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"action": schema.StringAttribute{MarkdownDescription: "metadata monitoring action", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("enable", "disable")}}, "exporters": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType}}}, "mode": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("none", "gtp-overlap")}}, "outer_traffic_path": schema.StringAttribute{MarkdownDescription: "Similar to inline-tool/inline-tool-group flex-traffic path, applicable for outer map", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("to-inline-tool", "bypass", "drop", "monitor")}}, "sa_apf_profile": schema.StringAttribute{MarkdownDescription: "ASF session profile", Optional: true, Computed: true}}}, "white_list_config": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"operation": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("ADD", "REPLACE", "DELETE")}}, "profile_list": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"file_source": schema.SingleNestedAttribute{MarkdownDescription: "Remote file source or destination", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"hostname": schema.StringAttribute{MarkdownDescription: "server address", Required: true}, "password": schema.StringAttribute{MarkdownDescription: "password to use for server login", Optional: true, Computed: true, Sensitive: true}, "path": schema.StringAttribute{MarkdownDescription: "file path on server", Required: true}, "protocol": schema.StringAttribute{MarkdownDescription: "Protocol to access server. HTTP and HTTPS only applicable for retrieving file", Required: true, Validators: []validator.String{stringvalidator.OneOf("scp", "sftp", "ftp", "tftp", "http", "https")}}, "username": schema.StringAttribute{MarkdownDescription: "user name to use for server login", Optional: true, Computed: true}}}, "list": schema.StringAttribute{MarkdownDescription: "The nodecryptlist or decryptlist . Mutually exclusive with 'fileSource'", Optional: true, Computed: true}}}}}}}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Inline SSL app only for that cluster is returned", Optional: true}, "cluster_name": schema.StringAttribute{Optional: true, Computed: true}, "config_status": schema.StringAttribute{MarkdownDescription: "Configuration status of this SSL app", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("SUCCESS", "PARTIAL_SUCCESS", "FAILED", "PENDING", "INCOMPLETE_DELETION", "DAMAGED")}}, "config_status_reasons": schema.ListAttribute{MarkdownDescription: "In case of configuration failure, this message provides details about the possible cause of the failure", Optional: true, Computed: true, ElementType: types.StringType}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("green", "yellow", "orange", "red")}}, "health_state_reasons": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Optional: true, Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("green", "yellow", "orange", "red")}}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("PORT_LOW_UTIL", "PORT_HIGH_UTIL", "PORT_PACKET_DROP", "PORT_PACKET_ERROR", "GIGASTREAM_IMBALANCE")}}}}}, "m_tls": schema.StringAttribute{MarkdownDescription: "Enable mTLS to do client Authentication", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("false", "true")}}, "ria_configs": schema.ListNestedAttribute{MarkdownDescription: "Configs applicable if RIA is enabled in ssl app, the listed properties here will not be applicable in appIntent configs if RIA is enabled", Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"cluster_name": schema.StringAttribute{Optional: true, Computed: true}, "gs_engines": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType}, "gs_group_alias": schema.StringAttribute{MarkdownDescription: "Alias of the gs group auto created by ssl app", Optional: true, Computed: true}, "gsop_alias": schema.StringAttribute{MarkdownDescription: "Alias of the gsop auto created by ssl app", Optional: true, Computed: true}, "network_access_configs": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"network_access": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "id of the defining cluster", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "dhcp": schema.BoolAttribute{Optional: true, Computed: true}, "dns": schema.StringAttribute{MarkdownDescription: "Only valid and required when 'dhcp' is false", Optional: true, Computed: true}, "eport": schema.StringAttribute{MarkdownDescription: "GigaSMART engine port", Required: true}, "gateway": schema.StringAttribute{MarkdownDescription: "Only valid and required when 'dhcp' is false", Optional: true, Computed: true}, "hw_address": schema.StringAttribute{Optional: true, Computed: true}, "interface": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("eth2", "eth3")}}, "ip_address": schema.StringAttribute{MarkdownDescription: "Only valid and required when 'dhcp' is false", Optional: true, Computed: true}, "ip_mask": schema.StringAttribute{MarkdownDescription: "Only valid and required when 'dhcp' is false", Optional: true, Computed: true}, "mtu": schema.Int64Attribute{MarkdownDescription: "Only valid when 'dhcp' is false", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(68, 1500)}}, "proxy_server_profile": schema.StringAttribute{Optional: true, Computed: true}, "status": schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("up", "down")}}, "vlan": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(20, 4094)}}}}, "operation": schema.StringAttribute{MarkdownDescription: "Adds or deletes(None) engine interface connectivity", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("Add", "None")}}}}}, "ssl_profile_alias": schema.StringAttribute{MarkdownDescription: "Alias of the ssl profile auto created by ssl app", Optional: true, Computed: true}, "vport_alias": schema.StringAttribute{MarkdownDescription: "Alias of the vport auto created by ssl app", Optional: true, Computed: true}}}}, "ria_enabled": schema.StringAttribute{MarkdownDescription: "Enable RIA to configure ssl app in two nodes which is needed for RIA SSL", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("false", "true")}}}, Blocks: map[string]schema.Block{"timeouts": schema.SingleNestedBlock{Attributes: map[string]schema.Attribute{"create": schema.Int64Attribute{MarkdownDescription: "Maximum time in seconds for the create operation. Defaults to 1200.", Optional: true}, "read": schema.Int64Attribute{MarkdownDescription: "Maximum time in seconds for the read operation. Defaults to 1200.", Optional: true}, "update": schema.Int64Attribute{MarkdownDescription: "Maximum time in seconds for the update operation. Defaults to 1200.", Optional: true}, "delete": schema.Int64Attribute{MarkdownDescription: "Maximum time in seconds for the delete operation. Defaults to 1200.", Optional: true}}, MarkdownDescription: "Per-operation timeouts in seconds."}}}
}

// IdentitySchema returns the resource identity schema shared with the paired list resource.
func (r *InlineSslAppResource) IdentitySchema(_ context.Context, _ resource.IdentitySchemaRequest, resp *resource.IdentitySchemaResponse) {
	resp.IdentitySchema = identityschema.Schema{Attributes: map[string]identityschema.Attribute{"alias": identityschema.StringAttribute{RequiredForImport: true, Description: "Alias of the inline ssl app"}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *InlineSslAppResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InlineSslAppResourceModel
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
func (r *InlineSslAppResource) createRemote(ctx context.Context, plan *InlineSslAppResourceModel, resp *resource.CreateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_inline_ssl_app", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_inline_ssl_app", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/flexInline/inlineSslApp"
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_inline_ssl_app", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_inline_ssl_app", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error creating gigavuecore_inline_ssl_app", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error creating gigavuecore_inline_ssl_app", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error creating gigavuecore_inline_ssl_app", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error creating gigavuecore_inline_ssl_app", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error creating gigavuecore_inline_ssl_app", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error creating gigavuecore_inline_ssl_app", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error creating gigavuecore_inline_ssl_app", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error creating gigavuecore_inline_ssl_app", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error creating gigavuecore_inline_ssl_app", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error creating gigavuecore_inline_ssl_app", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_inline_ssl_app", fmt.Sprintf("Could not map response to state: %s", err))
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
			resp.Diagnostics.AddError("Error creating gigavuecore_inline_ssl_app", "The create response did not contain an identifier and no Location header was returned, so the resource cannot be tracked in state.")
			return
		}
	}
}

// Read refreshes the Terraform state with the latest remote values.
func (r *InlineSslAppResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InlineSslAppResourceModel
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
func (r *InlineSslAppResource) readRemote(ctx context.Context, state *InlineSslAppResourceModel, resp *resource.ReadResponse) (removed bool) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/flexInline/inlineSslApp/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(state.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_inline_ssl_app", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !state.ClusterId.IsNull() {
		query.Set("clusterId", state.ClusterId.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_inline_ssl_app", fmt.Sprintf("Could not send request: %s", err))
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
			resp.Diagnostics.AddError("Error reading gigavuecore_inline_ssl_app", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_inline_ssl_app", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_inline_ssl_app", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_inline_ssl_app", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_inline_ssl_app", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_inline_ssl_app", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_inline_ssl_app", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_inline_ssl_app", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_inline_ssl_app", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_inline_ssl_app", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["gigaFlexInlineSslApp"]; ok {
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
		resp.Diagnostics.AddError("Error reading gigavuecore_inline_ssl_app", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
	return
}

// Update modifies the remote resource to match the desired plan.
func (r *InlineSslAppResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan InlineSslAppResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state InlineSslAppResourceModel
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
func (r *InlineSslAppResource) updateRemote(ctx context.Context, plan *InlineSslAppResourceModel, resp *resource.UpdateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_inline_ssl_app", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_inline_ssl_app", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/flexInline/inlineSslApp/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(plan.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_inline_ssl_app", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !plan.ClusterId.IsNull() {
		query.Set("clusterId", plan.ClusterId.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_inline_ssl_app", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error updating gigavuecore_inline_ssl_app", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error updating gigavuecore_inline_ssl_app", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error updating gigavuecore_inline_ssl_app", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error updating gigavuecore_inline_ssl_app", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error updating gigavuecore_inline_ssl_app", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error updating gigavuecore_inline_ssl_app", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error updating gigavuecore_inline_ssl_app", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error updating gigavuecore_inline_ssl_app", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error updating gigavuecore_inline_ssl_app", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error updating gigavuecore_inline_ssl_app", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_inline_ssl_app", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Delete destroys the remote resource.
func (r *InlineSslAppResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InlineSslAppResourceModel
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
func (r *InlineSslAppResource) deleteRemote(ctx context.Context, state *InlineSslAppResourceModel, resp *resource.DeleteResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/flexInline/inlineSslApp/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(state.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_inline_ssl_app", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !state.ClusterId.IsNull() {
		query.Set("clusterId", state.ClusterId.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_inline_ssl_app", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		return
	}
	if !(httpResp.StatusCode == 204) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error deleting gigavuecore_inline_ssl_app", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error deleting gigavuecore_inline_ssl_app", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error deleting gigavuecore_inline_ssl_app", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error deleting gigavuecore_inline_ssl_app", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error deleting gigavuecore_inline_ssl_app", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error deleting gigavuecore_inline_ssl_app", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error deleting gigavuecore_inline_ssl_app", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error deleting gigavuecore_inline_ssl_app", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error deleting gigavuecore_inline_ssl_app", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *InlineSslAppResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *InlineSslAppResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("alias"), req.ID)...)
}
