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
var _ action.Action = (*FetchInlineSslProfileListAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*FetchInlineSslProfileListAction)(nil)

// FetchInlineSslProfileListAction is the generated Terraform action implementation.
type FetchInlineSslProfileListAction struct {
	client *client.Client
}

// FetchInlineSslProfileListActionModel describes the action configuration shape.
type FetchInlineSslProfileListActionModel struct {
	Alias      types.String `tfsdk:"alias"`
	ClusterId  types.String `tfsdk:"cluster_id"`
	FileSource types.Object `tfsdk:"file_source" json:"fileSource"`
	List       types.String `tfsdk:"list"`
	ListType   types.String `tfsdk:"list_type"`
}

// NewFetchInlineSslProfileListAction returns a new instance of the generated action.
func NewFetchInlineSslProfileListAction() action.Action {
	return &FetchInlineSslProfileListAction{}
}

// Metadata returns the action type name.
func (r *FetchInlineSslProfileListAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_fetch_inline_ssl_profile_list"
}

// Schema returns the action schema.
func (r *FetchInlineSslProfileListAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Fetch inline SSL profile nodecryptlist or decryptlist", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Alias of the inline SSL profile", Required: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "file_source": schema.SingleNestedAttribute{MarkdownDescription: "Remote file source or destination", Optional: true, Attributes: map[string]schema.Attribute{"hostname": schema.StringAttribute{MarkdownDescription: "server address", Required: true}, "password": schema.StringAttribute{MarkdownDescription: "password to use for server login", Optional: true}, "path": schema.StringAttribute{MarkdownDescription: "file path on server", Required: true}, "protocol": schema.StringAttribute{MarkdownDescription: "Protocol to access server. HTTP and HTTPS only applicable for retrieving file", Required: true}, "username": schema.StringAttribute{MarkdownDescription: "user name to use for server login", Optional: true}}}, "list": schema.StringAttribute{MarkdownDescription: "The nodecryptlist or decryptlist . Mutually exclusive with 'fileSource'", Optional: true}, "list_type": schema.StringAttribute{MarkdownDescription: "specify a nodecryptlist or decryptlist. (deprecated: use nodecryptlist and decryptlist instead of whitelist and blacklist)", Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *FetchInlineSslProfileListAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config FetchInlineSslProfileListActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *FetchInlineSslProfileListAction) invokeRemote(ctx context.Context, config *FetchInlineSslProfileListActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/inlineSsl/profiles/{alias}/list/{listType}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	reqPath = strings.ReplaceAll(reqPath, "{listType}", url.PathEscape(config.ListType.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_fetch_inline_ssl_profile_list", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_fetch_inline_ssl_profile_list", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_fetch_inline_ssl_profile_list", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_fetch_inline_ssl_profile_list", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_fetch_inline_ssl_profile_list", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_fetch_inline_ssl_profile_list", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_fetch_inline_ssl_profile_list", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_fetch_inline_ssl_profile_list", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_fetch_inline_ssl_profile_list", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_fetch_inline_ssl_profile_list", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_fetch_inline_ssl_profile_list", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_fetch_inline_ssl_profile_list", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_fetch_inline_ssl_profile_list", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *FetchInlineSslProfileListAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
