package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*CreateHaGroupAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*CreateHaGroupAction)(nil)

// CreateHaGroupAction is the generated Terraform action implementation.
type CreateHaGroupAction struct {
	client *client.Client
}

// CreateHaGroupActionModel describes the action configuration shape.
type CreateHaGroupActionModel struct {
	FmHaTunnel types.Object `tfsdk:"fm_ha_tunnel" json:"fmHaTunnel"`
	Name       types.String `tfsdk:"name"`
	Nodes      types.Object `tfsdk:"nodes"`
}

// NewCreateHaGroupAction returns a new instance of the generated action.
func NewCreateHaGroupAction() action.Action {
	return &CreateHaGroupAction{}
}

// Metadata returns the action type name.
func (r *CreateHaGroupAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_create_ha_group"
}

// Schema returns the action schema.
func (r *CreateHaGroupAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "create HA group", Attributes: map[string]schema.Attribute{"fm_ha_tunnel": schema.SingleNestedAttribute{MarkdownDescription: "Auth mode for creating tunnels in new FMHA cluster", Required: true, Attributes: map[string]schema.Attribute{"tunnel_auth_mode": schema.StringAttribute{MarkdownDescription: "Auth mode for setting up FMHA tunnel. PSK is Pre Shared Key. PKI is Public Key Infrastructure(Certificate based)", Optional: true}}}, "name": schema.StringAttribute{MarkdownDescription: "Name of the new FMHA cluster", Required: true}, "nodes": schema.SingleNestedAttribute{MarkdownDescription: "Nodes of the new FMHA cluster", Required: true, Attributes: map[string]schema.Attribute{"cluster_ip_address": schema.StringAttribute{MarkdownDescription: "Cluster IP address for FM HA node", Required: true}, "entity_id": schema.StringAttribute{MarkdownDescription: "Entity Id for FM HA node", Required: true}, "hostname": schema.StringAttribute{MarkdownDescription: "DNS Name or IP Address for FM HA node", Required: true}, "idp_meta_data_url": schema.StringAttribute{MarkdownDescription: "IDP Meta data URL for FM HA node", Optional: true}, "management_ip_address": schema.StringAttribute{MarkdownDescription: "Management IP address for FM HA node ", Required: true}, "password": schema.StringAttribute{MarkdownDescription: "password for FM HA node", Required: true}, "public_ip_address": schema.StringAttribute{MarkdownDescription: "Public IP address for FM HA node", Required: true}, "reachable": schema.BoolAttribute{MarkdownDescription: "Reachable for FM HA node", Required: true}, "seed_node": schema.BoolAttribute{MarkdownDescription: "Seed node for FM HA node", Required: true}, "username": schema.StringAttribute{MarkdownDescription: "username for FM HA node", Required: true}}}}}
}

// Invoke executes the action against the remote API.
func (r *CreateHaGroupAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config CreateHaGroupActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *CreateHaGroupAction) invokeRemote(ctx context.Context, config *CreateHaGroupActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/fmHa"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_create_ha_group", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_create_ha_group", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_create_ha_group", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_create_ha_group", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201 || httpResp.StatusCode == 202) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_create_ha_group", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_create_ha_group", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_create_ha_group", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_create_ha_group", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_create_ha_group", "Entity Already Exists. See errors payload for details")
			return
		case 417:
			resp.Diagnostics.AddError("Error invoking gigavuecore_create_ha_group", "Deployment failed. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_create_ha_group", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_create_ha_group", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_create_ha_group", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_create_ha_group", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *CreateHaGroupAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
