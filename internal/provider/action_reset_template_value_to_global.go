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
var _ action.Action = (*ResetTemplateValueToGlobalAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*ResetTemplateValueToGlobalAction)(nil)

// ResetTemplateValueToGlobalAction is the generated Terraform action implementation.
type ResetTemplateValueToGlobalAction struct {
	client *client.Client
}

// ResetTemplateValueToGlobalActionModel describes the action configuration shape.
type ResetTemplateValueToGlobalActionModel struct {
	BodyConfigType   types.String  `tfsdk:"body_config_type" json:"configType"`
	Config           types.Dynamic `tfsdk:"config"`
	ConfigLevel      types.String  `tfsdk:"config_level" json:"configLevel"`
	ConfigLevelValue types.List    `tfsdk:"config_level_value" json:"configLevelValue"`
	ConfigResource   types.Dynamic `tfsdk:"config_resource" json:"configResource"`
	ConfigType       types.String  `tfsdk:"config_type"`
	Modifiable       types.Bool    `tfsdk:"modifiable"`
	RefCount         types.Int64   `tfsdk:"ref_count" json:"refCount"`
	RefObject        types.Dynamic `tfsdk:"ref_object" json:"refObject"`
	TemplateName     types.String  `tfsdk:"template_name" json:"templateName"`
	UpdateTime       types.String  `tfsdk:"update_time" json:"updateTime"`
}

// NewResetTemplateValueToGlobalAction returns a new instance of the generated action.
func NewResetTemplateValueToGlobalAction() action.Action {
	return &ResetTemplateValueToGlobalAction{}
}

// Metadata returns the action type name.
func (r *ResetTemplateValueToGlobalAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_reset_template_value_to_global"
}

// Schema returns the action schema.
func (r *ResetTemplateValueToGlobalAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "new in FM 5.14", Attributes: map[string]schema.Attribute{"body_config_type": schema.StringAttribute{MarkdownDescription: "Configuration Type of the FM template", Optional: true}, "config": schema.DynamicAttribute{Optional: true}, "config_level": schema.StringAttribute{MarkdownDescription: "Scope of the applied FM template", Optional: true}, "config_level_value": schema.ListAttribute{Optional: true, ElementType: types.StringType}, "config_resource": schema.DynamicAttribute{MarkdownDescription: "For a particular config type, there may be multiple templates with different levels. For example, cluster level the user can have more than one template and in same way in tag level there may be more than one template for each tag combination. In such cases, the configResource property can be used to contain the list of clusters associated to the template or the tag key values combination of the template. The data structure changes based on the level hence it is kept as type Object.", Optional: true}, "config_type": schema.StringAttribute{MarkdownDescription: "configType of the fm template", Required: true}, "modifiable": schema.BoolAttribute{Optional: true}, "ref_count": schema.Int64Attribute{Optional: true}, "ref_object": schema.DynamicAttribute{Optional: true}, "template_name": schema.StringAttribute{Optional: true}, "update_time": schema.StringAttribute{Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *ResetTemplateValueToGlobalAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config ResetTemplateValueToGlobalActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *ResetTemplateValueToGlobalAction) invokeRemote(ctx context.Context, config *ResetTemplateValueToGlobalActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/fm/templates/{configType}/reset"
	reqPath = strings.ReplaceAll(reqPath, "{configType}", url.PathEscape(config.ConfigType.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_reset_template_value_to_global", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_reset_template_value_to_global", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_reset_template_value_to_global", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_reset_template_value_to_global", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_reset_template_value_to_global", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_reset_template_value_to_global", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_reset_template_value_to_global", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_reset_template_value_to_global", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_reset_template_value_to_global", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_reset_template_value_to_global", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_reset_template_value_to_global", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_reset_template_value_to_global", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_reset_template_value_to_global", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *ResetTemplateValueToGlobalAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
