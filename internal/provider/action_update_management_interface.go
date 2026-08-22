package provider

import (
	"bytes"
	"context"
	"encoding/json"
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
var _ action.Action = (*UpdateManagementInterfaceAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*UpdateManagementInterfaceAction)(nil)

// UpdateManagementInterfaceAction is the generated Terraform action implementation.
type UpdateManagementInterfaceAction struct {
	client *client.Client
}

// UpdateManagementInterfaceActionModel describes the action configuration shape.
type UpdateManagementInterfaceActionModel struct {
	BodyBoxId         types.Int64  `tfsdk:"body_box_id" json:"boxId"`
	BodyClusterId     types.String `tfsdk:"body_cluster_id" json:"clusterId"`
	BodyInterfaceName types.String `tfsdk:"body_interface_name" json:"interfaceName"`
	BoxId             types.Int64  `tfsdk:"box_id"`
	ClusterId         types.String `tfsdk:"cluster_id"`
	DiscoveryProtocol types.String `tfsdk:"discovery_protocol" json:"discoveryProtocol"`
	GArp              types.Bool   `tfsdk:"g_arp" json:"gArp"`
	InterfaceName     types.String `tfsdk:"interface_name"`
}

// NewUpdateManagementInterfaceAction returns a new instance of the generated action.
func NewUpdateManagementInterfaceAction() action.Action {
	return &UpdateManagementInterfaceAction{}
}

// Metadata returns the action type name.
func (r *UpdateManagementInterfaceAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_update_management_interface"
}

// Schema returns the action schema.
func (r *UpdateManagementInterfaceAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "since FM 5.8", Attributes: map[string]schema.Attribute{"body_box_id": schema.Int64Attribute{Required: true}, "body_cluster_id": schema.StringAttribute{MarkdownDescription: "id of the defining cluster", Optional: true}, "body_interface_name": schema.StringAttribute{MarkdownDescription: "Management Interface Name", Required: true}, "box_id": schema.Int64Attribute{MarkdownDescription: "boxId", Required: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "discovery_protocol": schema.StringAttribute{Required: true}, "g_arp": schema.BoolAttribute{MarkdownDescription: "Enable or disable Gratuitous ARP", Required: true}, "interface_name": schema.StringAttribute{MarkdownDescription: "Interface Name", Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *UpdateManagementInterfaceAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UpdateManagementInterfaceActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *UpdateManagementInterfaceAction) invokeRemote(ctx context.Context, config *UpdateManagementInterfaceActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/system/interfaces/mgmt/{boxId}/{interfaceName}"
	reqPath = strings.ReplaceAll(reqPath, "{boxId}", url.PathEscape(strconv.FormatInt(config.BoxId.ValueInt64(), 10)))
	reqPath = strings.ReplaceAll(reqPath, "{interfaceName}", url.PathEscape(config.InterfaceName.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_management_interface", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_management_interface", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPatch, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_management_interface", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.BodyClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_management_interface", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_management_interface", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_management_interface", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_management_interface", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_management_interface", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_management_interface", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_management_interface", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_management_interface", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_update_management_interface", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_management_interface", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *UpdateManagementInterfaceAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
