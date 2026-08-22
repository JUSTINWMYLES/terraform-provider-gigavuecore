package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*GetAllEventNotificationConfigurationDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAllEventNotificationConfigurationDataSource)(nil)
)

// GetAllEventNotificationConfigurationDataSource is the generated Terraform data source implementation.
type GetAllEventNotificationConfigurationDataSource struct {
	client *client.Client
}

// GetAllEventNotificationConfigurationDataSourceModel describes the data source state shape.
type GetAllEventNotificationConfigurationDataSourceModel struct {
	EventDetails          types.String `tfsdk:"event_details" json:"eventDetails"`
	EventNotifTasks       types.List   `tfsdk:"event_notif_tasks" json:"eventNotifTasks"`
	ExternalTrapReceivers types.String `tfsdk:"external_trap_receivers" json:"externalTrapReceivers"`
	InstantRateLimit      types.Int64  `tfsdk:"instant_rate_limit" json:"instantRateLimit"`
	Recipients            types.String `tfsdk:"recipients"`
	Tags                  types.String `tfsdk:"tags"`
	TaskName              types.String `tfsdk:"task_name" json:"taskName"`
	TimeInterval          types.String `tfsdk:"time_interval" json:"timeInterval"`
	Type                  types.String `tfsdk:"type"`
}

// NewGetAllEventNotificationConfigurationDataSource returns a new instance of the generated data source.
func NewGetAllEventNotificationConfigurationDataSource() datasource.DataSource {
	return &GetAllEventNotificationConfigurationDataSource{}
}

// Metadata returns the data source type name.
func (d *GetAllEventNotificationConfigurationDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_event_notification_configuration"
}

// Schema returns the data source schema.
func (d *GetAllEventNotificationConfigurationDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get All Event Notification Configurations", Attributes: map[string]schema.Attribute{"event_details": schema.StringAttribute{MarkdownDescription: "eventDetails", Optional: true}, "event_notif_tasks": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"allow_attachment": schema.BoolAttribute{MarkdownDescription: "Allow attachment in notification email", Computed: true}, "attachment_limit": schema.Int64Attribute{MarkdownDescription: "Maximum number of events to be included in the attachment", Computed: true}, "comment": schema.StringAttribute{MarkdownDescription: "Comments", Computed: true}, "email_subject_prefix": schema.StringAttribute{MarkdownDescription: "Subject of the notification email", Computed: true}, "enabled": schema.BoolAttribute{MarkdownDescription: "Status of the notification task", Computed: true}, "event_details": schema.ListNestedAttribute{MarkdownDescription: "Event Details", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"description": schema.StringAttribute{MarkdownDescription: "Event Description", Computed: true}, "display_name": schema.StringAttribute{MarkdownDescription: "Event Display Name", Computed: true}, "event_type": schema.StringAttribute{MarkdownDescription: "Event Type", Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "Event Name", Computed: true}, "scope": schema.StringAttribute{MarkdownDescription: "Event Scope", Computed: true}, "severity": schema.ListAttribute{MarkdownDescription: "Event Severity", Computed: true, ElementType: types.StringType}, "severity_type": schema.StringAttribute{MarkdownDescription: "Event Severity Type", Computed: true}, "sub_type": schema.StringAttribute{MarkdownDescription: "Event Subtype", Computed: true}}}}, "external_trap_receivers": schema.ListAttribute{MarkdownDescription: "List of External Trap Receiver aliases", Computed: true, ElementType: types.StringType}, "instant_rate_limit": schema.Int64Attribute{MarkdownDescription: "Maximum number of instant emails that can be sent per minute", Computed: true}, "recipients": schema.ListAttribute{MarkdownDescription: "List of email recipients", Computed: true, ElementType: types.StringType}, "recurring_schedule": schema.StringAttribute{MarkdownDescription: "Cron expression", Computed: true}, "send_mail_if_empty": schema.BoolAttribute{MarkdownDescription: "Send email if no events are generated within time interval", Computed: true}, "severity": schema.ListAttribute{MarkdownDescription: "Event Severities", Computed: true, ElementType: types.StringType}, "tags": schema.ListNestedAttribute{MarkdownDescription: "Tags", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"tag_key": schema.StringAttribute{MarkdownDescription: "Name of the tag", Computed: true}, "tag_values": schema.ListAttribute{MarkdownDescription: "All possible values of the tag", Computed: true, ElementType: types.StringType}}}}, "task_id": schema.StringAttribute{MarkdownDescription: "ID of the notification task", Computed: true}, "task_name": schema.StringAttribute{MarkdownDescription: "Name of the task", Computed: true}, "template_details": schema.ListAttribute{MarkdownDescription: "List of templates used", Computed: true, ElementType: types.StringType}, "time_interval": schema.Int64Attribute{MarkdownDescription: "Time interval between two batch emails", Computed: true}, "time_left": schema.StringAttribute{MarkdownDescription: "Time left for the next batch task to execute", Computed: true}, "type": schema.StringAttribute{MarkdownDescription: "Type of the event notification task", Computed: true}}}}, "external_trap_receivers": schema.StringAttribute{MarkdownDescription: "Alias of External Trap Receivers", Optional: true}, "instant_rate_limit": schema.Int64Attribute{MarkdownDescription: "Maximum number of instant emails that can be sent per minute", Computed: true}, "recipients": schema.StringAttribute{MarkdownDescription: "Email Recipeints", Optional: true}, "tags": schema.StringAttribute{MarkdownDescription: "tags", Optional: true}, "task_name": schema.StringAttribute{MarkdownDescription: "Name of the task", Optional: true}, "time_interval": schema.StringAttribute{MarkdownDescription: "timeInterval", Optional: true}, "type": schema.StringAttribute{MarkdownDescription: "Type of the task", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetAllEventNotificationConfigurationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAllEventNotificationConfigurationDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	d.readRemote(ctx, &config, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}

// readRemote performs the read HTTP exchange and decodes the response into config. Extracted from Read so the request/response logic is unit-testable without a tfsdk.Config.
func (d *GetAllEventNotificationConfigurationDataSource) readRemote(ctx context.Context, config *GetAllEventNotificationConfigurationDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/notification/event/notifMetaConfig"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_event_notification_configuration", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.TaskName.IsNull() {
		query.Set("taskName", config.TaskName.ValueString())
	}
	if !config.Type.IsNull() {
		query.Set("type", config.Type.ValueString())
	}
	if !config.EventDetails.IsNull() {
		query.Set("eventDetails", config.EventDetails.ValueString())
	}
	if !config.Recipients.IsNull() {
		query.Set("recipients", config.Recipients.ValueString())
	}
	if !config.ExternalTrapReceivers.IsNull() {
		query.Set("externalTrapReceivers", config.ExternalTrapReceivers.ValueString())
	}
	if !config.TimeInterval.IsNull() {
		query.Set("timeInterval", config.TimeInterval.ValueString())
	}
	if !config.Tags.IsNull() {
		query.Set("tags", config.Tags.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_event_notification_configuration", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_event_notification_configuration", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_event_notification_configuration", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_event_notification_configuration", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_event_notification_configuration", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_event_notification_configuration", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_event_notification_configuration", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_event_notification_configuration", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_event_notification_configuration", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_all_event_notification_configuration", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_event_notification_configuration", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_event_notification_configuration", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_event_notification_configuration", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetAllEventNotificationConfigurationDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Data Source Configure Type", fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData))
		return
	}
	d.client = c
}
