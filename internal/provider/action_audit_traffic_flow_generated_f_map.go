package provider

import (
	"context"
	"fmt"
	"net/http"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*AuditTrafficFlowGeneratedFMapAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*AuditTrafficFlowGeneratedFMapAction)(nil)

// AuditTrafficFlowGeneratedFMapAction is the generated Terraform action implementation.
type AuditTrafficFlowGeneratedFMapAction struct {
	client *client.Client
}

// AuditTrafficFlowGeneratedFMapActionModel describes the action configuration shape.
type AuditTrafficFlowGeneratedFMapActionModel struct {
}

// NewAuditTrafficFlowGeneratedFMapAction returns a new instance of the generated action.
func NewAuditTrafficFlowGeneratedFMapAction() action.Action {
	return &AuditTrafficFlowGeneratedFMapAction{}
}

// Metadata returns the action type name.
func (r *AuditTrafficFlowGeneratedFMapAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_audit_traffic_flow_generated_f_map"
}

// Schema returns the action schema.
func (r *AuditTrafficFlowGeneratedFMapAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Check the integrity of traffic flow configurations"}
}

// Invoke executes the action against the remote API.
func (r *AuditTrafficFlowGeneratedFMapAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config AuditTrafficFlowGeneratedFMapActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *AuditTrafficFlowGeneratedFMapAction) invokeRemote(ctx context.Context, config *AuditTrafficFlowGeneratedFMapActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/traffic-flows/audit"
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_audit_traffic_flow_generated_f_map", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_audit_traffic_flow_generated_f_map", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 202) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_audit_traffic_flow_generated_f_map", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_audit_traffic_flow_generated_f_map", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_audit_traffic_flow_generated_f_map", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_audit_traffic_flow_generated_f_map", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_audit_traffic_flow_generated_f_map", "Entity Already Exists. See errors payload for details")
			return
		case 417:
			resp.Diagnostics.AddError("Error invoking gigavuecore_audit_traffic_flow_generated_f_map", "Unable to check the config integrity successfully. See the logs for more info")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_audit_traffic_flow_generated_f_map", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_audit_traffic_flow_generated_f_map", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_audit_traffic_flow_generated_f_map", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_audit_traffic_flow_generated_f_map", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *AuditTrafficFlowGeneratedFMapAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
