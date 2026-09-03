package provider

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*ClearPtpPortsCountersByPortIdAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*ClearPtpPortsCountersByPortIdAction)(nil)

// ClearPtpPortsCountersByPortIdAction is the generated Terraform action implementation.
type ClearPtpPortsCountersByPortIdAction struct {
	client *client.Client
}

// ClearPtpPortsCountersByPortIdActionModel describes the action configuration shape.
type ClearPtpPortsCountersByPortIdActionModel struct {
	ClusterId types.String `tfsdk:"cluster_id"`
	PortId    types.String `tfsdk:"port_id"`
}

// NewClearPtpPortsCountersByPortIdAction returns a new instance of the generated action.
func NewClearPtpPortsCountersByPortIdAction() action.Action {
	return &ClearPtpPortsCountersByPortIdAction{}
}

// Metadata returns the action type name.
func (r *ClearPtpPortsCountersByPortIdAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_clear_ptp_ports_counters_by_port_id"
}

// Schema returns the action schema.
func (r *ClearPtpPortsCountersByPortIdAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "clear PTP port counters by port ID", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "port_id": schema.StringAttribute{MarkdownDescription: "id of the target device Port (format: boxId_slotId_port, example: 1_1_c1)", Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *ClearPtpPortsCountersByPortIdAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config ClearPtpPortsCountersByPortIdActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *ClearPtpPortsCountersByPortIdAction) invokeRemote(ctx context.Context, config *ClearPtpPortsCountersByPortIdActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/nodeCounters/ptpPorts/{portId}"
	reqPath = strings.ReplaceAll(reqPath, "{portId}", url.PathEscape(config.PortId.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_clear_ptp_ports_counters_by_port_id", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_clear_ptp_ports_counters_by_port_id", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 204) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_clear_ptp_ports_counters_by_port_id", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_clear_ptp_ports_counters_by_port_id", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_clear_ptp_ports_counters_by_port_id", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_clear_ptp_ports_counters_by_port_id", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_clear_ptp_ports_counters_by_port_id", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_clear_ptp_ports_counters_by_port_id", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_clear_ptp_ports_counters_by_port_id", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_clear_ptp_ports_counters_by_port_id", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_clear_ptp_ports_counters_by_port_id", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *ClearPtpPortsCountersByPortIdAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
