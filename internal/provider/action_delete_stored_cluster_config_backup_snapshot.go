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
var _ action.Action = (*DeleteStoredClusterConfigBackupSnapshotAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*DeleteStoredClusterConfigBackupSnapshotAction)(nil)

// DeleteStoredClusterConfigBackupSnapshotAction is the generated Terraform action implementation.
type DeleteStoredClusterConfigBackupSnapshotAction struct {
	client *client.Client
}

// DeleteStoredClusterConfigBackupSnapshotActionModel describes the action configuration shape.
type DeleteStoredClusterConfigBackupSnapshotActionModel struct {
	BackupId  types.String `tfsdk:"backup_id"`
	ClusterId types.String `tfsdk:"cluster_id"`
}

// NewDeleteStoredClusterConfigBackupSnapshotAction returns a new instance of the generated action.
func NewDeleteStoredClusterConfigBackupSnapshotAction() action.Action {
	return &DeleteStoredClusterConfigBackupSnapshotAction{}
}

// Metadata returns the action type name.
func (r *DeleteStoredClusterConfigBackupSnapshotAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_delete_stored_cluster_config_backup_snapshot"
}

// Schema returns the action schema.
func (r *DeleteStoredClusterConfigBackupSnapshotAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Delete a labeled config backup snapshot", Attributes: map[string]schema.Attribute{"backup_id": schema.StringAttribute{MarkdownDescription: "id of the config backup snapshot to delete", Required: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "ID of the target cluster", Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *DeleteStoredClusterConfigBackupSnapshotAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config DeleteStoredClusterConfigBackupSnapshotActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *DeleteStoredClusterConfigBackupSnapshotAction) invokeRemote(ctx context.Context, config *DeleteStoredClusterConfigBackupSnapshotActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/clusterConfig/backup/repo/{clusterId}/{backupId}"
	reqPath = strings.ReplaceAll(reqPath, "{clusterId}", url.PathEscape(config.ClusterId.ValueString()))
	reqPath = strings.ReplaceAll(reqPath, "{backupId}", url.PathEscape(config.BackupId.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_delete_stored_cluster_config_backup_snapshot", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_delete_stored_cluster_config_backup_snapshot", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 204) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_stored_cluster_config_backup_snapshot", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_stored_cluster_config_backup_snapshot", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_stored_cluster_config_backup_snapshot", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_stored_cluster_config_backup_snapshot", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_stored_cluster_config_backup_snapshot", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_stored_cluster_config_backup_snapshot", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_stored_cluster_config_backup_snapshot", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_delete_stored_cluster_config_backup_snapshot", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_stored_cluster_config_backup_snapshot", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *DeleteStoredClusterConfigBackupSnapshotAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
