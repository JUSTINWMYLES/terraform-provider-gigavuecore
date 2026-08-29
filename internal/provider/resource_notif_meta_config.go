package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	timeouts "github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertions.
var (
	_ resource.Resource              = (*NotifMetaConfigResource)(nil)
	_ resource.ResourceWithConfigure = (*NotifMetaConfigResource)(nil)
)

// NotifMetaConfigResource is the generated Terraform managed resource implementation.
type NotifMetaConfigResource struct {
	client *client.Client
}

// NotifMetaConfigResourceModel describes the Terraform state and plan shape for NotifMetaConfigResource.
type NotifMetaConfigResourceModel struct {
	AllowAttachment       types.Bool     `tfsdk:"allow_attachment" json:"allowAttachment"`
	AttachmentLimit       types.Int64    `tfsdk:"attachment_limit" json:"attachmentLimit"`
	Comment               types.String   `tfsdk:"comment"`
	EmailSubjectPrefix    types.String   `tfsdk:"email_subject_prefix" json:"emailSubjectPrefix"`
	Enabled               types.Bool     `tfsdk:"enabled"`
	EventDetails          types.List     `tfsdk:"event_details" json:"eventDetails"`
	ExternalTrapReceivers types.List     `tfsdk:"external_trap_receivers" json:"externalTrapReceivers"`
	Id                    types.String   `tfsdk:"id"`
	InstantRateLimit      types.Int64    `tfsdk:"instant_rate_limit" json:"instantRateLimit"`
	Recipients            types.List     `tfsdk:"recipients"`
	RecurringSchedule     types.String   `tfsdk:"recurring_schedule" json:"recurringSchedule"`
	SendMailIfEmpty       types.Bool     `tfsdk:"send_mail_if_empty" json:"sendMailIfEmpty"`
	Severity              types.List     `tfsdk:"severity"`
	Tags                  types.List     `tfsdk:"tags"`
	TaskId                types.String   `tfsdk:"task_id" json:"taskId"`
	TaskName              types.String   `tfsdk:"task_name" json:"taskName"`
	TemplateDetails       types.List     `tfsdk:"template_details" json:"templateDetails"`
	TimeInterval          types.Int64    `tfsdk:"time_interval" json:"timeInterval"`
	TimeLeft              types.String   `tfsdk:"time_left" json:"timeLeft"`
	Type                  types.String   `tfsdk:"type"`
	Timeouts              timeouts.Value `tfsdk:"timeouts"`
}

// Metadata returns the resource type name.
func (r *NotifMetaConfigResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_notif_meta_config"
}

// Schema returns the Terraform schema for this resource.
func (r *NotifMetaConfigResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get Notification Configuration", Attributes: map[string]schema.Attribute{"allow_attachment": schema.BoolAttribute{MarkdownDescription: "Allow attachment in notification email", Optional: true, Computed: true}, "attachment_limit": schema.Int64Attribute{MarkdownDescription: "Maximum number of events to be included in the attachment", Optional: true, Computed: true}, "comment": schema.StringAttribute{MarkdownDescription: "Comments", Optional: true, Computed: true}, "email_subject_prefix": schema.StringAttribute{MarkdownDescription: "Subject of the notification email", Optional: true, Computed: true}, "enabled": schema.BoolAttribute{MarkdownDescription: "Status of the notification task", Optional: true, Computed: true}, "event_details": schema.ListNestedAttribute{MarkdownDescription: "Event Details", Required: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"description": schema.StringAttribute{MarkdownDescription: "Event Description", Optional: true, Computed: true}, "display_name": schema.StringAttribute{MarkdownDescription: "Event Display Name", Optional: true, Computed: true}, "event_type": schema.StringAttribute{MarkdownDescription: "Event Type", Optional: true, Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "Event Name", Optional: true, Computed: true}, "scope": schema.StringAttribute{MarkdownDescription: "Event Scope", Optional: true, Computed: true}, "severity": schema.ListAttribute{MarkdownDescription: "Event Severity", Optional: true, Computed: true, ElementType: types.StringType}, "severity_type": schema.StringAttribute{MarkdownDescription: "Event Severity Type", Optional: true, Computed: true}, "sub_type": schema.StringAttribute{MarkdownDescription: "Event Subtype", Optional: true, Computed: true}}}}, "external_trap_receivers": schema.ListAttribute{MarkdownDescription: "List of External Trap Receiver aliases", Optional: true, Computed: true, ElementType: types.StringType}, "id": schema.StringAttribute{Computed: true}, "instant_rate_limit": schema.Int64Attribute{MarkdownDescription: "Maximum number of instant emails that can be sent per minute", Optional: true, Computed: true}, "recipients": schema.ListAttribute{MarkdownDescription: "List of email recipients", Optional: true, Computed: true, ElementType: types.StringType}, "recurring_schedule": schema.StringAttribute{MarkdownDescription: "Cron expression", Optional: true, Computed: true}, "send_mail_if_empty": schema.BoolAttribute{MarkdownDescription: "Send email if no events are generated within time interval", Optional: true, Computed: true}, "severity": schema.ListAttribute{MarkdownDescription: "Event Severities", Optional: true, Computed: true, ElementType: types.StringType}, "tags": schema.ListNestedAttribute{MarkdownDescription: "Tags", Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"tag_key": schema.StringAttribute{MarkdownDescription: "Name of the tag", Required: true}, "tag_values": schema.ListAttribute{MarkdownDescription: "All possible values of the tag", Required: true, ElementType: types.StringType}}}}, "task_id": schema.StringAttribute{MarkdownDescription: "ID of the notification task", Required: true}, "task_name": schema.StringAttribute{MarkdownDescription: "Name of the task", Required: true}, "template_details": schema.ListAttribute{MarkdownDescription: "List of templates used", Optional: true, Computed: true, ElementType: types.StringType}, "time_interval": schema.Int64Attribute{MarkdownDescription: "Time interval between two batch emails", Optional: true, Computed: true}, "time_left": schema.StringAttribute{MarkdownDescription: "Time left for the next batch task to execute", Optional: true, Computed: true}, "type": schema.StringAttribute{MarkdownDescription: "Type of the event notification task", Required: true}}, Blocks: map[string]schema.Block{"timeouts": timeouts.Block(ctx, timeouts.Opts{Create: true, Read: true, Update: true, Delete: true})}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *NotifMetaConfigResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan NotifMetaConfigResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	timeout, diags := plan.Timeouts.Create(ctx, 20*time.Minute)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	r.createRemote(ctx, &plan, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// createRemote performs the create HTTP exchange and decodes the response into plan. Extracted from Create so the request/response logic is unit-testable without a tfsdk.Plan.
func (r *NotifMetaConfigResource) createRemote(ctx context.Context, plan *NotifMetaConfigResourceModel, resp *resource.CreateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_notif_meta_config", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_notif_meta_config", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/notification/event/notifMetaConfig/{notifType}/{taskId}"
	reqPath = strings.ReplaceAll(reqPath, "{notifType}", url.PathEscape("instant"))
	reqPath = strings.ReplaceAll(reqPath, "{taskId}", url.PathEscape(plan.Id.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_notif_meta_config", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_notif_meta_config", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error creating gigavuecore_notif_meta_config", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error creating gigavuecore_notif_meta_config", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error creating gigavuecore_notif_meta_config", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error creating gigavuecore_notif_meta_config", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error creating gigavuecore_notif_meta_config", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error creating gigavuecore_notif_meta_config", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error creating gigavuecore_notif_meta_config", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error creating gigavuecore_notif_meta_config", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error creating gigavuecore_notif_meta_config", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error creating gigavuecore_notif_meta_config", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error creating gigavuecore_notif_meta_config", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
	if plan.Id.IsNull() || plan.Id.IsUnknown() {
		loc := httpResp.Header.Get("Location")
		if loc != "" {
			loc = strings.TrimRight(loc, "/")
			i := strings.LastIndex(loc, "/")
			if i >= 0 {
				loc = loc[i+1:]
			}
			plan.Id = types.StringValue(loc)
		} else {
			resp.Diagnostics.AddError("Error creating gigavuecore_notif_meta_config", "The create response did not contain an identifier and no Location header was returned, so the resource cannot be tracked in state.")
			return
		}
	}
}

// Read refreshes the Terraform state with the latest remote values.
func (r *NotifMetaConfigResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state NotifMetaConfigResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	timeout, diags := state.Timeouts.Read(ctx, 10*time.Minute)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	if r.readRemote(ctx, &state, resp) {
		resp.State.RemoveResource(ctx)
		return
	}
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// readRemote performs the read HTTP exchange and decodes the response into state, returning removed=true when the API reports 404. Extracted from Read so the request/response logic is unit-testable without a tfsdk.State.
func (r *NotifMetaConfigResource) readRemote(ctx context.Context, state *NotifMetaConfigResourceModel, resp *resource.ReadResponse) (removed bool) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/notification/event/notifMetaConfig/{notifType}/{taskId}"
	reqPath = strings.ReplaceAll(reqPath, "{notifType}", url.PathEscape("instant"))
	reqPath = strings.ReplaceAll(reqPath, "{taskId}", url.PathEscape(state.Id.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_notif_meta_config", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_notif_meta_config", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		removed = true
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_notif_meta_config", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_notif_meta_config", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_notif_meta_config", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_notif_meta_config", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_notif_meta_config", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_notif_meta_config", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_notif_meta_config", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_notif_meta_config", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_notif_meta_config", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_notif_meta_config", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&state, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_notif_meta_config", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
	return
}

// Update modifies the remote resource to match the desired plan.
func (r *NotifMetaConfigResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan NotifMetaConfigResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state NotifMetaConfigResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	timeout, diags := plan.Timeouts.Update(ctx, 20*time.Minute)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	if plan.Id.IsNull() || plan.Id.IsUnknown() {
		if !state.Id.IsNull() && !state.Id.IsUnknown() {
			plan.Id = state.Id
		}
	}
	preserveStateIntoPlan(&plan, &state)
	r.updateRemote(ctx, &plan, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// updateRemote performs the update HTTP exchange and decodes the response into plan. Extracted from Update so the request/response logic is unit-testable without a tfsdk.Plan.
func (r *NotifMetaConfigResource) updateRemote(ctx context.Context, plan *NotifMetaConfigResourceModel, resp *resource.UpdateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_notif_meta_config", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_notif_meta_config", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/notification/event/notifMetaConfig/{notifType}/{taskId}"
	reqPath = strings.ReplaceAll(reqPath, "{notifType}", url.PathEscape("instant"))
	reqPath = strings.ReplaceAll(reqPath, "{taskId}", url.PathEscape(plan.Id.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_notif_meta_config", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_notif_meta_config", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error updating gigavuecore_notif_meta_config", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error updating gigavuecore_notif_meta_config", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error updating gigavuecore_notif_meta_config", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error updating gigavuecore_notif_meta_config", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error updating gigavuecore_notif_meta_config", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error updating gigavuecore_notif_meta_config", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error updating gigavuecore_notif_meta_config", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error updating gigavuecore_notif_meta_config", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error updating gigavuecore_notif_meta_config", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error updating gigavuecore_notif_meta_config", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error updating gigavuecore_notif_meta_config", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Delete destroys the remote resource.
func (r *NotifMetaConfigResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NotifMetaConfigResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	timeout, diags := state.Timeouts.Delete(ctx, 10*time.Minute)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	r.deleteRemote(ctx, &state, resp)
}

// deleteRemote performs the delete HTTP exchange, treating a 404 as already deleted. Extracted from Delete so the request/response logic is unit-testable without a tfsdk.State.
func (r *NotifMetaConfigResource) deleteRemote(ctx context.Context, state *NotifMetaConfigResourceModel, resp *resource.DeleteResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/notification/event/notifMetaConfig/{notifType}/{taskId}"
	reqPath = strings.ReplaceAll(reqPath, "{notifType}", url.PathEscape("instant"))
	reqPath = strings.ReplaceAll(reqPath, "{taskId}", url.PathEscape(state.Id.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_notif_meta_config", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting gigavuecore_notif_meta_config", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error deleting gigavuecore_notif_meta_config", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error deleting gigavuecore_notif_meta_config", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error deleting gigavuecore_notif_meta_config", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error deleting gigavuecore_notif_meta_config", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error deleting gigavuecore_notif_meta_config", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error deleting gigavuecore_notif_meta_config", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error deleting gigavuecore_notif_meta_config", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error deleting gigavuecore_notif_meta_config", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error deleting gigavuecore_notif_meta_config", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *NotifMetaConfigResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData))
		return
	}
	r.client = c
}
