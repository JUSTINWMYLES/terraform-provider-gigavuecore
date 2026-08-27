package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*RedefineSnmpServerSystemConfigAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*RedefineSnmpServerSystemConfigAction)(nil)

// RedefineSnmpServerSystemConfigAction is the generated Terraform action implementation.
type RedefineSnmpServerSystemConfigAction struct {
	client *client.Client
}

// RedefineSnmpServerSystemConfigActionModel describes the action configuration shape.
type RedefineSnmpServerSystemConfigActionModel struct {
	ClusterId   types.String `tfsdk:"cluster_id"`
	Enabled     types.Bool   `tfsdk:"enabled"`
	EngineId    types.String `tfsdk:"engine_id" json:"engineId"`
	Port        types.Int64  `tfsdk:"port"`
	SysContact  types.String `tfsdk:"sys_contact" json:"sysContact"`
	SysDescr    types.String `tfsdk:"sys_descr" json:"sysDescr"`
	SysLocation types.String `tfsdk:"sys_location" json:"sysLocation"`
	SysName     types.String `tfsdk:"sys_name" json:"sysName"`
}

// NewRedefineSnmpServerSystemConfigAction returns a new instance of the generated action.
func NewRedefineSnmpServerSystemConfigAction() action.Action {
	return &RedefineSnmpServerSystemConfigAction{}
}

// Metadata returns the action type name.
func (r *RedefineSnmpServerSystemConfigAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_redefine_snmp_server_system_config"
}

// Schema returns the action schema.
func (r *RedefineSnmpServerSystemConfigAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Redefine Snmp Server System config", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "enabled": schema.BoolAttribute{MarkdownDescription: "Enables SNMP Server on the node", Optional: true}, "engine_id": schema.StringAttribute{MarkdownDescription: "local EngineID", Optional: true}, "port": schema.Int64Attribute{Optional: true}, "sys_contact": schema.StringAttribute{MarkdownDescription: "MIB-II 'sysContact'", Optional: true}, "sys_descr": schema.StringAttribute{MarkdownDescription: "MIB-II 'sysDescr'", Optional: true}, "sys_location": schema.StringAttribute{MarkdownDescription: "MIB-II 'sysLocation'", Optional: true}, "sys_name": schema.StringAttribute{MarkdownDescription: "MIB-II 'sysName'", Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *RedefineSnmpServerSystemConfigAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config RedefineSnmpServerSystemConfigActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *RedefineSnmpServerSystemConfigAction) invokeRemote(ctx context.Context, config *RedefineSnmpServerSystemConfigActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/system/snmp/system"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_snmp_server_system_config", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_snmp_server_system_config", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_snmp_server_system_config", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_snmp_server_system_config", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_snmp_server_system_config", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_snmp_server_system_config", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_snmp_server_system_config", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_snmp_server_system_config", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_snmp_server_system_config", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_snmp_server_system_config", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_snmp_server_system_config", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_snmp_server_system_config", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_snmp_server_system_config", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *RedefineSnmpServerSystemConfigAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
