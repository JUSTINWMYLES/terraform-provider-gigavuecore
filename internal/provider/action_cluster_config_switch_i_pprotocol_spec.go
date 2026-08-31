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
var _ action.Action = (*ClusterConfigSwitchIPprotocolSpecAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*ClusterConfigSwitchIPprotocolSpecAction)(nil)

// ClusterConfigSwitchIPprotocolSpecAction is the generated Terraform action implementation.
type ClusterConfigSwitchIPprotocolSpecAction struct {
	client *client.Client
}

// ClusterConfigSwitchIPprotocolSpecActionModel describes the action configuration shape.
type ClusterConfigSwitchIPprotocolSpecActionModel struct {
	ClusterId          types.String `tfsdk:"cluster_id"`
	ClusterPrimaryIp   types.String `tfsdk:"cluster_primary_ip" json:"clusterPrimaryIp"`
	ClusterSecondaryIp types.String `tfsdk:"cluster_secondary_ip" json:"clusterSecondaryIp"`
	ClusterVip         types.String `tfsdk:"cluster_vip" json:"clusterVIP"`
	ClusterVipMaskLen  types.Int64  `tfsdk:"cluster_vip_mask_len" json:"clusterVipMaskLen"`
	IpProtocol         types.String `tfsdk:"ip_protocol" json:"ipProtocol"`
}

// NewClusterConfigSwitchIPprotocolSpecAction returns a new instance of the generated action.
func NewClusterConfigSwitchIPprotocolSpecAction() action.Action {
	return &ClusterConfigSwitchIPprotocolSpecAction{}
}

// Metadata returns the action type name.
func (r *ClusterConfigSwitchIPprotocolSpecAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_cluster_config_switch_i_pprotocol_spec"
}

// Schema returns the action schema.
func (r *ClusterConfigSwitchIPprotocolSpecAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "switch ip protocol for physical cluster nodes", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "The requested cluster ID", Required: true}, "cluster_primary_ip": schema.StringAttribute{MarkdownDescription: "Cluster primary IP", Optional: true}, "cluster_secondary_ip": schema.StringAttribute{MarkdownDescription: "Cluster secondary ip", Optional: true}, "cluster_vip": schema.StringAttribute{MarkdownDescription: "Cluster VIP", Optional: true}, "cluster_vip_mask_len": schema.Int64Attribute{MarkdownDescription: "cluster leader virtual ip mask length, valid and required when 'clusterVip' is specified", Optional: true}, "ip_protocol": schema.StringAttribute{MarkdownDescription: "ip protocol", Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *ClusterConfigSwitchIPprotocolSpecAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config ClusterConfigSwitchIPprotocolSpecActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *ClusterConfigSwitchIPprotocolSpecAction) invokeRemote(ctx context.Context, config *ClusterConfigSwitchIPprotocolSpecActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/clusterConfig/ipVerSwitch"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_switch_i_pprotocol_spec", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_switch_i_pprotocol_spec", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_switch_i_pprotocol_spec", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_switch_i_pprotocol_spec", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_switch_i_pprotocol_spec", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_switch_i_pprotocol_spec", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_switch_i_pprotocol_spec", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_switch_i_pprotocol_spec", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_switch_i_pprotocol_spec", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_switch_i_pprotocol_spec", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_switch_i_pprotocol_spec", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_switch_i_pprotocol_spec", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_cluster_config_switch_i_pprotocol_spec", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *ClusterConfigSwitchIPprotocolSpecAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
