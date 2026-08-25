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
var _ action.Action = (*UpgradeClusterGsCardImageAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*UpgradeClusterGsCardImageAction)(nil)

// UpgradeClusterGsCardImageAction is the generated Terraform action implementation.
type UpgradeClusterGsCardImageAction struct {
	client *client.Client
}

// UpgradeClusterGsCardImageActionModel describes the action configuration shape.
type UpgradeClusterGsCardImageActionModel struct {
	ClusterSpecs            types.Dynamic `tfsdk:"cluster_specs" json:"clusterSpecs"`
	ImageServer             types.String  `tfsdk:"image_server" json:"imageServer"`
	SkipNotReachableDevices types.Bool    `tfsdk:"skip_not_reachable_devices" json:"skipNotReachableDevices"`
	TaskName                types.String  `tfsdk:"task_name" json:"taskName"`
}

// NewUpgradeClusterGsCardImageAction returns a new instance of the generated action.
func NewUpgradeClusterGsCardImageAction() action.Action {
	return &UpgradeClusterGsCardImageAction{}
}

// Metadata returns the action type name.
func (r *UpgradeClusterGsCardImageAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_upgrade_cluster_gs_card_image"
}

// Schema returns the action schema.
func (r *UpgradeClusterGsCardImageAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "upgrade the GS cards for selected clusters, it is an async only", Attributes: map[string]schema.Attribute{"cluster_specs": schema.DynamicAttribute{MarkdownDescription: "Cluster wise gs image spec", Required: true}, "image_server": schema.StringAttribute{MarkdownDescription: "alias of an image file server. it has to reference one of the existing image file server profiles", Required: true}, "skip_not_reachable_devices": schema.BoolAttribute{MarkdownDescription: "indicates whether the selected GigaSMART cards which belong to not-reachable nodes(if any) should be skipped and continue with image upgrade", Optional: true}, "task_name": schema.StringAttribute{MarkdownDescription: "user provided task name for the image upgrade", Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *UpgradeClusterGsCardImageAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UpgradeClusterGsCardImageActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *UpgradeClusterGsCardImageAction) invokeRemote(ctx context.Context, config *UpgradeClusterGsCardImageActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/clusterConfig/gsImageUpgrade"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_upgrade_cluster_gs_card_image", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_upgrade_cluster_gs_card_image", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_upgrade_cluster_gs_card_image", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_upgrade_cluster_gs_card_image", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_upgrade_cluster_gs_card_image", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_upgrade_cluster_gs_card_image", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_upgrade_cluster_gs_card_image", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_upgrade_cluster_gs_card_image", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_upgrade_cluster_gs_card_image", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_upgrade_cluster_gs_card_image", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_upgrade_cluster_gs_card_image", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_upgrade_cluster_gs_card_image", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_upgrade_cluster_gs_card_image", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *UpgradeClusterGsCardImageAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
