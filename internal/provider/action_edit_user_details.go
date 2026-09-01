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
var _ action.Action = (*EditUserDetailsAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*EditUserDetailsAction)(nil)

// EditUserDetailsAction is the generated Terraform action implementation.
type EditUserDetailsAction struct {
	client *client.Client
}

// EditUserDetailsActionModel describes the action configuration shape.
type EditUserDetailsActionModel struct {
	BodyUsername types.String `tfsdk:"body_username" json:"username"`
	EmailId      types.String `tfsdk:"email_id" json:"emailId"`
	Enabled      types.Bool   `tfsdk:"enabled"`
	FullName     types.String `tfsdk:"full_name" json:"fullName"`
	Groups       types.List   `tfsdk:"groups"`
	Password     types.String `tfsdk:"password"`
	Username     types.String `tfsdk:"username"`
}

// NewEditUserDetailsAction returns a new instance of the generated action.
func NewEditUserDetailsAction() action.Action {
	return &EditUserDetailsAction{}
}

// Metadata returns the action type name.
func (r *EditUserDetailsAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_edit_user_details"
}

// Schema returns the action schema.
func (r *EditUserDetailsAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Edit User details", Attributes: map[string]schema.Attribute{"body_username": schema.StringAttribute{MarkdownDescription: "username", Required: true}, "email_id": schema.StringAttribute{MarkdownDescription: "email ID", Required: true}, "enabled": schema.BoolAttribute{Optional: true}, "full_name": schema.StringAttribute{MarkdownDescription: "user's full name", Optional: true}, "groups": schema.ListAttribute{Optional: true, ElementType: types.StringType}, "password": schema.StringAttribute{MarkdownDescription: "password", Required: true}, "username": schema.StringAttribute{MarkdownDescription: "Target Username", Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *EditUserDetailsAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config EditUserDetailsActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *EditUserDetailsAction) invokeRemote(ctx context.Context, config *EditUserDetailsActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/user"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_edit_user_details", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_edit_user_details", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_edit_user_details", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("username", config.BodyUsername.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_edit_user_details", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_edit_user_details", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_edit_user_details", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_edit_user_details", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_edit_user_details", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_edit_user_details", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_edit_user_details", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_edit_user_details", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_edit_user_details", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_edit_user_details", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *EditUserDetailsAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
