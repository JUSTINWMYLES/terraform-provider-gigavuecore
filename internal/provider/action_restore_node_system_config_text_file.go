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
var _ action.Action = (*RestoreNodeSystemConfigTextFileAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*RestoreNodeSystemConfigTextFileAction)(nil)

// RestoreNodeSystemConfigTextFileAction is the generated Terraform action implementation.
type RestoreNodeSystemConfigTextFileAction struct {
	client *client.Client
}

// RestoreNodeSystemConfigTextFileActionModel describes the action configuration shape.
type RestoreNodeSystemConfigTextFileActionModel struct {
	ClearConfig  types.Bool   `tfsdk:"clear_config"`
	ClusterId    types.String `tfsdk:"cluster_id"`
	Destination  types.Object `tfsdk:"destination"`
	FailContinue types.Bool   `tfsdk:"fail_continue"`
	Filename     types.String `tfsdk:"filename"`
	Protocol     types.String `tfsdk:"protocol"`
	Remote       types.Bool   `tfsdk:"remote"`
}

// NewRestoreNodeSystemConfigTextFileAction returns a new instance of the generated action.
func NewRestoreNodeSystemConfigTextFileAction() action.Action {
	return &RestoreNodeSystemConfigTextFileAction{}
}

// Metadata returns the action type name.
func (r *RestoreNodeSystemConfigTextFileAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_restore_node_system_config_text_file"
}

// Schema returns the action schema.
func (r *RestoreNodeSystemConfigTextFileAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Apply the text configuration to the running system from a remote host or local storage", Attributes: map[string]schema.Attribute{"clear_config": schema.BoolAttribute{MarkdownDescription: "If true, clear the existing traffic configuration before applying the text configuration file.", Optional: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "destination": schema.SingleNestedAttribute{Required: true, Attributes: map[string]schema.Attribute{"hostname": schema.StringAttribute{MarkdownDescription: "server address", Required: true}, "password": schema.StringAttribute{MarkdownDescription: "password to use for server login", Optional: true}, "path": schema.StringAttribute{MarkdownDescription: "configuration text file path on server", Required: true}, "username": schema.StringAttribute{MarkdownDescription: "user name to use for server login", Optional: true}}}, "fail_continue": schema.BoolAttribute{MarkdownDescription: "If true, while applying commands, continue execution even if one of them fails", Required: true}, "filename": schema.StringAttribute{MarkdownDescription: "File name of the text configuration file to be applied on system", Required: true}, "protocol": schema.StringAttribute{MarkdownDescription: "File Source/File Destination protocol. http and https only applicable for fileSource ", Required: true}, "remote": schema.BoolAttribute{MarkdownDescription: "If true, fetch the text configuration file from remote host.If false, fetch file from local storage", Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *RestoreNodeSystemConfigTextFileAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config RestoreNodeSystemConfigTextFileActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *RestoreNodeSystemConfigTextFileAction) invokeRemote(ctx context.Context, config *RestoreNodeSystemConfigTextFileActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/system/config/text/runningConfig"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_restore_node_system_config_text_file", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_restore_node_system_config_text_file", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_restore_node_system_config_text_file", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	query.Set("filename", config.Filename.ValueString())
	query.Set("remote", strconv.FormatBool(config.Remote.ValueBool()))
	query.Set("failContinue", strconv.FormatBool(config.FailContinue.ValueBool()))
	if !config.ClearConfig.IsNull() {
		query.Set("clearConfig", strconv.FormatBool(config.ClearConfig.ValueBool()))
	}
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_restore_node_system_config_text_file", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_restore_node_system_config_text_file", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_restore_node_system_config_text_file", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_restore_node_system_config_text_file", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_restore_node_system_config_text_file", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_restore_node_system_config_text_file", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_restore_node_system_config_text_file", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_restore_node_system_config_text_file", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_restore_node_system_config_text_file", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_restore_node_system_config_text_file", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *RestoreNodeSystemConfigTextFileAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
