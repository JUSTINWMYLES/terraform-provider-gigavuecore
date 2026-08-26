package provider

import (
	"context"
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
var _ action.Action = (*ClearPrecryptionPolicyConfigurationAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*ClearPrecryptionPolicyConfigurationAction)(nil)

// ClearPrecryptionPolicyConfigurationAction is the generated Terraform action implementation.
type ClearPrecryptionPolicyConfigurationAction struct {
	client *client.Client
}

// ClearPrecryptionPolicyConfigurationActionModel describes the action configuration shape.
type ClearPrecryptionPolicyConfigurationActionModel struct {
	Id types.String `tfsdk:"id"`
}

// NewClearPrecryptionPolicyConfigurationAction returns a new instance of the generated action.
func NewClearPrecryptionPolicyConfigurationAction() action.Action {
	return &ClearPrecryptionPolicyConfigurationAction{}
}

// Metadata returns the action type name.
func (r *ClearPrecryptionPolicyConfigurationAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_clear_precryption_policy_configuration"
}

// Schema returns the action schema.
func (r *ClearPrecryptionPolicyConfigurationAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Clear Precryption Policy configuration", Attributes: map[string]schema.Attribute{"id": schema.StringAttribute{MarkdownDescription: "policy graph id", Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *ClearPrecryptionPolicyConfigurationAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config ClearPrecryptionPolicyConfigurationActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *ClearPrecryptionPolicyConfigurationAction) invokeRemote(ctx context.Context, config *ClearPrecryptionPolicyConfigurationActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/unifiedResources/uctvPolicy/precryption/policyGraph/{id}/clear"
	reqPath = strings.ReplaceAll(reqPath, "{id}", url.PathEscape(config.Id.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodPatch, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_clear_precryption_policy_configuration", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_clear_precryption_policy_configuration", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 204) {
		switch httpResp.StatusCode {
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_clear_precryption_policy_configuration", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_clear_precryption_policy_configuration", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_clear_precryption_policy_configuration", "Entity Not Found. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_clear_precryption_policy_configuration", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_clear_precryption_policy_configuration", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_clear_precryption_policy_configuration", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_clear_precryption_policy_configuration", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *ClearPrecryptionPolicyConfigurationAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
