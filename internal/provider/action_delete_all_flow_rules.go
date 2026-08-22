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
var _ action.Action = (*DeleteAllFlowRulesAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*DeleteAllFlowRulesAction)(nil)

// DeleteAllFlowRulesAction is the generated Terraform action implementation.
type DeleteAllFlowRulesAction struct {
	client *client.Client
}

// DeleteAllFlowRulesActionModel describes the action configuration shape.
type DeleteAllFlowRulesActionModel struct {
	FlowAlias    types.String `tfsdk:"flow_alias"`
	PolicyAlias  types.String `tfsdk:"policy_alias"`
	SubFlowAlias types.String `tfsdk:"sub_flow_alias"`
}

// NewDeleteAllFlowRulesAction returns a new instance of the generated action.
func NewDeleteAllFlowRulesAction() action.Action {
	return &DeleteAllFlowRulesAction{}
}

// Metadata returns the action type name.
func (r *DeleteAllFlowRulesAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_delete_all_flow_rules"
}

// Schema returns the action schema.
func (r *DeleteAllFlowRulesAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Delete all flow rules for a subflow", Attributes: map[string]schema.Attribute{"flow_alias": schema.StringAttribute{MarkdownDescription: "Flow identifier", Required: true}, "policy_alias": schema.StringAttribute{MarkdownDescription: "Policy identifier", Required: true}, "sub_flow_alias": schema.StringAttribute{MarkdownDescription: "Subflow identifier", Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *DeleteAllFlowRulesAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config DeleteAllFlowRulesActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *DeleteAllFlowRulesAction) invokeRemote(ctx context.Context, config *DeleteAllFlowRulesActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/bulk/traffic-flows/{policyAlias}/flow/{flowAlias}/subFlow/{subFlowAlias}"
	reqPath = strings.ReplaceAll(reqPath, "{policyAlias}", url.PathEscape(config.PolicyAlias.ValueString()))
	reqPath = strings.ReplaceAll(reqPath, "{flowAlias}", url.PathEscape(config.FlowAlias.ValueString()))
	reqPath = strings.ReplaceAll(reqPath, "{subFlowAlias}", url.PathEscape(config.SubFlowAlias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_delete_all_flow_rules", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_delete_all_flow_rules", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 202) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_all_flow_rules", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_all_flow_rules", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_all_flow_rules", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_all_flow_rules", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_all_flow_rules", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_all_flow_rules", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_all_flow_rules", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_delete_all_flow_rules", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_all_flow_rules", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *DeleteAllFlowRulesAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
