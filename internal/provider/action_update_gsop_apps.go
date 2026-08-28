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
var _ action.Action = (*UpdateGsopAppsAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*UpdateGsopAppsAction)(nil)

// UpdateGsopAppsAction is the generated Terraform action implementation.
type UpdateGsopAppsAction struct {
	client *client.Client
}

// UpdateGsopAppsActionModel describes the action configuration shape.
type UpdateGsopAppsActionModel struct {
	Alias               types.String  `tfsdk:"alias"`
	Apf                 types.Dynamic `tfsdk:"apf"`
	ClusterId           types.String  `tfsdk:"cluster_id"`
	Dedup               types.Dynamic `tfsdk:"dedup"`
	DiameterWhitelist   types.Dynamic `tfsdk:"diameter_whitelist" json:"diameterWhitelist"`
	FlowFilter          types.Dynamic `tfsdk:"flow_filter" json:"flowFilter"`
	FlowSampling        types.Dynamic `tfsdk:"flow_sampling" json:"flowSampling"`
	GseriesHeaderAdd    types.Dynamic `tfsdk:"gseries_header_add" json:"gseriesHeaderAdd"`
	GseriesHeaderRemove types.Dynamic `tfsdk:"gseries_header_remove" json:"gseriesHeaderRemove"`
	GseriesLoadBalance  types.Dynamic `tfsdk:"gseries_load_balance" json:"gseriesLoadBalance"`
	GseriesPatternMatch types.Dynamic `tfsdk:"gseries_pattern_match" json:"gseriesPatternMatch"`
	GtpWhitelist        types.Dynamic `tfsdk:"gtp_whitelist" json:"gtpWhitelist"`
	HeaderAdd           types.Dynamic `tfsdk:"header_add" json:"headerAdd"`
	HeaderRemove        types.Dynamic `tfsdk:"header_remove" json:"headerRemove"`
	Icap                types.Dynamic `tfsdk:"icap"`
	InlineSsl           types.Dynamic `tfsdk:"inline_ssl" json:"inlineSsl"`
	LoadBalance         types.Dynamic `tfsdk:"load_balance" json:"loadBalance"`
	Masking             types.Dynamic `tfsdk:"masking"`
	MetadataExport      types.Dynamic `tfsdk:"metadata_export" json:"metadataExport"`
	Netflow             types.Dynamic `tfsdk:"netflow"`
	SaApf               types.Dynamic `tfsdk:"sa_apf" json:"saApf"`
	SipWhitelist        types.Dynamic `tfsdk:"sip_whitelist" json:"sipWhitelist"`
	Slicing             types.Dynamic `tfsdk:"slicing"`
	SslDecrypt          types.Dynamic `tfsdk:"ssl_decrypt" json:"sslDecrypt"`
	TrailerAdd          types.Dynamic `tfsdk:"trailer_add" json:"trailerAdd"`
	TrailerRemove       types.Dynamic `tfsdk:"trailer_remove" json:"trailerRemove"`
	TunnelDecap         types.Dynamic `tfsdk:"tunnel_decap" json:"tunnelDecap"`
	TunnelEncap         types.Dynamic `tfsdk:"tunnel_encap" json:"tunnelEncap"`
}

// NewUpdateGsopAppsAction returns a new instance of the generated action.
func NewUpdateGsopAppsAction() action.Action {
	return &UpdateGsopAppsAction{}
}

// Metadata returns the action type name.
func (r *UpdateGsopAppsAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_update_gsop_apps"
}

// Schema returns the action schema.
func (r *UpdateGsopAppsAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Update GSOP apps configuration", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "alias of the target GSOP", Required: true}, "apf": schema.DynamicAttribute{Optional: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "dedup": schema.DynamicAttribute{Optional: true}, "diameter_whitelist": schema.DynamicAttribute{Optional: true}, "flow_filter": schema.DynamicAttribute{MarkdownDescription: "GigaSMART 'Flow Filter' Application Configuration", Optional: true}, "flow_sampling": schema.DynamicAttribute{MarkdownDescription: "GigaSMART 'Flow Sampling' Application Configuration", Optional: true}, "gseries_header_add": schema.DynamicAttribute{MarkdownDescription: "Only applicable for G-series", Optional: true}, "gseries_header_remove": schema.DynamicAttribute{Optional: true}, "gseries_load_balance": schema.DynamicAttribute{MarkdownDescription: "Only applicable for G-series per-rule GSOP", Optional: true}, "gseries_pattern_match": schema.DynamicAttribute{MarkdownDescription: "Only applicable for G-series per-rule GSOP", Optional: true}, "gtp_whitelist": schema.DynamicAttribute{Optional: true}, "header_add": schema.DynamicAttribute{MarkdownDescription: "GigaSMART 'Add Header' Application Configuration", Optional: true}, "header_remove": schema.DynamicAttribute{MarkdownDescription: "GigaSMART 'Remove Header' Application Configuration", Optional: true}, "icap": schema.DynamicAttribute{MarkdownDescription: "GigaSMART ICAP Configuration", Optional: true}, "inline_ssl": schema.DynamicAttribute{MarkdownDescription: "GigaSMART Inline SSL Profile Configuration", Optional: true}, "load_balance": schema.DynamicAttribute{MarkdownDescription: "GigaSMART 'Load Balancing' Application Configuration", Optional: true}, "masking": schema.DynamicAttribute{MarkdownDescription: "GigaSMART 'Masking' Application Configuration", Optional: true}, "metadata_export": schema.DynamicAttribute{Optional: true}, "netflow": schema.DynamicAttribute{Optional: true}, "sa_apf": schema.DynamicAttribute{Optional: true}, "sip_whitelist": schema.DynamicAttribute{Optional: true}, "slicing": schema.DynamicAttribute{MarkdownDescription: "GigaSMART 'Slicing' Application Configuration", Optional: true}, "ssl_decrypt": schema.DynamicAttribute{MarkdownDescription: "GigaSMART 'SSL Decrypt' Application Configuration", Optional: true}, "trailer_add": schema.DynamicAttribute{MarkdownDescription: "GigaSMART 'Add Trailer' Application Configuration", Optional: true}, "trailer_remove": schema.DynamicAttribute{Optional: true}, "tunnel_decap": schema.DynamicAttribute{MarkdownDescription: "GigaSMART 'Decapsulate Tunnel' Application Configuration", Optional: true}, "tunnel_encap": schema.DynamicAttribute{MarkdownDescription: "GigaSMART 'Encapsulate Tunnel' Application Configuration", Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *UpdateGsopAppsAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UpdateGsopAppsActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *UpdateGsopAppsAction) invokeRemote(ctx context.Context, config *UpdateGsopAppsActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/gsops/{alias}/apps"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_gsop_apps", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_gsop_apps", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPatch, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_gsop_apps", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_gsop_apps", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_gsop_apps", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_gsop_apps", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_gsop_apps", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_gsop_apps", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_gsop_apps", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_gsop_apps", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_gsop_apps", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_update_gsop_apps", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_gsop_apps", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *UpdateGsopAppsAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
