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
var _ action.Action = (*UnacknowledgeMultipleAlarmsAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*UnacknowledgeMultipleAlarmsAction)(nil)

// UnacknowledgeMultipleAlarmsAction is the generated Terraform action implementation.
type UnacknowledgeMultipleAlarmsAction struct {
	client *client.Client
}

// UnacknowledgeMultipleAlarmsActionModel describes the action configuration shape.
type UnacknowledgeMultipleAlarmsActionModel struct {
	AlarmIds types.List   `tfsdk:"alarm_ids" json:"alarmIds"`
	Comment  types.String `tfsdk:"comment"`
}

// NewUnacknowledgeMultipleAlarmsAction returns a new instance of the generated action.
func NewUnacknowledgeMultipleAlarmsAction() action.Action {
	return &UnacknowledgeMultipleAlarmsAction{}
}

// Metadata returns the action type name.
func (r *UnacknowledgeMultipleAlarmsAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_unacknowledge_multiple_alarms"
}

// Schema returns the action schema.
func (r *UnacknowledgeMultipleAlarmsAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Unacknowledge Multiple Alarms", Attributes: map[string]schema.Attribute{"alarm_ids": schema.ListAttribute{MarkdownDescription: "IDs of alarms that needs to be unacknowledged", Required: true, ElementType: types.StringType}, "comment": schema.StringAttribute{MarkdownDescription: "Alarm comment", Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *UnacknowledgeMultipleAlarmsAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UnacknowledgeMultipleAlarmsActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *UnacknowledgeMultipleAlarmsAction) invokeRemote(ctx context.Context, config *UnacknowledgeMultipleAlarmsActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/alarms/bulk/unacknowledge"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_unacknowledge_multiple_alarms", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_unacknowledge_multiple_alarms", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_unacknowledge_multiple_alarms", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_unacknowledge_multiple_alarms", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_unacknowledge_multiple_alarms", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_unacknowledge_multiple_alarms", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_unacknowledge_multiple_alarms", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_unacknowledge_multiple_alarms", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_unacknowledge_multiple_alarms", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_unacknowledge_multiple_alarms", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_unacknowledge_multiple_alarms", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_unacknowledge_multiple_alarms", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_unacknowledge_multiple_alarms", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *UnacknowledgeMultipleAlarmsAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
