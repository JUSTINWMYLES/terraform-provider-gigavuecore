package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*CreateApiTokenFmSystemApiRateLimitingAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*CreateApiTokenFmSystemApiRateLimitingAction)(nil)

// CreateApiTokenFmSystemApiRateLimitingAction is the generated Terraform action implementation.
type CreateApiTokenFmSystemApiRateLimitingAction struct {
	client *client.Client
}

// CreateApiTokenFmSystemApiRateLimitingActionModel describes the action configuration shape.
type CreateApiTokenFmSystemApiRateLimitingActionModel struct {
	RequestLimitCount    types.Int64 `tfsdk:"request_limit_count" json:"requestLimitCount"`
	RequestLimitDuration types.Int64 `tfsdk:"request_limit_duration" json:"requestLimitDuration"`
}

// NewCreateApiTokenFmSystemApiRateLimitingAction returns a new instance of the generated action.
func NewCreateApiTokenFmSystemApiRateLimitingAction() action.Action {
	return &CreateApiTokenFmSystemApiRateLimitingAction{}
}

// Metadata returns the action type name.
func (r *CreateApiTokenFmSystemApiRateLimitingAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_create_api_token_fm_system_api_rate_limiting"
}

// Schema returns the action schema.
func (r *CreateApiTokenFmSystemApiRateLimitingAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Modify FM API rate limit configuration.", Attributes: map[string]schema.Attribute{"request_limit_count": schema.Int64Attribute{MarkdownDescription: "FM API Request Limit Count", Required: true}, "request_limit_duration": schema.Int64Attribute{MarkdownDescription: "FM API Request Limit Duration in seconds", Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *CreateApiTokenFmSystemApiRateLimitingAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config CreateApiTokenFmSystemApiRateLimitingActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *CreateApiTokenFmSystemApiRateLimitingAction) invokeRemote(ctx context.Context, config *CreateApiTokenFmSystemApiRateLimitingActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/fmSystem/apiRateLimiting"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_create_api_token_fm_system_api_rate_limiting", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_create_api_token_fm_system_api_rate_limiting", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_create_api_token_fm_system_api_rate_limiting", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_create_api_token_fm_system_api_rate_limiting", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_create_api_token_fm_system_api_rate_limiting", "Invalid request. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_create_api_token_fm_system_api_rate_limiting", "Access Denied. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_create_api_token_fm_system_api_rate_limiting", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_create_api_token_fm_system_api_rate_limiting", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_create_api_token_fm_system_api_rate_limiting", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_create_api_token_fm_system_api_rate_limiting", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *CreateApiTokenFmSystemApiRateLimitingAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
