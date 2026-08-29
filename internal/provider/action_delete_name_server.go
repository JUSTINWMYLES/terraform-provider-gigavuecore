package provider

import (
	"context"
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
var _ action.Action = (*DeleteNameServerAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*DeleteNameServerAction)(nil)

// DeleteNameServerAction is the generated Terraform action implementation.
type DeleteNameServerAction struct {
	client *client.Client
}

// DeleteNameServerActionModel describes the action configuration shape.
type DeleteNameServerActionModel struct {
	ServerIp types.String `tfsdk:"server_ip"`
}

// NewDeleteNameServerAction returns a new instance of the generated action.
func NewDeleteNameServerAction() action.Action {
	return &DeleteNameServerAction{}
}

// Metadata returns the action type name.
func (r *DeleteNameServerAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_delete_name_server"
}

// Schema returns the action schema.
func (r *DeleteNameServerAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Delete name server by address", Attributes: map[string]schema.Attribute{"server_ip": schema.StringAttribute{MarkdownDescription: "address of the nameserver to be deleted", Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *DeleteNameServerAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config DeleteNameServerActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *DeleteNameServerAction) invokeRemote(ctx context.Context, config *DeleteNameServerActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/sys/ipresolver/nameserver/{serverIp}"
	reqPath = strings.ReplaceAll(reqPath, "{serverIp}", url.PathEscape(config.ServerIp.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_delete_name_server", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_delete_name_server", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 204) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_name_server", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_name_server", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_name_server", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_name_server", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_name_server", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_name_server", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_name_server", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_delete_name_server", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_name_server", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *DeleteNameServerAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
