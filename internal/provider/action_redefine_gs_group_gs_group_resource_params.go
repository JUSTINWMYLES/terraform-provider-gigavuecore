package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*RedefineGsGroupGsGroupResourceParamsAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*RedefineGsGroupGsGroupResourceParamsAction)(nil)

// RedefineGsGroupGsGroupResourceParamsAction is the generated Terraform action implementation.
type RedefineGsGroupGsGroupResourceParamsAction struct {
	client *client.Client
}

// RedefineGsGroupGsGroupResourceParamsActionModel describes the action configuration shape.
type RedefineGsGroupGsGroupResourceParamsActionModel struct {
	Alias           types.String `tfsdk:"alias"`
	BufferAsfSize   types.Int64  `tfsdk:"buffer_asf_size" json:"bufferAsfSize"`
	Cpu             types.Object `tfsdk:"cpu"`
	HsmSsl          types.Object `tfsdk:"hsm_ssl" json:"hsmSsl"`
	InlineSsl       types.Object `tfsdk:"inline_ssl" json:"inlineSsl"`
	Metadata        types.Int64  `tfsdk:"metadata"`
	PacketBuffer    types.Object `tfsdk:"packet_buffer" json:"packetBuffer"`
	SessionOverload types.Object `tfsdk:"session_overload" json:"sessionOverload"`
	TunnelOverload  types.Object `tfsdk:"tunnel_overload" json:"tunnelOverload"`
	XpktMatch       types.Object `tfsdk:"xpkt_match" json:"xpktMatch"`
}

// NewRedefineGsGroupGsGroupResourceParamsAction returns a new instance of the generated action.
func NewRedefineGsGroupGsGroupResourceParamsAction() action.Action {
	return &RedefineGsGroupGsGroupResourceParamsAction{}
}

// Metadata returns the action type name.
func (r *RedefineGsGroupGsGroupResourceParamsAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_redefine_gs_group_gs_group_resource_params"
}

// Schema returns the action schema.
func (r *RedefineGsGroupGsGroupResourceParamsAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Redefine GS Group's resource Params", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "alias of the target GS Group", Required: true}, "buffer_asf_size": schema.Int64Attribute{MarkdownDescription: "Session aware APF buffer. Valid values: 2-5: in million sessions (0 = disabled). Changes take effect after card or system reboot", Optional: true}, "cpu": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"overload_threshold": schema.Int64Attribute{MarkdownDescription: "Valid values 50 - 90%. 0 is disabled. Changes take effect after card or system reboot", Optional: true}}}, "hsm_ssl": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Resource Hsm Ssl Parameters", Optional: true, Attributes: map[string]schema.Attribute{"buffer": schema.Int64Attribute{MarkdownDescription: "resource for application hsm-ssl buffer in MB. 0 to disable", Optional: true}, "packet_buffer": schema.Int64Attribute{MarkdownDescription: "resource for application hsm-ssl packet-buffer per connection", Optional: true}, "session_count": schema.Int64Attribute{MarkdownDescription: "resource for application hsm-ssl buffer session count in million, 0 to disable", Optional: true}}}, "inline_ssl": schema.SingleNestedAttribute{MarkdownDescription: "Used to configure other GS apps in addition to Inline SSL on a HC1 box", Optional: true, Attributes: map[string]schema.Attribute{"standalone": schema.BoolAttribute{MarkdownDescription: "If enabled , behaves in a normal way. If disabled, can configure other GS apps along side Inline ssl with Inline ssl occupying only 50% of the memory", Optional: true}}}, "metadata": schema.Int64Attribute{MarkdownDescription: "flows in millions, how many flows to support for metadata. 0 to disable", Optional: true}, "packet_buffer": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"overload_threshold": schema.Int64Attribute{MarkdownDescription: "Valid values 50 - 80%. 0 is disabled. Changes take effect after card or system reboot", Optional: true}}}, "session_overload": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"overload_threshold": schema.Int64Attribute{MarkdownDescription: "Session overload threshold value , Default value is 90 and 0 is disabled.", Optional: true}}}, "tunnel_overload": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"overload_threshold": schema.Int64Attribute{MarkdownDescription: "Tunnel overload threshold value , Default value is 90 and 0 is disabled.", Optional: true}}}, "xpkt_match": schema.SingleNestedAttribute{MarkdownDescription: "GsGroup Resource Cross Packet Match Parameters", Optional: true, Attributes: map[string]schema.Attribute{"flows": schema.Int64Attribute{MarkdownDescription: "num in 100K flows. 0 is disable", Optional: true}}}}}
}

// Invoke executes the action against the remote API.
func (r *RedefineGsGroupGsGroupResourceParamsAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config RedefineGsGroupGsGroupResourceParamsActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *RedefineGsGroupGsGroupResourceParamsAction) invokeRemote(ctx context.Context, config *RedefineGsGroupGsGroupResourceParamsActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/gsGroups/{alias}/params/resource"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_gs_group_resource_params", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_gs_group_resource_params", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_gs_group_resource_params", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_gs_group_resource_params", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_gs_group_resource_params", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_gs_group_resource_params", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_gs_group_resource_params", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_gs_group_resource_params", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_gs_group_resource_params", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_gs_group_resource_params", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_gs_group_resource_params", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_gs_group_resource_params", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_gs_group_resource_params", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *RedefineGsGroupGsGroupResourceParamsAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Action Configure Type", fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData))
		return
	}
	r.client = c
}
