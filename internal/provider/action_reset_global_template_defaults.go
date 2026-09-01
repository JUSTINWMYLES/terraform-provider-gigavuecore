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
var _ action.Action = (*ResetGlobalTemplateDefaultsAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*ResetGlobalTemplateDefaultsAction)(nil)

// ResetGlobalTemplateDefaultsAction is the generated Terraform action implementation.
type ResetGlobalTemplateDefaultsAction struct {
	client *client.Client
}

// ResetGlobalTemplateDefaultsActionModel describes the action configuration shape.
type ResetGlobalTemplateDefaultsActionModel struct {
	ConfigType types.String `tfsdk:"config_type"`
}

// NewResetGlobalTemplateDefaultsAction returns a new instance of the generated action.
func NewResetGlobalTemplateDefaultsAction() action.Action {
	return &ResetGlobalTemplateDefaultsAction{}
}

// Metadata returns the action type name.
func (r *ResetGlobalTemplateDefaultsAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_reset_global_template_defaults"
}

// Schema returns the action schema.
func (r *ResetGlobalTemplateDefaultsAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "new in FM 5.14", Attributes: map[string]schema.Attribute{"config_type": schema.StringAttribute{MarkdownDescription: "configType of the fm template", Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *ResetGlobalTemplateDefaultsAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config ResetGlobalTemplateDefaultsActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *ResetGlobalTemplateDefaultsAction) invokeRemote(ctx context.Context, config *ResetGlobalTemplateDefaultsActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/fm/templates/global/{configType}/reset"
	reqPath = strings.ReplaceAll(reqPath, "{configType}", url.PathEscape(config.ConfigType.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_reset_global_template_defaults", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_reset_global_template_defaults", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_reset_global_template_defaults", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_reset_global_template_defaults", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_reset_global_template_defaults", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_reset_global_template_defaults", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_reset_global_template_defaults", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_reset_global_template_defaults", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_reset_global_template_defaults", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_reset_global_template_defaults", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_reset_global_template_defaults", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *ResetGlobalTemplateDefaultsAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
