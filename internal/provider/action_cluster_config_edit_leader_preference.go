package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*ClusterConfigEditLeaderPreferenceAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*ClusterConfigEditLeaderPreferenceAction)(nil)

// ClusterConfigEditLeaderPreferenceAction is the generated Terraform action implementation.
type ClusterConfigEditLeaderPreferenceAction struct {
	client *client.Client
}

// ClusterConfigEditLeaderPreferenceActionModel describes the action configuration shape.
type ClusterConfigEditLeaderPreferenceActionModel struct {
	BoxId            types.String `tfsdk:"box_id"`
	ClusterId        types.String `tfsdk:"cluster_id"`
	LeaderPreference types.Int64  `tfsdk:"leader_preference" json:"leaderPreference"`
	MasterPreference types.Int64  `tfsdk:"master_preference" json:"masterPreference"`
}

// NewClusterConfigEditLeaderPreferenceAction returns a new instance of the generated action.
func NewClusterConfigEditLeaderPreferenceAction() action.Action {
	return &ClusterConfigEditLeaderPreferenceAction{}
}

// Metadata returns the action type name.
func (r *ClusterConfigEditLeaderPreferenceAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_cluster_config_edit_leader_preference"
}

// Schema returns the action schema.
func (r *ClusterConfigEditLeaderPreferenceAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "update the leader preference of the specified member", Attributes: map[string]schema.Attribute{"box_id": schema.StringAttribute{MarkdownDescription: "box id of the device whose leader preference needs to be changed", Required: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "cluster id to which the device belongs", Required: true}, "leader_preference": schema.Int64Attribute{MarkdownDescription: "leader election preference rank, 1 to 9 to exclude leadership", Optional: true}, "master_preference": schema.Int64Attribute{MarkdownDescription: "master election preference rank, 1 to 9 to exclude mastership. (deprecated: use leaderPreference)", Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *ClusterConfigEditLeaderPreferenceAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config ClusterConfigEditLeaderPreferenceActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *ClusterConfigEditLeaderPreferenceAction) invokeRemote(ctx context.Context, config *ClusterConfigEditLeaderPreferenceActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/clusterConfig/membership"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_edit_leader_preference", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_edit_leader_preference", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPatch, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_edit_leader_preference", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("boxId", config.BoxId.ValueString())
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_edit_leader_preference", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 202) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_edit_leader_preference", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_edit_leader_preference", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_edit_leader_preference", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_edit_leader_preference", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_edit_leader_preference", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_edit_leader_preference", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_edit_leader_preference", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_edit_leader_preference", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_edit_leader_preference", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *ClusterConfigEditLeaderPreferenceAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
