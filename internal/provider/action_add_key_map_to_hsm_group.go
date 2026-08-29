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
var _ action.Action = (*AddKeyMapToHsmGroupAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*AddKeyMapToHsmGroupAction)(nil)

// AddKeyMapToHsmGroupAction is the generated Terraform action implementation.
type AddKeyMapToHsmGroupAction struct {
	client *client.Client
}

// AddKeyMapToHsmGroupActionModel describes the action configuration shape.
type AddKeyMapToHsmGroupActionModel struct {
	Alias      types.String `tfsdk:"alias"`
	ClusterId  types.String `tfsdk:"cluster_id"`
	HsmKeyMaps types.List   `tfsdk:"hsm_key_maps" json:"hsmKeyMaps"`
}

// NewAddKeyMapToHsmGroupAction returns a new instance of the generated action.
func NewAddKeyMapToHsmGroupAction() action.Action {
	return &AddKeyMapToHsmGroupAction{}
}

// Metadata returns the action type name.
func (r *AddKeyMapToHsmGroupAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_add_key_map_to_hsm_group"
}

// Schema returns the action schema.
func (r *AddKeyMapToHsmGroupAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Add keymap to HSM Group", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "alias of target HSM Group", Required: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target cluster ID.", Required: true}, "hsm_key_maps": schema.ListNestedAttribute{Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"address": schema.StringAttribute{MarkdownDescription: "IPv4 address of the SSL endpoint (server)", Required: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "id of the defining cluster", Optional: true}, "key_name": schema.StringAttribute{MarkdownDescription: "key-name", Optional: true}, "key_token": schema.StringAttribute{MarkdownDescription: "key-token", Optional: true}, "port": schema.Int64Attribute{MarkdownDescription: "Port of the SSL endpoint (server). Value of 0 indicates any port", Optional: true}, "rfs_match": schema.StringAttribute{MarkdownDescription: "Is there a matching RFS key, yes or no", Optional: true}, "rule_id": schema.StringAttribute{MarkdownDescription: "Rule Id", Optional: true}}}}}}
}

// Invoke executes the action against the remote API.
func (r *AddKeyMapToHsmGroupAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config AddKeyMapToHsmGroupActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *AddKeyMapToHsmGroupAction) invokeRemote(ctx context.Context, config *AddKeyMapToHsmGroupActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/hsmGroup/{alias}/keyMap"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_add_key_map_to_hsm_group", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_add_key_map_to_hsm_group", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_add_key_map_to_hsm_group", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_add_key_map_to_hsm_group", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_key_map_to_hsm_group", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_key_map_to_hsm_group", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_key_map_to_hsm_group", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_key_map_to_hsm_group", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_key_map_to_hsm_group", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_key_map_to_hsm_group", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_key_map_to_hsm_group", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_add_key_map_to_hsm_group", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_key_map_to_hsm_group", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *AddKeyMapToHsmGroupAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
