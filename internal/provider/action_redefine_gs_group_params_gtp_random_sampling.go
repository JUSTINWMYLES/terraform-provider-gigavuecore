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
var _ action.Action = (*RedefineGsGroupParamsGtpRandomSamplingAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*RedefineGsGroupParamsGtpRandomSamplingAction)(nil)

// RedefineGsGroupParamsGtpRandomSamplingAction is the generated Terraform action implementation.
type RedefineGsGroupParamsGtpRandomSamplingAction struct {
	client *client.Client
}

// RedefineGsGroupParamsGtpRandomSamplingActionModel describes the action configuration shape.
type RedefineGsGroupParamsGtpRandomSamplingActionModel struct {
	Alias    types.String `tfsdk:"alias"`
	Enabled  types.Bool   `tfsdk:"enabled"`
	Interval types.Int64  `tfsdk:"interval"`
}

// NewRedefineGsGroupParamsGtpRandomSamplingAction returns a new instance of the generated action.
func NewRedefineGsGroupParamsGtpRandomSamplingAction() action.Action {
	return &RedefineGsGroupParamsGtpRandomSamplingAction{}
}

// Metadata returns the action type name.
func (r *RedefineGsGroupParamsGtpRandomSamplingAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_redefine_gs_group_params_gtp_random_sampling"
}

// Schema returns the action schema.
func (r *RedefineGsGroupParamsGtpRandomSamplingAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "new in H 5.7", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "alias of the target GS Group", Required: true}, "enabled": schema.BoolAttribute{MarkdownDescription: "When enabled, sampling of subscriber's sessions happens in random fashion", Optional: true}, "interval": schema.Int64Attribute{MarkdownDescription: "Rotation Interval in multiples of 12 (hrs)", Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *RedefineGsGroupParamsGtpRandomSamplingAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config RedefineGsGroupParamsGtpRandomSamplingActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *RedefineGsGroupParamsGtpRandomSamplingAction) invokeRemote(ctx context.Context, config *RedefineGsGroupParamsGtpRandomSamplingActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/gsGroups/{alias}/params/gtpRandomSampling"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_params_gtp_random_sampling", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_params_gtp_random_sampling", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_params_gtp_random_sampling", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_params_gtp_random_sampling", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_params_gtp_random_sampling", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_params_gtp_random_sampling", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_params_gtp_random_sampling", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_params_gtp_random_sampling", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_params_gtp_random_sampling", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_params_gtp_random_sampling", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_params_gtp_random_sampling", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_params_gtp_random_sampling", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_params_gtp_random_sampling", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *RedefineGsGroupParamsGtpRandomSamplingAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
