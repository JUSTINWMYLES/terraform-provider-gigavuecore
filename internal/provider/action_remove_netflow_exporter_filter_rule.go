package provider

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*RemoveNetflowExporterFilterRuleAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*RemoveNetflowExporterFilterRuleAction)(nil)

// RemoveNetflowExporterFilterRuleAction is the generated Terraform action implementation.
type RemoveNetflowExporterFilterRuleAction struct {
	client *client.Client
}

// RemoveNetflowExporterFilterRuleActionModel describes the action configuration shape.
type RemoveNetflowExporterFilterRuleActionModel struct {
	Alias     types.String `tfsdk:"alias"`
	ClusterId types.String `tfsdk:"cluster_id"`
	RuleId    types.Int64  `tfsdk:"rule_id"`
}

// NewRemoveNetflowExporterFilterRuleAction returns a new instance of the generated action.
func NewRemoveNetflowExporterFilterRuleAction() action.Action {
	return &RemoveNetflowExporterFilterRuleAction{}
}

// Metadata returns the action type name.
func (r *RemoveNetflowExporterFilterRuleAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_remove_netflow_exporter_filter_rule"
}

// Schema returns the action schema.
func (r *RemoveNetflowExporterFilterRuleAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Remove a rule from Netflow Exporter Filter", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "alias of the target Netflow Exporter", Required: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "rule_id": schema.Int64Attribute{MarkdownDescription: "filter rule id", Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *RemoveNetflowExporterFilterRuleAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config RemoveNetflowExporterFilterRuleActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *RemoveNetflowExporterFilterRuleAction) invokeRemote(ctx context.Context, config *RemoveNetflowExporterFilterRuleActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/netflow/exporters/{alias}/filter/rules/{ruleId}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	reqPath = strings.ReplaceAll(reqPath, "{ruleId}", url.PathEscape(strconv.FormatInt(config.RuleId.ValueInt64(), 10)))
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_remove_netflow_exporter_filter_rule", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_remove_netflow_exporter_filter_rule", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 204) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_remove_netflow_exporter_filter_rule", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_remove_netflow_exporter_filter_rule", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_remove_netflow_exporter_filter_rule", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_remove_netflow_exporter_filter_rule", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_remove_netflow_exporter_filter_rule", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_remove_netflow_exporter_filter_rule", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_remove_netflow_exporter_filter_rule", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_remove_netflow_exporter_filter_rule", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_remove_netflow_exporter_filter_rule", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *RemoveNetflowExporterFilterRuleAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
