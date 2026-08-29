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
var _ action.Action = (*RedefineSnmpThrottleConfigAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*RedefineSnmpThrottleConfigAction)(nil)

// RedefineSnmpThrottleConfigAction is the generated Terraform action implementation.
type RedefineSnmpThrottleConfigAction struct {
	client *client.Client
}

// RedefineSnmpThrottleConfigActionModel describes the action configuration shape.
type RedefineSnmpThrottleConfigActionModel struct {
	ThrottleConfigDetails types.List `tfsdk:"throttle_config_details" json:"throttleConfigDetails"`
}

// NewRedefineSnmpThrottleConfigAction returns a new instance of the generated action.
func NewRedefineSnmpThrottleConfigAction() action.Action {
	return &RedefineSnmpThrottleConfigAction{}
}

// Metadata returns the action type name.
func (r *RedefineSnmpThrottleConfigAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_redefine_snmp_throttle_config"
}

// Schema returns the action schema.
func (r *RedefineSnmpThrottleConfigAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Redefine Snmp Throttle config", Attributes: map[string]schema.Attribute{"throttle_config_details": schema.ListNestedAttribute{MarkdownDescription: "list of SNMP Throttle Config Details on the node", Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"interval": schema.Int64Attribute{MarkdownDescription: "Time interval at which the throttle should occur (in seconds)", Optional: true}, "notify_set": schema.StringAttribute{MarkdownDescription: "When 'notifySet' is 'select', throttleEvents represents the subset of event types for SNMP throttling. When set to 'none', effectively disables SNMP throttle from the node.", Optional: true}, "report_threshold": schema.Int64Attribute{MarkdownDescription: "Minimum count threshold to send the throttle report", Optional: true}, "throttle_events": schema.SetAttribute{MarkdownDescription: "The set of notification event types", Optional: true, ElementType: types.StringType}}}}}}
}

// Invoke executes the action against the remote API.
func (r *RedefineSnmpThrottleConfigAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config RedefineSnmpThrottleConfigActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *RedefineSnmpThrottleConfigAction) invokeRemote(ctx context.Context, config *RedefineSnmpThrottleConfigActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/snmp/throttle"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_snmp_throttle_config", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_snmp_throttle_config", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_snmp_throttle_config", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_snmp_throttle_config", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_snmp_throttle_config", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_snmp_throttle_config", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_snmp_throttle_config", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_snmp_throttle_config", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_snmp_throttle_config", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_snmp_throttle_config", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_snmp_throttle_config", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_snmp_throttle_config", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_snmp_throttle_config", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *RedefineSnmpThrottleConfigAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
