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
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*MigrateAllFlowMapToTrafficFlowsInClusterAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*MigrateAllFlowMapToTrafficFlowsInClusterAction)(nil)

// MigrateAllFlowMapToTrafficFlowsInClusterAction is the generated Terraform action implementation.
type MigrateAllFlowMapToTrafficFlowsInClusterAction struct {
	client *client.Client
}

// MigrateAllFlowMapToTrafficFlowsInClusterActionModel describes the action configuration shape.
type MigrateAllFlowMapToTrafficFlowsInClusterActionModel struct {
	ClusterId types.String `tfsdk:"cluster_id"`
	DryRun    types.Bool   `tfsdk:"dry_run"`
	Force     types.Bool   `tfsdk:"force"`
}

// NewMigrateAllFlowMapToTrafficFlowsInClusterAction returns a new instance of the generated action.
func NewMigrateAllFlowMapToTrafficFlowsInClusterAction() action.Action {
	return &MigrateAllFlowMapToTrafficFlowsInClusterAction{}
}

// Metadata returns the action type name.
func (r *MigrateAllFlowMapToTrafficFlowsInClusterAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_migrate_all_flow_map_to_traffic_flows_in_cluster"
}

// Schema returns the action schema.
func (r *MigrateAllFlowMapToTrafficFlowsInClusterAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Migrate all flow maps to traffic flows in a cluster", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "dry_run": schema.BoolAttribute{MarkdownDescription: "Simulate migration without making changes", Optional: true}, "force": schema.BoolAttribute{MarkdownDescription: "Force migration even if conflicts exist", Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *MigrateAllFlowMapToTrafficFlowsInClusterAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config MigrateAllFlowMapToTrafficFlowsInClusterActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *MigrateAllFlowMapToTrafficFlowsInClusterAction) invokeRemote(ctx context.Context, config *MigrateAllFlowMapToTrafficFlowsInClusterActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/traffic-flows/migrate/maps/{clusterId}/all"
	reqPath = strings.ReplaceAll(reqPath, "{clusterId}", url.PathEscape(config.ClusterId.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_all_flow_map_to_traffic_flows_in_cluster", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.Force.IsNull() {
		query.Set("force", strconv.FormatBool(config.Force.ValueBool()))
	}
	if !config.DryRun.IsNull() {
		query.Set("dryRun", strconv.FormatBool(config.DryRun.ValueBool()))
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_all_flow_map_to_traffic_flows_in_cluster", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_all_flow_map_to_traffic_flows_in_cluster", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_all_flow_map_to_traffic_flows_in_cluster", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_all_flow_map_to_traffic_flows_in_cluster", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_all_flow_map_to_traffic_flows_in_cluster", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_all_flow_map_to_traffic_flows_in_cluster", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_all_flow_map_to_traffic_flows_in_cluster", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_all_flow_map_to_traffic_flows_in_cluster", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_all_flow_map_to_traffic_flows_in_cluster", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_all_flow_map_to_traffic_flows_in_cluster", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *MigrateAllFlowMapToTrafficFlowsInClusterAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
