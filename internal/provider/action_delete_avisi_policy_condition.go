package provider

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*DeleteAvisiPolicyConditionAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*DeleteAvisiPolicyConditionAction)(nil)

// DeleteAvisiPolicyConditionAction is the generated Terraform action implementation.
type DeleteAvisiPolicyConditionAction struct {
	client *client.Client
}

// DeleteAvisiPolicyConditionActionModel describes the action configuration shape.
type DeleteAvisiPolicyConditionActionModel struct {
	ConditionId types.String `tfsdk:"condition_id"`
	PolicyId    types.String `tfsdk:"policy_id"`
}

// NewDeleteAvisiPolicyConditionAction returns a new instance of the generated action.
func NewDeleteAvisiPolicyConditionAction() action.Action {
	return &DeleteAvisiPolicyConditionAction{}
}

// Metadata returns the action type name.
func (r *DeleteAvisiPolicyConditionAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_delete_avisi_policy_condition"
}

// Schema returns the action schema.
func (r *DeleteAvisiPolicyConditionAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Delete Active Visibility Policy Condition", Attributes: map[string]schema.Attribute{"condition_id": schema.StringAttribute{MarkdownDescription: "Policy Condition Id", Required: true}, "policy_id": schema.StringAttribute{MarkdownDescription: "Policy Id", Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *DeleteAvisiPolicyConditionAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config DeleteAvisiPolicyConditionActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *DeleteAvisiPolicyConditionAction) invokeRemote(ctx context.Context, config *DeleteAvisiPolicyConditionActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/avisi/policies/{policyId}/conditions/{conditionId}"
	reqPath = strings.ReplaceAll(reqPath, "{policyId}", url.PathEscape(config.PolicyId.ValueString()))
	reqPath = strings.ReplaceAll(reqPath, "{conditionId}", url.PathEscape(config.ConditionId.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_delete_avisi_policy_condition", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_delete_avisi_policy_condition", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 204) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_avisi_policy_condition", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_avisi_policy_condition", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_avisi_policy_condition", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_avisi_policy_condition", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_avisi_policy_condition", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_avisi_policy_condition", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_avisi_policy_condition", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_delete_avisi_policy_condition", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_avisi_policy_condition", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *DeleteAvisiPolicyConditionAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
