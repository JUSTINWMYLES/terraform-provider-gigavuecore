package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*ConfigureEmailServerAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*ConfigureEmailServerAction)(nil)

// ConfigureEmailServerAction is the generated Terraform action implementation.
type ConfigureEmailServerAction struct {
	client *client.Client
}

// ConfigureEmailServerActionModel describes the action configuration shape.
type ConfigureEmailServerActionModel struct {
	EmailHost      types.String `tfsdk:"email_host" json:"emailHost"`
	EnableSmtpAuth types.Bool   `tfsdk:"enable_smtp_auth" json:"enableSmtpAuth"`
	From           types.String `tfsdk:"from"`
	Password       types.String `tfsdk:"password"`
	Port           types.Int64  `tfsdk:"port"`
	UpdateSecret   types.Bool   `tfsdk:"update_secret"`
	UserName       types.String `tfsdk:"user_name" json:"userName"`
}

// NewConfigureEmailServerAction returns a new instance of the generated action.
func NewConfigureEmailServerAction() action.Action {
	return &ConfigureEmailServerAction{}
}

// Metadata returns the action type name.
func (r *ConfigureEmailServerAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_configure_email_server"
}

// Schema returns the action schema.
func (r *ConfigureEmailServerAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Configure Email Server", Attributes: map[string]schema.Attribute{"email_host": schema.StringAttribute{MarkdownDescription: "Address of the SMTP Server", Optional: true}, "enable_smtp_auth": schema.BoolAttribute{MarkdownDescription: "Enable/Disable SMTP Authentication", Optional: true}, "from": schema.StringAttribute{MarkdownDescription: "From address for the email", Optional: true}, "password": schema.StringAttribute{MarkdownDescription: "Password for the SMTP Server", Optional: true}, "port": schema.Int64Attribute{MarkdownDescription: "SMTP Server Port", Optional: true}, "update_secret": schema.BoolAttribute{MarkdownDescription: "updateSecret", Optional: true}, "user_name": schema.StringAttribute{MarkdownDescription: "Username for the SMTP Server", Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *ConfigureEmailServerAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config ConfigureEmailServerActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *ConfigureEmailServerAction) invokeRemote(ctx context.Context, config *ConfigureEmailServerActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/notification/email/emailServer"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_configure_email_server", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_configure_email_server", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_configure_email_server", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.UpdateSecret.IsNull() {
		query.Set("updateSecret", strconv.FormatBool(config.UpdateSecret.ValueBool()))
	}
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_configure_email_server", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_configure_email_server", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_configure_email_server", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_configure_email_server", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_configure_email_server", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_configure_email_server", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_configure_email_server", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_configure_email_server", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_configure_email_server", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_configure_email_server", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *ConfigureEmailServerAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
