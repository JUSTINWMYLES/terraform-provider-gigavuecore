package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*DeleteFmNtpServerAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*DeleteFmNtpServerAction)(nil)

// DeleteFmNtpServerAction is the generated Terraform action implementation.
type DeleteFmNtpServerAction struct {
	client *client.Client
}

// DeleteFmNtpServerActionModel describes the action configuration shape.
type DeleteFmNtpServerActionModel struct {
	AuthRequired    types.Bool   `tfsdk:"auth_required" json:"authRequired"`
	BodyFmIp        types.String `tfsdk:"body_fm_ip" json:"fmIp"`
	FmIp            types.String `tfsdk:"fm_ip"`
	IsUserNtpServer types.Bool   `tfsdk:"is_user_ntp_server" json:"isUserNtpServer"`
	NtpAuth         types.Object `tfsdk:"ntp_auth" json:"ntpAuth"`
	ServerHost      types.String `tfsdk:"server_host" json:"serverHost"`
	ServerStatus    types.Object `tfsdk:"server_status" json:"serverStatus"`
	Version         types.Int64  `tfsdk:"version"`
}

// NewDeleteFmNtpServerAction returns a new instance of the generated action.
func NewDeleteFmNtpServerAction() action.Action {
	return &DeleteFmNtpServerAction{}
}

// Metadata returns the action type name.
func (r *DeleteFmNtpServerAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_delete_fm_ntp_server"
}

// Schema returns the action schema.
func (r *DeleteFmNtpServerAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Delete NTP Server by servername", Attributes: map[string]schema.Attribute{"auth_required": schema.BoolAttribute{MarkdownDescription: "authentication enabled status", Required: true}, "body_fm_ip": schema.StringAttribute{MarkdownDescription: "FM HighAvailability Node address/host", Optional: true}, "fm_ip": schema.StringAttribute{MarkdownDescription: "FMHighAvailability node IpAddress/domainName", Optional: true}, "is_user_ntp_server": schema.BoolAttribute{MarkdownDescription: "To differentiate user created ntp server and default ntp server", Optional: true}, "ntp_auth": schema.SingleNestedAttribute{MarkdownDescription: "NTP Auth Details", Optional: true, Attributes: map[string]schema.Attribute{"key": schema.StringAttribute{MarkdownDescription: "key for the ntp server", Required: true}, "type": schema.StringAttribute{MarkdownDescription: "Type of algorithm used for ntp server authentication", Required: true}, "value": schema.StringAttribute{MarkdownDescription: "value for the ntp server authentication", Required: true}}}, "server_host": schema.StringAttribute{MarkdownDescription: "Ip/Host Address", Required: true}, "server_status": schema.SingleNestedAttribute{MarkdownDescription: "Server Status", Optional: true, Attributes: map[string]schema.Attribute{"offset": schema.Int64Attribute{Optional: true}, "poll_interval": schema.StringAttribute{Optional: true}, "status": schema.StringAttribute{MarkdownDescription: "status of the server", Required: true}, "stratum": schema.StringAttribute{MarkdownDescription: "status of the server", Optional: true}}}, "version": schema.Int64Attribute{MarkdownDescription: "version of server", Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *DeleteFmNtpServerAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config DeleteFmNtpServerActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *DeleteFmNtpServerAction) invokeRemote(ctx context.Context, config *DeleteFmNtpServerActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/fm/system/time/ntp/servers"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_delete_fm_ntp_server", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_delete_fm_ntp_server", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_delete_fm_ntp_server", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.BodyFmIp.IsNull() {
		query.Set("fmIp", config.BodyFmIp.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_delete_fm_ntp_server", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 204) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_fm_ntp_server", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_fm_ntp_server", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_fm_ntp_server", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_fm_ntp_server", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_fm_ntp_server", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_fm_ntp_server", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_fm_ntp_server", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_delete_fm_ntp_server", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_fm_ntp_server", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *DeleteFmNtpServerAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
