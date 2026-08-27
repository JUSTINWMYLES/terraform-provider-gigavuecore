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
var _ action.Action = (*UndeployPolicyAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*UndeployPolicyAction)(nil)

// UndeployPolicyAction is the generated Terraform action implementation.
type UndeployPolicyAction struct {
	client *client.Client
}

// UndeployPolicyActionModel describes the action configuration shape.
type UndeployPolicyActionModel struct {
	Name types.String `tfsdk:"name"`
}

// NewUndeployPolicyAction returns a new instance of the generated action.
func NewUndeployPolicyAction() action.Action {
	return &UndeployPolicyAction{}
}

// Metadata returns the action type name.
func (r *UndeployPolicyAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_undeploy_policy"
}

// Schema returns the action schema.
func (r *UndeployPolicyAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Undeploy policy", Attributes: map[string]schema.Attribute{"name": schema.StringAttribute{MarkdownDescription: "policy name", Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *UndeployPolicyAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UndeployPolicyActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *UndeployPolicyAction) invokeRemote(ctx context.Context, config *UndeployPolicyActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/intent/policies/{name}/undeploy"
	reqPath = strings.ReplaceAll(reqPath, "{name}", url.PathEscape(config.Name.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_undeploy_policy", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_undeploy_policy", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 204) {
		apiErr, err := client.NewAPIError(httpResp)
		if err != nil {
			resp.Diagnostics.AddError("Error invoking gigavuecore_undeploy_policy", fmt.Sprintf("Could not read error response: %s", err))
			return
		}
		resp.Diagnostics.AddError("Error invoking gigavuecore_undeploy_policy", apiErr.Error())
		return
	}
}

// Configure stores the API client supplied by the provider.
func (r *UndeployPolicyAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
