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
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*RedefineSyslogConfigAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*RedefineSyslogConfigAction)(nil)

// RedefineSyslogConfigAction is the generated Terraform action implementation.
type RedefineSyslogConfigAction struct {
	client *client.Client
}

// RedefineSyslogConfigActionModel describes the action configuration shape.
type RedefineSyslogConfigActionModel struct {
	ClusterId        types.String `tfsdk:"cluster_id"`
	ClusterName      types.String `tfsdk:"cluster_name" json:"clusterName"`
	SyslogConfigList types.List   `tfsdk:"syslog_config_list" json:"syslogConfigList"`
}

// NewRedefineSyslogConfigAction returns a new instance of the generated action.
func NewRedefineSyslogConfigAction() action.Action {
	return &RedefineSyslogConfigAction{}
}

// Metadata returns the action type name.
func (r *RedefineSyslogConfigAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_redefine_syslog_config"
}

// Schema returns the action schema.
func (r *RedefineSyslogConfigAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Redefine Syslog config", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "cluster_name": schema.StringAttribute{MarkdownDescription: "Name of the cluster", Required: true}, "syslog_config_list": schema.ListNestedAttribute{MarkdownDescription: "List of the syslog configuration specification for every node in the cluster", Required: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"device_ip": schema.StringAttribute{MarkdownDescription: "IP address of the device", Required: true}, "log_severity": schema.StringAttribute{MarkdownDescription: "Minimum syslog logging severity level", Required: true}, "target_hosts": schema.ListNestedAttribute{MarkdownDescription: "List of syslog targets the device is streaming to", Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"log_severity": schema.StringAttribute{MarkdownDescription: "Minimum logging severity level the device will stream for this target", Required: true}, "port": schema.Int64Attribute{MarkdownDescription: "0 represents UDP and non zero is TCP", Required: true}, "server": schema.StringAttribute{MarkdownDescription: "ipv4 or ipv6 or hostname", Required: true}, "ssh_enabled": schema.BoolAttribute{MarkdownDescription: "Is Syslog target configured to receive logs via SSH", Optional: true}, "streaming_enabled": schema.BoolAttribute{MarkdownDescription: "Is Syslog Streaming enabled", Required: true}, "username": schema.StringAttribute{MarkdownDescription: "For syslog over UDP there won't be any user. Only valid for SSH type", Optional: true}}}}}}}}}
}

// Invoke executes the action against the remote API.
func (r *RedefineSyslogConfigAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config RedefineSyslogConfigActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *RedefineSyslogConfigAction) invokeRemote(ctx context.Context, config *RedefineSyslogConfigActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/system/syslog/config/{clusterId}"
	reqPath = strings.ReplaceAll(reqPath, "{clusterId}", url.PathEscape(config.ClusterId.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_syslog_config", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_syslog_config", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_syslog_config", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_syslog_config", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_syslog_config", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_syslog_config", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_syslog_config", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_syslog_config", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_syslog_config", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_syslog_config", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_syslog_config", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_syslog_config", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_syslog_config", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *RedefineSyslogConfigAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
