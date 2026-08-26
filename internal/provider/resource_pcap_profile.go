package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertions.
var (
	_ resource.Resource              = (*PcapProfileResource)(nil)
	_ resource.ResourceWithConfigure = (*PcapProfileResource)(nil)
)

// PcapProfileResource is the generated Terraform managed resource implementation.
type PcapProfileResource struct {
	client *client.Client
}

// PcapProfileResourceModel describes the Terraform state and plan shape for PcapProfileResource.
type PcapProfileResourceModel struct {
	Alias         types.String `tfsdk:"alias"`
	AliasList     types.String `tfsdk:"alias_list" json:"aliasList"`
	ChannelPort   types.String `tfsdk:"channel_port" json:"channelPort"`
	ClusterId     types.String `tfsdk:"cluster_id" json:"clusterId"`
	Direction     types.String `tfsdk:"direction"`
	Id            types.String `tfsdk:"id"`
	PacketLimit   types.Int64  `tfsdk:"packet_limit" json:"packetLimit"`
	PcapConfigs   types.List   `tfsdk:"pcap_configs" json:"pcapConfigs"`
	PcapRulesList types.List   `tfsdk:"pcap_rules_list" json:"pcapRulesList"`
	Port          types.String `tfsdk:"port"`
	PortIds       types.String `tfsdk:"port_ids" json:"portIds"`
}

// Metadata returns the resource type name.
func (r *PcapProfileResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_pcap_profile"
}

// Schema returns the Terraform schema for this resource.
func (r *PcapProfileResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "get All PCAP  profile", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{Computed: true}, "alias_list": schema.StringAttribute{MarkdownDescription: "List of alias to delete. Accepts multiple values, Comma separated eg: aliasList=alias1,alias2", Required: true}, "channel_port": schema.StringAttribute{MarkdownDescription: "port for capturing on egress direction", Computed: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "direction": schema.StringAttribute{MarkdownDescription: "tx, rx, both", Computed: true}, "id": schema.StringAttribute{Computed: true}, "packet_limit": schema.Int64Attribute{MarkdownDescription: "packet limit on a port for capturing", Computed: true}, "pcap_configs": schema.ListNestedAttribute{Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{Required: true}, "channel_port": schema.StringAttribute{MarkdownDescription: "port for capturing on egress direction", Optional: true}, "direction": schema.StringAttribute{MarkdownDescription: "tx, rx, both", Optional: true}, "packet_limit": schema.Int64Attribute{MarkdownDescription: "packet limit on a port for capturing", Optional: true}, "pcap_rules_list": schema.ListNestedAttribute{Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"dscp": schema.StringAttribute{MarkdownDescription: "DiffServ Code Point bits", Optional: true}, "dst_mac": schema.SingleNestedAttribute{MarkdownDescription: "MAC address and mask", Optional: true, Attributes: map[string]schema.Attribute{"address": schema.StringAttribute{MarkdownDescription: "Specifies a single MAC address to match", Optional: true}, "mask": schema.StringAttribute{MarkdownDescription: "Specifies a mask of MAC address", Optional: true}}}, "dstipv4_addrandmask": schema.SingleNestedAttribute{MarkdownDescription: "Ipv4 address and mask. Private class", Optional: true, Attributes: map[string]schema.Attribute{"address": schema.StringAttribute{MarkdownDescription: "Ipv4 address ex: 10.1.1.1", Optional: true}, "mask": schema.StringAttribute{MarkdownDescription: "Ipv4 mask ex:255.255.255.248, or <mask length>, e.g. /29", Optional: true}}}, "ether_type": schema.StringAttribute{MarkdownDescription: "Ether Type", Optional: true}, "inner_vlan": schema.Int64Attribute{MarkdownDescription: "Configure inner-vlan id. valid value is between 1 to 4094", Optional: true}, "ip4_frag": schema.StringAttribute{MarkdownDescription: "IP fragmentation bits, with in range [0..255]", Optional: true}, "ip4_ttl": schema.Int64Attribute{MarkdownDescription: "time to live value, with in range [0..255]", Optional: true}, "ip_ver": schema.StringAttribute{MarkdownDescription: "IP version number", Optional: true}, "packet_hit_count": schema.Int64Attribute{MarkdownDescription: "count of packets matches the configured rule. Applicable for GET", Optional: true}, "portdst": schema.Int64Attribute{MarkdownDescription: "destination port number.Valid value is between 0 to 65535", Optional: true}, "portsrc": schema.Int64Attribute{MarkdownDescription: "source port number.Valid value is between 0 to 65535", Optional: true}, "protocol": schema.Int64Attribute{MarkdownDescription: "protocol number in a range [0..255]. Well-known protocols numbers are: 0-ipv6Hop, 1-icmpIpv4, 2-igmp, 4-ipv4, 6-tcp, 17-udp, 41-ipv6, 46-rsvp, 47-gre, 58-icmpIpv6", Optional: true}, "rule_id": schema.Int64Attribute{MarkdownDescription: "Rule-id to track the filter rules", Optional: true}, "src_mac": schema.SingleNestedAttribute{MarkdownDescription: "MAC address and mask", Optional: true, Attributes: map[string]schema.Attribute{"address": schema.StringAttribute{MarkdownDescription: "Specifies a single MAC address to match", Optional: true}, "mask": schema.StringAttribute{MarkdownDescription: "Specifies a mask of MAC address", Optional: true}}}, "srcipv4_addrandmask": schema.SingleNestedAttribute{MarkdownDescription: "Ipv4 address and mask. Private class", Optional: true, Attributes: map[string]schema.Attribute{"address": schema.StringAttribute{MarkdownDescription: "Ipv4 address ex: 10.1.1.1", Optional: true}, "mask": schema.StringAttribute{MarkdownDescription: "Ipv4 mask ex:255.255.255.248, or <mask length>, e.g. /29", Optional: true}}}, "tcpctl": schema.Int64Attribute{MarkdownDescription: "Configure TCP control bits: URG, SYN, ACK, etc", Optional: true}, "vlan": schema.Int64Attribute{MarkdownDescription: "Configure vlan id. valid value is between 1 to 4094", Optional: true}}}}, "port": schema.StringAttribute{MarkdownDescription: "portID", Required: true}}}}, "pcap_rules_list": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"dscp": schema.StringAttribute{MarkdownDescription: "DiffServ Code Point bits", Computed: true}, "dst_mac": schema.SingleNestedAttribute{MarkdownDescription: "MAC address and mask", Computed: true, Attributes: map[string]schema.Attribute{"address": schema.StringAttribute{MarkdownDescription: "Specifies a single MAC address to match", Computed: true}, "mask": schema.StringAttribute{MarkdownDescription: "Specifies a mask of MAC address", Computed: true}}}, "dstipv4_addrandmask": schema.SingleNestedAttribute{MarkdownDescription: "Ipv4 address and mask. Private class", Computed: true, Attributes: map[string]schema.Attribute{"address": schema.StringAttribute{MarkdownDescription: "Ipv4 address ex: 10.1.1.1", Computed: true}, "mask": schema.StringAttribute{MarkdownDescription: "Ipv4 mask ex:255.255.255.248, or <mask length>, e.g. /29", Computed: true}}}, "ether_type": schema.StringAttribute{MarkdownDescription: "Ether Type", Computed: true}, "inner_vlan": schema.Int64Attribute{MarkdownDescription: "Configure inner-vlan id. valid value is between 1 to 4094", Computed: true}, "ip4_frag": schema.StringAttribute{MarkdownDescription: "IP fragmentation bits, with in range [0..255]", Computed: true}, "ip4_ttl": schema.Int64Attribute{MarkdownDescription: "time to live value, with in range [0..255]", Computed: true}, "ip_ver": schema.StringAttribute{MarkdownDescription: "IP version number", Computed: true}, "packet_hit_count": schema.Int64Attribute{MarkdownDescription: "count of packets matches the configured rule. Applicable for GET", Computed: true}, "portdst": schema.Int64Attribute{MarkdownDescription: "destination port number.Valid value is between 0 to 65535", Computed: true}, "portsrc": schema.Int64Attribute{MarkdownDescription: "source port number.Valid value is between 0 to 65535", Computed: true}, "protocol": schema.Int64Attribute{MarkdownDescription: "protocol number in a range [0..255]. Well-known protocols numbers are: 0-ipv6Hop, 1-icmpIpv4, 2-igmp, 4-ipv4, 6-tcp, 17-udp, 41-ipv6, 46-rsvp, 47-gre, 58-icmpIpv6", Computed: true}, "rule_id": schema.Int64Attribute{MarkdownDescription: "Rule-id to track the filter rules", Computed: true}, "src_mac": schema.SingleNestedAttribute{MarkdownDescription: "MAC address and mask", Computed: true, Attributes: map[string]schema.Attribute{"address": schema.StringAttribute{MarkdownDescription: "Specifies a single MAC address to match", Computed: true}, "mask": schema.StringAttribute{MarkdownDescription: "Specifies a mask of MAC address", Computed: true}}}, "srcipv4_addrandmask": schema.SingleNestedAttribute{MarkdownDescription: "Ipv4 address and mask. Private class", Computed: true, Attributes: map[string]schema.Attribute{"address": schema.StringAttribute{MarkdownDescription: "Ipv4 address ex: 10.1.1.1", Computed: true}, "mask": schema.StringAttribute{MarkdownDescription: "Ipv4 mask ex:255.255.255.248, or <mask length>, e.g. /29", Computed: true}}}, "tcpctl": schema.Int64Attribute{MarkdownDescription: "Configure TCP control bits: URG, SYN, ACK, etc", Computed: true}, "vlan": schema.Int64Attribute{MarkdownDescription: "Configure vlan id. valid value is between 1 to 4094", Computed: true}}}}, "port": schema.StringAttribute{MarkdownDescription: "portID", Computed: true}, "port_ids": schema.StringAttribute{MarkdownDescription: "List of ports to filterBy. Accepts multiple values, Comma separated eg: portIds=1/1/x1,1/1/x2", Required: true}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *PcapProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan PcapProfileResourceModel
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
func (r *PcapProfileResource) createRemote(ctx context.Context, plan *PcapProfileResourceModel, resp *resource.CreateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_pcap_profile", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_pcap_profile", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/pcap"
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_pcap_profile", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", plan.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_pcap_profile", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error creating gigavuecore_pcap_profile", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error creating gigavuecore_pcap_profile", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error creating gigavuecore_pcap_profile", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error creating gigavuecore_pcap_profile", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error creating gigavuecore_pcap_profile", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error creating gigavuecore_pcap_profile", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error creating gigavuecore_pcap_profile", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error creating gigavuecore_pcap_profile", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error creating gigavuecore_pcap_profile", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error creating gigavuecore_pcap_profile", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_pcap_profile", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
	if plan.Id.IsNull() || plan.Id.IsUnknown() {
		loc := httpResp.Header.Get("Location")
		if loc != "" {
			plan.Id = types.StringValue(loc)
		} else {
			resp.Diagnostics.AddError("Error creating gigavuecore_pcap_profile", "The create response did not contain an identifier and no Location header was returned, so the resource cannot be tracked in state.")
			return
		}
	}
}

// Read refreshes the Terraform state with the latest remote values.
func (r *PcapProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state PcapProfileResourceModel
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
func (r *PcapProfileResource) readRemote(ctx context.Context, state *PcapProfileResourceModel, resp *resource.ReadResponse) (removed bool) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/pcap"
	httpReq, err := r.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_pcap_profile", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", state.ClusterId.ValueString())
	query.Set("portIds", state.PortIds.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_pcap_profile", fmt.Sprintf("Could not send request: %s", err))
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
			resp.Diagnostics.AddError("Error reading gigavuecore_pcap_profile", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_pcap_profile", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_pcap_profile", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_pcap_profile", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_pcap_profile", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_pcap_profile", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_pcap_profile", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_pcap_profile", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_pcap_profile", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_pcap_profile", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["pcapAll"]; ok {
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
		resp.Diagnostics.AddError("Error reading gigavuecore_pcap_profile", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
	return
}

// Update modifies the remote resource to match the desired plan.
func (r *PcapProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan PcapProfileResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state PcapProfileResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if plan.Id.IsNull() || plan.Id.IsUnknown() {
		if !state.Id.IsNull() && !state.Id.IsUnknown() {
			plan.Id = state.Id
		}
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Update is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Delete destroys the remote resource.
func (r *PcapProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state PcapProfileResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.deleteRemote(ctx, &state, resp)
}

// deleteRemote performs the delete HTTP exchange, treating a 404 as already deleted. Extracted from Delete so the request/response logic is unit-testable without a tfsdk.State.
func (r *PcapProfileResource) deleteRemote(ctx context.Context, state *PcapProfileResourceModel, resp *resource.DeleteResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/pcap"
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_pcap_profile", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", state.ClusterId.ValueString())
	query.Set("aliasList", state.AliasList.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_pcap_profile", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		return
	}
	if !(httpResp.StatusCode == 204) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error deleting gigavuecore_pcap_profile", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error deleting gigavuecore_pcap_profile", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error deleting gigavuecore_pcap_profile", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error deleting gigavuecore_pcap_profile", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error deleting gigavuecore_pcap_profile", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error deleting gigavuecore_pcap_profile", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error deleting gigavuecore_pcap_profile", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error deleting gigavuecore_pcap_profile", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error deleting gigavuecore_pcap_profile", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *PcapProfileResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
