package provider

import (
	"context"
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
var _ action.Action = (*ConfigureVolumeUsageAlertEnableAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*ConfigureVolumeUsageAlertEnableAction)(nil)

// ConfigureVolumeUsageAlertEnableAction is the generated Terraform action implementation.
type ConfigureVolumeUsageAlertEnableAction struct {
	client *client.Client
}

// ConfigureVolumeUsageAlertEnableActionModel describes the action configuration shape.
type ConfigureVolumeUsageAlertEnableActionModel struct {
	Enable    types.Bool  `tfsdk:"enable"`
	Threshold types.Int64 `tfsdk:"threshold"`
}

// NewConfigureVolumeUsageAlertEnableAction returns a new instance of the generated action.
func NewConfigureVolumeUsageAlertEnableAction() action.Action {
	return &ConfigureVolumeUsageAlertEnableAction{}
}

// Metadata returns the action type name.
func (r *ConfigureVolumeUsageAlertEnableAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_configure_volume_usage_alert_enable"
}

// Schema returns the action schema.
func (r *ConfigureVolumeUsageAlertEnableAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Configure volume usage alert enable", Attributes: map[string]schema.Attribute{"enable": schema.BoolAttribute{MarkdownDescription: "the enable value", Required: true}, "threshold": schema.Int64Attribute{MarkdownDescription: "the threshold value, positive (alert is triggered when usage reaches 'threshold' percentage of allowance)", Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *ConfigureVolumeUsageAlertEnableAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config ConfigureVolumeUsageAlertEnableActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *ConfigureVolumeUsageAlertEnableAction) invokeRemote(ctx context.Context, config *ConfigureVolumeUsageAlertEnableActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/licensing/config/alert/volumeUsage"
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_configure_volume_usage_alert_enable", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("enable", strconv.FormatBool(config.Enable.ValueBool()))
	query.Set("threshold", strconv.FormatInt(config.Threshold.ValueInt64(), 10))
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_configure_volume_usage_alert_enable", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_configure_volume_usage_alert_enable", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_configure_volume_usage_alert_enable", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_configure_volume_usage_alert_enable", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_configure_volume_usage_alert_enable", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_configure_volume_usage_alert_enable", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_configure_volume_usage_alert_enable", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_configure_volume_usage_alert_enable", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_configure_volume_usage_alert_enable", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_configure_volume_usage_alert_enable", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *ConfigureVolumeUsageAlertEnableAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
