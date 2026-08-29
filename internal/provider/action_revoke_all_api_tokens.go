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
var _ action.Action = (*RevokeAllApiTokensAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*RevokeAllApiTokensAction)(nil)

// RevokeAllApiTokensAction is the generated Terraform action implementation.
type RevokeAllApiTokensAction struct {
	client *client.Client
}

// RevokeAllApiTokensActionModel describes the action configuration shape.
type RevokeAllApiTokensActionModel struct {
	Context                types.Object  `tfsdk:"context"`
	FmApiTokenUserEntities types.Dynamic `tfsdk:"fm_api_token_user_entities" json:"fmApiTokenUserEntities"`
}

// NewRevokeAllApiTokensAction returns a new instance of the generated action.
func NewRevokeAllApiTokensAction() action.Action {
	return &RevokeAllApiTokensAction{}
}

// Metadata returns the action type name.
func (r *RevokeAllApiTokensAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_revoke_all_api_tokens"
}

// Schema returns the action schema.
func (r *RevokeAllApiTokensAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Revoke token for FM Users", Attributes: map[string]schema.Attribute{"context": schema.SingleNestedAttribute{MarkdownDescription: "Gigamon query result context", Optional: true, Attributes: map[string]schema.Attribute{"page_no": schema.Int64Attribute{MarkdownDescription: "page number of the returned result set", Optional: true}, "page_size": schema.Int64Attribute{MarkdownDescription: "page size of the returned result set", Optional: true}, "sort": schema.ListAttribute{MarkdownDescription: "sorting info of the returned result set. list of fields in the array indicate sorting order", Optional: true, ElementType: types.StringType}, "total_items": schema.Int64Attribute{MarkdownDescription: "total number of items in the queried entity type", Required: true}}}, "fm_api_token_user_entities": schema.DynamicAttribute{Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *RevokeAllApiTokensAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config RevokeAllApiTokensActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *RevokeAllApiTokensAction) invokeRemote(ctx context.Context, config *RevokeAllApiTokensActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/tokens/revoke"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_revoke_all_api_tokens", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_revoke_all_api_tokens", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_revoke_all_api_tokens", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_revoke_all_api_tokens", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_revoke_all_api_tokens", "Invalid request. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_revoke_all_api_tokens", "Access Denied. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_revoke_all_api_tokens", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_revoke_all_api_tokens", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_revoke_all_api_tokens", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_revoke_all_api_tokens", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *RevokeAllApiTokensAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
