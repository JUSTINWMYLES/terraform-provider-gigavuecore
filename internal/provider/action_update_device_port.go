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
var _ action.Action = (*UpdateDevicePortAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*UpdateDevicePortAction)(nil)

// UpdateDevicePortAction is the generated Terraform action implementation.
type UpdateDevicePortAction struct {
	client *client.Client
}

// UpdateDevicePortActionModel describes the action configuration shape.
type UpdateDevicePortActionModel struct {
	AdminStatus  types.String `tfsdk:"admin_status" json:"adminStatus"`
	Alias        types.String `tfsdk:"alias"`
	AutoNeg      types.Bool   `tfsdk:"auto_neg" json:"autoNeg"`
	BodyPortId   types.String `tfsdk:"body_port_id" json:"portId"`
	BreakoutMode types.String `tfsdk:"breakout_mode" json:"breakoutMode"`
	CableLength  types.String `tfsdk:"cable_length" json:"cableLength"`
	ClusterId    types.String `tfsdk:"cluster_id"`
	Comment      types.String `tfsdk:"comment"`
	ConfigSpeed  types.String `tfsdk:"config_speed" json:"configSpeed"`
	Duplex       types.String `tfsdk:"duplex"`
	ForceLinkUp  types.Bool   `tfsdk:"force_link_up" json:"forceLinkUp"`
	Mtu          types.Int64  `tfsdk:"mtu"`
	PortId       types.String `tfsdk:"port_id"`
	PortType     types.String `tfsdk:"port_type" json:"portType"`
	Ude          types.Object `tfsdk:"ude"`
}

// NewUpdateDevicePortAction returns a new instance of the generated action.
func NewUpdateDevicePortAction() action.Action {
	return &UpdateDevicePortAction{}
}

// Metadata returns the action type name.
func (r *UpdateDevicePortAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_update_device_port"
}

// Schema returns the action schema.
func (r *UpdateDevicePortAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Update Device Port configuration", Attributes: map[string]schema.Attribute{"admin_status": schema.StringAttribute{Optional: true}, "alias": schema.StringAttribute{MarkdownDescription: "device port alias", Optional: true}, "auto_neg": schema.BoolAttribute{Optional: true}, "body_port_id": schema.StringAttribute{MarkdownDescription: "device port id. used to identify target port.", Required: true}, "breakout_mode": schema.StringAttribute{MarkdownDescription: "4x = 4x10G; 2q = 2x40G", Optional: true}, "cable_length": schema.StringAttribute{MarkdownDescription: "Attached cable length in meter", Optional: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "comment": schema.StringAttribute{Optional: true}, "config_speed": schema.StringAttribute{Optional: true}, "duplex": schema.StringAttribute{Optional: true}, "force_link_up": schema.BoolAttribute{Optional: true}, "mtu": schema.Int64Attribute{Optional: true}, "port_id": schema.StringAttribute{MarkdownDescription: "Device Port ID (format: boxId_slotId_port, example: 1_1_c1)", Required: true}, "port_type": schema.StringAttribute{Optional: true}, "ude": schema.SingleNestedAttribute{MarkdownDescription: "Unidirectional Ethernet", Optional: true, Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{MarkdownDescription: "Only applicable if 100g-bidi is detected", Optional: true}}}}}
}

// Invoke executes the action against the remote API.
func (r *UpdateDevicePortAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UpdateDevicePortActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *UpdateDevicePortAction) invokeRemote(ctx context.Context, config *UpdateDevicePortActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/inventory/ports/{portId}"
	reqPath = strings.ReplaceAll(reqPath, "{portId}", url.PathEscape(config.PortId.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_device_port", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_device_port", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPatch, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_device_port", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_device_port", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_device_port", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_device_port", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_device_port", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_device_port", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_device_port", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_device_port", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_device_port", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_update_device_port", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_device_port", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *UpdateDevicePortAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
