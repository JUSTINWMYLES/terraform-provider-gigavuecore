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
var _ action.Action = (*UpdateClusterConfigBackupAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*UpdateClusterConfigBackupAction)(nil)

// UpdateClusterConfigBackupAction is the generated Terraform action implementation.
type UpdateClusterConfigBackupAction struct {
	client *client.Client
}

// UpdateClusterConfigBackupActionModel describes the action configuration shape.
type UpdateClusterConfigBackupActionModel struct {
	Alias      types.String `tfsdk:"alias"`
	BackupId   types.String `tfsdk:"backup_id" json:"backupId"`
	ClusterId  types.String `tfsdk:"cluster_id"`
	Comment    types.String `tfsdk:"comment"`
	DoNotPurge types.Bool   `tfsdk:"do_not_purge" json:"doNotPurge"`
	Tags       types.List   `tfsdk:"tags"`
}

// NewUpdateClusterConfigBackupAction returns a new instance of the generated action.
func NewUpdateClusterConfigBackupAction() action.Action {
	return &UpdateClusterConfigBackupAction{}
}

// Metadata returns the action type name.
func (r *UpdateClusterConfigBackupAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_update_cluster_config_backup"
}

// Schema returns the action schema.
func (r *UpdateClusterConfigBackupAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Update meta data of selected config backup snapshot for a given cluster", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "User-entered alias for this snapshot", Optional: true}, "backup_id": schema.StringAttribute{MarkdownDescription: "Unique identifier of the cluster config snapshot for a given cluster", Required: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Cluster Id to restore cluster config for", Required: true}, "comment": schema.StringAttribute{MarkdownDescription: "User-entered comments for this snapshot", Optional: true}, "do_not_purge": schema.BoolAttribute{MarkdownDescription: "Indicates whether this snapshot should not be auto-aged/purged", Optional: true}, "tags": schema.ListAttribute{MarkdownDescription: "Tags associated with this file", Optional: true, ElementType: types.StringType}}}
}

// Invoke executes the action against the remote API.
func (r *UpdateClusterConfigBackupAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UpdateClusterConfigBackupActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *UpdateClusterConfigBackupAction) invokeRemote(ctx context.Context, config *UpdateClusterConfigBackupActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/clusterConfig/backup/repo/{clusterId}"
	reqPath = strings.ReplaceAll(reqPath, "{clusterId}", url.PathEscape(config.ClusterId.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_cluster_config_backup", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_cluster_config_backup", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPatch, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_cluster_config_backup", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_cluster_config_backup", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_cluster_config_backup", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_cluster_config_backup", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_cluster_config_backup", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_cluster_config_backup", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_cluster_config_backup", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_cluster_config_backup", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_cluster_config_backup", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_update_cluster_config_backup", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_cluster_config_backup", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *UpdateClusterConfigBackupAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
