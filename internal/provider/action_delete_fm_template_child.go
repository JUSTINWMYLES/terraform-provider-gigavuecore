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
var _ action.Action = (*DeleteFmTemplateChildAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*DeleteFmTemplateChildAction)(nil)

// DeleteFmTemplateChildAction is the generated Terraform action implementation.
type DeleteFmTemplateChildAction struct {
	client *client.Client
}

// DeleteFmTemplateChildActionModel describes the action configuration shape.
type DeleteFmTemplateChildActionModel struct {
	BodyChildType    types.String `tfsdk:"body_child_type" json:"childType"`
	BodyConfigType   types.String `tfsdk:"body_config_type" json:"configType"`
	ChildType        types.String `tfsdk:"child_type"`
	Config           types.Object `tfsdk:"config"`
	ConfigLevel      types.String `tfsdk:"config_level" json:"configLevel"`
	ConfigLevelValue types.List   `tfsdk:"config_level_value" json:"configLevelValue"`
	ConfigType       types.String `tfsdk:"config_type"`
}

// NewDeleteFmTemplateChildAction returns a new instance of the generated action.
func NewDeleteFmTemplateChildAction() action.Action {
	return &DeleteFmTemplateChildAction{}
}

// Metadata returns the action type name.
func (r *DeleteFmTemplateChildAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_delete_fm_template_child"
}

// Schema returns the action schema.
func (r *DeleteFmTemplateChildAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Delete child entry from template configuration", Attributes: map[string]schema.Attribute{"body_child_type": schema.StringAttribute{MarkdownDescription: "Child Type of the FM template Configuration", Optional: true}, "body_config_type": schema.StringAttribute{MarkdownDescription: "Configuration Type of the FM template", Optional: true}, "child_type": schema.StringAttribute{MarkdownDescription: "childType of the fm template", Required: true}, "config": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"ldap_servers": schema.ListNestedAttribute{MarkdownDescription: "List of LDAP Servers. Available for ConfigType LDAP_SERVERS_TEMPLATE ChildType ldapServers", Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"order": schema.StringAttribute{MarkdownDescription: "The order in which the server is to be reached. 1 means server will be contacted first", Optional: true}, "server_address": schema.StringAttribute{MarkdownDescription: "ipv4 or ipv6 or hostname. Specifies address of the LDAP server where authentication requests will be sent", Required: true}}}}, "remote_map_table": schema.ListNestedAttribute{MarkdownDescription: "List of LDAP User Group Mapping. Available for ConfigType LDAP_SYSTEM_CONFIG_TEMPLATE ChildType remoteMapTable", Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"local_account_name": schema.StringAttribute{MarkdownDescription: "Specifies local account to which remote base-dn needs to be mapped.", Optional: true}, "remote_base_dn": schema.StringAttribute{MarkdownDescription: "Specifies the base-dn of the remote user group to be mapped to local account", Optional: true}}}}}}, "config_level": schema.StringAttribute{MarkdownDescription: "Scope of the applied FM template", Optional: true}, "config_level_value": schema.ListAttribute{Optional: true, ElementType: types.StringType}, "config_type": schema.StringAttribute{MarkdownDescription: "configType of the fm template", Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *DeleteFmTemplateChildAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config DeleteFmTemplateChildActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *DeleteFmTemplateChildAction) invokeRemote(ctx context.Context, config *DeleteFmTemplateChildActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/fm/templates/{configType}/{childType}"
	reqPath = strings.ReplaceAll(reqPath, "{configType}", url.PathEscape(config.ConfigType.ValueString()))
	reqPath = strings.ReplaceAll(reqPath, "{childType}", url.PathEscape(config.ChildType.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_delete_fm_template_child", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_delete_fm_template_child", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_delete_fm_template_child", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_delete_fm_template_child", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 204) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_fm_template_child", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_fm_template_child", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_fm_template_child", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_fm_template_child", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_fm_template_child", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_fm_template_child", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_fm_template_child", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_delete_fm_template_child", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_fm_template_child", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *DeleteFmTemplateChildAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
