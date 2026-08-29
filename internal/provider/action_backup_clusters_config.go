package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*BackupClustersConfigAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*BackupClustersConfigAction)(nil)

// BackupClustersConfigAction is the generated Terraform action implementation.
type BackupClustersConfigAction struct {
	client *client.Client
}

// BackupClustersConfigActionModel describes the action configuration shape.
type BackupClustersConfigActionModel struct {
	ClustersConfigSpecs types.List `tfsdk:"clusters_config_specs" json:"clustersConfigSpecs"`
}

// NewBackupClustersConfigAction returns a new instance of the generated action.
func NewBackupClustersConfigAction() action.Action {
	return &BackupClustersConfigAction{}
}

// Metadata returns the action type name.
func (r *BackupClustersConfigAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_backup_clusters_config"
}

// Schema returns the action schema.
func (r *BackupClustersConfigAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Backup of multiple clusters can be requested", Attributes: map[string]schema.Attribute{"clusters_config_specs": schema.ListNestedAttribute{MarkdownDescription: "list of Clusters config backup references", Required: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Cluster Id", Required: true}}}}}}
}

// Invoke executes the action against the remote API.
func (r *BackupClustersConfigAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config BackupClustersConfigActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *BackupClustersConfigAction) invokeRemote(ctx context.Context, config *BackupClustersConfigActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/clusterConfig/backup/device/backup"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_backup_clusters_config", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_backup_clusters_config", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_backup_clusters_config", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_backup_clusters_config", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_backup_clusters_config", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_backup_clusters_config", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_backup_clusters_config", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_backup_clusters_config", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_backup_clusters_config", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_backup_clusters_config", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_backup_clusters_config", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_backup_clusters_config", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_backup_clusters_config", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *BackupClustersConfigAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
