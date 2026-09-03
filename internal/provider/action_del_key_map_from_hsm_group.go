package provider

import (
	"context"
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
var _ action.Action = (*DelKeyMapFromHsmGroupAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*DelKeyMapFromHsmGroupAction)(nil)

// DelKeyMapFromHsmGroupAction is the generated Terraform action implementation.
type DelKeyMapFromHsmGroupAction struct {
	client *client.Client
}

// DelKeyMapFromHsmGroupActionModel describes the action configuration shape.
type DelKeyMapFromHsmGroupActionModel struct {
	Alias     types.String `tfsdk:"alias"`
	ClusterId types.String `tfsdk:"cluster_id"`
	RuleId    types.String `tfsdk:"rule_id"`
}

// NewDelKeyMapFromHsmGroupAction returns a new instance of the generated action.
func NewDelKeyMapFromHsmGroupAction() action.Action {
	return &DelKeyMapFromHsmGroupAction{}
}

// Metadata returns the action type name.
func (r *DelKeyMapFromHsmGroupAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_del_key_map_from_hsm_group"
}

// Schema returns the action schema.
func (r *DelKeyMapFromHsmGroupAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Delete keymaps from HSM Group", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "alias of target HSM Group", Required: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target cluster ID.", Required: true}, "rule_id": schema.StringAttribute{MarkdownDescription: "ruleId of keymap to delete, 0 = all", Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *DelKeyMapFromHsmGroupAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config DelKeyMapFromHsmGroupActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *DelKeyMapFromHsmGroupAction) invokeRemote(ctx context.Context, config *DelKeyMapFromHsmGroupActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/hsmGroup/{alias}/keyMap/{ruleId}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	reqPath = strings.ReplaceAll(reqPath, "{ruleId}", url.PathEscape(config.RuleId.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_del_key_map_from_hsm_group", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_del_key_map_from_hsm_group", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_del_key_map_from_hsm_group", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_del_key_map_from_hsm_group", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_del_key_map_from_hsm_group", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_del_key_map_from_hsm_group", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_del_key_map_from_hsm_group", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_del_key_map_from_hsm_group", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_del_key_map_from_hsm_group", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_del_key_map_from_hsm_group", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_del_key_map_from_hsm_group", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *DelKeyMapFromHsmGroupAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
