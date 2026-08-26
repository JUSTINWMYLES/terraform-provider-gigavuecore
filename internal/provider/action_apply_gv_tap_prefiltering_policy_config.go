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
var _ action.Action = (*ApplyGvTapPrefilteringPolicyConfigAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*ApplyGvTapPrefilteringPolicyConfigAction)(nil)

// ApplyGvTapPrefilteringPolicyConfigAction is the generated Terraform action implementation.
type ApplyGvTapPrefilteringPolicyConfigAction struct {
	client *client.Client
}

// ApplyGvTapPrefilteringPolicyConfigActionModel describes the action configuration shape.
type ApplyGvTapPrefilteringPolicyConfigActionModel struct {
	Id    types.String  `tfsdk:"id"`
	Name  types.String  `tfsdk:"name"`
	Rules types.Dynamic `tfsdk:"rules"`
}

// NewApplyGvTapPrefilteringPolicyConfigAction returns a new instance of the generated action.
func NewApplyGvTapPrefilteringPolicyConfigAction() action.Action {
	return &ApplyGvTapPrefilteringPolicyConfigAction{}
}

// Metadata returns the action type name.
func (r *ApplyGvTapPrefilteringPolicyConfigAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_apply_gv_tap_prefiltering_policy_config"
}

// Schema returns the action schema.
func (r *ApplyGvTapPrefilteringPolicyConfigAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Apply Prefiltering Policy Config [marked for deprecation]", Attributes: map[string]schema.Attribute{"id": schema.StringAttribute{MarkdownDescription: "policy graph id", Required: true}, "name": schema.StringAttribute{MarkdownDescription: "Unique name of this traffic policy. It can be between 1 and 32 characters long and may contain only alpha numeric characters, underscores and dashes.", Required: true}, "rules": schema.DynamicAttribute{MarkdownDescription: "Rules defined for this traffic policy. At least one rule has to be specified and a maximum of 16 rules could be specified.", Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *ApplyGvTapPrefilteringPolicyConfigAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config ApplyGvTapPrefilteringPolicyConfigActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *ApplyGvTapPrefilteringPolicyConfigAction) invokeRemote(ctx context.Context, config *ApplyGvTapPrefilteringPolicyConfigActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/unifiedResources/gvtapAgentPolicy/prefiltering/policyGraph/{id}/apply"
	reqPath = strings.ReplaceAll(reqPath, "{id}", url.PathEscape(config.Id.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_apply_gv_tap_prefiltering_policy_config", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_apply_gv_tap_prefiltering_policy_config", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_apply_gv_tap_prefiltering_policy_config", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_apply_gv_tap_prefiltering_policy_config", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_apply_gv_tap_prefiltering_policy_config", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_apply_gv_tap_prefiltering_policy_config", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_apply_gv_tap_prefiltering_policy_config", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_apply_gv_tap_prefiltering_policy_config", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_apply_gv_tap_prefiltering_policy_config", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_apply_gv_tap_prefiltering_policy_config", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_apply_gv_tap_prefiltering_policy_config", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_apply_gv_tap_prefiltering_policy_config", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_apply_gv_tap_prefiltering_policy_config", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *ApplyGvTapPrefilteringPolicyConfigAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
