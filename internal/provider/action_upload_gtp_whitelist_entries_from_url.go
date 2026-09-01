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
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*UploadGtpWhitelistEntriesFromUrlAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*UploadGtpWhitelistEntriesFromUrlAction)(nil)

// UploadGtpWhitelistEntriesFromUrlAction is the generated Terraform action implementation.
type UploadGtpWhitelistEntriesFromUrlAction struct {
	client *client.Client
}

// UploadGtpWhitelistEntriesFromUrlActionModel describes the action configuration shape.
type UploadGtpWhitelistEntriesFromUrlActionModel struct {
	Alias          types.String `tfsdk:"alias"`
	ClusterId      types.String `tfsdk:"cluster_id"`
	Url            types.String `tfsdk:"url"`
	Usage          types.String `tfsdk:"usage"`
	WhitelistAlias types.String `tfsdk:"whitelist_alias" json:"whitelistAlias"`
}

// NewUploadGtpWhitelistEntriesFromUrlAction returns a new instance of the generated action.
func NewUploadGtpWhitelistEntriesFromUrlAction() action.Action {
	return &UploadGtpWhitelistEntriesFromUrlAction{}
}

// Metadata returns the action type name.
func (r *UploadGtpWhitelistEntriesFromUrlAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_upload_gtp_whitelist_entries_from_url"
}

// Schema returns the action schema.
func (r *UploadGtpWhitelistEntriesFromUrlAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Upload GTP Whitelist Entries from URL", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "alias of the target GTP Whitelist", Required: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "url": schema.StringAttribute{MarkdownDescription: "URL of the Whitelist file. format: protocol://[username[:password]]@hostname[:port]/path/filename", Required: true}, "usage": schema.StringAttribute{MarkdownDescription: "When 'delete' is specified, the uploaded list is treated as a delete request and all the entries in the list should be removed from the existing whitelist identified by the 'alias' parameter", Optional: true}, "whitelist_alias": schema.StringAttribute{MarkdownDescription: "alias of the target GTP Whitelist", Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *UploadGtpWhitelistEntriesFromUrlAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UploadGtpWhitelistEntriesFromUrlActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *UploadGtpWhitelistEntriesFromUrlAction) invokeRemote(ctx context.Context, config *UploadGtpWhitelistEntriesFromUrlActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/gtp/whitelists/{alias}/entries/url"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_upload_gtp_whitelist_entries_from_url", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_upload_gtp_whitelist_entries_from_url", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_upload_gtp_whitelist_entries_from_url", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_upload_gtp_whitelist_entries_from_url", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_upload_gtp_whitelist_entries_from_url", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_upload_gtp_whitelist_entries_from_url", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_upload_gtp_whitelist_entries_from_url", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_upload_gtp_whitelist_entries_from_url", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_upload_gtp_whitelist_entries_from_url", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_upload_gtp_whitelist_entries_from_url", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_upload_gtp_whitelist_entries_from_url", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_upload_gtp_whitelist_entries_from_url", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_upload_gtp_whitelist_entries_from_url", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *UploadGtpWhitelistEntriesFromUrlAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
