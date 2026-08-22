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
var _ action.Action = (*ClusterConfigAddMemberAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*ClusterConfigAddMemberAction)(nil)

// ClusterConfigAddMemberAction is the generated Terraform action implementation.
type ClusterConfigAddMemberAction struct {
	client *client.Client
}

// ClusterConfigAddMemberActionModel describes the action configuration shape.
type ClusterConfigAddMemberActionModel struct {
	Bulk               types.Bool    `tfsdk:"bulk"`
	ClusterId          types.String  `tfsdk:"cluster_id"`
	ClusterMemberSpecs types.Dynamic `tfsdk:"cluster_member_specs" json:"clusterMemberSpecs"`
}

// NewClusterConfigAddMemberAction returns a new instance of the generated action.
func NewClusterConfigAddMemberAction() action.Action {
	return &ClusterConfigAddMemberAction{}
}

// Metadata returns the action type name.
func (r *ClusterConfigAddMemberAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_cluster_config_add_member"
}

// Schema returns the action schema.
func (r *ClusterConfigAddMemberAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "adds the device to the specified cluster id", Attributes: map[string]schema.Attribute{"bulk": schema.BoolAttribute{MarkdownDescription: "if more than one member needs to be added then set it to true, otherwise false", Optional: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "cluster id to which the member(s) to be added", Required: true}, "cluster_member_specs": schema.DynamicAttribute{Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *ClusterConfigAddMemberAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config ClusterConfigAddMemberActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *ClusterConfigAddMemberAction) invokeRemote(ctx context.Context, config *ClusterConfigAddMemberActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/clusterConfig/membership"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_add_member", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_add_member", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_add_member", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.Bulk.IsNull() {
		query.Set("bulk", strconv.FormatBool(config.Bulk.ValueBool()))
	}
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_add_member", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 202) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_add_member", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_add_member", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_add_member", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_add_member", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_add_member", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_add_member", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_add_member", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_add_member", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_add_member", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *ClusterConfigAddMemberAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
