package provider

import (
	"context"
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
var _ action.Action = (*DeleteDiameterWhitelistEntryAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*DeleteDiameterWhitelistEntryAction)(nil)

// DeleteDiameterWhitelistEntryAction is the generated Terraform action implementation.
type DeleteDiameterWhitelistEntryAction struct {
	client *client.Client
}

// DeleteDiameterWhitelistEntryActionModel describes the action configuration shape.
type DeleteDiameterWhitelistEntryActionModel struct {
	Alias    types.String `tfsdk:"alias"`
	UserName types.String `tfsdk:"user_name"`
}

// NewDeleteDiameterWhitelistEntryAction returns a new instance of the generated action.
func NewDeleteDiameterWhitelistEntryAction() action.Action {
	return &DeleteDiameterWhitelistEntryAction{}
}

// Metadata returns the action type name.
func (r *DeleteDiameterWhitelistEntryAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_delete_diameter_whitelist_entry"
}

// Schema returns the action schema.
func (r *DeleteDiameterWhitelistEntryAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Delete Diameter Whitelist Entry", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "alias of the target Diameter Whitelist", Required: true}, "user_name": schema.StringAttribute{MarkdownDescription: "imsi-based whitelist entry being deleted.if not provided deletes all entries", Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *DeleteDiameterWhitelistEntryAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config DeleteDiameterWhitelistEntryActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *DeleteDiameterWhitelistEntryAction) invokeRemote(ctx context.Context, config *DeleteDiameterWhitelistEntryActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/diameter/whitelists/{alias}/entries"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_delete_diameter_whitelist_entry", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.UserName.IsNull() {
		query.Set("userName", config.UserName.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_delete_diameter_whitelist_entry", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 204) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_diameter_whitelist_entry", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_diameter_whitelist_entry", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_diameter_whitelist_entry", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_diameter_whitelist_entry", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_diameter_whitelist_entry", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_diameter_whitelist_entry", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_diameter_whitelist_entry", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_delete_diameter_whitelist_entry", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_diameter_whitelist_entry", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *DeleteDiameterWhitelistEntryAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
