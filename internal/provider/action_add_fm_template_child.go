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
var _ action.Action = (*AddFmTemplateChildAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*AddFmTemplateChildAction)(nil)

// AddFmTemplateChildAction is the generated Terraform action implementation.
type AddFmTemplateChildAction struct {
	client *client.Client
}

// AddFmTemplateChildActionModel describes the action configuration shape.
type AddFmTemplateChildActionModel struct {
	BodyChildType    types.String  `tfsdk:"body_child_type" json:"childType"`
	BodyConfigType   types.String  `tfsdk:"body_config_type" json:"configType"`
	ChildType        types.String  `tfsdk:"child_type"`
	Config           types.Dynamic `tfsdk:"config"`
	ConfigLevel      types.String  `tfsdk:"config_level" json:"configLevel"`
	ConfigLevelValue types.List    `tfsdk:"config_level_value" json:"configLevelValue"`
	ConfigType       types.String  `tfsdk:"config_type"`
}

// NewAddFmTemplateChildAction returns a new instance of the generated action.
func NewAddFmTemplateChildAction() action.Action {
	return &AddFmTemplateChildAction{}
}

// Metadata returns the action type name.
func (r *AddFmTemplateChildAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_add_fm_template_child"
}

// Schema returns the action schema.
func (r *AddFmTemplateChildAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "new in FM 5.13.01", Attributes: map[string]schema.Attribute{"body_child_type": schema.StringAttribute{MarkdownDescription: "Child Type of the FM template Configuration", Optional: true}, "body_config_type": schema.StringAttribute{MarkdownDescription: "Configuration Type of the FM template", Optional: true}, "child_type": schema.StringAttribute{MarkdownDescription: "childType of the fm template", Required: true}, "config": schema.DynamicAttribute{Optional: true}, "config_level": schema.StringAttribute{MarkdownDescription: "Scope of the applied FM template", Optional: true}, "config_level_value": schema.ListAttribute{Optional: true, ElementType: types.StringType}, "config_type": schema.StringAttribute{MarkdownDescription: "configType of the fm template", Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *AddFmTemplateChildAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config AddFmTemplateChildActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *AddFmTemplateChildAction) invokeRemote(ctx context.Context, config *AddFmTemplateChildActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/fm/templates/{configType}/{childType}"
	reqPath = strings.ReplaceAll(reqPath, "{configType}", url.PathEscape(config.ConfigType.ValueString()))
	reqPath = strings.ReplaceAll(reqPath, "{childType}", url.PathEscape(config.ChildType.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_add_fm_template_child", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_add_fm_template_child", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_add_fm_template_child", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_add_fm_template_child", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_fm_template_child", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_fm_template_child", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_fm_template_child", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_fm_template_child", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_fm_template_child", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_fm_template_child", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_fm_template_child", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_add_fm_template_child", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_fm_template_child", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *AddFmTemplateChildAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
