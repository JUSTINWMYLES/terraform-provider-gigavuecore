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
var _ action.Action = (*RedefineSnmpServerCommunityConfigAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*RedefineSnmpServerCommunityConfigAction)(nil)

// RedefineSnmpServerCommunityConfigAction is the generated Terraform action implementation.
type RedefineSnmpServerCommunityConfigAction struct {
	client *client.Client
}

// RedefineSnmpServerCommunityConfigActionModel describes the action configuration shape.
type RedefineSnmpServerCommunityConfigActionModel struct {
	ClusterId             types.String `tfsdk:"cluster_id"`
	CommunityStrings      types.List   `tfsdk:"community_strings" json:"communityStrings"`
	EnableCommunityAuth   types.Bool   `tfsdk:"enable_community_auth" json:"enableCommunityAuth"`
	EnableCommunityAuthV1 types.Bool   `tfsdk:"enable_community_auth_v1" json:"enableCommunityAuthV1"`
	EnableMultiCommunity  types.Bool   `tfsdk:"enable_multi_community" json:"enableMultiCommunity"`
}

// NewRedefineSnmpServerCommunityConfigAction returns a new instance of the generated action.
func NewRedefineSnmpServerCommunityConfigAction() action.Action {
	return &RedefineSnmpServerCommunityConfigAction{}
}

// Metadata returns the action type name.
func (r *RedefineSnmpServerCommunityConfigAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_redefine_snmp_server_community_config"
}

// Schema returns the action schema.
func (r *RedefineSnmpServerCommunityConfigAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Redefine Snmp Server Community config", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "community_strings": schema.ListAttribute{MarkdownDescription: "community string[s] used to connect to this node using SNMP. The default value is 'public'. If 'enableMultiCommunity' is enabled, multiple community strings for the node are allowed", Optional: true, ElementType: types.StringType}, "enable_community_auth": schema.BoolAttribute{MarkdownDescription: "turn on community-based authentication for the system", Optional: true}, "enable_community_auth_v1": schema.BoolAttribute{MarkdownDescription: "turn on community-based authentication for the SNMP v1", Optional: true}, "enable_multi_community": schema.BoolAttribute{MarkdownDescription: "allow configuration of multiple communities", Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *RedefineSnmpServerCommunityConfigAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config RedefineSnmpServerCommunityConfigActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *RedefineSnmpServerCommunityConfigAction) invokeRemote(ctx context.Context, config *RedefineSnmpServerCommunityConfigActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/system/snmp/community"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_snmp_server_community_config", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_snmp_server_community_config", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_snmp_server_community_config", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_snmp_server_community_config", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_snmp_server_community_config", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_snmp_server_community_config", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_snmp_server_community_config", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_snmp_server_community_config", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_snmp_server_community_config", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_snmp_server_community_config", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_snmp_server_community_config", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_snmp_server_community_config", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_snmp_server_community_config", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *RedefineSnmpServerCommunityConfigAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
