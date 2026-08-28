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
var _ action.Action = (*ModifyEmailServerAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*ModifyEmailServerAction)(nil)

// ModifyEmailServerAction is the generated Terraform action implementation.
type ModifyEmailServerAction struct {
	client *client.Client
}

// ModifyEmailServerActionModel describes the action configuration shape.
type ModifyEmailServerActionModel struct {
	DomainName             types.String `tfsdk:"domain_name" json:"domainName"`
	EnableAutoSupportNotif types.Bool   `tfsdk:"enable_auto_support_notif" json:"enableAutoSupportNotif"`
	EnableSmtpAuth         types.Bool   `tfsdk:"enable_smtp_auth" json:"enableSmtpAuth"`
	IncludeHostname        types.Bool   `tfsdk:"include_hostname" json:"includeHostname"`
	MailHubPort            types.Int64  `tfsdk:"mail_hub_port" json:"mailHubPort"`
	Password               types.String `tfsdk:"password"`
	ReturnAddress          types.String `tfsdk:"return_address" json:"returnAddress"`
	SmtpServer             types.String `tfsdk:"smtp_server" json:"smtpServer"`
	Username               types.String `tfsdk:"username"`
}

// NewModifyEmailServerAction returns a new instance of the generated action.
func NewModifyEmailServerAction() action.Action {
	return &ModifyEmailServerAction{}
}

// Metadata returns the action type name.
func (r *ModifyEmailServerAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_modify_email_server"
}

// Schema returns the action schema.
func (r *ModifyEmailServerAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Modify system email server", Attributes: map[string]schema.Attribute{"domain_name": schema.StringAttribute{MarkdownDescription: "Domain Name", Optional: true}, "enable_auto_support_notif": schema.BoolAttribute{MarkdownDescription: "Enable Auto Support Notifications", Optional: true}, "enable_smtp_auth": schema.BoolAttribute{MarkdownDescription: "Enable SMTP auth", Optional: true}, "include_hostname": schema.BoolAttribute{MarkdownDescription: "Include Hostname", Optional: true}, "mail_hub_port": schema.Int64Attribute{Optional: true}, "password": schema.StringAttribute{MarkdownDescription: "SMTP Password", Optional: true}, "return_address": schema.StringAttribute{MarkdownDescription: "Return Address", Optional: true}, "smtp_server": schema.StringAttribute{MarkdownDescription: "SMTP Server", Optional: true}, "username": schema.StringAttribute{MarkdownDescription: "SMTP Username", Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *ModifyEmailServerAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config ModifyEmailServerActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *ModifyEmailServerAction) invokeRemote(ctx context.Context, config *ModifyEmailServerActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/system/email/server"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_modify_email_server", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_modify_email_server", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPatch, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_modify_email_server", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_modify_email_server", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_modify_email_server", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_modify_email_server", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_modify_email_server", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_modify_email_server", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_modify_email_server", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_modify_email_server", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_modify_email_server", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_modify_email_server", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_modify_email_server", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *ModifyEmailServerAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
