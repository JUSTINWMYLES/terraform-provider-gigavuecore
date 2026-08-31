package provider

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*DeleteSipWhitelistEntryAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*DeleteSipWhitelistEntryAction)(nil)

// DeleteSipWhitelistEntryAction is the generated Terraform action implementation.
type DeleteSipWhitelistEntryAction struct {
	client *client.Client
}

// DeleteSipWhitelistEntryActionModel describes the action configuration shape.
type DeleteSipWhitelistEntryActionModel struct {
	Alias     types.String `tfsdk:"alias"`
	CallerId  types.String `tfsdk:"caller_id"`
	IdRange   types.String `tfsdk:"id_range"`
	IpAddress types.String `tfsdk:"ip_address"`
}

// NewDeleteSipWhitelistEntryAction returns a new instance of the generated action.
func NewDeleteSipWhitelistEntryAction() action.Action {
	return &DeleteSipWhitelistEntryAction{}
}

// Metadata returns the action type name.
func (r *DeleteSipWhitelistEntryAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_delete_sip_whitelist_entry"
}

// Schema returns the action schema.
func (r *DeleteSipWhitelistEntryAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Delete SIP Whitelist Entry", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "alias of the target SIP Whitelist", Required: true}, "caller_id": schema.StringAttribute{MarkdownDescription: "callerId-based whitelist entry being deleted. required till H 5.6", Optional: true}, "id_range": schema.StringAttribute{MarkdownDescription: "idrange based whitelist entry being deleted. Example:idRange=110..120", Optional: true}, "ip_address": schema.StringAttribute{MarkdownDescription: "ipAddress based whitelist entry being deleted. Example:ipAddress=1.1.1.1", Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *DeleteSipWhitelistEntryAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config DeleteSipWhitelistEntryActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *DeleteSipWhitelistEntryAction) invokeRemote(ctx context.Context, config *DeleteSipWhitelistEntryActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/sip/whitelists/{alias}/entries"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_delete_sip_whitelist_entry", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.CallerId.IsNull() {
		query.Set("callerId", config.CallerId.ValueString())
	}
	if !config.IdRange.IsNull() {
		query.Set("idRange", config.IdRange.ValueString())
	}
	if !config.IpAddress.IsNull() {
		query.Set("ipAddress", config.IpAddress.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_delete_sip_whitelist_entry", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 204) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_sip_whitelist_entry", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_sip_whitelist_entry", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_sip_whitelist_entry", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_sip_whitelist_entry", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_sip_whitelist_entry", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_sip_whitelist_entry", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_sip_whitelist_entry", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_delete_sip_whitelist_entry", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_sip_whitelist_entry", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *DeleteSipWhitelistEntryAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
