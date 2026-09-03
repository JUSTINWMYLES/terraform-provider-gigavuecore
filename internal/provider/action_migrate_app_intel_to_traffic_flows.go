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
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*MigrateAppIntelToTrafficFlowsAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*MigrateAppIntelToTrafficFlowsAction)(nil)

// MigrateAppIntelToTrafficFlowsAction is the generated Terraform action implementation.
type MigrateAppIntelToTrafficFlowsAction struct {
	client *client.Client
}

// MigrateAppIntelToTrafficFlowsActionModel describes the action configuration shape.
type MigrateAppIntelToTrafficFlowsActionModel struct {
	Alias          types.String `tfsdk:"alias"`
	ClusterId      types.String `tfsdk:"cluster_id" json:"clusterId"`
	DryRun         types.Bool   `tfsdk:"dry_run"`
	Force          types.Bool   `tfsdk:"force"`
	MigrationAlias types.String `tfsdk:"migration_alias" json:"migrationAlias"`
}

// NewMigrateAppIntelToTrafficFlowsAction returns a new instance of the generated action.
func NewMigrateAppIntelToTrafficFlowsAction() action.Action {
	return &MigrateAppIntelToTrafficFlowsAction{}
}

// Metadata returns the action type name.
func (r *MigrateAppIntelToTrafficFlowsAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_migrate_app_intel_to_traffic_flows"
}

// Schema returns the action schema.
func (r *MigrateAppIntelToTrafficFlowsAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Migrate AppIntel solution to Traffic Flows", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{Optional: true}, "cluster_id": schema.StringAttribute{Optional: true}, "dry_run": schema.BoolAttribute{MarkdownDescription: "Simulate migration without making changes", Optional: true}, "force": schema.BoolAttribute{MarkdownDescription: "Force migration even if conflicts exist", Optional: true}, "migration_alias": schema.StringAttribute{Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *MigrateAppIntelToTrafficFlowsAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config MigrateAppIntelToTrafficFlowsActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *MigrateAppIntelToTrafficFlowsAction) invokeRemote(ctx context.Context, config *MigrateAppIntelToTrafficFlowsActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/traffic-flows/migrate/appviz"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_app_intel_to_traffic_flows", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_app_intel_to_traffic_flows", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_app_intel_to_traffic_flows", fmt.Sprintf("Could not build request: %s", err))
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
		resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_app_intel_to_traffic_flows", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_app_intel_to_traffic_flows", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_app_intel_to_traffic_flows", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_app_intel_to_traffic_flows", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_app_intel_to_traffic_flows", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_app_intel_to_traffic_flows", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_app_intel_to_traffic_flows", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_app_intel_to_traffic_flows", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_app_intel_to_traffic_flows", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_migrate_app_intel_to_traffic_flows", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *MigrateAppIntelToTrafficFlowsAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
