package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*AddDevicesToDomainAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*AddDevicesToDomainAction)(nil)

// AddDevicesToDomainAction is the generated Terraform action implementation.
type AddDevicesToDomainAction struct {
	client *client.Client
}

// AddDevicesToDomainActionModel describes the action configuration shape.
type AddDevicesToDomainActionModel struct {
	Async        types.Bool `tfsdk:"async"`
	NodeAddSpecs types.List `tfsdk:"node_add_specs" json:"nodeAddSpecs"`
}

// NewAddDevicesToDomainAction returns a new instance of the generated action.
func NewAddDevicesToDomainAction() action.Action {
	return &AddDevicesToDomainAction{}
}

// Metadata returns the action type name.
func (r *AddDevicesToDomainAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_add_devices_to_domain"
}

// Schema returns the action schema.
func (r *AddDevicesToDomainAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "add clusters/devices to the FM-managed domain", Attributes: map[string]schema.Attribute{"async": schema.BoolAttribute{MarkdownDescription: "if provided, the call returns immediately with the [202 Accepted] HTTP status code, and the discovery process will complete in the background", Optional: true}, "node_add_specs": schema.ListNestedAttribute{Required: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"https_port": schema.StringAttribute{MarkdownDescription: "The default value used is 443. If changed in the device same should be given here", Optional: true}, "node_address": schema.StringAttribute{Required: true}, "password": schema.StringAttribute{Optional: true}, "snmp_version": schema.StringAttribute{MarkdownDescription: "SNMP version to use.", Optional: true}, "username": schema.StringAttribute{Optional: true}}}}}}
}

// Invoke executes the action against the remote API.
func (r *AddDevicesToDomainAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config AddDevicesToDomainActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *AddDevicesToDomainAction) invokeRemote(ctx context.Context, config *AddDevicesToDomainActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/nodes"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_add_devices_to_domain", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_add_devices_to_domain", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_add_devices_to_domain", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.Async.IsNull() {
		query.Set("async", strconv.FormatBool(config.Async.ValueBool()))
	}
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_add_devices_to_domain", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_devices_to_domain", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_devices_to_domain", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_devices_to_domain", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_devices_to_domain", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_devices_to_domain", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_devices_to_domain", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_devices_to_domain", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_add_devices_to_domain", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_add_devices_to_domain", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *AddDevicesToDomainAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
