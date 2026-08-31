package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*DownloadEventArchiveAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*DownloadEventArchiveAction)(nil)

// DownloadEventArchiveAction is the generated Terraform action implementation.
type DownloadEventArchiveAction struct {
	client *client.Client
}

// DownloadEventArchiveActionModel describes the action configuration shape.
type DownloadEventArchiveActionModel struct {
	Destination types.Object `tfsdk:"destination"`
	Purge       types.Bool   `tfsdk:"purge"`
	Scope       types.Object `tfsdk:"scope"`
}

// NewDownloadEventArchiveAction returns a new instance of the generated action.
func NewDownloadEventArchiveAction() action.Action {
	return &DownloadEventArchiveAction{}
}

// Metadata returns the action type name.
func (r *DownloadEventArchiveAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_download_event_archive"
}

// Schema returns the action schema.
func (r *DownloadEventArchiveAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Archive Events", Attributes: map[string]schema.Attribute{"destination": schema.SingleNestedAttribute{MarkdownDescription: "Destination for the archive", Required: true, Attributes: map[string]schema.Attribute{"sftp": schema.SingleNestedAttribute{MarkdownDescription: "Specifies the credentials of server where to save archive", Required: true, Attributes: map[string]schema.Attribute{"file_path": schema.StringAttribute{MarkdownDescription: "specifies the path on server to store the archive. Example: '/root/dir/archive.zip'", Required: true}, "host_address": schema.StringAttribute{MarkdownDescription: "Specifies the host address of server to save the archive", Required: true}, "password": schema.StringAttribute{MarkdownDescription: "specifies the password to use for server login", Required: true}, "username": schema.StringAttribute{MarkdownDescription: "specifies the username for the server login", Required: true}}}}}, "purge": schema.BoolAttribute{MarkdownDescription: "If true, the archived event entries will be purged", Optional: true}, "scope": schema.SingleNestedAttribute{MarkdownDescription: "Archiving scope", Required: true, Attributes: map[string]schema.Attribute{"start_date": schema.StringAttribute{MarkdownDescription: "Records older than startDate will be archived. In ISO 8601 format.", Optional: true}}}}}
}

// Invoke executes the action against the remote API.
func (r *DownloadEventArchiveAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config DownloadEventArchiveActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *DownloadEventArchiveAction) invokeRemote(ctx context.Context, config *DownloadEventArchiveActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/events/archive"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_download_event_archive", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_download_event_archive", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_download_event_archive", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.Purge.IsNull() {
		query.Set("purge", strconv.FormatBool(config.Purge.ValueBool()))
	}
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_download_event_archive", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_download_event_archive", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_download_event_archive", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_download_event_archive", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_download_event_archive", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_download_event_archive", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_download_event_archive", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_download_event_archive", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_download_event_archive", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_download_event_archive", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *DownloadEventArchiveAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
