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
var _ action.Action = (*RollbackTrafficFlowsMigrationAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*RollbackTrafficFlowsMigrationAction)(nil)

// RollbackTrafficFlowsMigrationAction is the generated Terraform action implementation.
type RollbackTrafficFlowsMigrationAction struct {
	client *client.Client
}

// RollbackTrafficFlowsMigrationActionModel describes the action configuration shape.
type RollbackTrafficFlowsMigrationActionModel struct {
	Force        types.Bool `tfsdk:"force"`
	TrafficFlows types.List `tfsdk:"traffic_flows" json:"trafficFlows"`
}

// NewRollbackTrafficFlowsMigrationAction returns a new instance of the generated action.
func NewRollbackTrafficFlowsMigrationAction() action.Action {
	return &RollbackTrafficFlowsMigrationAction{}
}

// Metadata returns the action type name.
func (r *RollbackTrafficFlowsMigrationAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_rollback_traffic_flows_migration"
}

// Schema returns the action schema.
func (r *RollbackTrafficFlowsMigrationAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Rollback traffic flows migration", Attributes: map[string]schema.Attribute{"force": schema.BoolAttribute{MarkdownDescription: "Force rollback even if conflicts exist", Optional: true}, "traffic_flows": schema.ListAttribute{Optional: true, ElementType: types.StringType}}}
}

// Invoke executes the action against the remote API.
func (r *RollbackTrafficFlowsMigrationAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config RollbackTrafficFlowsMigrationActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *RollbackTrafficFlowsMigrationAction) invokeRemote(ctx context.Context, config *RollbackTrafficFlowsMigrationActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/traffic-flows/migrate/rollback"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_rollback_traffic_flows_migration", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_rollback_traffic_flows_migration", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_rollback_traffic_flows_migration", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.Force.IsNull() {
		query.Set("force", strconv.FormatBool(config.Force.ValueBool()))
	}
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_rollback_traffic_flows_migration", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_rollback_traffic_flows_migration", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_rollback_traffic_flows_migration", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_rollback_traffic_flows_migration", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_rollback_traffic_flows_migration", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_rollback_traffic_flows_migration", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_rollback_traffic_flows_migration", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_rollback_traffic_flows_migration", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_rollback_traffic_flows_migration", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_rollback_traffic_flows_migration", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *RollbackTrafficFlowsMigrationAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
