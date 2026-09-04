package provider

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*MigrateAllFabricMapsToTrafficFlowsAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*MigrateAllFabricMapsToTrafficFlowsAction)(nil)

// MigrateAllFabricMapsToTrafficFlowsAction is the generated Terraform action implementation.
type MigrateAllFabricMapsToTrafficFlowsAction struct {
	client *client.Client
}

// MigrateAllFabricMapsToTrafficFlowsActionModel describes the action configuration shape.
type MigrateAllFabricMapsToTrafficFlowsActionModel struct {
	DryRun types.Bool `tfsdk:"dry_run"`
	Force  types.Bool `tfsdk:"force"`
}

// NewMigrateAllFabricMapsToTrafficFlowsAction returns a new instance of the generated action.
func NewMigrateAllFabricMapsToTrafficFlowsAction() action.Action {
	return &MigrateAllFabricMapsToTrafficFlowsAction{}
}

// Metadata returns the action type name.
func (r *MigrateAllFabricMapsToTrafficFlowsAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_migrate_all_fabric_maps_to_traffic_flows"
}

// Schema returns the action schema.
func (r *MigrateAllFabricMapsToTrafficFlowsAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Migrate all fabric maps to traffic flows", Attributes: map[string]schema.Attribute{"dry_run": schema.BoolAttribute{MarkdownDescription: "Simulate migration without making changes", Optional: true}, "force": schema.BoolAttribute{MarkdownDescription: "Force migration even if conflicts exist", Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *MigrateAllFabricMapsToTrafficFlowsAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config MigrateAllFabricMapsToTrafficFlowsActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *MigrateAllFabricMapsToTrafficFlowsAction) invokeRemote(ctx context.Context, config *MigrateAllFabricMapsToTrafficFlowsActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/traffic-flows/migrate/fabricMaps/all"
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_all_fabric_maps_to_traffic_flows", fmt.Sprintf("Could not build request: %s", err))
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
		resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_all_fabric_maps_to_traffic_flows", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_all_fabric_maps_to_traffic_flows", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_all_fabric_maps_to_traffic_flows", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_all_fabric_maps_to_traffic_flows", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_all_fabric_maps_to_traffic_flows", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_all_fabric_maps_to_traffic_flows", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_all_fabric_maps_to_traffic_flows", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_all_fabric_maps_to_traffic_flows", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_all_fabric_maps_to_traffic_flows", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_all_fabric_maps_to_traffic_flows", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *MigrateAllFabricMapsToTrafficFlowsAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
