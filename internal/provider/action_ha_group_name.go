package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*HaGroupNameAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*HaGroupNameAction)(nil)

// HaGroupNameAction is the generated Terraform action implementation.
type HaGroupNameAction struct {
	client *client.Client
}

// HaGroupNameActionModel describes the action configuration shape.
type HaGroupNameActionModel struct {
	EligibleNodeRole    types.String `tfsdk:"eligible_node_role" json:"eligibleNodeRole"`
	EntityId            types.String `tfsdk:"entity_id" json:"entityId"`
	HaGroupName         types.String `tfsdk:"ha_group_name"`
	HostName            types.String `tfsdk:"host_name" json:"hostName"`
	ManagementIpAddress types.String `tfsdk:"management_ip_address" json:"managementIpAddress"`
	Name                types.String `tfsdk:"name"`
	NodeRoles           types.Object `tfsdk:"node_roles" json:"nodeRoles"`
	Nodes               types.Object `tfsdk:"nodes"`
	Password            types.String `tfsdk:"password"`
	PublicIpAddress     types.String `tfsdk:"public_ip_address" json:"publicIpAddress"`
	Username            types.String `tfsdk:"username"`
}

// NewHaGroupNameAction returns a new instance of the generated action.
func NewHaGroupNameAction() action.Action {
	return &HaGroupNameAction{}
}

// Metadata returns the action type name.
func (r *HaGroupNameAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_ha_group_name"
}

// Schema returns the action schema.
func (r *HaGroupNameAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Add FM Instance to HA Group", Attributes: map[string]schema.Attribute{"eligible_node_role": schema.StringAttribute{MarkdownDescription: "Eligible Node Role of the HA group", Required: true}, "entity_id": schema.StringAttribute{MarkdownDescription: "Entity ID of the HA group", Required: true}, "ha_group_name": schema.StringAttribute{MarkdownDescription: "HA group name", Required: true}, "host_name": schema.StringAttribute{MarkdownDescription: "DNS Name or IP Address of the HA group", Required: true}, "management_ip_address": schema.StringAttribute{MarkdownDescription: "Management Ip Address of the HA group", Required: true}, "name": schema.StringAttribute{MarkdownDescription: "name of the HA group", Required: true}, "node_roles": schema.SingleNestedAttribute{Required: true, Attributes: map[string]schema.Attribute{"es_node_role": schema.StringAttribute{MarkdownDescription: "ES node role", Optional: true}, "mongo_node_role": schema.StringAttribute{MarkdownDescription: "Mongo node role", Optional: true}}}, "nodes": schema.SingleNestedAttribute{Required: true, Attributes: map[string]schema.Attribute{"cluster_ip_address": schema.StringAttribute{MarkdownDescription: "Cluster IP address for FM HA node", Required: true}, "entity_id": schema.StringAttribute{MarkdownDescription: "Entity Id for FM HA node", Required: true}, "hostname": schema.StringAttribute{MarkdownDescription: "DNS Name or IP Address for FM HA node", Required: true}, "idp_meta_data_url": schema.StringAttribute{MarkdownDescription: "IDP Meta data URL for FM HA node", Optional: true}, "management_ip_address": schema.StringAttribute{MarkdownDescription: "Management IP address for FM HA node ", Required: true}, "password": schema.StringAttribute{MarkdownDescription: "password for FM HA node", Required: true}, "public_ip_address": schema.StringAttribute{MarkdownDescription: "Public IP address for FM HA node", Required: true}, "reachable": schema.BoolAttribute{MarkdownDescription: "Reachable for FM HA node", Required: true}, "seed_node": schema.BoolAttribute{MarkdownDescription: "Seed node for FM HA node", Required: true}, "username": schema.StringAttribute{MarkdownDescription: "username for FM HA node", Required: true}}}, "password": schema.StringAttribute{MarkdownDescription: "Password of the HA group", Required: true}, "public_ip_address": schema.StringAttribute{MarkdownDescription: "Public Ip Address of the HA group", Required: true}, "username": schema.StringAttribute{MarkdownDescription: "Username of the HA group", Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *HaGroupNameAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config HaGroupNameActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *HaGroupNameAction) invokeRemote(ctx context.Context, config *HaGroupNameActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/fmHa/{haGroupName}"
	reqPath = strings.ReplaceAll(reqPath, "{haGroupName}", url.PathEscape(config.HaGroupName.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_ha_group_name", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_ha_group_name", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_ha_group_name", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_ha_group_name", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201 || httpResp.StatusCode == 202) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_ha_group_name", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_ha_group_name", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_ha_group_name", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_ha_group_name", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_ha_group_name", "Entity Already Exists. See errors payload for details")
			return
		case 417:
			resp.Diagnostics.AddError("Error invoking gigavuecore_ha_group_name", "Deployment failed. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_ha_group_name", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_ha_group_name", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_ha_group_name", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_ha_group_name", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *HaGroupNameAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
