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
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*CreateEventNotificationConfigurationAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*CreateEventNotificationConfigurationAction)(nil)

// CreateEventNotificationConfigurationAction is the generated Terraform action implementation.
type CreateEventNotificationConfigurationAction struct {
	client *client.Client
}

// CreateEventNotificationConfigurationActionModel describes the action configuration shape.
type CreateEventNotificationConfigurationActionModel struct {
	AllowAttachment       types.Bool   `tfsdk:"allow_attachment" json:"allowAttachment"`
	AttachmentLimit       types.Int64  `tfsdk:"attachment_limit" json:"attachmentLimit"`
	Comment               types.String `tfsdk:"comment"`
	EmailSubjectPrefix    types.String `tfsdk:"email_subject_prefix" json:"emailSubjectPrefix"`
	Enabled               types.Bool   `tfsdk:"enabled"`
	EventDetails          types.List   `tfsdk:"event_details" json:"eventDetails"`
	ExternalTrapReceivers types.List   `tfsdk:"external_trap_receivers" json:"externalTrapReceivers"`
	InstantRateLimit      types.Int64  `tfsdk:"instant_rate_limit" json:"instantRateLimit"`
	NotifType             types.String `tfsdk:"notif_type"`
	Recipients            types.List   `tfsdk:"recipients"`
	RecurringSchedule     types.String `tfsdk:"recurring_schedule" json:"recurringSchedule"`
	SendMailIfEmpty       types.Bool   `tfsdk:"send_mail_if_empty" json:"sendMailIfEmpty"`
	Severity              types.List   `tfsdk:"severity"`
	Tags                  types.List   `tfsdk:"tags"`
	TaskId                types.String `tfsdk:"task_id" json:"taskId"`
	TaskName              types.String `tfsdk:"task_name" json:"taskName"`
	TemplateDetails       types.List   `tfsdk:"template_details" json:"templateDetails"`
	TimeInterval          types.Int64  `tfsdk:"time_interval" json:"timeInterval"`
	TimeLeft              types.String `tfsdk:"time_left" json:"timeLeft"`
	Type                  types.String `tfsdk:"type"`
}

// NewCreateEventNotificationConfigurationAction returns a new instance of the generated action.
func NewCreateEventNotificationConfigurationAction() action.Action {
	return &CreateEventNotificationConfigurationAction{}
}

// Metadata returns the action type name.
func (r *CreateEventNotificationConfigurationAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_create_event_notification_configuration"
}

// Schema returns the action schema.
func (r *CreateEventNotificationConfigurationAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Create Notification Configuration", Attributes: map[string]schema.Attribute{"allow_attachment": schema.BoolAttribute{MarkdownDescription: "Allow attachment in notification email", Optional: true}, "attachment_limit": schema.Int64Attribute{MarkdownDescription: "Maximum number of events to be included in the attachment", Optional: true}, "comment": schema.StringAttribute{MarkdownDescription: "Comments", Optional: true}, "email_subject_prefix": schema.StringAttribute{MarkdownDescription: "Subject of the notification email", Optional: true}, "enabled": schema.BoolAttribute{MarkdownDescription: "Status of the notification task", Optional: true}, "event_details": schema.ListNestedAttribute{MarkdownDescription: "Event Details", Required: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"description": schema.StringAttribute{MarkdownDescription: "Event Description", Optional: true}, "display_name": schema.StringAttribute{MarkdownDescription: "Event Display Name", Optional: true}, "event_type": schema.StringAttribute{MarkdownDescription: "Event Type", Optional: true}, "name": schema.StringAttribute{MarkdownDescription: "Event Name", Optional: true}, "scope": schema.StringAttribute{MarkdownDescription: "Event Scope", Optional: true}, "severity": schema.ListAttribute{MarkdownDescription: "Event Severity", Optional: true, ElementType: types.StringType}, "severity_type": schema.StringAttribute{MarkdownDescription: "Event Severity Type", Optional: true}, "sub_type": schema.StringAttribute{MarkdownDescription: "Event Subtype", Optional: true}}}}, "external_trap_receivers": schema.ListAttribute{MarkdownDescription: "List of External Trap Receiver aliases", Optional: true, ElementType: types.StringType}, "instant_rate_limit": schema.Int64Attribute{MarkdownDescription: "Maximum number of instant emails that can be sent per minute", Optional: true}, "notif_type": schema.StringAttribute{MarkdownDescription: "Type of Notification", Required: true}, "recipients": schema.ListAttribute{MarkdownDescription: "List of email recipients", Optional: true, ElementType: types.StringType}, "recurring_schedule": schema.StringAttribute{MarkdownDescription: "Cron expression", Optional: true}, "send_mail_if_empty": schema.BoolAttribute{MarkdownDescription: "Send email if no events are generated within time interval", Optional: true}, "severity": schema.ListAttribute{MarkdownDescription: "Event Severities", Optional: true, ElementType: types.StringType}, "tags": schema.ListNestedAttribute{MarkdownDescription: "Tags", Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"tag_key": schema.StringAttribute{MarkdownDescription: "Name of the tag", Required: true}, "tag_values": schema.ListAttribute{MarkdownDescription: "All possible values of the tag", Required: true, ElementType: types.StringType}}}}, "task_id": schema.StringAttribute{MarkdownDescription: "ID of the notification task", Optional: true}, "task_name": schema.StringAttribute{MarkdownDescription: "Name of the task", Required: true}, "template_details": schema.ListAttribute{MarkdownDescription: "List of templates used", Optional: true, ElementType: types.StringType}, "time_interval": schema.Int64Attribute{MarkdownDescription: "Time interval between two batch emails", Optional: true}, "time_left": schema.StringAttribute{MarkdownDescription: "Time left for the next batch task to execute", Optional: true}, "type": schema.StringAttribute{MarkdownDescription: "Type of the event notification task", Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *CreateEventNotificationConfigurationAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config CreateEventNotificationConfigurationActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *CreateEventNotificationConfigurationAction) invokeRemote(ctx context.Context, config *CreateEventNotificationConfigurationActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/notification/event/notifMetaConfig/{notifType}"
	reqPath = strings.ReplaceAll(reqPath, "{notifType}", url.PathEscape(config.NotifType.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_create_event_notification_configuration", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_create_event_notification_configuration", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_create_event_notification_configuration", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_create_event_notification_configuration", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_create_event_notification_configuration", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_create_event_notification_configuration", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_create_event_notification_configuration", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_create_event_notification_configuration", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_create_event_notification_configuration", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_create_event_notification_configuration", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_create_event_notification_configuration", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_create_event_notification_configuration", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_create_event_notification_configuration", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *CreateEventNotificationConfigurationAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
