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
var _ action.Action = (*UpgradeFmImageAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*UpgradeFmImageAction)(nil)

// UpgradeFmImageAction is the generated Terraform action implementation.
type UpgradeFmImageAction struct {
	client *client.Client
}

// UpgradeFmImageActionModel describes the action configuration shape.
type UpgradeFmImageActionModel struct {
	Async       types.Bool   `tfsdk:"async"`
	FilePath    types.String `tfsdk:"file_path" json:"filePath"`
	ImageServer types.String `tfsdk:"image_server" json:"imageServer"`
	Reboot      types.Bool   `tfsdk:"reboot"`
}

// NewUpgradeFmImageAction returns a new instance of the generated action.
func NewUpgradeFmImageAction() action.Action {
	return &UpgradeFmImageAction{}
}

// Metadata returns the action type name.
func (r *UpgradeFmImageAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_upgrade_fm_image"
}

// Schema returns the action schema.
func (r *UpgradeFmImageAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Upgrade FM image", Attributes: map[string]schema.Attribute{"async": schema.BoolAttribute{MarkdownDescription: "if provided, the call returns immediately with the [202 Accepted] HTTP status code, and the image upgrade will complete in the background", Optional: true}, "file_path": schema.StringAttribute{MarkdownDescription: "image file path on the image server", Required: true}, "image_server": schema.StringAttribute{MarkdownDescription: "alias of an image file server. has to reference one of the existing image file server profiles", Required: true}, "reboot": schema.BoolAttribute{MarkdownDescription: "indicates whether nodes should reboot after image upgrade", Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *UpgradeFmImageAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UpgradeFmImageActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *UpgradeFmImageAction) invokeRemote(ctx context.Context, config *UpgradeFmImageActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/fmSystem/imageUpgrade"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_upgrade_fm_image", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_upgrade_fm_image", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_upgrade_fm_image", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.Async.IsNull() {
		query.Set("async", strconv.FormatBool(config.Async.ValueBool()))
	}
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_upgrade_fm_image", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_upgrade_fm_image", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_upgrade_fm_image", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_upgrade_fm_image", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_upgrade_fm_image", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_upgrade_fm_image", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_upgrade_fm_image", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_upgrade_fm_image", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_upgrade_fm_image", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_upgrade_fm_image", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *UpgradeFmImageAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
