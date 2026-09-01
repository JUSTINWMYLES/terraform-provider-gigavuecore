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
	_ resource.Resource                = (*RecordResource)(nil)
	_ resource.ResourceWithIdentity    = (*RecordResource)(nil)
	_ resource.ResourceWithImportState = (*RecordResource)(nil)
	_ resource.ResourceWithConfigure   = (*RecordResource)(nil)
)

// RecordResource is the generated Terraform managed resource implementation.
type RecordResource struct {
	client *client.Client
}

// RecordResourceModel describes the Terraform state and plan shape for RecordResource.
type RecordResourceModel struct {
	Alias          types.String   `tfsdk:"alias"`
	ClusterId      types.String   `tfsdk:"cluster_id" json:"clusterId"`
	Collect        types.Object   `tfsdk:"collect"`
	Collects       types.Dynamic  `tfsdk:"collects"`
	Description    types.String   `tfsdk:"description"`
	ExportBlankPen types.Bool     `tfsdk:"export_blank_pen" json:"exportBlankPen"`
	Exporters      types.List     `tfsdk:"exporters"`
	Match          types.Object   `tfsdk:"match"`
	NfVersion      types.String   `tfsdk:"nf_version" json:"nfVersion"`
	SamplingRate   types.Int64    `tfsdk:"sampling_rate" json:"samplingRate"`
	Timeouts       timeouts.Value `tfsdk:"timeouts"`
}

// Metadata returns the resource type name.
func (r *RecordResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_record"
}

// Schema returns the Terraform schema for this resource.
func (r *RecordResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Create a new Netflow Template", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{Required: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "id of the defining cluster", Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "collect": schema.SingleNestedAttribute{MarkdownDescription: "Netflow Record Collect spec", Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"collect_fields": schema.SetAttribute{Required: true, ElementType: types.StringType, Validators: []validator.Set{setvalidator.ValueStringsAre(stringvalidator.OneOf("counterBytes", "counterBytesLong", "counterPackets", "counterPacketsLong", "flowEndReason", "tsSysUptimeFirst", "tsSysUptimeLast", "tsFlowEndMseconds", "tsFlowEndSeconds", "tsFlowEndUseconds", "tsFlowStartMseconds", "tsFlowStartSeconds", "tsFlowStartUseconds", "phyIntfIn", "phyIntfOut", "intfNeighbor", "vlan", "macSrc", "macDst", "ipv4Version", "ipv4HeaderLen", "ipv4Tos", "ipv4TotalLen", "ipv4TotalLenMin", "ipv4TotalLenMax", "ipv4FragId", "ipv4FragFlags", "ipv4FragOffset", "ipv4Ttl", "ipv4Protocol", "ipv4Options", "ipv4Dscp", "ipv4Precedence", "ipv4AddrSrc", "ipv4AddrDst", "ipv4SectionHeaderSize", "ipv4SectionPayloadSize", "ipv6Version", "ipv6TrafficClass", "ipv6FlowLabel", "ipv6NextHeader", "ipv6HopLimit", "ipv6HopLimitMin", "ipv6HopLimitMax", "ipv6HeaderLen", "ipv6PayloadLen", "ipv6TotalLen", "ipv6TotalLenMin", "ipv6TotalLenMax", "ipv6ExtensionMap", "ipv6FragId", "ipv6FragFlags", "ipv6FragOffset", "ipv6Protocol", "ipv6Dscp", "ipv6Precedence", "ipv6AddrSrc", "ipv6AddrDst", "ipv6SectionHeaderSize", "ipv6SectionPayloadSize", "portSrc", "portDst", "icmpIpv4Code", "icmpIpv4Type", "icmpIpv6Code", "icmpIpv6Type", "udpPortSrc", "udpPortDst", "udpPayloadLen", "tcpPortSrc", "tcpPortDst", "tcpSeqNumber", "tcpAckNumber", "tcpWindow", "tcpUrgentPtr", "tcpHeaderLen", "tcpFlags", "penGigamonUrl", "penGigamonUserAgent", "penHttpResponseCode", "penHttpMethod", "penHttpVersion", "penHttpHost", "penDnsIdentifier", "penDnsOpCode", "penDnsResponseCode", "penDnsQueryName", "penDnsResponseName", "penDnsResponseTtl", "penDnsResponseIpv4Addr", "penDnsResponseIpv4AddrText", "penDnsResponseIpv6Addr", "penDnsResponseIpv6AddrText", "penDnsResponseClass", "penDnsResponseClassText", "penDnsResponseRdLength", "penDnsResponseRdata", "penDnsAuthorityName", "penDnsAuthorityType", "penDnsAuthorityTypeText", "penDnsAuthorityClass", "penDnsAuthorityClassText", "penDnsAuthorityTtl", "penDnsAuthorityRdLength", "penDnsAuthorityRdata", "penDnsAdditionalName", "penDnsAdditionalType", "penDnsAdditionalTypeText", "penDnsAdditionalClass", "penDnsAdditionalClassText", "penDnsAdditionalTtl", "penDnsAdditionalRdLength", "penDnsAdditionalRdata", "penDnsAnCount", "penDnsArCount", "penDnsBits", "penDnsNsCount", "penDnsQdCount", "penDnsQueryType", "penDnsQueryTypeText", "penDnsQueryClass", "penDnsQueryClassText", "penDnsResponseType", "penDnsResponseTypeText", "penSslCertificateIssuerCommonName", "penSslCertificateSubjectCommonName", "penSslCertificateIssuer", "penSslCertificateSubject", "penSslCertificateValidNotBefore", "penSslCertificateValidNotBeforeText", "penSslCertificateValidNotAfter", "penSslCertificateValidNotAfterText", "penSslCertificateSerialNumber", "penSslCertificateSerialNumberText", "penSslCertificateSignatureAlogrithm", "penSslCertificateSignatureAlgorithmText", "penSslCertificateSubjectAlgorithm", "penSslCertificateSubjectAlgorithmText", "penSslCertificateSubjectKeySize", "penSslCertificateSubjectAltName", "penSslServerNameIndication", "penSslServerVersion", "penSslServerVersionText", "penSslServerCipher", "penSslServerCipherText", "penSslServerCompressionMethod", "penSslServerSessionId", "exporterIpv4Addr", "exporterIpv6Addr"))}}, "intf_neighbor_width": schema.Int64Attribute{MarkdownDescription: "Only valid when 'intfNeighbor' is selected. Used for representing input interface name with width", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 32)}}, "ipv4_addr_dst_mask_len": schema.Int64Attribute{MarkdownDescription: "Only valid when 'ipv4AddrDst' is selected. defaults to 32, in which case it matches full address", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 32)}}, "ipv4_addr_src_mask_len": schema.Int64Attribute{MarkdownDescription: "Only valid when 'ipv4AddrSrc' is selected. Defaults to 32, in which case it matches full address", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 32)}}, "ipv4_section_header_size": schema.Int64Attribute{MarkdownDescription: "Number of bytes of raw data starting at the IPv4 header, to use as a key field. Only valid and required when ipv4SectionHeaderSize' is selected", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 128)}}, "ipv4_section_payload_size": schema.Int64Attribute{MarkdownDescription: "Number of bytes of raw data starting at the IPv4 payload, to use as a key field. Only valid and required when 'ipv4SectionPayloadSize' is selected", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 128)}}, "ipv6_addr_dst_mask_len": schema.Int64Attribute{MarkdownDescription: "Only valid when 'ipv6AddrDst' is selected. defaults to 128, in which case it matches full address", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 128)}}, "ipv6_addr_src_mask_len": schema.Int64Attribute{MarkdownDescription: "Only valid when 'ipv6AddrSrc' is selected. Defaults to 128, in which case it matches full address", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 128)}}, "ipv6_section_header_size": schema.Int64Attribute{MarkdownDescription: "Only valid and required when 'ipv6SectionHeaderSize' is selected", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 128)}}, "ipv6_section_payload_size": schema.Int64Attribute{MarkdownDescription: "Only valid and required when 'ipv6SectionPayloadSize' is selected", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 128)}}, "pen_gigamon_url_size": schema.Int64Attribute{MarkdownDescription: "Only valid when 'penGigamonUrl' is selected", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 250)}}, "pen_gigamon_user_agent_width": schema.Int64Attribute{MarkdownDescription: "Only valid when 'penGigamonUserAgent' is selected", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 250)}}, "pen_http_host_width": schema.Int64Attribute{MarkdownDescription: "Only valid when 'penHttpHost' is selected", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 250)}}, "pen_ssl_certificate_issuer_common_name_width": schema.Int64Attribute{MarkdownDescription: "Only valid when 'penSslCertificateIssuerCommonName' is selected", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 64)}}, "pen_ssl_certificate_issuer_width": schema.Int64Attribute{MarkdownDescription: "Only valid when 'penSslCertificateIssuer' is selected", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 250)}}, "pen_ssl_certificate_subject_alt_name_width": schema.Int64Attribute{MarkdownDescription: "Only valid when 'penSslCertificateSubjectAltName' is selected", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 64)}}, "pen_ssl_certificate_subject_common_name_width": schema.Int64Attribute{MarkdownDescription: "Only valid when 'penSslCertificateSubjectCommonName' is selected", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 64)}}, "pen_ssl_certificate_subject_width": schema.Int64Attribute{MarkdownDescription: "Only valid when 'penSslCertificateSubject' is selected", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 250)}}, "pen_ssl_server_name_indication_width": schema.Int64Attribute{MarkdownDescription: "Only valid when 'penSslServerNameIndication' is selected", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 64)}}, "phy_intf_in_width": schema.StringAttribute{MarkdownDescription: "Only valid when 'phyIntfIn' is selected", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("2", "4")}}, "phy_intf_out_width": schema.StringAttribute{MarkdownDescription: "Only valid when 'phyIntfOut' is selected", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("2", "4")}}, "tcp_flags": schema.SetAttribute{MarkdownDescription: "Represents a bit mask. Only valid and required when 'tcpFlags' is selected", Optional: true, Computed: true, ElementType: types.StringType, Validators: []validator.Set{setvalidator.ValueStringsAre(stringvalidator.OneOf("cwr", "ece", "urg", "ack", "psh", "rst", "syn", "fin"))}}}}, "collects": schema.DynamicAttribute{Optional: true, Computed: true}, "description": schema.StringAttribute{Optional: true, Computed: true}, "export_blank_pen": schema.BoolAttribute{MarkdownDescription: "if true and there is a mix of private elements and non-private elements and the private elements are blank, export the record", Optional: true, Computed: true}, "exporters": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType}, "match": schema.SingleNestedAttribute{MarkdownDescription: "Netflow Record Match spec", Required: true, Attributes: map[string]schema.Attribute{"description": schema.StringAttribute{Optional: true, Computed: true}, "ipv4_addr_dst_mask_len": schema.Int64Attribute{MarkdownDescription: "Only valid when 'ipv4AddrDst' is selected. defaults to 32, in which case it matches full address", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 32)}}, "ipv4_addr_src_mask_len": schema.Int64Attribute{MarkdownDescription: "Only valid when 'ipv4AddrSrc' is selected. Defaults to 32, in which case it matches full address", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 32)}}, "ipv4_section_header_size": schema.Int64Attribute{MarkdownDescription: "Number of bytes of raw data starting at the IPv4 header, to use as a key field. Only valid and required when ipv4SectionHeaderSize' is selected", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 128)}}, "ipv4_section_payload_size": schema.Int64Attribute{MarkdownDescription: "Number of bytes of raw data starting at the IPv4 payload, to use as a key field. Only valid and required when 'ipv4SectionPayloadSize' is selected", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 128)}}, "ipv6_addr_dst_mask_len": schema.Int64Attribute{MarkdownDescription: "Only valid when 'ipv6AddrDst' is selected. defaults to 128, in which case it matches full address", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 128)}}, "ipv6_addr_src_mask_len": schema.Int64Attribute{MarkdownDescription: "Only valid when 'ipv6AddrSrc' is selected. Defaults to 128, in which case it matches full address", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 128)}}, "ipv6_section_header_size": schema.Int64Attribute{MarkdownDescription: "Only valid and required when 'ipv6SectionHeaderSize' is selected", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 128)}}, "ipv6_section_payload_size": schema.Int64Attribute{MarkdownDescription: "Only valid and required when 'ipv6SectionPayloadSize' is selected", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 128)}}, "match_fields": schema.SetAttribute{Required: true, ElementType: types.StringType, Validators: []validator.Set{setvalidator.ValueStringsAre(stringvalidator.OneOf("phyIntfIn", "vlan", "macSrc", "macDst", "ipv4Version", "ipv4HeaderLen", "ipv4Tos", "ipv4TotalLen", "ipv4FragId", "ipv4FragFlags", "ipv4FragOffset", "ipv4Ttl", "ipv4Protocol", "ipv4Options", "ipv4Dscp", "ipv4Precedence", "ipv4AddrSrc", "ipv4AddrDst", "ipv4SectionHeaderSize", "ipv4SectionPayloadSize", "ipv6Version", "ipv6TrafficClass", "ipv6FlowLabel", "ipv6NextHeader", "ipv6HopLimit", "ipv6HeaderLen", "ipv6PayloadLen", "ipv6TotalLen", "ipv6TotalLenMin", "ipv6TotalLenMax", "ipv6ExtensionMap", "ipv6FragId", "ipv6FragFlags", "ipv6FragOffset", "ipv6Protocol", "ipv6Dscp", "ipv6Precedence", "ipv6AddrSrc", "ipv6AddrDst", "ipv6SectionHeaderSize", "ipv6SectionPayloadSize", "portSrc", "portDst", "icmpIpv4Code", "icmpIpv4Type", "icmpIpv6Code", "icmpIpv6Type", "udpPortSrc", "udpPortDst", "udpPayloadLen", "tcpPortSrc", "tcpPortDst", "tcpSeqNumber", "tcpAckNumber", "tcpWindow", "tcpUrgentPtr", "tcpHeaderLen", "tcpFlags"))}}, "phy_intf_in_width": schema.StringAttribute{MarkdownDescription: "Only valid when 'phyIntfIn' is selected", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("2", "4")}}, "tcp_flags": schema.SetAttribute{MarkdownDescription: "Represents a bit mask. Only valid and required when 'tcpFlags' is selected", Optional: true, Computed: true, ElementType: types.StringType, Validators: []validator.Set{setvalidator.ValueStringsAre(stringvalidator.OneOf("cwr", "ece", "urg", "ack", "psh", "rst", "syn", "fin"))}}}}, "nf_version": schema.StringAttribute{MarkdownDescription: "v5 is readOnly", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("v5", "v9", "ipfix")}}, "sampling_rate": schema.Int64Attribute{MarkdownDescription: "Packet interval window size. Valid values: 1-16000 (in packets); Associated monitor must 'multi-rate' sampling mode set; 0 is disabled", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 16000)}}}, Blocks: map[string]schema.Block{"timeouts": timeouts.Block(ctx, timeouts.Opts{Create: true, Read: true, Update: true, Delete: true})}}
}

// IdentitySchema returns the resource identity schema shared with the paired list resource.
func (r *RecordResource) IdentitySchema(_ context.Context, _ resource.IdentitySchemaRequest, resp *resource.IdentitySchemaResponse) {
	resp.IdentitySchema = identityschema.Schema{Attributes: map[string]identityschema.Attribute{"alias": identityschema.StringAttribute{RequiredForImport: true}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *RecordResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RecordResourceModel
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
	resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("alias"), plan.Alias)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// createRemote performs the create HTTP exchange and decodes the response into plan. Extracted from Create so the request/response logic is unit-testable without a tfsdk.Plan.
func (r *RecordResource) createRemote(ctx context.Context, plan *RecordResourceModel, resp *resource.CreateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_record", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_record", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/apps/netflow/records"
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_record", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", plan.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_record", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error creating gigavuecore_record", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error creating gigavuecore_record", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error creating gigavuecore_record", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error creating gigavuecore_record", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error creating gigavuecore_record", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error creating gigavuecore_record", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error creating gigavuecore_record", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error creating gigavuecore_record", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error creating gigavuecore_record", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error creating gigavuecore_record", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_record", fmt.Sprintf("Could not map response to state: %s", err))
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
			resp.Diagnostics.AddError("Error creating gigavuecore_record", "The create response did not contain an identifier and no Location header was returned, so the resource cannot be tracked in state.")
			return
		}
	}
}

// Read refreshes the Terraform state with the latest remote values.
func (r *RecordResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RecordResourceModel
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
	resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("alias"), state.Alias)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// readRemote performs the read HTTP exchange and decodes the response into state, returning removed=true when the API reports 404. Extracted from Read so the request/response logic is unit-testable without a tfsdk.State.
func (r *RecordResource) readRemote(ctx context.Context, state *RecordResourceModel, resp *resource.ReadResponse) (removed bool) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/netflow/records/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(state.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_record", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", state.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_record", fmt.Sprintf("Could not send request: %s", err))
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
			resp.Diagnostics.AddError("Error reading gigavuecore_record", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_record", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_record", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_record", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_record", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_record", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_record", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_record", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_record", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_record", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["nfRecord"]; ok {
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
		resp.Diagnostics.AddError("Error reading gigavuecore_record", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
	return
}

// Update modifies the remote resource to match the desired plan.
func (r *RecordResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan RecordResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state RecordResourceModel
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
	resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("alias"), plan.Alias)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// updateRemote performs the update HTTP exchange and decodes the response into plan. Extracted from Update so the request/response logic is unit-testable without a tfsdk.Plan.
func (r *RecordResource) updateRemote(ctx context.Context, plan *RecordResourceModel, resp *resource.UpdateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_record", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_record", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/apps/netflow/records/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(plan.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_record", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", plan.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_record", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error updating gigavuecore_record", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error updating gigavuecore_record", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error updating gigavuecore_record", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error updating gigavuecore_record", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error updating gigavuecore_record", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error updating gigavuecore_record", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error updating gigavuecore_record", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error updating gigavuecore_record", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error updating gigavuecore_record", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error updating gigavuecore_record", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_record", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Delete destroys the remote resource.
func (r *RecordResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RecordResourceModel
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
func (r *RecordResource) deleteRemote(ctx context.Context, state *RecordResourceModel, resp *resource.DeleteResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/netflow/records/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(state.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_record", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", state.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_record", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		return
	}
	if !(httpResp.StatusCode == 204) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error deleting gigavuecore_record", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error deleting gigavuecore_record", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error deleting gigavuecore_record", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error deleting gigavuecore_record", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error deleting gigavuecore_record", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error deleting gigavuecore_record", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error deleting gigavuecore_record", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error deleting gigavuecore_record", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error deleting gigavuecore_record", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *RecordResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *RecordResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	recordImportIDParts := strings.Split(req.ID, "/")
	if len(recordImportIDParts) != 2 {
		resp.Diagnostics.AddError("Unexpected Import Identifier", fmt.Sprintf("Expected import identifier with format \"{alias}/{cluster_id}\". Got %q.", req.ID))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("alias"), recordImportIDParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("cluster_id"), recordImportIDParts[1])...)
}
