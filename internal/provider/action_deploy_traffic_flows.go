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
var _ action.Action = (*DeployTrafficFlowsAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*DeployTrafficFlowsAction)(nil)

// DeployTrafficFlowsAction is the generated Terraform action implementation.
type DeployTrafficFlowsAction struct {
	client *client.Client
}

// DeployTrafficFlowsActionModel describes the action configuration shape.
type DeployTrafficFlowsActionModel struct {
	Alias types.String `tfsdk:"alias"`
}

// NewDeployTrafficFlowsAction returns a new instance of the generated action.
func NewDeployTrafficFlowsAction() action.Action {
	return &DeployTrafficFlowsAction{}
}

// Metadata returns the action type name.
func (r *DeployTrafficFlowsAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_deploy_traffic_flows"
}

// Schema returns the action schema.
func (r *DeployTrafficFlowsAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Deploy a Traffic Flows configuration", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Traffic Flows alias or ID", Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *DeployTrafficFlowsAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config DeployTrafficFlowsActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *DeployTrafficFlowsAction) invokeRemote(ctx context.Context, config *DeployTrafficFlowsActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/traffic-flows/deploy/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_deploy_traffic_flows", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_deploy_traffic_flows", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 202) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_deploy_traffic_flows", "Invalid request")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_deploy_traffic_flows", "Not Authenticated")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_deploy_traffic_flows", "Access Denied")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_deploy_traffic_flows", "Entity Not Found")
			return
		case 417:
			resp.Diagnostics.AddError("Error invoking gigavuecore_deploy_traffic_flows", "Policy deployment failed")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_deploy_traffic_flows", "Internal Server Error")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_deploy_traffic_flows", "Service Unavailable")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_deploy_traffic_flows", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_deploy_traffic_flows", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *DeployTrafficFlowsAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
