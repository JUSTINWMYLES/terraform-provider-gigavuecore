package provider

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*DeleteNtpAuthKeyAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*DeleteNtpAuthKeyAction)(nil)

// DeleteNtpAuthKeyAction is the generated Terraform action implementation.
type DeleteNtpAuthKeyAction struct {
	client *client.Client
}

// DeleteNtpAuthKeyActionModel describes the action configuration shape.
type DeleteNtpAuthKeyActionModel struct {
	AuthKeyId types.Int64  `tfsdk:"auth_key_id"`
	ClusterId types.String `tfsdk:"cluster_id"`
}

// NewDeleteNtpAuthKeyAction returns a new instance of the generated action.
func NewDeleteNtpAuthKeyAction() action.Action {
	return &DeleteNtpAuthKeyAction{}
}

// Metadata returns the action type name.
func (r *DeleteNtpAuthKeyAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_delete_ntp_auth_key"
}

// Schema returns the action schema.
func (r *DeleteNtpAuthKeyAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Delete NTP Auth Key by alias", Attributes: map[string]schema.Attribute{"auth_key_id": schema.Int64Attribute{MarkdownDescription: "NTP Authentication Key Id", Required: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *DeleteNtpAuthKeyAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config DeleteNtpAuthKeyActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *DeleteNtpAuthKeyAction) invokeRemote(ctx context.Context, config *DeleteNtpAuthKeyActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/system/time/ntp/authKeys/{authKeyId}"
	reqPath = strings.ReplaceAll(reqPath, "{authKeyId}", url.PathEscape(strconv.FormatInt(config.AuthKeyId.ValueInt64(), 10)))
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_delete_ntp_auth_key", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_delete_ntp_auth_key", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 204) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_ntp_auth_key", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_ntp_auth_key", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_ntp_auth_key", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_ntp_auth_key", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_ntp_auth_key", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_ntp_auth_key", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_ntp_auth_key", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_delete_ntp_auth_key", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_ntp_auth_key", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *DeleteNtpAuthKeyAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
