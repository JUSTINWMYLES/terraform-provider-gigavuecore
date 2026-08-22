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
var _ action.Action = (*UpdateIcapNrtStatsConfigAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*UpdateIcapNrtStatsConfigAction)(nil)

// UpdateIcapNrtStatsConfigAction is the generated Terraform action implementation.
type UpdateIcapNrtStatsConfigAction struct {
	client *client.Client
}

// UpdateIcapNrtStatsConfigActionModel describes the action configuration shape.
type UpdateIcapNrtStatsConfigActionModel struct {
	OperationType types.String `tfsdk:"operation_type"`
	SolutionAlias types.String `tfsdk:"solution_alias"`
}

// NewUpdateIcapNrtStatsConfigAction returns a new instance of the generated action.
func NewUpdateIcapNrtStatsConfigAction() action.Action {
	return &UpdateIcapNrtStatsConfigAction{}
}

// Metadata returns the action type name.
func (r *UpdateIcapNrtStatsConfigAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_update_icap_nrt_stats_config"
}

// Schema returns the action schema.
func (r *UpdateIcapNrtStatsConfigAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Register Icap solution to NRT stats", Attributes: map[string]schema.Attribute{"operation_type": schema.StringAttribute{MarkdownDescription: "adds icap solution to NRT", Required: true}, "solution_alias": schema.StringAttribute{MarkdownDescription: "Register icap solution to NRT", Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *UpdateIcapNrtStatsConfigAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UpdateIcapNrtStatsConfigActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *UpdateIcapNrtStatsConfigAction) invokeRemote(ctx context.Context, config *UpdateIcapNrtStatsConfigActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/icap/stats/{solutionAlias}"
	reqPath = strings.ReplaceAll(reqPath, "{solutionAlias}", url.PathEscape(config.SolutionAlias.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodPatch, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_icap_nrt_stats_config", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("operationType", config.OperationType.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_icap_nrt_stats_config", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_icap_nrt_stats_config", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_icap_nrt_stats_config", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_icap_nrt_stats_config", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_icap_nrt_stats_config", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_icap_nrt_stats_config", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_icap_nrt_stats_config", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_icap_nrt_stats_config", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_update_icap_nrt_stats_config", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_icap_nrt_stats_config", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *UpdateIcapNrtStatsConfigAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
