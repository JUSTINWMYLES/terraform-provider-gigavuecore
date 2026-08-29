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
var _ action.Action = (*ClearIpv6NeighborEntriesAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*ClearIpv6NeighborEntriesAction)(nil)

// ClearIpv6NeighborEntriesAction is the generated Terraform action implementation.
type ClearIpv6NeighborEntriesAction struct {
	client *client.Client
}

// ClearIpv6NeighborEntriesActionModel describes the action configuration shape.
type ClearIpv6NeighborEntriesActionModel struct {
}

// NewClearIpv6NeighborEntriesAction returns a new instance of the generated action.
func NewClearIpv6NeighborEntriesAction() action.Action {
	return &ClearIpv6NeighborEntriesAction{}
}

// Metadata returns the action type name.
func (r *ClearIpv6NeighborEntriesAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_clear_ipv6_neighbor_entries"
}

// Schema returns the action schema.
func (r *ClearIpv6NeighborEntriesAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Clear Ipv6 Neighbor Entries"}
}

// Invoke executes the action against the remote API.
func (r *ClearIpv6NeighborEntriesAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config ClearIpv6NeighborEntriesActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *ClearIpv6NeighborEntriesAction) invokeRemote(ctx context.Context, config *ClearIpv6NeighborEntriesActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/system/ipv6/neighbors"
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_clear_ipv6_neighbor_entries", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_clear_ipv6_neighbor_entries", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 204) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_clear_ipv6_neighbor_entries", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_clear_ipv6_neighbor_entries", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_clear_ipv6_neighbor_entries", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_clear_ipv6_neighbor_entries", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_clear_ipv6_neighbor_entries", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_clear_ipv6_neighbor_entries", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_clear_ipv6_neighbor_entries", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_clear_ipv6_neighbor_entries", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_clear_ipv6_neighbor_entries", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *ClearIpv6NeighborEntriesAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
