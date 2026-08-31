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
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*CollectHeartbeatfromGigaInsightNodeAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*CollectHeartbeatfromGigaInsightNodeAction)(nil)

// CollectHeartbeatfromGigaInsightNodeAction is the generated Terraform action implementation.
type CollectHeartbeatfromGigaInsightNodeAction struct {
	client *client.Client
}

// CollectHeartbeatfromGigaInsightNodeActionModel describes the action configuration shape.
type CollectHeartbeatfromGigaInsightNodeActionModel struct {
	BodyNodeId   types.String `tfsdk:"body_node_id" json:"nodeId"`
	NodeId       types.String `tfsdk:"node_id"`
	NodeVersion  types.String `tfsdk:"node_version" json:"nodeVersion"`
	PromptBundle types.Object `tfsdk:"prompt_bundle" json:"promptBundle"`
}

// NewCollectHeartbeatfromGigaInsightNodeAction returns a new instance of the generated action.
func NewCollectHeartbeatfromGigaInsightNodeAction() action.Action {
	return &CollectHeartbeatfromGigaInsightNodeAction{}
}

// Metadata returns the action type name.
func (r *CollectHeartbeatfromGigaInsightNodeAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_collect_heartbeatfrom_giga_insight_node"
}

// Schema returns the action schema.
func (r *CollectHeartbeatfromGigaInsightNodeAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Accept Heartbeat from Gigamon Insights in FM", Attributes: map[string]schema.Attribute{"body_node_id": schema.StringAttribute{MarkdownDescription: "The node identifier", Optional: true}, "node_id": schema.StringAttribute{MarkdownDescription: "The unique identifier of the GigaInsight Node", Required: true}, "node_version": schema.StringAttribute{MarkdownDescription: "The GigaInsight node version", Optional: true}, "prompt_bundle": schema.SingleNestedAttribute{MarkdownDescription: "The prompt bundle information installed on the node", Optional: true, Attributes: map[string]schema.Attribute{"prompt_bundle_last_updated_at": schema.StringAttribute{MarkdownDescription: "The date and time when the prompt bundle was last updated on the node", Optional: true}, "prompt_bundle_version": schema.StringAttribute{MarkdownDescription: "The prompt bundle version installed on the node", Optional: true}}}}}
}

// Invoke executes the action against the remote API.
func (r *CollectHeartbeatfromGigaInsightNodeAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config CollectHeartbeatfromGigaInsightNodeActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *CollectHeartbeatfromGigaInsightNodeAction) invokeRemote(ctx context.Context, config *CollectHeartbeatfromGigaInsightNodeActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/orchestrate/insights/heartbeat/{nodeId}"
	reqPath = strings.ReplaceAll(reqPath, "{nodeId}", url.PathEscape(config.NodeId.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_collect_heartbeatfrom_giga_insight_node", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_collect_heartbeatfrom_giga_insight_node", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPatch, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_collect_heartbeatfrom_giga_insight_node", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_collect_heartbeatfrom_giga_insight_node", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_collect_heartbeatfrom_giga_insight_node", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_collect_heartbeatfrom_giga_insight_node", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_collect_heartbeatfrom_giga_insight_node", "Access Denied. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_collect_heartbeatfrom_giga_insight_node", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_collect_heartbeatfrom_giga_insight_node", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_collect_heartbeatfrom_giga_insight_node", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_collect_heartbeatfrom_giga_insight_node", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *CollectHeartbeatfromGigaInsightNodeAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
