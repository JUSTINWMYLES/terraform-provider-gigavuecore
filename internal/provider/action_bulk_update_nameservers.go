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
var _ action.Action = (*BulkUpdateNameserversAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*BulkUpdateNameserversAction)(nil)

// BulkUpdateNameserversAction is the generated Terraform action implementation.
type BulkUpdateNameserversAction struct {
	client *client.Client
}

// BulkUpdateNameserversActionModel describes the action configuration shape.
type BulkUpdateNameserversActionModel struct {
	BodyInterfaceName types.String `tfsdk:"body_interface_name" json:"interfaceName"`
	InterfaceName     types.String `tfsdk:"interface_name"`
	Servers           types.List   `tfsdk:"servers"`
}

// NewBulkUpdateNameserversAction returns a new instance of the generated action.
func NewBulkUpdateNameserversAction() action.Action {
	return &BulkUpdateNameserversAction{}
}

// Metadata returns the action type name.
func (r *BulkUpdateNameserversAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_bulk_update_nameservers"
}

// Schema returns the action schema.
func (r *BulkUpdateNameserversAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "update name servers", Attributes: map[string]schema.Attribute{"body_interface_name": schema.StringAttribute{MarkdownDescription: "Interface in which the nameserver resides", Required: true}, "interface_name": schema.StringAttribute{MarkdownDescription: "interfaceName for which name servers are updated", Required: true}, "servers": schema.ListAttribute{MarkdownDescription: "List of name server address", Required: true, ElementType: types.StringType}}}
}

// Invoke executes the action against the remote API.
func (r *BulkUpdateNameserversAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config BulkUpdateNameserversActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *BulkUpdateNameserversAction) invokeRemote(ctx context.Context, config *BulkUpdateNameserversActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/sys/ipresolver/nameserver/{interfaceName}"
	reqPath = strings.ReplaceAll(reqPath, "{interfaceName}", url.PathEscape(config.InterfaceName.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_bulk_update_nameservers", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_bulk_update_nameservers", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_bulk_update_nameservers", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_bulk_update_nameservers", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_bulk_update_nameservers", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_bulk_update_nameservers", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_bulk_update_nameservers", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_bulk_update_nameservers", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_bulk_update_nameservers", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_bulk_update_nameservers", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_bulk_update_nameservers", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_bulk_update_nameservers", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_bulk_update_nameservers", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *BulkUpdateNameserversAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
