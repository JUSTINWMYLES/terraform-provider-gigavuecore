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
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*UpdateFmInstanceAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*UpdateFmInstanceAction)(nil)

// UpdateFmInstanceAction is the generated Terraform action implementation.
type UpdateFmInstanceAction struct {
	client *client.Client
}

// UpdateFmInstanceActionModel describes the action configuration shape.
type UpdateFmInstanceActionModel struct {
	EligibleNodeRole    types.String  `tfsdk:"eligible_node_role" json:"eligibleNodeRole"`
	EntityId            types.String  `tfsdk:"entity_id" json:"entityId"`
	HaGroupName         types.String  `tfsdk:"ha_group_name"`
	HostName            types.String  `tfsdk:"host_name" json:"hostName"`
	ManagementIpAddress types.String  `tfsdk:"management_ip_address" json:"managementIpAddress"`
	Name                types.String  `tfsdk:"name"`
	NodeRoles           types.Dynamic `tfsdk:"node_roles" json:"nodeRoles"`
	Nodes               types.Dynamic `tfsdk:"nodes"`
	Password            types.String  `tfsdk:"password"`
	PublicIpAddress     types.String  `tfsdk:"public_ip_address" json:"publicIpAddress"`
	Username            types.String  `tfsdk:"username"`
}

// NewUpdateFmInstanceAction returns a new instance of the generated action.
func NewUpdateFmInstanceAction() action.Action {
	return &UpdateFmInstanceAction{}
}

// Metadata returns the action type name.
func (r *UpdateFmInstanceAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_update_fm_instance"
}

// Schema returns the action schema.
func (r *UpdateFmInstanceAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Update FM Instances in HA Group", Attributes: map[string]schema.Attribute{"eligible_node_role": schema.StringAttribute{MarkdownDescription: "Eligible Node Role of the HA group", Required: true}, "entity_id": schema.StringAttribute{MarkdownDescription: "Entity ID of the HA group", Required: true}, "ha_group_name": schema.StringAttribute{MarkdownDescription: "HA group name", Required: true}, "host_name": schema.StringAttribute{MarkdownDescription: "DNS Name or IP Address of the HA group", Required: true}, "management_ip_address": schema.StringAttribute{MarkdownDescription: "Management Ip Address of the HA group", Required: true}, "name": schema.StringAttribute{MarkdownDescription: "name of the HA group", Required: true}, "node_roles": schema.DynamicAttribute{Required: true}, "nodes": schema.DynamicAttribute{Required: true}, "password": schema.StringAttribute{MarkdownDescription: "Password of the HA group", Required: true}, "public_ip_address": schema.StringAttribute{MarkdownDescription: "Public Ip Address of the HA group", Required: true}, "username": schema.StringAttribute{MarkdownDescription: "Username of the HA group", Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *UpdateFmInstanceAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UpdateFmInstanceActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *UpdateFmInstanceAction) invokeRemote(ctx context.Context, config *UpdateFmInstanceActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/fmHa/{haGroupName}/update"
	reqPath = strings.ReplaceAll(reqPath, "{haGroupName}", url.PathEscape(config.HaGroupName.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_fm_instance", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_fm_instance", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPatch, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_fm_instance", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_fm_instance", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_fm_instance", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_fm_instance", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_fm_instance", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_fm_instance", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_fm_instance", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_fm_instance", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_fm_instance", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_update_fm_instance", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_fm_instance", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *UpdateFmInstanceAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
