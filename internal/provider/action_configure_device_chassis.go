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
var _ action.Action = (*ConfigureDeviceChassisAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*ConfigureDeviceChassisAction)(nil)

// ConfigureDeviceChassisAction is the generated Terraform action implementation.
type ConfigureDeviceChassisAction struct {
	client *client.Client
}

// ConfigureDeviceChassisActionModel describes the action configuration shape.
type ConfigureDeviceChassisActionModel struct {
	BoxId        types.Int64   `tfsdk:"box_id" json:"boxId"`
	ChassisType  types.String  `tfsdk:"chassis_type" json:"chassisType"`
	ClusterId    types.String  `tfsdk:"cluster_id"`
	Gdp          types.Bool    `tfsdk:"gdp"`
	L2GreId      types.Int64   `tfsdk:"l2_gre_id" json:"l2greId"`
	LeafConfig   types.Dynamic `tfsdk:"leaf_config" json:"leafConfig"`
	Mode         types.String  `tfsdk:"mode"`
	NodeId       types.String  `tfsdk:"node_id"`
	SerialNumber types.String  `tfsdk:"serial_number" json:"serialNumber"`
	VxlanId      types.Int64   `tfsdk:"vxlan_id" json:"vxlanId"`
}

// NewConfigureDeviceChassisAction returns a new instance of the generated action.
func NewConfigureDeviceChassisAction() action.Action {
	return &ConfigureDeviceChassisAction{}
}

// Metadata returns the action type name.
func (r *ConfigureDeviceChassisAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_configure_device_chassis"
}

// Schema returns the action schema.
func (r *ConfigureDeviceChassisAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Configure Device Chassis", Attributes: map[string]schema.Attribute{"box_id": schema.Int64Attribute{MarkdownDescription: "Configure the chassis Box ID value", Required: true}, "chassis_type": schema.StringAttribute{MarkdownDescription: "Specify chassis type (to provision offline) valid values: hb1, hc1, hc2, hc2-v2, hc3, hd4-ccv1, hd4-ccv2, hd8-ccv1, hd8-ccv2, ly2r, ta1, ta10, ta10a, ta40, itac, tacx", Optional: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID. Either 'clusterId' or 'nodeId' is required", Required: true}, "gdp": schema.BoolAttribute{MarkdownDescription: "enable/disable GDP for chassis", Optional: true}, "l2_gre_id": schema.Int64Attribute{MarkdownDescription: "for type l2gre maximum value is 4294967295 , Value of 0 disables the vxlanId", Optional: true}, "leaf_config": schema.DynamicAttribute{MarkdownDescription: "Configuration of Leaf node for Spine-Link", Optional: true}, "mode": schema.StringAttribute{MarkdownDescription: "Use of 100G ports requires chassis mode value to be 100G. 100G applicable to HC2-v2 only. 100GLeft, Right refers to left, right side chassis bank.", Optional: true}, "node_id": schema.StringAttribute{MarkdownDescription: "ID of the target device", Optional: true}, "serial_number": schema.StringAttribute{MarkdownDescription: "Specify Chassis Serial Number (Defaults to the local chassis)", Optional: true}, "vxlan_id": schema.Int64Attribute{MarkdownDescription: "for type vxlan maximum value is 16777215 , Value of 0 disables the vxlanId", Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *ConfigureDeviceChassisAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config ConfigureDeviceChassisActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *ConfigureDeviceChassisAction) invokeRemote(ctx context.Context, config *ConfigureDeviceChassisActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/inventory/chassis/configure"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_configure_device_chassis", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_configure_device_chassis", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_configure_device_chassis", fmt.Sprintf("Could not build request: %s", err))
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
		resp.Diagnostics.AddError("Error invoking gigavuecore_configure_device_chassis", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_configure_device_chassis", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_configure_device_chassis", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_configure_device_chassis", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_configure_device_chassis", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_configure_device_chassis", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_configure_device_chassis", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_configure_device_chassis", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_configure_device_chassis", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_configure_device_chassis", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *ConfigureDeviceChassisAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
