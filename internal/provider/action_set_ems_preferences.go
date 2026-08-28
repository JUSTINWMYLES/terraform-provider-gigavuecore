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
var _ action.Action = (*SetEmsPreferencesAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*SetEmsPreferencesAction)(nil)

// SetEmsPreferencesAction is the generated Terraform action implementation.
type SetEmsPreferencesAction struct {
	client *client.Client
}

// SetEmsPreferencesActionModel describes the action configuration shape.
type SetEmsPreferencesActionModel struct {
	EmsEntitlementsEnabled types.Bool `tfsdk:"ems_entitlements_enabled" json:"emsEntitlementsEnabled"`
	EmsTelemetryEnabled    types.Bool `tfsdk:"ems_telemetry_enabled" json:"emsTelemetryEnabled"`
	EmsUsageEnabled        types.Bool `tfsdk:"ems_usage_enabled" json:"emsUsageEnabled"`
}

// NewSetEmsPreferencesAction returns a new instance of the generated action.
func NewSetEmsPreferencesAction() action.Action {
	return &SetEmsPreferencesAction{}
}

// Metadata returns the action type name.
func (r *SetEmsPreferencesAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_set_ems_preferences"
}

// Schema returns the action schema.
func (r *SetEmsPreferencesAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Set opt-in preferences for entitlements/automatic renewals, usage reporting and telemetry", Attributes: map[string]schema.Attribute{"ems_entitlements_enabled": schema.BoolAttribute{Optional: true}, "ems_telemetry_enabled": schema.BoolAttribute{Optional: true}, "ems_usage_enabled": schema.BoolAttribute{Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *SetEmsPreferencesAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config SetEmsPreferencesActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *SetEmsPreferencesAction) invokeRemote(ctx context.Context, config *SetEmsPreferencesActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/licensing/ems/preferences"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_set_ems_preferences", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_set_ems_preferences", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_set_ems_preferences", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_set_ems_preferences", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_set_ems_preferences", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_set_ems_preferences", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_set_ems_preferences", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_set_ems_preferences", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_set_ems_preferences", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_set_ems_preferences", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_set_ems_preferences", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_set_ems_preferences", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_set_ems_preferences", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *SetEmsPreferencesAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
