package provider

import (
	"context"
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
var _ action.Action = (*ClusterConfigDeleteMemberAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*ClusterConfigDeleteMemberAction)(nil)

// ClusterConfigDeleteMemberAction is the generated Terraform action implementation.
type ClusterConfigDeleteMemberAction struct {
	client *client.Client
}

// ClusterConfigDeleteMemberActionModel describes the action configuration shape.
type ClusterConfigDeleteMemberActionModel struct {
	BoxId     types.String `tfsdk:"box_id"`
	ClusterId types.String `tfsdk:"cluster_id"`
}

// NewClusterConfigDeleteMemberAction returns a new instance of the generated action.
func NewClusterConfigDeleteMemberAction() action.Action {
	return &ClusterConfigDeleteMemberAction{}
}

// Metadata returns the action type name.
func (r *ClusterConfigDeleteMemberAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_cluster_config_delete_member"
}

// Schema returns the action schema.
func (r *ClusterConfigDeleteMemberAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "removes the specified member from the cluster", Attributes: map[string]schema.Attribute{"box_id": schema.StringAttribute{MarkdownDescription: "box id of the device that needs to be removed from the cluster", Required: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "cluster id to which the device belongs", Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *ClusterConfigDeleteMemberAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config ClusterConfigDeleteMemberActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *ClusterConfigDeleteMemberAction) invokeRemote(ctx context.Context, config *ClusterConfigDeleteMemberActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/clusterConfig/membership"
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_delete_member", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("boxId", config.BoxId.ValueString())
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_delete_member", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 202) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_delete_member", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_delete_member", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_delete_member", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_delete_member", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_delete_member", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_delete_member", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_delete_member", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_delete_member", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_delete_member", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *ClusterConfigDeleteMemberAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
