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
var _ action.Action = (*UpdateDeviceCardAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*UpdateDeviceCardAction)(nil)

// UpdateDeviceCardAction is the generated Terraform action implementation.
type UpdateDeviceCardAction struct {
	client *client.Client
}

// UpdateDeviceCardActionModel describes the action configuration shape.
type UpdateDeviceCardActionModel struct {
	AdminStatus          types.String `tfsdk:"admin_status" json:"adminStatus"`
	AlarmBufferThreshold types.Int64  `tfsdk:"alarm_buffer_threshold" json:"alarmBufferThreshold"`
	BodySlotId           types.String `tfsdk:"body_slot_id" json:"slotId"`
	ClusterId            types.String `tfsdk:"cluster_id"`
	FabricHashAdv        types.Bool   `tfsdk:"fabric_hash_adv" json:"fabricHashAdv"`
	FilterTemplate       types.String `tfsdk:"filter_template" json:"filterTemplate"`
	Mode                 types.String `tfsdk:"mode"`
	NodeId               types.String `tfsdk:"node_id"`
	PldUpgrade           types.Bool   `tfsdk:"pld_upgrade" json:"pldUpgrade"`
	SlotId               types.String `tfsdk:"slot_id"`
}

// NewUpdateDeviceCardAction returns a new instance of the generated action.
func NewUpdateDeviceCardAction() action.Action {
	return &UpdateDeviceCardAction{}
}

// Metadata returns the action type name.
func (r *UpdateDeviceCardAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_update_device_card"
}

// Schema returns the action schema.
func (r *UpdateDeviceCardAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Update Device Card configuration", Attributes: map[string]schema.Attribute{"admin_status": schema.StringAttribute{MarkdownDescription: "up: bring card up; down: shutdown the card", Optional: true}, "alarm_buffer_threshold": schema.Int64Attribute{MarkdownDescription: "card micro burst threshold", Optional: true}, "body_slot_id": schema.StringAttribute{MarkdownDescription: "device card slot id. used to identify target card. not updatable", Required: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID. Either 'clusterId' or 'nodeId' is required", Required: true}, "fabric_hash_adv": schema.BoolAttribute{MarkdownDescription: "Advanced Fabric Hash. Supported only for Q02X32/Q08 cards", Optional: true}, "filter_template": schema.StringAttribute{MarkdownDescription: "alias of filter template or 'defaults'.'defaults' is special alias for predefined filter templates", Optional: true}, "mode": schema.StringAttribute{Optional: true}, "node_id": schema.StringAttribute{MarkdownDescription: "ID of the target device", Optional: true}, "pld_upgrade": schema.BoolAttribute{MarkdownDescription: "Upgrade PLD image", Optional: true}, "slot_id": schema.StringAttribute{MarkdownDescription: "Device card slot ID", Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *UpdateDeviceCardAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UpdateDeviceCardActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *UpdateDeviceCardAction) invokeRemote(ctx context.Context, config *UpdateDeviceCardActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/inventory/chassis/cards/{slotId}"
	reqPath = strings.ReplaceAll(reqPath, "{slotId}", url.PathEscape(config.SlotId.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_device_card", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_device_card", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPatch, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_device_card", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	if !config.NodeId.IsNull() {
		query.Set("nodeId", config.NodeId.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_device_card", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_device_card", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_device_card", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_device_card", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_device_card", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_device_card", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_device_card", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_device_card", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_update_device_card", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_device_card", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *UpdateDeviceCardAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
