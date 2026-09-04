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
var _ action.Action = (*DisableBatchEventNotificationConfigurationAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*DisableBatchEventNotificationConfigurationAction)(nil)

// DisableBatchEventNotificationConfigurationAction is the generated Terraform action implementation.
type DisableBatchEventNotificationConfigurationAction struct {
	client *client.Client
}

// DisableBatchEventNotificationConfigurationActionModel describes the action configuration shape.
type DisableBatchEventNotificationConfigurationActionModel struct {
	TaskId types.String `tfsdk:"task_id"`
}

// NewDisableBatchEventNotificationConfigurationAction returns a new instance of the generated action.
func NewDisableBatchEventNotificationConfigurationAction() action.Action {
	return &DisableBatchEventNotificationConfigurationAction{}
}

// Metadata returns the action type name.
func (r *DisableBatchEventNotificationConfigurationAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_disable_batch_event_notification_configuration"
}

// Schema returns the action schema.
func (r *DisableBatchEventNotificationConfigurationAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Disable Batch Notification Configuration", Attributes: map[string]schema.Attribute{"task_id": schema.StringAttribute{MarkdownDescription: "ID of the Task", Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *DisableBatchEventNotificationConfigurationAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config DisableBatchEventNotificationConfigurationActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *DisableBatchEventNotificationConfigurationAction) invokeRemote(ctx context.Context, config *DisableBatchEventNotificationConfigurationActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/notification/event/notifMetaConfig/batch/disable/{taskId}"
	reqPath = strings.ReplaceAll(reqPath, "{taskId}", url.PathEscape(config.TaskId.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_disable_batch_event_notification_configuration", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_disable_batch_event_notification_configuration", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_disable_batch_event_notification_configuration", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_disable_batch_event_notification_configuration", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_disable_batch_event_notification_configuration", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_disable_batch_event_notification_configuration", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_disable_batch_event_notification_configuration", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_disable_batch_event_notification_configuration", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_disable_batch_event_notification_configuration", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_disable_batch_event_notification_configuration", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_disable_batch_event_notification_configuration", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *DisableBatchEventNotificationConfigurationAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
