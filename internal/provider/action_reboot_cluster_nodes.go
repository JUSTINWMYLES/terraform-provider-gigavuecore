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
var _ action.Action = (*RebootClusterNodesAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*RebootClusterNodesAction)(nil)

// RebootClusterNodesAction is the generated Terraform action implementation.
type RebootClusterNodesAction struct {
	client *client.Client
}

// RebootClusterNodesActionModel describes the action configuration shape.
type RebootClusterNodesActionModel struct {
	ClusterIds              types.List `tfsdk:"cluster_ids" json:"clusterIds"`
	NodeIds                 types.List `tfsdk:"node_ids" json:"nodeIds"`
	SkipNotReachableDevices types.Bool `tfsdk:"skip_not_reachable_devices" json:"skipNotReachableDevices"`
}

// NewRebootClusterNodesAction returns a new instance of the generated action.
func NewRebootClusterNodesAction() action.Action {
	return &RebootClusterNodesAction{}
}

// Metadata returns the action type name.
func (r *RebootClusterNodesAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_reboot_cluster_nodes"
}

// Schema returns the action schema.
func (r *RebootClusterNodesAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "reboot physical cluster nodes", Attributes: map[string]schema.Attribute{"cluster_ids": schema.ListAttribute{MarkdownDescription: "ids of clusters to reboot. Every node in these clusters will be rebooted", Optional: true, ElementType: types.StringType}, "node_ids": schema.ListAttribute{MarkdownDescription: "ids of individual nodes to reboot", Optional: true, ElementType: types.StringType}, "skip_not_reachable_devices": schema.BoolAttribute{MarkdownDescription: "indicates whether not-reachable nodes(if any) should be skipped and continue with reboot operation", Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *RebootClusterNodesAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config RebootClusterNodesActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *RebootClusterNodesAction) invokeRemote(ctx context.Context, config *RebootClusterNodesActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/clusterConfig/reboot"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_reboot_cluster_nodes", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_reboot_cluster_nodes", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_reboot_cluster_nodes", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_reboot_cluster_nodes", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_reboot_cluster_nodes", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_reboot_cluster_nodes", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_reboot_cluster_nodes", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_reboot_cluster_nodes", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_reboot_cluster_nodes", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_reboot_cluster_nodes", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_reboot_cluster_nodes", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_reboot_cluster_nodes", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_reboot_cluster_nodes", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *RebootClusterNodesAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
