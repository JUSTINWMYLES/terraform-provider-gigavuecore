package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*MigrateFabricMapsToTrafficFlows1Action)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*MigrateFabricMapsToTrafficFlows1Action)(nil)

// MigrateFabricMapsToTrafficFlows1Action is the generated Terraform action implementation.
type MigrateFabricMapsToTrafficFlows1Action struct {
	client *client.Client
}

// MigrateFabricMapsToTrafficFlows1ActionModel describes the action configuration shape.
type MigrateFabricMapsToTrafficFlows1ActionModel struct {
	Alias          types.List   `tfsdk:"alias"`
	ClusterIds     types.List   `tfsdk:"cluster_ids" json:"clusterIds"`
	DryRun         types.Bool   `tfsdk:"dry_run"`
	Force          types.Bool   `tfsdk:"force"`
	MigrationAlias types.String `tfsdk:"migration_alias" json:"migrationAlias"`
}

// NewMigrateFabricMapsToTrafficFlows1Action returns a new instance of the generated action.
func NewMigrateFabricMapsToTrafficFlows1Action() action.Action {
	return &MigrateFabricMapsToTrafficFlows1Action{}
}

// Metadata returns the action type name.
func (r *MigrateFabricMapsToTrafficFlows1Action) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_migrate_fabric_maps_to_traffic_flows_1"
}

// Schema returns the action schema.
func (r *MigrateFabricMapsToTrafficFlows1Action) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Migrate specific fabric maps to traffic flows", Attributes: map[string]schema.Attribute{"alias": schema.ListAttribute{Optional: true, ElementType: types.StringType}, "cluster_ids": schema.ListAttribute{Optional: true, ElementType: types.StringType}, "dry_run": schema.BoolAttribute{MarkdownDescription: "Simulate migration without making changes", Optional: true}, "force": schema.BoolAttribute{MarkdownDescription: "Force migration even if conflicts exist", Optional: true}, "migration_alias": schema.StringAttribute{Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *MigrateFabricMapsToTrafficFlows1Action) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config MigrateFabricMapsToTrafficFlows1ActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *MigrateFabricMapsToTrafficFlows1Action) invokeRemote(ctx context.Context, config *MigrateFabricMapsToTrafficFlows1ActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/traffic-flows/migrate/fabricMaps"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_fabric_maps_to_traffic_flows_1", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_fabric_maps_to_traffic_flows_1", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_fabric_maps_to_traffic_flows_1", fmt.Sprintf("Could not build request: %s", err))
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
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_fabric_maps_to_traffic_flows_1", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_fabric_maps_to_traffic_flows_1", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_fabric_maps_to_traffic_flows_1", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_fabric_maps_to_traffic_flows_1", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_fabric_maps_to_traffic_flows_1", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_fabric_maps_to_traffic_flows_1", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_fabric_maps_to_traffic_flows_1", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_fabric_maps_to_traffic_flows_1", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_fabric_maps_to_traffic_flows_1", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_fabric_maps_to_traffic_flows_1", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *MigrateFabricMapsToTrafficFlows1Action) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
