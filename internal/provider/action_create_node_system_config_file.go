package provider

import (
	"context"
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
var _ action.Action = (*CreateNodeSystemConfigFileAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*CreateNodeSystemConfigFileAction)(nil)

// CreateNodeSystemConfigFileAction is the generated Terraform action implementation.
type CreateNodeSystemConfigFileAction struct {
	client *client.Client
}

// CreateNodeSystemConfigFileActionModel describes the action configuration shape.
type CreateNodeSystemConfigFileActionModel struct {
	ClusterId types.String `tfsdk:"cluster_id"`
	Filename  types.String `tfsdk:"filename"`
}

// NewCreateNodeSystemConfigFileAction returns a new instance of the generated action.
func NewCreateNodeSystemConfigFileAction() action.Action {
	return &CreateNodeSystemConfigFileAction{}
}

// Metadata returns the action type name.
func (r *CreateNodeSystemConfigFileAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_create_node_system_config_file"
}

// Schema returns the action schema.
func (r *CreateNodeSystemConfigFileAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Save the running configuration", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "filename": schema.StringAttribute{MarkdownDescription: "Save the configuration to a new file and make it active, if filename specified", Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *CreateNodeSystemConfigFileAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config CreateNodeSystemConfigFileActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *CreateNodeSystemConfigFileAction) invokeRemote(ctx context.Context, config *CreateNodeSystemConfigFileActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/system/config"
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_create_node_system_config_file", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.Filename.IsNull() {
		query.Set("filename", config.Filename.ValueString())
	}
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_create_node_system_config_file", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_create_node_system_config_file", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_create_node_system_config_file", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_create_node_system_config_file", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_create_node_system_config_file", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_create_node_system_config_file", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_create_node_system_config_file", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_create_node_system_config_file", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_create_node_system_config_file", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_create_node_system_config_file", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *CreateNodeSystemConfigFileAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
