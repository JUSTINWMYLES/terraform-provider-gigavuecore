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
var _ action.Action = (*PatchIpInterfaceSolutionAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*PatchIpInterfaceSolutionAction)(nil)

// PatchIpInterfaceSolutionAction is the generated Terraform action implementation.
type PatchIpInterfaceSolutionAction struct {
	client *client.Client
}

// PatchIpInterfaceSolutionActionModel describes the action configuration shape.
type PatchIpInterfaceSolutionActionModel struct {
	IpInterfaceConfigs types.List `tfsdk:"ip_interface_configs" json:"ipInterfaceConfigs"`
	Tags               types.List `tfsdk:"tags"`
}

// NewPatchIpInterfaceSolutionAction returns a new instance of the generated action.
func NewPatchIpInterfaceSolutionAction() action.Action {
	return &PatchIpInterfaceSolutionAction{}
}

// Metadata returns the action type name.
func (r *PatchIpInterfaceSolutionAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_patch_ip_interface_solution"
}

// Schema returns the action schema.
func (r *PatchIpInterfaceSolutionAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Modify existing Ip Interface solution configuration", Attributes: map[string]schema.Attribute{"ip_interface_configs": schema.ListNestedAttribute{Required: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Alias of the ip interface solution", Required: true}, "applications": schema.ListAttribute{MarkdownDescription: "GigaSMART applications for which the ipInterface is used", Required: true, ElementType: types.StringType}, "cluster_name": schema.StringAttribute{MarkdownDescription: "clusterId", Required: true}, "config_status": schema.StringAttribute{Optional: true}, "config_status_reasons": schema.StringAttribute{Optional: true}, "gateway": schema.StringAttribute{MarkdownDescription: "gateway ipv4 or ipv6 address", Required: true}, "interfaces": schema.ListAttribute{MarkdownDescription: "network ports ,tool ports or circuit ports", Required: true, ElementType: types.StringType}, "ip_address": schema.StringAttribute{MarkdownDescription: "ipv4/ipv6 address", Required: true}, "ip_mask": schema.StringAttribute{MarkdownDescription: "ipAddress netmask required with ipAddress", Required: true}, "managed_status": schema.StringAttribute{MarkdownDescription: "managed status of the ip interface configuration. ACTIVE if the cluster is managed in this FM , else it will be marked as INACTIVE", Optional: true}, "mtu": schema.Int64Attribute{Optional: true}, "ref_count": schema.Int64Attribute{MarkdownDescription: "number of control or user nodes using this ip interface configuration", Optional: true}}}}, "tags": schema.ListNestedAttribute{MarkdownDescription: "RBAC Tags", Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"tag_key": schema.StringAttribute{MarkdownDescription: "Name of the tag", Required: true}, "tag_values": schema.ListAttribute{MarkdownDescription: "All possible values of the tag", Required: true, ElementType: types.StringType}}}}}}
}

// Invoke executes the action against the remote API.
func (r *PatchIpInterfaceSolutionAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config PatchIpInterfaceSolutionActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *PatchIpInterfaceSolutionAction) invokeRemote(ctx context.Context, config *PatchIpInterfaceSolutionActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/ipInterfaceConfigs"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_patch_ip_interface_solution", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_patch_ip_interface_solution", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPatch, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_patch_ip_interface_solution", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_patch_ip_interface_solution", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 207) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_patch_ip_interface_solution", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_patch_ip_interface_solution", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_patch_ip_interface_solution", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_patch_ip_interface_solution", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_patch_ip_interface_solution", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_patch_ip_interface_solution", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_patch_ip_interface_solution", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_patch_ip_interface_solution", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_patch_ip_interface_solution", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *PatchIpInterfaceSolutionAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
