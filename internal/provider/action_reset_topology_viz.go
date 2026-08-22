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
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*ResetTopologyVizAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*ResetTopologyVizAction)(nil)

// ResetTopologyVizAction is the generated Terraform action implementation.
type ResetTopologyVizAction struct {
	client *client.Client
}

// ResetTopologyVizActionModel describes the action configuration shape.
type ResetTopologyVizActionModel struct {
}

// NewResetTopologyVizAction returns a new instance of the generated action.
func NewResetTopologyVizAction() action.Action {
	return &ResetTopologyVizAction{}
}

// Metadata returns the action type name.
func (r *ResetTopologyVizAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_reset_topology_viz"
}

// Schema returns the action schema.
func (r *ResetTopologyVizAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "resets topology data by reconstructing links and node neighbour details"}
}

// Invoke executes the action against the remote API.
func (r *ResetTopologyVizAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config ResetTopologyVizActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *ResetTopologyVizAction) invokeRemote(ctx context.Context, config *ResetTopologyVizActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/topoviz/links/reset"
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_reset_topology_viz", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_reset_topology_viz", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		apiErr, err := client.NewAPIError(httpResp)
		if err != nil {
			resp.Diagnostics.AddError("Error invoking gigavuecore_reset_topology_viz", fmt.Sprintf("Could not read error response: %s", err))
			return
		}
		resp.Diagnostics.AddError("Error invoking gigavuecore_reset_topology_viz", apiErr.Error())
		return
	}
}

// Configure stores the API client supplied by the provider.
func (r *ResetTopologyVizAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
