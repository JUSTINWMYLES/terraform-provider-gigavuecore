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
var _ action.Action = (*UpgradeClusterNodesImageAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*UpgradeClusterNodesImageAction)(nil)

// UpgradeClusterNodesImageAction is the generated Terraform action implementation.
type UpgradeClusterNodesImageAction struct {
	client *client.Client
}

// UpgradeClusterNodesImageActionModel describes the action configuration shape.
type UpgradeClusterNodesImageActionModel struct {
	Async                   types.Bool    `tfsdk:"async"`
	ClusterIds              types.List    `tfsdk:"cluster_ids" json:"clusterIds"`
	ImageFileSpecs          types.Dynamic `tfsdk:"image_file_specs" json:"imageFileSpecs"`
	ImageServer             types.String  `tfsdk:"image_server" json:"imageServer"`
	NodeIds                 types.List    `tfsdk:"node_ids" json:"nodeIds"`
	Reboot                  types.Bool    `tfsdk:"reboot"`
	SkipNotReachableDevices types.Bool    `tfsdk:"skip_not_reachable_devices" json:"skipNotReachableDevices"`
}

// NewUpgradeClusterNodesImageAction returns a new instance of the generated action.
func NewUpgradeClusterNodesImageAction() action.Action {
	return &UpgradeClusterNodesImageAction{}
}

// Metadata returns the action type name.
func (r *UpgradeClusterNodesImageAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_upgrade_cluster_nodes_image"
}

// Schema returns the action schema.
func (r *UpgradeClusterNodesImageAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "upgrade cluster nodes image", Attributes: map[string]schema.Attribute{"async": schema.BoolAttribute{MarkdownDescription: "if provided, the call returns immediately with the [202 Accepted] HTTP status code, and the image upgrade will complete in the background", Optional: true}, "cluster_ids": schema.ListAttribute{MarkdownDescription: "ids of clusters to upgrade image for. Every node in these clusters is upgraded", Required: true, ElementType: types.StringType}, "image_file_specs": schema.DynamicAttribute{MarkdownDescription: "lists image files to use", Required: true}, "image_server": schema.StringAttribute{MarkdownDescription: "alias of an image file server. has to reference one of the existing image file server profiles", Required: true}, "node_ids": schema.ListAttribute{MarkdownDescription: "ids of individual nodes to upgrade image for", Optional: true, ElementType: types.StringType}, "reboot": schema.BoolAttribute{MarkdownDescription: "indicates whether nodes should reboot after image upgrade", Required: true}, "skip_not_reachable_devices": schema.BoolAttribute{MarkdownDescription: "indicates whether not-reachable nodes(if any) should be skipped and continue with image upgrade", Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *UpgradeClusterNodesImageAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UpgradeClusterNodesImageActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *UpgradeClusterNodesImageAction) invokeRemote(ctx context.Context, config *UpgradeClusterNodesImageActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/clusterConfig/imageUpgrade"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_upgrade_cluster_nodes_image", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_upgrade_cluster_nodes_image", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_upgrade_cluster_nodes_image", fmt.Sprintf("Could not build request: %s", err))
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
		resp.Diagnostics.AddError("Error invoking gigavuecore_upgrade_cluster_nodes_image", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_upgrade_cluster_nodes_image", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_upgrade_cluster_nodes_image", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_upgrade_cluster_nodes_image", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_upgrade_cluster_nodes_image", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_upgrade_cluster_nodes_image", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_upgrade_cluster_nodes_image", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_upgrade_cluster_nodes_image", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_upgrade_cluster_nodes_image", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_upgrade_cluster_nodes_image", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *UpgradeClusterNodesImageAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
